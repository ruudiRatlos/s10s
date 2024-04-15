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

// checks if the GetStatus200ResponseLeaderboards type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatus200ResponseLeaderboards{}

// GetStatus200ResponseLeaderboards struct for GetStatus200ResponseLeaderboards
type GetStatus200ResponseLeaderboards struct {
	// Top agents with the most credits.
	MostCredits []GetStatus200ResponseLeaderboardsMostCreditsInner `json:"mostCredits"`
	// Top agents with the most charted submitted.
	MostSubmittedCharts  []GetStatus200ResponseLeaderboardsMostSubmittedChartsInner `json:"mostSubmittedCharts"`
	AdditionalProperties map[string]interface{}
}

type _GetStatus200ResponseLeaderboards GetStatus200ResponseLeaderboards

// NewGetStatus200ResponseLeaderboards instantiates a new GetStatus200ResponseLeaderboards object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatus200ResponseLeaderboards(mostCredits []GetStatus200ResponseLeaderboardsMostCreditsInner, mostSubmittedCharts []GetStatus200ResponseLeaderboardsMostSubmittedChartsInner) *GetStatus200ResponseLeaderboards {
	this := GetStatus200ResponseLeaderboards{}
	this.MostCredits = mostCredits
	this.MostSubmittedCharts = mostSubmittedCharts
	return &this
}

// NewGetStatus200ResponseLeaderboardsWithDefaults instantiates a new GetStatus200ResponseLeaderboards object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatus200ResponseLeaderboardsWithDefaults() *GetStatus200ResponseLeaderboards {
	this := GetStatus200ResponseLeaderboards{}
	return &this
}

// GetMostCredits returns the MostCredits field value
func (o *GetStatus200ResponseLeaderboards) GetMostCredits() []GetStatus200ResponseLeaderboardsMostCreditsInner {
	if o == nil {
		var ret []GetStatus200ResponseLeaderboardsMostCreditsInner
		return ret
	}

	return o.MostCredits
}

// GetMostCreditsOk returns a tuple with the MostCredits field value
// and a boolean to check if the value has been set.
func (o *GetStatus200ResponseLeaderboards) GetMostCreditsOk() ([]GetStatus200ResponseLeaderboardsMostCreditsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.MostCredits, true
}

// SetMostCredits sets field value
func (o *GetStatus200ResponseLeaderboards) SetMostCredits(v []GetStatus200ResponseLeaderboardsMostCreditsInner) {
	o.MostCredits = v
}

// GetMostSubmittedCharts returns the MostSubmittedCharts field value
func (o *GetStatus200ResponseLeaderboards) GetMostSubmittedCharts() []GetStatus200ResponseLeaderboardsMostSubmittedChartsInner {
	if o == nil {
		var ret []GetStatus200ResponseLeaderboardsMostSubmittedChartsInner
		return ret
	}

	return o.MostSubmittedCharts
}

// GetMostSubmittedChartsOk returns a tuple with the MostSubmittedCharts field value
// and a boolean to check if the value has been set.
func (o *GetStatus200ResponseLeaderboards) GetMostSubmittedChartsOk() ([]GetStatus200ResponseLeaderboardsMostSubmittedChartsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.MostSubmittedCharts, true
}

// SetMostSubmittedCharts sets field value
func (o *GetStatus200ResponseLeaderboards) SetMostSubmittedCharts(v []GetStatus200ResponseLeaderboardsMostSubmittedChartsInner) {
	o.MostSubmittedCharts = v
}

func (o GetStatus200ResponseLeaderboards) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatus200ResponseLeaderboards) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mostCredits"] = o.MostCredits
	toSerialize["mostSubmittedCharts"] = o.MostSubmittedCharts

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetStatus200ResponseLeaderboards) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mostCredits",
		"mostSubmittedCharts",
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

	varGetStatus200ResponseLeaderboards := _GetStatus200ResponseLeaderboards{}

	err = json.Unmarshal(data, &varGetStatus200ResponseLeaderboards)

	if err != nil {
		return err
	}

	*o = GetStatus200ResponseLeaderboards(varGetStatus200ResponseLeaderboards)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mostCredits")
		delete(additionalProperties, "mostSubmittedCharts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetStatus200ResponseLeaderboards struct {
	value *GetStatus200ResponseLeaderboards
	isSet bool
}

func (v NullableGetStatus200ResponseLeaderboards) Get() *GetStatus200ResponseLeaderboards {
	return v.value
}

func (v *NullableGetStatus200ResponseLeaderboards) Set(val *GetStatus200ResponseLeaderboards) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatus200ResponseLeaderboards) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatus200ResponseLeaderboards) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatus200ResponseLeaderboards(val *GetStatus200ResponseLeaderboards) *NullableGetStatus200ResponseLeaderboards {
	return &NullableGetStatus200ResponseLeaderboards{value: val, isSet: true}
}

func (v NullableGetStatus200ResponseLeaderboards) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatus200ResponseLeaderboards) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
