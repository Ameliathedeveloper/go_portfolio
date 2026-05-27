package lasagnamaster

import "fmt"

// Step 1: work out how much prep time is needed
// if average prep time is not provided, assume each layer takes 2 minutes to prepare
func PreparationTime(layers []string, averagePrepTime int) int {
	// if average prep time is not provided
	if averagePrepTime == 0 {
		//average prep time is 2 minutes per layer
		averagePrepTime = len(layers) * 2
	} else {
		// if average prep time is provided
		// average prep time is the number of layers multiplied by the average prep time
		averagePrepTime = len(layers) * averagePrepTime
	}
	return averagePrepTime
}

func Quantities(layers []string) (noodles int, sauce float64) {
	// for each noodle layer in the lasagna, you need 50 grams of noodles
	// for each sauce layer in the lasagna, you need 0.2 liters of sauce
	for _, layer := range layers {
		switch layer {
		// if the layer is noodles, add 50 grams to the total quantity
		case "noodles":
			noodles += 50
		// if the layer is sauce, add 0.2 liters to the total quantity
		case "sauce":
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendList, myList []string) {
	fmt.Println("read the last element of myList:", myList[len(myList)-1])
	fmt.Println("read the last element of friendList:", friendList[len(friendList)-1])
	//replace the last element of myList with the last element of friendList
	myList[len(myList)-1] = friendList[len(friendList)-1]
	fmt.Println("myList after replacing the last element:", myList)
}

func ScaleRecipe(amountNeededForTwoPortions []float64, numOfPortions int) []float64 {
	// Create a new slice to avoid modifying the original
	scaled := make([]float64, len(amountNeededForTwoPortions))

	// Scaling factor: how many times bigger than 2 portions
	factor := float64(numOfPortions) / 2.0

	// Scale each ingredient
	for i, amount := range amountNeededForTwoPortions {
		scaled[i] = amount * factor
	}

	return scaled
}
