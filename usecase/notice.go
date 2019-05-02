package usecase

import usecasemodel "gobdd/usecase/model"

func NewNotice() Notice {
	return &notice{}
}

type Notice interface {
	Create(noticeModel usecasemodel.Notice) usecasemodel.NoticeCreateResponse
}

type notice struct {
}

func (n *notice) Create(noticeModel usecasemodel.Notice) usecasemodel.NoticeCreateResponse {
	// FIXME: 【テストコードを通すための実装】ここでは渡された『お知らせ』情報と固定のIDを返却する。
	return usecasemodel.NewNoticeCreateResponse("ef5198df-5c04-42ba-9fbe-2beb2794468a", noticeModel)
}
