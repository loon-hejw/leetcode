package main

import (
	"fmt"
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {

	var (
		customerName []string
		customerSet  = map[string]bool{}
		tableName    []int
		tableSet     = map[int]bool{}
		tableMap     = map[int]map[string]int{}
	)

	for _, order := range orders {
		tableInt, _ := strconv.Atoi(order[1])
		if !tableSet[tableInt] {
			tableName = append(tableName, tableInt)
			tableSet[tableInt] = true
		}

		if _, ok := tableMap[tableInt]; !ok {
			tableMap[tableInt] = map[string]int{}
			tableMap[tableInt][order[2]] = 0
		}

		if !customerSet[order[2]] {
			customerSet[order[2]] = true
			customerName = append(customerName, order[2])
		}
		tableMap[tableInt][order[2]] ++
	}

	sort.Ints(tableName)
	sort.Strings(customerName)

	var res = make([][]string, len(tableName)+1)
	res[0] = append([]string{"Table"}, customerName...)

	for i := 0; i < len(tableName); i++ {
		res[i+1] = append(res[i+1], strconv.Itoa(tableName[i]))
		for j := 0; j < len(customerName); j++ {
			count, _ := tableMap[tableName[i]][customerName[j]]
			res[i+1] = append(res[i+1], strconv.Itoa(count))
		}
	}
	return res
}

func main() {
	res := displayTable([][]string{{"David", "3", "Ceviche"}, {"Corina", "10", "Beef Burrito"}, {"David", "3", "Fried Chicken"}, {"Carla", "5", "Water"}, {"Carla", "5", "Ceviche"}, {"Rous", "3", "Ceviche"}})
	fmt.Println(res)
}
