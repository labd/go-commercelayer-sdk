# ExternalGatewayUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The payment gateway&#39;s internal name. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**AuthorizeUrl** | Pointer to **string** | The endpoint used by the external gateway to authorize payments. | [optional] 
**CaptureUrl** | Pointer to **string** | The endpoint used by the external gateway to capture payments. | [optional] 
**VoidUrl** | Pointer to **string** | The endpoint used by the external gateway to void payments. | [optional] 
**RefundUrl** | Pointer to **string** | The endpoint used by the external gateway to refund payments. | [optional] 
**TokenUrl** | Pointer to **string** | The endpoint used by the external gateway to create a customer payment token. | [optional] 

## Methods

### NewExternalGatewayUpdateDataAttributes

`func NewExternalGatewayUpdateDataAttributes() *ExternalGatewayUpdateDataAttributes`

NewExternalGatewayUpdateDataAttributes instantiates a new ExternalGatewayUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalGatewayUpdateDataAttributesWithDefaults

`func NewExternalGatewayUpdateDataAttributesWithDefaults() *ExternalGatewayUpdateDataAttributes`

NewExternalGatewayUpdateDataAttributesWithDefaults instantiates a new ExternalGatewayUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExternalGatewayUpdateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalGatewayUpdateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalGatewayUpdateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalGatewayUpdateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReference

`func (o *ExternalGatewayUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ExternalGatewayUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ExternalGatewayUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ExternalGatewayUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *ExternalGatewayUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *ExternalGatewayUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *ExternalGatewayUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *ExternalGatewayUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *ExternalGatewayUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ExternalGatewayUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ExternalGatewayUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ExternalGatewayUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAuthorizeUrl

`func (o *ExternalGatewayUpdateDataAttributes) GetAuthorizeUrl() string`

GetAuthorizeUrl returns the AuthorizeUrl field if non-nil, zero value otherwise.

### GetAuthorizeUrlOk

`func (o *ExternalGatewayUpdateDataAttributes) GetAuthorizeUrlOk() (*string, bool)`

GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeUrl

`func (o *ExternalGatewayUpdateDataAttributes) SetAuthorizeUrl(v string)`

SetAuthorizeUrl sets AuthorizeUrl field to given value.

### HasAuthorizeUrl

`func (o *ExternalGatewayUpdateDataAttributes) HasAuthorizeUrl() bool`

HasAuthorizeUrl returns a boolean if a field has been set.

### GetCaptureUrl

`func (o *ExternalGatewayUpdateDataAttributes) GetCaptureUrl() string`

GetCaptureUrl returns the CaptureUrl field if non-nil, zero value otherwise.

### GetCaptureUrlOk

`func (o *ExternalGatewayUpdateDataAttributes) GetCaptureUrlOk() (*string, bool)`

GetCaptureUrlOk returns a tuple with the CaptureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureUrl

`func (o *ExternalGatewayUpdateDataAttributes) SetCaptureUrl(v string)`

SetCaptureUrl sets CaptureUrl field to given value.

### HasCaptureUrl

`func (o *ExternalGatewayUpdateDataAttributes) HasCaptureUrl() bool`

HasCaptureUrl returns a boolean if a field has been set.

### GetVoidUrl

`func (o *ExternalGatewayUpdateDataAttributes) GetVoidUrl() string`

GetVoidUrl returns the VoidUrl field if non-nil, zero value otherwise.

### GetVoidUrlOk

`func (o *ExternalGatewayUpdateDataAttributes) GetVoidUrlOk() (*string, bool)`

GetVoidUrlOk returns a tuple with the VoidUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidUrl

`func (o *ExternalGatewayUpdateDataAttributes) SetVoidUrl(v string)`

SetVoidUrl sets VoidUrl field to given value.

### HasVoidUrl

`func (o *ExternalGatewayUpdateDataAttributes) HasVoidUrl() bool`

HasVoidUrl returns a boolean if a field has been set.

### GetRefundUrl

`func (o *ExternalGatewayUpdateDataAttributes) GetRefundUrl() string`

GetRefundUrl returns the RefundUrl field if non-nil, zero value otherwise.

### GetRefundUrlOk

`func (o *ExternalGatewayUpdateDataAttributes) GetRefundUrlOk() (*string, bool)`

GetRefundUrlOk returns a tuple with the RefundUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundUrl

`func (o *ExternalGatewayUpdateDataAttributes) SetRefundUrl(v string)`

SetRefundUrl sets RefundUrl field to given value.

### HasRefundUrl

`func (o *ExternalGatewayUpdateDataAttributes) HasRefundUrl() bool`

HasRefundUrl returns a boolean if a field has been set.

### GetTokenUrl

`func (o *ExternalGatewayUpdateDataAttributes) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *ExternalGatewayUpdateDataAttributes) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *ExternalGatewayUpdateDataAttributes) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *ExternalGatewayUpdateDataAttributes) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


