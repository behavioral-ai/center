package operations

import (
	"github.com/behavioral-ai/core/messaging"
)

// Startup -
// TODO: Set origin for access logging.
//
//	Configure messaging agent with service name
//	Update collective with service name
func Startup(msg *messaging.Message) {
	agent.Message(msg)
	agent.Message(messaging.StartupMessage)
}
