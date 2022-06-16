package httpd

import (
	"github.com/emmvou/wints/logger"
	"github.com/emmvou/wints/schema"
	"github.com/emmvou/wints/session"
)

func streamLog(ex Exchange) error {
	kind := ex.V("k")
	if schema.IsAdminAtLeast(ex.s.RolesAsLevel()) {
		return session.ErrPermission
	}
	in, err := logger.Read(kind)
	return ex.out("text/plain", in, err)
}

func logs(ex Exchange) error {
	if schema.IsAdminAtLeast(ex.s.RolesAsLevel()) {
		return session.ErrPermission
	}
	in, err := logger.Logs()
	return ex.outJSON(in, err)
}
