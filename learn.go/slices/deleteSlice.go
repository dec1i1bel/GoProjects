package main

// два способа удаления эл-та из среза

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("1st: need an argument value")
		return
	}

	index := arguments[1]
	i, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Using index", i)
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", aSlice)

	// уд-ть эл-т с инд-м i
	if i > len(aSlice)-1 {
		fmt.Println("Cannot delete element", i)
		return
	}

	// способ 1: логически разделяем исходный срез на два. Два среза разделяются по индексу элемента, который необходимо удалить. После этого мы объединяем эти два среза с помощью операции <...>

	// оп-я <...>автом-ки разв-ет aSlice[i+1:] так, что его эл-ты можно доб-ть к aSlice[:i] посл-но
	aSlice = append(aSlice[:i], aSlice[i+1:]...) // склад-ем новый срез из эл-в до индекса i и после него.
	fmt.Println("After 1st deletion:", aSlice)

	// способ 2: заменяем элемент, который хотим удалить, на последний, используя оператор aSlice[i] = aSlice[len(aSlice)-1], а затем удаляем последний элемент с помощью оператора aSlice = aSlice[:len(aSlice)-1]

	if len(arguments) == 2 {
		fmt.Println("2nd: need an argument value")
		return
	}

	index = arguments[2]
	i, err = strconv.Atoi(index)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Using index", i)
	if i > len(aSlice)-1 {
		fmt.Println("2nd: cannot delete element", i)
		return
	}

	// зам-ть эл-т с инд i на посл-ий
	aSlice[i] = aSlice[len(aSlice)-1]
	// уд-ть посл-ий эл-т - то есть зам-ть сущ-ий срез на него же без последнего эл-та
	aSlice = aSlice[:len(aSlice)-1]
	fmt.Println("After 2nd deletion:", aSlice)
}
