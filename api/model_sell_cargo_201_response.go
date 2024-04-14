/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// checks if the SellCargo201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SellCargo201Response{}

// SellCargo201Response 
type SellCargo201Response struct {
	Data SellCargo201ResponseData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _SellCargo201Response SellCargo201Response

// NewSellCargo201Response instantiates a new SellCargo201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSellCargo201Response(data SellCargo201ResponseData) *SellCargo201Response {
	this := SellCargo201Response{}
	this.Data = data
	return &this
}

// NewSellCargo201ResponseWithDefaults instantiates a new SellCargo201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSellCargo201ResponseWithDefaults() *SellCargo201Response {
	this := SellCargo201Response{}
	return &this
}

// GetData returns the Data field value
func (o *SellCargo201Response) GetData() SellCargo201ResponseData {
	if o == nil {
		var ret SellCargo201ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SellCargo201Response) GetDataOk() (*SellCargo201ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SellCargo201Response) SetData(v SellCargo201ResponseData) {
	o.Data = v
}

func (o SellCargo201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SellCargo201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SellCargo201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varSellCargo201Response := _SellCargo201Response{}

	err = json.Unmarshal(data, &varSellCargo201Response)

	if err != nil {
		return err
	}

	*o = SellCargo201Response(varSellCargo201Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSellCargo201Response struct {
	value *SellCargo201Response
	isSet bool
}

func (v NullableSellCargo201Response) Get() *SellCargo201Response {
	return v.value
}

func (v *NullableSellCargo201Response) Set(val *SellCargo201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSellCargo201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSellCargo201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSellCargo201Response(val *SellCargo201Response) *NullableSellCargo201Response {
	return &NullableSellCargo201Response{value: val, isSet: true}
}

func (v NullableSellCargo201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSellCargo201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


