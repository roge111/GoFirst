package main

import (
	"fmt"
	"bufio"
	"os"
)
//На вход поступает три строки, вывести в том же порядке
func input()string{
	scanner := bufio.NewScanner(os.Stdin)

	_ = scanner.Scan()
	result:= scanner.Text()
	return result
}




func practic1() {
	str1 := input()
	str2:= input()
	str3 := input()

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
	
}



func practic2(){
	split := input()
	string1 := input()
	string2 := input()
	string3 := input()
	// fmt.Print(string1, split, string2, split, string3) //Вариант решения раз

	result := string1 + split + string2 + split + string3 // Вариант решения 2
	fmt.Println(result)
}

func main(){
	practic1()
	practic2()
}
