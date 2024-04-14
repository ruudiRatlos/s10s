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

// checks if the SystemFaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemFaction{}

// SystemFaction struct for SystemFaction
type SystemFaction struct {
	Symbol FactionSymbol `json:"symbol"`
	AdditionalProperties map[string]interface{}
}

type _SystemFaction SystemFaction

// NewSystemFaction instantiates a new SystemFaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemFaction(symbol FactionSymbol) *SystemFaction {
	this := SystemFaction{}
	this.Symbol = symbol
	return &this
}

// NewSystemFactionWithDefaults instantiates a new SystemFaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemFactionWithDefaults() *SystemFaction {
	this := SystemFaction{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *SystemFaction) GetSymbol() FactionSymbol {
	if o == nil {
		var ret FactionSymbol
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *SystemFaction) GetSymbolOk() (*FactionSymbol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *SystemFaction) SetSymbol(v FactionSymbol) {
	o.Symbol = v
}

func (o SystemFaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemFaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SystemFaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
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

	varSystemFaction := _SystemFaction{}

	err = json.Unmarshal(data, &varSystemFaction)

	if err != nil {
		return err
	}

	*o = SystemFaction(varSystemFaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSystemFaction struct {
	value *SystemFaction
	isSet bool
}

func (v NullableSystemFaction) Get() *SystemFaction {
	return v.value
}

func (v *NullableSystemFaction) Set(val *SystemFaction) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemFaction) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemFaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemFaction(val *SystemFaction) *NullableSystemFaction {
	return &NullableSystemFaction{value: val, isSet: true}
}

func (v NullableSystemFaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemFaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


