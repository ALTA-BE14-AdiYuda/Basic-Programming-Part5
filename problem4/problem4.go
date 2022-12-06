package main

import "fmt"

func FindaMaxSumSubArray(k int, arr []int) int {
	// your code here
	var lSum, rSum int
	var res int
	var lSlc, rSlc []int

	lSlc = arr[:k]

	if k*2 > len(arr) {
		rSlc = arr[k-1:]
	} else {
		rSlc = arr[k-1 : k*2-1]
	}
	for _, v := range lSlc {
		lSum += v
	}
	for _, v := range rSlc {
		rSum += v
	}
	if lSum > rSum {
		res = lSum
	} else {
		res = rSum
	}
	return res

}

func main() {
	fmt.Println(FindaMaxSumSubArray(3, []int{2, 1, 5, 1, 3, 2})) //9
	fmt.Println(FindaMaxSumSubArray(2, []int{2, 3, 4, 1, 5}))    //7
	fmt.Println(FindaMaxSumSubArray(2, []int{2, 1, 4, 1, 1}))    //5
	fmt.Println(FindaMaxSumSubArray(3, []int{2, 1, 4, 1, 1}))    //7
	fmt.Println(FindaMaxSumSubArray(4, []int{2, 1, 4, 1, 1}))    //8
}
