package main

import (
	"fmt"
	"math"
)

const input = 265149

func FloorSqrt(x int) int {
	return int(math.Floor(math.Sqrt(float64(x))))
}

func Ring(x int) int {
	return (FloorSqrt(x-1) + 1) / 2
}

func PosInRing(x int) int {
	return x - (2*Ring(x)-1)*(2*Ring(x)-1)
}

func SegmentInRing(x int) int {
	if Ring(x) == 0 {
		return 0
	}
	return ((PosInRing(x) - 1) / Ring(x)) + 1
}

func PosInSegment(x int) int {
	return PosInRing(x) - (SegmentInRing(x)-1)*Ring(x)
}

func ManhattanDistance(x int) int {
	if SegmentInRing(x)%2 == 0 {
		return Ring(x) + PosInSegment(x)
	}
	return 2*Ring(x) - PosInSegment(x)
}

func main() {
	fmt.Println(input, "must be carried", ManhattanDistance(input), "steps")
}
