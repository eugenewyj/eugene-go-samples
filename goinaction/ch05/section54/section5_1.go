// Sample program to show how polymorphic behavior with interfaces.
package main

import "fmt"

// user1 define a user in the program.
type user51 struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user51) notify51() {
	fmt.Printf("Send user email to %s<%s>\n", u.name, u.email)
}

// admin2 define a admin in the program
type admin51 struct {
	user51
	level string
}

// main is the entry point for the application.
func main() {
	ad := admin51{
		user51: user51{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "supper",
	}

	ad.user51.notify51()

	ad.notify51()
}
