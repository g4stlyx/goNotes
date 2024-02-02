package main

import "fmt"

// like arraylists, no need to specify and/or worry about the "array" length

func slices() {
	names := make([]string, 0)

	names = append(names, "name1")
	names = append(names, "name2")

	fmt.Println(names)
	fmt.Println(len(names))
}

func slices2() {
	cities := []string{"Ankara", "İstanbul", "İzmir"}

	// copying arrays/slices
	citiesCopy := make([]string, len(cities))
	copy(citiesCopy, cities)
	citiesCopy = append(citiesCopy, "Adana")

	// filtering
	fmt.Println(citiesCopy[1:3]) // [cities[1] cities[2]]
	fmt.Println(citiesCopy[3:])  // [cities[3] cities[4] ...]
	fmt.Println(citiesCopy[:35]) // [... cities[33] cities[34]]

	fmt.Println(cities)
	fmt.Println(citiesCopy)
}
