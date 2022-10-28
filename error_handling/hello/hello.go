package main

import "fmt"
import "log"
import "error_handling/input"

func main() {
	log.SetPrefix("Hello: ")
    log.SetFlags(0)

	message, err := input.Input("");
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message);
}