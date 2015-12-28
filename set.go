package snmp

import (
	"fmt"
	"os/exec"
	"strings"
)

func Set(addr, oid string, value interface{}, opts ...string) error {
	var options = strings.Join(opts, " ")
	return execSetCmd("snmpset -Lo %s %s %s = %v", options, addr, oid, value)
}

func execSetCmd(cmd string, args ...interface{}) error {
	if args != nil {
		cmd = fmt.Sprintf(cmd, args...)
	}
	fmt.Printf("[DEGUG]: %s\n", cmd)

	out, err := exec.Command("cmd.exe", "/c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("[ERROR]: %s\n", out)
	}

	return err
}
