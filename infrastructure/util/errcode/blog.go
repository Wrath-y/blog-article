package errcode

var (
	ArticleNotExists    = &ErrCode{40200, "文章不存在", ""}
	ArticleInsertFailed = &ErrCode{40201, "文章添加失败", ""}
)
