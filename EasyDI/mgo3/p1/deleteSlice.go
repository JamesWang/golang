package mgo3

import (
	"fmt"
	"strconv"
)

type DeleteMethod int16

const (
	Method_1 DeleteMethod = iota
	Method_2
)

func DoDeleteSlice(aSlice []int, args []string, deleteMethod DeleteMethod) {
	if len(args) == 0 {
		fmt.Println("Need an integer value.")
		return
	}
	index, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Deleting item at index:", index)
	DeleteSlice(aSlice, index, deleteMethod)
	//DeleteSlice(aSlice, index, Method_1)
}
func DeleteSlice(aSlice []int, index int, deleteMethod DeleteMethod) {
	fmt.Println("Original          slice:", aSlice)
	length := len(aSlice)
	if index > length-1 {
		fmt.Println("Cannot delete element:", index)
		return
	}
	if deleteMethod == Method_1 {
		aSlice = append(aSlice[:index], aSlice[index+1:]...)
		fmt.Println("After method_1 deletion:", aSlice)
	} else {
		aSlice[index] = aSlice[length-1]
		aSlice = aSlice[:length-1]
		fmt.Println("After method_2 deletion:", aSlice)
	}

}
