# GETOrders200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | Unique identifier for the order (numeric). | [optional] 
**Autorefresh** | Pointer to **bool** | Save this attribute as &#39;false&#39; if you want prevent the order to be refreshed automatically at each change (much faster). | [optional] 
**Status** | Pointer to **string** | The order status. One of &#39;draft&#39; (default), &#39;pending&#39;, &#39;placed&#39;, &#39;approved&#39;, or &#39;cancelled&#39;. | [optional] 
**PaymentStatus** | Pointer to **string** | The order&#39;s payment status. One of &#39;unpaid&#39; (default), &#39;authorized&#39;, &#39;paid&#39;, &#39;voided&#39;, or &#39;refunded&#39;. | [optional] 
**FulfillmentStatus** | Pointer to **string** | The order&#39;s fulfillment status. One of &#39;unfulfilled&#39; (default), &#39;in_progress&#39;, or &#39;fulfilled&#39;. | [optional] 
**Guest** | Pointer to **bool** | Indicates if the order has been placed as guest. | [optional] 
**Editable** | Pointer to **bool** | Indicates if the order can be edited. | [optional] 
**CustomerEmail** | Pointer to **string** | The email address of the associated customer. When creating or updating an order, this is a shortcut to find or create the associated customer by email. | [optional] 
**LanguageCode** | Pointer to **string** | The preferred language code (ISO 639-1) to be used when communicating with the customer. This can be useful when sending the order to 3rd party marketing tools and CRMs. If the language is supported, the hosted checkout will be localized accordingly. | [optional] 
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard, automatically inherited from the order&#39;s market. | [optional] 
**TaxIncluded** | Pointer to **bool** | Indicates if taxes are included in the order amounts, automatically inherited from the order&#39;s price list. | [optional] 
**TaxRate** | Pointer to **float32** | The tax rate for this order (if calculated). | [optional] 
**FreightTaxable** | Pointer to **bool** | Indicates if taxes are applied to shipping costs. | [optional] 
**RequiresBillingInfo** | Pointer to **bool** | Indicates if the billing address associated to this order requires billing info to be present. | [optional] 
**CountryCode** | Pointer to **string** | The international 2-letter country code as defined by the ISO 3166-1 standard, automatically inherited from the order&#39;s shipping address. | [optional] 
**ShippingCountryCodeLock** | Pointer to **string** | The country code that you want the shipping address to be locked to. This can be useful to make sure the shipping address belongs to a given shipping country, e.g. the one selected in a country selector page. | [optional] 
**CouponCode** | Pointer to **string** | The coupon code to be used for the order. If valid, it triggers a promotion adding a discount line item to the order. | [optional] 
**GiftCardCode** | Pointer to **string** | The gift card code (at least the first 8 characters) to be used for the order. If valid, it uses the gift card balance to pay for the order. | [optional] 
**GiftCardOrCouponCode** | Pointer to **string** | The gift card or coupon code (at least the first 8 characters) to be used for the order. If a gift card mathes, it uses the gift card balance to pay for the order. Otherwise it tries to find a valid coupon code and applies the associated discount. | [optional] 
**SubtotalAmountCents** | Pointer to **int32** | The sum of all the SKU line items total amounts, in cents. | [optional] 
**SubtotalAmountFloat** | Pointer to **float32** | The sum of all the SKU line items total amounts, float. | [optional] 
**FormattedSubtotalAmount** | Pointer to **string** | The sum of all the SKU line items total amounts, formatted. | [optional] 
**ShippingAmountCents** | Pointer to **int32** | The sum of all the shipping costs, in cents. | [optional] 
**ShippingAmountFloat** | Pointer to **float32** | The sum of all the shipping costs, float. | [optional] 
**FormattedShippingAmount** | Pointer to **string** | The sum of all the shipping costs, formatted. | [optional] 
**PaymentMethodAmountCents** | Pointer to **int32** | The payment method costs, in cents. | [optional] 
**PaymentMethodAmountFloat** | Pointer to **float32** | The payment method costs, float. | [optional] 
**FormattedPaymentMethodAmount** | Pointer to **string** | The payment method costs, formatted. | [optional] 
**DiscountAmountCents** | Pointer to **int32** | The sum of all the discounts applied to the order, in cents (negative amount). | [optional] 
**DiscountAmountFloat** | Pointer to **float32** | The sum of all the discounts applied to the order, float. | [optional] 
**FormattedDiscountAmount** | Pointer to **string** | The sum of all the discounts applied to the order, formatted. | [optional] 
**AdjustmentAmountCents** | Pointer to **int32** | The sum of all the adjustments applied to the order, in cents. | [optional] 
**AdjustmentAmountFloat** | Pointer to **float32** | The sum of all the adjustments applied to the order, float. | [optional] 
**FormattedAdjustmentAmount** | Pointer to **string** | The sum of all the adjustments applied to the order, formatted. | [optional] 
**GiftCardAmountCents** | Pointer to **int32** | The sum of all the gift_cards applied to the order, in cents. | [optional] 
**GiftCardAmountFloat** | Pointer to **float32** | The sum of all the gift_cards applied to the order, float. | [optional] 
**FormattedGiftCardAmount** | Pointer to **string** | The sum of all the gift_cards applied to the order, formatted. | [optional] 
**TotalTaxAmountCents** | Pointer to **int32** | The sum of all the taxes applied to the order, in cents. | [optional] 
**TotalTaxAmountFloat** | Pointer to **float32** | The sum of all the taxes applied to the order, float. | [optional] 
**FormattedTotalTaxAmount** | Pointer to **string** | The sum of all the taxes applied to the order, formatted. | [optional] 
**SubtotalTaxAmountCents** | Pointer to **int32** | The taxes applied to the order&#39;s subtotal, in cents. | [optional] 
**SubtotalTaxAmountFloat** | Pointer to **float32** | The taxes applied to the order&#39;s subtotal, float. | [optional] 
**FormattedSubtotalTaxAmount** | Pointer to **string** | The taxes applied to the order&#39;s subtotal, formatted. | [optional] 
**ShippingTaxAmountCents** | Pointer to **int32** | The taxes applied to the order&#39;s shipping costs, in cents. | [optional] 
**ShippingTaxAmountFloat** | Pointer to **float32** | The taxes applied to the order&#39;s shipping costs, float. | [optional] 
**FormattedShippingTaxAmount** | Pointer to **string** | The taxes applied to the order&#39;s shipping costs, formatted. | [optional] 
**PaymentMethodTaxAmountCents** | Pointer to **int32** | The taxes applied to the order&#39;s payment method costs, in cents. | [optional] 
**PaymentMethodTaxAmountFloat** | Pointer to **float32** | The taxes applied to the order&#39;s payment method costs, float. | [optional] 
**FormattedPaymentMethodTaxAmount** | Pointer to **string** | The taxes applied to the order&#39;s payment method costs, formatted. | [optional] 
**AdjustmentTaxAmountCents** | Pointer to **int32** | The taxes applied to the order adjustments, in cents. | [optional] 
**AdjustmentTaxAmountFloat** | Pointer to **float32** | The taxes applied to the order adjustments, float. | [optional] 
**FormattedAdjustmentTaxAmount** | Pointer to **string** | The taxes applied to the order adjustments, formatted. | [optional] 
**TotalAmountCents** | Pointer to **int32** | The order&#39;s total amount, in cents. | [optional] 
**TotalAmountFloat** | Pointer to **float32** | The order&#39;s total amount, float. | [optional] 
**FormattedTotalAmount** | Pointer to **string** | The order&#39;s total amount, formatted. | [optional] 
**TotalTaxableAmountCents** | Pointer to **int32** | The order&#39;s total taxable amount, in cents (without discounts). | [optional] 
**TotalTaxableAmountFloat** | Pointer to **float32** | The order&#39;s total taxable amount, float. | [optional] 
**FormattedTotalTaxableAmount** | Pointer to **string** | The order&#39;s total taxable amount, formatted. | [optional] 
**SubtotalTaxableAmountCents** | Pointer to **int32** | The order&#39;s subtotal taxable amount, in cents (equal to subtotal_amount_cents when prices don&#39;t include taxes). | [optional] 
**SubtotalTaxableAmountFloat** | Pointer to **float32** | The order&#39;s subtotal taxable amount, float. | [optional] 
**FormattedSubtotalTaxableAmount** | Pointer to **string** | The order&#39;s subtotal taxable amount, formatted. | [optional] 
**ShippingTaxableAmountCents** | Pointer to **int32** | The order&#39;s shipping taxable amount, in cents (equal to shipping_amount_cents when prices don&#39;t include taxes). | [optional] 
**ShippingTaxableAmountFloat** | Pointer to **float32** | The order&#39;s shipping taxable amount, float. | [optional] 
**FormattedShippingTaxableAmount** | Pointer to **string** | The order&#39;s shipping taxable amount, formatted. | [optional] 
**PaymentMethodTaxableAmountCents** | Pointer to **int32** | The order&#39;s payment method taxable amount, in cents (equal to payment_method_amount_cents when prices don&#39;t include taxes). | [optional] 
**PaymentMethodTaxableAmountFloat** | Pointer to **float32** | The order&#39;s payment method taxable amount, float. | [optional] 
**FormattedPaymentMethodTaxableAmount** | Pointer to **string** | The order&#39;s payment method taxable amount, formatted. | [optional] 
**AdjustmentTaxableAmountCents** | Pointer to **int32** | The order&#39;s adjustment taxable amount, in cents (equal to discount_adjustment_cents when prices don&#39;t include taxes). | [optional] 
**AdjustmentTaxableAmountFloat** | Pointer to **float32** | The order&#39;s adjustment taxable amount, float. | [optional] 
**FormattedAdjustmentTaxableAmount** | Pointer to **string** | The order&#39;s adjustment taxable amount, formatted. | [optional] 
**TotalAmountWithTaxesCents** | Pointer to **int32** | The order&#39;s total amount (when prices include taxes) or the order&#39;s total + taxes amount (when prices don&#39;t include taxes, e.g. US Markets or B2B). | [optional] 
**TotalAmountWithTaxesFloat** | Pointer to **float32** | The order&#39;s total amount with taxes, float. | [optional] 
**FormattedTotalAmountWithTaxes** | Pointer to **string** | The order&#39;s total amount with taxes, formatted. | [optional] 
**FeesAmountCents** | Pointer to **int32** | The fees amount that is applied by Commerce Layer, in cents. | [optional] 
**FeesAmountFloat** | Pointer to **float32** | The fees amount that is applied by Commerce Layer, float. | [optional] 
**FormattedFeesAmount** | Pointer to **string** | The fees amount that is applied by Commerce Layer, formatted. | [optional] 
**DutyAmountCents** | Pointer to **int32** | The duty amount that is calculated by external services, in cents. | [optional] 
**DutyAmountFloat** | Pointer to **float32** | The duty amount that is calculated by external services, float. | [optional] 
**FormattedDutyAmount** | Pointer to **string** | The duty amount that is calculated by external services, formatted. | [optional] 
**SkusCount** | Pointer to **int32** | The total number of SKUs in the order&#39;s line items. This can be useful to display a preview of the customer shopping cart content. | [optional] 
**LineItemOptionsCount** | Pointer to **int32** | The total number of line item options. This can be useful to display a preview of the customer shopping cart content. | [optional] 
**ShipmentsCount** | Pointer to **int32** | The total number of shipments. This can be useful to manage the shipping method(s) selection during checkout. | [optional] 
**TaxCalculationsCount** | Pointer to **int32** | The total number of tax calculations. This can be useful to monitor external tax service usage. | [optional] 
**PaymentSourceDetails** | Pointer to **map[string]interface{}** | An object that contains the shareable details of the order&#39;s payment source. | [optional] 
**Token** | Pointer to **string** | A unique token that can be shared more securely instead of the order&#39;s id. | [optional] 
**CartUrl** | Pointer to **string** | The cart url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**ReturnUrl** | Pointer to **string** | The return url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**TermsUrl** | Pointer to **string** | The terms and conditions url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**PrivacyUrl** | Pointer to **string** | The privacy policy url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**CheckoutUrl** | Pointer to **string** | The checkout url that was automatically generated for the order. Send the customers to this url to let them checkout the order securely on our hosted checkout application. | [optional] 
**PlacedAt** | Pointer to **string** | Time at which the order was placed. | [optional] 
**ApprovedAt** | Pointer to **string** | Time at which the order was approved. | [optional] 
**CancelledAt** | Pointer to **string** | Time at which the order was cancelled. | [optional] 
**PaymentUpdatedAt** | Pointer to **string** | Time at which the order&#39;s payment status was last updated. | [optional] 
**FulfillmentUpdatedAt** | Pointer to **string** | Time at which the order&#39;s fulfillment status was last updated. | [optional] 
**RefreshedAt** | Pointer to **string** | Last time at which an order was manually refreshed. | [optional] 
**ArchivedAt** | Pointer to **string** | Time at which the resource has been archived. | [optional] 
**ExpiresAt** | Pointer to **string** | Time at which an order is marked for cleanup. Any order will start with a default expire time of 2 months. Expiration is reset once a line item is added to the order. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETOrders200ResponseDataInnerAttributes

