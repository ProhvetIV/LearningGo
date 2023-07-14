package main

import "fmt"

// washingmachine struct
type washingmachine struct {
	brand string
	model string
}

// method to get the brand and model, returning a string
func (wm washingmachine) getModel() string {
	return fmt.Sprintf("Your washingmachine is from %s and the model is %s", wm.brand, wm.model)
}

// func to print the getModel() sting in the terminal
func test(info washingmachine) {
	fmt.Println(info.getModel())
	fmt.Println("========================================================================")
}

func main() {
	test(washingmachine{
		brand: "Samsung",
		model: "Eco-Life 205",
	})
	test(washingmachine{
		brand: "Sony",
		model: "The-Fastest 08",
	})
	test(washingmachine{
		brand: "Honda",
		model: "Multi-Purpose C960A",
	})
}
