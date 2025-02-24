package main
import ("fmt"
"math")

type Vertex struct{
	X,Y float64
}
func (v Vertex) Abs() float64{
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
func (v *Vertex) Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func main(){
	a:= Vertex{3,4}
	a.Scale(10)
	fmt.Println(a.Abs())
}