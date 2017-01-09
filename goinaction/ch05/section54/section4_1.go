// Sample program to show how polymorphic behavior with interfaces.
package main

import "fmt"

// notifier2 is an interface that defines notification
// type behavior.
type notifier41 interface {
	notify41()
}

// user1 define a user in the program.
type user41 struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user41) notify41() {
	fmt.Printf("Send user email to %s<%s>\n", u.name, u.email)
}

// admin2 define a admin in the program
type admin41 struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (a *admin41) notify41() {
	fmt.Printf("Send admin email to %s<%s>\n", a.name, a.email)
}

// main is the entry point for the application.
func main() {
	bill := user41{"Bill", "bill@email.com"}
	sendNotification41(&bill)

	lisa := admin41{"Lisa", "lisa@email.com"}
	sendNotification41(&lisa)
}
func sendNotification41(n notifier41) {
	n.notify41()
}
