package main

func main() {
	cards := newDeck()
	//cards.saveToFile("my_cards")
	cards.shuffe()
	cards.print()
}
