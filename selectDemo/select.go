package selectDemo

import (
	"fmt"
	"time"
)

// Đẩy dữ liệu cho channel và chờ 5 giây
func data1(ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "1"
}

// Đẩy dữ liệu cho channel và chờ 3 giây
func data2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "1"
}
func data3(ch chan string) {
	time.Sleep(4 * time.Second)
	ch <- "1"
}

func defaultChannel() {
	fmt.Println("Coming!")
	time.Sleep(5 * time.Second)
	fmt.Println("Default")
}

func Demo() {
	// Tạo các biến channel để truyền giá trị string
	chan1 := make(chan string)
	chan2 := make(chan string)
	chan3 := make(chan string)

	// Gọi các goroutine cùng với các biến channel
	go data1(chan1)
	go data2(chan2)
	go data3(chan3)

	// Cả hai câu lệnh case kiểm tra dữ liệu trong chan1 or chan2.
	// Nhưng dữ liệu không sẵn sàng (Cả 2 routines đều tạm dừng 2 giây và 4 giây)
	// Nên đoạn code trong default sẽ được chạy mà không chờ dữ liệu trong các channel.
	for {
		check := false
		time.Sleep(time.Second)
		select {
		case x := <-chan1:
			fmt.Println(x, "from ch1")
			check = true
		case y := <-chan2:
			fmt.Println(y, "from ch2")
			check = true
		case z := <-chan3:
			fmt.Println(z, "from ch3")
			check = true
		default:
			//defaultChannel()
			fmt.Println("Loading...")
		}
		if check {
			break
		}
	}
}
