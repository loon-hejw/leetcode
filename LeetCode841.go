package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	vis := make([]bool,len(rooms))
	queue := []int{}
	vis[0] = true
	queue = append(queue,0)
	num := 0
	for i := 0 ; i < len(queue) ; i ++ {
		for _ , j := range rooms[queue[i]] {
			if !vis[j] {
				vis[j] = true
				queue = append(queue,j)
			}
		}
		num ++
	}
	return num == len(rooms)
}

var number int
var vis    []bool
func canVisitAllRoomsDfs(rooms [][]int) bool {
	number = 0
	vis    = make([]bool,len(rooms))
	vis[0] = true
	canVisitAllRoomsDfsHerlper(rooms,0)
	return number == len(rooms)
}


func canVisitAllRoomsDfsHerlper(rooms [][]int , x int) {
	number ++
	for _ , i := range rooms[x] {
		if !vis[i] {
			vis[i] = true
			canVisitAllRoomsDfsHerlper(rooms,i)
		}
	}
}

func main() {
	fmt.Println(canVisitAllRoomsDfs([][]int{{1},{2},{3},{}}))
}
