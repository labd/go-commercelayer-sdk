# PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerPassword** | Pointer to **interface{}** | The customer new password. This will be accepted only if a valid &#39;_reset_password_token&#39; is sent with the request. | [optional] 
**ResetPasswordToken** | Pointer to **interface{}** | Send the &#39;reset_password_token&#39; that you got on create when updating the customer password. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes

`func NewPATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes() *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes`

NewPATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes instantiates a new PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributesWithDefaults

`func NewPATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributesWithDefaults() *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes`

NewPATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributesWithDefaults instantiates a new PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerPassword

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) GetCustomerPassword() interface{}`

GetCustomerPassword returns the CustomerPassword field if non-nil, zero value otherwise.

### GetCustomerPasswordOk

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) GetCustomerPasswordOk() (*interface{}, bool)`

GetCustomerPasswordOk returns a tuple with the CustomerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPassword

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) SetCustomerPassword(v interface{})`

SetCustomerPassword sets CustomerPassword field to given value.

### HasCustomerPassword

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) HasCustomerPassword() bool`

HasCustomerPassword returns a boolean if a field has been set.

### SetCustomerPasswordNil

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) SetCustomerPasswordNil(b bool)`

 SetCustomerPasswordNil sets the value for CustomerPassword to be an explicit nil

### UnsetCustomerPassword
`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) UnsetCustomerPassword()`

UnsetCustomerPassword ensures that no value is present for CustomerPassword, not even an explicit nil
### GetResetPasswordToken

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) GetResetPasswordToken() interface{}`

GetResetPasswordToken returns the ResetPasswordToken field if non-nil, zero value otherwise.

### GetResetPasswordTokenOk

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) GetResetPasswordTokenOk() (*interface{}, bool)`

GetResetPasswordTokenOk returns a tuple with the ResetPasswordToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPasswordToken

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) SetResetPasswordToken(v interface{})`

SetResetPasswordToken sets ResetPasswordToken field to given value.

### HasResetPasswordToken

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) HasResetPasswordToken() bool`

HasResetPasswordToken returns a boolean if a field has been set.

### SetResetPasswordTokenNil

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) SetResetPasswordTokenNil(b bool)`

 SetResetPasswordTokenNil sets the value for ResetPasswordToken to be an explicit nil

### UnsetResetPasswordToken
`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) UnsetResetPasswordToken()`

UnsetResetPasswordToken ensures that no value is present for ResetPasswordToken, not even an explicit nil
### GetReference

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


