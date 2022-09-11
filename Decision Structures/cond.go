package main

import "log"

func main() {
	var isTrue bool
	isTrue = true

	if isTrue {

		log.Println("isTrue is ", isTrue)

	}else {
		log.Panicln("isTrue is", isTrue)
	}

}