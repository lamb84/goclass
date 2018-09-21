package main

func main() {
	cards := deckFromFile("test1")
	shuffled := cards.shuffle()
	shuffled.print()
}
