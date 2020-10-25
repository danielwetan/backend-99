package routes

import (
	"net/http"

	"github.com/danielwetan/backend-99/soal-3-4/controllers"
)

func Palindrome() {
	http.HandleFunc("/api/palindrome", controllers.Palindrome)
}
