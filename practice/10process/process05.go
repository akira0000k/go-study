package main
import (
	"fmt"
	"os"
	"strings"
)
/*
   subject : Environmen Variables
*/
func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		//fmt.Println(e)
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], "=", pair[1])
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/10process/" -*-
// Compilation started at Fri Oct 15 21:14:39
//  
// go run process05.go 
// FOO: 1
// BAR: 
//  
// TERM = dumb
// TERMCAP = 
// COLUMNS = 117
// INSIDE_EMACS = 27.2,compile
// LANG = ja_JP.UTF-8
// XPC_FLAGS = 0x0
// XPC_SERVICE_NAME = application.org.gnu.Emacs.12997267185.12997267231
// TMPDIR = /var/folders/5m/29zwdxmj52q7klt6qnkzxn_40000gp/T/
// __CF_USER_TEXT_ENCODING = 0x1F6:0x1:0xE
// HOME = /Users/Akira
// SHELL = /bin/zsh
// SSH_AUTH_SOCK = /private/tmp/com.apple.launchd.TWZluEUVS0/Listeners
// PATH = /usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/Applications/Emacs.app/Contents/MacOS/bin-x86_64-10_14:/Applications/Emacs.app/Contents/MacOS/libexec-x86_64-10_14
// LOGNAME = Akira
// DISPLAY = /private/tmp/com.apple.launchd.TeIT7mxlPf/org.xquartz:0
// COMMAND_MODE = unix2003
// __CFBundleIdentifier = org.gnu.Emacs
// USER = Akira
// SHLVL = 0
// PWD = /Users/Akira/go/src/practice/10process
// OLDPWD = /Users/Akira/go/src/practice/10process
// _ = /usr/local/bin/go
// FOO = 1
//  
// Compilation finished at Fri Oct 15 21:14:39
