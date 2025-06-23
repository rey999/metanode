package main

import (
	"blog/common"
	"blog/config"
	"blog/server"
)

func main() {
	println("Hello, blog!")
	config.LoadConfig()
	common.InitDB()
	server.RunServer()
}
