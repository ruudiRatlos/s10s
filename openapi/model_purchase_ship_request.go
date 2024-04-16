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

// checks if the PurchaseShipRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseShipRequest{}

// PurchaseShipRequest struct for PurchaseShipRequest
type PurchaseShipRequest struct {
	ShipType ShipType `json:"shipType"`
	// The symbol of the waypoint you want to purchase the ship at.
	WaypointSymbol       string `json:"waypointSymbol"`
	AdditionalProperties map[string]interface{}
}

type _PurchaseShipRequest PurchaseShipRequest

// NewPurchaseShipRequest instantiates a new PurchaseShipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseShipRequest(shipType ShipType, waypointSymbol string) *PurchaseShipRequest {
	this := PurchaseShipRequest{}
	this.ShipType = shipType
	this.WaypointSymbol = waypointSymbol
	return &this
}

// NewPurchaseShipRequestWithDefaults instantiates a new PurchaseShipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseShipRequestWithDefaults() *PurchaseShipRequest {
	this := PurchaseShipRequest{}
	return &this
}

// GetShipType returns the ShipType field value
func (o *PurchaseShipRequest) GetShipType() ShipType {
	if o == nil {
		var ret ShipType
		return ret
	}

	return o.ShipType
}

// GetShipTypeOk returns a tuple with the ShipType field value
// and a boolean to check if the value has been set.
func (o *PurchaseShipRequest) GetShipTypeOk() (*ShipType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipType, true
}

// SetShipType sets field value
func (o *PurchaseShipRequest) SetShipType(v ShipType) {
	o.ShipType = v
}

// GetWaypointSymbol returns the WaypointSymbol field value
func (o *PurchaseShipRequest) GetWaypointSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WaypointSymbol
}

// GetWaypointSymbolOk returns a tuple with the WaypointSymbol field value
// and a boolean to check if the value has been set.
func (o *PurchaseShipRequest) GetWaypointSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaypointSymbol, true
}

// SetWaypointSymbol sets field value
func (o *PurchaseShipRequest) SetWaypointSymbol(v string) {
	o.WaypointSymbol = v
}

func (o PurchaseShipRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PurchaseShipRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipType"] = o.ShipType
	toSerialize["waypointSymbol"] = o.WaypointSymbol

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PurchaseShipRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"shipType",
		"waypointSymbol",
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

	varPurchaseShipRequest := _PurchaseShipRequest{}

	err = json.Unmarshal(data, &varPurchaseShipRequest)

	if err != nil {
		return err
	}

	*o = PurchaseShipRequest(varPurchaseShipRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "shipType")
		delete(additionalProperties, "waypointSymbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePurchaseShipRequest struct {
	value *PurchaseShipRequest
	isSet bool
}

func (v NullablePurchaseShipRequest) Get() *PurchaseShipRequest {
	return v.value
}

func (v *NullablePurchaseShipRequest) Set(val *PurchaseShipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseShipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseShipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseShipRequest(val *PurchaseShipRequest) *NullablePurchaseShipRequest {
	return &NullablePurchaseShipRequest{value: val, isSet: true}
}

func (v NullablePurchaseShipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseShipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
