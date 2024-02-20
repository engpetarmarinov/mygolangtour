package main

import (
	"fmt"
)

func main() {
	//append
	helloSlice := []byte("hello ")
	fmt.Printf("%v\n", string(append(helloSlice, "world"...))) //hello world
	fmt.Printf("%v\n", cap(helloSlice))                        //6
	//clear
	clear(helloSlice)
	fmt.Printf("%v\n", helloSlice) //[0 0 0 0 0 0]
	//make
	c := make(chan int)

	go func() {
		if v, ok := <-c; !ok {
			fmt.Printf("%v", ok)
		} else {
			fmt.Printf("%v", v)
		}

	}()

	c <- 1
	//only the sender invokes close
	close(c)

	copyDest := []string{"one", "two"}
	copySrc := []string{"cat"}
	//copy
	copy(copyDest, copySrc)
	fmt.Printf("%v\n", copyDest) //[cat two]
	copyDest[0] = "dog"
	fmt.Printf("%v\n", copySrc)  //[cat]
	fmt.Printf("%v\n", copyDest) //[dog two]

	var copyDest2 []string
	copy(copyDest2, copySrc)
	fmt.Printf("%v\n", copyDest2) //[]

	deleteMap := map[string]string{
		"one": "cat",
		"two": "dog",
	}
	//delete
	delete(deleteMap, "one")
	fmt.Printf("%v\n", deleteMap) //map[two:dog]

	//complex
	cn := complex(2.2, 2.3)
	fmt.Printf("%v\n", cn) //1(2.2+2.3i)

	//real
	fmt.Printf("%v\n", real(cn)) //2.2

	//imag
	fmt.Printf("%v\n", imag(cn)) //2.3 - imag returns the imaginary part of the complex number

	//min and max
	fmt.Printf("%v\n", min(1, 2, 3)) //1
	fmt.Printf("%v\n", max(1, 2, 3)) //3

	//new built-in function allocates memory, the value returned is a pointer to a newly allocated zero value of that type
	fmt.Printf("%v\n", new(int)) //0xc00000a140

	//panic and recover
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recovered: %v\n", r)
			}
		}()

		panic("panicus")
	}()

	fmt.Println("return to normal execution")

}
