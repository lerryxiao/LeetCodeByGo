//
//  Author:  lerry
//  Purpose: 
//  Created: 09/12/2016 21:04

package ThirdMaximumNumber414

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
)

func Test_thirdMax(t *testing.T){
	convey.Convey("Subject:Test thirdMax", t, func() {
		convey.Convey("Case:Input: [3, 2, 1]", func() {
			convey.So(thirdMax([]int{3,2,1}), convey.ShouldEqual, 1)
		})
		convey.Convey("Case:Input: [1, 2]", func() {
			convey.So(thirdMax([]int{1, 2}), convey.ShouldEqual, 2)
		})
		convey.Convey("Case:Input: [2, 2, 3, 1]", func() {
			convey.So(thirdMax([]int{2, 2, 3, 1}), convey.ShouldEqual, 1)
		})
	})
}