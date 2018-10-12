package server

import (
	"github.com/lisuiheng/fsm"
	"github.com/lisuiheng/fsm/dao"
)

// Service handles REST calls to the persistence
type Service struct {
	Logger    fsm.Logger
	client    *dao.Client
	UseAuth   bool
	Signature string
}
