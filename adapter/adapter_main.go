package adapter

type Print interface {
	PrintWeek()
	PrintStrong()
}

type Banner struct {
	Text string
}

func (banner *Banner) Banner(text string) {
	banner.Text = text
}

func (banner *Banner) showWithParen() string {
	return "(" + banner.Text + ")"
}

func (banner *Banner) showWithAster() string {
	return "*" + banner.Text + "*"
}
