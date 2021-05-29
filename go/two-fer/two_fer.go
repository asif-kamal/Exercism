// Package twofer prints out you in the phrase if there is no name arugment given.
package twofer

import "fmt"

// ShareWith is a function that either prints out "one for you, and one for me" or replaces you with given name argument.
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
