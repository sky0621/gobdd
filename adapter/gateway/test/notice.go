package testgateway

import domainmodel "gobdd/domain/model"

type NoticeImpl struct {
	ExpectID    string
	ExpectError error
}

func (n *NoticeImpl) Create(noticeModel *domainmodel.Notice) (string, error) {
	// 事前にセットされた期待値を返すだけ
	return n.ExpectID, n.ExpectError
}
