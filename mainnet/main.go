package main

import (
	"fmt"
	"time"
)

type hdaoThread struct {
	thread_name string
	index       int
	thread_func func()
}

var quit chan int = make(chan int)

func New(name string, num int, function func()) hdaoThread {
	t := hdaoThread{thread_name: name, index: num, thread_func: function}
	return t
}
func (t *hdaoThread) run() {
	t.thread_func()
	quit <- t.index
}
func main() {
	h_t := New("hello", 1, func() {
		fmt.Println("hello world.")
		time.Sleep(60 * time.Second)
	})
	go h_t.run()
	fmt.Println("ending ...")
	a := <-quit
	fmt.Println(a)
}
