package ssh

import (
	"testing"
	"time"
)

const (
	ssh_host = "127.0.0.1"
	ssh_port = 22
	ssh_user = "root"
	ssh_pwd  = "guowei"
)

func TestSShCmd(t *testing.T) {
	client := NewSShClient(ssh_host, ssh_port, ssh_user, ssh_pwd, time.Second*10)
	result, err := client.ExecCmd("ps -ef | grep mysql")
	if err != nil {
		t.Fatalf("execmd error:%v", err)
	}

	t.Logf("execmd success, result:%s", result)
}
