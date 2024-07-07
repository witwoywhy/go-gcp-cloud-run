package main

import (
	"go-gcp-cloud-run/httpserv"
	"go-gcp-cloud-run/infrastructure"
)

func init() {
	infrastructure.InitConfig()
}

func main() {
	infrastructure.InitLog()
	infrastructure.InitApp()
	httpserv.Run()
}
