// кастомная реализация методов Len(), Less(), Swap() интерфейса sort.Interface
package main

import (
	"fmt"
	"sort"
)

type S1 struct {
	F1 int
	F2 string
	F3 int
}

// F3.F1 эквивалентно S1.F1, так как F3 является структурой S1
type S2 struct {
	F1 int
	F2 string
	F3 S1
}

// нужно использовать срез, поскольку все операции сортировки работают со срезами. Именно для этого среза, который должен представлять собой новый тип данных, в этом случае S2slice, мы и будем реализовывать три метода типа sort.Interface
type S2slice []S2

// реализация Len() для типа данных S2slice
func (a S2slice) Len() int {
	return len(a)
}

// в sort.interface() Less() сообщает, должен ли элемент с индексом i сортироваться перед элементом с индексом j. здесь мы переопределяем это поведение
// реализация Less() для типа данных S2slice. метод определяет способ сортировки элементов. В данном случае — с помощью поля встроенной структуры (F3.F1)
func (a S2slice) Less(i, j int) bool {
	return a[i].F3.F1 < a[j].F3.F1
}

// реализация метода типа Swap(), который определяет способ обмена элементов среза во время сортировки
func (a S2slice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	data := []S2{
		S2{1, "One", S1{1, "S1_1", 10}},
		S2{2, "Two", S1{2, "S1_1", 20}},
		S2{-1, "Two", S1{-1, "S1_1", -20}},
	}
	fmt.Println("Before:", data)
	sort.Sort(S2slice(data))
	fmt.Println("After:", data)

	// обратная сортировка работает автоматически
	sort.Sort(sort.Reverse(S2slice(data)))
	fmt.Println("Reverse:", data)
}
