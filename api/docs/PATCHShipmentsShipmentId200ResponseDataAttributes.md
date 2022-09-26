# PATCHShipmentsShipmentId200ResponseDataAttributes

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

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetOnHold() bool`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetOnHoldOk() (*bool, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetOnHold(v bool)`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### GetPicking

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPicking() bool`

GetPicking returns the Picking field if non-nil, zero value otherwise.

### GetPickingOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPickingOk() (*bool, bool)`

GetPickingOk returns a tuple with the Picking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicking

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetPicking(v bool)`

SetPicking sets Picking field to given value.

### HasPicking

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasPicking() bool`

HasPicking returns a boolean if a field has been set.

### GetPacking

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPacking() bool`

GetPacking returns the Packing field if non-nil, zero value otherwise.

### GetPackingOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPackingOk() (*bool, bool)`

GetPackingOk returns a tuple with the Packing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacking

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetPacking(v bool)`

SetPacking sets Packing field to given value.

### HasPacking

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasPacking() bool`

HasPacking returns a boolean if a field has been set.

### GetReadyToShip

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReadyToShip() bool`

GetReadyToShip returns the ReadyToShip field if non-nil, zero value otherwise.

### GetReadyToShipOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReadyToShipOk() (*bool, bool)`

GetReadyToShipOk returns a tuple with the ReadyToShip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyToShip

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetReadyToShip(v bool)`

SetReadyToShip sets ReadyToShip field to given value.

### HasReadyToShip

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasReadyToShip() bool`

HasReadyToShip returns a boolean if a field has been set.

### GetShip

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetShip() bool`

GetShip returns the Ship field if non-nil, zero value otherwise.

### GetShipOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetShipOk() (*bool, bool)`

GetShipOk returns a tuple with the Ship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShip

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetShip(v bool)`

SetShip sets Ship field to given value.

### HasShip

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasShip() bool`

HasShip returns a boolean if a field has been set.

### GetGetRates

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetGetRates() bool`

GetGetRates returns the GetRates field if non-nil, zero value otherwise.

### GetGetRatesOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetGetRatesOk() (*bool, bool)`

GetGetRatesOk returns a tuple with the GetRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRates

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetGetRates(v bool)`

SetGetRates sets GetRates field to given value.

### HasGetRates

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasGetRates() bool`

HasGetRates returns a boolean if a field has been set.

### GetSelectedRateId

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetSelectedRateId() string`

GetSelectedRateId returns the SelectedRateId field if non-nil, zero value otherwise.

### GetSelectedRateIdOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetSelectedRateIdOk() (*string, bool)`

GetSelectedRateIdOk returns a tuple with the SelectedRateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRateId

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetSelectedRateId(v string)`

SetSelectedRateId sets SelectedRateId field to given value.

### HasSelectedRateId

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasSelectedRateId() bool`

HasSelectedRateId returns a boolean if a field has been set.

### GetPurchase

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPurchase() bool`

GetPurchase returns the Purchase field if non-nil, zero value otherwise.

### GetPurchaseOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPurchaseOk() (*bool, bool)`

GetPurchaseOk returns a tuple with the Purchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchase

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetPurchase(v bool)`

SetPurchase sets Purchase field to given value.

### HasPurchase

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasPurchase() bool`

HasPurchase returns a boolean if a field has been set.

### GetReference

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