`func NewGETOrders200ResponseDataInnerAttributes() *GETOrders200ResponseDataInnerAttributes`

NewGETOrders200ResponseDataInnerAttributes instantiates a new GETOrders200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETOrders200ResponseDataInnerAttributesWithDefaults

`func NewGETOrders200ResponseDataInnerAttributesWithDefaults() *GETOrders200ResponseDataInnerAttributes`

NewGETOrders200ResponseDataInnerAttributesWithDefaults instantiates a new GETOrders200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETOrders200ResponseDataInnerAttributes) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETOrders200ResponseDataInnerAttributes) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETOrders200ResponseDataInnerAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetAutorefresh

`func (o *GETOrders200ResponseDataInnerAttributes) GetAutorefresh() bool`

GetAutorefresh returns the Autorefresh field if non-nil, zero value otherwise.

### GetAutorefreshOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetAutorefreshOk() (*bool, bool)`

GetAutorefreshOk returns a tuple with the Autorefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutorefresh

`func (o *GETOrders200ResponseDataInnerAttributes) SetAutorefresh(v bool)`

SetAutorefresh sets Autorefresh field to given value.

### HasAutorefresh

`func (o *GETOrders200ResponseDataInnerAttributes) HasAutorefresh() bool`

HasAutorefresh returns a boolean if a field has been set.

### GetStatus

`func (o *GETOrders200ResponseDataInnerAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETOrders200ResponseDataInnerAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETOrders200ResponseDataInnerAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetFulfillmentStatus

`func (o *GETOrders200ResponseDataInnerAttributes) GetFulfillmentStatus() string`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFulfillmentStatusOk() (*string, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *GETOrders200ResponseDataInnerAttributes) SetFulfillmentStatus(v string)`

SetFulfillmentStatus sets FulfillmentStatus field to given value.

### HasFulfillmentStatus

`func (o *GETOrders200ResponseDataInnerAttributes) HasFulfillmentStatus() bool`

HasFulfillmentStatus returns a boolean if a field has been set.

### GetGuest

`func (o *GETOrders200ResponseDataInnerAttributes) GetGuest() bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetGuestOk() (*bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *GETOrders200ResponseDataInnerAttributes) SetGuest(v bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *GETOrders200ResponseDataInnerAttributes) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetEditable

`func (o *GETOrders200ResponseDataInnerAttributes) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *GETOrders200ResponseDataInnerAttributes) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *GETOrders200ResponseDataInnerAttributes) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *GETOrders200ResponseDataInnerAttributes) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *GETOrders200ResponseDataInnerAttributes) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *GETOrders200ResponseDataInnerAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetLanguageCode

`func (o *GETOrders200ResponseDataInnerAttributes) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *GETOrders200ResponseDataInnerAttributes) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *GETOrders200ResponseDataInnerAttributes) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GETOrders200ResponseDataInnerAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETOrders200ResponseDataInnerAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETOrders200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetTaxIncluded

`func (o *GETOrders200ResponseDataInnerAttributes) GetTaxIncluded() bool`

GetTaxIncluded returns the TaxIncluded field if non-nil, zero value otherwise.

### GetTaxIncludedOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTaxIncludedOk() (*bool, bool)`

GetTaxIncludedOk returns a tuple with the TaxIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncluded

`func (o *GETOrders200ResponseDataInnerAttributes) SetTaxIncluded(v bool)`

SetTaxIncluded sets TaxIncluded field to given value.

### HasTaxIncluded

