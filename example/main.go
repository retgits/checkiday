package main

import (
	"fmt"

	"github.com/retgits/checkiday"
)

func main() {
	days, err := checkiday.Today()
	if err != nil {
		fmt.Printf("Oh noes, an error occured: %s", err.Error())
	}

	for idx := range days.Holidays {
		fmt.Printf("Today is %s\n", days.Holidays[idx].Name)
	}

	days, err = checkiday.On("07/30/2018")
	if err != nil {
		fmt.Printf("Oh noes, an error occured: %s", err.Error())
	}

	for idx := range days.Holidays {
		fmt.Printf("On 07/30/2018 we celebrated %s\n", days.Holidays[idx].Name)
	}
}
