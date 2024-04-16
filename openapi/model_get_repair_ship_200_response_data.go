/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the GetRepairShip200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRepairShip200ResponseData{}

// GetRepairShip200ResponseData struct for GetRepairShip200ResponseData
type GetRepairShip200ResponseData struct {
	Transaction          RepairTransaction `json:"transaction"`
	AdditionalProperties map[string]interface{}
}

type _GetRepairShip200ResponseData GetRepairShip200ResponseData

// NewGetRepairShip200ResponseData instantiates a new GetRepairShip200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRepairShip200ResponseData(transaction RepairTransaction) *GetRepairShip200ResponseData {
	this := GetRepairShip200ResponseData{}
	this.Transaction = transaction
	return &this
}

// NewGetRepairShip200ResponseDataWithDefaults instantiates a new GetRepairShip200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRepairShip200ResponseDataWithDefaults() *GetRepairShip200ResponseData {
	this := GetRepairShip200ResponseData{}
	return &this
}

// GetTransaction returns the Transaction field value
func (o *GetRepairShip200ResponseData) GetTransaction() RepairTransaction {
	if o == nil {
		var ret RepairTransaction
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *GetRepairShip200ResponseData) GetTransactionOk() (*RepairTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *GetRepairShip200ResponseData) SetTransaction(v RepairTransaction) {
	o.Transaction = v
}

func (o GetRepairShip200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRepairShip200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction"] = o.Transaction

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetRepairShip200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetRepairShip200ResponseData := _GetRepairShip200ResponseData{}

	err = json.Unmarshal(data, &varGetRepairShip200ResponseData)

	if err != nil {
		return err
	}

	*o = GetRepairShip200ResponseData(varGetRepairShip200ResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "transaction")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRepairShip200ResponseData struct {
	value *GetRepairShip200ResponseData
	isSet bool
}

func (v NullableGetRepairShip200ResponseData) Get() *GetRepairShip200ResponseData {
	return v.value
}

func (v *NullableGetRepairShip200ResponseData) Set(val *GetRepairShip200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRepairShip200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRepairShip200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRepairShip200ResponseData(val *GetRepairShip200ResponseData) *NullableGetRepairShip200ResponseData {
	return &NullableGetRepairShip200ResponseData{value: val, isSet: true}
}

func (v NullableGetRepairShip200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRepairShip200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
