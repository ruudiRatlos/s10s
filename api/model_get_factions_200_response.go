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

// checks if the GetFactions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFactions200Response{}

// GetFactions200Response struct for GetFactions200Response
type GetFactions200Response struct {
	Data                 []Faction `json:"data"`
	Meta                 Meta      `json:"meta"`
	AdditionalProperties map[string]interface{}
}

type _GetFactions200Response GetFactions200Response

// NewGetFactions200Response instantiates a new GetFactions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFactions200Response(data []Faction, meta Meta) *GetFactions200Response {
	this := GetFactions200Response{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewGetFactions200ResponseWithDefaults instantiates a new GetFactions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFactions200ResponseWithDefaults() *GetFactions200Response {
	this := GetFactions200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetFactions200Response) GetData() []Faction {
	if o == nil {
		var ret []Faction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetFactions200Response) GetDataOk() ([]Faction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetFactions200Response) SetData(v []Faction) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *GetFactions200Response) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetFactions200Response) GetMetaOk() (*Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetFactions200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetFactions200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFactions200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFactions200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"meta",
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

	varGetFactions200Response := _GetFactions200Response{}

	err = json.Unmarshal(data, &varGetFactions200Response)

	if err != nil {
		return err
	}

	*o = GetFactions200Response(varGetFactions200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFactions200Response struct {
	value *GetFactions200Response
	isSet bool
}

func (v NullableGetFactions200Response) Get() *GetFactions200Response {
	return v.value
}

func (v *NullableGetFactions200Response) Set(val *GetFactions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFactions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFactions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFactions200Response(val *GetFactions200Response) *NullableGetFactions200Response {
	return &NullableGetFactions200Response{value: val, isSet: true}
}

func (v NullableGetFactions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFactions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
