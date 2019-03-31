package main

import (
	"os"
	"testing"
)

/*IN newDeck if we want to test we could like check :
1. has the four items we based it on
so should have main 4 cards here in the test
2.first is ace of spade and 3.last is like four of space */

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}
	if d[0] != "Spades of Ace" {
		t.Errorf("Got wrong card %v", d[0])
	}
	//d[-1] is invalid syntax
	if d[len(d)-1] != "Clubs of Four" {
		t.Errorf("Got wrong card %v", d[len(d)-1])
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Wrong number of cards:%v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
