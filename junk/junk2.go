package junk

import (
	"fmt"
	"os/exec"
	"sync"
)

func Junkerr(ip string, wg *sync.WaitGroup) {

	defer wg.Done()

	excmd := "Get-WmiObject -Class Win32_Product -ComputerName " + ip

	_, err := exec.Command("powershell", excmd, "|", "select name, vendor, version, InstallDate, caption, IdentifyingNumber, PackageName, ProductID, WarrantyDuration, Description, InstallSource, PackageCode, WarrantyStateDate", "|", "ConvertTo-Json -depth 100", "|", "out-file "+ip+".json").Output()

	if err != nil {
		fmt.Println("Error Occurred on ip", ip, "while saving installedSoftware.json file")
		// Warn(err)
		if recover() != nil {
			fmt.Println("!!recoving from empty state to avoid deadlock...")
		}
	} else {
		fmt.Println("File Saved for IP:", ip)
	}

}

// func Warn(err error) {
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
