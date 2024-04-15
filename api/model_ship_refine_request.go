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

// checks if the ShipRefineRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipRefineRequest{}

// ShipRefineRequest struct for ShipRefineRequest
type ShipRefineRequest struct {
	// The type of good to produce out of the refining process.
	Produce              string `json:"produce"`
	AdditionalProperties map[string]interface{}
}

type _ShipRefineRequest ShipRefineRequest

// NewShipRefineRequest instantiates a new ShipRefineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipRefineRequest(produce string) *ShipRefineRequest {
	this := ShipRefineRequest{}
	this.Produce = produce
	return &this
}

// NewShipRefineRequestWithDefaults instantiates a new ShipRefineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipRefineRequestWithDefaults() *ShipRefineRequest {
	this := ShipRefineRequest{}
	return &this
}

// GetProduce returns the Produce field value
func (o *ShipRefineRequest) GetProduce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Produce
}

// GetProduceOk returns a tuple with the Produce field value
// and a boolean to check if the value has been set.
func (o *ShipRefineRequest) GetProduceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Produce, true
}

// SetProduce sets field value
func (o *ShipRefineRequest) SetProduce(v string) {
	o.Produce = v
}

func (o ShipRefineRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipRefineRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["produce"] = o.Produce

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ShipRefineRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"produce",
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

	varShipRefineRequest := _ShipRefineRequest{}

	err = json.Unmarshal(data, &varShipRefineRequest)

	if err != nil {
		return err
	}

	*o = ShipRefineRequest(varShipRefineRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "produce")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableShipRefineRequest struct {
	value *ShipRefineRequest
	isSet bool
}

func (v NullableShipRefineRequest) Get() *ShipRefineRequest {
	return v.value
}

func (v *NullableShipRefineRequest) Set(val *ShipRefineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableShipRefineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableShipRefineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipRefineRequest(val *ShipRefineRequest) *NullableShipRefineRequest {
	return &NullableShipRefineRequest{value: val, isSet: true}
}

func (v NullableShipRefineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipRefineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
