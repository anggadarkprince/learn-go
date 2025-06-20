package main

import "fmt"

func main() {
	// Slice is a dynamic pointer or reference to array 
	// array[low:high] take index low to high
	// array[low:] take index low to last index
	// array[:high] take index 0 to high
	// array[:] take all array

	months := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	slice1 := months[4:6]
	fmt.Println(slice1, "length", len(slice1), "capacity", cap(slice1))

	slice2 := months[:2]
	fmt.Println(slice2, "length", len(slice2), "capacity", cap(slice2))

	slice3 := months[6:]
	fmt.Println(slice3, "length", len(slice3), "capacity", cap(slice3))

	allMonths := months[:]
	fmt.Println(allMonths, "length", len(allMonths), "capacity", cap(allMonths))

	// similar to manual slice declaration
	var slice []string = months[:];
	fmt.Println(slice);

	// Append slice (whe we change the slice, the reference of array is replaced)
	days := [...]string{
		"Sunday", "Monday", "Tuesday", "Wednesday",
	}
	daySlice := days[2:] // [0]Tuesday, [1]Wednesday
	daySlice[1] = "Rabu" // change Wednesday
	fmt.Println(daySlice)
	fmt.Println(days) // original array also change

	newDays := append(daySlice, "Thursday", "Friday", "Saturday") // append to last element, but if it full, then create new array (with length that fit to current capacity)
	fmt.Println(newDays)
	fmt.Println(days) // original array days not changed (not added)
	newDays[1] = "Wednesday" // revert to original value
	fmt.Println(newDays)
	fmt.Println(days) // because newDays is new array, then it not affecting the previous original array

	// New empty slice
	var newSlice = make([]string, 2, 5) // length 2 but capacity 5 (length can be assign like array, but the rest of capacity updated via append)
	newSlice[0] = "Angga"
	newSlice[1] = "Ari"
	// newSlice[2] = "Wijaya" // cannot assign index 2 because the length only 2
	newSliceName := append(newSlice, "Wijaya")
	fmt.Println(newSliceName, "length", len(newSliceName), "capacity", cap(newSliceName))

	// Copy slice
	newMonths := make([]string, len(months), cap(months))
	copy(newMonths, months[:]) // copy only copy slice to slice (cannot array)
	fmt.Println(newMonths)

	// declaration Slice and Array
	thisArray := [...]int{1, 2, 3}
	thisSLice := []int{1, 2, 3}
	fmt.Println(thisArray, thisSLice)
}