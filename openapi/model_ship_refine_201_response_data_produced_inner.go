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

// checks if the ShipRefine201ResponseDataProducedInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipRefine201ResponseDataProducedInner{}

// ShipRefine201ResponseDataProducedInner struct for ShipRefine201ResponseDataProducedInner
type ShipRefine201ResponseDataProducedInner struct {
	// Symbol of the good.
	TradeSymbol string `json:"tradeSymbol"`
	// Amount of units of the good.
	Units                int32 `json:"units"`
	AdditionalProperties map[string]interface{}
}

type _ShipRefine201ResponseDataProducedInner ShipRefine201ResponseDataProducedInner

// NewShipRefine201ResponseDataProducedInner instantiates a new ShipRefine201ResponseDataProducedInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipRefine201ResponseDataProducedInner(tradeSymbol string, units int32) *ShipRefine201ResponseDataProducedInner {
	this := ShipRefine201ResponseDataProducedInner{}
	this.TradeSymbol = tradeSymbol
	this.Units = units
	return &this
}

// NewShipRefine201ResponseDataProducedInnerWithDefaults instantiates a new ShipRefine201ResponseDataProducedInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipRefine201ResponseDataProducedInnerWithDefaults() *ShipRefine201ResponseDataProducedInner {
	this := ShipRefine201ResponseDataProducedInner{}
	return &this
}

// GetTradeSymbol returns the TradeSymbol field value
func (o *ShipRefine201ResponseDataProducedInner) GetTradeSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TradeSymbol
}

// GetTradeSymbolOk returns a tuple with the TradeSymbol field value
// and a boolean to check if the value has been set.
func (o *ShipRefine201ResponseDataProducedInner) GetTradeSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeSymbol, true
}

// SetTradeSymbol sets field value
func (o *ShipRefine201ResponseDataProducedInner) SetTradeSymbol(v string) {
	o.TradeSymbol = v
}

// GetUnits returns the Units field value
func (o *ShipRefine201ResponseDataProducedInner) GetUnits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *ShipRefine201ResponseDataProducedInner) GetUnitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *ShipRefine201ResponseDataProducedInner) SetUnits(v int32) {
	o.Units = v
}

func (o ShipRefine201ResponseDataProducedInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipRefine201ResponseDataProducedInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tradeSymbol"] = o.TradeSymbol
	toSerialize["units"] = o.Units

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ShipRefine201ResponseDataProducedInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tradeSymbol",
		"units",
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

	varShipRefine201ResponseDataProducedInner := _ShipRefine201ResponseDataProducedInner{}

	err = json.Unmarshal(data, &varShipRefine201ResponseDataProducedInner)

	if err != nil {
		return err
	}

	*o = ShipRefine201ResponseDataProducedInner(varShipRefine201ResponseDataProducedInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tradeSymbol")
		delete(additionalProperties, "units")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableShipRefine201ResponseDataProducedInner struct {
	value *ShipRefine201ResponseDataProducedInner
	isSet bool
}

func (v NullableShipRefine201ResponseDataProducedInner) Get() *ShipRefine201ResponseDataProducedInner {
	return v.value
}

func (v *NullableShipRefine201ResponseDataProducedInner) Set(val *ShipRefine201ResponseDataProducedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableShipRefine201ResponseDataProducedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableShipRefine201ResponseDataProducedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipRefine201ResponseDataProducedInner(val *ShipRefine201ResponseDataProducedInner) *NullableShipRefine201ResponseDataProducedInner {
	return &NullableShipRefine201ResponseDataProducedInner{value: val, isSet: true}
}

func (v NullableShipRefine201ResponseDataProducedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipRefine201ResponseDataProducedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
