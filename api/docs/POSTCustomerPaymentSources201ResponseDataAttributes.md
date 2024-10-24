# POSTCustomerPaymentSources201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerToken** | Pointer to **interface{}** | Returns the customer gateway token stored in the gateway. | [optional] 
**PaymentSourceToken** | Pointer to **interface{}** | Returns the payment source token stored in the gateway. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTCustomerPaymentSources201ResponseDataAttributes

`func NewPOSTCustomerPaymentSources201ResponseDataAttributes() *POSTCustomerPaymentSources201ResponseDataAttributes`

NewPOSTCustomerPaymentSources201ResponseDataAttributes instantiates a new POSTCustomerPaymentSources201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTCustomerPaymentSources201ResponseDataAttributesWithDefaults

`func NewPOSTCustomerPaymentSources201ResponseDataAttributesWithDefaults() *POSTCustomerPaymentSources201ResponseDataAttributes`

NewPOSTCustomerPaymentSources201ResponseDataAttributesWithDefaults instantiates a new POSTCustomerPaymentSources201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerToken

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) GetCustomerToken() interface{}`

GetCustomerToken returns the CustomerToken field if non-nil, zero value otherwise.

### GetCustomerTokenOk

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) GetCustomerTokenOk() (*interface{}, bool)`

GetCustomerTokenOk returns a tuple with the CustomerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerToken

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) SetCustomerToken(v interface{})`

SetCustomerToken sets CustomerToken field to given value.

### HasCustomerToken

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) HasCustomerToken() bool`

HasCustomerToken returns a boolean if a field has been set.

### SetCustomerTokenNil

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) SetCustomerTokenNil(b bool)`

 SetCustomerTokenNil sets the value for CustomerToken to be an explicit nil

### UnsetCustomerToken
`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) UnsetCustomerToken()`

UnsetCustomerToken ensures that no value is present for CustomerToken, not even an explicit nil
### GetPaymentSourceToken

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) GetPaymentSourceToken() interface{}`

GetPaymentSourceToken returns the PaymentSourceToken field if non-nil, zero value otherwise.

### GetPaymentSourceTokenOk

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) GetPaymentSourceTokenOk() (*interface{}, bool)`

GetPaymentSourceTokenOk returns a tuple with the PaymentSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceToken

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) SetPaymentSourceToken(v interface{})`

SetPaymentSourceToken sets PaymentSourceToken field to given value.

### HasPaymentSourceToken

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) HasPaymentSourceToken() bool`

HasPaymentSourceToken returns a boolean if a field has been set.

### SetPaymentSourceTokenNil

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) SetPaymentSourceTokenNil(b bool)`

 SetPaymentSourceTokenNil sets the value for PaymentSourceToken to be an explicit nil

### UnsetPaymentSourceToken
`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) UnsetPaymentSourceToken()`

UnsetPaymentSourceToken ensures that no value is present for PaymentSourceToken, not even an explicit nil
### GetReference

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTCustomerPaymentSources201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


