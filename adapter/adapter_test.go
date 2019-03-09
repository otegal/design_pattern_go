package adapter

import (
	"testing"
)

func TestAdapterInheritance(t *testing.T) {
	assert := "sample text"
	printBanner := NewPrintBannerByInheritance(assert)
	if printBanner.PrintWeek() != "("+assert+")" {
		t.Errorf("Expect result to %s, but %s", "("+assert+")", printBanner.PrintWeek())
	}
	if printBanner.PrintStrong() != "*"+assert+"*" {
		t.Errorf("Expect result to %s, but %s", "*"+assert+"*", printBanner.PrintWeek())
	}
}

func TestAdapterDelegation(t *testing.T) {
	assert := "sample text"
	printBanner := NewPrintBannerByDelegation(assert)
	if printBanner.PrintWeek() != "("+assert+")" {
		t.Errorf("Expect result to %s, but %s", "("+assert+")", printBanner.PrintWeek())
	}
	if printBanner.PrintStrong() != "*"+assert+"*" {
		t.Errorf("Expect result to %s, but %s", "*"+assert+"*", printBanner.PrintWeek())
	}
}
