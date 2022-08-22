# POSTAdyenGateways201ResponseDataAttributes

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

### NewPOSTAdyenGateways201ResponseDataAttributes

`func NewPOSTAdyenGateways201ResponseDataAttributes(name string, merchantAccount string, apiKey string, liveUrlPrefix string, ) *POSTAdyenGateways201ResponseDataAttributes`

NewPOSTAdyenGateways201ResponseDataAttributes instantiates a new POSTAdyenGateways201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTAdyenGateways201ResponseDataAttributesWithDefaults

`func NewPOSTAdyenGateways201ResponseDataAttributesWithDefaults() *POSTAdyenGateways201ResponseDataAttributes`

NewPOSTAdyenGateways201ResponseDataAttributesWithDefaults instantiates a new POSTAdyenGateways201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTAdyenGateways201ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTAdyenGateways201ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTAdyenGateways201ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetReference

`func (o *POSTAdyenGateways201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTAdyenGateways201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTAdyenGateways201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTAdyenGateways201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTAdyenGateways201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTAdyenGateways201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTAdyenGateways201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTAdyenGateways201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTAdyenGateways201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTAdyenGateways201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTAdyenGateways201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTAdyenGateways201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *POSTAdyenGateways201ResponseDataAttributes) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *POSTAdyenGateways201ResponseDataAttributes) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *POSTAdyenGateways201ResponseDataAttributes) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetApiKey

`func (o *POSTAdyenGateways201ResponseDataAttributes) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *POSTAdyenGateways201ResponseDataAttributes) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *POSTAdyenGateways201ResponseDataAttributes) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetPublicKey

`func (o *POSTAdyenGateways201ResponseDataAttributes) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *POSTAdyenGateways201ResponseDataAttributes) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *POSTAdyenGateways201ResponseDataAttributes) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *POSTAdyenGateways201ResponseDataAttributes) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetLiveUrlPrefix

`func (o *POSTAdyenGateways201ResponseDataAttributes) GetLiveUrlPrefix() string`

GetLiveUrlPrefix returns the LiveUrlPrefix field if non-nil, zero value otherwise.

### GetLiveUrlPrefixOk

`func (o *POSTAdyenGateways201ResponseDataAttributes) GetLiveUrlPrefixOk() (*string, bool)`

GetLiveUrlPrefixOk returns a tuple with the LiveUrlPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveUrlPrefix

`func (o *POSTAdyenGateways201ResponseDataAttributes) SetLiveUrlPrefix(v string)`

SetLiveUrlPrefix sets LiveUrlPrefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


