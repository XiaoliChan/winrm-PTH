package main

import (
	"context"
	"fmt"
	"os"

	"main/winrmntlm"

	"github.com/masterzen/winrm"
)

func main() {
	runExec_winrmntlm("192.168.183.253", 5985, false, "administrator", "e91d2eafde47de62c6c49a012b3a6af1") // works // unsupported action
}

func runExec_winrmntlm(address string, port int, https bool, userName string, password string) {
	endpoint := winrm.NewEndpoint(address, port, https, true, nil, nil, nil, 0)

	params := winrm.DefaultParameters
	enc, _ := winrmntlm.NewEncryption("ntlm", userName, password, endpoint, true) // true is means if password is hash, else false
	params.TransportDecorator = func() winrm.Transporter { return enc }

	client, err := winrm.NewClientWithParameters(endpoint, userName, password, params)
	if err != nil {
		fmt.Println(err)
	}

	exitCode, err := client.RunWithContext(context.Background(), "ipconfig /all", os.Stdout, os.Stderr)
	fmt.Printf("%d\n%v\nn", exitCode, err)
	if err != nil {
		_ = exitCode
		fmt.Println(err)
	} else {
		fmt.Println("Command Test Ok")
	}

	wmiQuery := `select * from Win32_ComputerSystem`
	psCommand := fmt.Sprintf(`$FormatEnumerationLimit=-1;  Get-WmiObject -Query "%s" | Out-String -Width 4096`, wmiQuery)
	stdOut, stdErr, exitCode, _ := client.RunPSWithContext(context.Background(), psCommand)

	fmt.Println(stdOut)
	fmt.Println(stdErr)
	fmt.Println(exitCode)
}
