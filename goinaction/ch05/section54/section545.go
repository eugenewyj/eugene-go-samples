// Sample program to show how polymorphic behavior with interfaces.
package main

import "fmt"

// notifier2 is an interface that defines notification
// type behavior.
type notifier2 interface {
	notify()
}

// user1 define a user in the program.
type user1 struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user1) notify() {
	fmt.Printf("Send user email to %s<%s>\n", u.name, u.email)
}

// admin2 define a admin in the program
type admin1 struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (a *admin1) notify() {
	fmt.Printf("Send admin email to %s<%s>\n", a.name, a.email)
}

// main is the entry point for the application.
func main() {
	bill := user1{"Bill", "bill@email.com"}
	sendNotification1(&bill)

	lisa := admin1{"Lisa", "lisa@email.com"}
	sendNotification1(&lisa)
}
func sendNotification1(n notifier) {
	n.notify()
}
