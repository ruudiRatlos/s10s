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

// checks if the GetSystemWaypoints200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSystemWaypoints200Response{}

// GetSystemWaypoints200Response
type GetSystemWaypoints200Response struct {
	Data                 []Waypoint `json:"data"`
	Meta                 Meta       `json:"meta"`
	AdditionalProperties map[string]interface{}
}

type _GetSystemWaypoints200Response GetSystemWaypoints200Response

// NewGetSystemWaypoints200Response instantiates a new GetSystemWaypoints200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSystemWaypoints200Response(data []Waypoint, meta Meta) *GetSystemWaypoints200Response {
	this := GetSystemWaypoints200Response{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewGetSystemWaypoints200ResponseWithDefaults instantiates a new GetSystemWaypoints200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSystemWaypoints200ResponseWithDefaults() *GetSystemWaypoints200Response {
	this := GetSystemWaypoints200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetSystemWaypoints200Response) GetData() []Waypoint {
	if o == nil {
		var ret []Waypoint
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetSystemWaypoints200Response) GetDataOk() ([]Waypoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetSystemWaypoints200Response) SetData(v []Waypoint) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *GetSystemWaypoints200Response) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetSystemWaypoints200Response) GetMetaOk() (*Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetSystemWaypoints200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetSystemWaypoints200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSystemWaypoints200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSystemWaypoints200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetSystemWaypoints200Response := _GetSystemWaypoints200Response{}

	err = json.Unmarshal(data, &varGetSystemWaypoints200Response)

	if err != nil {
		return err
	}

	*o = GetSystemWaypoints200Response(varGetSystemWaypoints200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSystemWaypoints200Response struct {
	value *GetSystemWaypoints200Response
	isSet bool
}

func (v NullableGetSystemWaypoints200Response) Get() *GetSystemWaypoints200Response {
	return v.value
}

func (v *NullableGetSystemWaypoints200Response) Set(val *GetSystemWaypoints200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSystemWaypoints200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSystemWaypoints200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSystemWaypoints200Response(val *GetSystemWaypoints200Response) *NullableGetSystemWaypoints200Response {
	return &NullableGetSystemWaypoints200Response{value: val, isSet: true}
}

func (v NullableGetSystemWaypoints200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSystemWaypoints200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
