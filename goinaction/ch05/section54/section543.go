// Sample program to show how to use an interface in Go
package main

import "fmt"

// notifier is an interface that defined notification
// type behavior
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a pointer receiver .
func (u *user) notify()  {
	fmt.Printf("Send user email to %s<%s>\n", u.name, u.email)
}

// main is the entry point for the application.
func main() {
	u := user{"Bill", "bill@email.com"}
	sendNotification(&u)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
 	n.notify()
}
