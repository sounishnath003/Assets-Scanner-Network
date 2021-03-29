// To parse and unparse this JSON data, add this code to your project and do:
//
//    metaDataInterface, err := UnmarshalMetaDataInterface(bytes)
//    bytes, err = metaDataInterface.Marshal()

package model

import "encoding/json"

type MetaDataInterface []MetaDataInterfaceElement

func UnmarshalMetaDataInterface(data *[]byte) (MetaDataInterface, error) {
	var r MetaDataInterface
	err := json.Unmarshal(*data, &r)
	return r, err
}

func (r *MetaDataInterface) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MetaDataInterfaceElement struct {
	Name              string      `json:"name"`
	Vendor            string      `json:"vendor"`
	Version           string      `json:"version"`
	InstallDate       string      `json:"InstallDate"`
	Caption           string      `json:"caption"`
	IdentifyingNumber string      `json:"IdentifyingNumber"`
	PackageName       string      `json:"PackageName"`
	ProductID         interface{} `json:"ProductID"`
	WarrantyDuration  interface{} `json:"WarrantyDuration"`
	Description       string      `json:"Description"`
	InstallSource     string      `json:"InstallSource"`
	PackageCode       string      `json:"PackageCode"`
	WarrantyStateDate interface{} `json:"WarrantyStateDate"`
}

// * Device Information Struct

func UnmarshalDeviceInfoInterface(data *[]byte) (DeviceInfoInterface, error) {
	var r DeviceInfoInterface
	err := json.Unmarshal(*data, &r)
	return r, err
}

func (r *DeviceInfoInterface) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeviceInfoInterface struct {
	WindowsProductID       string        `json:"WindowsProductId"`
	WindowsProductName     string        `json:"WindowsProductName"`
	WindowsRegisteredOwner string        `json:"WindowsRegisteredOwner"`
	WindowsVersion         string        `json:"WindowsVersion"`
	BIOSBIOSVersion        []string      `json:"BiosBIOSVersion"`
	BIOSManufacturer       string        `json:"BiosManufacturer"`
	BIOSSeralNumber        string        `json:"BiosSeralNumber"`
	CSDNSHostName          string        `json:"CsDNSHostName"`
	CSProcessors           []CSProcessor `json:"CsProcessors"`
	TimeZone               string        `json:"TimeZone"`
}

type CSProcessor struct {
	Name                      string `json:"Name"`
	Manufacturer              string `json:"Manufacturer"`
	Description               string `json:"Description"`
	Architecture              int64  `json:"Architecture"`
	AddressWidth              int64  `json:"AddressWidth"`
	DataWidth                 int64  `json:"DataWidth"`
	MaxClockSpeed             int64  `json:"MaxClockSpeed"`
	CurrentClockSpeed         int64  `json:"CurrentClockSpeed"`
	NumberOfCores             int64  `json:"NumberOfCores"`
	NumberOfLogicalProcessors int64  `json:"NumberOfLogicalProcessors"`
	ProcessorID               string `json:"ProcessorID"`
	SocketDesignation         string `json:"SocketDesignation"`
	ProcessorType             int64  `json:"ProcessorType"`
	Role                      string `json:"Role"`
	Status                    string `json:"Status"`
	CPUStatus                 int64  `json:"CpuStatus"`
	Availability              int64  `json:"Availability"`
}
