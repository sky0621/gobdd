package usecasemodel

// 入力値用『お知らせ』定義
type Notice struct {
	Title       string // お知らせのタイトル
	Text        string // お知らせの文章（現時点はテキストのみサポート）
	PublishFrom int    // お知らせの掲載開始日時
	PublishTo   int    // お知らせの掲載終了日時
}
