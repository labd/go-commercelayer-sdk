# POSTBraintreeGatewaysRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The payment gateway&#39;s internal name. | 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**MerchantAccountId** | **interface{}** | The gateway merchant account ID. | 
**MerchantId** | **interface{}** | The gateway merchant ID. | 
**PublicKey** | **interface{}** | The gateway API public key. | 
**PrivateKey** | **interface{}** | The gateway API private key. | 
**DescriptorName** | Pointer to **interface{}** | The dynamic descriptor name. Must be composed by business name (3, 7 or 12 chars), an asterisk (*) and the product name (18, 14 or 9 chars), for a total length of 22 chars. | [optional] 
**DescriptorPhone** | Pointer to **interface{}** | The dynamic descriptor phone number. Must be 10-14 characters and can only contain numbers, dashes, parentheses and periods. | [optional] 
**DescriptorUrl** | Pointer to **interface{}** | The dynamic descriptor URL. | [optional] 

## Methods

### NewPOSTBraintreeGatewaysRequestDataAttributes

`func NewPOSTBraintreeGatewaysRequestDataAttributes(name interface{}, merchantAccountId interface{}, merchantId interface{}, publicKey interface{}, privateKey interface{}, ) *POSTBraintreeGatewaysRequestDataAttributes`

NewPOSTBraintreeGatewaysRequestDataAttributes instantiates a new POSTBraintreeGatewaysRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTBraintreeGatewaysRequestDataAttributesWithDefaults

`func NewPOSTBraintreeGatewaysRequestDataAttributesWithDefaults() *POSTBraintreeGatewaysRequestDataAttributes`

NewPOSTBraintreeGatewaysRequestDataAttributesWithDefaults instantiates a new POSTBraintreeGatewaysRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTBraintreeGatewaysRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReference

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTBraintreeGatewaysRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTBraintreeGatewaysRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTBraintreeGatewaysRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTBraintreeGatewaysRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTBraintreeGatewaysRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTBraintreeGatewaysRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMerchantAccountId

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetMerchantAccountId() interface{}`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetMerchantAccountIdOk() (*interface{}, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetMerchantAccountId(v interface{})`

SetMerchantAccountId sets MerchantAccountId field to given value.


### SetMerchantAccountIdNil

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetMerchantAccountIdNil(b bool)`

 SetMerchantAccountIdNil sets the value for MerchantAccountId to be an explicit nil

### UnsetMerchantAccountId
`func (o *POSTBraintreeGatewaysRequestDataAttributes) UnsetMerchantAccountId()`

UnsetMerchantAccountId ensures that no value is present for MerchantAccountId, not even an explicit nil
### GetMerchantId

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetMerchantId() interface{}`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetMerchantIdOk() (*interface{}, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetMerchantId(v interface{})`

SetMerchantId sets MerchantId field to given value.


### SetMerchantIdNil

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetMerchantIdNil(b bool)`

 SetMerchantIdNil sets the value for MerchantId to be an explicit nil

### UnsetMerchantId
`func (o *POSTBraintreeGatewaysRequestDataAttributes) UnsetMerchantId()`

UnsetMerchantId ensures that no value is present for MerchantId, not even an explicit nil
### GetPublicKey

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetPublicKey() interface{}`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetPublicKeyOk() (*interface{}, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetPublicKey(v interface{})`

SetPublicKey sets PublicKey field to given value.


### SetPublicKeyNil

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *POSTBraintreeGatewaysRequestDataAttributes) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetPrivateKey

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetPrivateKey() interface{}`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetPrivateKeyOk() (*interface{}, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetPrivateKey(v interface{})`

SetPrivateKey sets PrivateKey field to given value.


### SetPrivateKeyNil

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *POSTBraintreeGatewaysRequestDataAttributes) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetDescriptorName

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetDescriptorName() interface{}`

GetDescriptorName returns the DescriptorName field if non-nil, zero value otherwise.

### GetDescriptorNameOk

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetDescriptorNameOk() (*interface{}, bool)`

GetDescriptorNameOk returns a tuple with the DescriptorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorName

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetDescriptorName(v interface{})`

SetDescriptorName sets DescriptorName field to given value.

### HasDescriptorName

`func (o *POSTBraintreeGatewaysRequestDataAttributes) HasDescriptorName() bool`

HasDescriptorName returns a boolean if a field has been set.

### SetDescriptorNameNil

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetDescriptorNameNil(b bool)`

 SetDescriptorNameNil sets the value for DescriptorName to be an explicit nil

### UnsetDescriptorName
`func (o *POSTBraintreeGatewaysRequestDataAttributes) UnsetDescriptorName()`

UnsetDescriptorName ensures that no value is present for DescriptorName, not even an explicit nil
### GetDescriptorPhone

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetDescriptorPhone() interface{}`

GetDescriptorPhone returns the DescriptorPhone field if non-nil, zero value otherwise.

### GetDescriptorPhoneOk

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetDescriptorPhoneOk() (*interface{}, bool)`

GetDescriptorPhoneOk returns a tuple with the DescriptorPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorPhone

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetDescriptorPhone(v interface{})`

SetDescriptorPhone sets DescriptorPhone field to given value.

### HasDescriptorPhone

`func (o *POSTBraintreeGatewaysRequestDataAttributes) HasDescriptorPhone() bool`

HasDescriptorPhone returns a boolean if a field has been set.

### SetDescriptorPhoneNil

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetDescriptorPhoneNil(b bool)`

 SetDescriptorPhoneNil sets the value for DescriptorPhone to be an explicit nil

### UnsetDescriptorPhone
`func (o *POSTBraintreeGatewaysRequestDataAttributes) UnsetDescriptorPhone()`

UnsetDescriptorPhone ensures that no value is present for DescriptorPhone, not even an explicit nil
### GetDescriptorUrl

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetDescriptorUrl() interface{}`

GetDescriptorUrl returns the DescriptorUrl field if non-nil, zero value otherwise.

### GetDescriptorUrlOk

`func (o *POSTBraintreeGatewaysRequestDataAttributes) GetDescriptorUrlOk() (*interface{}, bool)`

GetDescriptorUrlOk returns a tuple with the DescriptorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorUrl

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetDescriptorUrl(v interface{})`

SetDescriptorUrl sets DescriptorUrl field to given value.

### HasDescriptorUrl

`func (o *POSTBraintreeGatewaysRequestDataAttributes) HasDescriptorUrl() bool`

HasDescriptorUrl returns a boolean if a field has been set.

### SetDescriptorUrlNil

`func (o *POSTBraintreeGatewaysRequestDataAttributes) SetDescriptorUrlNil(b bool)`

 SetDescriptorUrlNil sets the value for DescriptorUrl to be an explicit nil

### UnsetDescriptorUrl
`func (o *POSTBraintreeGatewaysRequestDataAttributes) UnsetDescriptorUrl()`

UnsetDescriptorUrl ensures that no value is present for DescriptorUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


