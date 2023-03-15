package c

import (
	"c/multierr"
)

func f() {
	multierr.Errors() // 別パッケージのErrorsは検出しない
}
