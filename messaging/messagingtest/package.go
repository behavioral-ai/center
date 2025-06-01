package messagingtest

import (
	"fmt"
	center "github.com/behavioral-ai/center/messaging"
	"github.com/behavioral-ai/core/messaging"
)

// Comms -
var Comms = func() *center.Communication {
	return &center.Communication{
		Notify: func(msg *messaging.Message) {
			fmt.Printf("%v\n", msg)
		},
		Advise: func(msg *messaging.Message) {
			fmt.Printf("%v\n", msg)
		},
		Request: func(msg *messaging.Message) {
			fmt.Printf("%v\n", msg)
		},
		Respond: func(msg *messaging.Message) {
			fmt.Printf("%v\n", msg)
		},
		Subscribe: func(msg *messaging.Message) {
			fmt.Printf("%v\n", msg)
		},
		CancelSubscription: func(msg *messaging.Message) {
			fmt.Printf("%v\n", msg)
		},
	}
}()
