package main

import (
	"fmt"
	"time"
)

func main() {
	const MinuteCap = 5
	const FiveMinuteCap = 12
	const HourCap = 12
	var elapsed time.Duration 
	minuteStack := make([]int, 0, MinuteCap)
	fiveMinStack := make([]int, 0, FiveMinuteCap)
	hourStack := make([]int, 0, HourCap)

	minutesToRepeat := 0.00
	ballCount := userSetBallCount()

	// constant hour stack ball
	reserveQueue := make([]int, ballCount)
	initialOrder := make([]int, ballCount)

	reserveQueue = createInitialArray(reserveQueue, ballCount)
	initialOrder = createInitialArray(initialOrder, ballCount)

	start := time.Now()
	for {
		minutesToRepeat++
		minuteBall := reserveQueue[0]
		reserveQueue = reserveQueue[1:(len(reserveQueue))]
		minuteStack = append(minuteStack, minuteBall)

		if len(minuteStack) == MinuteCap {
			fiveMinuteBall := minuteStack[len(minuteStack) - 1]
			minuteStack = minuteStack[:len(minuteStack)-1]
			fiveMinStack = append(fiveMinStack, fiveMinuteBall)
			for i := len(minuteStack) - 1; i >= 0; i-- {
				reserveQueue = append(reserveQueue, minuteStack[i])
				minuteStack = minuteStack[:i]
			}
			if len(fiveMinStack) == FiveMinuteCap {
				hourBall := fiveMinStack[len(fiveMinStack) - 1]
				fiveMinStack = fiveMinStack[:len(fiveMinStack)-1]
				hourStack = append(hourStack, hourBall)
				for i := len(fiveMinStack) - 1; i >= 0; i-- {
					reserveQueue = append(reserveQueue, fiveMinStack[i])
					fiveMinStack = fiveMinStack[:i]
				}
				if len(hourStack) == HourCap {
					lastBall := hourStack[len(hourStack) - 1]
					hourStack = hourStack[:len(hourStack)-1]
					for i := len(hourStack) - 1; i >= 0; i-- {
						reserveQueue = append(reserveQueue, hourStack[i])
						hourStack = hourStack[:i]
					}
					reserveQueue = append(reserveQueue, lastBall)
					if ballsRepeat(initialOrder, reserveQueue){
						elapsed = time.Since(start)
						break
					}
				}
			}
		}
	}
	daysToRepeat := minutesToRepeat/1440
	fmt.Printf("%v balls took %v days to repeat (%v elapsed)\n", ballCount, daysToRepeat, elapsed);
}

func ballsRepeat(initial []int, current []int) bool {
	if(len(initial) == len(current)) {
		for i := len(initial)-1; i >= 0; i-- {
			if initial[i] == current[i] {
				continue
			} else {
				return false
			}
		}

		return true
	}
	return false
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

		if(ballCount >= 27 && ballCount <= 127) {
			needValidBallCount = false
		} else {
			fmt.Println("That was an invalid number. Please enter a number within the range specified.")
		}
	}
	return ballCount
}
