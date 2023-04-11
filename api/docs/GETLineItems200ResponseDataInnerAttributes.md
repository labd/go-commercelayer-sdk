# GETLineItems200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **interface{}** | The code of the associated SKU. | [optional] 
**BundleCode** | Pointer to **interface{}** | The code of the associated bundle. | [optional] 
**Quantity** | Pointer to **interface{}** | The line item quantity. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard, automatically inherited from the order&#39;s market. | [optional] 
**UnitAmountCents** | Pointer to **interface{}** | The unit amount of the line item, in cents. Can be specified without an item, otherwise is automatically populated from the price list associated to the order&#39;s market. | [optional] 
**UnitAmountFloat** | Pointer to **interface{}** | The unit amount of the line item, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce. | [optional] 
**FormattedUnitAmount** | Pointer to **interface{}** | The unit amount of the line item, formatted. This can be useful to display the amount with currency in you views. | [optional] 
**OptionsAmountCents** | Pointer to **interface{}** | The options amount of the line item, in cents. | [optional] 
**OptionsAmountFloat** | Pointer to **interface{}** | The options amount of the line item, float. | [optional] 
**FormattedOptionsAmount** | Pointer to **interface{}** | The options amount of the line item, formatted. | [optional] 
**DiscountCents** | Pointer to **interface{}** | The discount applied to the line item, in cents. When you apply a discount to an order, this is automatically calculated basing on the line item total_amount_cents value. | [optional] 
**DiscountFloat** | Pointer to **interface{}** | The discount applied to the line item, float. When you apply a discount to an order, this is automatically calculated basing on the line item total_amount_cents value. | [optional] 
**FormattedDiscount** | Pointer to **interface{}** | The discount applied to the line item, fromatted. When you apply a discount to an order, this is automatically calculated basing on the line item total_amount_cents value. | [optional] 
**TotalAmountCents** | Pointer to **interface{}** | Calculated as unit amount x quantity + options amount, in cents. | [optional] 
**TotalAmountFloat** | Pointer to **interface{}** | Calculated as unit amount x quantity + options amount, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce. | [optional] 
**FormattedTotalAmount** | Pointer to **interface{}** | Calculated as unit amount x quantity + options amount, formatted. This can be useful to display the amount with currency in you views. | [optional] 
**TaxAmountCents** | Pointer to **interface{}** | The collected tax amount, otherwise calculated as total amount cents - discount cent * tax rate, in cents. | [optional] 
**TaxAmountFloat** | Pointer to **interface{}** | The collected tax amount, otherwise calculated as total amount cents - discount cent * tax rate, float. | [optional] 
**FormattedTaxAmount** | Pointer to **interface{}** | The collected tax amount, otherwise calculated as total amount cents - discount cent * tax rate, formatted. | [optional] 
**Name** | Pointer to **interface{}** | The name of the line item. When blank, it gets populated with the name of the associated item (if present). | [optional] 
**ImageUrl** | Pointer to **interface{}** | The image_url of the line item. When blank, it gets populated with the image_url of the associated item (if present, SKU only). | [optional] 
**DiscountBreakdown** | Pointer to **interface{}** | The discount breakdown for this line item (if calculated). | [optional] 
**TaxRate** | Pointer to **interface{}** | The tax rate for this line item (if calculated). | [optional] 
**TaxBreakdown** | Pointer to **interface{}** | The tax breakdown for this line item (if calculated). | [optional] 
**ItemType** | Pointer to **interface{}** | The type of the associate item. Can be one of &#39;skus&#39;, &#39;bundles&#39;, &#39;shipments&#39;, &#39;payment_methods&#39;, &#39;adjustments&#39;, &#39;gift_cards&#39;, or a valid promotion type. | [optional] 
**Frequency** | Pointer to **interface{}** | The frequency which generates a subscription. Must be supported by existing associated subscription_model. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETLineItems200ResponseDataInnerAttributes

