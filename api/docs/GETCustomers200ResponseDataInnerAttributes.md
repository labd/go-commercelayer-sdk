# GETCustomers200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **interface{}** | The customer&#39;s email address | [optional] 
**Status** | Pointer to **interface{}** | The customer&#39;s status, one of &#39;prospect&#39;, &#39;acquired&#39;, or &#39;repeat&#39;. | [optional] 
**HasPassword** | Pointer to **interface{}** | Indicates if the customer has a password. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETCustomers200ResponseDataInnerAttributes

`func NewGETCustomers200ResponseDataInnerAttributes() *GETCustomers200ResponseDataInnerAttributes`

NewGETCustomers200ResponseDataInnerAttributes instantiates a new GETCustomers200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETCustomers200ResponseDataInnerAttributesWithDefaults

`func NewGETCustomers200ResponseDataInnerAttributesWithDefaults() *GETCustomers200ResponseDataInnerAttributes`

NewGETCustomers200ResponseDataInnerAttributesWithDefaults instantiates a new GETCustomers200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *GETCustomers200ResponseDataInnerAttributes) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GETCustomers200ResponseDataInnerAttributes) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GETCustomers200ResponseDataInnerAttributes) SetEmail(v interface{})`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GETCustomers200ResponseDataInnerAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *GETCustomers200ResponseDataInnerAttributes) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *GETCustomers200ResponseDataInnerAttributes) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetStatus

`func (o *GETCustomers200ResponseDataInnerAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETCustomers200ResponseDataInnerAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETCustomers200ResponseDataInnerAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETCustomers200ResponseDataInnerAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETCustomers200ResponseDataInnerAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETCustomers200ResponseDataInnerAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetHasPassword

`func (o *GETCustomers200ResponseDataInnerAttributes) GetHasPassword() interface{}`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *GETCustomers200ResponseDataInnerAttributes) GetHasPasswordOk() (*interface{}, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *GETCustomers200ResponseDataInnerAttributes) SetHasPassword(v interface{})`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *GETCustomers200ResponseDataInnerAttributes) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### SetHasPasswordNil

`func (o *GETCustomers200ResponseDataInnerAttributes) SetHasPasswordNil(b bool)`

 SetHasPasswordNil sets the value for HasPassword to be an explicit nil

### UnsetHasPassword
`func (o *GETCustomers200ResponseDataInnerAttributes) UnsetHasPassword()`

UnsetHasPassword ensures that no value is present for HasPassword, not even an explicit nil
### GetCreatedAt

`func (o *GETCustomers200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETCustomers200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETCustomers200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETCustomers200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETCustomers200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETCustomers200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETCustomers200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETCustomers200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETCustomers200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETCustomers200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETCustomers200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETCustomers200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETCustomers200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETCustomers200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETCustomers200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETCustomers200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETCustomers200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETCustomers200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETCustomers200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETCustomers200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETCustomers200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETCustomers200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETCustomers200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETCustomers200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETCustomers200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETCustomers200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETCustomers200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETCustomers200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETCustomers200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETCustomers200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


