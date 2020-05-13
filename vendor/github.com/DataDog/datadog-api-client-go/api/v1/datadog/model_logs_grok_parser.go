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

// LogsGrokParser Create custom grok rules to parse the full message or [a specific attribute of your raw event](https://docs.datadoghq.com/logs/processing/parsing/#advanced-settings). For more information, see the [parsing section](https://docs.datadoghq.com/logs/processing/parsing).
type LogsGrokParser struct {
	Grok LogsGrokParserRules `json:"grok"`
	// List of sample logs to test this grok parser.
	Samples *[]string `json:"samples,omitempty"`
	// Name of the log attribute to parse.
	Source string `json:"source"`
	// Type of processor.
	Type string `json:"type"`
	// Whether or not the processor is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the processor.
	Name *string `json:"name,omitempty"`
}

// NewLogsGrokParser instantiates a new LogsGrokParser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsGrokParser(grok LogsGrokParserRules, source string, type_ string) *LogsGrokParser {
	this := LogsGrokParser{}
	this.Grok = grok
	this.Source = source
	this.Type = type_
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	return &this
}

// NewLogsGrokParserWithDefaults instantiates a new LogsGrokParser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsGrokParserWithDefaults() *LogsGrokParser {
	this := LogsGrokParser{}
	var source string = "message"
	this.Source = source
	var type_ string = "grok-parser"
	this.Type = type_
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	return &this
}

// GetGrok returns the Grok field value
func (o *LogsGrokParser) GetGrok() LogsGrokParserRules {
	if o == nil {
		var ret LogsGrokParserRules
		return ret
	}

	return o.Grok
}

// GetGrokOk returns a tuple with the Grok field value
// and a boolean to check if the value has been set.
func (o *LogsGrokParser) GetGrokOk() (*LogsGrokParserRules, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Grok, true
}

// SetGrok sets field value
func (o *LogsGrokParser) SetGrok(v LogsGrokParserRules) {
	o.Grok = v
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *LogsGrokParser) GetSamples() []string {
	if o == nil || o.Samples == nil {
		var ret []string
		return ret
	}
	return *o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsGrokParser) GetSamplesOk() (*[]string, bool) {
	if o == nil || o.Samples == nil {
		return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *LogsGrokParser) HasSamples() bool {
	if o != nil && o.Samples != nil {
		return true
	}

	return false
}

// SetSamples gets a reference to the given []string and assigns it to the Samples field.
func (o *LogsGrokParser) SetSamples(v []string) {
	o.Samples = &v
}

// GetSource returns the Source field value
func (o *LogsGrokParser) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *LogsGrokParser) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *LogsGrokParser) SetSource(v string) {
	o.Source = v
}

// GetType returns the Type field value
func (o *LogsGrokParser) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsGrokParser) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LogsGrokParser) SetType(v string) {
	o.Type = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LogsGrokParser) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsGrokParser) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LogsGrokParser) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LogsGrokParser) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsGrokParser) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsGrokParser) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsGrokParser) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsGrokParser) SetName(v string) {
	o.Name = &v
}

func (o LogsGrokParser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["grok"] = o.Grok
	}
	if o.Samples != nil {
		toSerialize["samples"] = o.Samples
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

// AsLogsProcessor wraps this instance of LogsGrokParser in LogsProcessor
func (s *LogsGrokParser) AsLogsProcessor() LogsProcessor {
	return LogsProcessor{LogsProcessorInterface: s}
}

type NullableLogsGrokParser struct {
	value *LogsGrokParser
	isSet bool
}

func (v NullableLogsGrokParser) Get() *LogsGrokParser {
	return v.value
}

func (v *NullableLogsGrokParser) Set(val *LogsGrokParser) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsGrokParser) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsGrokParser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsGrokParser(val *LogsGrokParser) *NullableLogsGrokParser {
	return &NullableLogsGrokParser{value: val, isSet: true}
}

func (v NullableLogsGrokParser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsGrokParser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
