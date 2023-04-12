# POSTRecurringOrderCopies201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaceTargetOrder** | Pointer to **interface{}** | Indicates if the target order must be placed upon copy. | [optional] 
**ReuseWallet** | Pointer to **interface{}** | Indicates if the payment source within the source order customer&#39;s wallet must be copied. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTRecurringOrderCopies201ResponseDataAttributes

`func NewPOSTRecurringOrderCopies201ResponseDataAttributes() *POSTRecurringOrderCopies201ResponseDataAttributes`

NewPOSTRecurringOrderCopies201ResponseDataAttributes instantiates a new POSTRecurringOrderCopies201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTRecurringOrderCopies201ResponseDataAttributesWithDefaults

`func NewPOSTRecurringOrderCopies201ResponseDataAttributesWithDefaults() *POSTRecurringOrderCopies201ResponseDataAttributes`

NewPOSTRecurringOrderCopies201ResponseDataAttributesWithDefaults instantiates a new POSTRecurringOrderCopies201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaceTargetOrder

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) GetPlaceTargetOrder() interface{}`

GetPlaceTargetOrder returns the PlaceTargetOrder field if non-nil, zero value otherwise.

### GetPlaceTargetOrderOk

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) GetPlaceTargetOrderOk() (*interface{}, bool)`

GetPlaceTargetOrderOk returns a tuple with the PlaceTargetOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceTargetOrder

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) SetPlaceTargetOrder(v interface{})`

SetPlaceTargetOrder sets PlaceTargetOrder field to given value.

### HasPlaceTargetOrder

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) HasPlaceTargetOrder() bool`

HasPlaceTargetOrder returns a boolean if a field has been set.

### SetPlaceTargetOrderNil

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) SetPlaceTargetOrderNil(b bool)`

 SetPlaceTargetOrderNil sets the value for PlaceTargetOrder to be an explicit nil

### UnsetPlaceTargetOrder
`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) UnsetPlaceTargetOrder()`

UnsetPlaceTargetOrder ensures that no value is present for PlaceTargetOrder, not even an explicit nil
### GetReuseWallet

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) GetReuseWallet() interface{}`

GetReuseWallet returns the ReuseWallet field if non-nil, zero value otherwise.

### GetReuseWalletOk

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) GetReuseWalletOk() (*interface{}, bool)`

GetReuseWalletOk returns a tuple with the ReuseWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseWallet

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) SetReuseWallet(v interface{})`

SetReuseWallet sets ReuseWallet field to given value.

### HasReuseWallet

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) HasReuseWallet() bool`

HasReuseWallet returns a boolean if a field has been set.

### SetReuseWalletNil

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) SetReuseWalletNil(b bool)`

 SetReuseWalletNil sets the value for ReuseWallet to be an explicit nil

### UnsetReuseWallet
`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) UnsetReuseWallet()`

UnsetReuseWallet ensures that no value is present for ReuseWallet, not even an explicit nil
### GetReference

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTRecurringOrderCopies201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


