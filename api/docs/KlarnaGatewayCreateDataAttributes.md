# KlarnaGatewayCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The payment gateway&#39;s internal name. | 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**CountryCode** | **string** | The gateway country code one of EU, US, or OC. | 
**ApiKey** | **string** | The public key linked to your API credential. | 
**ApiSecret** | **string** | The gateway API key. | 

## Methods

### NewKlarnaGatewayCreateDataAttributes

`func NewKlarnaGatewayCreateDataAttributes(name string, countryCode string, apiKey string, apiSecret string, ) *KlarnaGatewayCreateDataAttributes`

NewKlarnaGatewayCreateDataAttributes instantiates a new KlarnaGatewayCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKlarnaGatewayCreateDataAttributesWithDefaults

`func NewKlarnaGatewayCreateDataAttributesWithDefaults() *KlarnaGatewayCreateDataAttributes`

NewKlarnaGatewayCreateDataAttributesWithDefaults instantiates a new KlarnaGatewayCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KlarnaGatewayCreateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KlarnaGatewayCreateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KlarnaGatewayCreateDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetReference

`func (o *KlarnaGatewayCreateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *KlarnaGatewayCreateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *KlarnaGatewayCreateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *KlarnaGatewayCreateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *KlarnaGatewayCreateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *KlarnaGatewayCreateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *KlarnaGatewayCreateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *KlarnaGatewayCreateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *KlarnaGatewayCreateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KlarnaGatewayCreateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KlarnaGatewayCreateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KlarnaGatewayCreateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCountryCode

`func (o *KlarnaGatewayCreateDataAttributes) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *KlarnaGatewayCreateDataAttributes) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *KlarnaGatewayCreateDataAttributes) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetApiKey

`func (o *KlarnaGatewayCreateDataAttributes) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *KlarnaGatewayCreateDataAttributes) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *KlarnaGatewayCreateDataAttributes) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetApiSecret

`func (o *KlarnaGatewayCreateDataAttributes) GetApiSecret() string`

GetApiSecret returns the ApiSecret field if non-nil, zero value otherwise.

### GetApiSecretOk

`func (o *KlarnaGatewayCreateDataAttributes) GetApiSecretOk() (*string, bool)`

GetApiSecretOk returns a tuple with the ApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSecret

`func (o *KlarnaGatewayCreateDataAttributes) SetApiSecret(v string)`

SetApiSecret sets ApiSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


