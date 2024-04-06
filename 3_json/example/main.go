package main

import "github.com/d-arken/workshop-go/tree/main/3_json/example/router"

func main() {
	r := router.Setup()
	r.Run(":8080")
}
