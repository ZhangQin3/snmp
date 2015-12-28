package snmp

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func Walk(addr, oid string, opts ...string) (map[string]string, error) {
	var options = strings.Join(opts, " ")
	return execWalkCmd("snmpwalk -Lo %s %s %s", options, addr, oid)
}

func execWalkCmd(cmd string, args ...interface{}) (map[string]string, error) {
	ret := make(map[string]string)

	if args != nil {
		cmd = fmt.Sprintf(cmd, args...)
	}
	fmt.Printf("[DEGUG]: %s\n", cmd)

	out, err := exec.Command("cmd.exe", "/c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("[ERROR]: %s\n", out)
	} else {
		fmt.Printf("[INFO]: %s\n", out)
		reg := regexp.MustCompile(`::(\S+) = \S+: (.+)`)
		if m := reg.FindAllSubmatch(out, -1); m != nil {
			for _, v := range m {
				ret[string(v[1])] = string(v[2])
			}
		}
	}

	return ret, err
}
