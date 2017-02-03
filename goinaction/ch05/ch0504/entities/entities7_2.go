package entities

// user7_2 defines a user in the program
type user7_2 struct {
	Name string
	Email string
}

// Admin7_2 defines an admin in the program.
type Admin7_2 struct {
	user7_2    // The embedded type is unexported
	Rights int
}

