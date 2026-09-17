package main

import (
	"fmt"
	"log"

	"github.com/supernurture/go-fundamentals/internal/embed"
)

func main() {
	fmt.Println(string(embed.Note))
	fmt.Println(embed.HW)

	// can be used to migrate SQL scripts.
	content, err := embed.UserScript.ReadFile("00001-auth.sql")
	if err != nil {
		log.Fatalf("unable to read embed: %s", err.Error())
	}
	fmt.Println(string(content))
}

// go build 					- go-fundamentals.exe
// go build -o app main.go 		- app
// go build -o app.exe main.go 	- app.exe
// go build main.go 			- main.exe

// -o stands for output. Name of the executable.
