package utility

import "fmt"

func ScanInt(message string) int {
	for {
		var input int
		fmt.Printf("%s >", message)
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return input
	}
}
