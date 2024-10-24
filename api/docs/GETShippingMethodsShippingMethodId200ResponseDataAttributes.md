# GETShippingMethodsShippingMethodId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The shipping method&#39;s name. | [optional] 
**Scheme** | Pointer to **interface{}** | The shipping method&#39;s scheme. One of &#39;flat&#39;, &#39;weight_tiered&#39;, or &#39;external&#39;. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**ExternalPricesUrl** | Pointer to **interface{}** | The URL used to overwrite prices by an external source. | [optional] 
**PriceAmountCents** | Pointer to **interface{}** | The price of this shipping method, in cents. | [optional] 
**PriceAmountFloat** | Pointer to **interface{}** | The price of this shipping method, float. | [optional] 
**FormattedPriceAmount** | Pointer to **interface{}** | The price of this shipping method, formatted. | [optional] 
**FreeOverAmountCents** | Pointer to **interface{}** | Apply free shipping if the order amount is over this value, in cents. | [optional] 
**FreeOverAmountFloat** | Pointer to **interface{}** | Apply free shipping if the order amount is over this value, float. | [optional] 
**FormattedFreeOverAmount** | Pointer to **interface{}** | Apply free shipping if the order amount is over this value, formatted. | [optional] 
**UseSubtotal** | Pointer to **interface{}** | Send this attribute if you want to compare the free over amount with order&#39;s subtotal (excluding discounts, if any). | [optional] 
**PriceAmountForShipmentCents** | Pointer to **interface{}** | The calculated price (zero or price amount) when associated to a shipment, in cents. | [optional] 
**PriceAmountForShipmentFloat** | Pointer to **interface{}** | The calculated price (zero or price amount) when associated to a shipment, float. | [optional] 
**FormattedPriceAmountForShipment** | Pointer to **interface{}** | The calculated price (zero or price amount) when associated to a shipment, formatted. | [optional] 
**MinWeight** | Pointer to **interface{}** | The minimum weight for which this shipping method is available. | [optional] 
**MaxWeight** | Pointer to **interface{}** | The maximum weight for which this shipping method is available. | [optional] 
**UnitOfWeight** | Pointer to **interface{}** | The unit of weight. One of &#39;gr&#39;, &#39;oz&#39;, or &#39;lb&#39;. | [optional] 
**DisabledAt** | Pointer to **interface{}** | Time at which this resource was disabled. | [optional] 
**CircuitState** | Pointer to **interface{}** | The circuit breaker state, by default it is &#39;closed&#39;. It can become &#39;open&#39; once the number of consecutive failures overlaps the specified threshold, in such case no further calls to the failing callback are made. | [optional] 
**CircuitFailureCount** | Pointer to **interface{}** | The number of consecutive failures recorded by the circuit breaker associated to this resource, will be reset on first successful call to callback. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETShippingMethodsShippingMethodId200ResponseDataAttributes

`func NewGETShippingMethodsShippingMethodId200ResponseDataAttributes() *GETShippingMethodsShippingMethodId200ResponseDataAttributes`

NewGETShippingMethodsShippingMethodId200ResponseDataAttributes instantiates a new GETShippingMethodsShippingMethodId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETShippingMethodsShippingMethodId200ResponseDataAttributesWithDefaults

`func NewGETShippingMethodsShippingMethodId200ResponseDataAttributesWithDefaults() *GETShippingMethodsShippingMethodId200ResponseDataAttributes`

NewGETShippingMethodsShippingMethodId200ResponseDataAttributesWithDefaults instantiates a new GETShippingMethodsShippingMethodId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetScheme

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetScheme() interface{}`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetSchemeOk() (*interface{}, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetScheme(v interface{})`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetCurrencyCode

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetExternalPricesUrl

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetExternalPricesUrl() interface{}`

GetExternalPricesUrl returns the ExternalPricesUrl field if non-nil, zero value otherwise.

### GetExternalPricesUrlOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetExternalPricesUrlOk() (*interface{}, bool)`

GetExternalPricesUrlOk returns a tuple with the ExternalPricesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPricesUrl

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetExternalPricesUrl(v interface{})`

SetExternalPricesUrl sets ExternalPricesUrl field to given value.

### HasExternalPricesUrl

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasExternalPricesUrl() bool`

HasExternalPricesUrl returns a boolean if a field has been set.

### SetExternalPricesUrlNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetExternalPricesUrlNil(b bool)`

 SetExternalPricesUrlNil sets the value for ExternalPricesUrl to be an explicit nil