`func NewGETLineItems200ResponseDataInnerAttributes() *GETLineItems200ResponseDataInnerAttributes`

NewGETLineItems200ResponseDataInnerAttributes instantiates a new GETLineItems200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETLineItems200ResponseDataInnerAttributesWithDefaults

`func NewGETLineItems200ResponseDataInnerAttributesWithDefaults() *GETLineItems200ResponseDataInnerAttributes`

NewGETLineItems200ResponseDataInnerAttributesWithDefaults instantiates a new GETLineItems200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *GETLineItems200ResponseDataInnerAttributes) GetSkuCode() interface{}`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetSkuCodeOk() (*interface{}, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *GETLineItems200ResponseDataInnerAttributes) SetSkuCode(v interface{})`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *GETLineItems200ResponseDataInnerAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### SetSkuCodeNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetSkuCodeNil(b bool)`

 SetSkuCodeNil sets the value for SkuCode to be an explicit nil

### UnsetSkuCode
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetSkuCode()`

UnsetSkuCode ensures that no value is present for SkuCode, not even an explicit nil
### GetBundleCode

`func (o *GETLineItems200ResponseDataInnerAttributes) GetBundleCode() interface{}`

GetBundleCode returns the BundleCode field if non-nil, zero value otherwise.

### GetBundleCodeOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetBundleCodeOk() (*interface{}, bool)`

GetBundleCodeOk returns a tuple with the BundleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleCode

`func (o *GETLineItems200ResponseDataInnerAttributes) SetBundleCode(v interface{})`

SetBundleCode sets BundleCode field to given value.

### HasBundleCode

`func (o *GETLineItems200ResponseDataInnerAttributes) HasBundleCode() bool`

HasBundleCode returns a boolean if a field has been set.

### SetBundleCodeNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetBundleCodeNil(b bool)`

 SetBundleCodeNil sets the value for BundleCode to be an explicit nil

### UnsetBundleCode
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetBundleCode()`

UnsetBundleCode ensures that no value is present for BundleCode, not even an explicit nil
### GetQuantity

`func (o *GETLineItems200ResponseDataInnerAttributes) GetQuantity() interface{}`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetQuantityOk() (*interface{}, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GETLineItems200ResponseDataInnerAttributes) SetQuantity(v interface{})`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GETLineItems200ResponseDataInnerAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetCurrencyCode

`func (o *GETLineItems200ResponseDataInnerAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETLineItems200ResponseDataInnerAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETLineItems200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetUnitAmountCents

`func (o *GETLineItems200ResponseDataInnerAttributes) GetUnitAmountCents() interface{}`

GetUnitAmountCents returns the UnitAmountCents field if non-nil, zero value otherwise.

### GetUnitAmountCentsOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetUnitAmountCentsOk() (*interface{}, bool)`

GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountCents

`func (o *GETLineItems200ResponseDataInnerAttributes) SetUnitAmountCents(v interface{})`

SetUnitAmountCents sets UnitAmountCents field to given value.

### HasUnitAmountCents

`func (o *GETLineItems200ResponseDataInnerAttributes) HasUnitAmountCents() bool`

HasUnitAmountCents returns a boolean if a field has been set.

### SetUnitAmountCentsNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetUnitAmountCentsNil(b bool)`

 SetUnitAmountCentsNil sets the value for UnitAmountCents to be an explicit nil

### UnsetUnitAmountCents
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetUnitAmountCents()`

UnsetUnitAmountCents ensures that no value is present for UnitAmountCents, not even an explicit nil
### GetUnitAmountFloat

`func (o *GETLineItems200ResponseDataInnerAttributes) GetUnitAmountFloat() interface{}`

GetUnitAmountFloat returns the UnitAmountFloat field if non-nil, zero value otherwise.

### GetUnitAmountFloatOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetUnitAmountFloatOk() (*interface{}, bool)`

