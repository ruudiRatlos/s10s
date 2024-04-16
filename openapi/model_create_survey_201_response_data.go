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

// checks if the CreateSurvey201ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSurvey201ResponseData{}

// CreateSurvey201ResponseData struct for CreateSurvey201ResponseData
type CreateSurvey201ResponseData struct {
	Cooldown Cooldown `json:"cooldown"`
	// Surveys created by this action.
	Surveys              []Survey `json:"surveys"`
	AdditionalProperties map[string]interface{}
}

type _CreateSurvey201ResponseData CreateSurvey201ResponseData

// NewCreateSurvey201ResponseData instantiates a new CreateSurvey201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSurvey201ResponseData(cooldown Cooldown, surveys []Survey) *CreateSurvey201ResponseData {
	this := CreateSurvey201ResponseData{}
	this.Cooldown = cooldown
	this.Surveys = surveys
	return &this
}

// NewCreateSurvey201ResponseDataWithDefaults instantiates a new CreateSurvey201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSurvey201ResponseDataWithDefaults() *CreateSurvey201ResponseData {
	this := CreateSurvey201ResponseData{}
	return &this
}

// GetCooldown returns the Cooldown field value
func (o *CreateSurvey201ResponseData) GetCooldown() Cooldown {
	if o == nil {
		var ret Cooldown
		return ret
	}

	return o.Cooldown
}

// GetCooldownOk returns a tuple with the Cooldown field value
// and a boolean to check if the value has been set.
func (o *CreateSurvey201ResponseData) GetCooldownOk() (*Cooldown, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cooldown, true
}

// SetCooldown sets field value
func (o *CreateSurvey201ResponseData) SetCooldown(v Cooldown) {
	o.Cooldown = v
}

// GetSurveys returns the Surveys field value
func (o *CreateSurvey201ResponseData) GetSurveys() []Survey {
	if o == nil {
		var ret []Survey
		return ret
	}

	return o.Surveys
}

// GetSurveysOk returns a tuple with the Surveys field value
// and a boolean to check if the value has been set.
func (o *CreateSurvey201ResponseData) GetSurveysOk() ([]Survey, bool) {
	if o == nil {
		return nil, false
	}
	return o.Surveys, true
}

// SetSurveys sets field value
func (o *CreateSurvey201ResponseData) SetSurveys(v []Survey) {
	o.Surveys = v
}

func (o CreateSurvey201ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSurvey201ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cooldown"] = o.Cooldown
	toSerialize["surveys"] = o.Surveys

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateSurvey201ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cooldown",
		"surveys",
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

	varCreateSurvey201ResponseData := _CreateSurvey201ResponseData{}

	err = json.Unmarshal(data, &varCreateSurvey201ResponseData)

	if err != nil {
		return err
	}

	*o = CreateSurvey201ResponseData(varCreateSurvey201ResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cooldown")
		delete(additionalProperties, "surveys")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSurvey201ResponseData struct {
	value *CreateSurvey201ResponseData
	isSet bool
}

func (v NullableCreateSurvey201ResponseData) Get() *CreateSurvey201ResponseData {
	return v.value
}

func (v *NullableCreateSurvey201ResponseData) Set(val *CreateSurvey201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSurvey201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSurvey201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSurvey201ResponseData(val *CreateSurvey201ResponseData) *NullableCreateSurvey201ResponseData {
	return &NullableCreateSurvey201ResponseData{value: val, isSet: true}
}

func (v NullableCreateSurvey201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSurvey201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
