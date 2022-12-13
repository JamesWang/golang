package regEx

import (
	"fmt"
	"regexp"
)

/* func getSubstring(s string, indices []int) string {
	return string(s[indices[0]:indices[1]])
} */

func Reg() {
	pattern := regexp.MustCompile("K[a-z]{4}| [A-z]oat")
	description := "Kayak. A boat for one person."

	firstMatch := pattern.FindString(description) //FindStringIndex(description)
	allMatches := pattern.FindAllString(description, -1)
	/*
		fmt.Println("First index", firstIndex[0], "-", firstIndex[1], "=",
			getSubstring(description, firstIndex)) */

	fmt.Println("First match:", firstMatch)
	for i, m := range allMatches {
		//fmt.Println("Index", i, "=", idx[0], "-", idx[1], "=", getSubstring(description, idx))
		fmt.Println("Match", i, "=", m)
	}
}

func MySplit() {
	pattern := regexp.MustCompile(" |boat|one")
	description := "Kayak, A boat for one person"
	split := pattern.Split(description, -1)
	for _, s := range split {
		if s != "" {
			fmt.Println("Substring:", s)
		}
	}
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func Display() {
	Kayak := map[string]string{
		"Name":     "Kayak",
		"Category": "Watersports",
		"Price":    "275",
	}

	Printfln("Value: %v", Kayak)
	Printfln("Go syntax: %#v", Kayak)
	Printfln("Type: %T", Kayak)
}
