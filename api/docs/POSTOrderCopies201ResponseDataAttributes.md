# POSTOrderCopies201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaceTargetOrder** | Pointer to **bool** | Indicates if the target order must be placed upon copy. | [optional] 
**CancelSourceOrder** | Pointer to **bool** | Indicates if the source order must be cancelled upon copy. | [optional] 
**ReuseWallet** | Pointer to **bool** | Indicates if the payment source within the source order customer&#39;s wallet must be copied. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTOrderCopies201ResponseDataAttributes

`func NewPOSTOrderCopies201ResponseDataAttributes() *POSTOrderCopies201ResponseDataAttributes`

NewPOSTOrderCopies201ResponseDataAttributes instantiates a new POSTOrderCopies201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrderCopies201ResponseDataAttributesWithDefaults

`func NewPOSTOrderCopies201ResponseDataAttributesWithDefaults() *POSTOrderCopies201ResponseDataAttributes`

NewPOSTOrderCopies201ResponseDataAttributesWithDefaults instantiates a new POSTOrderCopies201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaceTargetOrder

`func (o *POSTOrderCopies201ResponseDataAttributes) GetPlaceTargetOrder() bool`

GetPlaceTargetOrder returns the PlaceTargetOrder field if non-nil, zero value otherwise.

### GetPlaceTargetOrderOk

`func (o *POSTOrderCopies201ResponseDataAttributes) GetPlaceTargetOrderOk() (*bool, bool)`

GetPlaceTargetOrderOk returns a tuple with the PlaceTargetOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceTargetOrder

`func (o *POSTOrderCopies201ResponseDataAttributes) SetPlaceTargetOrder(v bool)`

SetPlaceTargetOrder sets PlaceTargetOrder field to given value.

### HasPlaceTargetOrder

`func (o *POSTOrderCopies201ResponseDataAttributes) HasPlaceTargetOrder() bool`

HasPlaceTargetOrder returns a boolean if a field has been set.

### GetCancelSourceOrder

`func (o *POSTOrderCopies201ResponseDataAttributes) GetCancelSourceOrder() bool`

GetCancelSourceOrder returns the CancelSourceOrder field if non-nil, zero value otherwise.

### GetCancelSourceOrderOk

`func (o *POSTOrderCopies201ResponseDataAttributes) GetCancelSourceOrderOk() (*bool, bool)`

GetCancelSourceOrderOk returns a tuple with the CancelSourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelSourceOrder

`func (o *POSTOrderCopies201ResponseDataAttributes) SetCancelSourceOrder(v bool)`

SetCancelSourceOrder sets CancelSourceOrder field to given value.

### HasCancelSourceOrder

`func (o *POSTOrderCopies201ResponseDataAttributes) HasCancelSourceOrder() bool`

HasCancelSourceOrder returns a boolean if a field has been set.

### GetReuseWallet

`func (o *POSTOrderCopies201ResponseDataAttributes) GetReuseWallet() bool`

GetReuseWallet returns the ReuseWallet field if non-nil, zero value otherwise.

### GetReuseWalletOk

`func (o *POSTOrderCopies201ResponseDataAttributes) GetReuseWalletOk() (*bool, bool)`

GetReuseWalletOk returns a tuple with the ReuseWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseWallet

`func (o *POSTOrderCopies201ResponseDataAttributes) SetReuseWallet(v bool)`

SetReuseWallet sets ReuseWallet field to given value.

### HasReuseWallet

`func (o *POSTOrderCopies201ResponseDataAttributes) HasReuseWallet() bool`

HasReuseWallet returns a boolean if a field has been set.

### GetReference

`func (o *POSTOrderCopies201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTOrderCopies201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTOrderCopies201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTOrderCopies201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTOrderCopies201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTOrderCopies201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTOrderCopies201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTOrderCopies201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTOrderCopies201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTOrderCopies201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTOrderCopies201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTOrderCopies201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


