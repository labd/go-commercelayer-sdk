# POSTBraintreeGateways201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The payment gateway&#39;s internal name. | 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**MerchantAccountId** | **string** | The gateway merchant account ID. | 
**MerchantId** | **string** | The gateway merchant ID. | 
**PublicKey** | **string** | The gateway API public key. | 
**PrivateKey** | **string** | The gateway API private key. | 
**DescriptorName** | Pointer to **string** | The dynamic descriptor name. Must be composed by business name (3, 7 or 12 chars), an asterisk (*) and the product name (18, 14 or 9 chars), for a total length of 22 chars. | [optional] 
**DescriptorPhone** | Pointer to **string** | The dynamic descriptor phone number. Must be 10-14 characters and can only contain numbers, dashes, parentheses and periods. | [optional] 
**DescriptorUrl** | Pointer to **string** | The dynamic descriptor URL. | [optional] 

## Methods

### NewPOSTBraintreeGateways201ResponseDataAttributes

`func NewPOSTBraintreeGateways201ResponseDataAttributes(name string, merchantAccountId string, merchantId string, publicKey string, privateKey string, ) *POSTBraintreeGateways201ResponseDataAttributes`

NewPOSTBraintreeGateways201ResponseDataAttributes instantiates a new POSTBraintreeGateways201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTBraintreeGateways201ResponseDataAttributesWithDefaults

`func NewPOSTBraintreeGateways201ResponseDataAttributesWithDefaults() *POSTBraintreeGateways201ResponseDataAttributes`

NewPOSTBraintreeGateways201ResponseDataAttributesWithDefaults instantiates a new POSTBraintreeGateways201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTBraintreeGateways201ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetReference

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTBraintreeGateways201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTBraintreeGateways201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTBraintreeGateways201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTBraintreeGateways201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTBraintreeGateways201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTBraintreeGateways201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *POSTBraintreeGateways201ResponseDataAttributes) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.


### GetMerchantId

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *POSTBraintreeGateways201ResponseDataAttributes) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetPublicKey

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *POSTBraintreeGateways201ResponseDataAttributes) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetPrivateKey

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *POSTBraintreeGateways201ResponseDataAttributes) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetDescriptorName

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetDescriptorName() string`

GetDescriptorName returns the DescriptorName field if non-nil, zero value otherwise.

### GetDescriptorNameOk

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetDescriptorNameOk() (*string, bool)`

GetDescriptorNameOk returns a tuple with the DescriptorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorName

`func (o *POSTBraintreeGateways201ResponseDataAttributes) SetDescriptorName(v string)`

SetDescriptorName sets DescriptorName field to given value.

### HasDescriptorName

`func (o *POSTBraintreeGateways201ResponseDataAttributes) HasDescriptorName() bool`

HasDescriptorName returns a boolean if a field has been set.

### GetDescriptorPhone

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetDescriptorPhone() string`

GetDescriptorPhone returns the DescriptorPhone field if non-nil, zero value otherwise.

### GetDescriptorPhoneOk

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetDescriptorPhoneOk() (*string, bool)`

GetDescriptorPhoneOk returns a tuple with the DescriptorPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorPhone

`func (o *POSTBraintreeGateways201ResponseDataAttributes) SetDescriptorPhone(v string)`

SetDescriptorPhone sets DescriptorPhone field to given value.

### HasDescriptorPhone

`func (o *POSTBraintreeGateways201ResponseDataAttributes) HasDescriptorPhone() bool`

HasDescriptorPhone returns a boolean if a field has been set.

### GetDescriptorUrl

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetDescriptorUrl() string`

GetDescriptorUrl returns the DescriptorUrl field if non-nil, zero value otherwise.

### GetDescriptorUrlOk

`func (o *POSTBraintreeGateways201ResponseDataAttributes) GetDescriptorUrlOk() (*string, bool)`

GetDescriptorUrlOk returns a tuple with the DescriptorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorUrl

`func (o *POSTBraintreeGateways201ResponseDataAttributes) SetDescriptorUrl(v string)`

SetDescriptorUrl sets DescriptorUrl field to given value.

### HasDescriptorUrl

`func (o *POSTBraintreeGateways201ResponseDataAttributes) HasDescriptorUrl() bool`

HasDescriptorUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


