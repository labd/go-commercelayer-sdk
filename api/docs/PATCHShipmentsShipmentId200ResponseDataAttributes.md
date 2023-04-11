# PATCHShipmentsShipmentId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnHold** | Pointer to **interface{}** | Send this attribute if you want to put this shipment on hold. | [optional] 
**Picking** | Pointer to **interface{}** | Send this attribute if you want to start picking this shipment. | [optional] 
**Packing** | Pointer to **interface{}** | Send this attribute if you want to start packing this shipment. | [optional] 
**ReadyToShip** | Pointer to **interface{}** | Send this attribute if you want to mark this shipment as ready to ship. | [optional] 
**Ship** | Pointer to **interface{}** | Send this attribute if you want to mark this shipment as shipped. | [optional] 
**GetRates** | Pointer to **interface{}** | Send this attribute if you want get the shipping rates from the associated carrier accounts. | [optional] 
**SelectedRateId** | Pointer to **interface{}** | The selected purchase rate from the available shipping rates. | [optional] 
**Purchase** | Pointer to **interface{}** | Send this attribute if you want to purchase this shipment with the selected rate. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHShipmentsShipmentId200ResponseDataAttributes

`func NewPATCHShipmentsShipmentId200ResponseDataAttributes() *PATCHShipmentsShipmentId200ResponseDataAttributes`

NewPATCHShipmentsShipmentId200ResponseDataAttributes instantiates a new PATCHShipmentsShipmentId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHShipmentsShipmentId200ResponseDataAttributesWithDefaults

`func NewPATCHShipmentsShipmentId200ResponseDataAttributesWithDefaults() *PATCHShipmentsShipmentId200ResponseDataAttributes`

NewPATCHShipmentsShipmentId200ResponseDataAttributesWithDefaults instantiates a new PATCHShipmentsShipmentId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnHold

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetOnHold() interface{}`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetOnHoldOk() (*interface{}, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetOnHold(v interface{})`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### SetOnHoldNil

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetOnHoldNil(b bool)`

 SetOnHoldNil sets the value for OnHold to be an explicit nil

### UnsetOnHold
`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) UnsetOnHold()`

UnsetOnHold ensures that no value is present for OnHold, not even an explicit nil
### GetPicking

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPicking() interface{}`

GetPicking returns the Picking field if non-nil, zero value otherwise.

### GetPickingOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPickingOk() (*interface{}, bool)`

GetPickingOk returns a tuple with the Picking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicking

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetPicking(v interface{})`

SetPicking sets Picking field to given value.

### HasPicking

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasPicking() bool`

HasPicking returns a boolean if a field has been set.

### SetPickingNil

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetPickingNil(b bool)`

 SetPickingNil sets the value for Picking to be an explicit nil

### UnsetPicking
`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) UnsetPicking()`

UnsetPicking ensures that no value is present for Picking, not even an explicit nil
### GetPacking

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPacking() interface{}`

GetPacking returns the Packing field if non-nil, zero value otherwise.

### GetPackingOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPackingOk() (*interface{}, bool)`

GetPackingOk returns a tuple with the Packing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacking

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetPacking(v interface{})`

SetPacking sets Packing field to given value.

### HasPacking

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasPacking() bool`

HasPacking returns a boolean if a field has been set.

### SetPackingNil

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetPackingNil(b bool)`

 SetPackingNil sets the value for Packing to be an explicit nil

### UnsetPacking
`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) UnsetPacking()`

UnsetPacking ensures that no value is present for Packing, not even an explicit nil
### GetReadyToShip

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReadyToShip() interface{}`

GetReadyToShip returns the ReadyToShip field if non-nil, zero value otherwise.

### GetReadyToShipOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReadyToShipOk() (*interface{}, bool)`

GetReadyToShipOk returns a tuple with the ReadyToShip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyToShip

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetReadyToShip(v interface{})`

SetReadyToShip sets ReadyToShip field to given value.

### HasReadyToShip

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasReadyToShip() bool`

HasReadyToShip returns a boolean if a field has been set.

### SetReadyToShipNil

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetReadyToShipNil(b bool)`

 SetReadyToShipNil sets the value for ReadyToShip to be an explicit nil

### UnsetReadyToShip
`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) UnsetReadyToShip()`

UnsetReadyToShip ensures that no value is present for ReadyToShip, not even an explicit nil
### GetShip

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetShip() interface{}`

GetShip returns the Ship field if non-nil, zero value otherwise.

### GetShipOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetShipOk() (*interface{}, bool)`

GetShipOk returns a tuple with the Ship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShip

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetShip(v interface{})`

SetShip sets Ship field to given value.

### HasShip

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasShip() bool`

HasShip returns a boolean if a field has been set.

### SetShipNil

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetShipNil(b bool)`

 SetShipNil sets the value for Ship to be an explicit nil

### UnsetShip
`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) UnsetShip()`

UnsetShip ensures that no value is present for Ship, not even an explicit nil
### GetGetRates

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetGetRates() interface{}`

GetGetRates returns the GetRates field if non-nil, zero value otherwise.

### GetGetRatesOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetGetRatesOk() (*interface{}, bool)`

GetGetRatesOk returns a tuple with the GetRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRates

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetGetRates(v interface{})`

SetGetRates sets GetRates field to given value.

### HasGetRates

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasGetRates() bool`

HasGetRates returns a boolean if a field has been set.

### SetGetRatesNil

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetGetRatesNil(b bool)`

 SetGetRatesNil sets the value for GetRates to be an explicit nil

### UnsetGetRates
`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) UnsetGetRates()`

UnsetGetRates ensures that no value is present for GetRates, not even an explicit nil
### GetSelectedRateId

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetSelectedRateId() interface{}`

GetSelectedRateId returns the SelectedRateId field if non-nil, zero value otherwise.

### GetSelectedRateIdOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetSelectedRateIdOk() (*interface{}, bool)`

GetSelectedRateIdOk returns a tuple with the SelectedRateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRateId

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetSelectedRateId(v interface{})`

SetSelectedRateId sets SelectedRateId field to given value.

### HasSelectedRateId

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasSelectedRateId() bool`

HasSelectedRateId returns a boolean if a field has been set.

### SetSelectedRateIdNil

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetSelectedRateIdNil(b bool)`

 SetSelectedRateIdNil sets the value for SelectedRateId to be an explicit nil

### UnsetSelectedRateId
`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) UnsetSelectedRateId()`

UnsetSelectedRateId ensures that no value is present for SelectedRateId, not even an explicit nil
### GetPurchase

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPurchase() interface{}`

GetPurchase returns the Purchase field if non-nil, zero value otherwise.

### GetPurchaseOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPurchaseOk() (*interface{}, bool)`

GetPurchaseOk returns a tuple with the Purchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchase

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetPurchase(v interface{})`

SetPurchase sets Purchase field to given value.

### HasPurchase

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasPurchase() bool`

HasPurchase returns a boolean if a field has been set.

### SetPurchaseNil

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetPurchaseNil(b bool)`

 SetPurchaseNil sets the value for Purchase to be an explicit nil

### UnsetPurchase
`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) UnsetPurchase()`

UnsetPurchase ensures that no value is present for Purchase, not even an explicit nil
### GetReference

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


