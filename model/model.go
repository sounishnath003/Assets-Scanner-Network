// This file was generated from JSON Schema using quicktype, do not modify it directly.
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
