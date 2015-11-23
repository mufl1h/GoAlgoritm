package main


import (
	"fmt"
)

func main(){
	fmt.Println(factorial(0))
}

func factorial(n int) float64 {
	var result float64 = 1
	for i:=1; i <= n; i ++ {
		result *= (float64)(i)
		if(i < n){
			fmt.Print(i, "x")
		}else{
			fmt.Print(i, "=")
		}
	}	
	return result
}
