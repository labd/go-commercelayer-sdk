# GETExternalGateways200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The payment gateway&#39;s internal name. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**SharedSecret** | Pointer to **string** | The shared secret used to sign the external request payload. | [optional] 
**AuthorizeUrl** | Pointer to **string** | The endpoint used by the external gateway to authorize payments. | [optional] 
**CaptureUrl** | Pointer to **string** | The endpoint used by the external gateway to capture payments. | [optional] 
**VoidUrl** | Pointer to **string** | The endpoint used by the external gateway to void payments. | [optional] 
**RefundUrl** | Pointer to **string** | The endpoint used by the external gateway to refund payments. | [optional] 
**TokenUrl** | Pointer to **string** | The endpoint used by the external gateway to create a customer payment token. | [optional] 

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

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSharedSecret

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetAuthorizeUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetAuthorizeUrl() string`

GetAuthorizeUrl returns the AuthorizeUrl field if non-nil, zero value otherwise.

### GetAuthorizeUrlOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetAuthorizeUrlOk() (*string, bool)`

GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetAuthorizeUrl(v string)`

SetAuthorizeUrl sets AuthorizeUrl field to given value.

### HasAuthorizeUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasAuthorizeUrl() bool`

HasAuthorizeUrl returns a boolean if a field has been set.

### GetCaptureUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetCaptureUrl() string`

GetCaptureUrl returns the CaptureUrl field if non-nil, zero value otherwise.

### GetCaptureUrlOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetCaptureUrlOk() (*string, bool)`

GetCaptureUrlOk returns a tuple with the CaptureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetCaptureUrl(v string)`

SetCaptureUrl sets CaptureUrl field to given value.

### HasCaptureUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasCaptureUrl() bool`

HasCaptureUrl returns a boolean if a field has been set.

### GetVoidUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetVoidUrl() string`

GetVoidUrl returns the VoidUrl field if non-nil, zero value otherwise.

### GetVoidUrlOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetVoidUrlOk() (*string, bool)`

GetVoidUrlOk returns a tuple with the VoidUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetVoidUrl(v string)`

SetVoidUrl sets VoidUrl field to given value.

### HasVoidUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasVoidUrl() bool`

HasVoidUrl returns a boolean if a field has been set.

### GetRefundUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetRefundUrl() string`

GetRefundUrl returns the RefundUrl field if non-nil, zero value otherwise.

### GetRefundUrlOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetRefundUrlOk() (*string, bool)`

GetRefundUrlOk returns a tuple with the RefundUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetRefundUrl(v string)`

SetRefundUrl sets RefundUrl field to given value.

### HasRefundUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasRefundUrl() bool`

HasRefundUrl returns a boolean if a field has been set.

### GetTokenUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *GETExternalGateways200ResponseDataInnerAttributes) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *GETExternalGateways200ResponseDataInnerAttributes) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


