# GETShippingMethods200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The shipping method&#39;s name | [optional] 
**Scheme** | Pointer to **interface{}** | The shipping method&#39;s scheme, one of &#39;flat&#39; or &#39;weight_tiered&#39;. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**DisabledAt** | Pointer to **interface{}** | Time at which the shipping method was disabled. | [optional] 
**PriceAmountCents** | Pointer to **interface{}** | The price of this shipping method, in cents. | [optional] 
**PriceAmountFloat** | Pointer to **interface{}** | The price of this shipping method, float. | [optional] 
**FormattedPriceAmount** | Pointer to **interface{}** | The price of this shipping method, formatted. | [optional] 
**FreeOverAmountCents** | Pointer to **interface{}** | Apply free shipping if the order amount is over this value, in cents. | [optional] 
**FreeOverAmountFloat** | Pointer to **interface{}** | Apply free shipping if the order amount is over this value, float. | [optional] 
**FormattedFreeOverAmount** | Pointer to **interface{}** | Apply free shipping if the order amount is over this value, formatted. | [optional] 
**PriceAmountForShipmentCents** | Pointer to **interface{}** | The calculated price (zero or price amount) when associated to a shipment, in cents. | [optional] 
**PriceAmountForShipmentFloat** | Pointer to **interface{}** | The calculated price (zero or price amount) when associated to a shipment, float. | [optional] 
**FormattedPriceAmountForShipment** | Pointer to **interface{}** | The calculated price (zero or price amount) when associated to a shipment, formatted. | [optional] 
**MinWeight** | Pointer to **interface{}** | The minimum weight for which this shipping method is available. | [optional] 
**MaxWeight** | Pointer to **interface{}** | The maximum weight for which this shipping method is available. | [optional] 
**UnitOfWeight** | Pointer to **interface{}** | Can be one of &#39;gr&#39;, &#39;lb&#39;, or &#39;oz&#39; | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETShippingMethods200ResponseDataInnerAttributes

`func NewGETShippingMethods200ResponseDataInnerAttributes() *GETShippingMethods200ResponseDataInnerAttributes`

NewGETShippingMethods200ResponseDataInnerAttributes instantiates a new GETShippingMethods200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETShippingMethods200ResponseDataInnerAttributesWithDefaults

`func NewGETShippingMethods200ResponseDataInnerAttributesWithDefaults() *GETShippingMethods200ResponseDataInnerAttributes`

NewGETShippingMethods200ResponseDataInnerAttributesWithDefaults instantiates a new GETShippingMethods200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetScheme

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetScheme() interface{}`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetSchemeOk() (*interface{}, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetScheme(v interface{})`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetCurrencyCode

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetDisabledAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetDisabledAt() interface{}`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetDisabledAtOk() (*interface{}, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetDisabledAt(v interface{})`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### SetDisabledAtNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetDisabledAtNil(b bool)`

 SetDisabledAtNil sets the value for DisabledAt to be an explicit nil

### UnsetDisabledAt
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetDisabledAt()`

UnsetDisabledAt ensures that no value is present for DisabledAt, not even an explicit nil
### GetPriceAmountCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### SetPriceAmountCentsNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetPriceAmountFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetPriceAmountFloat() interface{}`

GetPriceAmountFloat returns the PriceAmountFloat field if non-nil, zero value otherwise.

### GetPriceAmountFloatOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetPriceAmountFloatOk() (*interface{}, bool)`

GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetPriceAmountFloat(v interface{})`

SetPriceAmountFloat sets PriceAmountFloat field to given value.

### HasPriceAmountFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasPriceAmountFloat() bool`

HasPriceAmountFloat returns a boolean if a field has been set.

### SetPriceAmountFloatNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetPriceAmountFloatNil(b bool)`

 SetPriceAmountFloatNil sets the value for PriceAmountFloat to be an explicit nil

### UnsetPriceAmountFloat
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetPriceAmountFloat()`

UnsetPriceAmountFloat ensures that no value is present for PriceAmountFloat, not even an explicit nil
### GetFormattedPriceAmount

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFormattedPriceAmount() interface{}`

GetFormattedPriceAmount returns the FormattedPriceAmount field if non-nil, zero value otherwise.

### GetFormattedPriceAmountOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFormattedPriceAmountOk() (*interface{}, bool)`

GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmount

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetFormattedPriceAmount(v interface{})`

SetFormattedPriceAmount sets FormattedPriceAmount field to given value.

### HasFormattedPriceAmount

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasFormattedPriceAmount() bool`

HasFormattedPriceAmount returns a boolean if a field has been set.

### SetFormattedPriceAmountNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetFormattedPriceAmountNil(b bool)`

 SetFormattedPriceAmountNil sets the value for FormattedPriceAmount to be an explicit nil

### UnsetFormattedPriceAmount
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetFormattedPriceAmount()`

UnsetFormattedPriceAmount ensures that no value is present for FormattedPriceAmount, not even an explicit nil
### GetFreeOverAmountCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFreeOverAmountCents() interface{}`

GetFreeOverAmountCents returns the FreeOverAmountCents field if non-nil, zero value otherwise.

### GetFreeOverAmountCentsOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFreeOverAmountCentsOk() (*interface{}, bool)`

GetFreeOverAmountCentsOk returns a tuple with the FreeOverAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeOverAmountCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetFreeOverAmountCents(v interface{})`

SetFreeOverAmountCents sets FreeOverAmountCents field to given value.

### HasFreeOverAmountCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasFreeOverAmountCents() bool`

HasFreeOverAmountCents returns a boolean if a field has been set.

### SetFreeOverAmountCentsNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetFreeOverAmountCentsNil(b bool)`

 SetFreeOverAmountCentsNil sets the value for FreeOverAmountCents to be an explicit nil

### UnsetFreeOverAmountCents
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetFreeOverAmountCents()`

UnsetFreeOverAmountCents ensures that no value is present for FreeOverAmountCents, not even an explicit nil
### GetFreeOverAmountFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFreeOverAmountFloat() interface{}`

GetFreeOverAmountFloat returns the FreeOverAmountFloat field if non-nil, zero value otherwise.

### GetFreeOverAmountFloatOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFreeOverAmountFloatOk() (*interface{}, bool)`

GetFreeOverAmountFloatOk returns a tuple with the FreeOverAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeOverAmountFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetFreeOverAmountFloat(v interface{})`

SetFreeOverAmountFloat sets FreeOverAmountFloat field to given value.

### HasFreeOverAmountFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasFreeOverAmountFloat() bool`

HasFreeOverAmountFloat returns a boolean if a field has been set.

### SetFreeOverAmountFloatNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetFreeOverAmountFloatNil(b bool)`

 SetFreeOverAmountFloatNil sets the value for FreeOverAmountFloat to be an explicit nil

### UnsetFreeOverAmountFloat
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetFreeOverAmountFloat()`

UnsetFreeOverAmountFloat ensures that no value is present for FreeOverAmountFloat, not even an explicit nil
### GetFormattedFreeOverAmount

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFormattedFreeOverAmount() interface{}`

GetFormattedFreeOverAmount returns the FormattedFreeOverAmount field if non-nil, zero value otherwise.

### GetFormattedFreeOverAmountOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFormattedFreeOverAmountOk() (*interface{}, bool)`

GetFormattedFreeOverAmountOk returns a tuple with the FormattedFreeOverAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedFreeOverAmount

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetFormattedFreeOverAmount(v interface{})`

SetFormattedFreeOverAmount sets FormattedFreeOverAmount field to given value.

### HasFormattedFreeOverAmount

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasFormattedFreeOverAmount() bool`

HasFormattedFreeOverAmount returns a boolean if a field has been set.

### SetFormattedFreeOverAmountNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetFormattedFreeOverAmountNil(b bool)`

 SetFormattedFreeOverAmountNil sets the value for FormattedFreeOverAmount to be an explicit nil

### UnsetFormattedFreeOverAmount
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetFormattedFreeOverAmount()`

UnsetFormattedFreeOverAmount ensures that no value is present for FormattedFreeOverAmount, not even an explicit nil
### GetPriceAmountForShipmentCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetPriceAmountForShipmentCents() interface{}`

GetPriceAmountForShipmentCents returns the PriceAmountForShipmentCents field if non-nil, zero value otherwise.

### GetPriceAmountForShipmentCentsOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetPriceAmountForShipmentCentsOk() (*interface{}, bool)`

GetPriceAmountForShipmentCentsOk returns a tuple with the PriceAmountForShipmentCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountForShipmentCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetPriceAmountForShipmentCents(v interface{})`

SetPriceAmountForShipmentCents sets PriceAmountForShipmentCents field to given value.

### HasPriceAmountForShipmentCents

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasPriceAmountForShipmentCents() bool`

HasPriceAmountForShipmentCents returns a boolean if a field has been set.

### SetPriceAmountForShipmentCentsNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetPriceAmountForShipmentCentsNil(b bool)`

 SetPriceAmountForShipmentCentsNil sets the value for PriceAmountForShipmentCents to be an explicit nil

