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
	"time"
)

// checks if the ShipModificationTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipModificationTransaction{}

// ShipModificationTransaction Result of a transaction for a ship modification, such as installing a mount or a module.
type ShipModificationTransaction struct {
	// The symbol of the waypoint where the transaction took place.
	WaypointSymbol string `json:"waypointSymbol"`
	// The symbol of the ship that made the transaction.
	ShipSymbol string `json:"shipSymbol"`
	// The symbol of the trade good.
	TradeSymbol string `json:"tradeSymbol"`
	// The total price of the transaction.
	TotalPrice int32 `json:"totalPrice"`
	// The timestamp of the transaction.
	Timestamp            time.Time `json:"timestamp"`
	AdditionalProperties map[string]interface{}
}

type _ShipModificationTransaction ShipModificationTransaction

// NewShipModificationTransaction instantiates a new ShipModificationTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipModificationTransaction(waypointSymbol string, shipSymbol string, tradeSymbol string, totalPrice int32, timestamp time.Time) *ShipModificationTransaction {
	this := ShipModificationTransaction{}
	this.WaypointSymbol = waypointSymbol
	this.ShipSymbol = shipSymbol
	this.TradeSymbol = tradeSymbol
	this.TotalPrice = totalPrice
	this.Timestamp = timestamp
	return &this
}

// NewShipModificationTransactionWithDefaults instantiates a new ShipModificationTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipModificationTransactionWithDefaults() *ShipModificationTransaction {
	this := ShipModificationTransaction{}
	return &this
}

// GetWaypointSymbol returns the WaypointSymbol field value
func (o *ShipModificationTransaction) GetWaypointSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WaypointSymbol
}

// GetWaypointSymbolOk returns a tuple with the WaypointSymbol field value
// and a boolean to check if the value has been set.
func (o *ShipModificationTransaction) GetWaypointSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaypointSymbol, true
}

// SetWaypointSymbol sets field value
func (o *ShipModificationTransaction) SetWaypointSymbol(v string) {
	o.WaypointSymbol = v
}

// GetShipSymbol returns the ShipSymbol field value
func (o *ShipModificationTransaction) GetShipSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipSymbol
}

// GetShipSymbolOk returns a tuple with the ShipSymbol field value
// and a boolean to check if the value has been set.
func (o *ShipModificationTransaction) GetShipSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipSymbol, true
}

// SetShipSymbol sets field value
func (o *ShipModificationTransaction) SetShipSymbol(v string) {
	o.ShipSymbol = v
}

// GetTradeSymbol returns the TradeSymbol field value
func (o *ShipModificationTransaction) GetTradeSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TradeSymbol
}

// GetTradeSymbolOk returns a tuple with the TradeSymbol field value
// and a boolean to check if the value has been set.
func (o *ShipModificationTransaction) GetTradeSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeSymbol, true
}

// SetTradeSymbol sets field value
func (o *ShipModificationTransaction) SetTradeSymbol(v string) {
	o.TradeSymbol = v
}

// GetTotalPrice returns the TotalPrice field value
func (o *ShipModificationTransaction) GetTotalPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value
// and a boolean to check if the value has been set.
func (o *ShipModificationTransaction) GetTotalPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPrice, true
}

// SetTotalPrice sets field value
func (o *ShipModificationTransaction) SetTotalPrice(v int32) {
	o.TotalPrice = v
}

// GetTimestamp returns the Timestamp field value
func (o *ShipModificationTransaction) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ShipModificationTransaction) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ShipModificationTransaction) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o ShipModificationTransaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipModificationTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["waypointSymbol"] = o.WaypointSymbol
	toSerialize["shipSymbol"] = o.ShipSymbol
	toSerialize["tradeSymbol"] = o.TradeSymbol
	toSerialize["totalPrice"] = o.TotalPrice
	toSerialize["timestamp"] = o.Timestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ShipModificationTransaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"waypointSymbol",
		"shipSymbol",
		"tradeSymbol",
		"totalPrice",
		"timestamp",
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

	varShipModificationTransaction := _ShipModificationTransaction{}

	err = json.Unmarshal(data, &varShipModificationTransaction)

	if err != nil {
		return err
	}

	*o = ShipModificationTransaction(varShipModificationTransaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "waypointSymbol")
		delete(additionalProperties, "shipSymbol")
		delete(additionalProperties, "tradeSymbol")
		delete(additionalProperties, "totalPrice")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableShipModificationTransaction struct {
	value *ShipModificationTransaction
	isSet bool
}

func (v NullableShipModificationTransaction) Get() *ShipModificationTransaction {
	return v.value
}

func (v *NullableShipModificationTransaction) Set(val *ShipModificationTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableShipModificationTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableShipModificationTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipModificationTransaction(val *ShipModificationTransaction) *NullableShipModificationTransaction {
	return &NullableShipModificationTransaction{value: val, isSet: true}
}

func (v NullableShipModificationTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipModificationTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
