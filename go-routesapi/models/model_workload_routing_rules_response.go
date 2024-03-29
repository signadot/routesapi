/*
Signadot Routes API

The Routes API provides access to in-cluster routing configuration set up by  the Signadot Operator. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// checks if the WorkloadRoutingRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkloadRoutingRulesResponse{}

// WorkloadRoutingRulesResponse A WorkloadRoutingRulesResponse gives the set of WorkloadRoutingRules which match a given WorkloadRoutingRulesRequest.
type WorkloadRoutingRulesResponse struct {
	RoutingRules []WorkloadRoutingRule `json:"routingRules,omitempty"`
}

// NewWorkloadRoutingRulesResponse instantiates a new WorkloadRoutingRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadRoutingRulesResponse() *WorkloadRoutingRulesResponse {
	this := WorkloadRoutingRulesResponse{}
	return &this
}

// NewWorkloadRoutingRulesResponseWithDefaults instantiates a new WorkloadRoutingRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadRoutingRulesResponseWithDefaults() *WorkloadRoutingRulesResponse {
	this := WorkloadRoutingRulesResponse{}
	return &this
}

// GetRoutingRules returns the RoutingRules field value if set, zero value otherwise.
func (o *WorkloadRoutingRulesResponse) GetRoutingRules() []WorkloadRoutingRule {
	if o == nil || IsNil(o.RoutingRules) {
		var ret []WorkloadRoutingRule
		return ret
	}
	return o.RoutingRules
}

// GetRoutingRulesOk returns a tuple with the RoutingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadRoutingRulesResponse) GetRoutingRulesOk() ([]WorkloadRoutingRule, bool) {
	if o == nil || IsNil(o.RoutingRules) {
		return nil, false
	}
	return o.RoutingRules, true
}

// HasRoutingRules returns a boolean if a field has been set.
func (o *WorkloadRoutingRulesResponse) HasRoutingRules() bool {
	if o != nil && !IsNil(o.RoutingRules) {
		return true
	}

	return false
}

// SetRoutingRules gets a reference to the given []WorkloadRoutingRule and assigns it to the RoutingRules field.
func (o *WorkloadRoutingRulesResponse) SetRoutingRules(v []WorkloadRoutingRule) {
	o.RoutingRules = v
}

func (o WorkloadRoutingRulesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkloadRoutingRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoutingRules) {
		toSerialize["routingRules"] = o.RoutingRules
	}
	return toSerialize, nil
}

type NullableWorkloadRoutingRulesResponse struct {
	value *WorkloadRoutingRulesResponse
	isSet bool
}

func (v NullableWorkloadRoutingRulesResponse) Get() *WorkloadRoutingRulesResponse {
	return v.value
}

func (v *NullableWorkloadRoutingRulesResponse) Set(val *WorkloadRoutingRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadRoutingRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadRoutingRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadRoutingRulesResponse(val *WorkloadRoutingRulesResponse) *NullableWorkloadRoutingRulesResponse {
	return &NullableWorkloadRoutingRulesResponse{value: val, isSet: true}
}

func (v NullableWorkloadRoutingRulesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadRoutingRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


