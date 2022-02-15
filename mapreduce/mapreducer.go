package mapreduce
func MapReduce(nums []int,util func(int) int) []int{

	var res[]int
	for i:=0;i<len(nums);i++{
		if util(nums[i])==1{
			res=append(res,nums[i])
		}

		
	}
	return res

}