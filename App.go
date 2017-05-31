/*****
  RAIZ CUADRADA POR EL METODO DE APROXIMACION DE NEWTON 
  AUTHOR : ROLANDO ARRIAZA  rolignu90@gmail.com
  LANG : GO 
*****/

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64 , iterations int ) float64 {

	var begin float64 = 2.0 ;
	var z float64 = 0.0;
	
	for i:=0 ; i < iterations ; i++ {
		z += (0.5) * (begin + x/begin); 
		begin = (z/float64((i+1))) ;
	}
	
	z = z/float64(iterations) ;

	return z;
}

func main() {
	fmt.Println("Aproximacion por medio de newton ")
	fmt.Println(Sqrt(2 , 10))
	fmt.Println("Aproximacion libreria ")
	fmt.Println(math.Sqrt(2))
}
