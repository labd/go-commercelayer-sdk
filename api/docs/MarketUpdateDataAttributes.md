# MarketUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The market&#39;s internal name | [optional] 
**FacebookPixelId** | Pointer to **string** | The Facebook Pixed ID | [optional] 
**CheckoutUrl** | Pointer to **string** | The checkout URL for this market | [optional] 
**ExternalPricesUrl** | Pointer to **string** | The URL used to fetch prices from an external source | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewMarketUpdateDataAttributes

`func NewMarketUpdateDataAttributes() *MarketUpdateDataAttributes`

NewMarketUpdateDataAttributes instantiates a new MarketUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketUpdateDataAttributesWithDefaults

`func NewMarketUpdateDataAttributesWithDefaults() *MarketUpdateDataAttributes`

NewMarketUpdateDataAttributesWithDefaults instantiates a new MarketUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MarketUpdateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketUpdateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketUpdateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketUpdateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFacebookPixelId

`func (o *MarketUpdateDataAttributes) GetFacebookPixelId() string`

GetFacebookPixelId returns the FacebookPixelId field if non-nil, zero value otherwise.

### GetFacebookPixelIdOk

`func (o *MarketUpdateDataAttributes) GetFacebookPixelIdOk() (*string, bool)`

GetFacebookPixelIdOk returns a tuple with the FacebookPixelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookPixelId

`func (o *MarketUpdateDataAttributes) SetFacebookPixelId(v string)`

SetFacebookPixelId sets FacebookPixelId field to given value.

### HasFacebookPixelId

`func (o *MarketUpdateDataAttributes) HasFacebookPixelId() bool`

HasFacebookPixelId returns a boolean if a field has been set.

### GetCheckoutUrl

`func (o *MarketUpdateDataAttributes) GetCheckoutUrl() string`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *MarketUpdateDataAttributes) GetCheckoutUrlOk() (*string, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *MarketUpdateDataAttributes) SetCheckoutUrl(v string)`

SetCheckoutUrl sets CheckoutUrl field to given value.

### HasCheckoutUrl

`func (o *MarketUpdateDataAttributes) HasCheckoutUrl() bool`

HasCheckoutUrl returns a boolean if a field has been set.

### GetExternalPricesUrl

`func (o *MarketUpdateDataAttributes) GetExternalPricesUrl() string`

GetExternalPricesUrl returns the ExternalPricesUrl field if non-nil, zero value otherwise.

### GetExternalPricesUrlOk

`func (o *MarketUpdateDataAttributes) GetExternalPricesUrlOk() (*string, bool)`

GetExternalPricesUrlOk returns a tuple with the ExternalPricesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPricesUrl

`func (o *MarketUpdateDataAttributes) SetExternalPricesUrl(v string)`

SetExternalPricesUrl sets ExternalPricesUrl field to given value.

### HasExternalPricesUrl

`func (o *MarketUpdateDataAttributes) HasExternalPricesUrl() bool`

HasExternalPricesUrl returns a boolean if a field has been set.

### GetReference

`func (o *MarketUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *MarketUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *MarketUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *MarketUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *MarketUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *MarketUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *MarketUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *MarketUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *MarketUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MarketUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MarketUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MarketUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


