package main

import "fmt"

func main() {
	// len(cities) = cities.length
	var numbers [5]int
	cities:= [4]string{"Istanbul","Paris","London","New York"}

	numbers[2] = 50

	fmt.Println(numbers)

	for i:=0; i<len(cities); i++ {
		fmt.Println(cities[i])
	}

	var matrix [3][3]int
	var temp = 0
	for i:=0; i<len(matrix); i++ {
		for j:=0; j<len(matrix);j++{
			matrix[i][j] = temp
			temp ++
		}
	}

	for i:=0; i<len(matrix); i++ {
		fmt.Print("[")
		for j:=0; j<len(matrix);j++{
			fmt.Print(matrix[i][j])
			if j!= len(matrix)-1{
				fmt.Print(",")
			}
		}
		fmt.Println("]")
	}

}