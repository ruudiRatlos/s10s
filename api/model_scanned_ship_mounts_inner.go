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

// checks if the ScannedShipMountsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScannedShipMountsInner{}

// ScannedShipMountsInner A mount on the ship.
type ScannedShipMountsInner struct {
	// The symbol of the mount.
	Symbol string `json:"symbol"`
	AdditionalProperties map[string]interface{}
}

type _ScannedShipMountsInner ScannedShipMountsInner

// NewScannedShipMountsInner instantiates a new ScannedShipMountsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScannedShipMountsInner(symbol string) *ScannedShipMountsInner {
	this := ScannedShipMountsInner{}
	this.Symbol = symbol
	return &this
}

// NewScannedShipMountsInnerWithDefaults instantiates a new ScannedShipMountsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScannedShipMountsInnerWithDefaults() *ScannedShipMountsInner {
	this := ScannedShipMountsInner{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ScannedShipMountsInner) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ScannedShipMountsInner) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ScannedShipMountsInner) SetSymbol(v string) {
	o.Symbol = v
}

func (o ScannedShipMountsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScannedShipMountsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScannedShipMountsInner) UnmarshalJSON(data []byte) (err error) {
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

	varScannedShipMountsInner := _ScannedShipMountsInner{}

	err = json.Unmarshal(data, &varScannedShipMountsInner)

	if err != nil {
		return err
	}

	*o = ScannedShipMountsInner(varScannedShipMountsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScannedShipMountsInner struct {
	value *ScannedShipMountsInner
	isSet bool
}

func (v NullableScannedShipMountsInner) Get() *ScannedShipMountsInner {
	return v.value
}

func (v *NullableScannedShipMountsInner) Set(val *ScannedShipMountsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableScannedShipMountsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableScannedShipMountsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScannedShipMountsInner(val *ScannedShipMountsInner) *NullableScannedShipMountsInner {
	return &NullableScannedShipMountsInner{value: val, isSet: true}
}

func (v NullableScannedShipMountsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScannedShipMountsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


