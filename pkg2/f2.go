package pkg2

import (
	"fmt"

	"github.com/dattaray-basab/go_mod/pkg2"
	"github.com/dattaray-basab/go_mod/pkg2/pkg2-1"
)

func F2() {
	fmt.Println("pkg2.F2()")
}

func G2() {
	fmt.Println("pkg2.G2()")
	pkg2.F2_1()
}
