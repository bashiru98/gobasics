package main

import "fmt"

func main() {
	cards := deck{newCard(), newCard(), "ace of diamonds"}

	for i, card := range cards {

		fmt.Println(i, card)
	}

}

func newCard() string {
	return "nice cards"
}
