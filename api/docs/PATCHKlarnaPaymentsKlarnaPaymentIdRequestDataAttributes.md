# PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToken** | Pointer to **interface{}** | The token returned by a successful client authorization, mandatory to place the order. | [optional] 
**Update** | Pointer to **interface{}** | Send this attribute if you want to update the payment session with fresh order data. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes

`func NewPATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes() *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes`

NewPATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes instantiates a new PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributesWithDefaults

`func NewPATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributesWithDefaults() *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes`

NewPATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributesWithDefaults instantiates a new PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthToken

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) GetAuthToken() interface{}`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) GetAuthTokenOk() (*interface{}, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) SetAuthToken(v interface{})`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### SetAuthTokenNil

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) SetAuthTokenNil(b bool)`

 SetAuthTokenNil sets the value for AuthToken to be an explicit nil

### UnsetAuthToken
`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) UnsetAuthToken()`

UnsetAuthToken ensures that no value is present for AuthToken, not even an explicit nil
### GetUpdate

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) GetUpdate() interface{}`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) GetUpdateOk() (*interface{}, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) SetUpdate(v interface{})`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### SetUpdateNil

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) SetUpdateNil(b bool)`

 SetUpdateNil sets the value for Update to be an explicit nil

### UnsetUpdate
`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) UnsetUpdate()`

UnsetUpdate ensures that no value is present for Update, not even an explicit nil
### GetReference

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHKlarnaPaymentsKlarnaPaymentIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


