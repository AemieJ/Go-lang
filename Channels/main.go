package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://youtube.com",
		"http://golang.org",
	}

	c := make(chan string) //making a new channel

	for _, link := range links {
		//checkLink(l)
		go checkLink(link, c) // initiate go routine
		/*so when it goes into blocking call so the go routine
		passes back to main routine so for the next link ie next
		iteration it goes through the go routine and then it keeps going
		on and hence it doesnt wait for a error message in the blocking call
		so it calls all the links at the blocking call .
		*/
	}
	fmt.Println(<-c) //Here , the channel get the string and then prints immediately

	for l := range c { //infinite loop
		//go checkLink(l, c) //the <-c receives links and calls the go routine
		go func(link string) { //function literal
			time.Sleep(3 * time.Second)
			checkLink(link, c)
		}(l) // calling the function
	}
	/*
		The proble with defining the sleep function in the main routine is that
		lets say golang.org give the first resp and the main routine will sleep for 5
		seconds or whatever duration now there is a chance lets say google.com
		gives a response within the sleep duration but main routine wont excute it*/

}

/*So the problem of just writing go Checklink(link) is that
it will not print only the only reason is that when the main routine
is initiated it creates child go routine for each of the segment to run
simultaenously however when it does the looping the main routine just shut down
whether or not all the child go routine has completed the tasks
Hence you use channels .
*/

func checkLink(link string, c chan string) {

	_, err := http.Get(link) // this is called the blocking call
	if err != nil {
		fmt.Println(link, " might be down.")
		c <- link
		return //to write so we dont do anything else
	}
	fmt.Println(link, " is working.")
	c <- link

}

//So we go ask for request and wait for the response for every link
//which is not convinient for many links
// we use a go routine so we do all website links all simultaneously
