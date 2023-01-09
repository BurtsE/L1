package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Счетчик
type Counter struct {
	val int64
	mu  sync.Mutex
}

// Инициализация счетчика
func (c *Counter) Init() {
	c.val = 0
	c.mu = sync.Mutex{}
}

// Инкремент при помощи мьютекса для избежания гонки памяти
func (c *Counter) Increment1() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

// Инкремент при помощи атомарной операции исключает гонки памяти, мьютекс не нужен
func (c *Counter) Increment2() {
	atomic.AddInt64(&c.val, 1)
}
func main() {
	var c Counter
	c.Init()
	wg := new(sync.WaitGroup)
	// 10000 горутин одновременно пытаются увеличить счетчик
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Одновременное использование приведет к ошибке в расчетах
			//c.Increment1()
			c.Increment2()
		}()
	}
	wg.Wait()
	fmt.Println(c.val)
}
