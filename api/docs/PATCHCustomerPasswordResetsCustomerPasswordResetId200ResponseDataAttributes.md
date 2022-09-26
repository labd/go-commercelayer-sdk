# PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerPassword** | Pointer to **string** | The customer new password. This will be accepted only if a valid &#39;_reset_password_token&#39; is sent with the request. | [optional] 
**ResetPasswordToken** | Pointer to **string** | Send the &#39;reset_password_token&#39; that you got on create when updating the customer password. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetCustomerPassword() string`

GetCustomerPassword returns the CustomerPassword field if non-nil, zero value otherwise.

### GetCustomerPasswordOk

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetCustomerPasswordOk() (*string, bool)`

GetCustomerPasswordOk returns a tuple with the CustomerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPassword

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetCustomerPassword(v string)`

SetCustomerPassword sets CustomerPassword field to given value.

### HasCustomerPassword

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasCustomerPassword() bool`

HasCustomerPassword returns a boolean if a field has been set.

### GetResetPasswordToken

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetResetPasswordToken() string`

GetResetPasswordToken returns the ResetPasswordToken field if non-nil, zero value otherwise.

### GetResetPasswordTokenOk

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetResetPasswordTokenOk() (*string, bool)`

GetResetPasswordTokenOk returns a tuple with the ResetPasswordToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPasswordToken

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetResetPasswordToken(v string)`

SetResetPasswordToken sets ResetPasswordToken field to given value.

### HasResetPasswordToken

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasResetPasswordToken() bool`

HasResetPasswordToken returns a boolean if a field has been set.

### GetReference

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


