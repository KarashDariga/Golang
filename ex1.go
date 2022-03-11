package main

import (
	"fmt"
	"strings"
)

func scoreSummary(name string, score1, score2, score3 float64) {
	avg := (score1+score2+score3)/3
	fmt.Printf("%10s|%8.2f|%8.2f|%8.2f|%8.2f\n", name, score1, score2, score3, avg)
}



func main() {
	fmt.Printf("%10s|%8s|%8s|%8s|%8s\n", "Name", "Score 1", "Score 2", "Score 3", "Average")
	fmt.Println(strings.Repeat("-", 46))

	
	scoreSummary("Dariga", 85.40, 75.60, 89.0)
	scoreSummary("Saule", 90.50, 50.80, 90.80)
	scoreSummary("Zhans", 70.60, 50.0, 90.30)