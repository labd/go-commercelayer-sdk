# POSTShippingMethods201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The shipping method&#39;s name | 
**Scheme** | Pointer to **interface{}** | The shipping method&#39;s scheme, one of &#39;flat&#39; or &#39;weight_tiered&#39;. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**PriceAmountCents** | **interface{}** | The price of this shipping method, in cents. | 
**FreeOverAmountCents** | Pointer to **interface{}** | Apply free shipping if the order amount is over this value, in cents. | [optional] 
**MinWeight** | Pointer to **interface{}** | The minimum weight for which this shipping method is available. | [optional] 
**MaxWeight** | Pointer to **interface{}** | The maximum weight for which this shipping method is available. | [optional] 
**UnitOfWeight** | Pointer to **interface{}** | Can be one of &#39;gr&#39;, &#39;lb&#39;, or &#39;oz&#39; | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTShippingMethods201ResponseDataAttributes

`func NewPOSTShippingMethods201ResponseDataAttributes(name interface{}, priceAmountCents interface{}, ) *POSTShippingMethods201ResponseDataAttributes`

NewPOSTShippingMethods201ResponseDataAttributes instantiates a new POSTShippingMethods201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTShippingMethods201ResponseDataAttributesWithDefaults

`func NewPOSTShippingMethods201ResponseDataAttributesWithDefaults() *POSTShippingMethods201ResponseDataAttributes`

NewPOSTShippingMethods201ResponseDataAttributesWithDefaults instantiates a new POSTShippingMethods201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTShippingMethods201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTShippingMethods201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTShippingMethods201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTShippingMethods201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTShippingMethods201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetScheme

`func (o *POSTShippingMethods201ResponseDataAttributes) GetScheme() interface{}`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *POSTShippingMethods201ResponseDataAttributes) GetSchemeOk() (*interface{}, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *POSTShippingMethods201ResponseDataAttributes) SetScheme(v interface{})`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *POSTShippingMethods201ResponseDataAttributes) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *POSTShippingMethods201ResponseDataAttributes) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *POSTShippingMethods201ResponseDataAttributes) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetCurrencyCode

`func (o *POSTShippingMethods201ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *POSTShippingMethods201ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *POSTShippingMethods201ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *POSTShippingMethods201ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *POSTShippingMethods201ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *POSTShippingMethods201ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetPriceAmountCents

`func (o *POSTShippingMethods201ResponseDataAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *POSTShippingMethods201ResponseDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *POSTShippingMethods201ResponseDataAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.


### SetPriceAmountCentsNil

`func (o *POSTShippingMethods201ResponseDataAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *POSTShippingMethods201ResponseDataAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetFreeOverAmountCents

`func (o *POSTShippingMethods201ResponseDataAttributes) GetFreeOverAmountCents() interface{}`

GetFreeOverAmountCents returns the FreeOverAmountCents field if non-nil, zero value otherwise.

### GetFreeOverAmountCentsOk

`func (o *POSTShippingMethods201ResponseDataAttributes) GetFreeOverAmountCentsOk() (*interface{}, bool)`

GetFreeOverAmountCentsOk returns a tuple with the FreeOverAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeOverAmountCents

`func (o *POSTShippingMethods201ResponseDataAttributes) SetFreeOverAmountCents(v interface{})`

SetFreeOverAmountCents sets FreeOverAmountCents field to given value.

### HasFreeOverAmountCents

`func (o *POSTShippingMethods201ResponseDataAttributes) HasFreeOverAmountCents() bool`

HasFreeOverAmountCents returns a boolean if a field has been set.

### SetFreeOverAmountCentsNil

`func (o *POSTShippingMethods201ResponseDataAttributes) SetFreeOverAmountCentsNil(b bool)`

 SetFreeOverAmountCentsNil sets the value for FreeOverAmountCents to be an explicit nil

### UnsetFreeOverAmountCents
`func (o *POSTShippingMethods201ResponseDataAttributes) UnsetFreeOverAmountCents()`

UnsetFreeOverAmountCents ensures that no value is present for FreeOverAmountCents, not even an explicit nil
### GetMinWeight

`func (o *POSTShippingMethods201ResponseDataAttributes) GetMinWeight() interface{}`

GetMinWeight returns the MinWeight field if non-nil, zero value otherwise.

### GetMinWeightOk

`func (o *POSTShippingMethods201ResponseDataAttributes) GetMinWeightOk() (*interface{}, bool)`

GetMinWeightOk returns a tuple with the MinWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWeight

`func (o *POSTShippingMethods201ResponseDataAttributes) SetMinWeight(v interface{})`

SetMinWeight sets MinWeight field to given value.

### HasMinWeight

`func (o *POSTShippingMethods201ResponseDataAttributes) HasMinWeight() bool`

HasMinWeight returns a boolean if a field has been set.

### SetMinWeightNil

`func (o *POSTShippingMethods201ResponseDataAttributes) SetMinWeightNil(b bool)`

 SetMinWeightNil sets the value for MinWeight to be an explicit nil

### UnsetMinWeight
`func (o *POSTShippingMethods201ResponseDataAttributes) UnsetMinWeight()`

UnsetMinWeight ensures that no value is present for MinWeight, not even an explicit nil
### GetMaxWeight

`func (o *POSTShippingMethods201ResponseDataAttributes) GetMaxWeight() interface{}`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *POSTShippingMethods201ResponseDataAttributes) GetMaxWeightOk() (*interface{}, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *POSTShippingMethods201ResponseDataAttributes) SetMaxWeight(v interface{})`

SetMaxWeight sets MaxWeight field to given value.

### HasMaxWeight

`func (o *POSTShippingMethods201ResponseDataAttributes) HasMaxWeight() bool`

HasMaxWeight returns a boolean if a field has been set.

### SetMaxWeightNil

`func (o *POSTShippingMethods201ResponseDataAttributes) SetMaxWeightNil(b bool)`

 SetMaxWeightNil sets the value for MaxWeight to be an explicit nil

### UnsetMaxWeight
`func (o *POSTShippingMethods201ResponseDataAttributes) UnsetMaxWeight()`

UnsetMaxWeight ensures that no value is present for MaxWeight, not even an explicit nil
### GetUnitOfWeight

`func (o *POSTShippingMethods201ResponseDataAttributes) GetUnitOfWeight() interface{}`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *POSTShippingMethods201ResponseDataAttributes) GetUnitOfWeightOk() (*interface{}, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *POSTShippingMethods201ResponseDataAttributes) SetUnitOfWeight(v interface{})`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *POSTShippingMethods201ResponseDataAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### SetUnitOfWeightNil

`func (o *POSTShippingMethods201ResponseDataAttributes) SetUnitOfWeightNil(b bool)`

 SetUnitOfWeightNil sets the value for UnitOfWeight to be an explicit nil

### UnsetUnitOfWeight
`func (o *POSTShippingMethods201ResponseDataAttributes) UnsetUnitOfWeight()`

UnsetUnitOfWeight ensures that no value is present for UnitOfWeight, not even an explicit nil
### GetReference

`func (o *POSTShippingMethods201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTShippingMethods201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTShippingMethods201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTShippingMethods201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTShippingMethods201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTShippingMethods201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTShippingMethods201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTShippingMethods201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTShippingMethods201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTShippingMethods201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTShippingMethods201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTShippingMethods201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTShippingMethods201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTShippingMethods201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTShippingMethods201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTShippingMethods201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTShippingMethods201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTShippingMethods201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


