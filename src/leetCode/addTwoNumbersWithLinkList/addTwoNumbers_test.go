package AddTowNumbersLinkList

import (
	"testing"

	"leetCode/common"

	"github.com/smartystreets/goconvey/convey"
)

func Test_addTwoNumbers(t *testing.T) {
	convey.Convey("Subject:Test add two numbers by linkList", t, func(){
		convey.Convey("Case:Test [2, 4, 3] + [5, 6, 4] = [7, 0, 8]",func() {
			tArr := []int{7, 0, 8}
			res := addTwoNumbers(common.MakeArrayToLinkList([]int{2, 4, 3}), common.MakeArrayToLinkList([]int{5, 6, 4}))
			for i:=0; i < len(tArr); i++ {
				convey.So(res.Val, convey.ShouldEqual, tArr[i])
				res = res.Next
			}
		})
		convey.Convey("Case: Test [1, 9, 9] + [9, 2, 2] = [0, 2, 2, 1]",func(){
			tArr := []int{0, 2, 2, 1}
			res := addTwoNumbers(common.MakeArrayToLinkList([]int{1, 9, 9}), common.MakeArrayToLinkList([]int{9, 2, 2}))
			for i:=0; i < len(tArr); i++ {
				convey.So(res.Val, convey.ShouldEqual, tArr[i])
				res = res.Next
			}
		})
		convey.Convey("Case: Test [9, 8] + [1] = [0, 9]",func(){
			tArr := []int{0, 9}
			res := addTwoNumbers(common.MakeArrayToLinkList([]int{9,8}), common.MakeArrayToLinkList([]int{1}))
			for i := 0; i < len(tArr); i++ {
				convey.So(res.Val, convey.ShouldEqual, tArr[i])
				res = res.Next
			}
		})
		convey.Convey("Case: Test [0] + [7,3] = [7, 3]",func(){
			tArr := []int{7, 3}
			res := addTwoNumbers(common.MakeArrayToLinkList([]int{0}), common.MakeArrayToLinkList([]int{7, 3}))
			for i := 0; i < len(tArr); i++ {
				convey.So(res.Val, convey.ShouldEqual, tArr[i])
				res = res.Next
			}
		})
	})
}