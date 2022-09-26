# ShippingMethodUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The shipping method&#39;s name | [optional] 
**Scheme** | Pointer to **string** | The shipping method&#39;s scheme, one of &#39;flat&#39; or &#39;weight_tiered&#39;. | [optional] 
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Disable** | Pointer to **bool** | Send this attribute if you want to mark the shipping method as disabled. | [optional] 
**Enable** | Pointer to **bool** | Send this attribute if you want to mark the shipping method as enabled. | [optional] 
**PriceAmountCents** | Pointer to **int32** | The price of this shipping method, in cents. | [optional] 
**FreeOverAmountCents** | Pointer to **int32** | Apply free shipping if the order amount is over this value, in cents. | [optional] 
**MinWeight** | Pointer to **float32** | The minimum weight for which this shipping method is available. | [optional] 
**MaxWeight** | Pointer to **float32** | The maximum weight for which this shipping method is available. | [optional] 
**UnitOfWeight** | Pointer to **string** | Can be one of &#39;gr&#39;, &#39;lb&#39;, or &#39;oz&#39; | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewShippingMethodUpdateDataAttributes

`func NewShippingMethodUpdateDataAttributes() *ShippingMethodUpdateDataAttributes`

NewShippingMethodUpdateDataAttributes instantiates a new ShippingMethodUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingMethodUpdateDataAttributesWithDefaults

`func NewShippingMethodUpdateDataAttributesWithDefaults() *ShippingMethodUpdateDataAttributes`

NewShippingMethodUpdateDataAttributesWithDefaults instantiates a new ShippingMethodUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ShippingMethodUpdateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShippingMethodUpdateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShippingMethodUpdateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShippingMethodUpdateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheme

`func (o *ShippingMethodUpdateDataAttributes) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ShippingMethodUpdateDataAttributes) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ShippingMethodUpdateDataAttributes) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *ShippingMethodUpdateDataAttributes) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ShippingMethodUpdateDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ShippingMethodUpdateDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ShippingMethodUpdateDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ShippingMethodUpdateDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDisable

`func (o *ShippingMethodUpdateDataAttributes) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *ShippingMethodUpdateDataAttributes) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *ShippingMethodUpdateDataAttributes) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *ShippingMethodUpdateDataAttributes) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEnable

`func (o *ShippingMethodUpdateDataAttributes) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ShippingMethodUpdateDataAttributes) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ShippingMethodUpdateDataAttributes) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ShippingMethodUpdateDataAttributes) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetPriceAmountCents

`func (o *ShippingMethodUpdateDataAttributes) GetPriceAmountCents() int32`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *ShippingMethodUpdateDataAttributes) GetPriceAmountCentsOk() (*int32, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *ShippingMethodUpdateDataAttributes) SetPriceAmountCents(v int32)`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *ShippingMethodUpdateDataAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### GetFreeOverAmountCents

`func (o *ShippingMethodUpdateDataAttributes) GetFreeOverAmountCents() int32`

GetFreeOverAmountCents returns the FreeOverAmountCents field if non-nil, zero value otherwise.

### GetFreeOverAmountCentsOk

`func (o *ShippingMethodUpdateDataAttributes) GetFreeOverAmountCentsOk() (*int32, bool)`

GetFreeOverAmountCentsOk returns a tuple with the FreeOverAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeOverAmountCents

`func (o *ShippingMethodUpdateDataAttributes) SetFreeOverAmountCents(v int32)`

SetFreeOverAmountCents sets FreeOverAmountCents field to given value.

### HasFreeOverAmountCents

`func (o *ShippingMethodUpdateDataAttributes) HasFreeOverAmountCents() bool`

HasFreeOverAmountCents returns a boolean if a field has been set.

### GetMinWeight

`func (o *ShippingMethodUpdateDataAttributes) GetMinWeight() float32`

GetMinWeight returns the MinWeight field if non-nil, zero value otherwise.

### GetMinWeightOk

`func (o *ShippingMethodUpdateDataAttributes) GetMinWeightOk() (*float32, bool)`

GetMinWeightOk returns a tuple with the MinWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWeight

`func (o *ShippingMethodUpdateDataAttributes) SetMinWeight(v float32)`

SetMinWeight sets MinWeight field to given value.

### HasMinWeight

`func (o *ShippingMethodUpdateDataAttributes) HasMinWeight() bool`

HasMinWeight returns a boolean if a field has been set.

### GetMaxWeight

`func (o *ShippingMethodUpdateDataAttributes) GetMaxWeight() float32`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *ShippingMethodUpdateDataAttributes) GetMaxWeightOk() (*float32, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *ShippingMethodUpdateDataAttributes) SetMaxWeight(v float32)`

SetMaxWeight sets MaxWeight field to given value.

### HasMaxWeight

`func (o *ShippingMethodUpdateDataAttributes) HasMaxWeight() bool`

HasMaxWeight returns a boolean if a field has been set.

### GetUnitOfWeight

`func (o *ShippingMethodUpdateDataAttributes) GetUnitOfWeight() string`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *ShippingMethodUpdateDataAttributes) GetUnitOfWeightOk() (*string, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *ShippingMethodUpdateDataAttributes) SetUnitOfWeight(v string)`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *ShippingMethodUpdateDataAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### GetReference

`func (o *ShippingMethodUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ShippingMethodUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ShippingMethodUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ShippingMethodUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *ShippingMethodUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *ShippingMethodUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *ShippingMethodUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *ShippingMethodUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *ShippingMethodUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShippingMethodUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShippingMethodUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ShippingMethodUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


