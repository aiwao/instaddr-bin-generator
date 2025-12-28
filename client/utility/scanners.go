package utility

import (
	"common"
	"fmt"
)

func ScanInt(message string) int {
	for {
		var input int
		fmt.Printf("%s > ", message)
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Printf("%s%v%s", common.Red, err, common.Reset)
			continue
		}
		return input
	}
}

func ScanBool(message string) bool {
	for {
		var input bool
		fmt.Printf("%s > ", message)
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Printf("%s%v%s", common.Red, err, common.Reset)
			continue
		}
		return input
	}
}
