# LineItemDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **string** | The code of the associated sku. | [optional] 
**BundleCode** | Pointer to **string** | The code of the associated bundle. | [optional] 
**Quantity** | Pointer to **int32** | The line item quantity. | [optional] 
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard, automatically inherited from the order&#39;s market. | [optional] 
**UnitAmountCents** | Pointer to **int32** | The unit amount of the line item, in cents. Can be specified without an item, otherwise is automatically populated from the price list associated to the order&#39;s market. | [optional] 
**UnitAmountFloat** | Pointer to **float32** | The unit amount of the line item, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce. | [optional] 
**FormattedUnitAmount** | Pointer to **string** | The unit amount of the line item, formatted. This can be useful to display the amount with currency in you views. | [optional] 
**OptionsAmountCents** | Pointer to **int32** | The options amount of the line item, in cents. | [optional] 
**OptionsAmountFloat** | Pointer to **float32** | The options amount of the line item, float. | [optional] 
**FormattedOptionsAmount** | Pointer to **string** | The options amount of the line item, formatted. | [optional] 
**DiscountCents** | Pointer to **int32** | The discount applied to the line item, in cents. When you apply a discount to an order, this is automatically calculated basing on the line item total_amount_cents value. | [optional] 
**DiscountFloat** | Pointer to **float32** | The discount applied to the line item, float. When you apply a discount to an order, this is automatically calculated basing on the line item total_amount_cents value. | [optional] 
**FormattedDiscount** | Pointer to **string** | The discount applied to the line item, fromatted. When you apply a discount to an order, this is automatically calculated basing on the line item total_amount_cents value. | [optional] 
**TotalAmountCents** | Pointer to **int32** | Calculated as unit amount x quantity + options amount, in cents. | [optional] 
**TotalAmountFloat** | Pointer to **float32** | Calculated as unit amount x quantity + options amount, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce. | [optional] 
**FormattedTotalAmount** | Pointer to **string** | Calculated as unit amount x quantity + options amount, formatted. This can be useful to display the amount with currency in you views. | [optional] 
**TaxAmountCents** | Pointer to **int32** | Calculated as total amount cents - discount cent * tax rate, in cents. | [optional] 
**TaxAmountFloat** | Pointer to **float32** | Calculated as total amount cents - discount cent * tax rate, float. | [optional] 
**FormattedTaxAmount** | Pointer to **string** | Calculated as total amount cents - discount cent * tax rate, formatted. | [optional] 
**Name** | Pointer to **string** | The name of the line item. When blank, it gets populated with the name of the associated item (if present). | [optional] 
**ImageUrl** | Pointer to **string** | The image_url of the line item. When blank, it gets populated with the image_url of the associated item (if present, sku only). | [optional] 
**DiscountBreakdown** | Pointer to **map[string]interface{}** | The discount breakdown for this line item (if calculated). | [optional] 
**TaxRate** | Pointer to **float32** | The tax rate for this line item (if calculated). | [optional] 
**TaxBreakdown** | Pointer to **map[string]interface{}** | The tax breakdown for this line item (if calculated). | [optional] 
**ItemType** | Pointer to **string** | The type of the associate item. Can be one of &#39;sku&#39;, &#39;bundle&#39;, &#39;shipment&#39;, &#39;payment_method&#39;, &#39;adjustment&#39;, &#39;gift_card&#39;, or a valid promotion type. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewLineItemDataAttributes

`func NewLineItemDataAttributes() *LineItemDataAttributes`

NewLineItemDataAttributes instantiates a new LineItemDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemDataAttributesWithDefaults

`func NewLineItemDataAttributesWithDefaults() *LineItemDataAttributes`

NewLineItemDataAttributesWithDefaults instantiates a new LineItemDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *LineItemDataAttributes) GetSkuCode() string`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *LineItemDataAttributes) GetSkuCodeOk() (*string, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *LineItemDataAttributes) SetSkuCode(v string)`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *LineItemDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### GetBundleCode

`func (o *LineItemDataAttributes) GetBundleCode() string`

GetBundleCode returns the BundleCode field if non-nil, zero value otherwise.

### GetBundleCodeOk

`func (o *LineItemDataAttributes) GetBundleCodeOk() (*string, bool)`

GetBundleCodeOk returns a tuple with the BundleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleCode

`func (o *LineItemDataAttributes) SetBundleCode(v string)`

SetBundleCode sets BundleCode field to given value.

### HasBundleCode

`func (o *LineItemDataAttributes) HasBundleCode() bool`

HasBundleCode returns a boolean if a field has been set.

### GetQuantity

