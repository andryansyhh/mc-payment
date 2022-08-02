package main

// func main() {
// 	num := []int{1, 5, 8, 0, 9, 7, 4, 3, 2}
// 	max, min := MaxMin(num)
// 	fmt.Println(max, min)
// }

func MaxMin(num []int) (max, min int) {
	max = num[0]
	min = num[1]
	for _, v := range num {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}
	return max, min
}
