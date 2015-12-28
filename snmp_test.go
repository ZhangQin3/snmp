package snmp

import (
	"testing"
)

func TestSnmpGet(t *testing.T) {
	Get("10.89.156.1", "arrisCmDoc30AccessTelnetPassword.0")
}

func TestSnmpGetWithOptions(t *testing.T) {
	Get("10.89.156.1", "arrisCmDoc30AccessTelnetPassword.0", Version("2c"), Community("testing"), Timeout(5), Retries(3))
}

func TestSnmpGetWithStringOptions(t *testing.T) {
	Get("10.89.156.1", "arrisCmDoc30AccessTelnetPassword.0", "-v2c", "-ctesting", Timeout(5), Retries(3))
}

func TestSnmpWalk(t *testing.T) {
	Walk("10.89.156.1", "ifIndex")
}

func TestSnmpset(t *testing.T) {
	Set("10.89.156.1", "arrisCmDoc30AccessTelnetPassword.0", 12345678901, "-v2c")
}

func TestSnmpsetString(t *testing.T) {
	Set("10.89.156.1", "arrisCmDoc30AccessTelnetPassword.0", "12345678901")
}

func TestSnmpsetOverlong(t *testing.T) {
	Set("10.89.156.1", "arrisCmDoc30AccessTelnetPassword.0", 1234567890120)
}

func TestSnmpsetInt(t *testing.T) {
	Set("10.89.156.1", "arrisCmDoc30AccessTelnetPassword.0", 123456)
}
