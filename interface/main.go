package main
import(
	"fmt"
	"math"
)
type Shape interface{
	Area() float64
	Perimeter() float64
}
type Object interface{
	Volume() float64
}
type Rect struct {
	width float64
	height float64
}
type Circle struct {
	radius float64
}
func (c Circle) Volume() float64{
	return c.radius*c.radius*c.radius
}

func (r Rect) Area() float64{
	return r.width * r.height
}
func (r Rect) Perimeter() float64{
	return 2*(r.width + r.height)
}
func (c Circle) Area() float64{
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float64{
	return 2*math.Pi*c.radius
}
func TotalArea(shapes ...Shape) float64{
	total :=0.0
	for _, s := range shapes{
		total +=s.Area()
	}
	return total
}
func explain(i interface{}){
	switch i.(type){
	case string:
		fmt.Println(i)
	case int:
		fmt.Println("number=",i);
	default:	
		fmt.Printf("type=%T, value=%v\n", i, i)
	}
}
///function main
func main(){
	var s Shape = Rect{1,2}
	r:=s.(Rect)
	fmt.Printf("type:%T\n",r)
	fmt.Printf("v=%v,type=%T\n",s,s)

	var s1 Circle  = Circle{10}
	fmt.Printf("v=%v,type=%T\n",s,s)

	total:=TotalArea(s, s1)
	fmt.Printf("total=%v type=%T\n",total, total)

	explain(s)
	explain(s1)
	var s2 Shape = s1
	var o Object = s1
	fmt.Println("area=", s2.Area())
	fmt.Println("volumn=", o.Volume())	
	explain(s2)
	explain(o)
	explain(100)
	explain("Hello")



}
