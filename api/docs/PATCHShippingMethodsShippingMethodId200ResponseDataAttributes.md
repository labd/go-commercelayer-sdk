# PATCHShippingMethodsShippingMethodId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The shipping method&#39;s name | [optional] 
**Scheme** | Pointer to **interface{}** | The shipping method&#39;s scheme, one of &#39;flat&#39; or &#39;weight_tiered&#39;. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Disable** | Pointer to **interface{}** | Send this attribute if you want to mark the shipping method as disabled. | [optional] 
**Enable** | Pointer to **interface{}** | Send this attribute if you want to mark the shipping method as enabled. | [optional] 
**PriceAmountCents** | Pointer to **interface{}** | The price of this shipping method, in cents. | [optional] 
**FreeOverAmountCents** | Pointer to **interface{}** | Apply free shipping if the order amount is over this value, in cents. | [optional] 
**MinWeight** | Pointer to **interface{}** | The minimum weight for which this shipping method is available. | [optional] 
**MaxWeight** | Pointer to **interface{}** | The maximum weight for which this shipping method is available. | [optional] 
**UnitOfWeight** | Pointer to **interface{}** | Can be one of &#39;gr&#39;, &#39;lb&#39;, or &#39;oz&#39; | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetScheme

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetScheme() interface{}`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetSchemeOk() (*interface{}, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetScheme(v interface{})`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetCurrencyCode

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetDisable

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetDisable() interface{}`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetDisableOk() (*interface{}, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetDisable(v interface{})`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### SetDisableNil

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetDisableNil(b bool)`

 SetDisableNil sets the value for Disable to be an explicit nil

### UnsetDisable
`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetDisable()`

UnsetDisable ensures that no value is present for Disable, not even an explicit nil
### GetEnable

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetEnable() interface{}`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetEnableOk() (*interface{}, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetEnable(v interface{})`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetPriceAmountCents

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### SetPriceAmountCentsNil

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetFreeOverAmountCents

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetFreeOverAmountCents() interface{}`

GetFreeOverAmountCents returns the FreeOverAmountCents field if non-nil, zero value otherwise.

### GetFreeOverAmountCentsOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetFreeOverAmountCentsOk() (*interface{}, bool)`

GetFreeOverAmountCentsOk returns a tuple with the FreeOverAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeOverAmountCents

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetFreeOverAmountCents(v interface{})`

SetFreeOverAmountCents sets FreeOverAmountCents field to given value.

### HasFreeOverAmountCents

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasFreeOverAmountCents() bool`

HasFreeOverAmountCents returns a boolean if a field has been set.

### SetFreeOverAmountCentsNil

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetFreeOverAmountCentsNil(b bool)`

 SetFreeOverAmountCentsNil sets the value for FreeOverAmountCents to be an explicit nil

### UnsetFreeOverAmountCents
`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetFreeOverAmountCents()`

UnsetFreeOverAmountCents ensures that no value is present for FreeOverAmountCents, not even an explicit nil
### GetMinWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMinWeight() interface{}`

GetMinWeight returns the MinWeight field if non-nil, zero value otherwise.

### GetMinWeightOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMinWeightOk() (*interface{}, bool)`

GetMinWeightOk returns a tuple with the MinWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetMinWeight(v interface{})`

SetMinWeight sets MinWeight field to given value.

### HasMinWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasMinWeight() bool`

HasMinWeight returns a boolean if a field has been set.

### SetMinWeightNil

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetMinWeightNil(b bool)`

 SetMinWeightNil sets the value for MinWeight to be an explicit nil

### UnsetMinWeight
`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetMinWeight()`

UnsetMinWeight ensures that no value is present for MinWeight, not even an explicit nil
### GetMaxWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMaxWeight() interface{}`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMaxWeightOk() (*interface{}, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetMaxWeight(v interface{})`

SetMaxWeight sets MaxWeight field to given value.

### HasMaxWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasMaxWeight() bool`

HasMaxWeight returns a boolean if a field has been set.

### SetMaxWeightNil

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetMaxWeightNil(b bool)`

 SetMaxWeightNil sets the value for MaxWeight to be an explicit nil

### UnsetMaxWeight
`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetMaxWeight()`

UnsetMaxWeight ensures that no value is present for MaxWeight, not even an explicit nil
### GetUnitOfWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetUnitOfWeight() interface{}`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetUnitOfWeightOk() (*interface{}, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetUnitOfWeight(v interface{})`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### SetUnitOfWeightNil

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetUnitOfWeightNil(b bool)`

 SetUnitOfWeightNil sets the value for UnitOfWeight to be an explicit nil

### UnsetUnitOfWeight
`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetUnitOfWeight()`

UnsetUnitOfWeight ensures that no value is present for UnitOfWeight, not even an explicit nil
### GetReference

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


