package gcpgateway

import domainmodel "gobdd/domain/model"

type NoticeImpl struct {
	// FIXME: Cloud SQLへのコネクションを保持
}

func (n *NoticeImpl) Create(noticeModel *domainmodel.Notice) (string, error) {
	// FIXME: 実DBへの永続化処理を実装
	return "FIXME", nil
}
