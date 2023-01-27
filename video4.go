/* package main
  import "fmt"
  func main() {
	for i := 0; i<3;i++ {
		if i == 2{
			continue
		}
		fmt.Println(i)
	}
  }
*/

/*package main
import "fmt"
func main() {
	nums := []int{1,2,3,4,5,6,0}
	for k,v := range nums {
	    fmt.Println(k)
		fmt.Println(v)
	}
}
*/

/*
package main
import "fmt"
func main() {
	nums := []int{1,2,3,4,5,6,0}
	for _,v := range nums {

		fmt.Println(v)
	}
}
*/

/*
package main
import "fmt"
func main() {
	nums := []string{"fadhil","tarun","rahul"}
	for k,v := range nums {
	    fmt.Println(k)
		fmt.Println(v)
}
}
*/

/*package main
import "fmt"
func main() {
	result, x := c()
	fmt.Println(result)
	fmt.Println(x)
	r,_ := b(1,2,3,4,5)
	fmt.Println(r)
}
func a()(int,string){
	return 1,"fadh"
}
func b(args ...int)(bool,int) {
	for _,v := range args {
		fmt.Println(v)
	}
	return true,11
}
func c() (i int,j string) {
	i = 10
	j = "rahul"
	return
}
*/

//pointer

/*
package main
import "fmt"
func main() {
	i := 10
	//j := &i
	//*j = 1000 //only declaration
	//var j *int // declaration + assignment(j is not empty)
	j := new(int)
	j = &i
	*j = 100
	fmt.Println(j)
	fmt.Println(i)
	name := new(string)
	*name = "golang"
	fmt.Println(*name)
}
*/

/*
package main
import "fmt"
func main() {
	// const pi = 3.14 //implicit typing
	const pi float32 = 3.14 //explicit typing
	fmt.Println(pi)
}
*/

//closer

package main

import "fmt"
func main() {
	number := 10
	fmt.Println(number)
	 
	getInt := func() int {
		fmt.Println("in the function")
		return 10 
	}
	getInt()


}