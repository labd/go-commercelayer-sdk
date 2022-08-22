# PATCHStockTransfersStockTransferId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **string** | The code of the associated SKU. | [optional] 
**Upcoming** | Pointer to **bool** | Send this attribute if you want to mark this stock transfer as upcoming. | [optional] 
**Picking** | Pointer to **bool** | Send this attribute if you want to start picking this stock transfer. | [optional] 
**InTransit** | Pointer to **bool** | Send this attribute if you want to mark this stock transfer as in transit. | [optional] 
**Complete** | Pointer to **bool** | Send this attribute if you want to complete this stock transfer. | [optional] 
**Cancel** | Pointer to **bool** | Send this attribute if you want to cancel this stock transfer. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHStockTransfersStockTransferId200ResponseDataAttributes

`func NewPATCHStockTransfersStockTransferId200ResponseDataAttributes() *PATCHStockTransfersStockTransferId200ResponseDataAttributes`

NewPATCHStockTransfersStockTransferId200ResponseDataAttributes instantiates a new PATCHStockTransfersStockTransferId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHStockTransfersStockTransferId200ResponseDataAttributesWithDefaults

`func NewPATCHStockTransfersStockTransferId200ResponseDataAttributesWithDefaults() *PATCHStockTransfersStockTransferId200ResponseDataAttributes`

NewPATCHStockTransfersStockTransferId200ResponseDataAttributesWithDefaults instantiates a new PATCHStockTransfersStockTransferId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetSkuCode() string`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetSkuCodeOk() (*string, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetSkuCode(v string)`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### GetUpcoming

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetUpcoming() bool`

GetUpcoming returns the Upcoming field if non-nil, zero value otherwise.

### GetUpcomingOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetUpcomingOk() (*bool, bool)`

GetUpcomingOk returns a tuple with the Upcoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcoming

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetUpcoming(v bool)`

SetUpcoming sets Upcoming field to given value.

### HasUpcoming

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasUpcoming() bool`

HasUpcoming returns a boolean if a field has been set.

### GetPicking

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetPicking() bool`

GetPicking returns the Picking field if non-nil, zero value otherwise.

### GetPickingOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetPickingOk() (*bool, bool)`

GetPickingOk returns a tuple with the Picking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicking

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetPicking(v bool)`

SetPicking sets Picking field to given value.

### HasPicking

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasPicking() bool`

HasPicking returns a boolean if a field has been set.

### GetInTransit

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetInTransit() bool`

GetInTransit returns the InTransit field if non-nil, zero value otherwise.

### GetInTransitOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetInTransitOk() (*bool, bool)`

GetInTransitOk returns a tuple with the InTransit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTransit

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetInTransit(v bool)`

SetInTransit sets InTransit field to given value.

### HasInTransit

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasInTransit() bool`

HasInTransit returns a boolean if a field has been set.

### GetComplete

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetCancel

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetCancel() bool`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetCancelOk() (*bool, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetCancel(v bool)`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### GetReference

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


