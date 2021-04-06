package junk

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/sounishnath003/network-scanner/model"
)

var payload []model.RemoteDataStruct

func Junker(ip string, wg *sync.WaitGroup, mutex *sync.RWMutex) {

	defer wg.Done()
	defer mutex.Unlock()

	// commands to be executed on remote machines RPC server communications
	exec1 := "get-computerinfo | select WindowsProductId, WindowsProductName, WindowsRegisteredOwner, WindowsVersion, BiosBIOSVersion, BiosManufacturer, BiosSeralNumber, CsDNSHostName, CsProcessors, TimeZone | convertTo-json -depth 100"
	excmd := "Get-WmiObject -Class Win32_Product -ComputerName " + ip

	out1, err := exec.Command("powershell", exec1).Output()
	warn(err)

	out2, err := exec.Command("powershell", excmd, "|", "select name, vendor, version, InstallDate, caption, IdentifyingNumber, PackageName, ProductID, WarrantyDuration, Description, InstallSource, PackageCode, WarrantyStateDate", "|", "ConvertTo-Json -depth 100").Output()

	if err != nil {
		fmt.Println("*** Executing and connecting with RemoteIP:", ip, "refusing to connect...")
	} else {
		if len(string(out2[:])) < 1 {
			fmt.Println("device with IP:", ip, "is not neccesary to store data... skipping device...")
		} else {
			wg.Add(1)
			go parseData(ip, &out1, &out2, wg)
		}
	}
}

func fn0(ip string) {
	file, er := os.Create("payload.json")
	warn(er)
	dd, err := json.Marshal(payload)
	warn(err)
	file.Write([]byte(dd))
	fmt.Println("### host IP:", ip, " payload and detected softwares are stored in payload.json file....")
}

// writes collected payload to File
func WritePayloadToFile() {
	date := time.Now().Format("01-Jan-2020-15-04-15")
	fileName := "collected-" + date + ".json"

	dd, err := json.Marshal(payload)
	warn(err)
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	warn(err)
	defer f.Close()
	_, err = f.Write(dd)
	warn(err)

	fmt.Println("[SUCCEED]: collected data and payload has been written to the", fileName, "file...")
}

func parseData(ip string, out1, out2 *[]byte, wg *sync.WaitGroup) {
	defer wg.Done()

	// []bytes to JsonStructObject
	n, _ := model.UnmarshalDeviceInfoInterface(out1)                                       // for Deviceinfo
	m, _ := model.UnmarshalMetaDataInterface(out2)                                         // for InstalledSoftware Info
	d := model.RemoteDataStruct{IP: ip, DeviceInformation: n, InstalledSoftwareRecords: m} // actual payload formatting
	// saving in state [...payload]
	payload = append(payload, d)
}

func warn(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
