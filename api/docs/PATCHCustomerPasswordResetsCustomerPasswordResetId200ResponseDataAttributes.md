# PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerPassword** | Pointer to **interface{}** | The customer new password. This will be accepted only if a valid &#39;_reset_password_token&#39; is sent with the request. | [optional] 
**ResetPasswordToken** | Pointer to **interface{}** | Send the &#39;reset_password_token&#39; that you got on create when updating the customer password. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes

`func NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes() *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes`

NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes instantiates a new PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributesWithDefaults

`func NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributesWithDefaults() *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes`

NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributesWithDefaults instantiates a new PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerPassword

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetCustomerPassword() interface{}`

GetCustomerPassword returns the CustomerPassword field if non-nil, zero value otherwise.

### GetCustomerPasswordOk

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetCustomerPasswordOk() (*interface{}, bool)`

GetCustomerPasswordOk returns a tuple with the CustomerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPassword

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetCustomerPassword(v interface{})`

SetCustomerPassword sets CustomerPassword field to given value.

### HasCustomerPassword

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasCustomerPassword() bool`

HasCustomerPassword returns a boolean if a field has been set.

### SetCustomerPasswordNil

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetCustomerPasswordNil(b bool)`

 SetCustomerPasswordNil sets the value for CustomerPassword to be an explicit nil

### UnsetCustomerPassword
`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) UnsetCustomerPassword()`

UnsetCustomerPassword ensures that no value is present for CustomerPassword, not even an explicit nil
### GetResetPasswordToken

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetResetPasswordToken() interface{}`

GetResetPasswordToken returns the ResetPasswordToken field if non-nil, zero value otherwise.

### GetResetPasswordTokenOk

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetResetPasswordTokenOk() (*interface{}, bool)`

GetResetPasswordTokenOk returns a tuple with the ResetPasswordToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPasswordToken

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetResetPasswordToken(v interface{})`

SetResetPasswordToken sets ResetPasswordToken field to given value.

### HasResetPasswordToken

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasResetPasswordToken() bool`

HasResetPasswordToken returns a boolean if a field has been set.

### SetResetPasswordTokenNil

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetResetPasswordTokenNil(b bool)`

 SetResetPasswordTokenNil sets the value for ResetPasswordToken to be an explicit nil

### UnsetResetPasswordToken
`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) UnsetResetPasswordToken()`

UnsetResetPasswordToken ensures that no value is present for ResetPasswordToken, not even an explicit nil
### GetReference

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


