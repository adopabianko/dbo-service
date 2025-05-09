package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adopabianko/dbo-service/config"
	"github.com/adopabianko/dbo-service/pkg/http/middleware"
	"github.com/adopabianko/dbo-service/pkg/validation"
	"github.com/adopabianko/dbo-service/router"

	authHandler "github.com/adopabianko/dbo-service/internal/auth/handler"
	authRepository "github.com/adopabianko/dbo-service/internal/auth/repository"
	authService "github.com/adopabianko/dbo-service/internal/auth/service"
	customerHandler "github.com/adopabianko/dbo-service/internal/customer/handler"
	customerRepository "github.com/adopabianko/dbo-service/internal/customer/repository"
	customerService "github.com/adopabianko/dbo-service/internal/customer/service"
	orderHandler "github.com/adopabianko/dbo-service/internal/order/handler"
	orderRepository "github.com/adopabianko/dbo-service/internal/order/repository"
	orderService "github.com/adopabianko/dbo-service/internal/order/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func StartHTTPServer() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("failed to load config ", err)
		return
	}

	// load database connection
	db := config.InitPostgresDatabase(cfg)
	defer config.ClosedDB()

	// load server
	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	var (
		// Auth Module
		authRepository authRepository.IAuthRepository = authRepository.NewRepository(db)
		authService    authService.IAuthService       = authService.NewService(authRepository, cfg)
		authHandler    authHandler.IAuthHandler       = authHandler.NewHandler(authService)

		// Customer Module
		customerRepository customerRepository.ICustomerRepository = customerRepository.NewRepository(db)
		customerService    customerService.ICustomerService       = customerService.NewService(customerRepository)
		customerHandler    customerHandler.ICustomerHandler       = customerHandler.NewHandler(customerService)

		// Order Module
		orderRepository orderRepository.IOrderRepository = orderRepository.NewRepository(db)
		orderService    orderService.IOrderService       = orderService.NewService(orderRepository, customerRepository)
		orderHandler    orderHandler.IOrderHandler       = orderHandler.NewHandler(orderService)
	)

	// router
	router.Default(server)
	router.Swagger(server)
	router.Auth(server, authHandler)
	router.Customer(server, customerHandler)
	router.Order(server, orderHandler)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.App.Port),
		Handler: server.Handler(),
	}

	// get instance validator from Gin
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// register custom validator
		v.RegisterValidation("not_empty", validation.NotEmptySlice)
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen:", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")

	log.Println("Server exiting")
}
