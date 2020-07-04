package echo_test

import (
	echo "my-go/tgpl/1.Tutorial/ch1/echo1"
	"testing"
)

func BenchmarkPlusSign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo.PlusSign([]string{"a", "b", "b", "b", "b", "b", "b"})
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo.Join([]string{"a", "b", "b", "b", "b", "b", "b"})
	}
}
