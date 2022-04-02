package arraysum

func SumArray(ints []int) int {
	var result int
	for _, i := range ints {
		result += i
	}
	return result
}

func SumAllArrays(ints ...[]int) int {
	var result int
	for _, arr := range ints {
		result += SumArray(arr)
	}
	return result
}

func SumArrays(ints ...[]int) []int {
	var sums []int
	for _, arr := range ints {
		sums = append(sums, SumArray(arr))
	}
	return sums
}

func SumTails(ints ...[]int) []int {
	var sums []int
	for _, arr := range ints {
		if len(arr) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, SumArray(arr[1:]))
		}
	}
	return sums
}

func main() {

}
