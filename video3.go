/*
package main
import "fmt"
func main() {
	fmt.Println("Hello World")
}
*/

//datatypes

/*
package main
import "fmt"
func main() {
	var data int
	data = 10
	data1 := 100
	data = 1000
	fmt.Println(data)
	fmt.Println(data1)
	name := `rahul`
	fmt.Println(name)
}
*/

//packages/controls
/*
 package main
 import "fmt"
 func main() {
	var input int
	fmt.Println("Enter a number : ")
	fmt.Scanln(&input)
	if input%2 == 0{
		fmt.Println("even")
	}else {
		fmt.Println("odd")
	}
 }
*/

//switch
/*
 package main
 import "fmt"
 func main() {
	data := 21
	switch data {
	case 10:
		data = 100
		fmt.Println(data)
	case 100:
		data = 103
		fmt.Println(data)
		fallthrough //execute the next case also
	case 11:
		data = 104
		fmt.Println(data)
	case 15:
		data = 1002
		fmt.Println(data)
	default:
		data = 10909
		fmt.Println(data)
	}
 }
*/

// loops

package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}

}