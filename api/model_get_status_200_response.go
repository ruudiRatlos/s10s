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

// checks if the GetStatus200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatus200Response{}

// GetStatus200Response struct for GetStatus200Response
type GetStatus200Response struct {
	// The current status of the game server.
	Status string `json:"status"`
	// The current version of the API.
	Version string `json:"version"`
	// The date when the game server was last reset.
	ResetDate            string                                   `json:"resetDate"`
	Description          string                                   `json:"description"`
	Stats                GetStatus200ResponseStats                `json:"stats"`
	Leaderboards         GetStatus200ResponseLeaderboards         `json:"leaderboards"`
	ServerResets         GetStatus200ResponseServerResets         `json:"serverResets"`
	Announcements        []GetStatus200ResponseAnnouncementsInner `json:"announcements"`
	Links                []GetStatus200ResponseLinksInner         `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _GetStatus200Response GetStatus200Response

// NewGetStatus200Response instantiates a new GetStatus200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatus200Response(status string, version string, resetDate string, description string, stats GetStatus200ResponseStats, leaderboards GetStatus200ResponseLeaderboards, serverResets GetStatus200ResponseServerResets, announcements []GetStatus200ResponseAnnouncementsInner, links []GetStatus200ResponseLinksInner) *GetStatus200Response {
	this := GetStatus200Response{}
	this.Status = status
	this.Version = version
	this.ResetDate = resetDate
	this.Description = description
	this.Stats = stats
	this.Leaderboards = leaderboards
	this.ServerResets = serverResets
	this.Announcements = announcements
	this.Links = links
	return &this
}

// NewGetStatus200ResponseWithDefaults instantiates a new GetStatus200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatus200ResponseWithDefaults() *GetStatus200Response {
	this := GetStatus200Response{}
	return &this
}

// GetStatus returns the Status field value
func (o *GetStatus200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetStatus200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetStatus200Response) SetStatus(v string) {
	o.Status = v
}

// GetVersion returns the Version field value
func (o *GetStatus200Response) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetStatus200Response) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetStatus200Response) SetVersion(v string) {
	o.Version = v
}

// GetResetDate returns the ResetDate field value
func (o *GetStatus200Response) GetResetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResetDate
}

// GetResetDateOk returns a tuple with the ResetDate field value
// and a boolean to check if the value has been set.
func (o *GetStatus200Response) GetResetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResetDate, true
}

// SetResetDate sets field value
func (o *GetStatus200Response) SetResetDate(v string) {
	o.ResetDate = v
}

// GetDescription returns the Description field value
func (o *GetStatus200Response) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GetStatus200Response) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GetStatus200Response) SetDescription(v string) {
	o.Description = v
}

// GetStats returns the Stats field value
func (o *GetStatus200Response) GetStats() GetStatus200ResponseStats {
	if o == nil {
		var ret GetStatus200ResponseStats
		return ret
	}

	return o.Stats
}

// GetStatsOk returns a tuple with the Stats field value
// and a boolean to check if the value has been set.
func (o *GetStatus200Response) GetStatsOk() (*GetStatus200ResponseStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stats, true
}

// SetStats sets field value
func (o *GetStatus200Response) SetStats(v GetStatus200ResponseStats) {
	o.Stats = v
}

// GetLeaderboards returns the Leaderboards field value
func (o *GetStatus200Response) GetLeaderboards() GetStatus200ResponseLeaderboards {
	if o == nil {
		var ret GetStatus200ResponseLeaderboards
		return ret
	}

	return o.Leaderboards
}

// GetLeaderboardsOk returns a tuple with the Leaderboards field value
// and a boolean to check if the value has been set.
func (o *GetStatus200Response) GetLeaderboardsOk() (*GetStatus200ResponseLeaderboards, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Leaderboards, true
}

// SetLeaderboards sets field value
func (o *GetStatus200Response) SetLeaderboards(v GetStatus200ResponseLeaderboards) {
	o.Leaderboards = v
}

// GetServerResets returns the ServerResets field value
func (o *GetStatus200Response) GetServerResets() GetStatus200ResponseServerResets {
	if o == nil {
		var ret GetStatus200ResponseServerResets
		return ret
	}

	return o.ServerResets
}

// GetServerResetsOk returns a tuple with the ServerResets field value
// and a boolean to check if the value has been set.
func (o *GetStatus200Response) GetServerResetsOk() (*GetStatus200ResponseServerResets, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerResets, true
}

// SetServerResets sets field value
func (o *GetStatus200Response) SetServerResets(v GetStatus200ResponseServerResets) {
	o.ServerResets = v
}

// GetAnnouncements returns the Announcements field value
func (o *GetStatus200Response) GetAnnouncements() []GetStatus200ResponseAnnouncementsInner {
	if o == nil {
		var ret []GetStatus200ResponseAnnouncementsInner
		return ret
	}

	return o.Announcements
}

// GetAnnouncementsOk returns a tuple with the Announcements field value
// and a boolean to check if the value has been set.
func (o *GetStatus200Response) GetAnnouncementsOk() ([]GetStatus200ResponseAnnouncementsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Announcements, true
}

// SetAnnouncements sets field value
func (o *GetStatus200Response) SetAnnouncements(v []GetStatus200ResponseAnnouncementsInner) {
	o.Announcements = v
}

// GetLinks returns the Links field value
func (o *GetStatus200Response) GetLinks() []GetStatus200ResponseLinksInner {
	if o == nil {
		var ret []GetStatus200ResponseLinksInner
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetStatus200Response) GetLinksOk() ([]GetStatus200ResponseLinksInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *GetStatus200Response) SetLinks(v []GetStatus200ResponseLinksInner) {
	o.Links = v
}

func (o GetStatus200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatus200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["version"] = o.Version
	toSerialize["resetDate"] = o.ResetDate
	toSerialize["description"] = o.Description
	toSerialize["stats"] = o.Stats
	toSerialize["leaderboards"] = o.Leaderboards
	toSerialize["serverResets"] = o.ServerResets
	toSerialize["announcements"] = o.Announcements
	toSerialize["links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetStatus200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"version",
		"resetDate",
		"description",
		"stats",
		"leaderboards",
		"serverResets",
		"announcements",
		"links",
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

	varGetStatus200Response := _GetStatus200Response{}

	err = json.Unmarshal(data, &varGetStatus200Response)

	if err != nil {
		return err
	}

	*o = GetStatus200Response(varGetStatus200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "version")
		delete(additionalProperties, "resetDate")
		delete(additionalProperties, "description")
		delete(additionalProperties, "stats")
		delete(additionalProperties, "leaderboards")
		delete(additionalProperties, "serverResets")
		delete(additionalProperties, "announcements")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetStatus200Response struct {
	value *GetStatus200Response
	isSet bool
}

func (v NullableGetStatus200Response) Get() *GetStatus200Response {
	return v.value
}

func (v *NullableGetStatus200Response) Set(val *GetStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatus200Response(val *GetStatus200Response) *NullableGetStatus200Response {
	return &NullableGetStatus200Response{value: val, isSet: true}
}

func (v NullableGetStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
