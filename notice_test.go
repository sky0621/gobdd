package gobdd_test

import (
	"gobdd/usecase"
	usecasemodel "gobdd/usecase/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Notice", func() {
	Describe("『お知らせ』の登録", func() {
		Context("主体が「システム管理者」である場合", func() {
			var (
				// 登録対象の『お知らせ』情報
				noticeForCreateParam usecasemodel.Notice
				// 『お知らせ』の登録結果
				expectedResult usecasemodel.NoticeCreateResponse
			)
			BeforeEach(func() {
				noticeForCreateParam = usecasemodel.NewNotice("お知らせ１", "これはお知らせ１です。", 1556636400, 1557327599)
				expectedResult = usecasemodel.NewNoticeCreateResponse("ef5198df-5c04-42ba-9fbe-2beb2794468a", noticeForCreateParam)
			})
			It("表示期間を指定して『お知らせ』を登録できる。", func() {
				Expect(usecase.NewNotice().Create(noticeForCreateParam)).To(Equal(expectedResult))
			})
		})
	})
})
