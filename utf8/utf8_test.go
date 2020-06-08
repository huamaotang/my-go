package utf8

import (
	"testing"
	"unicode/utf8"
)

func TestName(t *testing.T) {
	t.Log(utf8.RuneCountInString("abc"), utf8.RuneCountInString("123"), utf8.RuneCountInString("汤华茂"))
}
