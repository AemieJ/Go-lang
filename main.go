package main

// func name and package name should be same

//go is a static type lang so you def the variable with its data type
func main() {
	//deck is the one which comprises off all slice of string
	/*cards := newDeck()
	deck1, deck2 := deal(cards, 5)
	deck1.print()
	deck2.print()*/

	/*cards := deck{"Ace of diamonds", newCard()} // A slice of string
	cards = append(cards, "Six of spades")
	cards.print()*/

	/*for i, card := range cards {
		fmt.Println(i, card)
	}*/

	/*var card string = "Ace of spades"
	card := newCard() // go figures card contains string
	// use := just for starting initializations
	//Assign values : card = "Five of Diamonds"
	fmt.Println(card)*/

	/*cards := newDeckFromFile("my_cards")
	cards.print()
	cards.saveToFile("my_cards")*/

	cards := newDeck()
	cards.shuffle()
	cards.print()
	//fmt.Println(cards.toString())

}

func newCard() string { // func returns type string
	// in new func you declare it as string to return it
	return "Five of Diamonds"
}

//slice is an array in which the datatype used must be same
