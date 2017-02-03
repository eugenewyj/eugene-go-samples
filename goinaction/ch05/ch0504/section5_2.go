// Sample program to show how polymorphic behavior with interfaces.
package main

import "fmt"

// notifier2 is an interface that defines notification
// type behavior.
type notifier52 interface {
	notify52()
}

// user1 define a user in the program.
type user52 struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user52) notify52() {
	fmt.Printf("Send user email to %s<%s>\n", u.name, u.email)
}

// admin2 define a admin in the program
type admin52 struct {
	user52
	level string
}

// main is the entry point for the application.
func main() {
	ad := admin52{
		user52: user52{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "supper",
	}

	sendNotification52(&ad)
}
func sendNotification52(n notifier52) {
	n.notify52()
}
