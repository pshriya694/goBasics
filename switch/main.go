package main

import "fmt"

func main() {
	//var day int
	//fmt.Println("Enter number b/w 1 to 7")
	//fmt.Scan(&day)
	//
	//switch day {
	//case 1:
	//	fmt.Println("Monday")
	//case 2:
	//	fmt.Println("Tuesday")
	//case 3:
	//	fmt.Println("Wednesday")
	//case 4:
	//	fmt.Println("Thursday")
	//case 5:
	//	fmt.Println("Friday")
	//case 6:
	//	fmt.Println("Saturday")
	//case 7:
	//	fmt.Println("Sunday")
	//default:
	//	fmt.Println("Wrong choice!")
	//}
	//

	//switch statement with multiple values
	month := "march" //CASE SENSITIVE

	switch month {
	case "january", "february", "december":
		fmt.Println("Winter")
	case "April", "march", "june":
		fmt.Println("Spring")
	default:
		fmt.Println("Others Season")

	}

	// can make switch cases like this too
	//
	//temperature := 12
	//switch {
	//case temperature > 0:
	//	fmt.Println("Freezing")
	//case temperature < 15:
	//	fmt.Println("Cold")
	//default:
	//	fmt.Println("Any other")
	//}

}
