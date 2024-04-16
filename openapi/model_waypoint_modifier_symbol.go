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

// WaypointModifierSymbol The unique identifier of the modifier.
type WaypointModifierSymbol string

// List of WaypointModifierSymbol
const (
	WAYPOINTMODIFIERSYMBOL_STRIPPED       WaypointModifierSymbol = "STRIPPED"
	WAYPOINTMODIFIERSYMBOL_UNSTABLE       WaypointModifierSymbol = "UNSTABLE"
	WAYPOINTMODIFIERSYMBOL_RADIATION_LEAK WaypointModifierSymbol = "RADIATION_LEAK"
	WAYPOINTMODIFIERSYMBOL_CRITICAL_LIMIT WaypointModifierSymbol = "CRITICAL_LIMIT"
	WAYPOINTMODIFIERSYMBOL_CIVIL_UNREST   WaypointModifierSymbol = "CIVIL_UNREST"
)

// All allowed values of WaypointModifierSymbol enum
var AllowedWaypointModifierSymbolEnumValues = []WaypointModifierSymbol{
	"STRIPPED",
	"UNSTABLE",
	"RADIATION_LEAK",
	"CRITICAL_LIMIT",
	"CIVIL_UNREST",
}

func (v *WaypointModifierSymbol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WaypointModifierSymbol(value)
	for _, existing := range AllowedWaypointModifierSymbolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WaypointModifierSymbol", value)
}

// NewWaypointModifierSymbolFromValue returns a pointer to a valid WaypointModifierSymbol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWaypointModifierSymbolFromValue(v string) (*WaypointModifierSymbol, error) {
	ev := WaypointModifierSymbol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WaypointModifierSymbol: valid values are %v", v, AllowedWaypointModifierSymbolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WaypointModifierSymbol) IsValid() bool {
	for _, existing := range AllowedWaypointModifierSymbolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WaypointModifierSymbol value
func (v WaypointModifierSymbol) Ptr() *WaypointModifierSymbol {
	return &v
}

type NullableWaypointModifierSymbol struct {
	value *WaypointModifierSymbol
	isSet bool
}

func (v NullableWaypointModifierSymbol) Get() *WaypointModifierSymbol {
	return v.value
}

func (v *NullableWaypointModifierSymbol) Set(val *WaypointModifierSymbol) {
	v.value = val
	v.isSet = true
}

func (v NullableWaypointModifierSymbol) IsSet() bool {
	return v.isSet
}

func (v *NullableWaypointModifierSymbol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWaypointModifierSymbol(val *WaypointModifierSymbol) *NullableWaypointModifierSymbol {
	return &NullableWaypointModifierSymbol{value: val, isSet: true}
}

func (v NullableWaypointModifierSymbol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWaypointModifierSymbol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
