package components

import (
	"github.com/1Panel-dev/1Panel/backend/utils/common"
	"strings"
)

const DefaultServer = "default_server"

type ServerListen struct {
	Bind          string
	DefaultServer string
	Parameters    []string
	Comment       string
}

func NewServerListen(params []string) *ServerListen {
	server := &ServerListen{
		Parameters: []string{},
	}
	for _, param := range params {
		if isBind(param) {
			server.Bind = param
		} else if param == DefaultServer {
			server.DefaultServer = DefaultServer
		} else {
			server.Parameters = append(server.Parameters, param)
		}
	}
	return server
}

func isBind(param string) bool {
	if common.IsNum(param) {
		return true
	}
	if strings.Contains(param, "*") || strings.Contains(param, ":") || strings.Contains(param, ".") {
		return true
	}
	return false
}

func (sl *ServerListen) GetName() string {
	return "listen"
}

func (sl *ServerListen) GetBlock() IBlock {
	return nil
}

func (sl *ServerListen) GetParameters() []string {
	params := []string{sl.Bind}
	params = append(params, sl.DefaultServer)
	params = append(params, sl.Parameters...)
	return params
}

func (sl *ServerListen) GetComment() string {
	return sl.Comment
}

func (sl *ServerListen) AddDefaultServer() {
	sl.DefaultServer = DefaultServer
}

func (sl *ServerListen) RemoveDefaultServe() {
	sl.DefaultServer = ""
}