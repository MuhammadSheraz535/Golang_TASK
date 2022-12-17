package main

import (
	"fmt"
	"math/big"
)

func main() {

	// 8.1. Hitting the ceiling

	const lightSpeeds = 299792 // km/s
	const secondsPerDays = 86400
	var distances int64 = 41.3e12
	fmt.Println("Alpha Centauri is", distances, "km away.")
	day := distances / lightSpeeds / secondsPerDays
	fmt.Println("That is", day, "days of travel at light speed.")

	fmt.Println("*********************")

	// 8.2 Big package

	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)
	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println("Andromeda Galaxy is", distance, "km away.")
	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)
	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	fmt.Println("That is", days, "days of travel at light speed.")
	fmt.Println("*********************")

	// 8.3. Constants of unusual size

	const distanc = 24000000000000000000
	const light_Speed = 299792
	const seconds_PerDay = 86400
	const days_ = distanc / light_Speed / seconds_PerDay
	fmt.Println("Andromeda Galaxy is", days_, "light days away.")
}
