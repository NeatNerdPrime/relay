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

// ErrorFormatted struct for ErrorFormatted
type ErrorFormatted struct {
	Friendly *string `json:"friendly,omitempty"`
	Technical *string `json:"technical,omitempty"`
}

// NewErrorFormatted instantiates a new ErrorFormatted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorFormatted() *ErrorFormatted {
	this := ErrorFormatted{}
	return &this
}

// NewErrorFormattedWithDefaults instantiates a new ErrorFormatted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorFormattedWithDefaults() *ErrorFormatted {
	this := ErrorFormatted{}
	return &this
}

// GetFriendly returns the Friendly field value if set, zero value otherwise.
func (o *ErrorFormatted) GetFriendly() string {
	if o == nil || o.Friendly == nil {
		var ret string
		return ret
	}
	return *o.Friendly
}

// GetFriendlyOk returns a tuple with the Friendly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorFormatted) GetFriendlyOk() (*string, bool) {
	if o == nil || o.Friendly == nil {
		return nil, false
	}
	return o.Friendly, true
}

// HasFriendly returns a boolean if a field has been set.
func (o *ErrorFormatted) HasFriendly() bool {
	if o != nil && o.Friendly != nil {
		return true
	}

	return false
}

// SetFriendly gets a reference to the given string and assigns it to the Friendly field.
func (o *ErrorFormatted) SetFriendly(v string) {
	o.Friendly = &v
}

// GetTechnical returns the Technical field value if set, zero value otherwise.
func (o *ErrorFormatted) GetTechnical() string {
	if o == nil || o.Technical == nil {
		var ret string
		return ret
	}
	return *o.Technical
}

// GetTechnicalOk returns a tuple with the Technical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorFormatted) GetTechnicalOk() (*string, bool) {
	if o == nil || o.Technical == nil {
		return nil, false
	}
	return o.Technical, true
}

// HasTechnical returns a boolean if a field has been set.
func (o *ErrorFormatted) HasTechnical() bool {
	if o != nil && o.Technical != nil {
		return true
	}

	return false
}

// SetTechnical gets a reference to the given string and assigns it to the Technical field.
func (o *ErrorFormatted) SetTechnical(v string) {
	o.Technical = &v
}

func (o ErrorFormatted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Friendly != nil {
		toSerialize["friendly"] = o.Friendly
	}
	if o.Technical != nil {
		toSerialize["technical"] = o.Technical
	}
	return json.Marshal(toSerialize)
}

type NullableErrorFormatted struct {
	value *ErrorFormatted
	isSet bool
}

func (v NullableErrorFormatted) Get() *ErrorFormatted {
	return v.value
}

func (v *NullableErrorFormatted) Set(val *ErrorFormatted) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorFormatted) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorFormatted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorFormatted(val *ErrorFormatted) *NullableErrorFormatted {
	return &NullableErrorFormatted{value: val, isSet: true}
}

func (v NullableErrorFormatted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorFormatted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

