// переключение поведения в зависимости от типа входных данных
package main

import "fmt"

type Secret struct {
	SecretValue string
}

type Entry struct {
	F1 int
	F2 string
	F3 Secret
}

// утверждение типа. поддерживает Secret и Entry. возвращает два значения: базовое и bool. Базовое значение — то, что вам‚ возможно‚ и нужно. Однако именно значение переменной bool сообщает, было ли утверждение типа успешным и, следовательно, можно ли использовать базовое значение
// Проверка того, относится ли переменная к типу int, требует использования нотации aVar.(int), которая возвращает два значения. В случае успеха вернется реальное значение int и true. В противном случае будет возвращено значение false в качестве второго значения, а это значит, утверждение типа не было успешным и реальное значение извлечь не удалось
func Teststruct(x interface{}) {
	switch T := x.(type) {
	case Secret:
		fmt.Println("Secret type")
	case Entry:
		fmt.Println("Entry type")
	default:
		fmt.Println("Not supported type: %T\n", T)
	}
}

func Learn(x interface{}) {
	switch T := x.(type) { // получам тип данных входного параметра
	default:
		fmt.Printf("Data type: %T\n", T)
	}
}

func main() {
	A := Entry{100, "F2", Secret{"my password"}}
	Teststruct(A)
	Teststruct(A.F3)
	Teststruct("a string")
	Learn(12.23)
	Learn('€')
}
