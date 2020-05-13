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

// SyntheticsCheckType Type of assertion to apply in an API test.
type SyntheticsCheckType string

// List of SyntheticsCheckType
const (
	SYNTHETICSCHECKTYPE_EQUALS          SyntheticsCheckType = "equals"
	SYNTHETICSCHECKTYPE_NOT_EQUALS      SyntheticsCheckType = "notEquals"
	SYNTHETICSCHECKTYPE_CONTAINS        SyntheticsCheckType = "contains"
	SYNTHETICSCHECKTYPE_NOT_CONTAINS    SyntheticsCheckType = "notContains"
	SYNTHETICSCHECKTYPE_STARTS_WITH     SyntheticsCheckType = "startsWith"
	SYNTHETICSCHECKTYPE_NOT_STARTS_WITH SyntheticsCheckType = "notStartsWith"
	SYNTHETICSCHECKTYPE_GREATER         SyntheticsCheckType = "greater"
	SYNTHETICSCHECKTYPE_LOWER           SyntheticsCheckType = "lower"
	SYNTHETICSCHECKTYPE_GREATER_EQUALS  SyntheticsCheckType = "greaterEquals"
	SYNTHETICSCHECKTYPE_LOWER_EQUALS    SyntheticsCheckType = "lowerEquals"
	SYNTHETICSCHECKTYPE_MATCH_REGEX     SyntheticsCheckType = "matchRegex"
)

// Ptr returns reference to SyntheticsCheckType value
func (v SyntheticsCheckType) Ptr() *SyntheticsCheckType {
	return &v
}

type NullableSyntheticsCheckType struct {
	value *SyntheticsCheckType
	isSet bool
}

func (v NullableSyntheticsCheckType) Get() *SyntheticsCheckType {
	return v.value
}

func (v *NullableSyntheticsCheckType) Set(val *SyntheticsCheckType) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsCheckType) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsCheckType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsCheckType(val *SyntheticsCheckType) *NullableSyntheticsCheckType {
	return &NullableSyntheticsCheckType{value: val, isSet: true}
}

func (v NullableSyntheticsCheckType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsCheckType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
