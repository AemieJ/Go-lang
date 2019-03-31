package main

type user struct{ name string }

//Ways you can write your interfaces func
type bot2 interface {
	getgreeting(string, int) (string, error)
	getbotversion() float64
	respondTouser(user) string
}

//Bot here is an interface type as you cant really get the values for it
//Concrete type like map , struct , int , string etc are one
//to which you can assign the values
