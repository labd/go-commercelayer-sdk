# PATCHShippingMethodsShippingMethodIdRequestDataAttributes

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

### NewPATCHShippingMethodsShippingMethodIdRequestDataAttributes

`func NewPATCHShippingMethodsShippingMethodIdRequestDataAttributes() *PATCHShippingMethodsShippingMethodIdRequestDataAttributes`

NewPATCHShippingMethodsShippingMethodIdRequestDataAttributes instantiates a new PATCHShippingMethodsShippingMethodIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHShippingMethodsShippingMethodIdRequestDataAttributesWithDefaults

`func NewPATCHShippingMethodsShippingMethodIdRequestDataAttributesWithDefaults() *PATCHShippingMethodsShippingMethodIdRequestDataAttributes`

NewPATCHShippingMethodsShippingMethodIdRequestDataAttributesWithDefaults instantiates a new PATCHShippingMethodsShippingMethodIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetScheme

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetScheme() interface{}`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetSchemeOk() (*interface{}, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetScheme(v interface{})`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetCurrencyCode

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetDisable

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetDisable() interface{}`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetDisableOk() (*interface{}, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetDisable(v interface{})`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### SetDisableNil

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetDisableNil(b bool)`

 SetDisableNil sets the value for Disable to be an explicit nil

### UnsetDisable
`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) UnsetDisable()`

UnsetDisable ensures that no value is present for Disable, not even an explicit nil
### GetEnable

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetEnable() interface{}`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetEnableOk() (*interface{}, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetEnable(v interface{})`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetPriceAmountCents

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### SetPriceAmountCentsNil

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetFreeOverAmountCents

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetFreeOverAmountCents() interface{}`

GetFreeOverAmountCents returns the FreeOverAmountCents field if non-nil, zero value otherwise.

### GetFreeOverAmountCentsOk

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetFreeOverAmountCentsOk() (*interface{}, bool)`

GetFreeOverAmountCentsOk returns a tuple with the FreeOverAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeOverAmountCents

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetFreeOverAmountCents(v interface{})`

SetFreeOverAmountCents sets FreeOverAmountCents field to given value.

### HasFreeOverAmountCents

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) HasFreeOverAmountCents() bool`

HasFreeOverAmountCents returns a boolean if a field has been set.

### SetFreeOverAmountCentsNil

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetFreeOverAmountCentsNil(b bool)`

 SetFreeOverAmountCentsNil sets the value for FreeOverAmountCents to be an explicit nil

### UnsetFreeOverAmountCents
`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) UnsetFreeOverAmountCents()`

UnsetFreeOverAmountCents ensures that no value is present for FreeOverAmountCents, not even an explicit nil
### GetMinWeight

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetMinWeight() interface{}`

GetMinWeight returns the MinWeight field if non-nil, zero value otherwise.

### GetMinWeightOk

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetMinWeightOk() (*interface{}, bool)`

GetMinWeightOk returns a tuple with the MinWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWeight

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetMinWeight(v interface{})`

SetMinWeight sets MinWeight field to given value.

### HasMinWeight

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) HasMinWeight() bool`

HasMinWeight returns a boolean if a field has been set.

### SetMinWeightNil

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetMinWeightNil(b bool)`

 SetMinWeightNil sets the value for MinWeight to be an explicit nil

### UnsetMinWeight
`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) UnsetMinWeight()`

UnsetMinWeight ensures that no value is present for MinWeight, not even an explicit nil
### GetMaxWeight

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetMaxWeight() interface{}`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetMaxWeightOk() (*interface{}, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetMaxWeight(v interface{})`

SetMaxWeight sets MaxWeight field to given value.

### HasMaxWeight

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) HasMaxWeight() bool`

HasMaxWeight returns a boolean if a field has been set.

### SetMaxWeightNil

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetMaxWeightNil(b bool)`

 SetMaxWeightNil sets the value for MaxWeight to be an explicit nil

### UnsetMaxWeight
`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) UnsetMaxWeight()`

UnsetMaxWeight ensures that no value is present for MaxWeight, not even an explicit nil
### GetUnitOfWeight

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetUnitOfWeight() interface{}`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetUnitOfWeightOk() (*interface{}, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetUnitOfWeight(v interface{})`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### SetUnitOfWeightNil

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetUnitOfWeightNil(b bool)`

 SetUnitOfWeightNil sets the value for UnitOfWeight to be an explicit nil

### UnsetUnitOfWeight
`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) UnsetUnitOfWeight()`

UnsetUnitOfWeight ensures that no value is present for UnitOfWeight, not even an explicit nil
### GetReference

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHShippingMethodsShippingMethodIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


