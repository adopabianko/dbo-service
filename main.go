package main

import (
	"github.com/adopabianko/dbo-service/cmd"
	_ "github.com/adopabianko/dbo-service/docs"
)

// @title DBO Service
// @description DBO Service handles Customer, Order process
// @schemes   http https
// @BasePath  /

// @securityDefinitions.apikey	BearerAuth
// @in 							header
// @name 						Authorization
func main() {
	cmd.Execute()
}
