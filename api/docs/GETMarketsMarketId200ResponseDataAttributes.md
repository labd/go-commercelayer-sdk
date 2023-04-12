# GETMarketsMarketId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | Unique identifier for the market (numeric) | [optional] 
**Name** | Pointer to **interface{}** | The market&#39;s internal name | [optional] 
**FacebookPixelId** | Pointer to **interface{}** | The Facebook Pixed ID | [optional] 
**CheckoutUrl** | Pointer to **interface{}** | The checkout URL for this market | [optional] 
**ExternalPricesUrl** | Pointer to **interface{}** | The URL used to overwrite prices by an external source. | [optional] 
**ExternalOrderValidationUrl** | Pointer to **interface{}** | The URL used to validate orders by an external source. | [optional] 
**SharedSecret** | Pointer to **interface{}** | The shared secret used to sign the external requests payload. | [optional] 
**Private** | Pointer to **interface{}** | Indicates if market belongs to a customer_group. | [optional] 
**DisabledAt** | Pointer to **interface{}** | Time at which the market was disabled. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETMarketsMarketId200ResponseDataAttributes

`func NewGETMarketsMarketId200ResponseDataAttributes() *GETMarketsMarketId200ResponseDataAttributes`

NewGETMarketsMarketId200ResponseDataAttributes instantiates a new GETMarketsMarketId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETMarketsMarketId200ResponseDataAttributesWithDefaults

`func NewGETMarketsMarketId200ResponseDataAttributesWithDefaults() *GETMarketsMarketId200ResponseDataAttributes`

NewGETMarketsMarketId200ResponseDataAttributesWithDefaults instantiates a new GETMarketsMarketId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETMarketsMarketId200ResponseDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETMarketsMarketId200ResponseDataAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetName

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETMarketsMarketId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETMarketsMarketId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFacebookPixelId

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetFacebookPixelId() interface{}`

GetFacebookPixelId returns the FacebookPixelId field if non-nil, zero value otherwise.

### GetFacebookPixelIdOk

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetFacebookPixelIdOk() (*interface{}, bool)`

GetFacebookPixelIdOk returns a tuple with the FacebookPixelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookPixelId

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetFacebookPixelId(v interface{})`

SetFacebookPixelId sets FacebookPixelId field to given value.

### HasFacebookPixelId

`func (o *GETMarketsMarketId200ResponseDataAttributes) HasFacebookPixelId() bool`

HasFacebookPixelId returns a boolean if a field has been set.

### SetFacebookPixelIdNil

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetFacebookPixelIdNil(b bool)`

 SetFacebookPixelIdNil sets the value for FacebookPixelId to be an explicit nil

### UnsetFacebookPixelId
`func (o *GETMarketsMarketId200ResponseDataAttributes) UnsetFacebookPixelId()`

UnsetFacebookPixelId ensures that no value is present for FacebookPixelId, not even an explicit nil
### GetCheckoutUrl

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetCheckoutUrl() interface{}`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetCheckoutUrlOk() (*interface{}, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetCheckoutUrl(v interface{})`

SetCheckoutUrl sets CheckoutUrl field to given value.

### HasCheckoutUrl

`func (o *GETMarketsMarketId200ResponseDataAttributes) HasCheckoutUrl() bool`

HasCheckoutUrl returns a boolean if a field has been set.

### SetCheckoutUrlNil

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetCheckoutUrlNil(b bool)`

 SetCheckoutUrlNil sets the value for CheckoutUrl to be an explicit nil

### UnsetCheckoutUrl
`func (o *GETMarketsMarketId200ResponseDataAttributes) UnsetCheckoutUrl()`

UnsetCheckoutUrl ensures that no value is present for CheckoutUrl, not even an explicit nil
### GetExternalPricesUrl

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetExternalPricesUrl() interface{}`

GetExternalPricesUrl returns the ExternalPricesUrl field if non-nil, zero value otherwise.

### GetExternalPricesUrlOk

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetExternalPricesUrlOk() (*interface{}, bool)`

GetExternalPricesUrlOk returns a tuple with the ExternalPricesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPricesUrl

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetExternalPricesUrl(v interface{})`

