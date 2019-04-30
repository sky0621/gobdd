package usecasemodel

// ------------------------------------
func NewNotice() Notice {
	return &notice{}
}

// 入力値用『お知らせ』定義
type Notice interface {
}

type notice struct {
}

// ------------------------------------
func NewNoticeCreateResponse() NoticeCreateResponse {
	return &noticeCreateResponse{}
}

// 登録処理の返り値用『お知らせ』定義
type NoticeCreateResponse interface {
}

type noticeCreateResponse struct {
}
