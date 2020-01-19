package webcmds

import (
	"fmt"

	"github.com/ibraimgm/libcmd"

	"dfis-utils/pkg/webutils"
)

func GetMails(cmd *libcmd.Cmd) error {
	url := cmd.Operand("url")
	if url == "" {
		fmt.Println("URL is requrired")
	}
	webutils.Getem(url)
	return nil
}

func GetKeyword(cmd *libcmd.Cmd) error {
	url := cmd.Operand("url")
	if url == "" {
		fmt.Println("URL is requrired")
	}

	keyword := cmd.GetString("keyword")
	return webutils.Keysearch(url, *keyword)
}

func GetHeads(cmd *libcmd.Cmd) error {
	url := cmd.Operand("url")
	if url == "" {
		fmt.Println("URL is requrired")
	}
	return webutils.Heads(url)
}

func SetCookie(cmd *libcmd.Cmd) error {
	url := cmd.Operand("url")
	if url == "" {
		fmt.Println("URL is requrired")
	}

	method := cmd.GetString("method")
	key := cmd.GetString("key")
	value := cmd.GetString("value")

	return webutils.Setcook(url, *key, *value, *method)
}

func GetComments(cmd *libcmd.Cmd) error {
	url := cmd.Operand("url")
	if url == "" {
		fmt.Println("URL is requrired")
	}
	return webutils.Getcoms(url)
}

func GetUnlisted(cmd *libcmd.Cmd) error {
	url := cmd.Operand("url")
	if url == "" {
		fmt.Println("URL is requrired")
	}
	wordlist := cmd.GetString("wordlist_file")
	return webutils.Getunlist(url, *wordlist)
}
