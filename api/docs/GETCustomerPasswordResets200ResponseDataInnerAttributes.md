# GETCustomerPasswordResets200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerEmail** | Pointer to **string** | The email of the customer that requires a password reset | [optional] 
**ResetPasswordToken** | Pointer to **string** | Automatically generated on create. Send its value as the &#39;_reset_password_token&#39; argument when updating the customer password. | [optional] 
**ResetPasswordAt** | Pointer to **string** | Time at which the password was reset. | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetResetPasswordToken

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetResetPasswordToken() string`

GetResetPasswordToken returns the ResetPasswordToken field if non-nil, zero value otherwise.

### GetResetPasswordTokenOk

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetResetPasswordTokenOk() (*string, bool)`

GetResetPasswordTokenOk returns a tuple with the ResetPasswordToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPasswordToken

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetResetPasswordToken(v string)`

SetResetPasswordToken sets ResetPasswordToken field to given value.

### HasResetPasswordToken

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) HasResetPasswordToken() bool`

HasResetPasswordToken returns a boolean if a field has been set.

### GetResetPasswordAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetResetPasswordAt() string`

GetResetPasswordAt returns the ResetPasswordAt field if non-nil, zero value otherwise.

### GetResetPasswordAtOk

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetResetPasswordAtOk() (*string, bool)`

GetResetPasswordAtOk returns a tuple with the ResetPasswordAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPasswordAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetResetPasswordAt(v string)`

SetResetPasswordAt sets ResetPasswordAt field to given value.

### HasResetPasswordAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) HasResetPasswordAt() bool`

HasResetPasswordAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETCustomerPasswordResets200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


