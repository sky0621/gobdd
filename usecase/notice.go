package usecase

import (
	"gobdd/domain"
	domainmodel "gobdd/domain/model"
	usecasemodel "gobdd/usecase/model"

	"github.com/google/uuid"
)

func NewNotice(noticeDomain domain.Notice) Notice {
	return &notice{noticeDomain: noticeDomain}
}

type Notice interface {
	Create(noticeModel *usecasemodel.Notice) (string, error)
}

type notice struct {
	noticeDomain domain.Notice
}

func (n *notice) Create(noticeModel *usecasemodel.Notice) (string, error) {
	// ドメインモデルへの変換（ユースケース層独自の構造からドメイン層独自の構造への変換（例：日時の持ち方や「姓」と「名」別持ちから「姓名」等））
	domainModel := &domainmodel.Notice{
		ID:          uuid.New().String(),
		Title:       noticeModel.Title,
		Text:        noticeModel.Text,
		PublishFrom: noticeModel.PublishFrom,
		PublishTo:   noticeModel.PublishTo,
	}
	return n.noticeDomain.Create(domainModel)
}
