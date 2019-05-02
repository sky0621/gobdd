package domain

import (
	domainmodel "gobdd/domain/model"
)

type Notice interface {
	Create(noticeModel *domainmodel.Notice) (string, error)
}
