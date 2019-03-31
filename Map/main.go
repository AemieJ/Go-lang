package main

import "fmt"

func main() {
	/*METHOD 3
	colors := make(map[int]string) //creates empty map
	colors[10] = "#0000ff"*/
	/*TO DELETE A KEY
	delete(colors, 10)*/

	//in structs you would do is like structName.white
	/*METHOD 2
	var colors map[string]string*/

	//METHOD 1
	//color says the keys are of type string and here val also are
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Println("Hex code for", key, "is:", val)
	}
}
