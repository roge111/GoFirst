package main

import (
	"fmt"
)

func main() {

	var name string
	fmt.Scan(&name) //Демонстрирую ввод одного значения в одну переменную.

	var s1, s2, s3 string
	countWords, error := fmt.Scan(&s1, &s2, &s3)

	fmt.Println(name)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println(countWords, error)
}
