// bp stands for best practice
// We must name our package using a single word and lowercase letters.
package bp

import (
	"errors"
	"fmt"
	"strconv"
)

// We must follow standard naming conventions when naming our elements and packages.
// Elements (variables, constants, interfaces, structs, and fields within structs).

//
// If we want an element to be accessible by other packages, we must capitalize its first letter.
// However, if we want to make it private, we must use a lowercase letter for the first character.
// For example

const (
	// Unexported
	pathURL = "..."
	// Exported
	PathURL = "..."
)

// For acronyms, we must use capital letters—for example, HTTP, API, URL, and so on.
var (
	// Unexported
	httpAddr         = "..."
	storyName string = "Story A"

	// Exported
	HTTPAddr        = "..."
	BookName string = "Book A"
)

func publishStory(name string) {
	fmt.Printf("The Story titled *%s* has been successfully published. \n", name)
}

func PublishBook(name string) {
	fmt.Printf("The Book titled *%s* has been successfully published. \n", name)
}

// We're going to check the default behavior of fmt.Errorf.
// We typically use it to add context to an error before passing it on to the next step.
func CheckBehaviorFmtErrorf() error {
	// Should produce an error
	_, err := strconv.Atoi("abc")
	if err != nil {
		// fmt.Errorf can be used to add context to an error.
		return fmt.Errorf("an error occurred while attempting to parse the string into a number: %w", err)
	}
	return nil
}

// We're going to check the default behavior of errors.Join.
// We typically use it to combine independent errors into one.
func CheckBehaviorErrorJoin() error {
	err := errors.New(
		"1st error")
	err2 := errors.New(
		"2nd error")
	return errors.Join(err, err2)
}

// We can create pre-declared global error type as a representation of spesific condition.
// The name of the global error type is sentinel error.
// Standard naming convention of the error is var Err + (condition)
var ErrUserNotFound = errors.New("can't find specified user")

func getUserByID(id int) error {
	fmt.Printf("failed getting user with id %d", id)
	return ErrUserNotFound
}

func srv() {
	err := getUserByID(10)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			// We can return a relevant response to the requester.
			// 404 or something else
			fmt.Println("invalid request")
		}
	}
}
