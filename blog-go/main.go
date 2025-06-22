package main

import (
	"blog/common"
	"blog/config"
	"blog/server"
)

func main() {
	println("Hello, blog!")
	common.InitDB()
	config.LoadConfig()
	server.RunServer()
}
