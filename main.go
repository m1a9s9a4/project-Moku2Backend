package main

import (
	"github.com/m1a9s9a4/route"
)

func init() {

}

func main() {
	router := route.Init()
	router.Start(":8080")
}
