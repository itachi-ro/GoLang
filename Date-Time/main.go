package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println("Current Date and Time:", currentTime) // Current Date and Time: 2025-05-22 13:18:24.5494494 +0530 IST m=+0.000000001

	// Format the date
	fmt.Println("Formatted Date:", currentTime.Format("02-Jan-2006")) //Formatted Date: 22-May-2025
	fmt.Println(currentTime.Format("2006-01-02"))                     //2025-05-22
	fmt.Println("Formatted Time:", currentTime.Format("15:04:05"))    // Formatted Time: 13:18:24
	year := currentTime.Format("2006")                                // Extract only the year
	hour := currentTime.Format("03 PM")                               // Extract only the hour in (12-H format)

	fmt.Println("Current Year:", year) //  -->  Current Year: 2025
	fmt.Println("Current Hour:", hour) // --> Current Hour: 01 PM

}

/*

15 is for 24-hour format. 03 is for 12-hour format.
You want	Use in Go Format string
Year			2006
Month			01
Day				02
Hour (24h)		15
Minute			04
Second			05

Explanation:
Part	Meaning
03		Hour (12-hour format)
04		Minutes
05		Seconds
PM		AM/PM indicator

*/
