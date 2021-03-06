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

// WidgetMargin Size of the margins around the image.
type WidgetMargin string

// List of WidgetMargin
const (
	WIDGETMARGIN_SMALL WidgetMargin = "small"
	WIDGETMARGIN_LARGE WidgetMargin = "large"
)

// Ptr returns reference to WidgetMargin value
func (v WidgetMargin) Ptr() *WidgetMargin {
	return &v
}

type NullableWidgetMargin struct {
	value *WidgetMargin
	isSet bool
}

func (v NullableWidgetMargin) Get() *WidgetMargin {
	return v.value
}

func (v *NullableWidgetMargin) Set(val *WidgetMargin) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetMargin) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetMargin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetMargin(val *WidgetMargin) *NullableWidgetMargin {
	return &NullableWidgetMargin{value: val, isSet: true}
}

func (v NullableWidgetMargin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetMargin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
