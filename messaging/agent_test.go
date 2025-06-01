package messaging

import "fmt"

func ExampleNewAgent() {
	a := NewAgent()

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [core:agent/communications/center]

}
