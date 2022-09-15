package main

import "fmt"

//二分搜索的经典写法。需要注意的三点：
func binarySearch(nums []int, target int) int {
	low, height := 0, len(nums)-1

	for low <= height {
		mid := low + (height-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			height = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

// 二分查找第一个与 target 相等的元素，时间复杂度 O(logn)。 左边界

func searchLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + ((right - left) >> 1)
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}

	}
	if nums[left] == target {
		return left
	}
	return -1
	//return nums[left] == target?left:
}

func searchRight(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + ((right - left) >> 1) + 1 // 注意
		// 收缩右边界
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	if nums[right] == target {
		return right
	}
	return -1
	//return nums[right] == target ? right : -1;
}

//
func searMax() {

}
func main() {
	arr3 := []int{1, 2, 2, 2, 2, 2, 5}

	fmt.Println(binarySearch(arr3, 2))
	fmt.Println(searchLeft(arr3, 2))
	fmt.Println(searchRight(arr3, 2))
}
