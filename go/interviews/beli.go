package interviews

import "fmt"

// list: The ingredients that needs to be bought
func (B *Beli) MulaiBeli(list []string) {
	// A variable in the form of:
	// ingredient_name: [ingredient_price, quantity_to_buy]
	// is the default value to store the ingredient name
	// available, the price of the ingredient and the
	// quantity to buy.
	hashMapper := map[string][2]int{
		"minyak":       {12000, 2},
		"bawang merah": {35000, 1}, // for every 2 unit bought, gets free extra 1 unit minyak
	}

	// Result variable is as the form of:
	// ingredient_name: [total_price, quantity]
	ansHashMap := make(map[string][2]int)

	for _, v := range list {

		// Checks whether the key exists in our hashMap.
		// If it doesn't we are just going to ignore.
		if _, exists := hashMapper[v]; exists {

			// Adds the total price bought and quantity taken to the answer variable.
			ansHashMap[v] = [2]int{ansHashMap[v][0] + hashMapper[v][0]*hashMapper[v][1], ansHashMap[v][1] + hashMapper[v][1]}

			// For the special occasion where the ingredient
			// bought is "bawang merah"	and the quantity is
			// over than 2, then we add the quantity for the
			// ingredient "minyak" as a promo.
			if v == "bawang merah" {
				if hashMapper[v][1] >= 2 {
					ansHashMap["minyak"] = [2]int{ansHashMap["minyak"][0], ansHashMap["minyak"][1] + 1*hashMapper[v][1]/2}
				}
			}
		} else {
			continue
		}
	}

	// Variable to hold the total price of the
	// ingredients bought.
	var totalPrice int

	// Display the result in the form of:
	// "Barang minya dibeli dengan harga total Rp. 12000 dan banyak 4 unit"
	for key, value := range ansHashMap {
		totalPrice += value[0]
		fmt.Println("Barang", key, "dibeli dengan harga total Rp.", value[0], "dan banyak", value[1], "unit")
	}
	fmt.Println("Harga total yang dikeluarkan adalah Rp.", totalPrice)
}
