package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*create a new type of deck
which is a slice of string */

type deck []string

// deck is like a slice under which you can
//set your own func
//type deck  consists of all slice of strings
//In func newDeck we make our deck of cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardVal := []string{"Ace", "Two", "Three", "Four"}
	/*In matlab when you dont want to use a variable you use ~ in
	go you use _ for same purpose */
	for _, cSuit := range cardSuits {
		for _, cVal := range cardVal {
			cards = append(cards, cSuit+" of "+cVal)
		}
	}
	return cards
}
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

/*EXPLANATION OF FUNC PRINT ACTING AS RECEIVER
any variable of type deck gets access to the functiom
print so like card is of type deck we can access print
func  . (d deck) : so here any variable of type deck
gets access to func and the d is like the instance
of the deck we can work on .
So in receiver you choose a single letter abbreviation related
with the type by convention . */

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

/*BYTE SLICE :
So if the string is "Hi there" then the byte slice will consist
of its corresponding ascii code so yea with byte slice you can
get the string */

//IN GO THE WRITEFILE FUNC MAKES USE OF A BYTE SLICE

//Type conversion : []datatype(another var with datatype)

//func to string changes from deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//TO WRITE IN A FILE
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

//TO READ A FILE
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		//Opt 1 : log the error and return call to newDeck()
		//Opt 2 : log the error and quit the prog.
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

//SHUFFLING THE DECK
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	/*we use a unix nano func which will return us the 
	int64 at  a particular time t and thus the seed or the source 
	will change everytime resulting in a diff sequence of random no 
	generated everytime 
	*/
	r := rand.New(source)
	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

/*GO uses the same seed which triggers the random generator so
we end up getting same random no so in our code we like have to
change the seed everytime */
