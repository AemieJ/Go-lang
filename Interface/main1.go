package main

import "fmt"

/*WITHOUT INTERFACES
func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
	//as you called the same func for different structures wont be supported
}

//AS WE ARE NOT REALLY USING EB AND SB IN THE RECEIVER PART SO BETTER NOT TO USE IT ONLY
//Used without interfaces so like no matter what we dont use the same func even if the paramteres used are totally diff
func (englishBot) getgreeting() string {
	return "Hello"
}

func (spanishBot) getgreeting() string {
	return "Hola"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getgreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getgreeting())
}
*/

//WITH INTERFACE
type bot interface {
	getgreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printgreeting(eb)
	printgreeting(sb)
}

func printgreeting(b bot) {
	fmt.Println(b.getgreeting())
}

func (englishBot) getgreeting() string {
	return "Hello"
}

func (spanishBot) getgreeting() string {
	return "Hola"
}

/*EXPLANATION : in this bot is defined with type interface
and this bot calls the getgreeting so any structure that also calls
the getgreeting() function and hence they are of type bot too
and thus they can also calling the printgreeting(b bot) func
now*/
