package usecase

import usecasemodel "gobdd/usecase/model"

func NewNotice() Notice {
	return &notice{}
}

type Notice interface {
	Create(usecasemodel.Notice) usecasemodel.NoticeCreateResponse
}

type notice struct {
}

func (n *notice) Create(usecasemodel.Notice) usecasemodel.NoticeCreateResponse {
	// FIXME: ユースケースに合わせて実装
	return nil
}
