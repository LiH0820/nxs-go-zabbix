package zabbix

import (
	"os"
	"testing"
)

const (
	ZABBIX_HOST     = "http://10.130.13.169:8080/api_jsonrpc.php"
	ZABBIX_USERNAME = "Admin"
	ZABBIX_PASSWORD = "zabbix"
)

func GetZabbixContext(t *testing.T) *Context {

	zbxHost := os.Getenv("ZABBIX_HOST")
	if zbxHost == "" {
		//t.Fatal("Login error: undefined env var `ZABBIX_HOST`")
		zbxHost = ZABBIX_HOST
	}

	zbxUsername := os.Getenv("ZABBIX_USERNAME")
	if zbxUsername == "" {
		//t.Fatal("Login error: undefined env var `ZABBIX_USERNAME`")
		zbxUsername = ZABBIX_USERNAME
	}

	zbxPassword := os.Getenv("ZABBIX_PASSWORD")
	if zbxPassword == "" {
		//t.Fatal("Login error: undefined env var `ZABBIX_PASSWORD`")
		zbxPassword = ZABBIX_PASSWORD
	}

	z, err := NewContext(zbxHost, zbxUsername, zbxPassword)
	if err != nil {
		t.Fatal("Login error: ", err)
	} else {
		t.Logf("Login: success")
	}

	return z
}
