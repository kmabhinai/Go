package main

import (
	"fmt"
	"sort"
)

func main() {
	var vegList = []string{"potato", "potato2"}
	vegList = append(vegList, "potato3")
	fmt.Println(vegList)

	vegList = append(vegList[1:])
	fmt.Println(vegList)

	highScores := make([]int, 4)
	highScores[0] = 100
	highScores[1] = 400
	highScores[2] = 300
	highScores[3] = 200

	highScores = append(highScores, 500, 600)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//remove values from slices based on index
	highScores = append(highScores[:2], highScores[3:]...)
	fmt.Println(highScores)

}
