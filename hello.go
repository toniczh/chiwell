package main

import (
	"fmt"
	"os"
)

func hello() {
	fmt.Printf("Hello")
	card := newCard()
	fmt.Println(card)
	cards := []string{"10", "11"}
	fmt.Println((cards))
	cards = append(cards, "12")
	fmt.Println((cards))
	for i, c := range cards {
		fmt.Println(i, c)
	}
	err := formattedError(0, 0)
	fmt.Println(err)
	fmt.Println(os.Args[1])
	const (
		Zero = 0
		One  = 1
	)
	fmt.Println(Zero)
	var f float64 = 12.13
	//Refer to pointer and get pointer's valua
	fmt.Println("Memory address of f: ", *(&f))

}
func newCard() string {
	return "Five of Diamond"
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d, and b %d, UserID: %d", a, b, os.Getuid())
	}
	return nil
}
