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

// checks if the Ship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ship{}

// Ship Ship details.
type Ship struct {
	// The globally unique identifier of the ship in the following format: `[AGENT_SYMBOL]-[HEX_ID]`
	Symbol string `json:"symbol"`
	Registration ShipRegistration `json:"registration"`
	Nav ShipNav `json:"nav"`
	Crew ShipCrew `json:"crew"`
	Frame ShipFrame `json:"frame"`
	Reactor ShipReactor `json:"reactor"`
	Engine ShipEngine `json:"engine"`
	Cooldown Cooldown `json:"cooldown"`
	// Modules installed in this ship.
	Modules []ShipModule `json:"modules"`
	// Mounts installed in this ship.
	Mounts []ShipMount `json:"mounts"`
	Cargo ShipCargo `json:"cargo"`
	Fuel ShipFuel `json:"fuel"`
	AdditionalProperties map[string]interface{}
}

type _Ship Ship

// NewShip instantiates a new Ship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShip(symbol string, registration ShipRegistration, nav ShipNav, crew ShipCrew, frame ShipFrame, reactor ShipReactor, engine ShipEngine, cooldown Cooldown, modules []ShipModule, mounts []ShipMount, cargo ShipCargo, fuel ShipFuel) *Ship {
	this := Ship{}
	this.Symbol = symbol
	this.Registration = registration
	this.Nav = nav
	this.Crew = crew
	this.Frame = frame
	this.Reactor = reactor
	this.Engine = engine
	this.Cooldown = cooldown
	this.Modules = modules
	this.Mounts = mounts
	this.Cargo = cargo
	this.Fuel = fuel
	return &this
}

// NewShipWithDefaults instantiates a new Ship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipWithDefaults() *Ship {
	this := Ship{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *Ship) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *Ship) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *Ship) SetSymbol(v string) {
	o.Symbol = v
}

