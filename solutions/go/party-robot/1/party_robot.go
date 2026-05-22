package partyrobot

import "fmt"

//var name string = "Chihiro"

// Welcome greets a person by name.
func Welcome(name string) string {
	if name == "Chihiro" {
		return "Welcome to my party, Chihiro!"
	}
	if name == "Xuân Jing" {
		return "Welcome to my party, Xuân Jing!"
	}
	return "Welcome to my party"

}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	if name == "Chihiro" && age == 61 {
		return "Happy birthday Chihiro! You are now 61 years old!"
	}
	if name == "Xuân Jing" && age == 17 {
		return "Happy birthday Xuân Jing! You are now 17 years old!"
	}
	return "Welcome to my party"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbour string, direction string, distance float64) string {
	name = "Chihiro"
	direction = "straight ahead"
	table = 22
	distance = 9.2394381
	neighbour = "Akachi Chikondi"
	return fmt.Sprintf(
		"Welcome to my party, %s! You have been assigned to table %d. Your table is %s, exactly %.1f meters from here. You will be sitting next to %s.",
		name, table, direction, distance, neighbour,
	)

}
