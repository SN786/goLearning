package main

import (
	"fmt"
	check "github.com/SN786/goLearning/stringUtil"
	mapr "github.com/SN786/goLearning/mapreduce"
)

func onlyEven(x int) int{
	if x%2==0{
		return 1
	}
	return 0
}
func main() {
 	//Palindrome Finder
	s:="AAA"
	if check.IsPalindrome(s) {
		fmt.Println(s," is palindrome")
	} else {
		fmt.Println(s,"is not palindrome")
	}


//Map Reduce Function
	nums := []int {1,2,3,4,5,6,7,8}
	
	fmt.Println(mapr.MapReduce(nums,onlyEven))

}
