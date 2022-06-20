# ExternalGatewayDataAttributes

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
**SharedSecret** | Pointer to **string** | The shared secret used to sign gateway payloads. | [optional] 
**AuthorizeUrl** | Pointer to **string** | The endpoint used by the external gateway to authorize payments. | [optional] 
**CaptureUrl** | Pointer to **string** | The endpoint used by the external gateway to capture payments. | [optional] 
**VoidUrl** | Pointer to **string** | The endpoint used by the external gateway to void payments. | [optional] 
**RefundUrl** | Pointer to **string** | The endpoint used by the external gateway to refund payments. | [optional] 
**TokenUrl** | Pointer to **string** | The endpoint used by the external gateway to create a customer payment token. | [optional] 

## Methods

### NewExternalGatewayDataAttributes

`func NewExternalGatewayDataAttributes() *ExternalGatewayDataAttributes`

NewExternalGatewayDataAttributes instantiates a new ExternalGatewayDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalGatewayDataAttributesWithDefaults

`func NewExternalGatewayDataAttributesWithDefaults() *ExternalGatewayDataAttributes`

NewExternalGatewayDataAttributesWithDefaults instantiates a new ExternalGatewayDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExternalGatewayDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalGatewayDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalGatewayDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalGatewayDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ExternalGatewayDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalGatewayDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalGatewayDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalGatewayDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ExternalGatewayDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExternalGatewayDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExternalGatewayDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ExternalGatewayDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ExternalGatewayDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExternalGatewayDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExternalGatewayDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ExternalGatewayDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *ExternalGatewayDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ExternalGatewayDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ExternalGatewayDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ExternalGatewayDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *ExternalGatewayDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *ExternalGatewayDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *ExternalGatewayDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *ExternalGatewayDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *ExternalGatewayDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ExternalGatewayDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ExternalGatewayDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ExternalGatewayDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSharedSecret

`func (o *ExternalGatewayDataAttributes) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *ExternalGatewayDataAttributes) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *ExternalGatewayDataAttributes) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *ExternalGatewayDataAttributes) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetAuthorizeUrl

`func (o *ExternalGatewayDataAttributes) GetAuthorizeUrl() string`

GetAuthorizeUrl returns the AuthorizeUrl field if non-nil, zero value otherwise.

### GetAuthorizeUrlOk

`func (o *ExternalGatewayDataAttributes) GetAuthorizeUrlOk() (*string, bool)`

GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeUrl

`func (o *ExternalGatewayDataAttributes) SetAuthorizeUrl(v string)`

SetAuthorizeUrl sets AuthorizeUrl field to given value.

### HasAuthorizeUrl

`func (o *ExternalGatewayDataAttributes) HasAuthorizeUrl() bool`

HasAuthorizeUrl returns a boolean if a field has been set.

### GetCaptureUrl

`func (o *ExternalGatewayDataAttributes) GetCaptureUrl() string`

GetCaptureUrl returns the CaptureUrl field if non-nil, zero value otherwise.

### GetCaptureUrlOk

`func (o *ExternalGatewayDataAttributes) GetCaptureUrlOk() (*string, bool)`

GetCaptureUrlOk returns a tuple with the CaptureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureUrl

`func (o *ExternalGatewayDataAttributes) SetCaptureUrl(v string)`

SetCaptureUrl sets CaptureUrl field to given value.

### HasCaptureUrl

`func (o *ExternalGatewayDataAttributes) HasCaptureUrl() bool`

HasCaptureUrl returns a boolean if a field has been set.

### GetVoidUrl

`func (o *ExternalGatewayDataAttributes) GetVoidUrl() string`

GetVoidUrl returns the VoidUrl field if non-nil, zero value otherwise.

### GetVoidUrlOk

`func (o *ExternalGatewayDataAttributes) GetVoidUrlOk() (*string, bool)`

GetVoidUrlOk returns a tuple with the VoidUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidUrl

`func (o *ExternalGatewayDataAttributes) SetVoidUrl(v string)`

SetVoidUrl sets VoidUrl field to given value.

### HasVoidUrl

`func (o *ExternalGatewayDataAttributes) HasVoidUrl() bool`

HasVoidUrl returns a boolean if a field has been set.

### GetRefundUrl

`func (o *ExternalGatewayDataAttributes) GetRefundUrl() string`

GetRefundUrl returns the RefundUrl field if non-nil, zero value otherwise.

### GetRefundUrlOk

`func (o *ExternalGatewayDataAttributes) GetRefundUrlOk() (*string, bool)`

GetRefundUrlOk returns a tuple with the RefundUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundUrl

`func (o *ExternalGatewayDataAttributes) SetRefundUrl(v string)`

SetRefundUrl sets RefundUrl field to given value.

### HasRefundUrl

`func (o *ExternalGatewayDataAttributes) HasRefundUrl() bool`

HasRefundUrl returns a boolean if a field has been set.

### GetTokenUrl

`func (o *ExternalGatewayDataAttributes) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *ExternalGatewayDataAttributes) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *ExternalGatewayDataAttributes) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *ExternalGatewayDataAttributes) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


