package guard

import (
	"fmt"
	"testing"
)

type Str string

func (s Str) String() string {
	return string(s)
}

func TestPanicWithError(t *testing.T) {
	tpl := "ERROR"
	err := ErrPanic{Data: fmt.Errorf(tpl)}

	manual := fmt.Sprintf("%s%s", errorPrefix, tpl)
	formatted := err.Error()
	if manual != formatted {
		t.Fatalf("%s expects output to be '%s' but has '%s'", t.Name(), manual, formatted)
	}
}

func TestPanicWithStringer(t *testing.T) {
	tpl := "STRING"
	err := ErrPanic{Data: Str(tpl)}

	manual := fmt.Sprintf("%s%s", errorPrefix, tpl)
	formatted := err.Error()
	if manual != formatted {
		t.Fatalf("%s expects output to be '%s' but has '%s'", t.Name(), manual, formatted)
	}
}

func TestPanicWithOther(t *testing.T) {
	tpl := 42
	err := ErrPanic{Data: tpl}

	manual := fmt.Sprintf("%s%d", errorPrefix, tpl)
	formatted := err.Error()
	if manual != formatted {
		t.Fatalf("%s expects output to be '%s' but has '%s'", t.Name(), manual, formatted)
	}
}