SetExternalPricesUrl sets ExternalPricesUrl field to given value.

### HasExternalPricesUrl

`func (o *GETMarketsMarketId200ResponseDataAttributes) HasExternalPricesUrl() bool`

HasExternalPricesUrl returns a boolean if a field has been set.

### SetExternalPricesUrlNil

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetExternalPricesUrlNil(b bool)`

 SetExternalPricesUrlNil sets the value for ExternalPricesUrl to be an explicit nil

### UnsetExternalPricesUrl
`func (o *GETMarketsMarketId200ResponseDataAttributes) UnsetExternalPricesUrl()`

UnsetExternalPricesUrl ensures that no value is present for ExternalPricesUrl, not even an explicit nil
### GetExternalOrderValidationUrl

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetExternalOrderValidationUrl() interface{}`

GetExternalOrderValidationUrl returns the ExternalOrderValidationUrl field if non-nil, zero value otherwise.

### GetExternalOrderValidationUrlOk

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetExternalOrderValidationUrlOk() (*interface{}, bool)`

GetExternalOrderValidationUrlOk returns a tuple with the ExternalOrderValidationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOrderValidationUrl

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetExternalOrderValidationUrl(v interface{})`

SetExternalOrderValidationUrl sets ExternalOrderValidationUrl field to given value.

### HasExternalOrderValidationUrl

`func (o *GETMarketsMarketId200ResponseDataAttributes) HasExternalOrderValidationUrl() bool`

HasExternalOrderValidationUrl returns a boolean if a field has been set.

### SetExternalOrderValidationUrlNil

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetExternalOrderValidationUrlNil(b bool)`

 SetExternalOrderValidationUrlNil sets the value for ExternalOrderValidationUrl to be an explicit nil

### UnsetExternalOrderValidationUrl
`func (o *GETMarketsMarketId200ResponseDataAttributes) UnsetExternalOrderValidationUrl()`

UnsetExternalOrderValidationUrl ensures that no value is present for ExternalOrderValidationUrl, not even an explicit nil
### GetSharedSecret

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetSharedSecret() interface{}`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetSharedSecretOk() (*interface{}, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetSharedSecret(v interface{})`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *GETMarketsMarketId200ResponseDataAttributes) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### SetSharedSecretNil

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetSharedSecretNil(b bool)`

 SetSharedSecretNil sets the value for SharedSecret to be an explicit nil

### UnsetSharedSecret
`func (o *GETMarketsMarketId200ResponseDataAttributes) UnsetSharedSecret()`

UnsetSharedSecret ensures that no value is present for SharedSecret, not even an explicit nil
### GetPrivate

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetPrivate() interface{}`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetPrivateOk() (*interface{}, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetPrivate(v interface{})`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *GETMarketsMarketId200ResponseDataAttributes) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### SetPrivateNil

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetPrivateNil(b bool)`

 SetPrivateNil sets the value for Private to be an explicit nil

### UnsetPrivate
`func (o *GETMarketsMarketId200ResponseDataAttributes) UnsetPrivate()`

UnsetPrivate ensures that no value is present for Private, not even an explicit nil
### GetDisabledAt

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetDisabledAt() interface{}`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetDisabledAtOk() (*interface{}, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetDisabledAt(v interface{})`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *GETMarketsMarketId200ResponseDataAttributes) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### SetDisabledAtNil

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetDisabledAtNil(b bool)`

 SetDisabledAtNil sets the value for DisabledAt to be an explicit nil

### UnsetDisabledAt
`func (o *GETMarketsMarketId200ResponseDataAttributes) UnsetDisabledAt()`

UnsetDisabledAt ensures that no value is present for DisabledAt, not even an explicit nil
### GetCreatedAt

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETMarketsMarketId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETMarketsMarketId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETMarketsMarketId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETMarketsMarketId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETMarketsMarketId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETMarketsMarketId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETMarketsMarketId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETMarketsMarketId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETMarketsMarketId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETMarketsMarketId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETMarketsMarketId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETMarketsMarketId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