GetUnitAmountFloatOk returns a tuple with the UnitAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountFloat

`func (o *GETLineItems200ResponseDataInnerAttributes) SetUnitAmountFloat(v interface{})`

SetUnitAmountFloat sets UnitAmountFloat field to given value.

### HasUnitAmountFloat

`func (o *GETLineItems200ResponseDataInnerAttributes) HasUnitAmountFloat() bool`

HasUnitAmountFloat returns a boolean if a field has been set.

### SetUnitAmountFloatNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetUnitAmountFloatNil(b bool)`

 SetUnitAmountFloatNil sets the value for UnitAmountFloat to be an explicit nil

### UnsetUnitAmountFloat
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetUnitAmountFloat()`

UnsetUnitAmountFloat ensures that no value is present for UnitAmountFloat, not even an explicit nil
### GetFormattedUnitAmount

`func (o *GETLineItems200ResponseDataInnerAttributes) GetFormattedUnitAmount() interface{}`

GetFormattedUnitAmount returns the FormattedUnitAmount field if non-nil, zero value otherwise.

### GetFormattedUnitAmountOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetFormattedUnitAmountOk() (*interface{}, bool)`

GetFormattedUnitAmountOk returns a tuple with the FormattedUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedUnitAmount

`func (o *GETLineItems200ResponseDataInnerAttributes) SetFormattedUnitAmount(v interface{})`

SetFormattedUnitAmount sets FormattedUnitAmount field to given value.

### HasFormattedUnitAmount

`func (o *GETLineItems200ResponseDataInnerAttributes) HasFormattedUnitAmount() bool`

HasFormattedUnitAmount returns a boolean if a field has been set.

### SetFormattedUnitAmountNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetFormattedUnitAmountNil(b bool)`

 SetFormattedUnitAmountNil sets the value for FormattedUnitAmount to be an explicit nil

### UnsetFormattedUnitAmount
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetFormattedUnitAmount()`

UnsetFormattedUnitAmount ensures that no value is present for FormattedUnitAmount, not even an explicit nil
### GetOptionsAmountCents

`func (o *GETLineItems200ResponseDataInnerAttributes) GetOptionsAmountCents() interface{}`

GetOptionsAmountCents returns the OptionsAmountCents field if non-nil, zero value otherwise.

### GetOptionsAmountCentsOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetOptionsAmountCentsOk() (*interface{}, bool)`

GetOptionsAmountCentsOk returns a tuple with the OptionsAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsAmountCents

`func (o *GETLineItems200ResponseDataInnerAttributes) SetOptionsAmountCents(v interface{})`

SetOptionsAmountCents sets OptionsAmountCents field to given value.

### HasOptionsAmountCents

`func (o *GETLineItems200ResponseDataInnerAttributes) HasOptionsAmountCents() bool`

HasOptionsAmountCents returns a boolean if a field has been set.

### SetOptionsAmountCentsNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetOptionsAmountCentsNil(b bool)`

 SetOptionsAmountCentsNil sets the value for OptionsAmountCents to be an explicit nil

### UnsetOptionsAmountCents
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetOptionsAmountCents()`

UnsetOptionsAmountCents ensures that no value is present for OptionsAmountCents, not even an explicit nil
### GetOptionsAmountFloat

`func (o *GETLineItems200ResponseDataInnerAttributes) GetOptionsAmountFloat() interface{}`

GetOptionsAmountFloat returns the OptionsAmountFloat field if non-nil, zero value otherwise.

### GetOptionsAmountFloatOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetOptionsAmountFloatOk() (*interface{}, bool)`

GetOptionsAmountFloatOk returns a tuple with the OptionsAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsAmountFloat

`func (o *GETLineItems200ResponseDataInnerAttributes) SetOptionsAmountFloat(v interface{})`

SetOptionsAmountFloat sets OptionsAmountFloat field to given value.

### HasOptionsAmountFloat

