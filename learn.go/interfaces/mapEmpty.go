package main

import "fmt"

// значение по умолчанию при отсутствии пользовательского ввода
var JSONrecord = `{
	Flag": true,
    "Array": ["a","b","c"],
    "Entity": {
      "a1": "b1",
      "a2": "b2",
      "Value": -456,
      "Null": null
    },

    "Message": "Hello Go!"
}`

func typeSwitch(m map[string]interface{}) {
	for k, v := range m {
		switch c := v.(type) {
		case string:
			fmt.Println("Is a string", k, c)
		case float64:
			fmt.Println("Is a float64", k, c)
		case bool:
			fmt.Println("Is a boolean", k, c)
		case map[string]interface{}:
			fmt.Println("Is a map", k, c)
			// если в параметре-карте найдено значение-карта, то рекурсивно вызываем эту же функцию и углубляемся в него
			typeSwitch(v.(map[string]interface{}))
		default:
			fmt.Printf("...Is %v: %T\n", k, c)
		}
	}
	return
}

func exploreMap(m map[string]interface{}) {

}
