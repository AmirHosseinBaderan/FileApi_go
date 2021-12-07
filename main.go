package main

import (
	Server "File_api/app"
)

func main() {
	Server.AddRoutes()
	Server.RunServer()
}
