package main

import (
	"fmt"

	"github.com/supernurture/go-fundamentals/internal/embed"
)

func main() {
	fmt.Println(string(embed.Note))
	fmt.Println(embed.HW)
}

// go build 					- go-fundamentals.exe
// go build -o app main.go 		- app
// go build -o app.exe main.go 	- app.exe
// go build main.go 			- main.exe

// -o stands for output. Name of the executable.
