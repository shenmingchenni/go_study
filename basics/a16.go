package main

import "fmt"

func printcitymap(citymap map[string]string) {
	for key, value := range citymap {
		fmt.Println("key=", key, "value=", value)
	}
}

func changecitymap(citymap map[string]string) {
	citymap["uk"] = "lundon"
}

func main() {
	citymap := make(map[string]string)
	citymap["china"] = "beijing"
	citymap["tyko"] = "dj"
	citymap["amaric"] = "hsd"

	printcitymap(citymap)

	delete(citymap, "china")

	citymap["tyko"] = "ddjj"
	changecitymap(citymap)

	fmt.Println("......")

	printcitymap(citymap)

}
