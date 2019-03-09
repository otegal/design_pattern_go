package adapter

type PrintBannerByDelegation struct {
	banner *Banner
}

func NewPrintBannerByDelegation(text string) *PrintBannerByDelegation {
	return &PrintBannerByDelegation{&Banner{Text: text}}
}

func (printBanner *PrintBannerByDelegation) PrintWeek() string {
	return printBanner.banner.showWithParen()
}

func (printBanner *PrintBannerByDelegation) PrintStrong() string {
	return printBanner.banner.showWithAster()
}
