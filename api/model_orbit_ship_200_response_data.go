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

// checks if the OrbitShip200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrbitShip200ResponseData{}

// OrbitShip200ResponseData struct for OrbitShip200ResponseData
type OrbitShip200ResponseData struct {
	Nav ShipNav `json:"nav"`
	AdditionalProperties map[string]interface{}
}

type _OrbitShip200ResponseData OrbitShip200ResponseData

// NewOrbitShip200ResponseData instantiates a new OrbitShip200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrbitShip200ResponseData(nav ShipNav) *OrbitShip200ResponseData {
	this := OrbitShip200ResponseData{}
	this.Nav = nav
	return &this
}

// NewOrbitShip200ResponseDataWithDefaults instantiates a new OrbitShip200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrbitShip200ResponseDataWithDefaults() *OrbitShip200ResponseData {
	this := OrbitShip200ResponseData{}
	return &this
}

// GetNav returns the Nav field value
func (o *OrbitShip200ResponseData) GetNav() ShipNav {
	if o == nil {
		var ret ShipNav
		return ret
	}

	return o.Nav
}

// GetNavOk returns a tuple with the Nav field value
// and a boolean to check if the value has been set.
func (o *OrbitShip200ResponseData) GetNavOk() (*ShipNav, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nav, true
}

// SetNav sets field value
func (o *OrbitShip200ResponseData) SetNav(v ShipNav) {
	o.Nav = v
}

func (o OrbitShip200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrbitShip200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nav"] = o.Nav

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrbitShip200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nav",
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

	varOrbitShip200ResponseData := _OrbitShip200ResponseData{}

	err = json.Unmarshal(data, &varOrbitShip200ResponseData)

	if err != nil {
		return err
	}

	*o = OrbitShip200ResponseData(varOrbitShip200ResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "nav")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrbitShip200ResponseData struct {
	value *OrbitShip200ResponseData
	isSet bool
}

func (v NullableOrbitShip200ResponseData) Get() *OrbitShip200ResponseData {
	return v.value
}

func (v *NullableOrbitShip200ResponseData) Set(val *OrbitShip200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrbitShip200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrbitShip200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrbitShip200ResponseData(val *OrbitShip200ResponseData) *NullableOrbitShip200ResponseData {
	return &NullableOrbitShip200ResponseData{value: val, isSet: true}
}

func (v NullableOrbitShip200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrbitShip200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


