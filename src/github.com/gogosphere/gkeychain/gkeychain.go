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
