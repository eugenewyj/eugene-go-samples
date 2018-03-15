package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s1)

	s2 := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2)

	s2 = s2[1:4]
	fmt.Println(s2)
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2)

	s2 = s2[:2]
	fmt.Println(s2)
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2)

	s2 = s2[1:]
	fmt.Println(s2)
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2)

	var n []int
	fmt.Println(n, len(n), cap(n))
	if n == nil {
		fmt.Println("nil!")
	}
}
