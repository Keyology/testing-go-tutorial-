package main 

import (
	"fmt"
)

// function that that add x + 2 and return the results
func Calculate(x int ) (result int){
	result = x + 2 
	return result 


}

func main(){
	fmt.Println(Calculate(2))
}