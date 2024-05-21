package kata


func GetSum(a, b int) int {
  min:=a
  max:=b
  sum:=0;
  if a > b {
    max = a
    min = b
  }else if a==b{
    return a
  }
  for i:= min; i<=max;i++{
    sum += i
  }
  return sum
}