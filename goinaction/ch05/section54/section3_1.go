// Sample program to show how to use an interface in Go
package main

import "fmt"

// notifier is an interface that defined notification
// type behavior
type notifier31 interface {
	notify31()
}

// user defines a user in the program.
type user31 struct {
	name  string
	email string
}

// notify implements a method with a pointer receiver .
func (u *user31) notify31()  {
	fmt.Printf("Send user email to %s<%s>\n", u.name, u.email)
}

// main is the entry point for the application.
func main() {
	u := user31{"Bill", "bill@email.com"}
	sendNotification31(&u)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification31(n notifier31) {
 	n.notify31()
}
