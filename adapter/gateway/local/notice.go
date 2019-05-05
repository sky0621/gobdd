package localgateway

import (
	middlewaremodel "gobdd/adapter/middleware/model"
	"gobdd/adapter/middleware/persistence"
	"gobdd/domain"
	domainmodel "gobdd/domain/model"
	"time"
)

func NewNotice(rdbMiddleware persistence.RDBMiddleware) domain.Notice {
	return &noticeImpl{rdbMiddleware: rdbMiddleware}
}

type noticeImpl struct {
	rdbMiddleware persistence.RDBMiddleware
}

func (n *noticeImpl) Create(noticeModel *domainmodel.Notice) (string, error) {
	m := &middlewaremodel.Notice{
		ID:          noticeModel.ID,
		Title:       noticeModel.Title,
		Text:        noticeModel.Text,
		PublishFrom: noticeModel.PublishFrom,
		PublishTo:   noticeModel.PublishTo,
		CreatedAt:   time.Now(),
	}

	if err := n.rdbMiddleware.Create(m); err != nil {
		return "", err
	}
	return noticeModel.ID, nil
}
