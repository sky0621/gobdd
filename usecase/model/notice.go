package usecasemodel

// ------------------------------------
func NewNotice(title, text string, publishFrom, publishTo int) Notice {
	return &notice{
		title:       title,       // お知らせのタイトル
		text:        text,        // お知らせの文章（現時点はテキストのみサポート）
		publishFrom: publishFrom, // お知らせの掲載開始日時
		publishTo:   publishTo,   // お知らせの掲載終了日時
	}
}

// 入力値用『お知らせ』定義
type Notice interface {
}

type notice struct {
	title       string
	text        string
	publishFrom int
	publishTo   int
}

// ------------------------------------
func NewNoticeCreateResponse(id string, n Notice) NoticeCreateResponse {
	return &noticeCreateResponse{
		id: id,
		n:  n,
	}
}

// 登録処理の返り値用『お知らせ』定義（登録された『お知らせ』をユニークに特定するIDとともに登録された情報を保持）
type NoticeCreateResponse interface {
}

type noticeCreateResponse struct {
	id string
	n  Notice
}
