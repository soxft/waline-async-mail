package engine

type Platform string

const (
	PlatformSmtp     Platform = "smtp"
	PlatformAliyun   Platform = "aliyun"
	PlatformBark     Platform = "bark"
	PlatformTelegram Platform = "telegram"
)

type Mail struct {
	Subject   string
	Content   string
	ToAddress string
	Typ       string // 邮件类型
}

type OwnerArgs struct {
	Permalink string `args:"permalink"`  // 文章链接
	SiteTitle string `args:"site_title"` // 网站标题
	Author    string `args:"author"`     // 作者
	Text      string `args:"text"`       // 评论内容
	Ip        string `args:"ip"`
	Mail      string `args:"mail"`
	Time      string `args:"time"`
	Status    string `args:"status"`
}

type GuestArgs struct {
	AuthorP   string `args:"author_p"`   // 我的昵称
	Permalink string `args:"permalink"`  // 文章链接
	Author    string `args:"author"`     // 回复者昵称
	TextP     string `args:"text_p"`     // 我的评论
	Text      string `args:"text"`       // 别人回复的评论
	SiteTitle string `args:"site_title"` // 网站标题
}
