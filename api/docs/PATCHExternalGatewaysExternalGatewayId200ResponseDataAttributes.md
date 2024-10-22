# PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The payment gateway&#39;s internal name. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**AuthorizeUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to authorize payments. | [optional] 
**CaptureUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to capture payments. | [optional] 
**VoidUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to void payments. | [optional] 
**RefundUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to refund payments. | [optional] 
**TokenUrl** | Pointer to **interface{}** | The endpoint used by the external gateway to create a customer payment token. | [optional] 
**ResetCircuit** | Pointer to **interface{}** | Send this attribute if you want to reset the circuit breaker associated to this resource to &#39;closed&#39; state and zero failures count. Cannot be passed by sales channels. | [optional] 

## Methods

### NewPATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes

`func NewPATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes() *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes`

NewPATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes instantiates a new PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHExternalGatewaysExternalGatewayId200ResponseDataAttributesWithDefaults

`func NewPATCHExternalGatewaysExternalGatewayId200ResponseDataAttributesWithDefaults() *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes`

NewPATCHExternalGatewaysExternalGatewayId200ResponseDataAttributesWithDefaults instantiates a new PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReference

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAuthorizeUrl

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetAuthorizeUrl() interface{}`

GetAuthorizeUrl returns the AuthorizeUrl field if non-nil, zero value otherwise.

### GetAuthorizeUrlOk

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetAuthorizeUrlOk() (*interface{}, bool)`

GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeUrl

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetAuthorizeUrl(v interface{})`

SetAuthorizeUrl sets AuthorizeUrl field to given value.

### HasAuthorizeUrl

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasAuthorizeUrl() bool`

HasAuthorizeUrl returns a boolean if a field has been set.

### SetAuthorizeUrlNil

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetAuthorizeUrlNil(b bool)`

 SetAuthorizeUrlNil sets the value for AuthorizeUrl to be an explicit nil

### UnsetAuthorizeUrl
`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetAuthorizeUrl()`

UnsetAuthorizeUrl ensures that no value is present for AuthorizeUrl, not even an explicit nil
### GetCaptureUrl

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetCaptureUrl() interface{}`

GetCaptureUrl returns the CaptureUrl field if non-nil, zero value otherwise.

### GetCaptureUrlOk

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetCaptureUrlOk() (*interface{}, bool)`

GetCaptureUrlOk returns a tuple with the CaptureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureUrl

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetCaptureUrl(v interface{})`

SetCaptureUrl sets CaptureUrl field to given value.

### HasCaptureUrl

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasCaptureUrl() bool`

HasCaptureUrl returns a boolean if a field has been set.

### SetCaptureUrlNil

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetCaptureUrlNil(b bool)`

 SetCaptureUrlNil sets the value for CaptureUrl to be an explicit nil

### UnsetCaptureUrl
`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetCaptureUrl()`

UnsetCaptureUrl ensures that no value is present for CaptureUrl, not even an explicit nil
### GetVoidUrl

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetVoidUrl() interface{}`

GetVoidUrl returns the VoidUrl field if non-nil, zero value otherwise.

### GetVoidUrlOk

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetVoidUrlOk() (*interface{}, bool)`

GetVoidUrlOk returns a tuple with the VoidUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidUrl

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetVoidUrl(v interface{})`

SetVoidUrl sets VoidUrl field to given value.

### HasVoidUrl

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasVoidUrl() bool`

HasVoidUrl returns a boolean if a field has been set.

### SetVoidUrlNil

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetVoidUrlNil(b bool)`

 SetVoidUrlNil sets the value for VoidUrl to be an explicit nil

### UnsetVoidUrl
`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetVoidUrl()`

UnsetVoidUrl ensures that no value is present for VoidUrl, not even an explicit nil
### GetRefundUrl

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetRefundUrl() interface{}`

GetRefundUrl returns the RefundUrl field if non-nil, zero value otherwise.

### GetRefundUrlOk

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetRefundUrlOk() (*interface{}, bool)`

GetRefundUrlOk returns a tuple with the RefundUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundUrl

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetRefundUrl(v interface{})`

SetRefundUrl sets RefundUrl field to given value.

### HasRefundUrl

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasRefundUrl() bool`

HasRefundUrl returns a boolean if a field has been set.

### SetRefundUrlNil

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetRefundUrlNil(b bool)`

 SetRefundUrlNil sets the value for RefundUrl to be an explicit nil

### UnsetRefundUrl
`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetRefundUrl()`

UnsetRefundUrl ensures that no value is present for RefundUrl, not even an explicit nil
### GetTokenUrl

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetTokenUrl() interface{}`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetTokenUrlOk() (*interface{}, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetTokenUrl(v interface{})`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### SetTokenUrlNil

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetTokenUrlNil(b bool)`

 SetTokenUrlNil sets the value for TokenUrl to be an explicit nil

### UnsetTokenUrl
`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetTokenUrl()`

UnsetTokenUrl ensures that no value is present for TokenUrl, not even an explicit nil
### GetResetCircuit

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetResetCircuit() interface{}`

GetResetCircuit returns the ResetCircuit field if non-nil, zero value otherwise.

### GetResetCircuitOk

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetResetCircuitOk() (*interface{}, bool)`

GetResetCircuitOk returns a tuple with the ResetCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCircuit

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetResetCircuit(v interface{})`

SetResetCircuit sets ResetCircuit field to given value.

### HasResetCircuit

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasResetCircuit() bool`

HasResetCircuit returns a boolean if a field has been set.

### SetResetCircuitNil

`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetResetCircuitNil(b bool)`

 SetResetCircuitNil sets the value for ResetCircuit to be an explicit nil

### UnsetResetCircuit
`func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnsetResetCircuit()`

UnsetResetCircuit ensures that no value is present for ResetCircuit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


