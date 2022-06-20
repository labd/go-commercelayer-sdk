# AdyenGatewayUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The payment gateway&#39;s internal name. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**MerchantAccount** | Pointer to **string** | The gateway merchant account. | [optional] 
**ApiKey** | Pointer to **string** | The gateway API key. | [optional] 
**PublicKey** | Pointer to **string** | The public key linked to your API credential. | [optional] 
**LiveUrlPrefix** | Pointer to **string** | The prefix of the endpoint used for live transactions. | [optional] 

## Methods

### NewAdyenGatewayUpdateDataAttributes

`func NewAdyenGatewayUpdateDataAttributes() *AdyenGatewayUpdateDataAttributes`

NewAdyenGatewayUpdateDataAttributes instantiates a new AdyenGatewayUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdyenGatewayUpdateDataAttributesWithDefaults

`func NewAdyenGatewayUpdateDataAttributesWithDefaults() *AdyenGatewayUpdateDataAttributes`

NewAdyenGatewayUpdateDataAttributesWithDefaults instantiates a new AdyenGatewayUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdyenGatewayUpdateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdyenGatewayUpdateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdyenGatewayUpdateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdyenGatewayUpdateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReference

`func (o *AdyenGatewayUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AdyenGatewayUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AdyenGatewayUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AdyenGatewayUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *AdyenGatewayUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *AdyenGatewayUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *AdyenGatewayUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *AdyenGatewayUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *AdyenGatewayUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AdyenGatewayUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AdyenGatewayUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AdyenGatewayUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *AdyenGatewayUpdateDataAttributes) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *AdyenGatewayUpdateDataAttributes) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *AdyenGatewayUpdateDataAttributes) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.

### HasMerchantAccount

`func (o *AdyenGatewayUpdateDataAttributes) HasMerchantAccount() bool`

HasMerchantAccount returns a boolean if a field has been set.

### GetApiKey

`func (o *AdyenGatewayUpdateDataAttributes) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AdyenGatewayUpdateDataAttributes) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AdyenGatewayUpdateDataAttributes) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *AdyenGatewayUpdateDataAttributes) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *AdyenGatewayUpdateDataAttributes) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *AdyenGatewayUpdateDataAttributes) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *AdyenGatewayUpdateDataAttributes) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *AdyenGatewayUpdateDataAttributes) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetLiveUrlPrefix

`func (o *AdyenGatewayUpdateDataAttributes) GetLiveUrlPrefix() string`

GetLiveUrlPrefix returns the LiveUrlPrefix field if non-nil, zero value otherwise.

### GetLiveUrlPrefixOk

`func (o *AdyenGatewayUpdateDataAttributes) GetLiveUrlPrefixOk() (*string, bool)`

GetLiveUrlPrefixOk returns a tuple with the LiveUrlPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveUrlPrefix

`func (o *AdyenGatewayUpdateDataAttributes) SetLiveUrlPrefix(v string)`

SetLiveUrlPrefix sets LiveUrlPrefix field to given value.

### HasLiveUrlPrefix

`func (o *AdyenGatewayUpdateDataAttributes) HasLiveUrlPrefix() bool`

HasLiveUrlPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


