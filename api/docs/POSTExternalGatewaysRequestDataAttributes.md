# POSTExternalGatewaysRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The payment gateway&#39;s internal name. | 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**AuthorizeUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to authorize payments. | [optional] 
**CaptureUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to capture payments. | [optional] 
**VoidUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to void payments. | [optional] 
**RefundUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to refund payments. | [optional] 
**TokenUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to create a customer payment token. | [optional] 

## Methods

### NewPOSTExternalGatewaysRequestDataAttributes

`func NewPOSTExternalGatewaysRequestDataAttributes(name interface{}, ) *POSTExternalGatewaysRequestDataAttributes`

NewPOSTExternalGatewaysRequestDataAttributes instantiates a new POSTExternalGatewaysRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTExternalGatewaysRequestDataAttributesWithDefaults

`func NewPOSTExternalGatewaysRequestDataAttributesWithDefaults() *POSTExternalGatewaysRequestDataAttributes`

NewPOSTExternalGatewaysRequestDataAttributesWithDefaults instantiates a new POSTExternalGatewaysRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTExternalGatewaysRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTExternalGatewaysRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTExternalGatewaysRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTExternalGatewaysRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTExternalGatewaysRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReference

`func (o *POSTExternalGatewaysRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTExternalGatewaysRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTExternalGatewaysRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTExternalGatewaysRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTExternalGatewaysRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTExternalGatewaysRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTExternalGatewaysRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTExternalGatewaysRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTExternalGatewaysRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTExternalGatewaysRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTExternalGatewaysRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTExternalGatewaysRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTExternalGatewaysRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTExternalGatewaysRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTExternalGatewaysRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTExternalGatewaysRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTExternalGatewaysRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTExternalGatewaysRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAuthorizeUrl

`func (o *POSTExternalGatewaysRequestDataAttributes) GetAuthorizeUrl() interface{}`

GetAuthorizeUrl returns the AuthorizeUrl field if non-nil, zero value otherwise.

### GetAuthorizeUrlOk

`func (o *POSTExternalGatewaysRequestDataAttributes) GetAuthorizeUrlOk() (*interface{}, bool)`

GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeUrl

`func (o *POSTExternalGatewaysRequestDataAttributes) SetAuthorizeUrl(v interface{})`

SetAuthorizeUrl sets AuthorizeUrl field to given value.

### HasAuthorizeUrl

`func (o *POSTExternalGatewaysRequestDataAttributes) HasAuthorizeUrl() bool`

HasAuthorizeUrl returns a boolean if a field has been set.

### SetAuthorizeUrlNil

`func (o *POSTExternalGatewaysRequestDataAttributes) SetAuthorizeUrlNil(b bool)`

 SetAuthorizeUrlNil sets the value for AuthorizeUrl to be an explicit nil

### UnsetAuthorizeUrl
`func (o *POSTExternalGatewaysRequestDataAttributes) UnsetAuthorizeUrl()`

UnsetAuthorizeUrl ensures that no value is present for AuthorizeUrl, not even an explicit nil
### GetCaptureUrl

`func (o *POSTExternalGatewaysRequestDataAttributes) GetCaptureUrl() interface{}`

GetCaptureUrl returns the CaptureUrl field if non-nil, zero value otherwise.

### GetCaptureUrlOk

`func (o *POSTExternalGatewaysRequestDataAttributes) GetCaptureUrlOk() (*interface{}, bool)`

GetCaptureUrlOk returns a tuple with the CaptureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureUrl

`func (o *POSTExternalGatewaysRequestDataAttributes) SetCaptureUrl(v interface{})`

SetCaptureUrl sets CaptureUrl field to given value.

### HasCaptureUrl

`func (o *POSTExternalGatewaysRequestDataAttributes) HasCaptureUrl() bool`

HasCaptureUrl returns a boolean if a field has been set.

### SetCaptureUrlNil

`func (o *POSTExternalGatewaysRequestDataAttributes) SetCaptureUrlNil(b bool)`

 SetCaptureUrlNil sets the value for CaptureUrl to be an explicit nil

### UnsetCaptureUrl
`func (o *POSTExternalGatewaysRequestDataAttributes) UnsetCaptureUrl()`

UnsetCaptureUrl ensures that no value is present for CaptureUrl, not even an explicit nil
### GetVoidUrl

`func (o *POSTExternalGatewaysRequestDataAttributes) GetVoidUrl() interface{}`

GetVoidUrl returns the VoidUrl field if non-nil, zero value otherwise.

### GetVoidUrlOk

`func (o *POSTExternalGatewaysRequestDataAttributes) GetVoidUrlOk() (*interface{}, bool)`

GetVoidUrlOk returns a tuple with the VoidUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidUrl

`func (o *POSTExternalGatewaysRequestDataAttributes) SetVoidUrl(v interface{})`

SetVoidUrl sets VoidUrl field to given value.

### HasVoidUrl

`func (o *POSTExternalGatewaysRequestDataAttributes) HasVoidUrl() bool`

HasVoidUrl returns a boolean if a field has been set.

### SetVoidUrlNil

`func (o *POSTExternalGatewaysRequestDataAttributes) SetVoidUrlNil(b bool)`

 SetVoidUrlNil sets the value for VoidUrl to be an explicit nil

### UnsetVoidUrl
`func (o *POSTExternalGatewaysRequestDataAttributes) UnsetVoidUrl()`

UnsetVoidUrl ensures that no value is present for VoidUrl, not even an explicit nil
### GetRefundUrl

`func (o *POSTExternalGatewaysRequestDataAttributes) GetRefundUrl() interface{}`

GetRefundUrl returns the RefundUrl field if non-nil, zero value otherwise.

### GetRefundUrlOk

`func (o *POSTExternalGatewaysRequestDataAttributes) GetRefundUrlOk() (*interface{}, bool)`

GetRefundUrlOk returns a tuple with the RefundUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundUrl

`func (o *POSTExternalGatewaysRequestDataAttributes) SetRefundUrl(v interface{})`

SetRefundUrl sets RefundUrl field to given value.

### HasRefundUrl

`func (o *POSTExternalGatewaysRequestDataAttributes) HasRefundUrl() bool`

HasRefundUrl returns a boolean if a field has been set.

### SetRefundUrlNil

`func (o *POSTExternalGatewaysRequestDataAttributes) SetRefundUrlNil(b bool)`

 SetRefundUrlNil sets the value for RefundUrl to be an explicit nil

### UnsetRefundUrl
`func (o *POSTExternalGatewaysRequestDataAttributes) UnsetRefundUrl()`

UnsetRefundUrl ensures that no value is present for RefundUrl, not even an explicit nil
### GetTokenUrl

`func (o *POSTExternalGatewaysRequestDataAttributes) GetTokenUrl() interface{}`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *POSTExternalGatewaysRequestDataAttributes) GetTokenUrlOk() (*interface{}, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *POSTExternalGatewaysRequestDataAttributes) SetTokenUrl(v interface{})`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *POSTExternalGatewaysRequestDataAttributes) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### SetTokenUrlNil

`func (o *POSTExternalGatewaysRequestDataAttributes) SetTokenUrlNil(b bool)`

 SetTokenUrlNil sets the value for TokenUrl to be an explicit nil

### UnsetTokenUrl
`func (o *POSTExternalGatewaysRequestDataAttributes) UnsetTokenUrl()`

UnsetTokenUrl ensures that no value is present for TokenUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


