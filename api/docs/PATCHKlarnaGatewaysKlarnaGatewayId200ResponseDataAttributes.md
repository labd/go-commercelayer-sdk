# PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The payment gateway&#39;s internal name. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**CountryCode** | Pointer to **string** | The gateway country code one of EU, US, or OC. | [optional] 
**ApiKey** | Pointer to **string** | The public key linked to your API credential. | [optional] 
**ApiSecret** | Pointer to **string** | The gateway API key. | [optional] 

## Methods

### NewPATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes

`func NewPATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes() *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes`

NewPATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes instantiates a new PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributesWithDefaults

`func NewPATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributesWithDefaults() *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes`

NewPATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributesWithDefaults instantiates a new PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReference

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCountryCode

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetApiKey

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetApiSecret

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetApiSecret() string`

GetApiSecret returns the ApiSecret field if non-nil, zero value otherwise.

### GetApiSecretOk

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetApiSecretOk() (*string, bool)`

GetApiSecretOk returns a tuple with the ApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSecret

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) SetApiSecret(v string)`

SetApiSecret sets ApiSecret field to given value.

### HasApiSecret

`func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) HasApiSecret() bool`

HasApiSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