`func (o *GETLineItems200ResponseDataInnerAttributes) HasOptionsAmountFloat() bool`

HasOptionsAmountFloat returns a boolean if a field has been set.

### SetOptionsAmountFloatNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetOptionsAmountFloatNil(b bool)`

 SetOptionsAmountFloatNil sets the value for OptionsAmountFloat to be an explicit nil

### UnsetOptionsAmountFloat
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetOptionsAmountFloat()`

UnsetOptionsAmountFloat ensures that no value is present for OptionsAmountFloat, not even an explicit nil
### GetFormattedOptionsAmount

`func (o *GETLineItems200ResponseDataInnerAttributes) GetFormattedOptionsAmount() interface{}`

GetFormattedOptionsAmount returns the FormattedOptionsAmount field if non-nil, zero value otherwise.

### GetFormattedOptionsAmountOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetFormattedOptionsAmountOk() (*interface{}, bool)`

GetFormattedOptionsAmountOk returns a tuple with the FormattedOptionsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedOptionsAmount

`func (o *GETLineItems200ResponseDataInnerAttributes) SetFormattedOptionsAmount(v interface{})`

SetFormattedOptionsAmount sets FormattedOptionsAmount field to given value.

### HasFormattedOptionsAmount

`func (o *GETLineItems200ResponseDataInnerAttributes) HasFormattedOptionsAmount() bool`

HasFormattedOptionsAmount returns a boolean if a field has been set.

### SetFormattedOptionsAmountNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetFormattedOptionsAmountNil(b bool)`

 SetFormattedOptionsAmountNil sets the value for FormattedOptionsAmount to be an explicit nil

### UnsetFormattedOptionsAmount
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetFormattedOptionsAmount()`

UnsetFormattedOptionsAmount ensures that no value is present for FormattedOptionsAmount, not even an explicit nil
### GetDiscountCents

`func (o *GETLineItems200ResponseDataInnerAttributes) GetDiscountCents() interface{}`

GetDiscountCents returns the DiscountCents field if non-nil, zero value otherwise.

### GetDiscountCentsOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetDiscountCentsOk() (*interface{}, bool)`

GetDiscountCentsOk returns a tuple with the DiscountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCents

`func (o *GETLineItems200ResponseDataInnerAttributes) SetDiscountCents(v interface{})`

SetDiscountCents sets DiscountCents field to given value.

### HasDiscountCents

`func (o *GETLineItems200ResponseDataInnerAttributes) HasDiscountCents() bool`

HasDiscountCents returns a boolean if a field has been set.

### SetDiscountCentsNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetDiscountCentsNil(b bool)`

 SetDiscountCentsNil sets the value for DiscountCents to be an explicit nil

### UnsetDiscountCents
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetDiscountCents()`

UnsetDiscountCents ensures that no value is present for DiscountCents, not even an explicit nil
### GetDiscountFloat

`func (o *GETLineItems200ResponseDataInnerAttributes) GetDiscountFloat() interface{}`

GetDiscountFloat returns the DiscountFloat field if non-nil, zero value otherwise.

### GetDiscountFloatOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetDiscountFloatOk() (*interface{}, bool)`

GetDiscountFloatOk returns a tuple with the DiscountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountFloat

`func (o *GETLineItems200ResponseDataInnerAttributes) SetDiscountFloat(v interface{})`

SetDiscountFloat sets DiscountFloat field to given value.

### HasDiscountFloat

`func (o *GETLineItems200ResponseDataInnerAttributes) HasDiscountFloat() bool`

HasDiscountFloat returns a boolean if a field has been set.

### SetDiscountFloatNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetDiscountFloatNil(b bool)`

 SetDiscountFloatNil sets the value for DiscountFloat to be an explicit nil

