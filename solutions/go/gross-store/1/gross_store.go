package gross

// Units returns the standard unit sizes.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new empty bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// CheckItemExists checks if a unit exists and returns its quantity.
func CheckItemExists(units map[string]int, unit string) (int, bool) {
	itemQty, exists := units[unit]
	return itemQty, exists
}

// AddItem adds an item to the bill using the unit size.
func AddItem(bill, units map[string]int, item, unit string) bool {
	itemQty, exists := CheckItemExists(units, unit)
	if !exists {
		return false
	}

	bill[item] += itemQty
	return true
}

// RemoveItem removes an item quantity from the bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemQty, exists := CheckItemExists(units, unit)
	if !exists {
		return false
	}

	currentQty, itemExists := bill[item]
	if !itemExists {
		return false
	}

	newQty := currentQty - itemQty
	if newQty < 0 {
		return false
	}
	if newQty == 0 {
		delete(bill, item)
		return true
	}

	bill[item] = newQty
	return true
}

// GetItem returns the quantity of an item in the bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, exists := bill[item]
	return qty, exists
}
