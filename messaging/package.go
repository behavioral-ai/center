package messaging

import "github.com/behavioral-ai/core/messaging"

// Notifications from agents:
//  1. Status - some error that needs triage or awareness
//  2. Event - some event needs to be raised for a single member, or member listeners.

// Communication - in the real world
type Communication struct {
	Notify             func(msg *messaging.Message)
	Advise             func(msg *messaging.Message)
	Request            func(msg *messaging.Message)
	Respond            func(msg *messaging.Message)
	Subscribe          func(msg *messaging.Message)
	CancelSubscription func(msg *messaging.Message)
}

// Comms -
var Comms = func() *Communication {
	return &Communication{
		Notify: func(msg *messaging.Message) {
			agent.notify(msg)
		},
		Advise: func(msg *messaging.Message) {
			agent.advise(msg)
		},
		Request: func(msg *messaging.Message) {
			agent.request(msg)
		},
		Respond: func(msg *messaging.Message) {
			agent.respond(msg)
		},
		Subscribe: func(msg *messaging.Message) {
			agent.subscribe(msg)
		},
		CancelSubscription: func(msg *messaging.Message) {
			agent.cancel(msg)
		},
	}
}()
