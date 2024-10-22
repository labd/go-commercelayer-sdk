# POSTCustomerPasswordResets201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerEmail** | **interface{}** | The email of the customer that requires a password reset. | 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTCustomerPasswordResets201ResponseDataAttributes

`func NewPOSTCustomerPasswordResets201ResponseDataAttributes(customerEmail interface{}, ) *POSTCustomerPasswordResets201ResponseDataAttributes`

NewPOSTCustomerPasswordResets201ResponseDataAttributes instantiates a new POSTCustomerPasswordResets201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTCustomerPasswordResets201ResponseDataAttributesWithDefaults

`func NewPOSTCustomerPasswordResets201ResponseDataAttributesWithDefaults() *POSTCustomerPasswordResets201ResponseDataAttributes`

NewPOSTCustomerPasswordResets201ResponseDataAttributesWithDefaults instantiates a new POSTCustomerPasswordResets201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerEmail

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) GetCustomerEmail() interface{}`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) GetCustomerEmailOk() (*interface{}, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) SetCustomerEmail(v interface{})`

SetCustomerEmail sets CustomerEmail field to given value.


### SetCustomerEmailNil

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) SetCustomerEmailNil(b bool)`

 SetCustomerEmailNil sets the value for CustomerEmail to be an explicit nil

### UnsetCustomerEmail
`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) UnsetCustomerEmail()`

UnsetCustomerEmail ensures that no value is present for CustomerEmail, not even an explicit nil
### GetReference

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTCustomerPasswordResets201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


