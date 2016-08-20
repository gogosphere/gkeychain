Example to pull keychain information as a library.

It requires  `github.com/iwat/go-seckeychain` for the C bindings to OSX security.

```
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

```

```
package gkeychain

import (
	"os"

	"github.com/iwat/go-seckeychain"
)

// Secret should return the password of the object
func Secret(account string) string {

	// general-password is the Name in the keychain and "me" is the account name
	password, err := seckeychain.FindGenericPassword("general-password", account)
	if err != nil {
		println(err.Error())
		os.Exit(2)
	}

	return password
}
}
```
