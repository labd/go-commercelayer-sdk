# MarketDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | Unique identifier for the market (numeric) | [optional] 
**Name** | Pointer to **string** | The market&#39;s internal name | [optional] 
**FacebookPixelId** | Pointer to **string** | The Facebook Pixed ID | [optional] 
**CheckoutUrl** | Pointer to **string** | The checkout URL for this market | [optional] 
**ExternalPricesUrl** | Pointer to **string** | The URL used to fetch prices from an external source | [optional] 
**Private** | Pointer to **bool** | Indicates if market belongs to a customer_group. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewMarketDataAttributes

`func NewMarketDataAttributes() *MarketDataAttributes`

NewMarketDataAttributes instantiates a new MarketDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketDataAttributesWithDefaults

`func NewMarketDataAttributesWithDefaults() *MarketDataAttributes`

NewMarketDataAttributesWithDefaults instantiates a new MarketDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *MarketDataAttributes) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *MarketDataAttributes) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *MarketDataAttributes) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *MarketDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetName

`func (o *MarketDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFacebookPixelId

`func (o *MarketDataAttributes) GetFacebookPixelId() string`

GetFacebookPixelId returns the FacebookPixelId field if non-nil, zero value otherwise.

### GetFacebookPixelIdOk

`func (o *MarketDataAttributes) GetFacebookPixelIdOk() (*string, bool)`

GetFacebookPixelIdOk returns a tuple with the FacebookPixelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookPixelId

`func (o *MarketDataAttributes) SetFacebookPixelId(v string)`

SetFacebookPixelId sets FacebookPixelId field to given value.

### HasFacebookPixelId

`func (o *MarketDataAttributes) HasFacebookPixelId() bool`

HasFacebookPixelId returns a boolean if a field has been set.

### GetCheckoutUrl

`func (o *MarketDataAttributes) GetCheckoutUrl() string`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *MarketDataAttributes) GetCheckoutUrlOk() (*string, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *MarketDataAttributes) SetCheckoutUrl(v string)`

SetCheckoutUrl sets CheckoutUrl field to given value.

### HasCheckoutUrl

`func (o *MarketDataAttributes) HasCheckoutUrl() bool`

HasCheckoutUrl returns a boolean if a field has been set.

### GetExternalPricesUrl

`func (o *MarketDataAttributes) GetExternalPricesUrl() string`

GetExternalPricesUrl returns the ExternalPricesUrl field if non-nil, zero value otherwise.

### GetExternalPricesUrlOk

`func (o *MarketDataAttributes) GetExternalPricesUrlOk() (*string, bool)`

GetExternalPricesUrlOk returns a tuple with the ExternalPricesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPricesUrl

`func (o *MarketDataAttributes) SetExternalPricesUrl(v string)`

SetExternalPricesUrl sets ExternalPricesUrl field to given value.

### HasExternalPricesUrl

`func (o *MarketDataAttributes) HasExternalPricesUrl() bool`

HasExternalPricesUrl returns a boolean if a field has been set.

### GetPrivate

`func (o *MarketDataAttributes) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *MarketDataAttributes) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *MarketDataAttributes) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *MarketDataAttributes) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetId

`func (o *MarketDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MarketDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarketDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarketDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MarketDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MarketDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MarketDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MarketDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MarketDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *MarketDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *MarketDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *MarketDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *MarketDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *MarketDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *MarketDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *MarketDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *MarketDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *MarketDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MarketDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MarketDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MarketDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


