package main

import (
	"fmt"
	"sort"
)

func violence(nums[] int) [][] int{
	var result [][]int
	sort.Ints(nums)
	if nums == nil || len(nums) < 3  {
		return result
	}

	HashMap := map[string]string{}

	for i := 0 ; i < len(nums) - 2 ; i ++  {
		for j := i + 1; j< len(nums) - 1 ; j ++ {
			for k := j + 1 ; k < len(nums) ; k ++  {
				if nums[i] + nums[j] + nums[k] == 0 {
					tmp := []int{nums[i],nums[j],nums[k]}

					if _ , ok := HashMap[fmt.Sprint(tmp)] ; !ok {
						result = append(result,tmp)
					}

					HashMap[fmt.Sprint(tmp)] = "1"
				}
			}
		}
	}

	return result
}

func HashdirectlySolution(nums[] int) [][] int{

	var result [][]int;
	if nums == nil || len(nums) < 3  {
		return result
	}
	sort.Ints(nums)
	HashAll := map[string]int{}
	for i:=0 ; i<len(nums) - 2 ; i++  {
		targer := -nums[i]
		HashMap := map[int]int{}
		for j:= i+1; j < len(nums) ; j++ {
			tmp := targer - nums[j]

			if v , ok := HashMap[tmp] ; ok {

				arr := []int{nums[i],v,nums[j]}
				if  _ , ok2 := HashAll[fmt.Sprint(arr)] ; !ok2 {
					result = append(result, arr)
				}
				HashAll[fmt.Sprint(arr)] = v
			}

			HashMap[nums[j]] = nums[j]
		}
	}
	return result
}

func doubledirectlySolution(nums[] int) [][] int{
	var result [][]int
	if nums == nil || len(nums) < 3 {
		return result
	}

	sort.Ints(nums)
	for i := 0 ; i < len(nums) - 2 ; i ++  {
		if nums[i] > 0 { break }
		if i > 0 && nums[i] == nums[i - 1] { continue }
		k , j := i + 1 , len(nums) - 1

		for k < j {
			sum := nums[i] + nums[k] + nums[j]
			if ( sum > 0 ) {
				for k < j && nums[j] == nums[j - 1]  { j -- }
				j = j - 1
			} else if ( sum < 0 ){
				for k < j && nums[k] == nums[k + 1]  { k ++ }
				k = k + 1
			} else {
				result = append(result, []int{nums[i] , nums[k] , nums[j]})
				for k < j && nums[j] == nums[j - 1]  { j -- }
				for k < j && nums[k] == nums[k + 1]  { k ++ }
				j = j - 1
				k = k + 1
			}
		}
	}
	return result
}
func main() {
 	fmt.Println(doubledirectlySolution([]int{-4,-2,1,-5,-4,-4,4,-2,0,4,0,-2,3,1,-5,0}))
}
