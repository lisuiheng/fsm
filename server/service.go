package server

import (
	"github.com/lisuiheng/fsm"
)

// Service handles REST calls to the persistence
type Service struct {
	Logger fsm.Logger
}
