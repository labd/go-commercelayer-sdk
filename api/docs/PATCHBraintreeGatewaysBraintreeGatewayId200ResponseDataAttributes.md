# PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The payment gateway&#39;s internal name. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**MerchantAccountId** | Pointer to **interface{}** | The gateway merchant account ID. | [optional] 
**MerchantId** | Pointer to **interface{}** | The gateway merchant ID. | [optional] 
**PublicKey** | Pointer to **interface{}** | The gateway API public key. | [optional] 
**PrivateKey** | Pointer to **interface{}** | The gateway API private key. | [optional] 
**DescriptorName** | Pointer to **interface{}** | The dynamic descriptor name. Must be composed by business name (3, 7 or 12 chars), an asterisk (*) and the product name (18, 14 or 9 chars), for a total length of 22 chars. | [optional] 
**DescriptorPhone** | Pointer to **interface{}** | The dynamic descriptor phone number. Must be 10-14 characters and can only contain numbers, dashes, parentheses and periods. | [optional] 
**DescriptorUrl** | Pointer to **interface{}** | The dynamic descriptor URL. | [optional] 

## Methods

### NewPATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes

`func NewPATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes() *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes`

NewPATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes instantiates a new PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributesWithDefaults

`func NewPATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributesWithDefaults() *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes`

NewPATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributesWithDefaults instantiates a new PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReference

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMerchantAccountId

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetMerchantAccountId() interface{}`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetMerchantAccountIdOk() (*interface{}, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetMerchantAccountId(v interface{})`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### SetMerchantAccountIdNil

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetMerchantAccountIdNil(b bool)`

 SetMerchantAccountIdNil sets the value for MerchantAccountId to be an explicit nil

### UnsetMerchantAccountId
`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetMerchantAccountId()`

UnsetMerchantAccountId ensures that no value is present for MerchantAccountId, not even an explicit nil
### GetMerchantId

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetMerchantId() interface{}`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetMerchantIdOk() (*interface{}, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetMerchantId(v interface{})`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### SetMerchantIdNil

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetMerchantIdNil(b bool)`

 SetMerchantIdNil sets the value for MerchantId to be an explicit nil

### UnsetMerchantId
`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetMerchantId()`

UnsetMerchantId ensures that no value is present for MerchantId, not even an explicit nil
### GetPublicKey

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetPublicKey() interface{}`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetPublicKeyOk() (*interface{}, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetPublicKey(v interface{})`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetPrivateKey

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetPrivateKey() interface{}`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetPrivateKeyOk() (*interface{}, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetPrivateKey(v interface{})`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetDescriptorName

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorName() interface{}`

GetDescriptorName returns the DescriptorName field if non-nil, zero value otherwise.

### GetDescriptorNameOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorNameOk() (*interface{}, bool)`

GetDescriptorNameOk returns a tuple with the DescriptorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorName

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorName(v interface{})`

SetDescriptorName sets DescriptorName field to given value.

### HasDescriptorName

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasDescriptorName() bool`

HasDescriptorName returns a boolean if a field has been set.

### SetDescriptorNameNil

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorNameNil(b bool)`

 SetDescriptorNameNil sets the value for DescriptorName to be an explicit nil

### UnsetDescriptorName
`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetDescriptorName()`

UnsetDescriptorName ensures that no value is present for DescriptorName, not even an explicit nil
### GetDescriptorPhone

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorPhone() interface{}`

GetDescriptorPhone returns the DescriptorPhone field if non-nil, zero value otherwise.

### GetDescriptorPhoneOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorPhoneOk() (*interface{}, bool)`

GetDescriptorPhoneOk returns a tuple with the DescriptorPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorPhone

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorPhone(v interface{})`

SetDescriptorPhone sets DescriptorPhone field to given value.

### HasDescriptorPhone

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasDescriptorPhone() bool`

HasDescriptorPhone returns a boolean if a field has been set.

### SetDescriptorPhoneNil

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorPhoneNil(b bool)`

 SetDescriptorPhoneNil sets the value for DescriptorPhone to be an explicit nil

### UnsetDescriptorPhone
`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetDescriptorPhone()`

UnsetDescriptorPhone ensures that no value is present for DescriptorPhone, not even an explicit nil
### GetDescriptorUrl

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorUrl() interface{}`

GetDescriptorUrl returns the DescriptorUrl field if non-nil, zero value otherwise.

### GetDescriptorUrlOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorUrlOk() (*interface{}, bool)`

GetDescriptorUrlOk returns a tuple with the DescriptorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorUrl

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorUrl(v interface{})`

SetDescriptorUrl sets DescriptorUrl field to given value.

### HasDescriptorUrl

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasDescriptorUrl() bool`

HasDescriptorUrl returns a boolean if a field has been set.

### SetDescriptorUrlNil

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorUrlNil(b bool)`

 SetDescriptorUrlNil sets the value for DescriptorUrl to be an explicit nil

### UnsetDescriptorUrl
`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnsetDescriptorUrl()`

UnsetDescriptorUrl ensures that no value is present for DescriptorUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


