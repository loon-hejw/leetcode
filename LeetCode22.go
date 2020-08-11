package main

import (
	"fmt"
)


func generateParenthesis(n int) []string {
	var generate []string
	return _generate(0,0,n,"",generate)
}

func _generate(left int ,right int ,max int,s string,generate []string) []string {

	// 终止条件
	if left == max && right == max {
		generate = append(generate ,s)
		return generate
	}

	// 处理逻辑过程
	if left < max {
		// 递归
		generate = _generate(left + 1 , right , max , s + "(",generate)
	}

	if right < left {
		// 递归
		generate = _generate(left , right + 1 , max , s + ")",generate)
	}

	return generate
}


func main() {
	generate := generateParenthesis(3);

	for v , k := range generate {
		fmt.Println(v,k)
	}

}
