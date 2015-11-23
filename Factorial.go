package main


import (
	"fmt"
)

func main(){
	var n int
	fmt.Print("Input n! : ")
	_, err := fmt.Scanf("%d", &n)
	if(err != nil){
		fmt.Println(err)
	}
	fmt.Println(n, "!")
	fmt.Println(factorial(n))
}

func factorial(n int) float64 {
	var result float64 = 1
	for i:=1; i <= n; i ++ {
		result *= (float64)(i)
		if(i < n){
			fmt.Print(i, " x ")
		}else{
			fmt.Print(i, " = ")
		}
	}	
	return result
}
