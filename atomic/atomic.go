package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total uint64

func worker(wg *sync.WaitGroup) {
	// wg thông báo hoàn thành khi ra khỏi hàm
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		// lệnh cộng atomic.AddUint64 total được đảm bảo là atomic (đơn nguyên)
		atomic.AddUint64(&total, i)
	}
}
func Demo() {
	// wg được dùng để dừng hàm main đợi các Goroutines khác
	var wg sync.WaitGroup
	// wg cần đợi hai Goroutines gọi lệnh Done() mới thực thi tiếp
	wg.Add(2)
	// tạo Goroutines thứ nhất
	go worker(&wg)
	// tạo Goroutines thứ hai
	go worker(&wg)
	// bắt đầu việc đợi
	wg.Wait()
	// in ra kết quả
	fmt.Println(total)
}
