package operations

import (
	center "github.com/behavioral-ai/center/messaging"
	ops "github.com/behavioral-ai/collective/operations"
	"github.com/behavioral-ai/collective/repository"
	"github.com/behavioral-ai/core/messaging"
)

const (
	NamespaceName = "core:agent/operations/center"
)

var (
	agent *agentT
)

type agentT struct {
	agents *messaging.Exchange
}

func init() {
	repository.RegisterConstructor(NamespaceName, func() messaging.Agent {
		return newAgent()
	})
}

func newAgent() *agentT {
	a := new(agentT)
	a.agents = messaging.NewExchange()
	a.agents.Register(center.NewAgent())
	agent = a
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
	if m.Name == messaging.ConfigEvent {
		ops.Startup(m)
	}
	if m.Name != messaging.ConfigEvent {
		a.agents.Broadcast(m)
	}
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
