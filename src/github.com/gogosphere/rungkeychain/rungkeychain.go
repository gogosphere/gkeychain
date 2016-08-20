package main

import (
	"fmt"
	"os"
	"os/exec"

	gk "github.com/gogosphere/gkeychain"
)

func main() {
	mea, err := exec.Command("whoami").Output()
	if err != nil {
		println(err.Error())
		os.Exit(2)
	}
	me := fmt.Sprintf(string(mea)[0 : len(mea)-1])
	fmt.Println(gk.Secret(me))
}
