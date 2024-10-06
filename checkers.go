package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
)

func init() {
	log.SetFlags(0)
}

func progressChecker(progress string) (uint64, float64) {
	timeFloat, err := strconv.ParseFloat(progress, 64)

	if err != nil {
		log.Fatal(err)
	}
	ipart := uint64(timeFloat)
	decpartS := fmt.Sprintf("%.6f", timeFloat-float64(ipart))
	decpart, err := strconv.ParseFloat(decpartS, 64)

	if err != nil {
		log.Fatal(err)
	}

	leftMin := math.Round(60 - decpart*60)
	if leftMin == 60 {
		leftMin = 0
	}

	leftHour := 20 - ipart

	if leftMin > 0 {
		leftHour--
	}

	fmt.Printf("Progress: %vh %vm\n", ipart, math.Round(decpart*60))
	fmt.Print("left: ", leftHour, "h ", leftMin, "m\n")

	return leftHour, leftMin
}

func survivalChecker(leftHour uint64, leftMin float64) {
	currentTime := time.Now()
	weekdayAsInt := int(currentTime.Weekday())

	if weekdayAsInt == 0 {
		weekdayAsInt = 7
	}

	possibleHours := (8 - weekdayAsInt) * 24

	maxHour := possibleHours - currentTime.Hour()
	maxMin := 60 - currentTime.Minute()

	if maxMin > 0 {
		maxHour -= 1
	}
	excess := (maxHour*60 + maxMin) - (int(leftHour)*60 + int(leftMin))

	printSurvive(maxHour, maxMin, excess, 20)
}

func printSurvive(maxHour, maxMin, excess, rounding int) {

	fmt.Println("===================================")
	fmt.Print("Before weekly restart: ", maxHour, "h ", maxMin, "m\n")

	if excess+rounding >= 0 {
		fmt.Println("You will survive 😁")
		fmt.Print("Excess without ", rounding, "m rounding: ", excess, "m\n")
		fmt.Print("Excess with ", rounding, "m rounding: ", excess+rounding, "m\n")
	} else {
		fmt.Println("You will lose one life 😭")
	}
	fmt.Println("===================================")
}

func timeChecker(progress string) {

	leftHour, leftMin := progressChecker(progress)

	survivalChecker(leftHour, leftMin)
}
