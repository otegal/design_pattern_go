package adapter

type PrintBannerByInheritance struct {
	*Banner
}

// Goには継承がないのでPrintBannerはBannerを含む構造体にして
// 擬似的にBannerを親クラス、PrintBannerを子クラスのようにして利用する
func NewPrintBannerByInheritance(text string) *PrintBannerByInheritance {
	return &PrintBannerByInheritance{&Banner{Text: text}}
}

func (printBanner *PrintBannerByInheritance) PrintWeek() string {
	return printBanner.showWithParen()
}

func (printBanner *PrintBannerByInheritance) PrintStrong() string {
	return printBanner.showWithAster()
}
