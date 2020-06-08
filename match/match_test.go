package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestFindAllStringSubMatch(t *testing.T) {
	x := "var hq_str_sh000001=\"上证指数,2796.,2789.2537,2774.4822,2826.9068,2715.2213,0,0,188272494,192721831750,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2020-03-17,10:57:46,00,\";\n"
	reg := regexp.MustCompile(`^var hq_str_([\w|\d]+)=\"(.*)\";`)
	match := reg.FindAllStringSubmatch(x, -1)
	fmt.Println(match[0][2])
}

func TestMatchString(t *testing.T) {
	x := "var hq_str_sh000001=\"上证指数,2796.,2789.2537,2774.4822,2826.9068,2715.2213,0,0,188272494,192721831750,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2020-03-17,10:57:46,00,\";\n"
	reg := regexp.MustCompile(`^var hq_str_sh000001="`)
	t.Log(reg.MatchString(x))
}
