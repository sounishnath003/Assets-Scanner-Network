package junk

import (
	"fmt"
	"log"
	"os/exec"
)

func Junker() {
	bytes, err := exec.Command("powershell", "Get-WmiObject -Class Win32_Product -ComputerName 192.168.0.3", "|", "select name, vendor, version, InstallDate, caption, IdentifyingNumber, PackageName, ProductID, WarrantyDuration, Description, InstallSource, PackageCode, WarrantyStateDate", "|", "ConvertTo-Json -depth 100", "|", "out-file outt.json").Output()

	if err != nil {
		panic(err)
	} else {
		fmt.Println("File Saved", bytes)
	}

	Warn(err)
}

func Warn(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
