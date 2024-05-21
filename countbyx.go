package kata

func CountBy(x, n int) []int {
	mySlice := []int{} // creating an empty slice
	for i := 1; i <= n; i++ {
		mySlice = append(mySlice, i*x)
	}
	return mySlice
}
