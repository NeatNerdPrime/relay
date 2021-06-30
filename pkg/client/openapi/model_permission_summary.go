/*
 * Relay API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v20200615
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PermissionSummary struct for PermissionSummary
type PermissionSummary struct {
	// The name of the permission
	Name string `json:"name"`
	// A human-readable explanation of the access provided by this permission
	Description *string `json:"description,omitempty"`
}

// NewPermissionSummary instantiates a new PermissionSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionSummary(name string) *PermissionSummary {
	this := PermissionSummary{}
	this.Name = name
	return &this
}

// NewPermissionSummaryWithDefaults instantiates a new PermissionSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionSummaryWithDefaults() *PermissionSummary {
	this := PermissionSummary{}
	return &this
}

// GetName returns the Name field value
func (o *PermissionSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PermissionSummary) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PermissionSummary) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PermissionSummary) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionSummary) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PermissionSummary) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PermissionSummary) SetDescription(v string) {
	o.Description = &v
}

func (o PermissionSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionSummary struct {
	value *PermissionSummary
	isSet bool
}

func (v NullablePermissionSummary) Get() *PermissionSummary {
	return v.value
}

func (v *NullablePermissionSummary) Set(val *PermissionSummary) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionSummary) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionSummary(val *PermissionSummary) *NullablePermissionSummary {
	return &NullablePermissionSummary{value: val, isSet: true}
}

func (v NullablePermissionSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


