package model

import "encoding/json"

type MetaDataInterface struct {
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

func UnmarshalMetaDataInterface(data []byte) (MetaDataInterface, error) {
	var r MetaDataInterface
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MetaDataInterface) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
