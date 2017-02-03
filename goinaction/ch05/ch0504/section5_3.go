// Sample program to show how polymorphic behavior with interfaces.
package main

import "fmt"

// notifier2 is an interface that defines notification
// type behavior.
type notifier53 interface {
	notify53()
}

// user1 define a user in the program.
type user53 struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user53) notify53() {
	fmt.Printf("Send user email to %s<%s>\n", u.name, u.email)
}

// admin2 define a admin in the program
type admin53 struct {
	user53
	level string
}

// notify53 implements a method that can be called via a value of type admin
func (a *admin53) notify53() {
	fmt.Printf("Send admin email to %s<%s>\n", a.name, a.email)
}

// main is the entry point for the application.
func main() {
	ad := admin53{
		user53: user53{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "supper",
	}

	sendNotification53(&ad)

	ad.user53.notify53()

	ad.notify53()
}
func sendNotification53(n notifier53) {
	n.notify53()
}
