package main

import (
	"fmt"
	"net/http"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Person

	// 1. Get the data from you in the terminal first
	fmt.Print("Enter your name: ")
	fmt.Scanln(&pers1.name)
	fmt.Print("Enter your job: ")
	fmt.Scanln(&pers1.job)

	// 2. Set up the Web Route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 1. Tell the browser we are sending HTML
		w.Header().Set("Content-Type", "text/html")

		// 2. Use <b> tags to make the name bold
		fmt.Fprintf(w, "<h1>User Profile</h1>")
		fmt.Fprintf(w, "<p><b>Name:</b> %s</p>", pers1.name)
		fmt.Fprintf(w, "<p><b>Job:</b> %s</p>", pers1.job)
	})

	// 3. Start the server
	fmt.Println("Server starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
