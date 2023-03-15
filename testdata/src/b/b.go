package b

import (
	"fmt"

	// 別名でimport
	me "go.uber.org/multierr"
)

func f() {
	me.Errors(fmt.Errorf("hoge")) // want "multierr.Errors found"
}