`func (o *LineItemDataAttributes) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *LineItemDataAttributes) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *LineItemDataAttributes) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *LineItemDataAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *LineItemDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *LineItemDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *LineItemDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *LineItemDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetUnitAmountCents

`func (o *LineItemDataAttributes) GetUnitAmountCents() int32`

GetUnitAmountCents returns the UnitAmountCents field if non-nil, zero value otherwise.

### GetUnitAmountCentsOk

`func (o *LineItemDataAttributes) GetUnitAmountCentsOk() (*int32, bool)`

GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountCents

`func (o *LineItemDataAttributes) SetUnitAmountCents(v int32)`

SetUnitAmountCents sets UnitAmountCents field to given value.

### HasUnitAmountCents

`func (o *LineItemDataAttributes) HasUnitAmountCents() bool`

HasUnitAmountCents returns a boolean if a field has been set.

### GetUnitAmountFloat

`func (o *LineItemDataAttributes) GetUnitAmountFloat() float32`

GetUnitAmountFloat returns the UnitAmountFloat field if non-nil, zero value otherwise.

### GetUnitAmountFloatOk

`func (o *LineItemDataAttributes) GetUnitAmountFloatOk() (*float32, bool)`

GetUnitAmountFloatOk returns a tuple with the UnitAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountFloat

`func (o *LineItemDataAttributes) SetUnitAmountFloat(v float32)`

SetUnitAmountFloat sets UnitAmountFloat field to given value.

### HasUnitAmountFloat

`func (o *LineItemDataAttributes) HasUnitAmountFloat() bool`

HasUnitAmountFloat returns a boolean if a field has been set.

### GetFormattedUnitAmount

`func (o *LineItemDataAttributes) GetFormattedUnitAmount() string`

GetFormattedUnitAmount returns the FormattedUnitAmount field if non-nil, zero value otherwise.

### GetFormattedUnitAmountOk

`func (o *LineItemDataAttributes) GetFormattedUnitAmountOk() (*string, bool)`

GetFormattedUnitAmountOk returns a tuple with the FormattedUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedUnitAmount

`func (o *LineItemDataAttributes) SetFormattedUnitAmount(v string)`

SetFormattedUnitAmount sets FormattedUnitAmount field to given value.

### HasFormattedUnitAmount

`func (o *LineItemDataAttributes) HasFormattedUnitAmount() bool`

HasFormattedUnitAmount returns a boolean if a field has been set.

### GetOptionsAmountCents

`func (o *LineItemDataAttributes) GetOptionsAmountCents() int32`

GetOptionsAmountCents returns the OptionsAmountCents field if non-nil, zero value otherwise.

### GetOptionsAmountCentsOk

`func (o *LineItemDataAttributes) GetOptionsAmountCentsOk() (*int32, bool)`

GetOptionsAmountCentsOk returns a tuple with the OptionsAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsAmountCents

`func (o *LineItemDataAttributes) SetOptionsAmountCents(v int32)`

SetOptionsAmountCents sets OptionsAmountCents field to given value.

### HasOptionsAmountCents

`func (o *LineItemDataAttributes) HasOptionsAmountCents() bool`

HasOptionsAmountCents returns a boolean if a field has been set.

### GetOptionsAmountFloat

`func (o *LineItemDataAttributes) GetOptionsAmountFloat() float32`

GetOptionsAmountFloat returns the OptionsAmountFloat field if non-nil, zero value otherwise.

### GetOptionsAmountFloatOk

`func (o *LineItemDataAttributes) GetOptionsAmountFloatOk() (*float32, bool)`

GetOptionsAmountFloatOk returns a tuple with the OptionsAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsAmountFloat

`func (o *LineItemDataAttributes) SetOptionsAmountFloat(v float32)`

SetOptionsAmountFloat sets OptionsAmountFloat field to given value.

### HasOptionsAmountFloat

`func (o *LineItemDataAttributes) HasOptionsAmountFloat() bool`

HasOptionsAmountFloat returns a boolean if a field has been set.

### GetFormattedOptionsAmount

`func (o *LineItemDataAttributes) GetFormattedOptionsAmount() string`

GetFormattedOptionsAmount returns the FormattedOptionsAmount field if non-nil, zero value otherwise.

### GetFormattedOptionsAmountOk

`func (o *LineItemDataAttributes) GetFormattedOptionsAmountOk() (*string, bool)`

GetFormattedOptionsAmountOk returns a tuple with the FormattedOptionsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedOptionsAmount

`func (o *LineItemDataAttributes) SetFormattedOptionsAmount(v string)`

SetFormattedOptionsAmount sets FormattedOptionsAmount field to given value.

### HasFormattedOptionsAmount

`func (o *LineItemDataAttributes) HasFormattedOptionsAmount() bool`

HasFormattedOptionsAmount returns a boolean if a field has been set.

### GetDiscountCents

`func (o *LineItemDataAttributes) GetDiscountCents() int32`

GetDiscountCents returns the DiscountCents field if non-nil, zero value otherwise.

### GetDiscountCentsOk

`func (o *LineItemDataAttributes) GetDiscountCentsOk() (*int32, bool)`

GetDiscountCentsOk returns a tuple with the DiscountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCents

`func (o *LineItemDataAttributes) SetDiscountCents(v int32)`

SetDiscountCents sets DiscountCents field to given value.

### HasDiscountCents

`func (o *LineItemDataAttributes) HasDiscountCents() bool`

HasDiscountCents returns a boolean if a field has been set.

### GetDiscountFloat

`func (o *LineItemDataAttributes) GetDiscountFloat() float32`

GetDiscountFloat returns the DiscountFloat field if non-nil, zero value otherwise.

### GetDiscountFloatOk

`func (o *LineItemDataAttributes) GetDiscountFloatOk() (*float32, bool)`

GetDiscountFloatOk returns a tuple with the DiscountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountFloat

`func (o *LineItemDataAttributes) SetDiscountFloat(v float32)`

SetDiscountFloat sets DiscountFloat field to given value.

### HasDiscountFloat

`func (o *LineItemDataAttributes) HasDiscountFloat() bool`

HasDiscountFloat returns a boolean if a field has been set.

### GetFormattedDiscount

`func (o *LineItemDataAttributes) GetFormattedDiscount() string`

GetFormattedDiscount returns the FormattedDiscount field if non-nil, zero value otherwise.

### GetFormattedDiscountOk

`func (o *LineItemDataAttributes) GetFormattedDiscountOk() (*string, bool)`

GetFormattedDiscountOk returns a tuple with the FormattedDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedDiscount

`func (o *LineItemDataAttributes) SetFormattedDiscount(v string)`

SetFormattedDiscount sets FormattedDiscount field to given value.

### HasFormattedDiscount

`func (o *LineItemDataAttributes) HasFormattedDiscount() bool`

HasFormattedDiscount returns a boolean if a field has been set.

### GetTotalAmountCents

`func (o *LineItemDataAttributes) GetTotalAmountCents() int32`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *LineItemDataAttributes) GetTotalAmountCentsOk() (*int32, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *LineItemDataAttributes) SetTotalAmountCents(v int32)`

