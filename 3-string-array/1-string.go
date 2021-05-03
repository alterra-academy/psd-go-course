import (
	"fmt"
	"strings"
)
   
const (
	str    = "something"
	substr = "some"
)
 
func main() {
	// 1. len string
	sentence := "Hello";
	lenSentence := len(sentence)
	fmt.Println(lenSentence)
   
	// 2. compare string
	str1 := "abc"
	str2 := "abd"
	fmt.Println(str1 == str2)
   
	// 3. Contains
	res := strings.Contains(str, substr)
	fmt.Println(res) // true
}
   