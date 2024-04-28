package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"

	"main/winrm"
	"main/winrmntlm"

	"golang.org/x/crypto/md4"
	"golang.org/x/text/encoding/unicode"
)

func main() {
	plaintext := "111qqq..."
	plaintext_, _ := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder().Bytes([]byte(plaintext))
	ntlm := hex.EncodeToString(hashMD4(plaintext_))

	/*
		or
		ntlm := "e91d2eafde47de62c6c49a012b3a6af1"
	*/

	runExec_winrmntlm("192.168.1.128", 5985, false, "administrator", ntlm) // works // unsupported action
}

func runExec_winrmntlm(address string, port int, https bool, userName, ntlm string) {
	endpoint := winrm.NewEndpoint(address, port, https, true, nil, nil, nil, 0)

	params := winrm.DefaultParameters
	enc, _ := winrmntlm.NewEncryption("ntlm", userName, ntlm, endpoint)
	params.TransportDecorator = func() winrm.Transporter { return enc }
	client, err := winrm.NewClientWithParameters(endpoint, userName, "111qqq...", params)

	//client, err := winrm.NewClient(endpoint, userName, ntlm)
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

func hashMD4(b []byte) []byte {
	md4 := md4.New()
	md4.Write(b)

	return md4.Sum(nil)
}