SetTotalAmountCents sets TotalAmountCents field to given value.

### HasTotalAmountCents

`func (o *LineItemDataAttributes) HasTotalAmountCents() bool`

HasTotalAmountCents returns a boolean if a field has been set.

### GetTotalAmountFloat

`func (o *LineItemDataAttributes) GetTotalAmountFloat() float32`

GetTotalAmountFloat returns the TotalAmountFloat field if non-nil, zero value otherwise.

### GetTotalAmountFloatOk

`func (o *LineItemDataAttributes) GetTotalAmountFloatOk() (*float32, bool)`

GetTotalAmountFloatOk returns a tuple with the TotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountFloat

`func (o *LineItemDataAttributes) SetTotalAmountFloat(v float32)`

SetTotalAmountFloat sets TotalAmountFloat field to given value.

### HasTotalAmountFloat

`func (o *LineItemDataAttributes) HasTotalAmountFloat() bool`

HasTotalAmountFloat returns a boolean if a field has been set.

### GetFormattedTotalAmount

`func (o *LineItemDataAttributes) GetFormattedTotalAmount() string`

GetFormattedTotalAmount returns the FormattedTotalAmount field if non-nil, zero value otherwise.

### GetFormattedTotalAmountOk

`func (o *LineItemDataAttributes) GetFormattedTotalAmountOk() (*string, bool)`

GetFormattedTotalAmountOk returns a tuple with the FormattedTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalAmount

`func (o *LineItemDataAttributes) SetFormattedTotalAmount(v string)`

SetFormattedTotalAmount sets FormattedTotalAmount field to given value.

### HasFormattedTotalAmount

`func (o *LineItemDataAttributes) HasFormattedTotalAmount() bool`

HasFormattedTotalAmount returns a boolean if a field has been set.

### GetTaxAmountCents

`func (o *LineItemDataAttributes) GetTaxAmountCents() int32`

GetTaxAmountCents returns the TaxAmountCents field if non-nil, zero value otherwise.

### GetTaxAmountCentsOk

