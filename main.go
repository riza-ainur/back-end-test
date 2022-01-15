package main

import "fmt"

func main() {
	var value, temp int
	size := 0

	fmt.Print("jumlah array n=... ")
	fmt.Scanln(&size)
	fmt.Println("masukkan value array:... ")

	elements := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&elements[i])
	}

	fmt.Println("value array yang sudah dimasukkan:... ", elements)

	for _, element := range elements {
		if element > temp {
			temp = element
			value = temp
		}
	}

	fmt.Println("nilai terbesarnya:...  ", value)
}
