//  Author:  lerry
//  Purpose: 
//  Created: 09/12/2016 20:34

package ThirdMaximumNumber414

func thirdMax(nums []int) int {
	switch len(nums) {
	case 0:
		return 0;
	case 1:
		return nums[1];
	case 2:
		if(nums[0] > nums[1]){
			return nums[0]
		}else{
			return nums[1];
		}
	}
	for i, v := range nums {

	}
	return nums[2]
}
