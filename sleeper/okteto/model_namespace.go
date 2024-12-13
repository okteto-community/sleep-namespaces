/*
Okteto Public API

The Okteto Public API is defined using OpenAPI, providing access to resources managed by Okteto. Combined with the Okteto CLI or other tools, it allows you to automate workflows, like deleting unused development volumes within a Namespace or redeploying applications. 

API version: 0.1.0
Contact: support@okteto.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package okteto

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the Namespace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Namespace{}

// Namespace struct for Namespace
type Namespace struct {
	// Date and time when the namespace was created.
	CreationDate time.Time `json:"creationDate"`
	// Date and time when the namespace was last updated.
	LastUpdated time.Time `json:"lastUpdated"`
	// Name of the namespace.
	Name string `json:"name"`
	// Indicates if the namespace is persistent.
	Persistent bool `json:"persistent"`
	// Indicates if the namespace is the default namespace for any user in the instance.
	Personal bool `json:"personal"`
	// Status of the namespace. Possible values are 'Active', 'Inactive', 'Deleted'.
	Status string `json:"status"`
	// Type of the namespace. Possible values are 'Development' and 'Preview'.
	Type string `json:"type"`
	// Unique identifier of the namespace.
	Uuid string `json:"uuid"`
}

type _Namespace Namespace

// NewNamespace instantiates a new Namespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamespace(creationDate time.Time, lastUpdated time.Time, name string, persistent bool, personal bool, status string, type_ string, uuid string) *Namespace {
	this := Namespace{}
	this.CreationDate = creationDate
	this.LastUpdated = lastUpdated
	this.Name = name
	this.Persistent = persistent
	this.Personal = personal
	this.Status = status
	this.Type = type_
	this.Uuid = uuid
	return &this
}

// NewNamespaceWithDefaults instantiates a new Namespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamespaceWithDefaults() *Namespace {
	this := Namespace{}
	return &this
}

// GetCreationDate returns the CreationDate field value
func (o *Namespace) GetCreationDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value
// and a boolean to check if the value has been set.
func (o *Namespace) GetCreationDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationDate, true
}

// SetCreationDate sets field value
func (o *Namespace) SetCreationDate(v time.Time) {
	o.CreationDate = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *Namespace) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *Namespace) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *Namespace) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetName returns the Name field value
func (o *Namespace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Namespace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Namespace) SetName(v string) {
	o.Name = v
}

// GetPersistent returns the Persistent field value
func (o *Namespace) GetPersistent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Persistent
}

// GetPersistentOk returns a tuple with the Persistent field value
// and a boolean to check if the value has been set.
func (o *Namespace) GetPersistentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Persistent, true
}

// SetPersistent sets field value
func (o *Namespace) SetPersistent(v bool) {
	o.Persistent = v
}

// GetPersonal returns the Personal field value
func (o *Namespace) GetPersonal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Personal
}

// GetPersonalOk returns a tuple with the Personal field value
// and a boolean to check if the value has been set.
func (o *Namespace) GetPersonalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Personal, true
}

// SetPersonal sets field value
func (o *Namespace) SetPersonal(v bool) {
	o.Personal = v
}

// GetStatus returns the Status field value
func (o *Namespace) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Namespace) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Namespace) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *Namespace) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Namespace) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Namespace) SetType(v string) {
	o.Type = v
}

// GetUuid returns the Uuid field value
func (o *Namespace) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *Namespace) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *Namespace) SetUuid(v string) {
	o.Uuid = v
}

func (o Namespace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Namespace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["creationDate"] = o.CreationDate
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["name"] = o.Name
	toSerialize["persistent"] = o.Persistent
	toSerialize["personal"] = o.Personal
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	toSerialize["uuid"] = o.Uuid
	return toSerialize, nil
}

func (o *Namespace) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"creationDate",
		"lastUpdated",
		"name",
		"persistent",
		"personal",
		"status",
		"type",
		"uuid",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNamespace := _Namespace{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNamespace)

	if err != nil {
		return err
	}

	*o = Namespace(varNamespace)

	return err
}

type NullableNamespace struct {
	value *Namespace
	isSet bool
}

func (v NullableNamespace) Get() *Namespace {
	return v.value
}

func (v *NullableNamespace) Set(val *Namespace) {
	v.value = val
	v.isSet = true
}

func (v NullableNamespace) IsSet() bool {
	return v.isSet
}

func (v *NullableNamespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamespace(val *Namespace) *NullableNamespace {
	return &NullableNamespace{value: val, isSet: true}
}

func (v NullableNamespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


