# BraintreeGatewayCreateDataAttributes

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

### NewBraintreeGatewayCreateDataAttributes

`func NewBraintreeGatewayCreateDataAttributes(name string, merchantAccountId string, merchantId string, publicKey string, privateKey string, ) *BraintreeGatewayCreateDataAttributes`

NewBraintreeGatewayCreateDataAttributes instantiates a new BraintreeGatewayCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBraintreeGatewayCreateDataAttributesWithDefaults

`func NewBraintreeGatewayCreateDataAttributesWithDefaults() *BraintreeGatewayCreateDataAttributes`

NewBraintreeGatewayCreateDataAttributesWithDefaults instantiates a new BraintreeGatewayCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BraintreeGatewayCreateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BraintreeGatewayCreateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BraintreeGatewayCreateDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetReference

`func (o *BraintreeGatewayCreateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *BraintreeGatewayCreateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *BraintreeGatewayCreateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *BraintreeGatewayCreateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *BraintreeGatewayCreateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *BraintreeGatewayCreateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *BraintreeGatewayCreateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *BraintreeGatewayCreateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *BraintreeGatewayCreateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BraintreeGatewayCreateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BraintreeGatewayCreateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BraintreeGatewayCreateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *BraintreeGatewayCreateDataAttributes) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *BraintreeGatewayCreateDataAttributes) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *BraintreeGatewayCreateDataAttributes) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.


### GetMerchantId

`func (o *BraintreeGatewayCreateDataAttributes) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *BraintreeGatewayCreateDataAttributes) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *BraintreeGatewayCreateDataAttributes) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetPublicKey

`func (o *BraintreeGatewayCreateDataAttributes) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *BraintreeGatewayCreateDataAttributes) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *BraintreeGatewayCreateDataAttributes) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetPrivateKey

`func (o *BraintreeGatewayCreateDataAttributes) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *BraintreeGatewayCreateDataAttributes) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *BraintreeGatewayCreateDataAttributes) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetDescriptorName

`func (o *BraintreeGatewayCreateDataAttributes) GetDescriptorName() string`

GetDescriptorName returns the DescriptorName field if non-nil, zero value otherwise.

### GetDescriptorNameOk

`func (o *BraintreeGatewayCreateDataAttributes) GetDescriptorNameOk() (*string, bool)`

GetDescriptorNameOk returns a tuple with the DescriptorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorName

`func (o *BraintreeGatewayCreateDataAttributes) SetDescriptorName(v string)`

SetDescriptorName sets DescriptorName field to given value.

### HasDescriptorName

`func (o *BraintreeGatewayCreateDataAttributes) HasDescriptorName() bool`

HasDescriptorName returns a boolean if a field has been set.

### GetDescriptorPhone

`func (o *BraintreeGatewayCreateDataAttributes) GetDescriptorPhone() string`

GetDescriptorPhone returns the DescriptorPhone field if non-nil, zero value otherwise.

### GetDescriptorPhoneOk

`func (o *BraintreeGatewayCreateDataAttributes) GetDescriptorPhoneOk() (*string, bool)`

GetDescriptorPhoneOk returns a tuple with the DescriptorPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorPhone

`func (o *BraintreeGatewayCreateDataAttributes) SetDescriptorPhone(v string)`

SetDescriptorPhone sets DescriptorPhone field to given value.

### HasDescriptorPhone

`func (o *BraintreeGatewayCreateDataAttributes) HasDescriptorPhone() bool`

HasDescriptorPhone returns a boolean if a field has been set.

### GetDescriptorUrl

`func (o *BraintreeGatewayCreateDataAttributes) GetDescriptorUrl() string`

GetDescriptorUrl returns the DescriptorUrl field if non-nil, zero value otherwise.

### GetDescriptorUrlOk

`func (o *BraintreeGatewayCreateDataAttributes) GetDescriptorUrlOk() (*string, bool)`

GetDescriptorUrlOk returns a tuple with the DescriptorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorUrl

`func (o *BraintreeGatewayCreateDataAttributes) SetDescriptorUrl(v string)`

SetDescriptorUrl sets DescriptorUrl field to given value.

### HasDescriptorUrl

`func (o *BraintreeGatewayCreateDataAttributes) HasDescriptorUrl() bool`

HasDescriptorUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


