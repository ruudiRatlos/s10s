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

// checks if the ShipNav type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipNav{}

// ShipNav The navigation information of the ship.
type ShipNav struct {
	// The symbol of the system.
	SystemSymbol string `json:"systemSymbol"`
	// The symbol of the waypoint.
	WaypointSymbol       string            `json:"waypointSymbol"`
	Route                ShipNavRoute      `json:"route"`
	Status               ShipNavStatus     `json:"status"`
	FlightMode           ShipNavFlightMode `json:"flightMode"`
	AdditionalProperties map[string]interface{}
}

type _ShipNav ShipNav

// NewShipNav instantiates a new ShipNav object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipNav(systemSymbol string, waypointSymbol string, route ShipNavRoute, status ShipNavStatus, flightMode ShipNavFlightMode) *ShipNav {
	this := ShipNav{}
	this.SystemSymbol = systemSymbol
	this.WaypointSymbol = waypointSymbol
	this.Route = route
	this.Status = status
	this.FlightMode = flightMode
	return &this
}

// NewShipNavWithDefaults instantiates a new ShipNav object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipNavWithDefaults() *ShipNav {
	this := ShipNav{}
	var flightMode ShipNavFlightMode = SHIPNAVFLIGHTMODE_CRUISE
	this.FlightMode = flightMode
	return &this
}

// GetSystemSymbol returns the SystemSymbol field value
func (o *ShipNav) GetSystemSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemSymbol
}

// GetSystemSymbolOk returns a tuple with the SystemSymbol field value
// and a boolean to check if the value has been set.
func (o *ShipNav) GetSystemSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemSymbol, true
}

// SetSystemSymbol sets field value
func (o *ShipNav) SetSystemSymbol(v string) {
	o.SystemSymbol = v
}

// GetWaypointSymbol returns the WaypointSymbol field value
func (o *ShipNav) GetWaypointSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WaypointSymbol
}

// GetWaypointSymbolOk returns a tuple with the WaypointSymbol field value
// and a boolean to check if the value has been set.
func (o *ShipNav) GetWaypointSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaypointSymbol, true
}

// SetWaypointSymbol sets field value
func (o *ShipNav) SetWaypointSymbol(v string) {
	o.WaypointSymbol = v
}

// GetRoute returns the Route field value
func (o *ShipNav) GetRoute() ShipNavRoute {
	if o == nil {
		var ret ShipNavRoute
		return ret
	}

	return o.Route
}

// GetRouteOk returns a tuple with the Route field value
// and a boolean to check if the value has been set.
func (o *ShipNav) GetRouteOk() (*ShipNavRoute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Route, true
}

// SetRoute sets field value
func (o *ShipNav) SetRoute(v ShipNavRoute) {
	o.Route = v
}

// GetStatus returns the Status field value
func (o *ShipNav) GetStatus() ShipNavStatus {
	if o == nil {
		var ret ShipNavStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ShipNav) GetStatusOk() (*ShipNavStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ShipNav) SetStatus(v ShipNavStatus) {
	o.Status = v
}

// GetFlightMode returns the FlightMode field value
func (o *ShipNav) GetFlightMode() ShipNavFlightMode {
	if o == nil {
		var ret ShipNavFlightMode
		return ret
	}

	return o.FlightMode
}

// GetFlightModeOk returns a tuple with the FlightMode field value
// and a boolean to check if the value has been set.
func (o *ShipNav) GetFlightModeOk() (*ShipNavFlightMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlightMode, true
}

// SetFlightMode sets field value
func (o *ShipNav) SetFlightMode(v ShipNavFlightMode) {
	o.FlightMode = v
}

func (o ShipNav) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipNav) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["systemSymbol"] = o.SystemSymbol
	toSerialize["waypointSymbol"] = o.WaypointSymbol
	toSerialize["route"] = o.Route
	toSerialize["status"] = o.Status
	toSerialize["flightMode"] = o.FlightMode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ShipNav) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"systemSymbol",
		"waypointSymbol",
		"route",
		"status",
		"flightMode",
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

	varShipNav := _ShipNav{}

	err = json.Unmarshal(data, &varShipNav)

	if err != nil {
		return err
	}

	*o = ShipNav(varShipNav)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "systemSymbol")
		delete(additionalProperties, "waypointSymbol")
		delete(additionalProperties, "route")
		delete(additionalProperties, "status")
		delete(additionalProperties, "flightMode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableShipNav struct {
	value *ShipNav
	isSet bool
}

func (v NullableShipNav) Get() *ShipNav {
	return v.value
}

func (v *NullableShipNav) Set(val *ShipNav) {
	v.value = val
	v.isSet = true
}

func (v NullableShipNav) IsSet() bool {
	return v.isSet
}

func (v *NullableShipNav) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipNav(val *ShipNav) *NullableShipNav {
	return &NullableShipNav{value: val, isSet: true}
}

func (v NullableShipNav) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipNav) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
