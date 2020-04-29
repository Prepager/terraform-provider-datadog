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

// OrganizationSettingsSaml Set the boolean property enabled to enable or disable single sign on with SAML. See the SAML documentation for more information about all SAML settings.
type OrganizationSettingsSaml struct {
	// TODO.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewOrganizationSettingsSaml instantiates a new OrganizationSettingsSaml object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationSettingsSaml() *OrganizationSettingsSaml {
	this := OrganizationSettingsSaml{}
	return &this
}

// NewOrganizationSettingsSamlWithDefaults instantiates a new OrganizationSettingsSaml object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationSettingsSamlWithDefaults() *OrganizationSettingsSaml {
	this := OrganizationSettingsSaml{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrganizationSettingsSaml) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSettingsSaml) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrganizationSettingsSaml) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrganizationSettingsSaml) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o OrganizationSettingsSaml) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationSettingsSaml struct {
	value *OrganizationSettingsSaml
	isSet bool
}

func (v NullableOrganizationSettingsSaml) Get() *OrganizationSettingsSaml {
	return v.value
}

func (v *NullableOrganizationSettingsSaml) Set(val *OrganizationSettingsSaml) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationSettingsSaml) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationSettingsSaml) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationSettingsSaml(val *OrganizationSettingsSaml) *NullableOrganizationSettingsSaml {
	return &NullableOrganizationSettingsSaml{value: val, isSet: true}
}

func (v NullableOrganizationSettingsSaml) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationSettingsSaml) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}