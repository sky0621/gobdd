package gateway

import (
	gcpgateway "gobdd/adapter/gateway/gcp"
	localgateway "gobdd/adapter/gateway/local"
	"gobdd/adapter/middleware/persistence"
	"gobdd/domain"
	"gobdd/global"
)

func NewNotice(rdbMiddleware persistence.RDBMiddleware) domain.Notice {
	var n domaiin.Notice
	if global.IsLocal() {
		n = localgateway.NewNotice(rdbMiddleware)
	} else {
		n = gcpgateway.NewNotice(rdbMiddleware)
	}
	return n
}
