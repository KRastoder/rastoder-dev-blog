package main

import "fmt"

func plus(a int , b int )int {
	return a + b
}
func plusplus(a,b,c int ) int {
	return a + b + c
}
func sum(nums ...int ){
	fmt.Print(nums, " ")
	total := 0

	for _,num := range nums{
		total += num
	}
	fmt.Println(" =",total)
}

func main() {
	res := plus(1,2)
	fmt.Println("1 + 2 =",res)


	res  = plusplus(1,2 , plus(1,2))
	fmt.Println("1 + 2 + 3 =",res)

	fmt.Println("##############################\n")
	sum(1,2,)
	sum(10,10,10,10,10,10,7)
	
}
