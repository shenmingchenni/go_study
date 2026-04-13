package main

import "fmt"

func main() {
	var maymap map[string]string
	if maymap == nil {
		fmt.Println("maymap为空")
	}

	mymap1 := make(map[string]string, 10)
	mymap1["one"] = "a"
	mymap1["two"] = "b"
	fmt.Println(mymap1)

	mymap2 := make(map[int]string)
	mymap2[1] = "aa"
	mymap2[2] = "bb"
	fmt.Println(mymap2)

	mymap3 := map[string]string{
		"1": "aaa",
		"2": "bbb"}
	fmt.Println(mymap3)
}
