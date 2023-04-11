# GETCustomerPasswordResets200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerEmail** | Pointer to **interface{}** | The email of the customer that requires a password reset | [optional] 
**ResetPasswordToken** | Pointer to **interface{}** | Automatically generated on create. Send its value as the &#39;_reset_password_token&#39; argument when updating the customer password. | [optional] 
**ResetPasswordAt** | Pointer to **interface{}** | Time at which the password was reset. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETCustomerPasswordResets200ResponseDataInnerAttributes

`func NewGETCustomerPasswordResets200ResponseDataInnerAttributes() *GETCustomerPasswordResets200ResponseDataInnerAttributes`

NewGETCustomerPasswordResets200ResponseDataInnerAttributes instantiates a new GETCustomerPasswordResets200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETCustomerPasswordResets200ResponseDataInnerAttributesWithDefaults

`func NewGETCustomerPasswordResets200ResponseDataInnerAttributesWithDefaults() *GETCustomerPasswordResets200ResponseDataInnerAttributes`

NewGETCustomerPasswordResets200ResponseDataInnerAttributesWithDefaults instantiates a new GETCustomerPasswordResets200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerEmail

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetCustomerEmail() interface{}`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetCustomerEmailOk() (*interface{}, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetCustomerEmail(v interface{})`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### SetCustomerEmailNil

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetCustomerEmailNil(b bool)`

 SetCustomerEmailNil sets the value for CustomerEmail to be an explicit nil

### UnsetCustomerEmail
`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) UnsetCustomerEmail()`

UnsetCustomerEmail ensures that no value is present for CustomerEmail, not even an explicit nil
### GetResetPasswordToken

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetResetPasswordToken() interface{}`

GetResetPasswordToken returns the ResetPasswordToken field if non-nil, zero value otherwise.

### GetResetPasswordTokenOk

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetResetPasswordTokenOk() (*interface{}, bool)`

GetResetPasswordTokenOk returns a tuple with the ResetPasswordToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPasswordToken

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetResetPasswordToken(v interface{})`

SetResetPasswordToken sets ResetPasswordToken field to given value.

### HasResetPasswordToken

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) HasResetPasswordToken() bool`

HasResetPasswordToken returns a boolean if a field has been set.

### SetResetPasswordTokenNil

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetResetPasswordTokenNil(b bool)`

 SetResetPasswordTokenNil sets the value for ResetPasswordToken to be an explicit nil

### UnsetResetPasswordToken
`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) UnsetResetPasswordToken()`

UnsetResetPasswordToken ensures that no value is present for ResetPasswordToken, not even an explicit nil
### GetResetPasswordAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetResetPasswordAt() interface{}`

GetResetPasswordAt returns the ResetPasswordAt field if non-nil, zero value otherwise.

### GetResetPasswordAtOk

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetResetPasswordAtOk() (*interface{}, bool)`

GetResetPasswordAtOk returns a tuple with the ResetPasswordAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPasswordAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetResetPasswordAt(v interface{})`

SetResetPasswordAt sets ResetPasswordAt field to given value.

### HasResetPasswordAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) HasResetPasswordAt() bool`

HasResetPasswordAt returns a boolean if a field has been set.

### SetResetPasswordAtNil

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetResetPasswordAtNil(b bool)`

 SetResetPasswordAtNil sets the value for ResetPasswordAt to be an explicit nil

### UnsetResetPasswordAt
`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) UnsetResetPasswordAt()`

UnsetResetPasswordAt ensures that no value is present for ResetPasswordAt, not even an explicit nil
### GetCreatedAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


