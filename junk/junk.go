package junk

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/sounishnath003/network-scanner/model"
)

type RemoteDataStruct struct {
	IP                string                    `json:"ip"`
	InstalledSoftware []model.MetaDataInterface `json:"installedSoftwareDetails"`
}

var payload []RemoteDataStruct

func Junker(ip string, wg *sync.WaitGroup) {

	defer wg.Done()

	excmd := "Get-WmiObject -Class Win32_Product -ComputerName " + ip

	out, err := exec.Command("powershell", excmd, "|", "select name, vendor, version, InstallDate, caption, IdentifyingNumber, PackageName, ProductID, WarrantyDuration, Description, InstallSource, PackageCode, WarrantyStateDate", "|", "ConvertTo-Json -depth 100").Output()

	if err != nil {
		fmt.Println("*** Executing and connecting with RemoteIP:", ip, "refusing to connect...")
	} else {
		if len(string(out[:])) < 1 {
			fmt.Println("device with IP:", ip, "is not neccesary to store data... skipping device...")
		} else {
			wg.Add(1)
			go parseData(out, wg)
		}
	}
}

func parseData(out []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	file, _ := os.Create("192.168.0.3.json")
	file.Write(out)

}
