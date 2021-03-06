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

// UserWorkflowSubscriptionsSummary struct for UserWorkflowSubscriptionsSummary
type UserWorkflowSubscriptionsSummary struct {
	Name string `json:"name"`
	Subscriptions *UserWorkflowSubscriptions `json:"subscriptions,omitempty"`
}

// NewUserWorkflowSubscriptionsSummary instantiates a new UserWorkflowSubscriptionsSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserWorkflowSubscriptionsSummary(name string) *UserWorkflowSubscriptionsSummary {
	this := UserWorkflowSubscriptionsSummary{}
	this.Name = name
	return &this
}

// NewUserWorkflowSubscriptionsSummaryWithDefaults instantiates a new UserWorkflowSubscriptionsSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWorkflowSubscriptionsSummaryWithDefaults() *UserWorkflowSubscriptionsSummary {
	this := UserWorkflowSubscriptionsSummary{}
	return &this
}

// GetName returns the Name field value
func (o *UserWorkflowSubscriptionsSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserWorkflowSubscriptionsSummary) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserWorkflowSubscriptionsSummary) SetName(v string) {
	o.Name = v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *UserWorkflowSubscriptionsSummary) GetSubscriptions() UserWorkflowSubscriptions {
	if o == nil || o.Subscriptions == nil {
		var ret UserWorkflowSubscriptions
		return ret
	}
	return *o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWorkflowSubscriptionsSummary) GetSubscriptionsOk() (*UserWorkflowSubscriptions, bool) {
	if o == nil || o.Subscriptions == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *UserWorkflowSubscriptionsSummary) HasSubscriptions() bool {
	if o != nil && o.Subscriptions != nil {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given UserWorkflowSubscriptions and assigns it to the Subscriptions field.
func (o *UserWorkflowSubscriptionsSummary) SetSubscriptions(v UserWorkflowSubscriptions) {
	o.Subscriptions = &v
}

func (o UserWorkflowSubscriptionsSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Subscriptions != nil {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	return json.Marshal(toSerialize)
}

type NullableUserWorkflowSubscriptionsSummary struct {
	value *UserWorkflowSubscriptionsSummary
	isSet bool
}

func (v NullableUserWorkflowSubscriptionsSummary) Get() *UserWorkflowSubscriptionsSummary {
	return v.value
}

func (v *NullableUserWorkflowSubscriptionsSummary) Set(val *UserWorkflowSubscriptionsSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableUserWorkflowSubscriptionsSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableUserWorkflowSubscriptionsSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserWorkflowSubscriptionsSummary(val *UserWorkflowSubscriptionsSummary) *NullableUserWorkflowSubscriptionsSummary {
	return &NullableUserWorkflowSubscriptionsSummary{value: val, isSet: true}
}

func (v NullableUserWorkflowSubscriptionsSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserWorkflowSubscriptionsSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


