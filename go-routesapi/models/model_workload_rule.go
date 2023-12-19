/*
Signadot Routes API

The Signadot Routes API provides access to routing rules pertinent to Signadot Sandboxes on a cluster with the Signadot Operator (>= v0.14.2) installed. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
)

// checks if the WorkloadRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkloadRule{}

// WorkloadRule A WorkloadRule r means: if a request    1. has r.RoutingKey; and    2. is originally destined to r.SandboxedWorkload.Baseline; and   3. is sent on a port indicated in one of r.PortRules pr    then send it to the host and port indicated in any destination host and port indicated in pr.destinations. Moreover, these destinations are all addresses of r.SandboxedWorkload, any one of them can be used.
type WorkloadRule struct {
	// The routing key
	RoutingKey string `json:"routingKey"`
	SandboxedWorkload SandboxedWorkload `json:"sandboxedWorkload"`
	// Workload port rules
	PortRules []WorkloadPortRule `json:"portRules,omitempty"`
}

type _WorkloadRule WorkloadRule

// NewWorkloadRule instantiates a new WorkloadRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadRule(routingKey string, sandboxedWorkload SandboxedWorkload) *WorkloadRule {
	this := WorkloadRule{}
	this.RoutingKey = routingKey
	this.SandboxedWorkload = sandboxedWorkload
	return &this
}

// NewWorkloadRuleWithDefaults instantiates a new WorkloadRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadRuleWithDefaults() *WorkloadRule {
	this := WorkloadRule{}
	return &this
}

// GetRoutingKey returns the RoutingKey field value
func (o *WorkloadRule) GetRoutingKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingKey
}

// GetRoutingKeyOk returns a tuple with the RoutingKey field value
// and a boolean to check if the value has been set.
func (o *WorkloadRule) GetRoutingKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoutingKey, true
}

// SetRoutingKey sets field value
func (o *WorkloadRule) SetRoutingKey(v string) {
	o.RoutingKey = v
}

// GetSandboxedWorkload returns the SandboxedWorkload field value
func (o *WorkloadRule) GetSandboxedWorkload() SandboxedWorkload {
	if o == nil {
		var ret SandboxedWorkload
		return ret
	}

	return o.SandboxedWorkload
}

// GetSandboxedWorkloadOk returns a tuple with the SandboxedWorkload field value
// and a boolean to check if the value has been set.
func (o *WorkloadRule) GetSandboxedWorkloadOk() (*SandboxedWorkload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SandboxedWorkload, true
}

// SetSandboxedWorkload sets field value
func (o *WorkloadRule) SetSandboxedWorkload(v SandboxedWorkload) {
	o.SandboxedWorkload = v
}

// GetPortRules returns the PortRules field value if set, zero value otherwise.
func (o *WorkloadRule) GetPortRules() []WorkloadPortRule {
	if o == nil || IsNil(o.PortRules) {
		var ret []WorkloadPortRule
		return ret
	}
	return o.PortRules
}

// GetPortRulesOk returns a tuple with the PortRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadRule) GetPortRulesOk() ([]WorkloadPortRule, bool) {
	if o == nil || IsNil(o.PortRules) {
		return nil, false
	}
	return o.PortRules, true
}

// HasPortRules returns a boolean if a field has been set.
func (o *WorkloadRule) HasPortRules() bool {
	if o != nil && !IsNil(o.PortRules) {
		return true
	}

	return false
}

// SetPortRules gets a reference to the given []WorkloadPortRule and assigns it to the PortRules field.
func (o *WorkloadRule) SetPortRules(v []WorkloadPortRule) {
	o.PortRules = v
}

func (o WorkloadRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkloadRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["routingKey"] = o.RoutingKey
	toSerialize["sandboxedWorkload"] = o.SandboxedWorkload
	if !IsNil(o.PortRules) {
		toSerialize["portRules"] = o.PortRules
	}
	return toSerialize, nil
}

func (o *WorkloadRule) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"routingKey",
		"sandboxedWorkload",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWorkloadRule := _WorkloadRule{}

	err = json.Unmarshal(bytes, &varWorkloadRule)

	if err != nil {
		return err
	}

	*o = WorkloadRule(varWorkloadRule)

	return err
}

type NullableWorkloadRule struct {
	value *WorkloadRule
	isSet bool
}

func (v NullableWorkloadRule) Get() *WorkloadRule {
	return v.value
}

func (v *NullableWorkloadRule) Set(val *WorkloadRule) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadRule) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadRule(val *WorkloadRule) *NullableWorkloadRule {
	return &NullableWorkloadRule{value: val, isSet: true}
}

func (v NullableWorkloadRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