### UnsetDiscountFloat
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetDiscountFloat()`

UnsetDiscountFloat ensures that no value is present for DiscountFloat, not even an explicit nil
### GetFormattedDiscount

`func (o *GETLineItems200ResponseDataInnerAttributes) GetFormattedDiscount() interface{}`

GetFormattedDiscount returns the FormattedDiscount field if non-nil, zero value otherwise.

### GetFormattedDiscountOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetFormattedDiscountOk() (*interface{}, bool)`

GetFormattedDiscountOk returns a tuple with the FormattedDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedDiscount

`func (o *GETLineItems200ResponseDataInnerAttributes) SetFormattedDiscount(v interface{})`

SetFormattedDiscount sets FormattedDiscount field to given value.

### HasFormattedDiscount

`func (o *GETLineItems200ResponseDataInnerAttributes) HasFormattedDiscount() bool`

HasFormattedDiscount returns a boolean if a field has been set.

### SetFormattedDiscountNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetFormattedDiscountNil(b bool)`

 SetFormattedDiscountNil sets the value for FormattedDiscount to be an explicit nil

### UnsetFormattedDiscount
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetFormattedDiscount()`

UnsetFormattedDiscount ensures that no value is present for FormattedDiscount, not even an explicit nil
### GetTotalAmountCents

`func (o *GETLineItems200ResponseDataInnerAttributes) GetTotalAmountCents() interface{}`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetTotalAmountCentsOk() (*interface{}, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *GETLineItems200ResponseDataInnerAttributes) SetTotalAmountCents(v interface{})`

SetTotalAmountCents sets TotalAmountCents field to given value.

### HasTotalAmountCents

`func (o *GETLineItems200ResponseDataInnerAttributes) HasTotalAmountCents() bool`

HasTotalAmountCents returns a boolean if a field has been set.

### SetTotalAmountCentsNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetTotalAmountCentsNil(b bool)`

 SetTotalAmountCentsNil sets the value for TotalAmountCents to be an explicit nil

### UnsetTotalAmountCents
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetTotalAmountCents()`

UnsetTotalAmountCents ensures that no value is present for TotalAmountCents, not even an explicit nil
### GetTotalAmountFloat

`func (o *GETLineItems200ResponseDataInnerAttributes) GetTotalAmountFloat() interface{}`

GetTotalAmountFloat returns the TotalAmountFloat field if non-nil, zero value otherwise.

### GetTotalAmountFloatOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetTotalAmountFloatOk() (*interface{}, bool)`

GetTotalAmountFloatOk returns a tuple with the TotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountFloat

`func (o *GETLineItems200ResponseDataInnerAttributes) SetTotalAmountFloat(v interface{})`

SetTotalAmountFloat sets TotalAmountFloat field to given value.

### HasTotalAmountFloat

`func (o *GETLineItems200ResponseDataInnerAttributes) HasTotalAmountFloat() bool`

HasTotalAmountFloat returns a boolean if a field has been set.

### SetTotalAmountFloatNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetTotalAmountFloatNil(b bool)`

 SetTotalAmountFloatNil sets the value for TotalAmountFloat to be an explicit nil

### UnsetTotalAmountFloat
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetTotalAmountFloat()`

UnsetTotalAmountFloat ensures that no value is present for TotalAmountFloat, not even an explicit nil
### GetFormattedTotalAmount

`func (o *GETLineItems200ResponseDataInnerAttributes) GetFormattedTotalAmount() interface{}`

GetFormattedTotalAmount returns the FormattedTotalAmount field if non-nil, zero value otherwise.

### GetFormattedTotalAmountOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetFormattedTotalAmountOk() (*interface{}, bool)`

GetFormattedTotalAmountOk returns a tuple with the FormattedTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalAmount

`func (o *GETLineItems200ResponseDataInnerAttributes) SetFormattedTotalAmount(v interface{})`

SetFormattedTotalAmount sets FormattedTotalAmount field to given value.

### HasFormattedTotalAmount

`func (o *GETLineItems200ResponseDataInnerAttributes) HasFormattedTotalAmount() bool`

