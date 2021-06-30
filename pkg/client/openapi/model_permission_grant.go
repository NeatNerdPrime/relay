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

// PermissionGrant A permission resolved in a given context, such as a user in an entity
type PermissionGrant struct {
	Permission *PermissionSummary `json:"permission,omitempty"`
	// Whether this permission has been granted in the context or a sub-context
	Granted *string `json:"granted,omitempty"`
}

// NewPermissionGrant instantiates a new PermissionGrant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionGrant() *PermissionGrant {
	this := PermissionGrant{}
	return &this
}

// NewPermissionGrantWithDefaults instantiates a new PermissionGrant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionGrantWithDefaults() *PermissionGrant {
	this := PermissionGrant{}
	return &this
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *PermissionGrant) GetPermission() PermissionSummary {
	if o == nil || o.Permission == nil {
		var ret PermissionSummary
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGrant) GetPermissionOk() (*PermissionSummary, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *PermissionGrant) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given PermissionSummary and assigns it to the Permission field.
func (o *PermissionGrant) SetPermission(v PermissionSummary) {
	o.Permission = &v
}

// GetGranted returns the Granted field value if set, zero value otherwise.
func (o *PermissionGrant) GetGranted() string {
	if o == nil || o.Granted == nil {
		var ret string
		return ret
	}
	return *o.Granted
}

// GetGrantedOk returns a tuple with the Granted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGrant) GetGrantedOk() (*string, bool) {
	if o == nil || o.Granted == nil {
		return nil, false
	}
	return o.Granted, true
}

// HasGranted returns a boolean if a field has been set.
func (o *PermissionGrant) HasGranted() bool {
	if o != nil && o.Granted != nil {
		return true
	}

	return false
}

// SetGranted gets a reference to the given string and assigns it to the Granted field.
func (o *PermissionGrant) SetGranted(v string) {
	o.Granted = &v
}

func (o PermissionGrant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Permission != nil {
		toSerialize["permission"] = o.Permission
	}
	if o.Granted != nil {
		toSerialize["granted"] = o.Granted
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionGrant struct {
	value *PermissionGrant
	isSet bool
}

func (v NullablePermissionGrant) Get() *PermissionGrant {
	return v.value
}

func (v *NullablePermissionGrant) Set(val *PermissionGrant) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionGrant) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionGrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionGrant(val *PermissionGrant) *NullablePermissionGrant {
	return &NullablePermissionGrant{value: val, isSet: true}
}

func (v NullablePermissionGrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionGrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


