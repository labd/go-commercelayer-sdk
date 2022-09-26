# PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The payment gateway&#39;s internal name. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**MerchantAccountId** | Pointer to **string** | The gateway merchant account ID. | [optional] 
**MerchantId** | Pointer to **string** | The gateway merchant ID. | [optional] 
**PublicKey** | Pointer to **string** | The gateway API public key. | [optional] 
**PrivateKey** | Pointer to **string** | The gateway API private key. | [optional] 
**DescriptorName** | Pointer to **string** | The dynamic descriptor name. Must be composed by business name (3, 7 or 12 chars), an asterisk (*) and the product name (18, 14 or 9 chars), for a total length of 22 chars. | [optional] 
**DescriptorPhone** | Pointer to **string** | The dynamic descriptor phone number. Must be 10-14 characters and can only contain numbers, dashes, parentheses and periods. | [optional] 
**DescriptorUrl** | Pointer to **string** | The dynamic descriptor URL. | [optional] 

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

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReference

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### GetMerchantId

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPublicKey

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetPrivateKey

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetDescriptorName

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorName() string`

GetDescriptorName returns the DescriptorName field if non-nil, zero value otherwise.

### GetDescriptorNameOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorNameOk() (*string, bool)`

GetDescriptorNameOk returns a tuple with the DescriptorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorName

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorName(v string)`

SetDescriptorName sets DescriptorName field to given value.

### HasDescriptorName

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasDescriptorName() bool`

HasDescriptorName returns a boolean if a field has been set.

### GetDescriptorPhone

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorPhone() string`

GetDescriptorPhone returns the DescriptorPhone field if non-nil, zero value otherwise.

### GetDescriptorPhoneOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorPhoneOk() (*string, bool)`

GetDescriptorPhoneOk returns a tuple with the DescriptorPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorPhone

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorPhone(v string)`

SetDescriptorPhone sets DescriptorPhone field to given value.

### HasDescriptorPhone

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasDescriptorPhone() bool`

HasDescriptorPhone returns a boolean if a field has been set.

### GetDescriptorUrl

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorUrl() string`

GetDescriptorUrl returns the DescriptorUrl field if non-nil, zero value otherwise.

### GetDescriptorUrlOk

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorUrlOk() (*string, bool)`

GetDescriptorUrlOk returns a tuple with the DescriptorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorUrl

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorUrl(v string)`

SetDescriptorUrl sets DescriptorUrl field to given value.

### HasDescriptorUrl

`func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasDescriptorUrl() bool`

HasDescriptorUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


