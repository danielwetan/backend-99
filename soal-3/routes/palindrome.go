package routes

import (
	"net/http"

	"github.com/danielwetan/backend-99/controllers"
)

func Palindrome() {
	http.HandleFunc("/api/palindrome", controllers.Palindrome)
}
