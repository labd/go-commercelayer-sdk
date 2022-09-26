# ShippingMethodTierDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The shipping method tier&#39;s name | [optional] 
**UpTo** | Pointer to **float32** | The tier upper limit. When &#39;null&#39; it means infinity (useful to have an always matching tier). | [optional] 
**PriceAmountCents** | Pointer to **int32** | The price of this shipping method tier, in cents. | [optional] 
**PriceAmountFloat** | Pointer to **float32** | The price of this shipping method tier, float. | [optional] 
**FormattedPriceAmount** | Pointer to **string** | The price of this shipping method tier, formatted. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewShippingMethodTierDataAttributes

`func NewShippingMethodTierDataAttributes() *ShippingMethodTierDataAttributes`

NewShippingMethodTierDataAttributes instantiates a new ShippingMethodTierDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingMethodTierDataAttributesWithDefaults

`func NewShippingMethodTierDataAttributesWithDefaults() *ShippingMethodTierDataAttributes`

NewShippingMethodTierDataAttributesWithDefaults instantiates a new ShippingMethodTierDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ShippingMethodTierDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShippingMethodTierDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShippingMethodTierDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShippingMethodTierDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpTo

`func (o *ShippingMethodTierDataAttributes) GetUpTo() float32`

GetUpTo returns the UpTo field if non-nil, zero value otherwise.

### GetUpToOk

`func (o *ShippingMethodTierDataAttributes) GetUpToOk() (*float32, bool)`

GetUpToOk returns a tuple with the UpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTo

`func (o *ShippingMethodTierDataAttributes) SetUpTo(v float32)`

SetUpTo sets UpTo field to given value.

### HasUpTo

`func (o *ShippingMethodTierDataAttributes) HasUpTo() bool`

HasUpTo returns a boolean if a field has been set.

### GetPriceAmountCents

`func (o *ShippingMethodTierDataAttributes) GetPriceAmountCents() int32`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *ShippingMethodTierDataAttributes) GetPriceAmountCentsOk() (*int32, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *ShippingMethodTierDataAttributes) SetPriceAmountCents(v int32)`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *ShippingMethodTierDataAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### GetPriceAmountFloat

`func (o *ShippingMethodTierDataAttributes) GetPriceAmountFloat() float32`

GetPriceAmountFloat returns the PriceAmountFloat field if non-nil, zero value otherwise.

### GetPriceAmountFloatOk

`func (o *ShippingMethodTierDataAttributes) GetPriceAmountFloatOk() (*float32, bool)`

GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountFloat

`func (o *ShippingMethodTierDataAttributes) SetPriceAmountFloat(v float32)`

SetPriceAmountFloat sets PriceAmountFloat field to given value.

### HasPriceAmountFloat

`func (o *ShippingMethodTierDataAttributes) HasPriceAmountFloat() bool`

HasPriceAmountFloat returns a boolean if a field has been set.

### GetFormattedPriceAmount

`func (o *ShippingMethodTierDataAttributes) GetFormattedPriceAmount() string`

GetFormattedPriceAmount returns the FormattedPriceAmount field if non-nil, zero value otherwise.

### GetFormattedPriceAmountOk

`func (o *ShippingMethodTierDataAttributes) GetFormattedPriceAmountOk() (*string, bool)`

GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmount

`func (o *ShippingMethodTierDataAttributes) SetFormattedPriceAmount(v string)`

SetFormattedPriceAmount sets FormattedPriceAmount field to given value.

### HasFormattedPriceAmount

`func (o *ShippingMethodTierDataAttributes) HasFormattedPriceAmount() bool`

HasFormattedPriceAmount returns a boolean if a field has been set.

### GetId

`func (o *ShippingMethodTierDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingMethodTierDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingMethodTierDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShippingMethodTierDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ShippingMethodTierDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ShippingMethodTierDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ShippingMethodTierDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ShippingMethodTierDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ShippingMethodTierDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ShippingMethodTierDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ShippingMethodTierDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ShippingMethodTierDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *ShippingMethodTierDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ShippingMethodTierDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ShippingMethodTierDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ShippingMethodTierDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *ShippingMethodTierDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *ShippingMethodTierDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *ShippingMethodTierDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *ShippingMethodTierDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *ShippingMethodTierDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShippingMethodTierDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShippingMethodTierDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ShippingMethodTierDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


