package kata
import "math"

func PowersOfTwo(n int) []uint64 {
  // your code goes here
  mySlice:= []uint64{}
  for i:=0; i<= n; i++{
    res := math.Pow(2, float64(i))
    mySlice = append(mySlice, uint64(res))
  }
  return mySlice
}