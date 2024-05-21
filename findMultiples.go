package kata

func FindMultiples(integer, limit int) []int {
  // Your code here!
  mySlice:= []int{}
  for i:=1; i<= limit;i++{
    if(i % integer == 0){
      mySlice =append(mySlice,i)
    }
  }
  return mySlice
}
