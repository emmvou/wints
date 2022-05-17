package httpd

import (
	"github.com/emmvou/wints/logger"
	"github.com/emmvou/wints/schema"
	"github.com/emmvou/wints/session"
)

func streamLog(ex Exchange) error {
	kind := ex.V("k")
	if ex.s.Me().Role.Level() < schema.AdminLevel {
		return session.ErrPermission
	}
	in, err := logger.Read(kind)
	return ex.out("text/plain", in, err)
}

func logs(ex Exchange) error {
	if ex.s.Me().Role.Level() < schema.AdminLevel {
		return session.ErrPermission
	}
	in, err := logger.Logs()
	return ex.outJSON(in, err)
}
