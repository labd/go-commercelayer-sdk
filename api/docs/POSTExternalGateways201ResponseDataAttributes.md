# POSTExternalGateways201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The payment gateway&#39;s internal name. | 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**AuthorizeUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to authorize payments. | [optional] 
**CaptureUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to capture payments. | [optional] 
**VoidUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to void payments. | [optional] 
**RefundUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to refund payments. | [optional] 
**TokenUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to create a customer payment token. | [optional] 

## Methods

### NewPOSTExternalGateways201ResponseDataAttributes

`func NewPOSTExternalGateways201ResponseDataAttributes(name interface{}, ) *POSTExternalGateways201ResponseDataAttributes`

NewPOSTExternalGateways201ResponseDataAttributes instantiates a new POSTExternalGateways201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTExternalGateways201ResponseDataAttributesWithDefaults

`func NewPOSTExternalGateways201ResponseDataAttributesWithDefaults() *POSTExternalGateways201ResponseDataAttributes`

NewPOSTExternalGateways201ResponseDataAttributesWithDefaults instantiates a new POSTExternalGateways201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTExternalGateways201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTExternalGateways201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTExternalGateways201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTExternalGateways201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTExternalGateways201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReference

`func (o *POSTExternalGateways201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTExternalGateways201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTExternalGateways201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTExternalGateways201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTExternalGateways201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTExternalGateways201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTExternalGateways201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTExternalGateways201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTExternalGateways201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTExternalGateways201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTExternalGateways201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTExternalGateways201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTExternalGateways201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTExternalGateways201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTExternalGateways201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTExternalGateways201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTExternalGateways201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTExternalGateways201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAuthorizeUrl

`func (o *POSTExternalGateways201ResponseDataAttributes) GetAuthorizeUrl() interface{}`

GetAuthorizeUrl returns the AuthorizeUrl field if non-nil, zero value otherwise.

### GetAuthorizeUrlOk

`func (o *POSTExternalGateways201ResponseDataAttributes) GetAuthorizeUrlOk() (*interface{}, bool)`

GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeUrl

`func (o *POSTExternalGateways201ResponseDataAttributes) SetAuthorizeUrl(v interface{})`

SetAuthorizeUrl sets AuthorizeUrl field to given value.

### HasAuthorizeUrl

`func (o *POSTExternalGateways201ResponseDataAttributes) HasAuthorizeUrl() bool`

HasAuthorizeUrl returns a boolean if a field has been set.

### SetAuthorizeUrlNil

`func (o *POSTExternalGateways201ResponseDataAttributes) SetAuthorizeUrlNil(b bool)`

 SetAuthorizeUrlNil sets the value for AuthorizeUrl to be an explicit nil

### UnsetAuthorizeUrl
`func (o *POSTExternalGateways201ResponseDataAttributes) UnsetAuthorizeUrl()`

UnsetAuthorizeUrl ensures that no value is present for AuthorizeUrl, not even an explicit nil
### GetCaptureUrl

`func (o *POSTExternalGateways201ResponseDataAttributes) GetCaptureUrl() interface{}`

GetCaptureUrl returns the CaptureUrl field if non-nil, zero value otherwise.

### GetCaptureUrlOk

`func (o *POSTExternalGateways201ResponseDataAttributes) GetCaptureUrlOk() (*interface{}, bool)`

GetCaptureUrlOk returns a tuple with the CaptureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureUrl

`func (o *POSTExternalGateways201ResponseDataAttributes) SetCaptureUrl(v interface{})`

SetCaptureUrl sets CaptureUrl field to given value.

### HasCaptureUrl

`func (o *POSTExternalGateways201ResponseDataAttributes) HasCaptureUrl() bool`

HasCaptureUrl returns a boolean if a field has been set.

### SetCaptureUrlNil

`func (o *POSTExternalGateways201ResponseDataAttributes) SetCaptureUrlNil(b bool)`

 SetCaptureUrlNil sets the value for CaptureUrl to be an explicit nil

### UnsetCaptureUrl
`func (o *POSTExternalGateways201ResponseDataAttributes) UnsetCaptureUrl()`

UnsetCaptureUrl ensures that no value is present for CaptureUrl, not even an explicit nil
### GetVoidUrl

`func (o *POSTExternalGateways201ResponseDataAttributes) GetVoidUrl() interface{}`

GetVoidUrl returns the VoidUrl field if non-nil, zero value otherwise.

### GetVoidUrlOk

`func (o *POSTExternalGateways201ResponseDataAttributes) GetVoidUrlOk() (*interface{}, bool)`

GetVoidUrlOk returns a tuple with the VoidUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidUrl

`func (o *POSTExternalGateways201ResponseDataAttributes) SetVoidUrl(v interface{})`

SetVoidUrl sets VoidUrl field to given value.

### HasVoidUrl

`func (o *POSTExternalGateways201ResponseDataAttributes) HasVoidUrl() bool`

HasVoidUrl returns a boolean if a field has been set.

### SetVoidUrlNil

`func (o *POSTExternalGateways201ResponseDataAttributes) SetVoidUrlNil(b bool)`

 SetVoidUrlNil sets the value for VoidUrl to be an explicit nil

### UnsetVoidUrl
`func (o *POSTExternalGateways201ResponseDataAttributes) UnsetVoidUrl()`

UnsetVoidUrl ensures that no value is present for VoidUrl, not even an explicit nil
### GetRefundUrl

`func (o *POSTExternalGateways201ResponseDataAttributes) GetRefundUrl() interface{}`

GetRefundUrl returns the RefundUrl field if non-nil, zero value otherwise.

### GetRefundUrlOk

`func (o *POSTExternalGateways201ResponseDataAttributes) GetRefundUrlOk() (*interface{}, bool)`

GetRefundUrlOk returns a tuple with the RefundUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundUrl

`func (o *POSTExternalGateways201ResponseDataAttributes) SetRefundUrl(v interface{})`

SetRefundUrl sets RefundUrl field to given value.

### HasRefundUrl

`func (o *POSTExternalGateways201ResponseDataAttributes) HasRefundUrl() bool`

HasRefundUrl returns a boolean if a field has been set.

### SetRefundUrlNil

`func (o *POSTExternalGateways201ResponseDataAttributes) SetRefundUrlNil(b bool)`

 SetRefundUrlNil sets the value for RefundUrl to be an explicit nil

### UnsetRefundUrl
`func (o *POSTExternalGateways201ResponseDataAttributes) UnsetRefundUrl()`

UnsetRefundUrl ensures that no value is present for RefundUrl, not even an explicit nil
### GetTokenUrl

`func (o *POSTExternalGateways201ResponseDataAttributes) GetTokenUrl() interface{}`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *POSTExternalGateways201ResponseDataAttributes) GetTokenUrlOk() (*interface{}, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *POSTExternalGateways201ResponseDataAttributes) SetTokenUrl(v interface{})`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *POSTExternalGateways201ResponseDataAttributes) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### SetTokenUrlNil

`func (o *POSTExternalGateways201ResponseDataAttributes) SetTokenUrlNil(b bool)`

 SetTokenUrlNil sets the value for TokenUrl to be an explicit nil

### UnsetTokenUrl
`func (o *POSTExternalGateways201ResponseDataAttributes) UnsetTokenUrl()`

UnsetTokenUrl ensures that no value is present for TokenUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


