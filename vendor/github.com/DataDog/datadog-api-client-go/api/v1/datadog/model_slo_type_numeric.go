/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// SLOTypeNumeric A numeric representation of the type of the service level objective (`0` for monitor, `1` for metric). Always included in service level objective responses. Ignored in create/update requests.
type SLOTypeNumeric int32

// List of SLOTypeNumeric
const (
	SLOTYPENUMERIC_MONITOR SLOTypeNumeric = 0
	SLOTYPENUMERIC_METRIC  SLOTypeNumeric = 1
)

// Ptr returns reference to SLOTypeNumeric value
func (v SLOTypeNumeric) Ptr() *SLOTypeNumeric {
	return &v
}

type NullableSLOTypeNumeric struct {
	value *SLOTypeNumeric
	isSet bool
}

func (v NullableSLOTypeNumeric) Get() *SLOTypeNumeric {
	return v.value
}

func (v *NullableSLOTypeNumeric) Set(val *SLOTypeNumeric) {
	v.value = val
	v.isSet = true
}

func (v NullableSLOTypeNumeric) IsSet() bool {
	return v.isSet
}

func (v *NullableSLOTypeNumeric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLOTypeNumeric(val *SLOTypeNumeric) *NullableSLOTypeNumeric {
	return &NullableSLOTypeNumeric{value: val, isSet: true}
}

func (v NullableSLOTypeNumeric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLOTypeNumeric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
