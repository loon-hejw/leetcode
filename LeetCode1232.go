package main

func checkStraightLine(coordinates [][]int) bool {

	if len(coordinates) <= 2 {
		return true
	}
	// y = kx + b
	k := float64(coordinates[1][1] - coordinates[0][1]) / float64(coordinates[1][0] - coordinates[0][0])
	b := -((k * float64(coordinates[0][0])) - float64(coordinates[0][1]))


	for i := 2 ; i < len(coordinates) ; i ++ {
		if float64(coordinates[i][1]) != k * float64(coordinates[i][0]) + b {
			return false
		}
	}
	return true
}

func main() {
	
}
