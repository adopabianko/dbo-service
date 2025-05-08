package handler

import (
	"net/http"

	"github.com/adopabianko/dbo-service/internal/auth/dto"
	"github.com/adopabianko/dbo-service/pkg/http/response"
	"github.com/adopabianko/dbo-service/pkg/stacktrace"
	"github.com/gin-gonic/gin"
)

// @Tags Auth
// @Summary user Loginl
// @Description API for user Loginl
// @Accept json
// @Produce json
// @Param payload body dto.LoginRequest true "Payload User Login"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 409 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /auth [post]
func (h *handler) Login(ctx *gin.Context) {
	var opt dto.LoginRequest
	if err := ctx.ShouldBindJSON(&opt); err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, http.StatusBadRequest, "Login"),
			"email or password is invalid",
		)
		return
	}

	token, err := h.service.Login(ctx, opt)
	if err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, stacktrace.GetCode(err), "Login"),
			"email or password is invalid",
		)
		return
	}

	response.ResponseSuccess(
		ctx,
		http.StatusOK,
		"login successfully",
		token,
		nil,
	)
}

// @Tags Auth
// @Summary user Register
// @Description API for user Register
// @Accept json
// @Produce json
// @Param payload body dto.RegisterRequest true "Payload User Register"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 409 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/register [post]
func (h *handler) Register(ctx *gin.Context) {
	var opt dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&opt); err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, http.StatusBadRequest, "Register"),
			"email or password is invalid",
		)
		return
	}

	err := h.service.Register(ctx, opt)
	if err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, stacktrace.GetCode(err), "Register"),
			"email or password is invalid",
		)
		return
	}

	response.ResponseSuccess(
		ctx,
		http.StatusCreated,
		"register successfully",
		nil,
		nil,
	)
}
