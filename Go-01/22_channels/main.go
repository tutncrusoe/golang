package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello Gopher")
	done <- true // Gửi dữ liệu vào channel
}

func main() {
	// Khởi tạo channel
	done := make(chan bool)
	go hello(done)

	<-done // Nhận dữ liệu từ channel, chương trình sẽ bị block tại đây cho đến khi nhận được dữ liệu từ channel

	fmt.Println("Main")
}
