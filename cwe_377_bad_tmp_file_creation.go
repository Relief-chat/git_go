package samples

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// bad
	err := ioutil.WriteFile("/tmp/demo2", []byte("This is some data"), 0644)
	if err != nil {
		fmt.Println("Error while writing!")
	}
}
func main_good() {
	// ok 
	err := ioutil.Tempfile("/tmp", "my_temp")
	if err != nil {
		fmt.Println("Error while writing!")
	}
}
