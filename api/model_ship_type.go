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

// ShipType Type of ship
type ShipType string

// List of ShipType
const (
	SHIPTYPE_PROBE              ShipType = "SHIP_PROBE"
	SHIPTYPE_MINING_DRONE       ShipType = "SHIP_MINING_DRONE"
	SHIPTYPE_SIPHON_DRONE       ShipType = "SHIP_SIPHON_DRONE"
	SHIPTYPE_INTERCEPTOR        ShipType = "SHIP_INTERCEPTOR"
	SHIPTYPE_LIGHT_HAULER       ShipType = "SHIP_LIGHT_HAULER"
	SHIPTYPE_COMMAND_FRIGATE    ShipType = "SHIP_COMMAND_FRIGATE"
	SHIPTYPE_EXPLORER           ShipType = "SHIP_EXPLORER"
	SHIPTYPE_HEAVY_FREIGHTER    ShipType = "SHIP_HEAVY_FREIGHTER"
	SHIPTYPE_LIGHT_SHUTTLE      ShipType = "SHIP_LIGHT_SHUTTLE"
	SHIPTYPE_ORE_HOUND          ShipType = "SHIP_ORE_HOUND"
	SHIPTYPE_REFINING_FREIGHTER ShipType = "SHIP_REFINING_FREIGHTER"
	SHIPTYPE_SURVEYOR           ShipType = "SHIP_SURVEYOR"
)

// All allowed values of ShipType enum
var AllowedShipTypeEnumValues = []ShipType{
	"SHIP_PROBE",
	"SHIP_MINING_DRONE",
	"SHIP_SIPHON_DRONE",
	"SHIP_INTERCEPTOR",
	"SHIP_LIGHT_HAULER",
	"SHIP_COMMAND_FRIGATE",
	"SHIP_EXPLORER",
	"SHIP_HEAVY_FREIGHTER",
	"SHIP_LIGHT_SHUTTLE",
	"SHIP_ORE_HOUND",
	"SHIP_REFINING_FREIGHTER",
	"SHIP_SURVEYOR",
}

func (v *ShipType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ShipType(value)
	for _, existing := range AllowedShipTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ShipType", value)
}

// NewShipTypeFromValue returns a pointer to a valid ShipType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShipTypeFromValue(v string) (*ShipType, error) {
	ev := ShipType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShipType: valid values are %v", v, AllowedShipTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShipType) IsValid() bool {
	for _, existing := range AllowedShipTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ShipType value
func (v ShipType) Ptr() *ShipType {
	return &v
}

type NullableShipType struct {
	value *ShipType
	isSet bool
}

func (v NullableShipType) Get() *ShipType {
	return v.value
}

func (v *NullableShipType) Set(val *ShipType) {
	v.value = val
	v.isSet = true
}

func (v NullableShipType) IsSet() bool {
	return v.isSet
}

func (v *NullableShipType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipType(val *ShipType) *NullableShipType {
	return &NullableShipType{value: val, isSet: true}
}

func (v NullableShipType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
