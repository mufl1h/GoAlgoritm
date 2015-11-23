package main

import (
	"fmt"
)

func main(){
	var i int
	fmt.Print("Input String (fmt.Scanf): ")
	_, err := fmt.Scanf("%d", &i)
	if(err != nil){
		fmt.Println(err)
	}
	fmt.Println(i)
}