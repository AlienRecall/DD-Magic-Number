package main

import (
	"testing"
)

func TestMagicNumber(t *testing.T) {
	ddmn := &DatadomeMagicNumber{
		language:  "en-GB",
		languages: []string{"en-GB", "en-US", "en", "it"},
		r:         ".AmpSWXKxiOOE-dZx0XOos9WBjFdz5sINnD4ma-tBP2mzH3wZlSTjDWQ-vS9-W2Vif7SLH8o~CKpsVhkWh2apXl03PnDfJl7i0N_nBvDiZg_9JhGz0biH6qc~uU8gu.g",
		t:         10,
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.85 Safari/537.36",
	}
	magicNumber := ddmn.Generate()
	t.Logf("%d", magicNumber)
}
