package lesson

import (
	"fmt"
	"sync"
)

// WaitGroupWithUnsafeCounter waitGroup是线程安全的，但如果goroutine中的操作不是原子性的，就会导致预期之外的结果
func WaitGroupWithUnsafeCounter() {
	var wg sync.WaitGroup
	counter := 0
	total := 5000
	for i := 0; i < total; i++ {
		wg.Add(1)
		go func() {
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("WaitGroupWithUnsafeCounter: %v/%v\n" ,counter,total)
}


func WaitGroupWithSafeCounter() {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	total := 5000
	for i := 0; i < total; i++ {
		wg.Add(1)
		go func() {
			defer func() { mut.Unlock() }()
			mut.Lock()
			counter++
			//fmt.Println(counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("WaitGroupWithSafeCounter: %v/%v\n" ,counter,total)
}