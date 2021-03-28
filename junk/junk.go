package junk

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/sounishnath003/network-scanner/model"
)

type RemoteDataStruct struct {
	IP                       string      `json:"ip"`
	InstalledSoftwareRecords interface{} `json:"installedSoftwareRecords"`
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
			go parseData(ip, &out, wg)
		}
	}

	file, _ := os.Create("payload.json")
	dd, _ := json.Marshal(payload)
	file.Write([]byte(dd))

	fmt.Println("### host IP:", ip, " payload and detected softwares are stored in payload.json file....")
}

func parseData(ip string, out *[]byte, wg *sync.WaitGroup) {
	defer wg.Done()

	m, _ := model.UnmarshalMetaDataInterface(out)              // []bytes to JsonStructObject
	d := RemoteDataStruct{IP: ip, InstalledSoftwareRecords: m} // actual payload formatting
	// saving in state [...payload]
	payload = append(payload, d)
}
