package main

import "fmt"

//  values of a, vo, and so.
// The returned function remembers these values,
func genDisplaceFn(a, vo, so float64) func(float64) float64 {
	// Anonymous function (closure)
	return func(t float64) float64 {
		// Formula:
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	fmt.Println("GenDisplace Function Program")

	// read acceleration
	var a float64
	fmt.Print("Enter a: ")
	fmt.Scan(&a)

	// read initial velocity
	var vo float64
	fmt.Print("Enter Vo: ")
	fmt.Scan(&vo)

	// read initial displacement
	var so float64
	fmt.Print("Enter So: ")
	fmt.Scan(&so)

	// Generate a new displacement function, returned function already remembers.
	fn := genDisplaceFn(a, vo, so)

	// read time
	var t float64
	fmt.Print("Enter t: ")
	fmt.Scan(&t)

	// Example: fn(3) -> calculate displacement after 3 seconds.
	fmt.Println(fn(t))
}
