# ShipmentUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnHold** | Pointer to **bool** | Send this attribute if you want to put this shipment on hold. | [optional] 
**Picking** | Pointer to **bool** | Send this attribute if you want to start picking this shipment. | [optional] 
**Packing** | Pointer to **bool** | Send this attribute if you want to start packing this shipment. | [optional] 
**ReadyToShip** | Pointer to **bool** | Send this attribute if you want to mark this shipment as ready to ship. | [optional] 
**Ship** | Pointer to **bool** | Send this attribute if you want to mark this shipment as shipped. | [optional] 
**GetRates** | Pointer to **bool** | Send this attribute if you want get the shipping rates from the associated carrier accounts. | [optional] 
**SelectedRateId** | Pointer to **string** | The selected purchase rate from the available shipping rates. | [optional] 
**Purchase** | Pointer to **bool** | Send this attribute if you want to purchase this shipment with the selected rate. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewShipmentUpdateDataAttributes

`func NewShipmentUpdateDataAttributes() *ShipmentUpdateDataAttributes`

NewShipmentUpdateDataAttributes instantiates a new ShipmentUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentUpdateDataAttributesWithDefaults

`func NewShipmentUpdateDataAttributesWithDefaults() *ShipmentUpdateDataAttributes`

NewShipmentUpdateDataAttributesWithDefaults instantiates a new ShipmentUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnHold

`func (o *ShipmentUpdateDataAttributes) GetOnHold() bool`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *ShipmentUpdateDataAttributes) GetOnHoldOk() (*bool, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *ShipmentUpdateDataAttributes) SetOnHold(v bool)`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *ShipmentUpdateDataAttributes) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### GetPicking

`func (o *ShipmentUpdateDataAttributes) GetPicking() bool`

GetPicking returns the Picking field if non-nil, zero value otherwise.

### GetPickingOk

`func (o *ShipmentUpdateDataAttributes) GetPickingOk() (*bool, bool)`

GetPickingOk returns a tuple with the Picking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicking

`func (o *ShipmentUpdateDataAttributes) SetPicking(v bool)`

SetPicking sets Picking field to given value.

### HasPicking

`func (o *ShipmentUpdateDataAttributes) HasPicking() bool`

HasPicking returns a boolean if a field has been set.

### GetPacking

`func (o *ShipmentUpdateDataAttributes) GetPacking() bool`

GetPacking returns the Packing field if non-nil, zero value otherwise.

### GetPackingOk

`func (o *ShipmentUpdateDataAttributes) GetPackingOk() (*bool, bool)`

GetPackingOk returns a tuple with the Packing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacking

`func (o *ShipmentUpdateDataAttributes) SetPacking(v bool)`

SetPacking sets Packing field to given value.

### HasPacking

`func (o *ShipmentUpdateDataAttributes) HasPacking() bool`

HasPacking returns a boolean if a field has been set.

### GetReadyToShip

`func (o *ShipmentUpdateDataAttributes) GetReadyToShip() bool`

GetReadyToShip returns the ReadyToShip field if non-nil, zero value otherwise.

### GetReadyToShipOk

`func (o *ShipmentUpdateDataAttributes) GetReadyToShipOk() (*bool, bool)`

GetReadyToShipOk returns a tuple with the ReadyToShip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyToShip

`func (o *ShipmentUpdateDataAttributes) SetReadyToShip(v bool)`

SetReadyToShip sets ReadyToShip field to given value.

### HasReadyToShip

`func (o *ShipmentUpdateDataAttributes) HasReadyToShip() bool`

HasReadyToShip returns a boolean if a field has been set.

### GetShip

`func (o *ShipmentUpdateDataAttributes) GetShip() bool`

GetShip returns the Ship field if non-nil, zero value otherwise.

### GetShipOk

`func (o *ShipmentUpdateDataAttributes) GetShipOk() (*bool, bool)`

GetShipOk returns a tuple with the Ship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShip

`func (o *ShipmentUpdateDataAttributes) SetShip(v bool)`

SetShip sets Ship field to given value.

### HasShip

`func (o *ShipmentUpdateDataAttributes) HasShip() bool`

HasShip returns a boolean if a field has been set.

### GetGetRates

`func (o *ShipmentUpdateDataAttributes) GetGetRates() bool`

GetGetRates returns the GetRates field if non-nil, zero value otherwise.

### GetGetRatesOk

`func (o *ShipmentUpdateDataAttributes) GetGetRatesOk() (*bool, bool)`

GetGetRatesOk returns a tuple with the GetRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRates

`func (o *ShipmentUpdateDataAttributes) SetGetRates(v bool)`

SetGetRates sets GetRates field to given value.

### HasGetRates

`func (o *ShipmentUpdateDataAttributes) HasGetRates() bool`

HasGetRates returns a boolean if a field has been set.

### GetSelectedRateId

`func (o *ShipmentUpdateDataAttributes) GetSelectedRateId() string`

GetSelectedRateId returns the SelectedRateId field if non-nil, zero value otherwise.

### GetSelectedRateIdOk

`func (o *ShipmentUpdateDataAttributes) GetSelectedRateIdOk() (*string, bool)`

GetSelectedRateIdOk returns a tuple with the SelectedRateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRateId

`func (o *ShipmentUpdateDataAttributes) SetSelectedRateId(v string)`

SetSelectedRateId sets SelectedRateId field to given value.

### HasSelectedRateId

`func (o *ShipmentUpdateDataAttributes) HasSelectedRateId() bool`

HasSelectedRateId returns a boolean if a field has been set.

### GetPurchase

`func (o *ShipmentUpdateDataAttributes) GetPurchase() bool`

GetPurchase returns the Purchase field if non-nil, zero value otherwise.

### GetPurchaseOk

`func (o *ShipmentUpdateDataAttributes) GetPurchaseOk() (*bool, bool)`

GetPurchaseOk returns a tuple with the Purchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchase

`func (o *ShipmentUpdateDataAttributes) SetPurchase(v bool)`

SetPurchase sets Purchase field to given value.

### HasPurchase

`func (o *ShipmentUpdateDataAttributes) HasPurchase() bool`

HasPurchase returns a boolean if a field has been set.

### GetReference

`func (o *ShipmentUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ShipmentUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ShipmentUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ShipmentUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *ShipmentUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *ShipmentUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *ShipmentUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *ShipmentUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *ShipmentUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShipmentUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShipmentUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ShipmentUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


