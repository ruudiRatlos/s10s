# RepairTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WaypointSymbol** | **string** | The symbol of the waypoint. | 
**ShipSymbol** | **string** | The symbol of the ship. | 
**TotalPrice** | **int32** | The total price of the transaction. | 
**Timestamp** | **time.Time** | The timestamp of the transaction. | 

## Methods

### NewRepairTransaction

`func NewRepairTransaction(waypointSymbol string, shipSymbol string, totalPrice int32, timestamp time.Time, ) *RepairTransaction`

NewRepairTransaction instantiates a new RepairTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepairTransactionWithDefaults

`func NewRepairTransactionWithDefaults() *RepairTransaction`

NewRepairTransactionWithDefaults instantiates a new RepairTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWaypointSymbol

`func (o *RepairTransaction) GetWaypointSymbol() string`

GetWaypointSymbol returns the WaypointSymbol field if non-nil, zero value otherwise.

### GetWaypointSymbolOk

`func (o *RepairTransaction) GetWaypointSymbolOk() (*string, bool)`

GetWaypointSymbolOk returns a tuple with the WaypointSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypointSymbol

`func (o *RepairTransaction) SetWaypointSymbol(v string)`

SetWaypointSymbol sets WaypointSymbol field to given value.


### GetShipSymbol

`func (o *RepairTransaction) GetShipSymbol() string`

GetShipSymbol returns the ShipSymbol field if non-nil, zero value otherwise.

### GetShipSymbolOk

`func (o *RepairTransaction) GetShipSymbolOk() (*string, bool)`

GetShipSymbolOk returns a tuple with the ShipSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipSymbol

`func (o *RepairTransaction) SetShipSymbol(v string)`

SetShipSymbol sets ShipSymbol field to given value.


### GetTotalPrice

`func (o *RepairTransaction) GetTotalPrice() int32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *RepairTransaction) GetTotalPriceOk() (*int32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *RepairTransaction) SetTotalPrice(v int32)`

SetTotalPrice sets TotalPrice field to given value.


### GetTimestamp

`func (o *RepairTransaction) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RepairTransaction) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RepairTransaction) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


