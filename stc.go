// stc.go (small transport company)

package main

import (
	"fmt"
	"math"
)

type stnslist []int

func carsNeededAtStation(passengersAtStation int, busCapacity int, carCapacity int) (int, int) {
	var cwb int

	//cwob stands for cars needed without the bus
	cwob := int(math.Ceil(float64(passengersAtStation) / float64(carCapacity)))

	d := passengersAtStation - busCapacity
	//cwb stands for cars needed with the bus
	if d <= 0 {
		cwb = 0
	} else {
		cwb = int(math.Ceil(float64(d) / float64(carCapacity)))
	}
	return cwb, cwob

}

func CarsNeeded(busCapacity int, carCapacity int, stations stnslist) int {
	numberOfStations := len(stations)
	carsNeededWithoutBus := make([]int, numberOfStations)
	carsNeededWithBus := make([]int, numberOfStations)

	fmt.Println("Bus capacity:", busCapacity, "- Car capacity:", carCapacity)
	fmt.Println("Number of stations:", numberOfStations, "- Passengers at stations:", stations)

	for i, passengersAtStation := range stations {
		cwb, cwob := carsNeededAtStation(passengersAtStation, busCapacity, carCapacity)
		carsNeededWithoutBus[i] = cwob
		carsNeededWithBus[i] = cwb
	}

	minCarsNeeded := math.MaxInt64
	carsNeeded := 0

	for i := 0; i < numberOfStations; i++ {
		carsNeeded = carsNeededWithBus[i]

		for j := 0; j < numberOfStations; j++ {
			if i == j {
				continue
			}
			carsNeeded = carsNeeded + carsNeededWithoutBus[j]
		}
		if carsNeeded < minCarsNeeded {
			minCarsNeeded = carsNeeded
		}
	}

	fmt.Println("Cars needed:", carsNeeded)
	fmt.Println()

	return carsNeeded
}

func main() {

	fmt.Println("Hello from Small Transport Company")

	//STATIONS = [15, 10]
	//B = 12
	//C = 5
	//Returns: 3
	B := 12
	C := 5
	stations := stnslist{15, 10}

	CarsNeeded(B, C, stations)

	//STATIONS = [11, 15, 13]
	//B = 9
	//C = 5
	//Returns: 7
	B = 9
	C = 5
	stations = stnslist{11, 15, 13}

	CarsNeeded(B, C, stations)

}