// GetRegistration returns the Registration field value
func (o *Ship) GetRegistration() ShipRegistration {
	if o == nil {
		var ret ShipRegistration
		return ret
	}

	return o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value
// and a boolean to check if the value has been set.
func (o *Ship) GetRegistrationOk() (*ShipRegistration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Registration, true
}

// SetRegistration sets field value
func (o *Ship) SetRegistration(v ShipRegistration) {
	o.Registration = v
}

// GetNav returns the Nav field value
func (o *Ship) GetNav() ShipNav {
	if o == nil {
		var ret ShipNav
		return ret
	}

	return o.Nav
}

// GetNavOk returns a tuple with the Nav field value
// and a boolean to check if the value has been set.
func (o *Ship) GetNavOk() (*ShipNav, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nav, true
}

// SetNav sets field value
func (o *Ship) SetNav(v ShipNav) {
	o.Nav = v
}

// GetCrew returns the Crew field value
func (o *Ship) GetCrew() ShipCrew {
	if o == nil {
		var ret ShipCrew
		return ret
	}

	return o.Crew
}

// GetCrewOk returns a tuple with the Crew field value
// and a boolean to check if the value has been set.
func (o *Ship) GetCrewOk() (*ShipCrew, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Crew, true
}

// SetCrew sets field value
func (o *Ship) SetCrew(v ShipCrew) {
	o.Crew = v
}

// GetFrame returns the Frame field value
func (o *Ship) GetFrame() ShipFrame {
	if o == nil {
		var ret ShipFrame
		return ret
	}

	return o.Frame
}

// GetFrameOk returns a tuple with the Frame field value
// and a boolean to check if the value has been set.
func (o *Ship) GetFrameOk() (*ShipFrame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frame, true
}

// SetFrame sets field value
func (o *Ship) SetFrame(v ShipFrame) {
	o.Frame = v
}

// GetReactor returns the Reactor field value
func (o *Ship) GetReactor() ShipReactor {
	if o == nil {
		var ret ShipReactor
		return ret
	}

	return o.Reactor
}

// GetReactorOk returns a tuple with the Reactor field value
// and a boolean to check if the value has been set.
func (o *Ship) GetReactorOk() (*ShipReactor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reactor, true
}

// SetReactor sets field value
func (o *Ship) SetReactor(v ShipReactor) {
	o.Reactor = v
}

// GetEngine returns the Engine field value
func (o *Ship) GetEngine() ShipEngine {
	if o == nil {
		var ret ShipEngine
		return ret
	}

	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *Ship) GetEngineOk() (*ShipEngine, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value
func (o *Ship) SetEngine(v ShipEngine) {
	o.Engine = v
}

// GetCooldown returns the Cooldown field value
func (o *Ship) GetCooldown() Cooldown {
	if o == nil {
		var ret Cooldown
		return ret
	}

	return o.Cooldown
}

// GetCooldownOk returns a tuple with the Cooldown field value
// and a boolean to check if the value has been set.
func (o *Ship) GetCooldownOk() (*Cooldown, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cooldown, true
}

// SetCooldown sets field value
func (o *Ship) SetCooldown(v Cooldown) {
	o.Cooldown = v
}

// GetModules returns the Modules field value
func (o *Ship) GetModules() []ShipModule {
	if o == nil {
		var ret []ShipModule
		return ret
	}

	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value
// and a boolean to check if the value has been set.
func (o *Ship) GetModulesOk() ([]ShipModule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modules, true
}

// SetModules sets field value
func (o *Ship) SetModules(v []ShipModule) {
	o.Modules = v
}

// GetMounts returns the Mounts field value
func (o *Ship) GetMounts() []ShipMount {
	if o == nil {
		var ret []ShipMount
		return ret
	}

	return o.Mounts
}

// GetMountsOk returns a tuple with the Mounts field value
// and a boolean to check if the value has been set.
func (o *Ship) GetMountsOk() ([]ShipMount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mounts, true
}

// SetMounts sets field value
func (o *Ship) SetMounts(v []ShipMount) {
	o.Mounts = v
}

// GetCargo returns the Cargo field value
func (o *Ship) GetCargo() ShipCargo {
	if o == nil {
		var ret ShipCargo
		return ret
	}

	return o.Cargo
}

// GetCargoOk returns a tuple with the Cargo field value
// and a boolean to check if the value has been set.
func (o *Ship) GetCargoOk() (*ShipCargo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cargo, true
}

// SetCargo sets field value
func (o *Ship) SetCargo(v ShipCargo) {
	o.Cargo = v
}

// GetFuel returns the Fuel field value
func (o *Ship) GetFuel() ShipFuel {
	if o == nil {
		var ret ShipFuel
		return ret
	}

	return o.Fuel
}

// GetFuelOk returns a tuple with the Fuel field value
// and a boolean to check if the value has been set.
func (o *Ship) GetFuelOk() (*ShipFuel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fuel, true
}

// SetFuel sets field value
func (o *Ship) SetFuel(v ShipFuel) {
	o.Fuel = v
}

func (o Ship) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["registration"] = o.Registration
	toSerialize["nav"] = o.Nav
	toSerialize["crew"] = o.Crew
	toSerialize["frame"] = o.Frame
	toSerialize["reactor"] = o.Reactor
	toSerialize["engine"] = o.Engine
	toSerialize["cooldown"] = o.Cooldown
	toSerialize["modules"] = o.Modules
	toSerialize["mounts"] = o.Mounts
	toSerialize["cargo"] = o.Cargo
	toSerialize["fuel"] = o.Fuel

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Ship) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"registration",
		"nav",
		"crew",
		"frame",
		"reactor",
		"engine",
		"cooldown",
		"modules",
		"mounts",
		"cargo",
		"fuel",
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

	varShip := _Ship{}

	err = json.Unmarshal(data, &varShip)

	if err != nil {
		return err
	}

	*o = Ship(varShip)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "registration")
		delete(additionalProperties, "nav")
		delete(additionalProperties, "crew")
		delete(additionalProperties, "frame")
		delete(additionalProperties, "reactor")
		delete(additionalProperties, "engine")
		delete(additionalProperties, "cooldown")
		delete(additionalProperties, "modules")
		delete(additionalProperties, "mounts")
		delete(additionalProperties, "cargo")
		delete(additionalProperties, "fuel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableShip struct {
	value *Ship
	isSet bool
}

func (v NullableShip) Get() *Ship {
	return v.value
}

func (v *NullableShip) Set(val *Ship) {
	v.value = val
	v.isSet = true
}

func (v NullableShip) IsSet() bool {
	return v.isSet
}

func (v *NullableShip) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShip(val *Ship) *NullableShip {
	return &NullableShip{value: val, isSet: true}
}

func (v NullableShip) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShip) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


