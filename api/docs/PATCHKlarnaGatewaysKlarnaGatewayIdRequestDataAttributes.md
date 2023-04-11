# PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The payment gateway&#39;s internal name. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**CountryCode** | Pointer to **interface{}** | The gateway country code one of EU, US, or OC. | [optional] 
**ApiKey** | Pointer to **interface{}** | The public key linked to your API credential. | [optional] 
**ApiSecret** | Pointer to **interface{}** | The gateway API key. | [optional] 

## Methods

### NewPATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes

`func NewPATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes() *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes`

NewPATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes instantiates a new PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributesWithDefaults

`func NewPATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributesWithDefaults() *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes`

NewPATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributesWithDefaults instantiates a new PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReference

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCountryCode

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) GetCountryCode() interface{}`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) GetCountryCodeOk() (*interface{}, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) SetCountryCode(v interface{})`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetApiKey

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) GetApiKey() interface{}`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) GetApiKeyOk() (*interface{}, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) SetApiKey(v interface{})`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetApiSecret

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) GetApiSecret() interface{}`

GetApiSecret returns the ApiSecret field if non-nil, zero value otherwise.

### GetApiSecretOk

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) GetApiSecretOk() (*interface{}, bool)`

GetApiSecretOk returns a tuple with the ApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSecret

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) SetApiSecret(v interface{})`

SetApiSecret sets ApiSecret field to given value.

### HasApiSecret

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) HasApiSecret() bool`

HasApiSecret returns a boolean if a field has been set.

### SetApiSecretNil

`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) SetApiSecretNil(b bool)`

 SetApiSecretNil sets the value for ApiSecret to be an explicit nil

### UnsetApiSecret
`func (o *PATCHKlarnaGatewaysKlarnaGatewayIdRequestDataAttributes) UnsetApiSecret()`

UnsetApiSecret ensures that no value is present for ApiSecret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


