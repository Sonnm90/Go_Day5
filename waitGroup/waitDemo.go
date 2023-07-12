package waitGroup

import (
	"fmt"
	"sync"
)

func printNumber(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func printChar(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+6; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
func Demo() {
	//B1: Tạo biến có kiểu syncWait
	var wg sync.WaitGroup
	//B2:wg.Add( số lượng goroutine cần đợi)
	wg.Add(2)
	//B3: ở mỗi goroutine gọi method  wg.Done() trước khi return
	go printNumber(&wg)
	go printChar(&wg)
	//B4: gọi method wg.Wait
	wg.Wait()
	fmt.Println("Finished")
}
