package main

import "fmt"

func modifyCategory() {
	var modCat, newCatName string
	
	fmt.Println("")
	fmt.Println("Modify category")
	fmt.Println("===============")
	fmt.Printf("Which category do you wish to modify? ")
	fmt.Scanln(&modCat)

	indexCat := searchCategory(modCat)
	if indexCat == -1 {
		fmt.Println("Category not found.")
	} else {
		fmt.Printf("You have selected Category: %v. What would you like the new name to be? ", categories[indexCat])
		fmt.Scanln(&newCatName)
		
		categories[indexCat] = newCatName
		fmt.Printf("Category name has been changed to %v.\n", newCatName)
	}
}

func searchCategory(category string) (index int) {
	index = -1
	for i, element := range categories {
		if category == element {
			index = i
		}
	}
	return index
}