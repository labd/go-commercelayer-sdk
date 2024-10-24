# PATCHRefundsRefundId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Succeeded** | Pointer to **interface{}** | Indicates if the transaction is successful. | [optional] 
**Forward** | Pointer to **interface{}** | Send this attribute if you want to forward a stuck transaction to succeeded and update associated order states accordingly. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHRefundsRefundId200ResponseDataAttributes

`func NewPATCHRefundsRefundId200ResponseDataAttributes() *PATCHRefundsRefundId200ResponseDataAttributes`

NewPATCHRefundsRefundId200ResponseDataAttributes instantiates a new PATCHRefundsRefundId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHRefundsRefundId200ResponseDataAttributesWithDefaults

`func NewPATCHRefundsRefundId200ResponseDataAttributesWithDefaults() *PATCHRefundsRefundId200ResponseDataAttributes`

NewPATCHRefundsRefundId200ResponseDataAttributesWithDefaults instantiates a new PATCHRefundsRefundId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSucceeded

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) GetSucceeded() interface{}`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) GetSucceededOk() (*interface{}, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) SetSucceeded(v interface{})`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### SetSucceededNil

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) SetSucceededNil(b bool)`

 SetSucceededNil sets the value for Succeeded to be an explicit nil

### UnsetSucceeded
`func (o *PATCHRefundsRefundId200ResponseDataAttributes) UnsetSucceeded()`

UnsetSucceeded ensures that no value is present for Succeeded, not even an explicit nil
### GetForward

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) GetForward() interface{}`

GetForward returns the Forward field if non-nil, zero value otherwise.

### GetForwardOk

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) GetForwardOk() (*interface{}, bool)`

GetForwardOk returns a tuple with the Forward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForward

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) SetForward(v interface{})`

SetForward sets Forward field to given value.

### HasForward

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) HasForward() bool`

HasForward returns a boolean if a field has been set.

### SetForwardNil

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) SetForwardNil(b bool)`

 SetForwardNil sets the value for Forward to be an explicit nil

### UnsetForward
`func (o *PATCHRefundsRefundId200ResponseDataAttributes) UnsetForward()`

UnsetForward ensures that no value is present for Forward, not even an explicit nil
### GetReference

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHRefundsRefundId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHRefundsRefundId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHRefundsRefundId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHRefundsRefundId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


