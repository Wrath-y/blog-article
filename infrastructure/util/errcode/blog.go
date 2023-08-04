package errcode

var (
	ArticleNotExists  = &ErrCode{40200, "文章不存在", ""}
	ArticleFindFailed = &ErrCode{40201, "文章列表获取失败", ""}
)