`func (o *GETOrders200ResponseDataInnerAttributes) HasTaxIncluded() bool`

HasTaxIncluded returns a boolean if a field has been set.

### GetTaxRate

`func (o *GETOrders200ResponseDataInnerAttributes) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *GETOrders200ResponseDataInnerAttributes) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *GETOrders200ResponseDataInnerAttributes) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetFreightTaxable

`func (o *GETOrders200ResponseDataInnerAttributes) GetFreightTaxable() bool`

GetFreightTaxable returns the FreightTaxable field if non-nil, zero value otherwise.

### GetFreightTaxableOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFreightTaxableOk() (*bool, bool)`

GetFreightTaxableOk returns a tuple with the FreightTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTaxable

`func (o *GETOrders200ResponseDataInnerAttributes) SetFreightTaxable(v bool)`

SetFreightTaxable sets FreightTaxable field to given value.

### HasFreightTaxable

`func (o *GETOrders200ResponseDataInnerAttributes) HasFreightTaxable() bool`

HasFreightTaxable returns a boolean if a field has been set.

### GetRequiresBillingInfo

`func (o *GETOrders200ResponseDataInnerAttributes) GetRequiresBillingInfo() bool`

GetRequiresBillingInfo returns the RequiresBillingInfo field if non-nil, zero value otherwise.

### GetRequiresBillingInfoOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetRequiresBillingInfoOk() (*bool, bool)`

GetRequiresBillingInfoOk returns a tuple with the RequiresBillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresBillingInfo

`func (o *GETOrders200ResponseDataInnerAttributes) SetRequiresBillingInfo(v bool)`

SetRequiresBillingInfo sets RequiresBillingInfo field to given value.

### HasRequiresBillingInfo

`func (o *GETOrders200ResponseDataInnerAttributes) HasRequiresBillingInfo() bool`

HasRequiresBillingInfo returns a boolean if a field has been set.

### GetCountryCode

`func (o *GETOrders200ResponseDataInnerAttributes) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *GETOrders200ResponseDataInnerAttributes) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *GETOrders200ResponseDataInnerAttributes) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetShippingCountryCodeLock

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingCountryCodeLock() string`

GetShippingCountryCodeLock returns the ShippingCountryCodeLock field if non-nil, zero value otherwise.

### GetShippingCountryCodeLockOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingCountryCodeLockOk() (*string, bool)`

GetShippingCountryCodeLockOk returns a tuple with the ShippingCountryCodeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCountryCodeLock

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingCountryCodeLock(v string)`

SetShippingCountryCodeLock sets ShippingCountryCodeLock field to given value.

### HasShippingCountryCodeLock

`func (o *GETOrders200ResponseDataInnerAttributes) HasShippingCountryCodeLock() bool`

HasShippingCountryCodeLock returns a boolean if a field has been set.

### GetCouponCode

`func (o *GETOrders200ResponseDataInnerAttributes) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *GETOrders200ResponseDataInnerAttributes) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *GETOrders200ResponseDataInnerAttributes) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.

### GetGiftCardCode

`func (o *GETOrders200ResponseDataInnerAttributes) GetGiftCardCode() string`

GetGiftCardCode returns the GiftCardCode field if non-nil, zero value otherwise.

### GetGiftCardCodeOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetGiftCardCodeOk() (*string, bool)`

GetGiftCardCodeOk returns a tuple with the GiftCardCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardCode

`func (o *GETOrders200ResponseDataInnerAttributes) SetGiftCardCode(v string)`

SetGiftCardCode sets GiftCardCode field to given value.

### HasGiftCardCode

`func (o *GETOrders200ResponseDataInnerAttributes) HasGiftCardCode() bool`

HasGiftCardCode returns a boolean if a field has been set.

### GetGiftCardOrCouponCode

`func (o *GETOrders200ResponseDataInnerAttributes) GetGiftCardOrCouponCode() string`

GetGiftCardOrCouponCode returns the GiftCardOrCouponCode field if non-nil, zero value otherwise.

### GetGiftCardOrCouponCodeOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetGiftCardOrCouponCodeOk() (*string, bool)`

GetGiftCardOrCouponCodeOk returns a tuple with the GiftCardOrCouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardOrCouponCode

`func (o *GETOrders200ResponseDataInnerAttributes) SetGiftCardOrCouponCode(v string)`

SetGiftCardOrCouponCode sets GiftCardOrCouponCode field to given value.

### HasGiftCardOrCouponCode

`func (o *GETOrders200ResponseDataInnerAttributes) HasGiftCardOrCouponCode() bool`

HasGiftCardOrCouponCode returns a boolean if a field has been set.

### GetSubtotalAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalAmountCents() int32`

GetSubtotalAmountCents returns the SubtotalAmountCents field if non-nil, zero value otherwise.

### GetSubtotalAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalAmountCentsOk() (*int32, bool)`

GetSubtotalAmountCentsOk returns a tuple with the SubtotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalAmountCents(v int32)`

SetSubtotalAmountCents sets SubtotalAmountCents field to given value.

### HasSubtotalAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasSubtotalAmountCents() bool`

HasSubtotalAmountCents returns a boolean if a field has been set.

### GetSubtotalAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalAmountFloat() float32`

GetSubtotalAmountFloat returns the SubtotalAmountFloat field if non-nil, zero value otherwise.

### GetSubtotalAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalAmountFloatOk() (*float32, bool)`

GetSubtotalAmountFloatOk returns a tuple with the SubtotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalAmountFloat(v float32)`

SetSubtotalAmountFloat sets SubtotalAmountFloat field to given value.

### HasSubtotalAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasSubtotalAmountFloat() bool`

HasSubtotalAmountFloat returns a boolean if a field has been set.

### GetFormattedSubtotalAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedSubtotalAmount() string`

GetFormattedSubtotalAmount returns the FormattedSubtotalAmount field if non-nil, zero value otherwise.

### GetFormattedSubtotalAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedSubtotalAmountOk() (*string, bool)`

GetFormattedSubtotalAmountOk returns a tuple with the FormattedSubtotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSubtotalAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedSubtotalAmount(v string)`

SetFormattedSubtotalAmount sets FormattedSubtotalAmount field to given value.

### HasFormattedSubtotalAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedSubtotalAmount() bool`

HasFormattedSubtotalAmount returns a boolean if a field has been set.

### GetShippingAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingAmountCents() int32`

GetShippingAmountCents returns the ShippingAmountCents field if non-nil, zero value otherwise.

### GetShippingAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingAmountCentsOk() (*int32, bool)`

GetShippingAmountCentsOk returns a tuple with the ShippingAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingAmountCents(v int32)`

SetShippingAmountCents sets ShippingAmountCents field to given value.

### HasShippingAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasShippingAmountCents() bool`

HasShippingAmountCents returns a boolean if a field has been set.

### GetShippingAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingAmountFloat() float32`

GetShippingAmountFloat returns the ShippingAmountFloat field if non-nil, zero value otherwise.

### GetShippingAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingAmountFloatOk() (*float32, bool)`

GetShippingAmountFloatOk returns a tuple with the ShippingAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingAmountFloat(v float32)`

SetShippingAmountFloat sets ShippingAmountFloat field to given value.

### HasShippingAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasShippingAmountFloat() bool`

HasShippingAmountFloat returns a boolean if a field has been set.

### GetFormattedShippingAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedShippingAmount() string`

GetFormattedShippingAmount returns the FormattedShippingAmount field if non-nil, zero value otherwise.

### GetFormattedShippingAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedShippingAmountOk() (*string, bool)`

GetFormattedShippingAmountOk returns a tuple with the FormattedShippingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedShippingAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedShippingAmount(v string)`

SetFormattedShippingAmount sets FormattedShippingAmount field to given value.

### HasFormattedShippingAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedShippingAmount() bool`

HasFormattedShippingAmount returns a boolean if a field has been set.

### GetPaymentMethodAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodAmountCents() int32`

GetPaymentMethodAmountCents returns the PaymentMethodAmountCents field if non-nil, zero value otherwise.

### GetPaymentMethodAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodAmountCentsOk() (*int32, bool)`

GetPaymentMethodAmountCentsOk returns a tuple with the PaymentMethodAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodAmountCents(v int32)`

SetPaymentMethodAmountCents sets PaymentMethodAmountCents field to given value.

### HasPaymentMethodAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentMethodAmountCents() bool`

HasPaymentMethodAmountCents returns a boolean if a field has been set.

### GetPaymentMethodAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodAmountFloat() float32`

GetPaymentMethodAmountFloat returns the PaymentMethodAmountFloat field if non-nil, zero value otherwise.

### GetPaymentMethodAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodAmountFloatOk() (*float32, bool)`

GetPaymentMethodAmountFloatOk returns a tuple with the PaymentMethodAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodAmountFloat(v float32)`

SetPaymentMethodAmountFloat sets PaymentMethodAmountFloat field to given value.

### HasPaymentMethodAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentMethodAmountFloat() bool`

HasPaymentMethodAmountFloat returns a boolean if a field has been set.

### GetFormattedPaymentMethodAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedPaymentMethodAmount() string`

GetFormattedPaymentMethodAmount returns the FormattedPaymentMethodAmount field if non-nil, zero value otherwise.

### GetFormattedPaymentMethodAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedPaymentMethodAmountOk() (*string, bool)`

GetFormattedPaymentMethodAmountOk returns a tuple with the FormattedPaymentMethodAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPaymentMethodAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedPaymentMethodAmount(v string)`

SetFormattedPaymentMethodAmount sets FormattedPaymentMethodAmount field to given value.

### HasFormattedPaymentMethodAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedPaymentMethodAmount() bool`

HasFormattedPaymentMethodAmount returns a boolean if a field has been set.

### GetDiscountAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetDiscountAmountCents() int32`

GetDiscountAmountCents returns the DiscountAmountCents field if non-nil, zero value otherwise.

### GetDiscountAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetDiscountAmountCentsOk() (*int32, bool)`

GetDiscountAmountCentsOk returns a tuple with the DiscountAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetDiscountAmountCents(v int32)`

SetDiscountAmountCents sets DiscountAmountCents field to given value.

### HasDiscountAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasDiscountAmountCents() bool`

HasDiscountAmountCents returns a boolean if a field has been set.

### GetDiscountAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetDiscountAmountFloat() float32`

GetDiscountAmountFloat returns the DiscountAmountFloat field if non-nil, zero value otherwise.

### GetDiscountAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetDiscountAmountFloatOk() (*float32, bool)`

GetDiscountAmountFloatOk returns a tuple with the DiscountAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetDiscountAmountFloat(v float32)`

SetDiscountAmountFloat sets DiscountAmountFloat field to given value.

### HasDiscountAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasDiscountAmountFloat() bool`

HasDiscountAmountFloat returns a boolean if a field has been set.

### GetFormattedDiscountAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedDiscountAmount() string`

GetFormattedDiscountAmount returns the FormattedDiscountAmount field if non-nil, zero value otherwise.

### GetFormattedDiscountAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedDiscountAmountOk() (*string, bool)`

GetFormattedDiscountAmountOk returns a tuple with the FormattedDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedDiscountAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedDiscountAmount(v string)`

SetFormattedDiscountAmount sets FormattedDiscountAmount field to given value.

### HasFormattedDiscountAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedDiscountAmount() bool`

HasFormattedDiscountAmount returns a boolean if a field has been set.

### GetAdjustmentAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentAmountCents() int32`

GetAdjustmentAmountCents returns the AdjustmentAmountCents field if non-nil, zero value otherwise.

### GetAdjustmentAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentAmountCentsOk() (*int32, bool)`

GetAdjustmentAmountCentsOk returns a tuple with the AdjustmentAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentAmountCents(v int32)`

SetAdjustmentAmountCents sets AdjustmentAmountCents field to given value.

### HasAdjustmentAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasAdjustmentAmountCents() bool`

HasAdjustmentAmountCents returns a boolean if a field has been set.

### GetAdjustmentAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentAmountFloat() float32`

GetAdjustmentAmountFloat returns the AdjustmentAmountFloat field if non-nil, zero value otherwise.

### GetAdjustmentAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentAmountFloatOk() (*float32, bool)`

GetAdjustmentAmountFloatOk returns a tuple with the AdjustmentAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentAmountFloat(v float32)`

SetAdjustmentAmountFloat sets AdjustmentAmountFloat field to given value.

### HasAdjustmentAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasAdjustmentAmountFloat() bool`

HasAdjustmentAmountFloat returns a boolean if a field has been set.

### GetFormattedAdjustmentAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedAdjustmentAmount() string`

GetFormattedAdjustmentAmount returns the FormattedAdjustmentAmount field if non-nil, zero value otherwise.

### GetFormattedAdjustmentAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedAdjustmentAmountOk() (*string, bool)`

GetFormattedAdjustmentAmountOk returns a tuple with the FormattedAdjustmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAdjustmentAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedAdjustmentAmount(v string)`

SetFormattedAdjustmentAmount sets FormattedAdjustmentAmount field to given value.

### HasFormattedAdjustmentAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedAdjustmentAmount() bool`

HasFormattedAdjustmentAmount returns a boolean if a field has been set.

### GetGiftCardAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetGiftCardAmountCents() int32`

GetGiftCardAmountCents returns the GiftCardAmountCents field if non-nil, zero value otherwise.

### GetGiftCardAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetGiftCardAmountCentsOk() (*int32, bool)`

GetGiftCardAmountCentsOk returns a tuple with the GiftCardAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetGiftCardAmountCents(v int32)`

SetGiftCardAmountCents sets GiftCardAmountCents field to given value.

### HasGiftCardAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasGiftCardAmountCents() bool`

HasGiftCardAmountCents returns a boolean if a field has been set.

### GetGiftCardAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetGiftCardAmountFloat() float32`

GetGiftCardAmountFloat returns the GiftCardAmountFloat field if non-nil, zero value otherwise.

### GetGiftCardAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetGiftCardAmountFloatOk() (*float32, bool)`

GetGiftCardAmountFloatOk returns a tuple with the GiftCardAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetGiftCardAmountFloat(v float32)`

SetGiftCardAmountFloat sets GiftCardAmountFloat field to given value.

### HasGiftCardAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasGiftCardAmountFloat() bool`

HasGiftCardAmountFloat returns a boolean if a field has been set.

### GetFormattedGiftCardAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedGiftCardAmount() string`

GetFormattedGiftCardAmount returns the FormattedGiftCardAmount field if non-nil, zero value otherwise.

### GetFormattedGiftCardAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedGiftCardAmountOk() (*string, bool)`

GetFormattedGiftCardAmountOk returns a tuple with the FormattedGiftCardAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedGiftCardAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedGiftCardAmount(v string)`

SetFormattedGiftCardAmount sets FormattedGiftCardAmount field to given value.

### HasFormattedGiftCardAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedGiftCardAmount() bool`

HasFormattedGiftCardAmount returns a boolean if a field has been set.

### GetTotalTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalTaxAmountCents() int32`

GetTotalTaxAmountCents returns the TotalTaxAmountCents field if non-nil, zero value otherwise.

### GetTotalTaxAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalTaxAmountCentsOk() (*int32, bool)`

GetTotalTaxAmountCentsOk returns a tuple with the TotalTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalTaxAmountCents(v int32)`

SetTotalTaxAmountCents sets TotalTaxAmountCents field to given value.

### HasTotalTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasTotalTaxAmountCents() bool`

HasTotalTaxAmountCents returns a boolean if a field has been set.

### GetTotalTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalTaxAmountFloat() float32`

GetTotalTaxAmountFloat returns the TotalTaxAmountFloat field if non-nil, zero value otherwise.

### GetTotalTaxAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalTaxAmountFloatOk() (*float32, bool)`

GetTotalTaxAmountFloatOk returns a tuple with the TotalTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalTaxAmountFloat(v float32)`

SetTotalTaxAmountFloat sets TotalTaxAmountFloat field to given value.

### HasTotalTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasTotalTaxAmountFloat() bool`

HasTotalTaxAmountFloat returns a boolean if a field has been set.

### GetFormattedTotalTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedTotalTaxAmount() string`

GetFormattedTotalTaxAmount returns the FormattedTotalTaxAmount field if non-nil, zero value otherwise.

### GetFormattedTotalTaxAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedTotalTaxAmountOk() (*string, bool)`

GetFormattedTotalTaxAmountOk returns a tuple with the FormattedTotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedTotalTaxAmount(v string)`

SetFormattedTotalTaxAmount sets FormattedTotalTaxAmount field to given value.

### HasFormattedTotalTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedTotalTaxAmount() bool`

