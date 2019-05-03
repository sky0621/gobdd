package gcpgateway

import (
	"gobdd/adapter/middleware/persistence"
	"gobdd/domain"
	domainmodel "gobdd/domain/model"
)

func NewNotice(rdbMiddleware persistence.RDBMiddleware) domain.Notice {
	return &noticeImpl{rdbMiddleware: rdbMiddleware}
}

type noticeImpl struct {
	rdbMiddleware persistence.RDBMiddleware
}

func (n *noticeImpl) Create(noticeModel *domainmodel.Notice) (string, error) {
	// FIXME: 実DBへの永続化処理を実装
	return "FIXME", nil
}
