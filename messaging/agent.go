package messaging

import (
	"github.com/behavioral-ai/core/messaging"
)

const (
	NamespaceName = "core:agent/communications/center"
)

var (
	agent *agentT
)

type agentT struct {
	origin Origin
}

func NewAgent() *agentT {
	agent = new(agentT)
	return agent
}

// String - identity
func (a *agentT) String() string { return a.Name() }

// Name - agent identifier
func (a *agentT) Name() string { return NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	// Create origin
	if m.Name == messaging.ConfigEvent {

	}
	if m.Name == messaging.ShutdownEvent {

	}
}

func (a *agentT) notify(m *messaging.Message) {
}

func (a *agentT) advise(m *messaging.Message) {
}

func (a *agentT) request(m *messaging.Message) {
}

func (a *agentT) respond(m *messaging.Message) {
}

func (a *agentT) subscribe(m *messaging.Message) {
}

func (a *agentT) cancel(m *messaging.Message) {
}

func (a *agentT) trace(name, task, observation, action string) {
}

/*
func (a *agentT) configure(m *messaging.Message) {
	//ur := messaging.messaging.ConfigMapContent(m)
	//if cfg == nil {
	//	messaging.Reply(m, messaging.ConfigEmptyStatusError(a), a.Uri())
	//}
	// configure
	//messaging.Reply(m, messaging.StatusOK(), a.Uri())
	agents.Broadcast(m)
}


*/
