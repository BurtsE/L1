package main

import (
	"fmt"
	"sync"
	"time"
)

// Структура для записи
type myMap struct {
	mp map[interface{}]interface{}
	// Поскольку запись в map не потокобезопасна, необходимо блокировать попытки параллельной записи. Для этого добавляется мьютекс
	mu sync.Mutex
}

func main() {
	var m = myMap{
		mp: make(map[interface{}]interface{}),
		mu: sync.Mutex{},
	}
	// Создание горутин, производящих запись в map.
	for i := 0; i < 10; i++ {
		go m.Set(i, "sdf")
	}
	// Sleep для ожидания результата работы
	time.Sleep(time.Second * 3)
	fmt.Println(m.mp)
}

// Функция записи в словарь
func (m *myMap) Set(key, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mp[key] = value
}
