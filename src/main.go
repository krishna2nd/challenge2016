package main

import (
	"commands"
)

func main() {
	/*regions.GenerateMapFromCSV("cities.csv");
	//fmt.Println(regions.Find([]string {"TAMILNADU", "INDIA" }))
	//fmt.Println(regions.Find([]string {"INDIA"}))
	var distributor = distributor.New("D1", []string{
		"INCLUDE: INDIA",
		"INCLUDE: UNITEDSTATES",
		"EXCLUDE: KARNATAKA-INDIA",
		"EXCLUDE: CHENNAI-TAMILNADU-INDIA",
	})

	fmt.Print(distributor, distributor.GetPermissions())
	*/
	// Interactive shell will provide to user to operate
	// with file data and console process
	commands.NewShell().Process()
}
