# GETExternalGateways200ResponseDataInnerAttributes

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

### NewGETExternalGateways200ResponseDataInnerAttributes

`func NewGETExternalGateways200ResponseDataInnerAttributes() *GETExternalGateways200ResponseDataInnerAttributes`

NewGETExternalGateways200ResponseDataInnerAttributes instantiates a new GETExternalGateways200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETExternalGateways200ResponseDataInnerAttributesWithDefaults

`func NewGETExternalGateways200ResponseDataInnerAttributesWithDefaults() *GETExternalGateways200ResponseDataInnerAttributes`

NewGETExternalGateways200ResponseDataInnerAttributesWithDefaults instantiates a new GETExternalGateways200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETExternalGateways200ResponseDataInnerAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreatedAt

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETExternalGateways200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETExternalGateways200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETExternalGateways200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETExternalGateways200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETExternalGateways200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetSharedSecret

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetSharedSecret() interface{}`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetSharedSecretOk() (*interface{}, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetSharedSecret(v interface{})`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### SetSharedSecretNil

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetSharedSecretNil(b bool)`

 SetSharedSecretNil sets the value for SharedSecret to be an explicit nil

### UnsetSharedSecret
`func (o *GETExternalGateways200ResponseDataInnerAttributes) UnsetSharedSecret()`

UnsetSharedSecret ensures that no value is present for SharedSecret, not even an explicit nil
### GetAuthorizeUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetAuthorizeUrl() interface{}`

GetAuthorizeUrl returns the AuthorizeUrl field if non-nil, zero value otherwise.

### GetAuthorizeUrlOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetAuthorizeUrlOk() (*interface{}, bool)`

GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetAuthorizeUrl(v interface{})`

SetAuthorizeUrl sets AuthorizeUrl field to given value.

### HasAuthorizeUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasAuthorizeUrl() bool`

HasAuthorizeUrl returns a boolean if a field has been set.

### SetAuthorizeUrlNil

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetAuthorizeUrlNil(b bool)`

 SetAuthorizeUrlNil sets the value for AuthorizeUrl to be an explicit nil

### UnsetAuthorizeUrl
`func (o *GETExternalGateways200ResponseDataInnerAttributes) UnsetAuthorizeUrl()`

UnsetAuthorizeUrl ensures that no value is present for AuthorizeUrl, not even an explicit nil
### GetCaptureUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetCaptureUrl() interface{}`

GetCaptureUrl returns the CaptureUrl field if non-nil, zero value otherwise.

### GetCaptureUrlOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetCaptureUrlOk() (*interface{}, bool)`

GetCaptureUrlOk returns a tuple with the CaptureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetCaptureUrl(v interface{})`

SetCaptureUrl sets CaptureUrl field to given value.

### HasCaptureUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasCaptureUrl() bool`

HasCaptureUrl returns a boolean if a field has been set.

### SetCaptureUrlNil

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetCaptureUrlNil(b bool)`

 SetCaptureUrlNil sets the value for CaptureUrl to be an explicit nil

### UnsetCaptureUrl
`func (o *GETExternalGateways200ResponseDataInnerAttributes) UnsetCaptureUrl()`

UnsetCaptureUrl ensures that no value is present for CaptureUrl, not even an explicit nil
### GetVoidUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetVoidUrl() interface{}`

GetVoidUrl returns the VoidUrl field if non-nil, zero value otherwise.

### GetVoidUrlOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetVoidUrlOk() (*interface{}, bool)`

GetVoidUrlOk returns a tuple with the VoidUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetVoidUrl(v interface{})`

SetVoidUrl sets VoidUrl field to given value.

### HasVoidUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasVoidUrl() bool`

HasVoidUrl returns a boolean if a field has been set.

### SetVoidUrlNil

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetVoidUrlNil(b bool)`

 SetVoidUrlNil sets the value for VoidUrl to be an explicit nil

### UnsetVoidUrl
`func (o *GETExternalGateways200ResponseDataInnerAttributes) UnsetVoidUrl()`

UnsetVoidUrl ensures that no value is present for VoidUrl, not even an explicit nil
### GetRefundUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetRefundUrl() interface{}`

GetRefundUrl returns the RefundUrl field if non-nil, zero value otherwise.

### GetRefundUrlOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetRefundUrlOk() (*interface{}, bool)`

GetRefundUrlOk returns a tuple with the RefundUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetRefundUrl(v interface{})`

SetRefundUrl sets RefundUrl field to given value.

### HasRefundUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasRefundUrl() bool`

HasRefundUrl returns a boolean if a field has been set.

### SetRefundUrlNil

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetRefundUrlNil(b bool)`

 SetRefundUrlNil sets the value for RefundUrl to be an explicit nil

### UnsetRefundUrl
`func (o *GETExternalGateways200ResponseDataInnerAttributes) UnsetRefundUrl()`

UnsetRefundUrl ensures that no value is present for RefundUrl, not even an explicit nil
### GetTokenUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetTokenUrl() interface{}`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetTokenUrlOk() (*interface{}, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetTokenUrl(v interface{})`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### SetTokenUrlNil

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetTokenUrlNil(b bool)`

 SetTokenUrlNil sets the value for TokenUrl to be an explicit nil

### UnsetTokenUrl
`func (o *GETExternalGateways200ResponseDataInnerAttributes) UnsetTokenUrl()`

UnsetTokenUrl ensures that no value is present for TokenUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


