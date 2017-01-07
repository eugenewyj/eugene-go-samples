package main

import "fmt"

type user struct {
	name       string
	email      string
}

func (u user) notify() {
	fmt.Printf("Send User Email To %s<%s>\n", u.name, u.email)
}

func (u user) changeEmail1(email string) {
	u.email = email
}

func (u *user) changeEmail2(email string) {
	u.email = email
}

func main() {
	bill1 := user{"bill1", "bill1@emial.com"}
	bill1.notify()

	bill2 := &user{name: "bill2", email: "bill2@email.com"}
	bill2.notify()
	//相当于
	(*bill2).notify()

	bill1.changeEmail1("bill1-1@email.com")
	bill1.notify()

	bill1.changeEmail2("bill1-2@email.com")
	//相当于
	(&bill1).changeEmail2("bill-2-1@email.com")
	bill1.notify()

	bill2.changeEmail1("bill2-1@email.com")
	bill2.notify()

	bill2.changeEmail2("bill2-2@email.com")
	bill2.notify()
}
