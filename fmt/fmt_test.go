package fmt

import (
	"fmt"
	"testing"
)

func TestSprintf(t *testing.T) {
	yearRate := fmt.Sprintf("%.2f", (2777.4181-3050.12)/3050.12*100)
	fmt.Println(yearRate, 2777.4181-3050.12)
}
