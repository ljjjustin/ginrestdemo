package main

import (
	"github.com/ljjjustin/ginrestdemo/apis"
)

func main() {
	router := apis.NewRouter()

	router.Run(":8080")
}
