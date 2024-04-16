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

// checks if the CreateShipShipScan201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateShipShipScan201Response{}

// CreateShipShipScan201Response struct for CreateShipShipScan201Response
type CreateShipShipScan201Response struct {
	Data                 CreateShipShipScan201ResponseData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _CreateShipShipScan201Response CreateShipShipScan201Response

// NewCreateShipShipScan201Response instantiates a new CreateShipShipScan201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShipShipScan201Response(data CreateShipShipScan201ResponseData) *CreateShipShipScan201Response {
	this := CreateShipShipScan201Response{}
	this.Data = data
	return &this
}

// NewCreateShipShipScan201ResponseWithDefaults instantiates a new CreateShipShipScan201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShipShipScan201ResponseWithDefaults() *CreateShipShipScan201Response {
	this := CreateShipShipScan201Response{}
	return &this
}

// GetData returns the Data field value
func (o *CreateShipShipScan201Response) GetData() CreateShipShipScan201ResponseData {
	if o == nil {
		var ret CreateShipShipScan201ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateShipShipScan201Response) GetDataOk() (*CreateShipShipScan201ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateShipShipScan201Response) SetData(v CreateShipShipScan201ResponseData) {
	o.Data = v
}

func (o CreateShipShipScan201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateShipShipScan201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateShipShipScan201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varCreateShipShipScan201Response := _CreateShipShipScan201Response{}

	err = json.Unmarshal(data, &varCreateShipShipScan201Response)

	if err != nil {
		return err
	}

	*o = CreateShipShipScan201Response(varCreateShipShipScan201Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateShipShipScan201Response struct {
	value *CreateShipShipScan201Response
	isSet bool
}

func (v NullableCreateShipShipScan201Response) Get() *CreateShipShipScan201Response {
	return v.value
}

func (v *NullableCreateShipShipScan201Response) Set(val *CreateShipShipScan201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShipShipScan201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShipShipScan201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShipShipScan201Response(val *CreateShipShipScan201Response) *NullableCreateShipShipScan201Response {
	return &NullableCreateShipShipScan201Response{value: val, isSet: true}
}

func (v NullableCreateShipShipScan201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShipShipScan201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
