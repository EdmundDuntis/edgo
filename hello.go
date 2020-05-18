package main

import "fmt"

func main() {
	var nums []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(bs(nums, 8))

	var mud []int = []int{7, 8, 5, 3, 9, 4, 6, 1, 2}
	fmt.Println(mud)
	quicksort(mud)
	fmt.Println(mud)

}

func bs(nums []int, target int) int {
	var begin = 0
	var end = len(nums) - 1
	if target > nums[end] || target < nums[begin] {
		return -1
	}
	for begin <= end {
		var mid = begin + (end-begin)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			begin = mid + 1
		}
	}
	return -1
}

func quicksort(nums []int) {
	sort(nums, 0, len(nums)-1)
}
func sort(nums []int, left int, right int) {
	if left > right {
		return
	}

	var base, i, j = nums[left], left, right

	for i != j {
		for nums[j] >= base && i < j {
			j--
		}
		for nums[i] <= base && i < j {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	nums[left], nums[i] = nums[i], base
	sort(nums, left, i-1)
	sort(nums, i+1, right)
}