### UnsetExternalPricesUrl
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetExternalPricesUrl()`

UnsetExternalPricesUrl ensures that no value is present for ExternalPricesUrl, not even an explicit nil
### GetPriceAmountCents

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### SetPriceAmountCentsNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetPriceAmountFloat

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetPriceAmountFloat() interface{}`

GetPriceAmountFloat returns the PriceAmountFloat field if non-nil, zero value otherwise.

### GetPriceAmountFloatOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetPriceAmountFloatOk() (*interface{}, bool)`

GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountFloat

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetPriceAmountFloat(v interface{})`

SetPriceAmountFloat sets PriceAmountFloat field to given value.

### HasPriceAmountFloat

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasPriceAmountFloat() bool`

HasPriceAmountFloat returns a boolean if a field has been set.

### SetPriceAmountFloatNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetPriceAmountFloatNil(b bool)`

 SetPriceAmountFloatNil sets the value for PriceAmountFloat to be an explicit nil

### UnsetPriceAmountFloat
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetPriceAmountFloat()`

UnsetPriceAmountFloat ensures that no value is present for PriceAmountFloat, not even an explicit nil
### GetFormattedPriceAmount

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetFormattedPriceAmount() interface{}`

GetFormattedPriceAmount returns the FormattedPriceAmount field if non-nil, zero value otherwise.

### GetFormattedPriceAmountOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetFormattedPriceAmountOk() (*interface{}, bool)`

GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmount

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetFormattedPriceAmount(v interface{})`

SetFormattedPriceAmount sets FormattedPriceAmount field to given value.

### HasFormattedPriceAmount

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasFormattedPriceAmount() bool`

HasFormattedPriceAmount returns a boolean if a field has been set.

### SetFormattedPriceAmountNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetFormattedPriceAmountNil(b bool)`

 SetFormattedPriceAmountNil sets the value for FormattedPriceAmount to be an explicit nil

### UnsetFormattedPriceAmount
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetFormattedPriceAmount()`

UnsetFormattedPriceAmount ensures that no value is present for FormattedPriceAmount, not even an explicit nil
### GetFreeOverAmountCents

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetFreeOverAmountCents() interface{}`

GetFreeOverAmountCents returns the FreeOverAmountCents field if non-nil, zero value otherwise.

### GetFreeOverAmountCentsOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetFreeOverAmountCentsOk() (*interface{}, bool)`

GetFreeOverAmountCentsOk returns a tuple with the FreeOverAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeOverAmountCents

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetFreeOverAmountCents(v interface{})`

SetFreeOverAmountCents sets FreeOverAmountCents field to given value.

### HasFreeOverAmountCents

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasFreeOverAmountCents() bool`

HasFreeOverAmountCents returns a boolean if a field has been set.

### SetFreeOverAmountCentsNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetFreeOverAmountCentsNil(b bool)`

 SetFreeOverAmountCentsNil sets the value for FreeOverAmountCents to be an explicit nil

### UnsetFreeOverAmountCents
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetFreeOverAmountCents()`

UnsetFreeOverAmountCents ensures that no value is present for FreeOverAmountCents, not even an explicit nil
### GetFreeOverAmountFloat

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetFreeOverAmountFloat() interface{}`

GetFreeOverAmountFloat returns the FreeOverAmountFloat field if non-nil, zero value otherwise.

### GetFreeOverAmountFloatOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetFreeOverAmountFloatOk() (*interface{}, bool)`

GetFreeOverAmountFloatOk returns a tuple with the FreeOverAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeOverAmountFloat

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetFreeOverAmountFloat(v interface{})`

SetFreeOverAmountFloat sets FreeOverAmountFloat field to given value.

### HasFreeOverAmountFloat

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasFreeOverAmountFloat() bool`

HasFreeOverAmountFloat returns a boolean if a field has been set.

### SetFreeOverAmountFloatNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetFreeOverAmountFloatNil(b bool)`

 SetFreeOverAmountFloatNil sets the value for FreeOverAmountFloat to be an explicit nil

### UnsetFreeOverAmountFloat
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetFreeOverAmountFloat()`

UnsetFreeOverAmountFloat ensures that no value is present for FreeOverAmountFloat, not even an explicit nil
### GetFormattedFreeOverAmount

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetFormattedFreeOverAmount() interface{}`

GetFormattedFreeOverAmount returns the FormattedFreeOverAmount field if non-nil, zero value otherwise.

### GetFormattedFreeOverAmountOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetFormattedFreeOverAmountOk() (*interface{}, bool)`

GetFormattedFreeOverAmountOk returns a tuple with the FormattedFreeOverAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedFreeOverAmount

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetFormattedFreeOverAmount(v interface{})`

SetFormattedFreeOverAmount sets FormattedFreeOverAmount field to given value.

### HasFormattedFreeOverAmount

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasFormattedFreeOverAmount() bool`

HasFormattedFreeOverAmount returns a boolean if a field has been set.

### SetFormattedFreeOverAmountNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetFormattedFreeOverAmountNil(b bool)`

 SetFormattedFreeOverAmountNil sets the value for FormattedFreeOverAmount to be an explicit nil

### UnsetFormattedFreeOverAmount
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetFormattedFreeOverAmount()`

UnsetFormattedFreeOverAmount ensures that no value is present for FormattedFreeOverAmount, not even an explicit nil
### GetUseSubtotal

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetUseSubtotal() interface{}`

GetUseSubtotal returns the UseSubtotal field if non-nil, zero value otherwise.

### GetUseSubtotalOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetUseSubtotalOk() (*interface{}, bool)`

GetUseSubtotalOk returns a tuple with the UseSubtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSubtotal

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetUseSubtotal(v interface{})`

SetUseSubtotal sets UseSubtotal field to given value.

### HasUseSubtotal

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasUseSubtotal() bool`

HasUseSubtotal returns a boolean if a field has been set.

### SetUseSubtotalNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetUseSubtotalNil(b bool)`

 SetUseSubtotalNil sets the value for UseSubtotal to be an explicit nil

### UnsetUseSubtotal
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetUseSubtotal()`

UnsetUseSubtotal ensures that no value is present for UseSubtotal, not even an explicit nil
### GetPriceAmountForShipmentCents

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetPriceAmountForShipmentCents() interface{}`

GetPriceAmountForShipmentCents returns the PriceAmountForShipmentCents field if non-nil, zero value otherwise.

### GetPriceAmountForShipmentCentsOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetPriceAmountForShipmentCentsOk() (*interface{}, bool)`

GetPriceAmountForShipmentCentsOk returns a tuple with the PriceAmountForShipmentCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountForShipmentCents

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetPriceAmountForShipmentCents(v interface{})`

SetPriceAmountForShipmentCents sets PriceAmountForShipmentCents field to given value.

### HasPriceAmountForShipmentCents

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasPriceAmountForShipmentCents() bool`

HasPriceAmountForShipmentCents returns a boolean if a field has been set.

### SetPriceAmountForShipmentCentsNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetPriceAmountForShipmentCentsNil(b bool)`

 SetPriceAmountForShipmentCentsNil sets the value for PriceAmountForShipmentCents to be an explicit nil

### UnsetPriceAmountForShipmentCents
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetPriceAmountForShipmentCents()`

UnsetPriceAmountForShipmentCents ensures that no value is present for PriceAmountForShipmentCents, not even an explicit nil
### GetPriceAmountForShipmentFloat

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetPriceAmountForShipmentFloat() interface{}`

GetPriceAmountForShipmentFloat returns the PriceAmountForShipmentFloat field if non-nil, zero value otherwise.

### GetPriceAmountForShipmentFloatOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetPriceAmountForShipmentFloatOk() (*interface{}, bool)`

GetPriceAmountForShipmentFloatOk returns a tuple with the PriceAmountForShipmentFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountForShipmentFloat

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetPriceAmountForShipmentFloat(v interface{})`

SetPriceAmountForShipmentFloat sets PriceAmountForShipmentFloat field to given value.

### HasPriceAmountForShipmentFloat

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasPriceAmountForShipmentFloat() bool`

HasPriceAmountForShipmentFloat returns a boolean if a field has been set.

### SetPriceAmountForShipmentFloatNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetPriceAmountForShipmentFloatNil(b bool)`

 SetPriceAmountForShipmentFloatNil sets the value for PriceAmountForShipmentFloat to be an explicit nil

### UnsetPriceAmountForShipmentFloat
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetPriceAmountForShipmentFloat()`

UnsetPriceAmountForShipmentFloat ensures that no value is present for PriceAmountForShipmentFloat, not even an explicit nil
### GetFormattedPriceAmountForShipment

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetFormattedPriceAmountForShipment() interface{}`

GetFormattedPriceAmountForShipment returns the FormattedPriceAmountForShipment field if non-nil, zero value otherwise.

### GetFormattedPriceAmountForShipmentOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetFormattedPriceAmountForShipmentOk() (*interface{}, bool)`

GetFormattedPriceAmountForShipmentOk returns a tuple with the FormattedPriceAmountForShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmountForShipment

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetFormattedPriceAmountForShipment(v interface{})`

SetFormattedPriceAmountForShipment sets FormattedPriceAmountForShipment field to given value.

### HasFormattedPriceAmountForShipment

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasFormattedPriceAmountForShipment() bool`

HasFormattedPriceAmountForShipment returns a boolean if a field has been set.

### SetFormattedPriceAmountForShipmentNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetFormattedPriceAmountForShipmentNil(b bool)`

 SetFormattedPriceAmountForShipmentNil sets the value for FormattedPriceAmountForShipment to be an explicit nil

### UnsetFormattedPriceAmountForShipment
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetFormattedPriceAmountForShipment()`

UnsetFormattedPriceAmountForShipment ensures that no value is present for FormattedPriceAmountForShipment, not even an explicit nil
### GetMinWeight

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetMinWeight() interface{}`

GetMinWeight returns the MinWeight field if non-nil, zero value otherwise.

### GetMinWeightOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetMinWeightOk() (*interface{}, bool)`

GetMinWeightOk returns a tuple with the MinWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWeight

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetMinWeight(v interface{})`

SetMinWeight sets MinWeight field to given value.

### HasMinWeight

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasMinWeight() bool`

HasMinWeight returns a boolean if a field has been set.

### SetMinWeightNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetMinWeightNil(b bool)`

 SetMinWeightNil sets the value for MinWeight to be an explicit nil

### UnsetMinWeight
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetMinWeight()`

UnsetMinWeight ensures that no value is present for MinWeight, not even an explicit nil
### GetMaxWeight

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetMaxWeight() interface{}`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetMaxWeightOk() (*interface{}, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetMaxWeight(v interface{})`

SetMaxWeight sets MaxWeight field to given value.

### HasMaxWeight

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasMaxWeight() bool`

HasMaxWeight returns a boolean if a field has been set.

### SetMaxWeightNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetMaxWeightNil(b bool)`

 SetMaxWeightNil sets the value for MaxWeight to be an explicit nil

### UnsetMaxWeight
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetMaxWeight()`

UnsetMaxWeight ensures that no value is present for MaxWeight, not even an explicit nil
### GetUnitOfWeight

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetUnitOfWeight() interface{}`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetUnitOfWeightOk() (*interface{}, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetUnitOfWeight(v interface{})`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### SetUnitOfWeightNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetUnitOfWeightNil(b bool)`

 SetUnitOfWeightNil sets the value for UnitOfWeight to be an explicit nil

### UnsetUnitOfWeight
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetUnitOfWeight()`

UnsetUnitOfWeight ensures that no value is present for UnitOfWeight, not even an explicit nil
### GetDisabledAt

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetDisabledAt() interface{}`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetDisabledAtOk() (*interface{}, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetDisabledAt(v interface{})`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### SetDisabledAtNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetDisabledAtNil(b bool)`

 SetDisabledAtNil sets the value for DisabledAt to be an explicit nil

### UnsetDisabledAt
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetDisabledAt()`

UnsetDisabledAt ensures that no value is present for DisabledAt, not even an explicit nil
### GetCircuitState

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetCircuitState() interface{}`

GetCircuitState returns the CircuitState field if non-nil, zero value otherwise.

### GetCircuitStateOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetCircuitStateOk() (*interface{}, bool)`

GetCircuitStateOk returns a tuple with the CircuitState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitState

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetCircuitState(v interface{})`

SetCircuitState sets CircuitState field to given value.

### HasCircuitState

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasCircuitState() bool`

HasCircuitState returns a boolean if a field has been set.

### SetCircuitStateNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetCircuitStateNil(b bool)`

 SetCircuitStateNil sets the value for CircuitState to be an explicit nil

### UnsetCircuitState
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetCircuitState()`

UnsetCircuitState ensures that no value is present for CircuitState, not even an explicit nil
### GetCircuitFailureCount

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetCircuitFailureCount() interface{}`

GetCircuitFailureCount returns the CircuitFailureCount field if non-nil, zero value otherwise.

### GetCircuitFailureCountOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetCircuitFailureCountOk() (*interface{}, bool)`

GetCircuitFailureCountOk returns a tuple with the CircuitFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitFailureCount

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetCircuitFailureCount(v interface{})`

SetCircuitFailureCount sets CircuitFailureCount field to given value.

### HasCircuitFailureCount

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasCircuitFailureCount() bool`

HasCircuitFailureCount returns a boolean if a field has been set.

### SetCircuitFailureCountNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetCircuitFailureCountNil(b bool)`

 SetCircuitFailureCountNil sets the value for CircuitFailureCount to be an explicit nil

### UnsetCircuitFailureCount
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetCircuitFailureCount()`

UnsetCircuitFailureCount ensures that no value is present for CircuitFailureCount, not even an explicit nil
### GetCreatedAt

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETShippingMethodsShippingMethodId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


