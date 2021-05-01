package main

import (
	"fmt"
	"time"
)

func main() {
	c :=make(chan string)
	people := [5]string{"aaron", "jang", "jenny", "sam", "kate"}
	for _, person := range people {
		go isSexy(person, c) // goroutine has 2 arguments : person & channel
	} // goroutine을 실행// channel을 통해서 결과를 받는것을 기다림
	for i:=0; i <len(people); i++ {
		// channel을 통해서 결과를 받는것을 기다림 (Receiving message through channel : blocking operation)
		s := fmt.Sprintf("Received this message: %s", <-c)
		fmt.Print("Waiting for ", i)
		fmt.Println(s)
	}

}

func isSexy(person string, channel chan string) {
	time.Sleep(time.Second *5)
	channel <- person + " is Sexy"
}

/*
1. goroutine process ends when main function process ends
2. you have to specify the type of data you are going to send&receive via channel
3. sending message : channel <- message
   receiving message : <- channel
*/