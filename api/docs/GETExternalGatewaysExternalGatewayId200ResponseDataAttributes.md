# GETExternalGatewaysExternalGatewayId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The payment gateway&#39;s internal name. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**SharedSecret** | Pointer to **interface{}** | The shared secret used to sign the external request payload. | [optional] 
**AuthorizeUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to authorize payments. | [optional] 
**CaptureUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to capture payments. | [optional] 
**VoidUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to void payments. | [optional] 
**RefundUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to refund payments. | [optional] 
**TokenUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to create a customer payment token. | [optional] 

## Methods

### NewGETExternalGatewaysExternalGatewayId200ResponseDataAttributes

`func NewGETExternalGatewaysExternalGatewayId200ResponseDataAttributes() *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes`

NewGETExternalGatewaysExternalGatewayId200ResponseDataAttributes instantiates a new GETExternalGatewaysExternalGatewayId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETExternalGatewaysExternalGatewayId200ResponseDataAttributesWithDefaults

`func NewGETExternalGatewaysExternalGatewayId200ResponseDataAttributesWithDefaults() *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes`

NewGETExternalGatewaysExternalGatewayId200ResponseDataAttributesWithDefaults instantiates a new GETExternalGatewaysExternalGatewayId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreatedAt

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetSharedSecret

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetSharedSecret() interface{}`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetSharedSecretOk() (*interface{}, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetSharedSecret(v interface{})`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### SetSharedSecretNil

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetSharedSecretNil(b bool)`

 SetSharedSecretNil sets the value for SharedSecret to be an explicit nil

### UnsetSharedSecret
`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetSharedSecret()`

UnsetSharedSecret ensures that no value is present for SharedSecret, not even an explicit nil
### GetAuthorizeUrl

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetAuthorizeUrl() interface{}`

GetAuthorizeUrl returns the AuthorizeUrl field if non-nil, zero value otherwise.

### GetAuthorizeUrlOk

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetAuthorizeUrlOk() (*interface{}, bool)`

GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeUrl

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetAuthorizeUrl(v interface{})`

SetAuthorizeUrl sets AuthorizeUrl field to given value.

### HasAuthorizeUrl

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasAuthorizeUrl() bool`

HasAuthorizeUrl returns a boolean if a field has been set.

### SetAuthorizeUrlNil

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetAuthorizeUrlNil(b bool)`

 SetAuthorizeUrlNil sets the value for AuthorizeUrl to be an explicit nil

### UnsetAuthorizeUrl
`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetAuthorizeUrl()`

UnsetAuthorizeUrl ensures that no value is present for AuthorizeUrl, not even an explicit nil
### GetCaptureUrl

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetCaptureUrl() interface{}`

GetCaptureUrl returns the CaptureUrl field if non-nil, zero value otherwise.

### GetCaptureUrlOk

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetCaptureUrlOk() (*interface{}, bool)`

GetCaptureUrlOk returns a tuple with the CaptureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureUrl

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetCaptureUrl(v interface{})`

SetCaptureUrl sets CaptureUrl field to given value.

### HasCaptureUrl

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasCaptureUrl() bool`

HasCaptureUrl returns a boolean if a field has been set.

### SetCaptureUrlNil

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetCaptureUrlNil(b bool)`

 SetCaptureUrlNil sets the value for CaptureUrl to be an explicit nil

### UnsetCaptureUrl
`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetCaptureUrl()`

UnsetCaptureUrl ensures that no value is present for CaptureUrl, not even an explicit nil
### GetVoidUrl

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetVoidUrl() interface{}`

GetVoidUrl returns the VoidUrl field if non-nil, zero value otherwise.

### GetVoidUrlOk

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetVoidUrlOk() (*interface{}, bool)`

GetVoidUrlOk returns a tuple with the VoidUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidUrl

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetVoidUrl(v interface{})`

SetVoidUrl sets VoidUrl field to given value.

### HasVoidUrl

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasVoidUrl() bool`

HasVoidUrl returns a boolean if a field has been set.

### SetVoidUrlNil

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetVoidUrlNil(b bool)`

 SetVoidUrlNil sets the value for VoidUrl to be an explicit nil

### UnsetVoidUrl
`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetVoidUrl()`

UnsetVoidUrl ensures that no value is present for VoidUrl, not even an explicit nil
### GetRefundUrl

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetRefundUrl() interface{}`

GetRefundUrl returns the RefundUrl field if non-nil, zero value otherwise.

### GetRefundUrlOk

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetRefundUrlOk() (*interface{}, bool)`

GetRefundUrlOk returns a tuple with the RefundUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundUrl

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetRefundUrl(v interface{})`

SetRefundUrl sets RefundUrl field to given value.

### HasRefundUrl

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasRefundUrl() bool`

HasRefundUrl returns a boolean if a field has been set.

### SetRefundUrlNil

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetRefundUrlNil(b bool)`

 SetRefundUrlNil sets the value for RefundUrl to be an explicit nil

### UnsetRefundUrl
`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetRefundUrl()`

UnsetRefundUrl ensures that no value is present for RefundUrl, not even an explicit nil
### GetTokenUrl

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetTokenUrl() interface{}`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetTokenUrlOk() (*interface{}, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetTokenUrl(v interface{})`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### SetTokenUrlNil

`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetTokenUrlNil(b bool)`

 SetTokenUrlNil sets the value for TokenUrl to be an explicit nil

### UnsetTokenUrl
`func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetTokenUrl()`

UnsetTokenUrl ensures that no value is present for TokenUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


