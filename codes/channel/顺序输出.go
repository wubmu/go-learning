package main

import (
	"fmt"
	"sync"
)

func PrintCat(dogChan, catChan chan bool) {
	defer wg.Done()
	defer close(catChan)
	for i := 0; i < 10; i++ {
		<-dogChan
		fmt.Println("cat ...")
		catChan <- true
	}
}

func PrintDog(fishChan, dogChan chan bool) {
	defer wg.Done()
	defer close(dogChan)
	for i := 0; i < 10; i++ {
		<-fishChan
		fmt.Println("dog ...")
		dogChan <- true
	}
}

func PrintFish(catChan, fishChan chan bool) {
	defer wg.Done()
	defer close(fishChan)
	for i := 0; i < 10; i++ {
		<-catChan
		fmt.Println("fish ...")
		fishChan <- true
	}
}

var wg sync.WaitGroup

func main() {
	dogChan := make(chan bool, 1)
	catChan := make(chan bool, 1)
	fishChan := make(chan bool, 1)
	fishChan <- true

	go PrintDog(fishChan, dogChan)
	go PrintFish(catChan, fishChan)
	go PrintCat(dogChan, catChan)

	wg.Add(3)
	wg.Wait()

}
