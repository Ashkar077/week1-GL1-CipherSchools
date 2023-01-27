//  factorial
/*
package main
import "fmt"
func fact(n int)int{
	if n == 1 || n == 0 {
		return 1
	}
	if n < 0 {
		fmt.Println("Invalid number")
		return -1
	}
	result := n * fact(n-1)
	return result
}
func main() {
	fmt.Println(fact(10))
}
*/

//fibonacci

//fibonacci
/*package main
import "fmt"
func fibonacci(n int)[]int{
	fib := make([]int,n)
	fib[0] = 0
	fib[1]= 1
	for i :=2 ; i<n ; i++ {
		fib[i] = fib[i-1]+fib[i-2]
	}
	return fib
}
func main() {
	fmt.Println(fibonacci(10))
}
*/
package main

import "fmt"
func main() {
	fmt.Println(fib(5))
}
func fib(n int) int {
	if n ==0 || n==1 {
		return n
	}
	result := fib(n-1) + fib(n-2)
	return result
}