package main

import "fmt"

func mainsymplemap() {
	element := make(map[string]string)
	element["H"] = "Hydrogen"
	element["He"] = "Helium"
	element["Li"] = "Lithium"
	element["Be"] = "Beryllium"
	element["B"] = "Boron"
	element["C"] = "Carbon"
	element["N"] = "Nitrogen"
	element["O"] = "Oxygen"
	element["F"] = "Fluorine"
	element["Ne"] = "Neon"
	element["Na"] = "Sodium"
	element["Mg"] = "Magnesium"

	nome, ok := element["Mg"]
	fmt.Println(nome, ok)
}

func mainfullmap() {
	element := map[string]map[string]string{

		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gás",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gás",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "gás",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "gás",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "gás",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "gás",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gás",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gás",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gás",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gás",
		},
		"Na": map[string]string{
			"name":  "Sodium",
			"state": "gás",
		},
		"Mg": map[string]string{
			"name":  "Magnesium",
			"state": "gás",
		},
	}

	if el, ok := element["Li"]; ok {

		fmt.Println(el["name"], el["state"])
	}
}

func main() {
	mainfullmap()
}
