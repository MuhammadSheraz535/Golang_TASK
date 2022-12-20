// Program to print numbers for natural numbers 1 + 2 + 3 + ... +n

package main
import "fmt"

func main() {
  
 for i:=0; i<5; i++ {
	for j:=0; j<i+1; j++ {

		fmt.Print(" * ")

	}
fmt.Println()
  
}

}