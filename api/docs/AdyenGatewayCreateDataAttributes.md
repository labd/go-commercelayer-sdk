# AdyenGatewayCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The payment gateway&#39;s internal name. | 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**MerchantAccount** | **string** | The gateway merchant account. | 
**ApiKey** | **string** | The gateway API key. | 
**PublicKey** | Pointer to **string** | The public key linked to your API credential. | [optional] 
**LiveUrlPrefix** | **string** | The prefix of the endpoint used for live transactions. | 

## Methods

### NewAdyenGatewayCreateDataAttributes

`func NewAdyenGatewayCreateDataAttributes(name string, merchantAccount string, apiKey string, liveUrlPrefix string, ) *AdyenGatewayCreateDataAttributes`

NewAdyenGatewayCreateDataAttributes instantiates a new AdyenGatewayCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdyenGatewayCreateDataAttributesWithDefaults

`func NewAdyenGatewayCreateDataAttributesWithDefaults() *AdyenGatewayCreateDataAttributes`

NewAdyenGatewayCreateDataAttributesWithDefaults instantiates a new AdyenGatewayCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdyenGatewayCreateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdyenGatewayCreateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdyenGatewayCreateDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetReference

`func (o *AdyenGatewayCreateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AdyenGatewayCreateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AdyenGatewayCreateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AdyenGatewayCreateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *AdyenGatewayCreateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *AdyenGatewayCreateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *AdyenGatewayCreateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *AdyenGatewayCreateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *AdyenGatewayCreateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AdyenGatewayCreateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AdyenGatewayCreateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AdyenGatewayCreateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *AdyenGatewayCreateDataAttributes) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *AdyenGatewayCreateDataAttributes) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *AdyenGatewayCreateDataAttributes) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetApiKey

`func (o *AdyenGatewayCreateDataAttributes) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AdyenGatewayCreateDataAttributes) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AdyenGatewayCreateDataAttributes) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetPublicKey

`func (o *AdyenGatewayCreateDataAttributes) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *AdyenGatewayCreateDataAttributes) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *AdyenGatewayCreateDataAttributes) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *AdyenGatewayCreateDataAttributes) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetLiveUrlPrefix

`func (o *AdyenGatewayCreateDataAttributes) GetLiveUrlPrefix() string`

GetLiveUrlPrefix returns the LiveUrlPrefix field if non-nil, zero value otherwise.

### GetLiveUrlPrefixOk

`func (o *AdyenGatewayCreateDataAttributes) GetLiveUrlPrefixOk() (*string, bool)`

GetLiveUrlPrefixOk returns a tuple with the LiveUrlPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveUrlPrefix

`func (o *AdyenGatewayCreateDataAttributes) SetLiveUrlPrefix(v string)`

SetLiveUrlPrefix sets LiveUrlPrefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


