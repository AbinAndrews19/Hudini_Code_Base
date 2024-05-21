package kata

func SquareSum(numbers []int) int {
    // your code here
  sum:=0
  for i:=0;i<len(numbers);i++{
    sqr := numbers[i] * numbers[i]
    sum += sqr
  }
  return sum
}