package main

import (
	server "fanServer"
	"fanServer/pkg"
)


func main(){
	srv := new(server.Server)
	handler := pkg.NewHandler()
	srv.Run("80", handler.InitRoutes())
}
