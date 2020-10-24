package main

import (
	"fmt"
	"net/http"
	"github.com/danielwetan/backend-99/routes"
)

func main() {
	routes.Palindrome()

	PORT := ":3000"
	fmt.Println("App running on PORT", PORT[1:])
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
