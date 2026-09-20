package main

import (
	"fmt"
	"log"

	"github.com/supernurture/go-fundamentals/internal/embed"
	ex "github.com/supernurture/go-fundamentals/internal/exercises"

	"github.com/rs/zerolog"
	zg "github.com/rs/zerolog/log"
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

	fmt.Println("===================")

	fmt.Println(ex.FizzBuzz(3))
	fmt.Println(ex.FizzBuzz(5))
	fmt.Println(ex.FizzBuzz(15))

	fmt.Println("===================")

	fmt.Println(ex.TwoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(ex.TwoSum([]int{3, 2, 4}, 6))
	fmt.Println(ex.TwoSum([]int{3, 3}, 6))
	fmt.Println(ex.TwoSum([]int{-1, -2, -3, -4, -5}, -8))

	fmt.Println("===================")

	helloChars := []string{"h", "e", "l", "l", "o"}
	ex.ReverseString(helloChars)
	fmt.Println(helloChars)

	nameChars := []string{"H", "a", "n", "n", "a", "h"}
	ex.ReverseString(nameChars)
	fmt.Println(nameChars)

	char := []string{"a"}
	ex.ReverseString(char)
	fmt.Println(char)

	fmt.Println("===================")

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zg.Print("zerolog test")
	zg.Debug().Str("Name", "Steve").Int("Age", 28).Msg("Information")
}

// go build 					- go-fundamentals.exe
// go build -o app main.go 		- app
// go build -o app.exe main.go 	- app.exe
// go build main.go 			- main.exe

// -o stands for output. Name of the executable.