HasFormattedTotalTaxAmount returns a boolean if a field has been set.

### GetSubtotalTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalTaxAmountCents() int32`

GetSubtotalTaxAmountCents returns the SubtotalTaxAmountCents field if non-nil, zero value otherwise.

### GetSubtotalTaxAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalTaxAmountCentsOk() (*int32, bool)`

GetSubtotalTaxAmountCentsOk returns a tuple with the SubtotalTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalTaxAmountCents(v int32)`

SetSubtotalTaxAmountCents sets SubtotalTaxAmountCents field to given value.

### HasSubtotalTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasSubtotalTaxAmountCents() bool`

HasSubtotalTaxAmountCents returns a boolean if a field has been set.

### GetSubtotalTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalTaxAmountFloat() float32`

GetSubtotalTaxAmountFloat returns the SubtotalTaxAmountFloat field if non-nil, zero value otherwise.

### GetSubtotalTaxAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalTaxAmountFloatOk() (*float32, bool)`

GetSubtotalTaxAmountFloatOk returns a tuple with the SubtotalTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalTaxAmountFloat(v float32)`

SetSubtotalTaxAmountFloat sets SubtotalTaxAmountFloat field to given value.

### HasSubtotalTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasSubtotalTaxAmountFloat() bool`

HasSubtotalTaxAmountFloat returns a boolean if a field has been set.

### GetFormattedSubtotalTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedSubtotalTaxAmount() string`

GetFormattedSubtotalTaxAmount returns the FormattedSubtotalTaxAmount field if non-nil, zero value otherwise.

### GetFormattedSubtotalTaxAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedSubtotalTaxAmountOk() (*string, bool)`

GetFormattedSubtotalTaxAmountOk returns a tuple with the FormattedSubtotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSubtotalTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedSubtotalTaxAmount(v string)`

SetFormattedSubtotalTaxAmount sets FormattedSubtotalTaxAmount field to given value.

### HasFormattedSubtotalTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedSubtotalTaxAmount() bool`

HasFormattedSubtotalTaxAmount returns a boolean if a field has been set.

### GetShippingTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingTaxAmountCents() int32`

GetShippingTaxAmountCents returns the ShippingTaxAmountCents field if non-nil, zero value otherwise.

### GetShippingTaxAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingTaxAmountCentsOk() (*int32, bool)`

GetShippingTaxAmountCentsOk returns a tuple with the ShippingTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingTaxAmountCents(v int32)`

SetShippingTaxAmountCents sets ShippingTaxAmountCents field to given value.

### HasShippingTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasShippingTaxAmountCents() bool`

HasShippingTaxAmountCents returns a boolean if a field has been set.

### GetShippingTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingTaxAmountFloat() float32`

GetShippingTaxAmountFloat returns the ShippingTaxAmountFloat field if non-nil, zero value otherwise.

### GetShippingTaxAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingTaxAmountFloatOk() (*float32, bool)`

GetShippingTaxAmountFloatOk returns a tuple with the ShippingTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingTaxAmountFloat(v float32)`

SetShippingTaxAmountFloat sets ShippingTaxAmountFloat field to given value.

### HasShippingTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasShippingTaxAmountFloat() bool`

HasShippingTaxAmountFloat returns a boolean if a field has been set.

### GetFormattedShippingTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedShippingTaxAmount() string`

GetFormattedShippingTaxAmount returns the FormattedShippingTaxAmount field if non-nil, zero value otherwise.

### GetFormattedShippingTaxAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedShippingTaxAmountOk() (*string, bool)`

GetFormattedShippingTaxAmountOk returns a tuple with the FormattedShippingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedShippingTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedShippingTaxAmount(v string)`

SetFormattedShippingTaxAmount sets FormattedShippingTaxAmount field to given value.

### HasFormattedShippingTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedShippingTaxAmount() bool`

HasFormattedShippingTaxAmount returns a boolean if a field has been set.

### GetPaymentMethodTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodTaxAmountCents() int32`

GetPaymentMethodTaxAmountCents returns the PaymentMethodTaxAmountCents field if non-nil, zero value otherwise.

### GetPaymentMethodTaxAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodTaxAmountCentsOk() (*int32, bool)`

GetPaymentMethodTaxAmountCentsOk returns a tuple with the PaymentMethodTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodTaxAmountCents(v int32)`

SetPaymentMethodTaxAmountCents sets PaymentMethodTaxAmountCents field to given value.

### HasPaymentMethodTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentMethodTaxAmountCents() bool`

HasPaymentMethodTaxAmountCents returns a boolean if a field has been set.

### GetPaymentMethodTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodTaxAmountFloat() float32`

GetPaymentMethodTaxAmountFloat returns the PaymentMethodTaxAmountFloat field if non-nil, zero value otherwise.

### GetPaymentMethodTaxAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodTaxAmountFloatOk() (*float32, bool)`

GetPaymentMethodTaxAmountFloatOk returns a tuple with the PaymentMethodTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodTaxAmountFloat(v float32)`

SetPaymentMethodTaxAmountFloat sets PaymentMethodTaxAmountFloat field to given value.

### HasPaymentMethodTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentMethodTaxAmountFloat() bool`

HasPaymentMethodTaxAmountFloat returns a boolean if a field has been set.

### GetFormattedPaymentMethodTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedPaymentMethodTaxAmount() string`

GetFormattedPaymentMethodTaxAmount returns the FormattedPaymentMethodTaxAmount field if non-nil, zero value otherwise.

### GetFormattedPaymentMethodTaxAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedPaymentMethodTaxAmountOk() (*string, bool)`

GetFormattedPaymentMethodTaxAmountOk returns a tuple with the FormattedPaymentMethodTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPaymentMethodTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedPaymentMethodTaxAmount(v string)`

SetFormattedPaymentMethodTaxAmount sets FormattedPaymentMethodTaxAmount field to given value.

### HasFormattedPaymentMethodTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedPaymentMethodTaxAmount() bool`

HasFormattedPaymentMethodTaxAmount returns a boolean if a field has been set.

### GetAdjustmentTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentTaxAmountCents() int32`

GetAdjustmentTaxAmountCents returns the AdjustmentTaxAmountCents field if non-nil, zero value otherwise.

### GetAdjustmentTaxAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentTaxAmountCentsOk() (*int32, bool)`

GetAdjustmentTaxAmountCentsOk returns a tuple with the AdjustmentTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentTaxAmountCents(v int32)`

SetAdjustmentTaxAmountCents sets AdjustmentTaxAmountCents field to given value.

### HasAdjustmentTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasAdjustmentTaxAmountCents() bool`

HasAdjustmentTaxAmountCents returns a boolean if a field has been set.

### GetAdjustmentTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentTaxAmountFloat() float32`

GetAdjustmentTaxAmountFloat returns the AdjustmentTaxAmountFloat field if non-nil, zero value otherwise.

### GetAdjustmentTaxAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentTaxAmountFloatOk() (*float32, bool)`

GetAdjustmentTaxAmountFloatOk returns a tuple with the AdjustmentTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentTaxAmountFloat(v float32)`

SetAdjustmentTaxAmountFloat sets AdjustmentTaxAmountFloat field to given value.

### HasAdjustmentTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasAdjustmentTaxAmountFloat() bool`

HasAdjustmentTaxAmountFloat returns a boolean if a field has been set.

### GetFormattedAdjustmentTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedAdjustmentTaxAmount() string`

GetFormattedAdjustmentTaxAmount returns the FormattedAdjustmentTaxAmount field if non-nil, zero value otherwise.

### GetFormattedAdjustmentTaxAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedAdjustmentTaxAmountOk() (*string, bool)`

GetFormattedAdjustmentTaxAmountOk returns a tuple with the FormattedAdjustmentTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAdjustmentTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedAdjustmentTaxAmount(v string)`

SetFormattedAdjustmentTaxAmount sets FormattedAdjustmentTaxAmount field to given value.

### HasFormattedAdjustmentTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedAdjustmentTaxAmount() bool`

HasFormattedAdjustmentTaxAmount returns a boolean if a field has been set.

### GetTotalAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalAmountCents() int32`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalAmountCentsOk() (*int32, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalAmountCents(v int32)`

SetTotalAmountCents sets TotalAmountCents field to given value.

### HasTotalAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasTotalAmountCents() bool`

HasTotalAmountCents returns a boolean if a field has been set.

### GetTotalAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalAmountFloat() float32`

