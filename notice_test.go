package gobdd_test

import (
	testgateway "gobdd/adapter/gateway/test"
	"gobdd/domain"
	"gobdd/usecase"
	usecasemodel "gobdd/usecase/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Notice", func() {
	Describe("『お知らせ』の登録", func() {
		Context("主体が「システム管理者」である場合", func() {
			const (
				ExpectID = "ef5198df-5c04-42ba-9fbe-2beb2794468a"
			)
			var (
				// 『お知らせ』情報を登録するロジック
				noticeDomain domain.Notice
				// 登録対象の『お知らせ』情報
				noticeForCreateParam *usecasemodel.Notice
			)
			BeforeEach(func() {
				// 期待値をセットできるテスト用のスタブドメインロジックを使うことで、外部サービス接続ロジックを回避したテストが可能
				noticeDomain = &testgateway.NoticeImpl{
					ExpectID:    ExpectID,
					ExpectError: nil,
				}
				noticeForCreateParam = &usecasemodel.Notice{
					Title:       "お知らせ１",
					Text:        "これはお知らせ１です。",
					PublishFrom: 1556636400,
					PublishTo:   1557327599,
				}
			})
			It("表示期間を指定して『お知らせ』を登録できる。", func() {
				id, err := usecase.NewNotice(noticeDomain).Create(noticeForCreateParam)
				Expect(id).To(Equal(ExpectID))
				Expect(err).To(BeNil())
			})
		})
	})
})
