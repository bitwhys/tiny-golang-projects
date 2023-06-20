package main

import (
	"fmt"
)

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

// language represents the language's code
type language string

// phrasebook hold greeting for each supported language
var phrasebook = map[language]string {
	"el": "",
	"en":"",
	"fr":"",
	"he":"",
	"ur":"",
	"vi":"Xin ch"
}

// greet returns a greeting to the world.
func greet(l language) string {
	switch l {
	case "en":
		return "Hello world"
	case "fr":
		return "Bonjour le monde"
	default:
		return ""
	}
}
