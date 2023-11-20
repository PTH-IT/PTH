package main

import (
	_ "Golang/docs"
	InforLog "Golang/log/infor"
	service "Golang/service"
	"fmt"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email haupham404
// @host 103.195.239.33:1909
func main() {
	InforLog.PrintLog(fmt.Sprintf("af.Run call"))
	service.Run()
}
