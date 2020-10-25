package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danielwetan/backend-99/soal-3-4/helpers"
)

func Palindrome(w http.ResponseWriter, r *http.Request) {
	helpers.Headers(&w)

	if r.Method == "POST" {
		r.ParseForm()
		var (
			start, _ = strconv.ParseInt(r.FormValue("start"), 10, 32)
			end, _   = strconv.ParseInt(r.FormValue("end"), 10, 32)
		)

		palindrome := helpers.Palindrome(int(start), int(end))

		body := map[string]int{
			"totalPalindrome": palindrome,
		}
		res := helpers.ResponseMsg(true, body)
		json.NewEncoder(w).Encode(res)
	} else {
		body := "Invalid HTTP method"
		res := helpers.ResponseMsg(false, body)
		json.NewEncoder(w).Encode(res)
	}
}
