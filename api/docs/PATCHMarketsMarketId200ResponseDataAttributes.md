# PATCHMarketsMarketId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The market&#39;s internal name | [optional] 
**FacebookPixelId** | Pointer to **string** | The Facebook Pixed ID | [optional] 
**CheckoutUrl** | Pointer to **string** | The checkout URL for this market | [optional] 
**ExternalPricesUrl** | Pointer to **string** | The URL used to overwrite prices by an external source. | [optional] 
**ExternalOrderValidationUrl** | Pointer to **string** | The URL used to validate orders by an external source. | [optional] 
**Disable** | Pointer to **bool** | Send this attribute if you want to mark the market as disabled. | [optional] 
**Enable** | Pointer to **bool** | Send this attribute if you want to mark the market as enabled. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHMarketsMarketId200ResponseDataAttributes

`func NewPATCHMarketsMarketId200ResponseDataAttributes() *PATCHMarketsMarketId200ResponseDataAttributes`

NewPATCHMarketsMarketId200ResponseDataAttributes instantiates a new PATCHMarketsMarketId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHMarketsMarketId200ResponseDataAttributesWithDefaults

`func NewPATCHMarketsMarketId200ResponseDataAttributesWithDefaults() *PATCHMarketsMarketId200ResponseDataAttributes`

NewPATCHMarketsMarketId200ResponseDataAttributesWithDefaults instantiates a new PATCHMarketsMarketId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFacebookPixelId

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetFacebookPixelId() string`

GetFacebookPixelId returns the FacebookPixelId field if non-nil, zero value otherwise.

### GetFacebookPixelIdOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetFacebookPixelIdOk() (*string, bool)`

GetFacebookPixelIdOk returns a tuple with the FacebookPixelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookPixelId

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetFacebookPixelId(v string)`

SetFacebookPixelId sets FacebookPixelId field to given value.

### HasFacebookPixelId

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasFacebookPixelId() bool`

HasFacebookPixelId returns a boolean if a field has been set.

### GetCheckoutUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetCheckoutUrl() string`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetCheckoutUrlOk() (*string, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetCheckoutUrl(v string)`

SetCheckoutUrl sets CheckoutUrl field to given value.

### HasCheckoutUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasCheckoutUrl() bool`

HasCheckoutUrl returns a boolean if a field has been set.

### GetExternalPricesUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetExternalPricesUrl() string`

GetExternalPricesUrl returns the ExternalPricesUrl field if non-nil, zero value otherwise.

### GetExternalPricesUrlOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetExternalPricesUrlOk() (*string, bool)`

GetExternalPricesUrlOk returns a tuple with the ExternalPricesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPricesUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetExternalPricesUrl(v string)`

SetExternalPricesUrl sets ExternalPricesUrl field to given value.

### HasExternalPricesUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasExternalPricesUrl() bool`

HasExternalPricesUrl returns a boolean if a field has been set.

### GetExternalOrderValidationUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetExternalOrderValidationUrl() string`

GetExternalOrderValidationUrl returns the ExternalOrderValidationUrl field if non-nil, zero value otherwise.

### GetExternalOrderValidationUrlOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetExternalOrderValidationUrlOk() (*string, bool)`

GetExternalOrderValidationUrlOk returns a tuple with the ExternalOrderValidationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOrderValidationUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetExternalOrderValidationUrl(v string)`

SetExternalOrderValidationUrl sets ExternalOrderValidationUrl field to given value.

### HasExternalOrderValidationUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasExternalOrderValidationUrl() bool`

HasExternalOrderValidationUrl returns a boolean if a field has been set.

### GetDisable

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEnable

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetReference

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


