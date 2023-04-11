# PATCHStockTransfersStockTransferIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **interface{}** | The code of the associated SKU. | [optional] 
**Upcoming** | Pointer to **interface{}** | Send this attribute if you want to mark this stock transfer as upcoming. | [optional] 
**Picking** | Pointer to **interface{}** | Send this attribute if you want to start picking this stock transfer. | [optional] 
**InTransit** | Pointer to **interface{}** | Send this attribute if you want to mark this stock transfer as in transit. | [optional] 
**Complete** | Pointer to **interface{}** | Send this attribute if you want to complete this stock transfer. | [optional] 
**Cancel** | Pointer to **interface{}** | Send this attribute if you want to cancel this stock transfer. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHStockTransfersStockTransferIdRequestDataAttributes

`func NewPATCHStockTransfersStockTransferIdRequestDataAttributes() *PATCHStockTransfersStockTransferIdRequestDataAttributes`

NewPATCHStockTransfersStockTransferIdRequestDataAttributes instantiates a new PATCHStockTransfersStockTransferIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHStockTransfersStockTransferIdRequestDataAttributesWithDefaults

`func NewPATCHStockTransfersStockTransferIdRequestDataAttributesWithDefaults() *PATCHStockTransfersStockTransferIdRequestDataAttributes`

NewPATCHStockTransfersStockTransferIdRequestDataAttributesWithDefaults instantiates a new PATCHStockTransfersStockTransferIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetSkuCode() interface{}`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetSkuCodeOk() (*interface{}, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetSkuCode(v interface{})`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### SetSkuCodeNil

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetSkuCodeNil(b bool)`

 SetSkuCodeNil sets the value for SkuCode to be an explicit nil

### UnsetSkuCode
`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) UnsetSkuCode()`

UnsetSkuCode ensures that no value is present for SkuCode, not even an explicit nil
### GetUpcoming

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetUpcoming() interface{}`

GetUpcoming returns the Upcoming field if non-nil, zero value otherwise.

### GetUpcomingOk

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetUpcomingOk() (*interface{}, bool)`

GetUpcomingOk returns a tuple with the Upcoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcoming

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetUpcoming(v interface{})`

SetUpcoming sets Upcoming field to given value.

### HasUpcoming

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) HasUpcoming() bool`

HasUpcoming returns a boolean if a field has been set.

### SetUpcomingNil

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetUpcomingNil(b bool)`

 SetUpcomingNil sets the value for Upcoming to be an explicit nil

### UnsetUpcoming
`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) UnsetUpcoming()`

UnsetUpcoming ensures that no value is present for Upcoming, not even an explicit nil
### GetPicking

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetPicking() interface{}`

GetPicking returns the Picking field if non-nil, zero value otherwise.

### GetPickingOk

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetPickingOk() (*interface{}, bool)`

GetPickingOk returns a tuple with the Picking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicking

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetPicking(v interface{})`

SetPicking sets Picking field to given value.

### HasPicking

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) HasPicking() bool`

HasPicking returns a boolean if a field has been set.

### SetPickingNil

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetPickingNil(b bool)`

 SetPickingNil sets the value for Picking to be an explicit nil

### UnsetPicking
`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) UnsetPicking()`

UnsetPicking ensures that no value is present for Picking, not even an explicit nil
### GetInTransit

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetInTransit() interface{}`

GetInTransit returns the InTransit field if non-nil, zero value otherwise.

### GetInTransitOk

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetInTransitOk() (*interface{}, bool)`

GetInTransitOk returns a tuple with the InTransit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTransit

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetInTransit(v interface{})`

SetInTransit sets InTransit field to given value.

### HasInTransit

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) HasInTransit() bool`

HasInTransit returns a boolean if a field has been set.

### SetInTransitNil

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetInTransitNil(b bool)`

 SetInTransitNil sets the value for InTransit to be an explicit nil

### UnsetInTransit
`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) UnsetInTransit()`

UnsetInTransit ensures that no value is present for InTransit, not even an explicit nil
### GetComplete

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetComplete() interface{}`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetCompleteOk() (*interface{}, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetComplete(v interface{})`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### SetCompleteNil

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetCompleteNil(b bool)`

 SetCompleteNil sets the value for Complete to be an explicit nil

### UnsetComplete
`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) UnsetComplete()`

UnsetComplete ensures that no value is present for Complete, not even an explicit nil
### GetCancel

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetCancel() interface{}`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetCancelOk() (*interface{}, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetCancel(v interface{})`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### SetCancelNil

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetCancelNil(b bool)`

 SetCancelNil sets the value for Cancel to be an explicit nil

### UnsetCancel
`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) UnsetCancel()`

UnsetCancel ensures that no value is present for Cancel, not even an explicit nil
### GetReference

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHStockTransfersStockTransferIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


