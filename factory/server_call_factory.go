package factory

import (
	"github.com/Peanoquio/goserverrequest/config"
	"github.com/Peanoquio/goserverrequest/core"
	"github.com/Peanoquio/goserverrequest/interfaces"
)

// NewServerCallUtil creates a new instance of the ServerCall class
func NewServerCallUtil(config *config.ServerCallConfig) interfaces.ServerCallInterface {
	var serverCallInterface interfaces.ServerCallInterface
	serverCall := &core.ServerCall{}
	serverCall.SetConfig(config)
	serverCall.Init()
	serverCallInterface = serverCall
	return serverCallInterface
}
