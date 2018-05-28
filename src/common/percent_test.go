package common

import (
	"testing"
	"fmt"
	"math/rand"
)

func TestPercent(t *testing.T) {
	for i:=0;i<1000 ;i++  {
		a1:=rand.Intn(10000)
		a2:=rand.Intn(10000)
		//fmt.Println(a1)
		//fmt.Println(a2)
		fmt.Println(Percent(float64(a1),float64(a2)))
	}
}

func BenchmarkPercent(b *testing.B) {
	for i:=0;i<1000 ;i++  {
		a1:=rand.Intn(10000)
		a2:=rand.Intn(10000)
		//fmt.Println(a1)
		//fmt.Println(a2)
		fmt.Println(Percent(float64(a1),float64(a2)))
	}
}