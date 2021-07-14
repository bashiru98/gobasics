package main

func main() {
	cards := deck{newCard(), newCard(), "ace of diamonds"}

	cards.print()

}

func newCard() string {
	return "nice cards"
}
