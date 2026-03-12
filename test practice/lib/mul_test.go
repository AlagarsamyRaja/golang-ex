package lib

import (
	"fmt"
	"testing"
)

func TestMul(t *testing.T) {

	tt:=[]struct{
		name string
		x int
		y int
		v int
	}{
		{"two * three",2,3,6},
		{"four * five",4,5,20},
	}

	for _, ti := range tt {
		t.Run(ti.name,func(t *testing.T) {
			v:=Mul(ti.x,ti.y)
			if v != ti.v{
				t.Errorf("Expected %d Got %d",ti.v,v)
			}
		})
	}
}

func ExampleMul()  {
	fmt.Println(Mul(2,3))
	//output:6
	fmt.Println(Mul(2,4))
	//output:8
	fmt.Println(Mul(4,5))
	//output:20
}