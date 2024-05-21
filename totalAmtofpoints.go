package kata
import "strings"

func Points(games []string) int {
  // your code here!
  var score int
  for i:= 0; i<len(games);i++{
    res:= strings.Split(games[i],":")
    if res[0]>res[1]{
       score +=3
    }else if res[0] == res[1]{
      score +=1
    }else{
      score +=0
    }
    
  }
  return score
}