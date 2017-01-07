// Sample program to show how polymorphic behavior with interfaces.
package main

import "fmt"

// notifier2 is an interface that defines notification
// type behavior.
type notifier3 interface {
	notify()
}

// user1 define a user in the program.
type user3 struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user3) notify() {
	fmt.Printf("Send user email to %s<%s>\n", u.name, u.email)
}

// admin2 define a admin in the program
type admin3 struct {
	user3
	level string
}

// main is the entry point for the application.
func main() {
	ad := admin3{
		user3: user3{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "supper",
	}

	ad.user3.notify()

	ad.notify()
}