`func (o *LineItemDataAttributes) GetTaxAmountCentsOk() (*int32, bool)`

GetTaxAmountCentsOk returns a tuple with the TaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountCents

`func (o *LineItemDataAttributes) SetTaxAmountCents(v int32)`

SetTaxAmountCents sets TaxAmountCents field to given value.

### HasTaxAmountCents

`func (o *LineItemDataAttributes) HasTaxAmountCents() bool`

HasTaxAmountCents returns a boolean if a field has been set.

### GetTaxAmountFloat

`func (o *LineItemDataAttributes) GetTaxAmountFloat() float32`

GetTaxAmountFloat returns the TaxAmountFloat field if non-nil, zero value otherwise.

### GetTaxAmountFloatOk

`func (o *LineItemDataAttributes) GetTaxAmountFloatOk() (*float32, bool)`

GetTaxAmountFloatOk returns a tuple with the TaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountFloat

`func (o *LineItemDataAttributes) SetTaxAmountFloat(v float32)`

SetTaxAmountFloat sets TaxAmountFloat field to given value.

### HasTaxAmountFloat

`func (o *LineItemDataAttributes) HasTaxAmountFloat() bool`

HasTaxAmountFloat returns a boolean if a field has been set.

### GetFormattedTaxAmount

`func (o *LineItemDataAttributes) GetFormattedTaxAmount() string`

GetFormattedTaxAmount returns the FormattedTaxAmount field if non-nil, zero value otherwise.

### GetFormattedTaxAmountOk

`func (o *LineItemDataAttributes) GetFormattedTaxAmountOk() (*string, bool)`

GetFormattedTaxAmountOk returns a tuple with the FormattedTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTaxAmount

`func (o *LineItemDataAttributes) SetFormattedTaxAmount(v string)`

SetFormattedTaxAmount sets FormattedTaxAmount field to given value.

### HasFormattedTaxAmount

`func (o *LineItemDataAttributes) HasFormattedTaxAmount() bool`

HasFormattedTaxAmount returns a boolean if a field has been set.

### GetName

`func (o *LineItemDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LineItemDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LineItemDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LineItemDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImageUrl

`func (o *LineItemDataAttributes) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *LineItemDataAttributes) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *LineItemDataAttributes) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *LineItemDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetDiscountBreakdown

`func (o *LineItemDataAttributes) GetDiscountBreakdown() map[string]interface{}`

GetDiscountBreakdown returns the DiscountBreakdown field if non-nil, zero value otherwise.

### GetDiscountBreakdownOk

`func (o *LineItemDataAttributes) GetDiscountBreakdownOk() (*map[string]interface{}, bool)`

GetDiscountBreakdownOk returns a tuple with the DiscountBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountBreakdown

`func (o *LineItemDataAttributes) SetDiscountBreakdown(v map[string]interface{})`

SetDiscountBreakdown sets DiscountBreakdown field to given value.

### HasDiscountBreakdown

`func (o *LineItemDataAttributes) HasDiscountBreakdown() bool`

HasDiscountBreakdown returns a boolean if a field has been set.

### GetTaxRate

`func (o *LineItemDataAttributes) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *LineItemDataAttributes) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *LineItemDataAttributes) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *LineItemDataAttributes) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetTaxBreakdown

`func (o *LineItemDataAttributes) GetTaxBreakdown() map[string]interface{}`

GetTaxBreakdown returns the TaxBreakdown field if non-nil, zero value otherwise.

### GetTaxBreakdownOk

`func (o *LineItemDataAttributes) GetTaxBreakdownOk() (*map[string]interface{}, bool)`

GetTaxBreakdownOk returns a tuple with the TaxBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBreakdown

`func (o *LineItemDataAttributes) SetTaxBreakdown(v map[string]interface{})`

SetTaxBreakdown sets TaxBreakdown field to given value.

### HasTaxBreakdown

`func (o *LineItemDataAttributes) HasTaxBreakdown() bool`

HasTaxBreakdown returns a boolean if a field has been set.

### GetItemType

`func (o *LineItemDataAttributes) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *LineItemDataAttributes) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *LineItemDataAttributes) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *LineItemDataAttributes) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetId

`func (o *LineItemDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LineItemDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LineItemDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LineItemDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LineItemDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LineItemDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LineItemDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LineItemDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LineItemDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LineItemDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LineItemDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LineItemDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *LineItemDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *LineItemDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *LineItemDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *LineItemDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *LineItemDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *LineItemDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *LineItemDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *LineItemDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *LineItemDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LineItemDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LineItemDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LineItemDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


