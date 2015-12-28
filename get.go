package snmp

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func Get(addr, oid string, opts ...string) (string, error) {
	var options = strings.Join(opts, " ")
	out, err := execCmd("snmpget -Lo %s %s %s", options, addr, oid)
	return string(out), err
}

func execCmd(cmd string, args ...interface{}) ([]byte, error) {
	if args != nil {
		cmd = fmt.Sprintf(cmd, args...)
	}
	fmt.Printf("[DEGUG]: %s\n", cmd)

	out, err := exec.Command("cmd.exe", "/c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("[ERROR]: %s\n", out)
	} else {
		fmt.Printf("[INFO]: %s\n", out)
		reg := regexp.MustCompile(`= \S+: (.+)`)
		if m := reg.FindSubmatch(out); m != nil {
			out = m[1]
		}
	}

	return out, err
}