HasFormattedTotalAmount returns a boolean if a field has been set.

### SetFormattedTotalAmountNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetFormattedTotalAmountNil(b bool)`

 SetFormattedTotalAmountNil sets the value for FormattedTotalAmount to be an explicit nil

### UnsetFormattedTotalAmount
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetFormattedTotalAmount()`

UnsetFormattedTotalAmount ensures that no value is present for FormattedTotalAmount, not even an explicit nil
### GetTaxAmountCents

`func (o *GETLineItems200ResponseDataInnerAttributes) GetTaxAmountCents() interface{}`

GetTaxAmountCents returns the TaxAmountCents field if non-nil, zero value otherwise.

### GetTaxAmountCentsOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetTaxAmountCentsOk() (*interface{}, bool)`

GetTaxAmountCentsOk returns a tuple with the TaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountCents

`func (o *GETLineItems200ResponseDataInnerAttributes) SetTaxAmountCents(v interface{})`

SetTaxAmountCents sets TaxAmountCents field to given value.

### HasTaxAmountCents

`func (o *GETLineItems200ResponseDataInnerAttributes) HasTaxAmountCents() bool`

HasTaxAmountCents returns a boolean if a field has been set.

### SetTaxAmountCentsNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetTaxAmountCentsNil(b bool)`

 SetTaxAmountCentsNil sets the value for TaxAmountCents to be an explicit nil

### UnsetTaxAmountCents
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetTaxAmountCents()`

UnsetTaxAmountCents ensures that no value is present for TaxAmountCents, not even an explicit nil
### GetTaxAmountFloat

`func (o *GETLineItems200ResponseDataInnerAttributes) GetTaxAmountFloat() interface{}`

GetTaxAmountFloat returns the TaxAmountFloat field if non-nil, zero value otherwise.

### GetTaxAmountFloatOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetTaxAmountFloatOk() (*interface{}, bool)`

GetTaxAmountFloatOk returns a tuple with the TaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountFloat

`func (o *GETLineItems200ResponseDataInnerAttributes) SetTaxAmountFloat(v interface{})`

SetTaxAmountFloat sets TaxAmountFloat field to given value.

### HasTaxAmountFloat

`func (o *GETLineItems200ResponseDataInnerAttributes) HasTaxAmountFloat() bool`

HasTaxAmountFloat returns a boolean if a field has been set.

### SetTaxAmountFloatNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetTaxAmountFloatNil(b bool)`

 SetTaxAmountFloatNil sets the value for TaxAmountFloat to be an explicit nil

### UnsetTaxAmountFloat
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetTaxAmountFloat()`

UnsetTaxAmountFloat ensures that no value is present for TaxAmountFloat, not even an explicit nil
### GetFormattedTaxAmount

`func (o *GETLineItems200ResponseDataInnerAttributes) GetFormattedTaxAmount() interface{}`

GetFormattedTaxAmount returns the FormattedTaxAmount field if non-nil, zero value otherwise.

### GetFormattedTaxAmountOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetFormattedTaxAmountOk() (*interface{}, bool)`

GetFormattedTaxAmountOk returns a tuple with the FormattedTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTaxAmount

`func (o *GETLineItems200ResponseDataInnerAttributes) SetFormattedTaxAmount(v interface{})`

SetFormattedTaxAmount sets FormattedTaxAmount field to given value.

### HasFormattedTaxAmount

`func (o *GETLineItems200ResponseDataInnerAttributes) HasFormattedTaxAmount() bool`

HasFormattedTaxAmount returns a boolean if a field has been set.

### SetFormattedTaxAmountNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetFormattedTaxAmountNil(b bool)`

 SetFormattedTaxAmountNil sets the value for FormattedTaxAmount to be an explicit nil

