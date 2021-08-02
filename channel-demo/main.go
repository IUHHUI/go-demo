package main

import (
	"fmt"
	"sync"
)

func pCat() {
	fmt.Printf("cat -> ")
}
func pDog() {
	fmt.Printf("dog -> ")
}
func pFish() {
	fmt.Printf("fish -> ")
}
func main() {
	//无缓存的 channel, 必须先有接收者,不然会死锁
	//有缓存的 channel,无此限制
	catCh := make(chan struct{}, 1)
	dogCh := make(chan struct{}, 1)
	fishCh := make(chan struct{}, 1)

	var wg = sync.WaitGroup{}
	wg.Add(3)
	catCh <- struct{}{}

	loop := 3

	go func() {
		for i := 0; i < loop; i++ {
			<-catCh
			pCat()
			dogCh <- struct{}{}
		}
		wg.Done()
		return
	}()

	go func() {
		for i := 0; i < loop; i++ {
			<-dogCh
			pDog()
			fishCh <- struct{}{}
		}
		wg.Done()
		return
	}()

	go func() {
		for i := 0; i < loop; i++ {
			<-fishCh
			pFish()
			catCh <- struct{}{}
		}
		wg.Done()
		return
	}()

	wg.Wait()
	fmt.Println("\nprint finish")

	//cat -> dog -> fish -> cat -> dog -> fish -> cat -> dog -> fish ->
}
