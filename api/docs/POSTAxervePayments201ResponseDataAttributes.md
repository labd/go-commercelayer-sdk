# POSTAxervePayments201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnUrl** | **interface{}** | The URL where the payer is redirected after they approve the payment. | 
**ClientIp** | Pointer to **interface{}** | The IP adress of the client creating the payment. | [optional] 
**BuyerDetails** | Pointer to **interface{}** | The details of the buyer creating the payment. | [optional] 
**RequestToken** | Pointer to **interface{}** | Requires the creation of a token to represent this payment, mandatory to use customer&#39;s wallet and order subscriptions. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTAxervePayments201ResponseDataAttributes

`func NewPOSTAxervePayments201ResponseDataAttributes(returnUrl interface{}, ) *POSTAxervePayments201ResponseDataAttributes`

NewPOSTAxervePayments201ResponseDataAttributes instantiates a new POSTAxervePayments201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTAxervePayments201ResponseDataAttributesWithDefaults

`func NewPOSTAxervePayments201ResponseDataAttributesWithDefaults() *POSTAxervePayments201ResponseDataAttributes`

NewPOSTAxervePayments201ResponseDataAttributesWithDefaults instantiates a new POSTAxervePayments201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnUrl

`func (o *POSTAxervePayments201ResponseDataAttributes) GetReturnUrl() interface{}`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *POSTAxervePayments201ResponseDataAttributes) GetReturnUrlOk() (*interface{}, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *POSTAxervePayments201ResponseDataAttributes) SetReturnUrl(v interface{})`

SetReturnUrl sets ReturnUrl field to given value.


### SetReturnUrlNil

`func (o *POSTAxervePayments201ResponseDataAttributes) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *POSTAxervePayments201ResponseDataAttributes) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetClientIp

`func (o *POSTAxervePayments201ResponseDataAttributes) GetClientIp() interface{}`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *POSTAxervePayments201ResponseDataAttributes) GetClientIpOk() (*interface{}, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *POSTAxervePayments201ResponseDataAttributes) SetClientIp(v interface{})`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *POSTAxervePayments201ResponseDataAttributes) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### SetClientIpNil

`func (o *POSTAxervePayments201ResponseDataAttributes) SetClientIpNil(b bool)`

 SetClientIpNil sets the value for ClientIp to be an explicit nil

### UnsetClientIp
`func (o *POSTAxervePayments201ResponseDataAttributes) UnsetClientIp()`

UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
### GetBuyerDetails

`func (o *POSTAxervePayments201ResponseDataAttributes) GetBuyerDetails() interface{}`

GetBuyerDetails returns the BuyerDetails field if non-nil, zero value otherwise.

### GetBuyerDetailsOk

`func (o *POSTAxervePayments201ResponseDataAttributes) GetBuyerDetailsOk() (*interface{}, bool)`

GetBuyerDetailsOk returns a tuple with the BuyerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerDetails

`func (o *POSTAxervePayments201ResponseDataAttributes) SetBuyerDetails(v interface{})`

SetBuyerDetails sets BuyerDetails field to given value.

### HasBuyerDetails

`func (o *POSTAxervePayments201ResponseDataAttributes) HasBuyerDetails() bool`

HasBuyerDetails returns a boolean if a field has been set.

### SetBuyerDetailsNil

`func (o *POSTAxervePayments201ResponseDataAttributes) SetBuyerDetailsNil(b bool)`

 SetBuyerDetailsNil sets the value for BuyerDetails to be an explicit nil

### UnsetBuyerDetails
`func (o *POSTAxervePayments201ResponseDataAttributes) UnsetBuyerDetails()`

UnsetBuyerDetails ensures that no value is present for BuyerDetails, not even an explicit nil
### GetRequestToken

`func (o *POSTAxervePayments201ResponseDataAttributes) GetRequestToken() interface{}`

GetRequestToken returns the RequestToken field if non-nil, zero value otherwise.

### GetRequestTokenOk

`func (o *POSTAxervePayments201ResponseDataAttributes) GetRequestTokenOk() (*interface{}, bool)`

GetRequestTokenOk returns a tuple with the RequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToken

`func (o *POSTAxervePayments201ResponseDataAttributes) SetRequestToken(v interface{})`

SetRequestToken sets RequestToken field to given value.

### HasRequestToken

`func (o *POSTAxervePayments201ResponseDataAttributes) HasRequestToken() bool`

HasRequestToken returns a boolean if a field has been set.

### SetRequestTokenNil

`func (o *POSTAxervePayments201ResponseDataAttributes) SetRequestTokenNil(b bool)`

 SetRequestTokenNil sets the value for RequestToken to be an explicit nil

### UnsetRequestToken
`func (o *POSTAxervePayments201ResponseDataAttributes) UnsetRequestToken()`

UnsetRequestToken ensures that no value is present for RequestToken, not even an explicit nil
### GetReference

`func (o *POSTAxervePayments201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTAxervePayments201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTAxervePayments201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTAxervePayments201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTAxervePayments201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTAxervePayments201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTAxervePayments201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTAxervePayments201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTAxervePayments201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTAxervePayments201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTAxervePayments201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTAxervePayments201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTAxervePayments201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTAxervePayments201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTAxervePayments201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTAxervePayments201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTAxervePayments201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTAxervePayments201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