### UnsetPriceAmountForShipmentCents
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetPriceAmountForShipmentCents()`

UnsetPriceAmountForShipmentCents ensures that no value is present for PriceAmountForShipmentCents, not even an explicit nil
### GetPriceAmountForShipmentFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetPriceAmountForShipmentFloat() interface{}`

GetPriceAmountForShipmentFloat returns the PriceAmountForShipmentFloat field if non-nil, zero value otherwise.

### GetPriceAmountForShipmentFloatOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetPriceAmountForShipmentFloatOk() (*interface{}, bool)`

GetPriceAmountForShipmentFloatOk returns a tuple with the PriceAmountForShipmentFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountForShipmentFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetPriceAmountForShipmentFloat(v interface{})`

SetPriceAmountForShipmentFloat sets PriceAmountForShipmentFloat field to given value.

### HasPriceAmountForShipmentFloat

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasPriceAmountForShipmentFloat() bool`

HasPriceAmountForShipmentFloat returns a boolean if a field has been set.

### SetPriceAmountForShipmentFloatNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetPriceAmountForShipmentFloatNil(b bool)`

 SetPriceAmountForShipmentFloatNil sets the value for PriceAmountForShipmentFloat to be an explicit nil

### UnsetPriceAmountForShipmentFloat
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetPriceAmountForShipmentFloat()`

UnsetPriceAmountForShipmentFloat ensures that no value is present for PriceAmountForShipmentFloat, not even an explicit nil
### GetFormattedPriceAmountForShipment

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFormattedPriceAmountForShipment() interface{}`

GetFormattedPriceAmountForShipment returns the FormattedPriceAmountForShipment field if non-nil, zero value otherwise.

### GetFormattedPriceAmountForShipmentOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetFormattedPriceAmountForShipmentOk() (*interface{}, bool)`

GetFormattedPriceAmountForShipmentOk returns a tuple with the FormattedPriceAmountForShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmountForShipment

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetFormattedPriceAmountForShipment(v interface{})`

SetFormattedPriceAmountForShipment sets FormattedPriceAmountForShipment field to given value.

### HasFormattedPriceAmountForShipment

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasFormattedPriceAmountForShipment() bool`

HasFormattedPriceAmountForShipment returns a boolean if a field has been set.

### SetFormattedPriceAmountForShipmentNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetFormattedPriceAmountForShipmentNil(b bool)`

 SetFormattedPriceAmountForShipmentNil sets the value for FormattedPriceAmountForShipment to be an explicit nil

### UnsetFormattedPriceAmountForShipment
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetFormattedPriceAmountForShipment()`

UnsetFormattedPriceAmountForShipment ensures that no value is present for FormattedPriceAmountForShipment, not even an explicit nil
### GetMinWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetMinWeight() interface{}`

GetMinWeight returns the MinWeight field if non-nil, zero value otherwise.

### GetMinWeightOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetMinWeightOk() (*interface{}, bool)`

GetMinWeightOk returns a tuple with the MinWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetMinWeight(v interface{})`

SetMinWeight sets MinWeight field to given value.

### HasMinWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasMinWeight() bool`

HasMinWeight returns a boolean if a field has been set.

### SetMinWeightNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetMinWeightNil(b bool)`

 SetMinWeightNil sets the value for MinWeight to be an explicit nil

### UnsetMinWeight
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetMinWeight()`

UnsetMinWeight ensures that no value is present for MinWeight, not even an explicit nil
### GetMaxWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetMaxWeight() interface{}`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetMaxWeightOk() (*interface{}, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetMaxWeight(v interface{})`

SetMaxWeight sets MaxWeight field to given value.

### HasMaxWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasMaxWeight() bool`

HasMaxWeight returns a boolean if a field has been set.

### SetMaxWeightNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetMaxWeightNil(b bool)`

 SetMaxWeightNil sets the value for MaxWeight to be an explicit nil

### UnsetMaxWeight
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetMaxWeight()`

UnsetMaxWeight ensures that no value is present for MaxWeight, not even an explicit nil
### GetUnitOfWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetUnitOfWeight() interface{}`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetUnitOfWeightOk() (*interface{}, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetUnitOfWeight(v interface{})`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### SetUnitOfWeightNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetUnitOfWeightNil(b bool)`

 SetUnitOfWeightNil sets the value for UnitOfWeight to be an explicit nil

### UnsetUnitOfWeight
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetUnitOfWeight()`

UnsetUnitOfWeight ensures that no value is present for UnitOfWeight, not even an explicit nil
### GetCreatedAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETShippingMethods200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETShippingMethods200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETShippingMethods200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETShippingMethods200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