GetTotalAmountFloat returns the TotalAmountFloat field if non-nil, zero value otherwise.

### GetTotalAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalAmountFloatOk() (*float32, bool)`

GetTotalAmountFloatOk returns a tuple with the TotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalAmountFloat(v float32)`

SetTotalAmountFloat sets TotalAmountFloat field to given value.

### HasTotalAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasTotalAmountFloat() bool`

HasTotalAmountFloat returns a boolean if a field has been set.

### GetFormattedTotalAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedTotalAmount() string`

GetFormattedTotalAmount returns the FormattedTotalAmount field if non-nil, zero value otherwise.

### GetFormattedTotalAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedTotalAmountOk() (*string, bool)`

GetFormattedTotalAmountOk returns a tuple with the FormattedTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedTotalAmount(v string)`

SetFormattedTotalAmount sets FormattedTotalAmount field to given value.

### HasFormattedTotalAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedTotalAmount() bool`

HasFormattedTotalAmount returns a boolean if a field has been set.

### GetTotalTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalTaxableAmountCents() int32`

GetTotalTaxableAmountCents returns the TotalTaxableAmountCents field if non-nil, zero value otherwise.

### GetTotalTaxableAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalTaxableAmountCentsOk() (*int32, bool)`

GetTotalTaxableAmountCentsOk returns a tuple with the TotalTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalTaxableAmountCents(v int32)`

SetTotalTaxableAmountCents sets TotalTaxableAmountCents field to given value.

### HasTotalTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasTotalTaxableAmountCents() bool`

HasTotalTaxableAmountCents returns a boolean if a field has been set.

### GetTotalTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalTaxableAmountFloat() float32`

GetTotalTaxableAmountFloat returns the TotalTaxableAmountFloat field if non-nil, zero value otherwise.

### GetTotalTaxableAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalTaxableAmountFloatOk() (*float32, bool)`

GetTotalTaxableAmountFloatOk returns a tuple with the TotalTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalTaxableAmountFloat(v float32)`

SetTotalTaxableAmountFloat sets TotalTaxableAmountFloat field to given value.

### HasTotalTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasTotalTaxableAmountFloat() bool`

HasTotalTaxableAmountFloat returns a boolean if a field has been set.

### GetFormattedTotalTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedTotalTaxableAmount() string`

GetFormattedTotalTaxableAmount returns the FormattedTotalTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedTotalTaxableAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedTotalTaxableAmountOk() (*string, bool)`

GetFormattedTotalTaxableAmountOk returns a tuple with the FormattedTotalTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedTotalTaxableAmount(v string)`

SetFormattedTotalTaxableAmount sets FormattedTotalTaxableAmount field to given value.

### HasFormattedTotalTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedTotalTaxableAmount() bool`

HasFormattedTotalTaxableAmount returns a boolean if a field has been set.

### GetSubtotalTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalTaxableAmountCents() int32`

GetSubtotalTaxableAmountCents returns the SubtotalTaxableAmountCents field if non-nil, zero value otherwise.

### GetSubtotalTaxableAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalTaxableAmountCentsOk() (*int32, bool)`

GetSubtotalTaxableAmountCentsOk returns a tuple with the SubtotalTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalTaxableAmountCents(v int32)`

SetSubtotalTaxableAmountCents sets SubtotalTaxableAmountCents field to given value.

### HasSubtotalTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasSubtotalTaxableAmountCents() bool`

HasSubtotalTaxableAmountCents returns a boolean if a field has been set.

### GetSubtotalTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalTaxableAmountFloat() float32`

GetSubtotalTaxableAmountFloat returns the SubtotalTaxableAmountFloat field if non-nil, zero value otherwise.

### GetSubtotalTaxableAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalTaxableAmountFloatOk() (*float32, bool)`

GetSubtotalTaxableAmountFloatOk returns a tuple with the SubtotalTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalTaxableAmountFloat(v float32)`

SetSubtotalTaxableAmountFloat sets SubtotalTaxableAmountFloat field to given value.

### HasSubtotalTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasSubtotalTaxableAmountFloat() bool`

HasSubtotalTaxableAmountFloat returns a boolean if a field has been set.

### GetFormattedSubtotalTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedSubtotalTaxableAmount() string`

GetFormattedSubtotalTaxableAmount returns the FormattedSubtotalTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedSubtotalTaxableAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedSubtotalTaxableAmountOk() (*string, bool)`

GetFormattedSubtotalTaxableAmountOk returns a tuple with the FormattedSubtotalTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSubtotalTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedSubtotalTaxableAmount(v string)`

SetFormattedSubtotalTaxableAmount sets FormattedSubtotalTaxableAmount field to given value.

### HasFormattedSubtotalTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedSubtotalTaxableAmount() bool`

HasFormattedSubtotalTaxableAmount returns a boolean if a field has been set.

### GetShippingTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingTaxableAmountCents() int32`

GetShippingTaxableAmountCents returns the ShippingTaxableAmountCents field if non-nil, zero value otherwise.

### GetShippingTaxableAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingTaxableAmountCentsOk() (*int32, bool)`

GetShippingTaxableAmountCentsOk returns a tuple with the ShippingTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingTaxableAmountCents(v int32)`

SetShippingTaxableAmountCents sets ShippingTaxableAmountCents field to given value.

### HasShippingTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasShippingTaxableAmountCents() bool`

HasShippingTaxableAmountCents returns a boolean if a field has been set.

### GetShippingTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingTaxableAmountFloat() float32`

GetShippingTaxableAmountFloat returns the ShippingTaxableAmountFloat field if non-nil, zero value otherwise.

### GetShippingTaxableAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingTaxableAmountFloatOk() (*float32, bool)`

GetShippingTaxableAmountFloatOk returns a tuple with the ShippingTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingTaxableAmountFloat(v float32)`

SetShippingTaxableAmountFloat sets ShippingTaxableAmountFloat field to given value.

### HasShippingTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasShippingTaxableAmountFloat() bool`

HasShippingTaxableAmountFloat returns a boolean if a field has been set.

### GetFormattedShippingTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedShippingTaxableAmount() string`

GetFormattedShippingTaxableAmount returns the FormattedShippingTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedShippingTaxableAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedShippingTaxableAmountOk() (*string, bool)`

GetFormattedShippingTaxableAmountOk returns a tuple with the FormattedShippingTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedShippingTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedShippingTaxableAmount(v string)`

SetFormattedShippingTaxableAmount sets FormattedShippingTaxableAmount field to given value.

### HasFormattedShippingTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedShippingTaxableAmount() bool`

HasFormattedShippingTaxableAmount returns a boolean if a field has been set.

### GetPaymentMethodTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodTaxableAmountCents() int32`

GetPaymentMethodTaxableAmountCents returns the PaymentMethodTaxableAmountCents field if non-nil, zero value otherwise.

### GetPaymentMethodTaxableAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodTaxableAmountCentsOk() (*int32, bool)`

GetPaymentMethodTaxableAmountCentsOk returns a tuple with the PaymentMethodTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodTaxableAmountCents(v int32)`

SetPaymentMethodTaxableAmountCents sets PaymentMethodTaxableAmountCents field to given value.

### HasPaymentMethodTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentMethodTaxableAmountCents() bool`

HasPaymentMethodTaxableAmountCents returns a boolean if a field has been set.

### GetPaymentMethodTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodTaxableAmountFloat() float32`

GetPaymentMethodTaxableAmountFloat returns the PaymentMethodTaxableAmountFloat field if non-nil, zero value otherwise.

### GetPaymentMethodTaxableAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodTaxableAmountFloatOk() (*float32, bool)`

GetPaymentMethodTaxableAmountFloatOk returns a tuple with the PaymentMethodTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodTaxableAmountFloat(v float32)`

SetPaymentMethodTaxableAmountFloat sets PaymentMethodTaxableAmountFloat field to given value.

### HasPaymentMethodTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentMethodTaxableAmountFloat() bool`

HasPaymentMethodTaxableAmountFloat returns a boolean if a field has been set.

### GetFormattedPaymentMethodTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedPaymentMethodTaxableAmount() string`

GetFormattedPaymentMethodTaxableAmount returns the FormattedPaymentMethodTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedPaymentMethodTaxableAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedPaymentMethodTaxableAmountOk() (*string, bool)`

GetFormattedPaymentMethodTaxableAmountOk returns a tuple with the FormattedPaymentMethodTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPaymentMethodTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedPaymentMethodTaxableAmount(v string)`

