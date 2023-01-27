/*
package main
import "fmt"
func main(){
rec(4)
}
func rec(n int) {
	if n == 0 {
		return
	}
	if n%2 == 0 {
		rec(n - 1)
		fmt.Println(n - 1)
	}else{
		rec(n-1)
		fmt.Println(n-1)
	}
	fmt.Println(n -1)
}
*/

package main

import "fmt"

func main(){
rec(4)
}
func rec(n int) {
	if n <= 0 {
		return
	}
	rec(n - 1)
	rec(n-2)
	
	fmt.Println(n -1)
}