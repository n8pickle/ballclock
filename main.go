package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	ballCount, err := strconv.Atoi(os.Args[1])
	valid := ballCount >= 27 && ballCount <= 127 && err == nil
	if(!valid) {
		return;
	}

	start := time.Now()
	const MinuteCap = 5
	const FiveMinuteCap = 12
	const HourCap = 12

	var elapsed time.Duration 
	var minuteStack [MinuteCap]int
	var fiveMinStack [FiveMinuteCap]int
	var hourStack [HourCap]int

	// indexes

	// ReserveQueue
	rqRemIndex := 0
	rqInsIndex := 0

	// Minute Stack
	msIndex := 0

	// Five Minute Stack
	fmIndex := 0

	// Hour Stack
	hsIndex := 0

	baseArr := [127]int{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24,25, 26, 27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127 }
	reserveQueue := baseArr[:ballCount]

	
	// rqRemIndex = (rqRemIndex + 1) % ballCount

	minutesToRepeat := 0.00
	for {
		minutesToRepeat++
		minuteBall := reserveQueue[rqRemIndex]
		rqRemIndex = (rqRemIndex + 1) % ballCount
		minuteStack[msIndex] = minuteBall

		if msIndex == 4 {
			fiveMinBall := minuteStack[msIndex]
			fiveMinStack[fmIndex] = fiveMinBall
			for i := 3; i >= 0; i-- {
				reserveQueue[rqInsIndex] = minuteStack[i]
				rqInsIndex = (rqInsIndex + 1) % ballCount
			}
			if fmIndex == 11 {
				hourBall := fiveMinStack[fmIndex]
				hourStack[hsIndex] = hourBall
				for i := 10; i >= 0; i-- {
					reserveQueue[rqInsIndex] = fiveMinStack[i]
					rqInsIndex = (rqInsIndex + 1) % ballCount
				}
				if hsIndex == 11 {
					lastBall := hourStack[hsIndex]
					for i := 10; i >= 0; i-- {
						reserveQueue[rqInsIndex] = hourStack[i]
						rqInsIndex = (rqInsIndex + 1) % ballCount
					}
					reserveQueue[rqInsIndex] = lastBall
					rqInsIndex = (rqInsIndex + 1) % ballCount
					if rqRemIndex == 0 && rqInsIndex == 0 {
						if ballsRepeat(reserveQueue, ballCount){
							break
						}
					}
				}
				hsIndex = (hsIndex + 1) % HourCap
			}
			fmIndex = (fmIndex + 1) % FiveMinuteCap
		}
		msIndex = (msIndex + 1) % MinuteCap
	}
	elapsed = time.Since(start)
	daysToRepeat := minutesToRepeat/1440
	fmt.Printf("%v balls took %v days to repeat (%v elapsed)\n", ballCount, daysToRepeat, elapsed);
}

func ballsRepeat(reserveQueue []int, ballCount int) bool {
		for i := ballCount - 1; i >= 0; i-- {
			if reserveQueue[i] == i + 1 {
				continue
			} else { 
				return false 
			}
		}
		return true
}

func createInitialArray(arr []int, ballCount int) []int {
	for i := 0; i < ballCount; i++ {
		arr[i] = i
	}
	return arr
}

func userSetBallCount() int {
	needValidBallCount := true
	ballCount := 0
	
	for needValidBallCount {
		fmt.Println("How many balls should be used? (27-127) ")

		fmt.Scanln(&ballCount)


	}
	return ballCount
}
