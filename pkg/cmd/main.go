package main

import "github.com/kelsonniiz/Service-for-calculating-arithmetic-expressions/pkg/internal/application"

func main() {
	app := application.New()
	app.RunServer()
}
