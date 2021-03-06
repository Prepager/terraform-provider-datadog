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

// WidgetDisplayType Type of display to use for the request.
type WidgetDisplayType string

// List of WidgetDisplayType
const (
	WIDGETDISPLAYTYPE_AREA WidgetDisplayType = "area"
	WIDGETDISPLAYTYPE_BARS WidgetDisplayType = "bars"
	WIDGETDISPLAYTYPE_LINE WidgetDisplayType = "line"
)

// Ptr returns reference to WidgetDisplayType value
func (v WidgetDisplayType) Ptr() *WidgetDisplayType {
	return &v
}

type NullableWidgetDisplayType struct {
	value *WidgetDisplayType
	isSet bool
}

func (v NullableWidgetDisplayType) Get() *WidgetDisplayType {
	return v.value
}

func (v *NullableWidgetDisplayType) Set(val *WidgetDisplayType) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetDisplayType) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetDisplayType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetDisplayType(val *WidgetDisplayType) *NullableWidgetDisplayType {
	return &NullableWidgetDisplayType{value: val, isSet: true}
}

func (v NullableWidgetDisplayType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetDisplayType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
