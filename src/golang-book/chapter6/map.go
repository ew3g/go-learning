package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)

	z := make(map[int]int)
	z[1] = 10
	fmt.Println(z[1])

	delete(z, 1)
	fmt.Println(z)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Un"])
	//name, ok := elements["Un"]
	//fmt.Println(name, ok)

	// get the value and if the search was ok
	// if ok == true, run the code
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}

	// Short way to create a map
	shortElements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(shortElements)

	// Storing general information to maps
	generalInfoElements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
	}

	if el, ok := generalInfoElements["H"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	xa := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(xa[2:5])
}
