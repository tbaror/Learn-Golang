package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var whatToSay string

	whatToSay = "goodbye crule worl"

	fmt.Println(whatToSay)
    
	whatWasSaid, otherthing := saySomthing()

	fmt.Println("The functon says", whatWasSaid, otherthing)

}

func saySomthing() (string, string) {
	return "somthing", "second thing"

	
}