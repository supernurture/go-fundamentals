// bp stands for best practice
// We must name our package using a single word and lowercase letters.
package bp

import "fmt"

// We must follow standard naming conventions when naming our elements and packages.
// Elements (variables, constants, interfaces, structs, and fields within structs).

//
// If we want an element to be accessible by other packages, we must capitalize its first letter.
// However, if we want to make it private, we must use a lowercase letter for the first character.
// For example

// Unexported
var storyName string = "Happy"

func publishStory(name string) {
	fmt.Printf("The Story titled *%s* has been successfully published. \n", name)
}

// Exported
var BookName string = "Grand Theft Auto VI"

func PublishBook(name string) {
	fmt.Printf("The Book titled *%s* has been successfully published. \n", name)
}

// For acronyms, we must use capital letters—for example, HTTP, API, URL, and so on.
// For example

// Unexported
var ftpURL = "..."

// Exported
var HTTPAddr = "..."
