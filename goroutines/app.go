package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
   done <- true

}

func main() {
    done := make(chan bool)
	greet("Nice to meet you")
	greet("How are you?")
	go slowGreet("How...are...you", done)
	greet("how are you liking the course")
  <-done

}
