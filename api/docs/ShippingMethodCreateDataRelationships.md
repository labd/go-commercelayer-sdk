# ShippingMethodCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**AvalaraAccountDataRelationshipsMarkets**](AvalaraAccountDataRelationshipsMarkets.md) |  | [optional] 
**ShippingZone** | [**ShippingMethodDataRelationshipsShippingZone**](ShippingMethodDataRelationshipsShippingZone.md) |  | 
**ShippingCategory** | [**ShipmentDataRelationshipsShippingCategory**](ShipmentDataRelationshipsShippingCategory.md) |  | 

## Methods

### NewShippingMethodCreateDataRelationships

`func NewShippingMethodCreateDataRelationships(shippingZone ShippingMethodDataRelationshipsShippingZone, shippingCategory ShipmentDataRelationshipsShippingCategory, ) *ShippingMethodCreateDataRelationships`

NewShippingMethodCreateDataRelationships instantiates a new ShippingMethodCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingMethodCreateDataRelationshipsWithDefaults

`func NewShippingMethodCreateDataRelationshipsWithDefaults() *ShippingMethodCreateDataRelationships`

NewShippingMethodCreateDataRelationshipsWithDefaults instantiates a new ShippingMethodCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *ShippingMethodCreateDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *ShippingMethodCreateDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *ShippingMethodCreateDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *ShippingMethodCreateDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetShippingZone

`func (o *ShippingMethodCreateDataRelationships) GetShippingZone() ShippingMethodDataRelationshipsShippingZone`

GetShippingZone returns the ShippingZone field if non-nil, zero value otherwise.

### GetShippingZoneOk

`func (o *ShippingMethodCreateDataRelationships) GetShippingZoneOk() (*ShippingMethodDataRelationshipsShippingZone, bool)`

GetShippingZoneOk returns a tuple with the ShippingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingZone

`func (o *ShippingMethodCreateDataRelationships) SetShippingZone(v ShippingMethodDataRelationshipsShippingZone)`

SetShippingZone sets ShippingZone field to given value.


### GetShippingCategory

`func (o *ShippingMethodCreateDataRelationships) GetShippingCategory() ShipmentDataRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *ShippingMethodCreateDataRelationships) GetShippingCategoryOk() (*ShipmentDataRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *ShippingMethodCreateDataRelationships) SetShippingCategory(v ShipmentDataRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


