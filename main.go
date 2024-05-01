package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

var target int
var attemps int

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/guess", guessHandler)

	target = rand.Intn(100)
	fmt.Println("Server is running on port 8080")
	fmt.Printf("Target number is %d\n", target)
	http.ListenAndServe(":8080", nil)
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	guess, err := strconv.Atoi(r.FormValue("guess"))
	if err != nil {
		http.Error(w, "Invalid guess", http.StatusBadRequest)
		return
	}

	attemps++
	if guess == target {
		fmt.Fprintf(w, "Congratulations! You have guessed the number in %d attemps", attemps)
		target = rand.Intn(100)
		attemps = 0
		fmt.Printf("Target number is %d\n", target)
	} else if guess < target {
		fmt.Fprintf(w, "Your guess is too low")
	} else {
		fmt.Fprintf(w, "Your guess is too high")
	}
}
