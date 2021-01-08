package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 *int = &var1
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(*var2)
}

//【Go】基本文法③(ポインタ・構造体)
// https://qiita.com/k-penguin-sato/items/62dfe0f93f56e4bf9157　より
