// +build windows

package shell

import (
	"bytes"
	"os/exec"
)

// 阻塞式执行shell命令，等待执行完毕并返回结果
func ExecShell(s string) (string, error) {
	cmd := exec.Command("cmd", "/C", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return ``, err
	}
	return out.String(), err
}
