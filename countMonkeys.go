package kata

func monkeyCount(n int) []int {
   // Your code here, happy coding!
   mySlice := [] int{}
   for i:=1;i<=n;i++{
     mySlice = append(mySlice,i)
   }
   return mySlice
}