SetFormattedPaymentMethodTaxableAmount sets FormattedPaymentMethodTaxableAmount field to given value.

### HasFormattedPaymentMethodTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedPaymentMethodTaxableAmount() bool`

HasFormattedPaymentMethodTaxableAmount returns a boolean if a field has been set.

### GetAdjustmentTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentTaxableAmountCents() int32`

GetAdjustmentTaxableAmountCents returns the AdjustmentTaxableAmountCents field if non-nil, zero value otherwise.

### GetAdjustmentTaxableAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentTaxableAmountCentsOk() (*int32, bool)`

GetAdjustmentTaxableAmountCentsOk returns a tuple with the AdjustmentTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentTaxableAmountCents(v int32)`

SetAdjustmentTaxableAmountCents sets AdjustmentTaxableAmountCents field to given value.

### HasAdjustmentTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasAdjustmentTaxableAmountCents() bool`

HasAdjustmentTaxableAmountCents returns a boolean if a field has been set.

### GetAdjustmentTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentTaxableAmountFloat() float32`

GetAdjustmentTaxableAmountFloat returns the AdjustmentTaxableAmountFloat field if non-nil, zero value otherwise.

### GetAdjustmentTaxableAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentTaxableAmountFloatOk() (*float32, bool)`

GetAdjustmentTaxableAmountFloatOk returns a tuple with the AdjustmentTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentTaxableAmountFloat(v float32)`

SetAdjustmentTaxableAmountFloat sets AdjustmentTaxableAmountFloat field to given value.

### HasAdjustmentTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasAdjustmentTaxableAmountFloat() bool`

HasAdjustmentTaxableAmountFloat returns a boolean if a field has been set.

### GetFormattedAdjustmentTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedAdjustmentTaxableAmount() string`

GetFormattedAdjustmentTaxableAmount returns the FormattedAdjustmentTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedAdjustmentTaxableAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedAdjustmentTaxableAmountOk() (*string, bool)`

GetFormattedAdjustmentTaxableAmountOk returns a tuple with the FormattedAdjustmentTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAdjustmentTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedAdjustmentTaxableAmount(v string)`

SetFormattedAdjustmentTaxableAmount sets FormattedAdjustmentTaxableAmount field to given value.

### HasFormattedAdjustmentTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedAdjustmentTaxableAmount() bool`

HasFormattedAdjustmentTaxableAmount returns a boolean if a field has been set.

### GetTotalAmountWithTaxesCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalAmountWithTaxesCents() int32`

GetTotalAmountWithTaxesCents returns the TotalAmountWithTaxesCents field if non-nil, zero value otherwise.

### GetTotalAmountWithTaxesCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalAmountWithTaxesCentsOk() (*int32, bool)`

GetTotalAmountWithTaxesCentsOk returns a tuple with the TotalAmountWithTaxesCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountWithTaxesCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalAmountWithTaxesCents(v int32)`

SetTotalAmountWithTaxesCents sets TotalAmountWithTaxesCents field to given value.

### HasTotalAmountWithTaxesCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasTotalAmountWithTaxesCents() bool`

HasTotalAmountWithTaxesCents returns a boolean if a field has been set.

### GetTotalAmountWithTaxesFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalAmountWithTaxesFloat() float32`

GetTotalAmountWithTaxesFloat returns the TotalAmountWithTaxesFloat field if non-nil, zero value otherwise.

### GetTotalAmountWithTaxesFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalAmountWithTaxesFloatOk() (*float32, bool)`

GetTotalAmountWithTaxesFloatOk returns a tuple with the TotalAmountWithTaxesFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountWithTaxesFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalAmountWithTaxesFloat(v float32)`

SetTotalAmountWithTaxesFloat sets TotalAmountWithTaxesFloat field to given value.

### HasTotalAmountWithTaxesFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasTotalAmountWithTaxesFloat() bool`

HasTotalAmountWithTaxesFloat returns a boolean if a field has been set.

### GetFormattedTotalAmountWithTaxes

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedTotalAmountWithTaxes() string`

GetFormattedTotalAmountWithTaxes returns the FormattedTotalAmountWithTaxes field if non-nil, zero value otherwise.

### GetFormattedTotalAmountWithTaxesOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedTotalAmountWithTaxesOk() (*string, bool)`

GetFormattedTotalAmountWithTaxesOk returns a tuple with the FormattedTotalAmountWithTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalAmountWithTaxes

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedTotalAmountWithTaxes(v string)`

SetFormattedTotalAmountWithTaxes sets FormattedTotalAmountWithTaxes field to given value.

### HasFormattedTotalAmountWithTaxes

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedTotalAmountWithTaxes() bool`

HasFormattedTotalAmountWithTaxes returns a boolean if a field has been set.

### GetFeesAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetFeesAmountCents() int32`

GetFeesAmountCents returns the FeesAmountCents field if non-nil, zero value otherwise.

### GetFeesAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFeesAmountCentsOk() (*int32, bool)`

GetFeesAmountCentsOk returns a tuple with the FeesAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetFeesAmountCents(v int32)`

SetFeesAmountCents sets FeesAmountCents field to given value.

### HasFeesAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasFeesAmountCents() bool`

HasFeesAmountCents returns a boolean if a field has been set.

### GetFeesAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetFeesAmountFloat() float32`

GetFeesAmountFloat returns the FeesAmountFloat field if non-nil, zero value otherwise.

### GetFeesAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFeesAmountFloatOk() (*float32, bool)`

GetFeesAmountFloatOk returns a tuple with the FeesAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetFeesAmountFloat(v float32)`

SetFeesAmountFloat sets FeesAmountFloat field to given value.

### HasFeesAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasFeesAmountFloat() bool`

HasFeesAmountFloat returns a boolean if a field has been set.

### GetFormattedFeesAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedFeesAmount() string`

GetFormattedFeesAmount returns the FormattedFeesAmount field if non-nil, zero value otherwise.

### GetFormattedFeesAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedFeesAmountOk() (*string, bool)`

GetFormattedFeesAmountOk returns a tuple with the FormattedFeesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedFeesAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedFeesAmount(v string)`

SetFormattedFeesAmount sets FormattedFeesAmount field to given value.

### HasFormattedFeesAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedFeesAmount() bool`

HasFormattedFeesAmount returns a boolean if a field has been set.

### GetDutyAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetDutyAmountCents() int32`

GetDutyAmountCents returns the DutyAmountCents field if non-nil, zero value otherwise.

### GetDutyAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetDutyAmountCentsOk() (*int32, bool)`

GetDutyAmountCentsOk returns a tuple with the DutyAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetDutyAmountCents(v int32)`

SetDutyAmountCents sets DutyAmountCents field to given value.

### HasDutyAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasDutyAmountCents() bool`

HasDutyAmountCents returns a boolean if a field has been set.

### GetDutyAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetDutyAmountFloat() float32`

GetDutyAmountFloat returns the DutyAmountFloat field if non-nil, zero value otherwise.

### GetDutyAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetDutyAmountFloatOk() (*float32, bool)`

GetDutyAmountFloatOk returns a tuple with the DutyAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetDutyAmountFloat(v float32)`

SetDutyAmountFloat sets DutyAmountFloat field to given value.

### HasDutyAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasDutyAmountFloat() bool`

HasDutyAmountFloat returns a boolean if a field has been set.

### GetFormattedDutyAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedDutyAmount() string`

GetFormattedDutyAmount returns the FormattedDutyAmount field if non-nil, zero value otherwise.

### GetFormattedDutyAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedDutyAmountOk() (*string, bool)`

GetFormattedDutyAmountOk returns a tuple with the FormattedDutyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedDutyAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedDutyAmount(v string)`

SetFormattedDutyAmount sets FormattedDutyAmount field to given value.

### HasFormattedDutyAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedDutyAmount() bool`

HasFormattedDutyAmount returns a boolean if a field has been set.

### GetSkusCount

`func (o *GETOrders200ResponseDataInnerAttributes) GetSkusCount() int32`

GetSkusCount returns the SkusCount field if non-nil, zero value otherwise.

### GetSkusCountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetSkusCountOk() (*int32, bool)`

GetSkusCountOk returns a tuple with the SkusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusCount

`func (o *GETOrders200ResponseDataInnerAttributes) SetSkusCount(v int32)`

