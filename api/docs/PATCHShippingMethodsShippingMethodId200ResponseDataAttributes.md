# PATCHShippingMethodsShippingMethodId200ResponseDataAttributes

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

### NewPATCHShippingMethodsShippingMethodId200ResponseDataAttributes

`func NewPATCHShippingMethodsShippingMethodId200ResponseDataAttributes() *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes`

NewPATCHShippingMethodsShippingMethodId200ResponseDataAttributes instantiates a new PATCHShippingMethodsShippingMethodId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHShippingMethodsShippingMethodId200ResponseDataAttributesWithDefaults

`func NewPATCHShippingMethodsShippingMethodId200ResponseDataAttributesWithDefaults() *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes`

NewPATCHShippingMethodsShippingMethodId200ResponseDataAttributesWithDefaults instantiates a new PATCHShippingMethodsShippingMethodId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheme

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDisable

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEnable

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetPriceAmountCents

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetPriceAmountCents() int32`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetPriceAmountCentsOk() (*int32, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetPriceAmountCents(v int32)`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### GetFreeOverAmountCents

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetFreeOverAmountCents() int32`

GetFreeOverAmountCents returns the FreeOverAmountCents field if non-nil, zero value otherwise.

### GetFreeOverAmountCentsOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetFreeOverAmountCentsOk() (*int32, bool)`

GetFreeOverAmountCentsOk returns a tuple with the FreeOverAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeOverAmountCents

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetFreeOverAmountCents(v int32)`

SetFreeOverAmountCents sets FreeOverAmountCents field to given value.

### HasFreeOverAmountCents

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasFreeOverAmountCents() bool`

HasFreeOverAmountCents returns a boolean if a field has been set.

### GetMinWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMinWeight() float32`

GetMinWeight returns the MinWeight field if non-nil, zero value otherwise.

### GetMinWeightOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMinWeightOk() (*float32, bool)`

GetMinWeightOk returns a tuple with the MinWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetMinWeight(v float32)`

SetMinWeight sets MinWeight field to given value.

### HasMinWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasMinWeight() bool`

HasMinWeight returns a boolean if a field has been set.

### GetMaxWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMaxWeight() float32`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMaxWeightOk() (*float32, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetMaxWeight(v float32)`

SetMaxWeight sets MaxWeight field to given value.

### HasMaxWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasMaxWeight() bool`

HasMaxWeight returns a boolean if a field has been set.

### GetUnitOfWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetUnitOfWeight() string`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetUnitOfWeightOk() (*string, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetUnitOfWeight(v string)`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### GetReference

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