### UnsetFormattedTaxAmount
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetFormattedTaxAmount()`

UnsetFormattedTaxAmount ensures that no value is present for FormattedTaxAmount, not even an explicit nil
### GetName

`func (o *GETLineItems200ResponseDataInnerAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETLineItems200ResponseDataInnerAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETLineItems200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImageUrl

`func (o *GETLineItems200ResponseDataInnerAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GETLineItems200ResponseDataInnerAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GETLineItems200ResponseDataInnerAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetDiscountBreakdown

`func (o *GETLineItems200ResponseDataInnerAttributes) GetDiscountBreakdown() interface{}`

GetDiscountBreakdown returns the DiscountBreakdown field if non-nil, zero value otherwise.

### GetDiscountBreakdownOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetDiscountBreakdownOk() (*interface{}, bool)`

GetDiscountBreakdownOk returns a tuple with the DiscountBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountBreakdown

`func (o *GETLineItems200ResponseDataInnerAttributes) SetDiscountBreakdown(v interface{})`

SetDiscountBreakdown sets DiscountBreakdown field to given value.

### HasDiscountBreakdown

`func (o *GETLineItems200ResponseDataInnerAttributes) HasDiscountBreakdown() bool`

HasDiscountBreakdown returns a boolean if a field has been set.

### SetDiscountBreakdownNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetDiscountBreakdownNil(b bool)`

 SetDiscountBreakdownNil sets the value for DiscountBreakdown to be an explicit nil

### UnsetDiscountBreakdown
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetDiscountBreakdown()`

UnsetDiscountBreakdown ensures that no value is present for DiscountBreakdown, not even an explicit nil
### GetTaxRate

`func (o *GETLineItems200ResponseDataInnerAttributes) GetTaxRate() interface{}`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetTaxRateOk() (*interface{}, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *GETLineItems200ResponseDataInnerAttributes) SetTaxRate(v interface{})`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *GETLineItems200ResponseDataInnerAttributes) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetTaxBreakdown

`func (o *GETLineItems200ResponseDataInnerAttributes) GetTaxBreakdown() interface{}`

GetTaxBreakdown returns the TaxBreakdown field if non-nil, zero value otherwise.

### GetTaxBreakdownOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetTaxBreakdownOk() (*interface{}, bool)`

GetTaxBreakdownOk returns a tuple with the TaxBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBreakdown

`func (o *GETLineItems200ResponseDataInnerAttributes) SetTaxBreakdown(v interface{})`

SetTaxBreakdown sets TaxBreakdown field to given value.

### HasTaxBreakdown

`func (o *GETLineItems200ResponseDataInnerAttributes) HasTaxBreakdown() bool`

HasTaxBreakdown returns a boolean if a field has been set.

### SetTaxBreakdownNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetTaxBreakdownNil(b bool)`

 SetTaxBreakdownNil sets the value for TaxBreakdown to be an explicit nil

### UnsetTaxBreakdown
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetTaxBreakdown()`

UnsetTaxBreakdown ensures that no value is present for TaxBreakdown, not even an explicit nil
### GetItemType

`func (o *GETLineItems200ResponseDataInnerAttributes) GetItemType() interface{}`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetItemTypeOk() (*interface{}, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *GETLineItems200ResponseDataInnerAttributes) SetItemType(v interface{})`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *GETLineItems200ResponseDataInnerAttributes) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetFrequency

`func (o *GETLineItems200ResponseDataInnerAttributes) GetFrequency() interface{}`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetFrequencyOk() (*interface{}, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GETLineItems200ResponseDataInnerAttributes) SetFrequency(v interface{})`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GETLineItems200ResponseDataInnerAttributes) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetCreatedAt

`func (o *GETLineItems200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETLineItems200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETLineItems200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETLineItems200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETLineItems200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETLineItems200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETLineItems200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETLineItems200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETLineItems200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETLineItems200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETLineItems200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETLineItems200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETLineItems200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETLineItems200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETLineItems200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETLineItems200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETLineItems200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETLineItems200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


