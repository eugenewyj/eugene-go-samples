package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

type Duration int64

func main() {
	var bill user
	fmt.Println("user bill:", bill)

	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}
	fmt.Println("user lisa:", lisa)
	lisa.notify()

	lisa2 := user{"lisa2", "lisa2@email.com", 321, false}
	fmt.Println("user lisa2:", lisa2)

	fred := admin{
		person: user{
			name:       "lisa",
			email:      "lisa@email.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}
	fmt.Println("fred:", fred)

	var dur Duration
	//dur = int64(1000)
	fmt.Println("dur init:", dur)
	dur = Duration(1000)
	fmt.Println("dur inited:", dur)
}
