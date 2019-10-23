package main 

import "fmt"

func counter() func() int{
	var i = 0
	return func() int{
		i++
		return i
	}
}

func main(){
   getCount := counter()

   for i := 0; i<5; i++{
	   fmt.Printf("Counter: %d\n",getCount())
   }
}