SetSkusCount sets SkusCount field to given value.

### HasSkusCount

`func (o *GETOrders200ResponseDataInnerAttributes) HasSkusCount() bool`

HasSkusCount returns a boolean if a field has been set.

### GetLineItemOptionsCount

`func (o *GETOrders200ResponseDataInnerAttributes) GetLineItemOptionsCount() int32`

GetLineItemOptionsCount returns the LineItemOptionsCount field if non-nil, zero value otherwise.

### GetLineItemOptionsCountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetLineItemOptionsCountOk() (*int32, bool)`

GetLineItemOptionsCountOk returns a tuple with the LineItemOptionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemOptionsCount

`func (o *GETOrders200ResponseDataInnerAttributes) SetLineItemOptionsCount(v int32)`

SetLineItemOptionsCount sets LineItemOptionsCount field to given value.

### HasLineItemOptionsCount

`func (o *GETOrders200ResponseDataInnerAttributes) HasLineItemOptionsCount() bool`

HasLineItemOptionsCount returns a boolean if a field has been set.

### GetShipmentsCount

`func (o *GETOrders200ResponseDataInnerAttributes) GetShipmentsCount() int32`

GetShipmentsCount returns the ShipmentsCount field if non-nil, zero value otherwise.

### GetShipmentsCountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetShipmentsCountOk() (*int32, bool)`

GetShipmentsCountOk returns a tuple with the ShipmentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentsCount

`func (o *GETOrders200ResponseDataInnerAttributes) SetShipmentsCount(v int32)`

SetShipmentsCount sets ShipmentsCount field to given value.

### HasShipmentsCount

`func (o *GETOrders200ResponseDataInnerAttributes) HasShipmentsCount() bool`

HasShipmentsCount returns a boolean if a field has been set.

### GetTaxCalculationsCount

`func (o *GETOrders200ResponseDataInnerAttributes) GetTaxCalculationsCount() int32`

GetTaxCalculationsCount returns the TaxCalculationsCount field if non-nil, zero value otherwise.

### GetTaxCalculationsCountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTaxCalculationsCountOk() (*int32, bool)`

GetTaxCalculationsCountOk returns a tuple with the TaxCalculationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationsCount

`func (o *GETOrders200ResponseDataInnerAttributes) SetTaxCalculationsCount(v int32)`

SetTaxCalculationsCount sets TaxCalculationsCount field to given value.

### HasTaxCalculationsCount

`func (o *GETOrders200ResponseDataInnerAttributes) HasTaxCalculationsCount() bool`

HasTaxCalculationsCount returns a boolean if a field has been set.

### GetPaymentSourceDetails

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentSourceDetails() map[string]interface{}`

GetPaymentSourceDetails returns the PaymentSourceDetails field if non-nil, zero value otherwise.

### GetPaymentSourceDetailsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentSourceDetailsOk() (*map[string]interface{}, bool)`

GetPaymentSourceDetailsOk returns a tuple with the PaymentSourceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceDetails

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentSourceDetails(v map[string]interface{})`

SetPaymentSourceDetails sets PaymentSourceDetails field to given value.

### HasPaymentSourceDetails

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentSourceDetails() bool`

HasPaymentSourceDetails returns a boolean if a field has been set.

### GetToken

`func (o *GETOrders200ResponseDataInnerAttributes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GETOrders200ResponseDataInnerAttributes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GETOrders200ResponseDataInnerAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetCartUrl

`func (o *GETOrders200ResponseDataInnerAttributes) GetCartUrl() string`

GetCartUrl returns the CartUrl field if non-nil, zero value otherwise.

### GetCartUrlOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetCartUrlOk() (*string, bool)`

GetCartUrlOk returns a tuple with the CartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartUrl

`func (o *GETOrders200ResponseDataInnerAttributes) SetCartUrl(v string)`

SetCartUrl sets CartUrl field to given value.

### HasCartUrl

`func (o *GETOrders200ResponseDataInnerAttributes) HasCartUrl() bool`

HasCartUrl returns a boolean if a field has been set.

### GetReturnUrl

`func (o *GETOrders200ResponseDataInnerAttributes) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *GETOrders200ResponseDataInnerAttributes) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *GETOrders200ResponseDataInnerAttributes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetTermsUrl

`func (o *GETOrders200ResponseDataInnerAttributes) GetTermsUrl() string`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTermsUrlOk() (*string, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *GETOrders200ResponseDataInnerAttributes) SetTermsUrl(v string)`

SetTermsUrl sets TermsUrl field to given value.

### HasTermsUrl

`func (o *GETOrders200ResponseDataInnerAttributes) HasTermsUrl() bool`

HasTermsUrl returns a boolean if a field has been set.

### GetPrivacyUrl

`func (o *GETOrders200ResponseDataInnerAttributes) GetPrivacyUrl() string`

GetPrivacyUrl returns the PrivacyUrl field if non-nil, zero value otherwise.

### GetPrivacyUrlOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPrivacyUrlOk() (*string, bool)`

GetPrivacyUrlOk returns a tuple with the PrivacyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyUrl

`func (o *GETOrders200ResponseDataInnerAttributes) SetPrivacyUrl(v string)`

SetPrivacyUrl sets PrivacyUrl field to given value.

### HasPrivacyUrl

`func (o *GETOrders200ResponseDataInnerAttributes) HasPrivacyUrl() bool`

HasPrivacyUrl returns a boolean if a field has been set.

### GetCheckoutUrl

`func (o *GETOrders200ResponseDataInnerAttributes) GetCheckoutUrl() string`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetCheckoutUrlOk() (*string, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *GETOrders200ResponseDataInnerAttributes) SetCheckoutUrl(v string)`

SetCheckoutUrl sets CheckoutUrl field to given value.

### HasCheckoutUrl

`func (o *GETOrders200ResponseDataInnerAttributes) HasCheckoutUrl() bool`

HasCheckoutUrl returns a boolean if a field has been set.

### GetPlacedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetPlacedAt() string`

GetPlacedAt returns the PlacedAt field if non-nil, zero value otherwise.

### GetPlacedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPlacedAtOk() (*string, bool)`

GetPlacedAtOk returns a tuple with the PlacedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetPlacedAt(v string)`

SetPlacedAt sets PlacedAt field to given value.

### HasPlacedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasPlacedAt() bool`

HasPlacedAt returns a boolean if a field has been set.

### GetApprovedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetApprovedAt() string`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetApprovedAtOk() (*string, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetApprovedAt(v string)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### GetCancelledAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetCancelledAt() string`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetCancelledAtOk() (*string, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetCancelledAt(v string)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### GetPaymentUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentUpdatedAt() string`

GetPaymentUpdatedAt returns the PaymentUpdatedAt field if non-nil, zero value otherwise.

### GetPaymentUpdatedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentUpdatedAtOk() (*string, bool)`

GetPaymentUpdatedAtOk returns a tuple with the PaymentUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentUpdatedAt(v string)`

SetPaymentUpdatedAt sets PaymentUpdatedAt field to given value.

### HasPaymentUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentUpdatedAt() bool`

HasPaymentUpdatedAt returns a boolean if a field has been set.

### GetFulfillmentUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetFulfillmentUpdatedAt() string`

GetFulfillmentUpdatedAt returns the FulfillmentUpdatedAt field if non-nil, zero value otherwise.

### GetFulfillmentUpdatedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFulfillmentUpdatedAtOk() (*string, bool)`

GetFulfillmentUpdatedAtOk returns a tuple with the FulfillmentUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetFulfillmentUpdatedAt(v string)`

SetFulfillmentUpdatedAt sets FulfillmentUpdatedAt field to given value.

### HasFulfillmentUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasFulfillmentUpdatedAt() bool`

HasFulfillmentUpdatedAt returns a boolean if a field has been set.

### GetRefreshedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetRefreshedAt() string`

GetRefreshedAt returns the RefreshedAt field if non-nil, zero value otherwise.

### GetRefreshedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetRefreshedAtOk() (*string, bool)`

GetRefreshedAtOk returns a tuple with the RefreshedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetRefreshedAt(v string)`

SetRefreshedAt sets RefreshedAt field to given value.

### HasRefreshedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasRefreshedAt() bool`

HasRefreshedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *GETOrders200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETOrders200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETOrders200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETOrders200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETOrders200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETOrders200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETOrders200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETOrders200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETOrders200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETOrders200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETOrders200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETOrders200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


