package main

import (
	"fmt"
	"strconv"
)

// Паттерн заключается в приведении интерфейса к другому виду при помощи дополнительного класса

// Мок базы данных, все методы принимают строки в качестве входных параметров
type DB struct {
}

func NewDatabase() *DB {
	return &DB{}
}
func (db *DB) Get(id string) ([]string, error) {
	// Выполнение запросов...
	return []string{"sdfdf", "sdffs"}, nil
}
func (db *DB) Add(id string) error {
	// Выполнение запросов...
	return nil
}
func (db *DB) Update(id, val string) error {
	// Выполнение запросов...
	return nil
}



// Структура, преобразующая интерфейс работы. Все методы принимают целые числа в качестве параметра
type workWithDB struct {
	db *DB
}

func NewDB() *workWithDB {
	return &workWithDB{
		db: NewDatabase(),
	}
}

func (w *workWithDB) Get(id int) ([]string, error) {
	// Преобразование введенных параметров
	newId := strconv.Itoa(id)
	// ...
	return w.db.Get(newId)
}
func (w *workWithDB) Add(id int) error {
	// Преобразование введенных параметров
	newId := strconv.Itoa(id)
	// ...
	err := w.db.Add(newId)
	return err
}
func (w *workWithDB) Update(id, val int) error {
	// Преобразование введенных параметров
	newId := strconv.Itoa(id)
	newVal := strconv.Itoa(val)
	// ...
	err := w.db.Update(newId, newVal)
	return err
}

func main() {
	var db = NewDB()
	worker(db)
}

func worker(db *workWithDB) {
	fmt.Println(db.Get(1))
	fmt.Println(db.Add(2))
	fmt.Println(db.Update(2, 3))
}
