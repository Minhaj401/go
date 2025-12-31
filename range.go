package main

import "fmt"

// iterating over data structures
func main() {
	m := map[string]string{"fname": "john", "lname": "doe"}

for k, v := range m {
	fmt.Println(k, v)
}

	

}