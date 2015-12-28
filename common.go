package snmp

import (
	"fmt"
)

// set SNMP version to use
func Version(v string) string {
	return fmt.Sprintf("-v%s", v)
}

// set the community string
func Community(c string) string {
	return fmt.Sprintf("-c%s", c)
}

// set authentication protocol (MD5|SHA)
func Protocol(p string) string {
	return fmt.Sprintf("-a%s", p)
}

// set authentication protocol pass phrase
func PassPhrase(p string) string {
	return fmt.Sprintf("-A%s", p)
}

// set the request timeout (in seconds)
func Timeout(t int32) string {
	return fmt.Sprintf("-t%d", t)
}

// set the number of retries
func Retries(r int32) string {
	return fmt.Sprintf("-r%d", r)
}
