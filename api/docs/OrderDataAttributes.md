# OrderDataAttributes

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

### NewOrderDataAttributes

`func NewOrderDataAttributes() *OrderDataAttributes`

NewOrderDataAttributes instantiates a new OrderDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDataAttributesWithDefaults

`func NewOrderDataAttributesWithDefaults() *OrderDataAttributes`

NewOrderDataAttributesWithDefaults instantiates a new OrderDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *OrderDataAttributes) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OrderDataAttributes) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OrderDataAttributes) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *OrderDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetAutorefresh

`func (o *OrderDataAttributes) GetAutorefresh() bool`

GetAutorefresh returns the Autorefresh field if non-nil, zero value otherwise.

### GetAutorefreshOk

`func (o *OrderDataAttributes) GetAutorefreshOk() (*bool, bool)`

GetAutorefreshOk returns a tuple with the Autorefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutorefresh

`func (o *OrderDataAttributes) SetAutorefresh(v bool)`

SetAutorefresh sets Autorefresh field to given value.

### HasAutorefresh

`func (o *OrderDataAttributes) HasAutorefresh() bool`

HasAutorefresh returns a boolean if a field has been set.

### GetStatus

`func (o *OrderDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *OrderDataAttributes) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *OrderDataAttributes) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *OrderDataAttributes) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *OrderDataAttributes) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetFulfillmentStatus

`func (o *OrderDataAttributes) GetFulfillmentStatus() string`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *OrderDataAttributes) GetFulfillmentStatusOk() (*string, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *OrderDataAttributes) SetFulfillmentStatus(v string)`

SetFulfillmentStatus sets FulfillmentStatus field to given value.

### HasFulfillmentStatus

`func (o *OrderDataAttributes) HasFulfillmentStatus() bool`

HasFulfillmentStatus returns a boolean if a field has been set.

### GetGuest

`func (o *OrderDataAttributes) GetGuest() bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *OrderDataAttributes) GetGuestOk() (*bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *OrderDataAttributes) SetGuest(v bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *OrderDataAttributes) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetEditable

`func (o *OrderDataAttributes) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *OrderDataAttributes) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *OrderDataAttributes) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *OrderDataAttributes) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *OrderDataAttributes) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *OrderDataAttributes) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *OrderDataAttributes) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *OrderDataAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetLanguageCode

`func (o *OrderDataAttributes) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *OrderDataAttributes) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *OrderDataAttributes) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *OrderDataAttributes) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *OrderDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *OrderDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *OrderDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *OrderDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetTaxIncluded

`func (o *OrderDataAttributes) GetTaxIncluded() bool`

GetTaxIncluded returns the TaxIncluded field if non-nil, zero value otherwise.

### GetTaxIncludedOk

`func (o *OrderDataAttributes) GetTaxIncludedOk() (*bool, bool)`

GetTaxIncludedOk returns a tuple with the TaxIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncluded

`func (o *OrderDataAttributes) SetTaxIncluded(v bool)`

SetTaxIncluded sets TaxIncluded field to given value.

### HasTaxIncluded

`func (o *OrderDataAttributes) HasTaxIncluded() bool`

HasTaxIncluded returns a boolean if a field has been set.

### GetTaxRate

`func (o *OrderDataAttributes) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *OrderDataAttributes) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *OrderDataAttributes) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *OrderDataAttributes) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetFreightTaxable

`func (o *OrderDataAttributes) GetFreightTaxable() bool`

GetFreightTaxable returns the FreightTaxable field if non-nil, zero value otherwise.

### GetFreightTaxableOk

`func (o *OrderDataAttributes) GetFreightTaxableOk() (*bool, bool)`

GetFreightTaxableOk returns a tuple with the FreightTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTaxable

`func (o *OrderDataAttributes) SetFreightTaxable(v bool)`

SetFreightTaxable sets FreightTaxable field to given value.

### HasFreightTaxable

`func (o *OrderDataAttributes) HasFreightTaxable() bool`

HasFreightTaxable returns a boolean if a field has been set.

### GetRequiresBillingInfo

`func (o *OrderDataAttributes) GetRequiresBillingInfo() bool`

GetRequiresBillingInfo returns the RequiresBillingInfo field if non-nil, zero value otherwise.

### GetRequiresBillingInfoOk

`func (o *OrderDataAttributes) GetRequiresBillingInfoOk() (*bool, bool)`

GetRequiresBillingInfoOk returns a tuple with the RequiresBillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresBillingInfo

`func (o *OrderDataAttributes) SetRequiresBillingInfo(v bool)`

SetRequiresBillingInfo sets RequiresBillingInfo field to given value.

### HasRequiresBillingInfo

`func (o *OrderDataAttributes) HasRequiresBillingInfo() bool`

HasRequiresBillingInfo returns a boolean if a field has been set.

### GetCountryCode

`func (o *OrderDataAttributes) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *OrderDataAttributes) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *OrderDataAttributes) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *OrderDataAttributes) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetShippingCountryCodeLock

`func (o *OrderDataAttributes) GetShippingCountryCodeLock() string`

GetShippingCountryCodeLock returns the ShippingCountryCodeLock field if non-nil, zero value otherwise.

### GetShippingCountryCodeLockOk

`func (o *OrderDataAttributes) GetShippingCountryCodeLockOk() (*string, bool)`

GetShippingCountryCodeLockOk returns a tuple with the ShippingCountryCodeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCountryCodeLock

`func (o *OrderDataAttributes) SetShippingCountryCodeLock(v string)`

SetShippingCountryCodeLock sets ShippingCountryCodeLock field to given value.

### HasShippingCountryCodeLock

`func (o *OrderDataAttributes) HasShippingCountryCodeLock() bool`

HasShippingCountryCodeLock returns a boolean if a field has been set.

### GetCouponCode

`func (o *OrderDataAttributes) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *OrderDataAttributes) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *OrderDataAttributes) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *OrderDataAttributes) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.

### GetGiftCardCode

`func (o *OrderDataAttributes) GetGiftCardCode() string`

GetGiftCardCode returns the GiftCardCode field if non-nil, zero value otherwise.

### GetGiftCardCodeOk

`func (o *OrderDataAttributes) GetGiftCardCodeOk() (*string, bool)`

GetGiftCardCodeOk returns a tuple with the GiftCardCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardCode

`func (o *OrderDataAttributes) SetGiftCardCode(v string)`

SetGiftCardCode sets GiftCardCode field to given value.

### HasGiftCardCode

`func (o *OrderDataAttributes) HasGiftCardCode() bool`

HasGiftCardCode returns a boolean if a field has been set.

### GetGiftCardOrCouponCode

`func (o *OrderDataAttributes) GetGiftCardOrCouponCode() string`

GetGiftCardOrCouponCode returns the GiftCardOrCouponCode field if non-nil, zero value otherwise.

### GetGiftCardOrCouponCodeOk

`func (o *OrderDataAttributes) GetGiftCardOrCouponCodeOk() (*string, bool)`

GetGiftCardOrCouponCodeOk returns a tuple with the GiftCardOrCouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardOrCouponCode

`func (o *OrderDataAttributes) SetGiftCardOrCouponCode(v string)`

SetGiftCardOrCouponCode sets GiftCardOrCouponCode field to given value.

### HasGiftCardOrCouponCode

`func (o *OrderDataAttributes) HasGiftCardOrCouponCode() bool`

HasGiftCardOrCouponCode returns a boolean if a field has been set.

### GetSubtotalAmountCents

`func (o *OrderDataAttributes) GetSubtotalAmountCents() int32`

GetSubtotalAmountCents returns the SubtotalAmountCents field if non-nil, zero value otherwise.

### GetSubtotalAmountCentsOk

`func (o *OrderDataAttributes) GetSubtotalAmountCentsOk() (*int32, bool)`

GetSubtotalAmountCentsOk returns a tuple with the SubtotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalAmountCents

`func (o *OrderDataAttributes) SetSubtotalAmountCents(v int32)`

SetSubtotalAmountCents sets SubtotalAmountCents field to given value.

### HasSubtotalAmountCents

`func (o *OrderDataAttributes) HasSubtotalAmountCents() bool`

HasSubtotalAmountCents returns a boolean if a field has been set.

### GetSubtotalAmountFloat

`func (o *OrderDataAttributes) GetSubtotalAmountFloat() float32`

GetSubtotalAmountFloat returns the SubtotalAmountFloat field if non-nil, zero value otherwise.

### GetSubtotalAmountFloatOk

`func (o *OrderDataAttributes) GetSubtotalAmountFloatOk() (*float32, bool)`

GetSubtotalAmountFloatOk returns a tuple with the SubtotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalAmountFloat

`func (o *OrderDataAttributes) SetSubtotalAmountFloat(v float32)`

SetSubtotalAmountFloat sets SubtotalAmountFloat field to given value.

### HasSubtotalAmountFloat

`func (o *OrderDataAttributes) HasSubtotalAmountFloat() bool`

HasSubtotalAmountFloat returns a boolean if a field has been set.

### GetFormattedSubtotalAmount

`func (o *OrderDataAttributes) GetFormattedSubtotalAmount() string`

GetFormattedSubtotalAmount returns the FormattedSubtotalAmount field if non-nil, zero value otherwise.

### GetFormattedSubtotalAmountOk

`func (o *OrderDataAttributes) GetFormattedSubtotalAmountOk() (*string, bool)`

GetFormattedSubtotalAmountOk returns a tuple with the FormattedSubtotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSubtotalAmount

`func (o *OrderDataAttributes) SetFormattedSubtotalAmount(v string)`

SetFormattedSubtotalAmount sets FormattedSubtotalAmount field to given value.

### HasFormattedSubtotalAmount

`func (o *OrderDataAttributes) HasFormattedSubtotalAmount() bool`

HasFormattedSubtotalAmount returns a boolean if a field has been set.

### GetShippingAmountCents

`func (o *OrderDataAttributes) GetShippingAmountCents() int32`

GetShippingAmountCents returns the ShippingAmountCents field if non-nil, zero value otherwise.

### GetShippingAmountCentsOk

`func (o *OrderDataAttributes) GetShippingAmountCentsOk() (*int32, bool)`

GetShippingAmountCentsOk returns a tuple with the ShippingAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAmountCents

`func (o *OrderDataAttributes) SetShippingAmountCents(v int32)`

SetShippingAmountCents sets ShippingAmountCents field to given value.

### HasShippingAmountCents

`func (o *OrderDataAttributes) HasShippingAmountCents() bool`

HasShippingAmountCents returns a boolean if a field has been set.

### GetShippingAmountFloat

`func (o *OrderDataAttributes) GetShippingAmountFloat() float32`

GetShippingAmountFloat returns the ShippingAmountFloat field if non-nil, zero value otherwise.

### GetShippingAmountFloatOk

`func (o *OrderDataAttributes) GetShippingAmountFloatOk() (*float32, bool)`

GetShippingAmountFloatOk returns a tuple with the ShippingAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAmountFloat

`func (o *OrderDataAttributes) SetShippingAmountFloat(v float32)`

SetShippingAmountFloat sets ShippingAmountFloat field to given value.

### HasShippingAmountFloat

`func (o *OrderDataAttributes) HasShippingAmountFloat() bool`

HasShippingAmountFloat returns a boolean if a field has been set.

### GetFormattedShippingAmount

`func (o *OrderDataAttributes) GetFormattedShippingAmount() string`

GetFormattedShippingAmount returns the FormattedShippingAmount field if non-nil, zero value otherwise.

### GetFormattedShippingAmountOk

`func (o *OrderDataAttributes) GetFormattedShippingAmountOk() (*string, bool)`

GetFormattedShippingAmountOk returns a tuple with the FormattedShippingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedShippingAmount

`func (o *OrderDataAttributes) SetFormattedShippingAmount(v string)`

SetFormattedShippingAmount sets FormattedShippingAmount field to given value.

### HasFormattedShippingAmount

`func (o *OrderDataAttributes) HasFormattedShippingAmount() bool`

HasFormattedShippingAmount returns a boolean if a field has been set.

### GetPaymentMethodAmountCents

`func (o *OrderDataAttributes) GetPaymentMethodAmountCents() int32`

GetPaymentMethodAmountCents returns the PaymentMethodAmountCents field if non-nil, zero value otherwise.

### GetPaymentMethodAmountCentsOk

`func (o *OrderDataAttributes) GetPaymentMethodAmountCentsOk() (*int32, bool)`

GetPaymentMethodAmountCentsOk returns a tuple with the PaymentMethodAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodAmountCents

`func (o *OrderDataAttributes) SetPaymentMethodAmountCents(v int32)`

SetPaymentMethodAmountCents sets PaymentMethodAmountCents field to given value.

### HasPaymentMethodAmountCents

`func (o *OrderDataAttributes) HasPaymentMethodAmountCents() bool`

HasPaymentMethodAmountCents returns a boolean if a field has been set.

### GetPaymentMethodAmountFloat

`func (o *OrderDataAttributes) GetPaymentMethodAmountFloat() float32`

GetPaymentMethodAmountFloat returns the PaymentMethodAmountFloat field if non-nil, zero value otherwise.

### GetPaymentMethodAmountFloatOk

`func (o *OrderDataAttributes) GetPaymentMethodAmountFloatOk() (*float32, bool)`

GetPaymentMethodAmountFloatOk returns a tuple with the PaymentMethodAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodAmountFloat

`func (o *OrderDataAttributes) SetPaymentMethodAmountFloat(v float32)`

SetPaymentMethodAmountFloat sets PaymentMethodAmountFloat field to given value.

### HasPaymentMethodAmountFloat

`func (o *OrderDataAttributes) HasPaymentMethodAmountFloat() bool`

HasPaymentMethodAmountFloat returns a boolean if a field has been set.

### GetFormattedPaymentMethodAmount

`func (o *OrderDataAttributes) GetFormattedPaymentMethodAmount() string`

GetFormattedPaymentMethodAmount returns the FormattedPaymentMethodAmount field if non-nil, zero value otherwise.

### GetFormattedPaymentMethodAmountOk

`func (o *OrderDataAttributes) GetFormattedPaymentMethodAmountOk() (*string, bool)`

GetFormattedPaymentMethodAmountOk returns a tuple with the FormattedPaymentMethodAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPaymentMethodAmount

`func (o *OrderDataAttributes) SetFormattedPaymentMethodAmount(v string)`

SetFormattedPaymentMethodAmount sets FormattedPaymentMethodAmount field to given value.

### HasFormattedPaymentMethodAmount

`func (o *OrderDataAttributes) HasFormattedPaymentMethodAmount() bool`

HasFormattedPaymentMethodAmount returns a boolean if a field has been set.

### GetDiscountAmountCents

`func (o *OrderDataAttributes) GetDiscountAmountCents() int32`

GetDiscountAmountCents returns the DiscountAmountCents field if non-nil, zero value otherwise.

### GetDiscountAmountCentsOk

`func (o *OrderDataAttributes) GetDiscountAmountCentsOk() (*int32, bool)`

GetDiscountAmountCentsOk returns a tuple with the DiscountAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmountCents

`func (o *OrderDataAttributes) SetDiscountAmountCents(v int32)`

SetDiscountAmountCents sets DiscountAmountCents field to given value.

### HasDiscountAmountCents

`func (o *OrderDataAttributes) HasDiscountAmountCents() bool`

HasDiscountAmountCents returns a boolean if a field has been set.

### GetDiscountAmountFloat

`func (o *OrderDataAttributes) GetDiscountAmountFloat() float32`

GetDiscountAmountFloat returns the DiscountAmountFloat field if non-nil, zero value otherwise.

### GetDiscountAmountFloatOk

`func (o *OrderDataAttributes) GetDiscountAmountFloatOk() (*float32, bool)`

GetDiscountAmountFloatOk returns a tuple with the DiscountAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmountFloat

`func (o *OrderDataAttributes) SetDiscountAmountFloat(v float32)`

SetDiscountAmountFloat sets DiscountAmountFloat field to given value.

### HasDiscountAmountFloat

`func (o *OrderDataAttributes) HasDiscountAmountFloat() bool`

HasDiscountAmountFloat returns a boolean if a field has been set.

### GetFormattedDiscountAmount

`func (o *OrderDataAttributes) GetFormattedDiscountAmount() string`

GetFormattedDiscountAmount returns the FormattedDiscountAmount field if non-nil, zero value otherwise.

### GetFormattedDiscountAmountOk

`func (o *OrderDataAttributes) GetFormattedDiscountAmountOk() (*string, bool)`

GetFormattedDiscountAmountOk returns a tuple with the FormattedDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedDiscountAmount

`func (o *OrderDataAttributes) SetFormattedDiscountAmount(v string)`

SetFormattedDiscountAmount sets FormattedDiscountAmount field to given value.

### HasFormattedDiscountAmount

`func (o *OrderDataAttributes) HasFormattedDiscountAmount() bool`

HasFormattedDiscountAmount returns a boolean if a field has been set.

### GetAdjustmentAmountCents

`func (o *OrderDataAttributes) GetAdjustmentAmountCents() int32`

GetAdjustmentAmountCents returns the AdjustmentAmountCents field if non-nil, zero value otherwise.

### GetAdjustmentAmountCentsOk

`func (o *OrderDataAttributes) GetAdjustmentAmountCentsOk() (*int32, bool)`

GetAdjustmentAmountCentsOk returns a tuple with the AdjustmentAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentAmountCents

`func (o *OrderDataAttributes) SetAdjustmentAmountCents(v int32)`

SetAdjustmentAmountCents sets AdjustmentAmountCents field to given value.

### HasAdjustmentAmountCents

`func (o *OrderDataAttributes) HasAdjustmentAmountCents() bool`

HasAdjustmentAmountCents returns a boolean if a field has been set.

### GetAdjustmentAmountFloat

`func (o *OrderDataAttributes) GetAdjustmentAmountFloat() float32`

GetAdjustmentAmountFloat returns the AdjustmentAmountFloat field if non-nil, zero value otherwise.

### GetAdjustmentAmountFloatOk

`func (o *OrderDataAttributes) GetAdjustmentAmountFloatOk() (*float32, bool)`

GetAdjustmentAmountFloatOk returns a tuple with the AdjustmentAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentAmountFloat

`func (o *OrderDataAttributes) SetAdjustmentAmountFloat(v float32)`

SetAdjustmentAmountFloat sets AdjustmentAmountFloat field to given value.

### HasAdjustmentAmountFloat

`func (o *OrderDataAttributes) HasAdjustmentAmountFloat() bool`

HasAdjustmentAmountFloat returns a boolean if a field has been set.

### GetFormattedAdjustmentAmount

`func (o *OrderDataAttributes) GetFormattedAdjustmentAmount() string`

GetFormattedAdjustmentAmount returns the FormattedAdjustmentAmount field if non-nil, zero value otherwise.

### GetFormattedAdjustmentAmountOk

`func (o *OrderDataAttributes) GetFormattedAdjustmentAmountOk() (*string, bool)`

GetFormattedAdjustmentAmountOk returns a tuple with the FormattedAdjustmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAdjustmentAmount

`func (o *OrderDataAttributes) SetFormattedAdjustmentAmount(v string)`

SetFormattedAdjustmentAmount sets FormattedAdjustmentAmount field to given value.

### HasFormattedAdjustmentAmount

`func (o *OrderDataAttributes) HasFormattedAdjustmentAmount() bool`

HasFormattedAdjustmentAmount returns a boolean if a field has been set.

### GetGiftCardAmountCents

`func (o *OrderDataAttributes) GetGiftCardAmountCents() int32`

GetGiftCardAmountCents returns the GiftCardAmountCents field if non-nil, zero value otherwise.

### GetGiftCardAmountCentsOk

`func (o *OrderDataAttributes) GetGiftCardAmountCentsOk() (*int32, bool)`

GetGiftCardAmountCentsOk returns a tuple with the GiftCardAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardAmountCents

`func (o *OrderDataAttributes) SetGiftCardAmountCents(v int32)`

SetGiftCardAmountCents sets GiftCardAmountCents field to given value.

### HasGiftCardAmountCents

`func (o *OrderDataAttributes) HasGiftCardAmountCents() bool`

HasGiftCardAmountCents returns a boolean if a field has been set.

### GetGiftCardAmountFloat

`func (o *OrderDataAttributes) GetGiftCardAmountFloat() float32`

GetGiftCardAmountFloat returns the GiftCardAmountFloat field if non-nil, zero value otherwise.

### GetGiftCardAmountFloatOk

`func (o *OrderDataAttributes) GetGiftCardAmountFloatOk() (*float32, bool)`

GetGiftCardAmountFloatOk returns a tuple with the GiftCardAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardAmountFloat

`func (o *OrderDataAttributes) SetGiftCardAmountFloat(v float32)`

SetGiftCardAmountFloat sets GiftCardAmountFloat field to given value.

### HasGiftCardAmountFloat

`func (o *OrderDataAttributes) HasGiftCardAmountFloat() bool`

HasGiftCardAmountFloat returns a boolean if a field has been set.

### GetFormattedGiftCardAmount

`func (o *OrderDataAttributes) GetFormattedGiftCardAmount() string`

GetFormattedGiftCardAmount returns the FormattedGiftCardAmount field if non-nil, zero value otherwise.

### GetFormattedGiftCardAmountOk

`func (o *OrderDataAttributes) GetFormattedGiftCardAmountOk() (*string, bool)`

GetFormattedGiftCardAmountOk returns a tuple with the FormattedGiftCardAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedGiftCardAmount

`func (o *OrderDataAttributes) SetFormattedGiftCardAmount(v string)`

SetFormattedGiftCardAmount sets FormattedGiftCardAmount field to given value.

### HasFormattedGiftCardAmount

`func (o *OrderDataAttributes) HasFormattedGiftCardAmount() bool`

HasFormattedGiftCardAmount returns a boolean if a field has been set.

### GetTotalTaxAmountCents

`func (o *OrderDataAttributes) GetTotalTaxAmountCents() int32`

GetTotalTaxAmountCents returns the TotalTaxAmountCents field if non-nil, zero value otherwise.

### GetTotalTaxAmountCentsOk

`func (o *OrderDataAttributes) GetTotalTaxAmountCentsOk() (*int32, bool)`

GetTotalTaxAmountCentsOk returns a tuple with the TotalTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmountCents

`func (o *OrderDataAttributes) SetTotalTaxAmountCents(v int32)`

SetTotalTaxAmountCents sets TotalTaxAmountCents field to given value.

### HasTotalTaxAmountCents

`func (o *OrderDataAttributes) HasTotalTaxAmountCents() bool`

HasTotalTaxAmountCents returns a boolean if a field has been set.

### GetTotalTaxAmountFloat

`func (o *OrderDataAttributes) GetTotalTaxAmountFloat() float32`

GetTotalTaxAmountFloat returns the TotalTaxAmountFloat field if non-nil, zero value otherwise.

### GetTotalTaxAmountFloatOk

`func (o *OrderDataAttributes) GetTotalTaxAmountFloatOk() (*float32, bool)`

GetTotalTaxAmountFloatOk returns a tuple with the TotalTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmountFloat

`func (o *OrderDataAttributes) SetTotalTaxAmountFloat(v float32)`

SetTotalTaxAmountFloat sets TotalTaxAmountFloat field to given value.

### HasTotalTaxAmountFloat

`func (o *OrderDataAttributes) HasTotalTaxAmountFloat() bool`

HasTotalTaxAmountFloat returns a boolean if a field has been set.

### GetFormattedTotalTaxAmount

`func (o *OrderDataAttributes) GetFormattedTotalTaxAmount() string`

GetFormattedTotalTaxAmount returns the FormattedTotalTaxAmount field if non-nil, zero value otherwise.

### GetFormattedTotalTaxAmountOk

`func (o *OrderDataAttributes) GetFormattedTotalTaxAmountOk() (*string, bool)`

GetFormattedTotalTaxAmountOk returns a tuple with the FormattedTotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalTaxAmount

`func (o *OrderDataAttributes) SetFormattedTotalTaxAmount(v string)`

SetFormattedTotalTaxAmount sets FormattedTotalTaxAmount field to given value.

### HasFormattedTotalTaxAmount

`func (o *OrderDataAttributes) HasFormattedTotalTaxAmount() bool`

HasFormattedTotalTaxAmount returns a boolean if a field has been set.

### GetSubtotalTaxAmountCents

`func (o *OrderDataAttributes) GetSubtotalTaxAmountCents() int32`

GetSubtotalTaxAmountCents returns the SubtotalTaxAmountCents field if non-nil, zero value otherwise.

### GetSubtotalTaxAmountCentsOk

`func (o *OrderDataAttributes) GetSubtotalTaxAmountCentsOk() (*int32, bool)`

GetSubtotalTaxAmountCentsOk returns a tuple with the SubtotalTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalTaxAmountCents

`func (o *OrderDataAttributes) SetSubtotalTaxAmountCents(v int32)`

SetSubtotalTaxAmountCents sets SubtotalTaxAmountCents field to given value.

### HasSubtotalTaxAmountCents

`func (o *OrderDataAttributes) HasSubtotalTaxAmountCents() bool`

HasSubtotalTaxAmountCents returns a boolean if a field has been set.

### GetSubtotalTaxAmountFloat

`func (o *OrderDataAttributes) GetSubtotalTaxAmountFloat() float32`

GetSubtotalTaxAmountFloat returns the SubtotalTaxAmountFloat field if non-nil, zero value otherwise.

### GetSubtotalTaxAmountFloatOk

`func (o *OrderDataAttributes) GetSubtotalTaxAmountFloatOk() (*float32, bool)`

GetSubtotalTaxAmountFloatOk returns a tuple with the SubtotalTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalTaxAmountFloat

`func (o *OrderDataAttributes) SetSubtotalTaxAmountFloat(v float32)`

SetSubtotalTaxAmountFloat sets SubtotalTaxAmountFloat field to given value.

### HasSubtotalTaxAmountFloat

`func (o *OrderDataAttributes) HasSubtotalTaxAmountFloat() bool`

HasSubtotalTaxAmountFloat returns a boolean if a field has been set.

### GetFormattedSubtotalTaxAmount

`func (o *OrderDataAttributes) GetFormattedSubtotalTaxAmount() string`

GetFormattedSubtotalTaxAmount returns the FormattedSubtotalTaxAmount field if non-nil, zero value otherwise.

### GetFormattedSubtotalTaxAmountOk

`func (o *OrderDataAttributes) GetFormattedSubtotalTaxAmountOk() (*string, bool)`

GetFormattedSubtotalTaxAmountOk returns a tuple with the FormattedSubtotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSubtotalTaxAmount

`func (o *OrderDataAttributes) SetFormattedSubtotalTaxAmount(v string)`

SetFormattedSubtotalTaxAmount sets FormattedSubtotalTaxAmount field to given value.

### HasFormattedSubtotalTaxAmount

`func (o *OrderDataAttributes) HasFormattedSubtotalTaxAmount() bool`

HasFormattedSubtotalTaxAmount returns a boolean if a field has been set.

### GetShippingTaxAmountCents

`func (o *OrderDataAttributes) GetShippingTaxAmountCents() int32`

GetShippingTaxAmountCents returns the ShippingTaxAmountCents field if non-nil, zero value otherwise.

### GetShippingTaxAmountCentsOk

`func (o *OrderDataAttributes) GetShippingTaxAmountCentsOk() (*int32, bool)`

GetShippingTaxAmountCentsOk returns a tuple with the ShippingTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTaxAmountCents

`func (o *OrderDataAttributes) SetShippingTaxAmountCents(v int32)`

SetShippingTaxAmountCents sets ShippingTaxAmountCents field to given value.

### HasShippingTaxAmountCents

`func (o *OrderDataAttributes) HasShippingTaxAmountCents() bool`

HasShippingTaxAmountCents returns a boolean if a field has been set.

### GetShippingTaxAmountFloat

`func (o *OrderDataAttributes) GetShippingTaxAmountFloat() float32`

GetShippingTaxAmountFloat returns the ShippingTaxAmountFloat field if non-nil, zero value otherwise.

### GetShippingTaxAmountFloatOk

`func (o *OrderDataAttributes) GetShippingTaxAmountFloatOk() (*float32, bool)`

GetShippingTaxAmountFloatOk returns a tuple with the ShippingTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTaxAmountFloat

`func (o *OrderDataAttributes) SetShippingTaxAmountFloat(v float32)`

SetShippingTaxAmountFloat sets ShippingTaxAmountFloat field to given value.

### HasShippingTaxAmountFloat

`func (o *OrderDataAttributes) HasShippingTaxAmountFloat() bool`

HasShippingTaxAmountFloat returns a boolean if a field has been set.

### GetFormattedShippingTaxAmount

`func (o *OrderDataAttributes) GetFormattedShippingTaxAmount() string`

GetFormattedShippingTaxAmount returns the FormattedShippingTaxAmount field if non-nil, zero value otherwise.

### GetFormattedShippingTaxAmountOk

`func (o *OrderDataAttributes) GetFormattedShippingTaxAmountOk() (*string, bool)`

GetFormattedShippingTaxAmountOk returns a tuple with the FormattedShippingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedShippingTaxAmount

`func (o *OrderDataAttributes) SetFormattedShippingTaxAmount(v string)`

SetFormattedShippingTaxAmount sets FormattedShippingTaxAmount field to given value.

### HasFormattedShippingTaxAmount

`func (o *OrderDataAttributes) HasFormattedShippingTaxAmount() bool`

HasFormattedShippingTaxAmount returns a boolean if a field has been set.

### GetPaymentMethodTaxAmountCents

`func (o *OrderDataAttributes) GetPaymentMethodTaxAmountCents() int32`

GetPaymentMethodTaxAmountCents returns the PaymentMethodTaxAmountCents field if non-nil, zero value otherwise.

### GetPaymentMethodTaxAmountCentsOk

`func (o *OrderDataAttributes) GetPaymentMethodTaxAmountCentsOk() (*int32, bool)`

GetPaymentMethodTaxAmountCentsOk returns a tuple with the PaymentMethodTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxAmountCents

`func (o *OrderDataAttributes) SetPaymentMethodTaxAmountCents(v int32)`

SetPaymentMethodTaxAmountCents sets PaymentMethodTaxAmountCents field to given value.

### HasPaymentMethodTaxAmountCents

`func (o *OrderDataAttributes) HasPaymentMethodTaxAmountCents() bool`

HasPaymentMethodTaxAmountCents returns a boolean if a field has been set.

### GetPaymentMethodTaxAmountFloat

`func (o *OrderDataAttributes) GetPaymentMethodTaxAmountFloat() float32`

GetPaymentMethodTaxAmountFloat returns the PaymentMethodTaxAmountFloat field if non-nil, zero value otherwise.

### GetPaymentMethodTaxAmountFloatOk

`func (o *OrderDataAttributes) GetPaymentMethodTaxAmountFloatOk() (*float32, bool)`

GetPaymentMethodTaxAmountFloatOk returns a tuple with the PaymentMethodTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxAmountFloat

`func (o *OrderDataAttributes) SetPaymentMethodTaxAmountFloat(v float32)`

SetPaymentMethodTaxAmountFloat sets PaymentMethodTaxAmountFloat field to given value.

### HasPaymentMethodTaxAmountFloat

`func (o *OrderDataAttributes) HasPaymentMethodTaxAmountFloat() bool`

HasPaymentMethodTaxAmountFloat returns a boolean if a field has been set.

### GetFormattedPaymentMethodTaxAmount

`func (o *OrderDataAttributes) GetFormattedPaymentMethodTaxAmount() string`

GetFormattedPaymentMethodTaxAmount returns the FormattedPaymentMethodTaxAmount field if non-nil, zero value otherwise.

### GetFormattedPaymentMethodTaxAmountOk

`func (o *OrderDataAttributes) GetFormattedPaymentMethodTaxAmountOk() (*string, bool)`

GetFormattedPaymentMethodTaxAmountOk returns a tuple with the FormattedPaymentMethodTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPaymentMethodTaxAmount

`func (o *OrderDataAttributes) SetFormattedPaymentMethodTaxAmount(v string)`

SetFormattedPaymentMethodTaxAmount sets FormattedPaymentMethodTaxAmount field to given value.

### HasFormattedPaymentMethodTaxAmount

`func (o *OrderDataAttributes) HasFormattedPaymentMethodTaxAmount() bool`

HasFormattedPaymentMethodTaxAmount returns a boolean if a field has been set.

### GetAdjustmentTaxAmountCents

`func (o *OrderDataAttributes) GetAdjustmentTaxAmountCents() int32`

GetAdjustmentTaxAmountCents returns the AdjustmentTaxAmountCents field if non-nil, zero value otherwise.

### GetAdjustmentTaxAmountCentsOk

`func (o *OrderDataAttributes) GetAdjustmentTaxAmountCentsOk() (*int32, bool)`

GetAdjustmentTaxAmountCentsOk returns a tuple with the AdjustmentTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxAmountCents

`func (o *OrderDataAttributes) SetAdjustmentTaxAmountCents(v int32)`

SetAdjustmentTaxAmountCents sets AdjustmentTaxAmountCents field to given value.

### HasAdjustmentTaxAmountCents

`func (o *OrderDataAttributes) HasAdjustmentTaxAmountCents() bool`

HasAdjustmentTaxAmountCents returns a boolean if a field has been set.

### GetAdjustmentTaxAmountFloat

`func (o *OrderDataAttributes) GetAdjustmentTaxAmountFloat() float32`

GetAdjustmentTaxAmountFloat returns the AdjustmentTaxAmountFloat field if non-nil, zero value otherwise.

### GetAdjustmentTaxAmountFloatOk

`func (o *OrderDataAttributes) GetAdjustmentTaxAmountFloatOk() (*float32, bool)`

GetAdjustmentTaxAmountFloatOk returns a tuple with the AdjustmentTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxAmountFloat

`func (o *OrderDataAttributes) SetAdjustmentTaxAmountFloat(v float32)`

SetAdjustmentTaxAmountFloat sets AdjustmentTaxAmountFloat field to given value.

### HasAdjustmentTaxAmountFloat

`func (o *OrderDataAttributes) HasAdjustmentTaxAmountFloat() bool`

HasAdjustmentTaxAmountFloat returns a boolean if a field has been set.

### GetFormattedAdjustmentTaxAmount

`func (o *OrderDataAttributes) GetFormattedAdjustmentTaxAmount() string`

GetFormattedAdjustmentTaxAmount returns the FormattedAdjustmentTaxAmount field if non-nil, zero value otherwise.

### GetFormattedAdjustmentTaxAmountOk

`func (o *OrderDataAttributes) GetFormattedAdjustmentTaxAmountOk() (*string, bool)`

GetFormattedAdjustmentTaxAmountOk returns a tuple with the FormattedAdjustmentTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAdjustmentTaxAmount

`func (o *OrderDataAttributes) SetFormattedAdjustmentTaxAmount(v string)`

SetFormattedAdjustmentTaxAmount sets FormattedAdjustmentTaxAmount field to given value.

### HasFormattedAdjustmentTaxAmount

`func (o *OrderDataAttributes) HasFormattedAdjustmentTaxAmount() bool`

HasFormattedAdjustmentTaxAmount returns a boolean if a field has been set.

### GetTotalAmountCents

`func (o *OrderDataAttributes) GetTotalAmountCents() int32`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *OrderDataAttributes) GetTotalAmountCentsOk() (*int32, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *OrderDataAttributes) SetTotalAmountCents(v int32)`

SetTotalAmountCents sets TotalAmountCents field to given value.

### HasTotalAmountCents

`func (o *OrderDataAttributes) HasTotalAmountCents() bool`

HasTotalAmountCents returns a boolean if a field has been set.

### GetTotalAmountFloat

`func (o *OrderDataAttributes) GetTotalAmountFloat() float32`

GetTotalAmountFloat returns the TotalAmountFloat field if non-nil, zero value otherwise.

### GetTotalAmountFloatOk

`func (o *OrderDataAttributes) GetTotalAmountFloatOk() (*float32, bool)`

GetTotalAmountFloatOk returns a tuple with the TotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountFloat

`func (o *OrderDataAttributes) SetTotalAmountFloat(v float32)`

SetTotalAmountFloat sets TotalAmountFloat field to given value.

### HasTotalAmountFloat

`func (o *OrderDataAttributes) HasTotalAmountFloat() bool`

HasTotalAmountFloat returns a boolean if a field has been set.

### GetFormattedTotalAmount

`func (o *OrderDataAttributes) GetFormattedTotalAmount() string`

GetFormattedTotalAmount returns the FormattedTotalAmount field if non-nil, zero value otherwise.

### GetFormattedTotalAmountOk

`func (o *OrderDataAttributes) GetFormattedTotalAmountOk() (*string, bool)`

GetFormattedTotalAmountOk returns a tuple with the FormattedTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalAmount

`func (o *OrderDataAttributes) SetFormattedTotalAmount(v string)`

SetFormattedTotalAmount sets FormattedTotalAmount field to given value.

### HasFormattedTotalAmount

`func (o *OrderDataAttributes) HasFormattedTotalAmount() bool`

HasFormattedTotalAmount returns a boolean if a field has been set.

### GetTotalTaxableAmountCents

`func (o *OrderDataAttributes) GetTotalTaxableAmountCents() int32`

GetTotalTaxableAmountCents returns the TotalTaxableAmountCents field if non-nil, zero value otherwise.

### GetTotalTaxableAmountCentsOk

`func (o *OrderDataAttributes) GetTotalTaxableAmountCentsOk() (*int32, bool)`

GetTotalTaxableAmountCentsOk returns a tuple with the TotalTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxableAmountCents

`func (o *OrderDataAttributes) SetTotalTaxableAmountCents(v int32)`

SetTotalTaxableAmountCents sets TotalTaxableAmountCents field to given value.

### HasTotalTaxableAmountCents

`func (o *OrderDataAttributes) HasTotalTaxableAmountCents() bool`

HasTotalTaxableAmountCents returns a boolean if a field has been set.

### GetTotalTaxableAmountFloat

`func (o *OrderDataAttributes) GetTotalTaxableAmountFloat() float32`

GetTotalTaxableAmountFloat returns the TotalTaxableAmountFloat field if non-nil, zero value otherwise.

### GetTotalTaxableAmountFloatOk

`func (o *OrderDataAttributes) GetTotalTaxableAmountFloatOk() (*float32, bool)`

GetTotalTaxableAmountFloatOk returns a tuple with the TotalTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxableAmountFloat

`func (o *OrderDataAttributes) SetTotalTaxableAmountFloat(v float32)`

SetTotalTaxableAmountFloat sets TotalTaxableAmountFloat field to given value.

### HasTotalTaxableAmountFloat

`func (o *OrderDataAttributes) HasTotalTaxableAmountFloat() bool`

HasTotalTaxableAmountFloat returns a boolean if a field has been set.

### GetFormattedTotalTaxableAmount

`func (o *OrderDataAttributes) GetFormattedTotalTaxableAmount() string`

GetFormattedTotalTaxableAmount returns the FormattedTotalTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedTotalTaxableAmountOk

`func (o *OrderDataAttributes) GetFormattedTotalTaxableAmountOk() (*string, bool)`

GetFormattedTotalTaxableAmountOk returns a tuple with the FormattedTotalTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalTaxableAmount

`func (o *OrderDataAttributes) SetFormattedTotalTaxableAmount(v string)`

SetFormattedTotalTaxableAmount sets FormattedTotalTaxableAmount field to given value.

### HasFormattedTotalTaxableAmount

`func (o *OrderDataAttributes) HasFormattedTotalTaxableAmount() bool`

HasFormattedTotalTaxableAmount returns a boolean if a field has been set.

### GetSubtotalTaxableAmountCents

`func (o *OrderDataAttributes) GetSubtotalTaxableAmountCents() int32`

GetSubtotalTaxableAmountCents returns the SubtotalTaxableAmountCents field if non-nil, zero value otherwise.

### GetSubtotalTaxableAmountCentsOk

`func (o *OrderDataAttributes) GetSubtotalTaxableAmountCentsOk() (*int32, bool)`

GetSubtotalTaxableAmountCentsOk returns a tuple with the SubtotalTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalTaxableAmountCents

`func (o *OrderDataAttributes) SetSubtotalTaxableAmountCents(v int32)`

SetSubtotalTaxableAmountCents sets SubtotalTaxableAmountCents field to given value.

### HasSubtotalTaxableAmountCents

`func (o *OrderDataAttributes) HasSubtotalTaxableAmountCents() bool`

HasSubtotalTaxableAmountCents returns a boolean if a field has been set.

### GetSubtotalTaxableAmountFloat

`func (o *OrderDataAttributes) GetSubtotalTaxableAmountFloat() float32`

GetSubtotalTaxableAmountFloat returns the SubtotalTaxableAmountFloat field if non-nil, zero value otherwise.

### GetSubtotalTaxableAmountFloatOk

`func (o *OrderDataAttributes) GetSubtotalTaxableAmountFloatOk() (*float32, bool)`

GetSubtotalTaxableAmountFloatOk returns a tuple with the SubtotalTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalTaxableAmountFloat

`func (o *OrderDataAttributes) SetSubtotalTaxableAmountFloat(v float32)`

SetSubtotalTaxableAmountFloat sets SubtotalTaxableAmountFloat field to given value.

### HasSubtotalTaxableAmountFloat

`func (o *OrderDataAttributes) HasSubtotalTaxableAmountFloat() bool`

HasSubtotalTaxableAmountFloat returns a boolean if a field has been set.

### GetFormattedSubtotalTaxableAmount

`func (o *OrderDataAttributes) GetFormattedSubtotalTaxableAmount() string`

GetFormattedSubtotalTaxableAmount returns the FormattedSubtotalTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedSubtotalTaxableAmountOk

`func (o *OrderDataAttributes) GetFormattedSubtotalTaxableAmountOk() (*string, bool)`

GetFormattedSubtotalTaxableAmountOk returns a tuple with the FormattedSubtotalTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSubtotalTaxableAmount

`func (o *OrderDataAttributes) SetFormattedSubtotalTaxableAmount(v string)`

SetFormattedSubtotalTaxableAmount sets FormattedSubtotalTaxableAmount field to given value.

### HasFormattedSubtotalTaxableAmount

`func (o *OrderDataAttributes) HasFormattedSubtotalTaxableAmount() bool`

HasFormattedSubtotalTaxableAmount returns a boolean if a field has been set.

### GetShippingTaxableAmountCents

`func (o *OrderDataAttributes) GetShippingTaxableAmountCents() int32`

GetShippingTaxableAmountCents returns the ShippingTaxableAmountCents field if non-nil, zero value otherwise.

### GetShippingTaxableAmountCentsOk

`func (o *OrderDataAttributes) GetShippingTaxableAmountCentsOk() (*int32, bool)`

GetShippingTaxableAmountCentsOk returns a tuple with the ShippingTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTaxableAmountCents

`func (o *OrderDataAttributes) SetShippingTaxableAmountCents(v int32)`

SetShippingTaxableAmountCents sets ShippingTaxableAmountCents field to given value.

### HasShippingTaxableAmountCents

`func (o *OrderDataAttributes) HasShippingTaxableAmountCents() bool`

HasShippingTaxableAmountCents returns a boolean if a field has been set.

### GetShippingTaxableAmountFloat

`func (o *OrderDataAttributes) GetShippingTaxableAmountFloat() float32`

GetShippingTaxableAmountFloat returns the ShippingTaxableAmountFloat field if non-nil, zero value otherwise.

### GetShippingTaxableAmountFloatOk

`func (o *OrderDataAttributes) GetShippingTaxableAmountFloatOk() (*float32, bool)`

GetShippingTaxableAmountFloatOk returns a tuple with the ShippingTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTaxableAmountFloat

`func (o *OrderDataAttributes) SetShippingTaxableAmountFloat(v float32)`

SetShippingTaxableAmountFloat sets ShippingTaxableAmountFloat field to given value.

### HasShippingTaxableAmountFloat

`func (o *OrderDataAttributes) HasShippingTaxableAmountFloat() bool`

HasShippingTaxableAmountFloat returns a boolean if a field has been set.

### GetFormattedShippingTaxableAmount

`func (o *OrderDataAttributes) GetFormattedShippingTaxableAmount() string`

GetFormattedShippingTaxableAmount returns the FormattedShippingTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedShippingTaxableAmountOk

`func (o *OrderDataAttributes) GetFormattedShippingTaxableAmountOk() (*string, bool)`

GetFormattedShippingTaxableAmountOk returns a tuple with the FormattedShippingTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedShippingTaxableAmount

`func (o *OrderDataAttributes) SetFormattedShippingTaxableAmount(v string)`

SetFormattedShippingTaxableAmount sets FormattedShippingTaxableAmount field to given value.

### HasFormattedShippingTaxableAmount

`func (o *OrderDataAttributes) HasFormattedShippingTaxableAmount() bool`

HasFormattedShippingTaxableAmount returns a boolean if a field has been set.

### GetPaymentMethodTaxableAmountCents

`func (o *OrderDataAttributes) GetPaymentMethodTaxableAmountCents() int32`

GetPaymentMethodTaxableAmountCents returns the PaymentMethodTaxableAmountCents field if non-nil, zero value otherwise.

### GetPaymentMethodTaxableAmountCentsOk

`func (o *OrderDataAttributes) GetPaymentMethodTaxableAmountCentsOk() (*int32, bool)`

GetPaymentMethodTaxableAmountCentsOk returns a tuple with the PaymentMethodTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxableAmountCents

`func (o *OrderDataAttributes) SetPaymentMethodTaxableAmountCents(v int32)`

SetPaymentMethodTaxableAmountCents sets PaymentMethodTaxableAmountCents field to given value.

### HasPaymentMethodTaxableAmountCents

`func (o *OrderDataAttributes) HasPaymentMethodTaxableAmountCents() bool`

HasPaymentMethodTaxableAmountCents returns a boolean if a field has been set.

### GetPaymentMethodTaxableAmountFloat

`func (o *OrderDataAttributes) GetPaymentMethodTaxableAmountFloat() float32`

GetPaymentMethodTaxableAmountFloat returns the PaymentMethodTaxableAmountFloat field if non-nil, zero value otherwise.

### GetPaymentMethodTaxableAmountFloatOk

`func (o *OrderDataAttributes) GetPaymentMethodTaxableAmountFloatOk() (*float32, bool)`

GetPaymentMethodTaxableAmountFloatOk returns a tuple with the PaymentMethodTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxableAmountFloat

`func (o *OrderDataAttributes) SetPaymentMethodTaxableAmountFloat(v float32)`

SetPaymentMethodTaxableAmountFloat sets PaymentMethodTaxableAmountFloat field to given value.

### HasPaymentMethodTaxableAmountFloat

`func (o *OrderDataAttributes) HasPaymentMethodTaxableAmountFloat() bool`

HasPaymentMethodTaxableAmountFloat returns a boolean if a field has been set.

### GetFormattedPaymentMethodTaxableAmount

`func (o *OrderDataAttributes) GetFormattedPaymentMethodTaxableAmount() string`

GetFormattedPaymentMethodTaxableAmount returns the FormattedPaymentMethodTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedPaymentMethodTaxableAmountOk

`func (o *OrderDataAttributes) GetFormattedPaymentMethodTaxableAmountOk() (*string, bool)`

GetFormattedPaymentMethodTaxableAmountOk returns a tuple with the FormattedPaymentMethodTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPaymentMethodTaxableAmount

`func (o *OrderDataAttributes) SetFormattedPaymentMethodTaxableAmount(v string)`

SetFormattedPaymentMethodTaxableAmount sets FormattedPaymentMethodTaxableAmount field to given value.

### HasFormattedPaymentMethodTaxableAmount

`func (o *OrderDataAttributes) HasFormattedPaymentMethodTaxableAmount() bool`

HasFormattedPaymentMethodTaxableAmount returns a boolean if a field has been set.

### GetAdjustmentTaxableAmountCents

`func (o *OrderDataAttributes) GetAdjustmentTaxableAmountCents() int32`

GetAdjustmentTaxableAmountCents returns the AdjustmentTaxableAmountCents field if non-nil, zero value otherwise.

### GetAdjustmentTaxableAmountCentsOk

`func (o *OrderDataAttributes) GetAdjustmentTaxableAmountCentsOk() (*int32, bool)`

GetAdjustmentTaxableAmountCentsOk returns a tuple with the AdjustmentTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxableAmountCents

`func (o *OrderDataAttributes) SetAdjustmentTaxableAmountCents(v int32)`

SetAdjustmentTaxableAmountCents sets AdjustmentTaxableAmountCents field to given value.

### HasAdjustmentTaxableAmountCents

`func (o *OrderDataAttributes) HasAdjustmentTaxableAmountCents() bool`

HasAdjustmentTaxableAmountCents returns a boolean if a field has been set.

### GetAdjustmentTaxableAmountFloat

`func (o *OrderDataAttributes) GetAdjustmentTaxableAmountFloat() float32`

GetAdjustmentTaxableAmountFloat returns the AdjustmentTaxableAmountFloat field if non-nil, zero value otherwise.

### GetAdjustmentTaxableAmountFloatOk

`func (o *OrderDataAttributes) GetAdjustmentTaxableAmountFloatOk() (*float32, bool)`

GetAdjustmentTaxableAmountFloatOk returns a tuple with the AdjustmentTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxableAmountFloat

`func (o *OrderDataAttributes) SetAdjustmentTaxableAmountFloat(v float32)`

SetAdjustmentTaxableAmountFloat sets AdjustmentTaxableAmountFloat field to given value.

### HasAdjustmentTaxableAmountFloat

`func (o *OrderDataAttributes) HasAdjustmentTaxableAmountFloat() bool`

HasAdjustmentTaxableAmountFloat returns a boolean if a field has been set.

### GetFormattedAdjustmentTaxableAmount

`func (o *OrderDataAttributes) GetFormattedAdjustmentTaxableAmount() string`

GetFormattedAdjustmentTaxableAmount returns the FormattedAdjustmentTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedAdjustmentTaxableAmountOk

`func (o *OrderDataAttributes) GetFormattedAdjustmentTaxableAmountOk() (*string, bool)`

GetFormattedAdjustmentTaxableAmountOk returns a tuple with the FormattedAdjustmentTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAdjustmentTaxableAmount

`func (o *OrderDataAttributes) SetFormattedAdjustmentTaxableAmount(v string)`

SetFormattedAdjustmentTaxableAmount sets FormattedAdjustmentTaxableAmount field to given value.

### HasFormattedAdjustmentTaxableAmount

`func (o *OrderDataAttributes) HasFormattedAdjustmentTaxableAmount() bool`

HasFormattedAdjustmentTaxableAmount returns a boolean if a field has been set.

### GetTotalAmountWithTaxesCents

`func (o *OrderDataAttributes) GetTotalAmountWithTaxesCents() int32`

GetTotalAmountWithTaxesCents returns the TotalAmountWithTaxesCents field if non-nil, zero value otherwise.

### GetTotalAmountWithTaxesCentsOk

`func (o *OrderDataAttributes) GetTotalAmountWithTaxesCentsOk() (*int32, bool)`

GetTotalAmountWithTaxesCentsOk returns a tuple with the TotalAmountWithTaxesCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountWithTaxesCents

`func (o *OrderDataAttributes) SetTotalAmountWithTaxesCents(v int32)`

SetTotalAmountWithTaxesCents sets TotalAmountWithTaxesCents field to given value.

### HasTotalAmountWithTaxesCents

`func (o *OrderDataAttributes) HasTotalAmountWithTaxesCents() bool`

HasTotalAmountWithTaxesCents returns a boolean if a field has been set.

### GetTotalAmountWithTaxesFloat

`func (o *OrderDataAttributes) GetTotalAmountWithTaxesFloat() float32`

GetTotalAmountWithTaxesFloat returns the TotalAmountWithTaxesFloat field if non-nil, zero value otherwise.

### GetTotalAmountWithTaxesFloatOk

`func (o *OrderDataAttributes) GetTotalAmountWithTaxesFloatOk() (*float32, bool)`

GetTotalAmountWithTaxesFloatOk returns a tuple with the TotalAmountWithTaxesFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountWithTaxesFloat

`func (o *OrderDataAttributes) SetTotalAmountWithTaxesFloat(v float32)`

SetTotalAmountWithTaxesFloat sets TotalAmountWithTaxesFloat field to given value.

### HasTotalAmountWithTaxesFloat

`func (o *OrderDataAttributes) HasTotalAmountWithTaxesFloat() bool`

HasTotalAmountWithTaxesFloat returns a boolean if a field has been set.

### GetFormattedTotalAmountWithTaxes

`func (o *OrderDataAttributes) GetFormattedTotalAmountWithTaxes() string`

GetFormattedTotalAmountWithTaxes returns the FormattedTotalAmountWithTaxes field if non-nil, zero value otherwise.

### GetFormattedTotalAmountWithTaxesOk

`func (o *OrderDataAttributes) GetFormattedTotalAmountWithTaxesOk() (*string, bool)`

GetFormattedTotalAmountWithTaxesOk returns a tuple with the FormattedTotalAmountWithTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalAmountWithTaxes

`func (o *OrderDataAttributes) SetFormattedTotalAmountWithTaxes(v string)`

SetFormattedTotalAmountWithTaxes sets FormattedTotalAmountWithTaxes field to given value.

### HasFormattedTotalAmountWithTaxes

`func (o *OrderDataAttributes) HasFormattedTotalAmountWithTaxes() bool`

HasFormattedTotalAmountWithTaxes returns a boolean if a field has been set.

### GetFeesAmountCents

`func (o *OrderDataAttributes) GetFeesAmountCents() int32`

GetFeesAmountCents returns the FeesAmountCents field if non-nil, zero value otherwise.

### GetFeesAmountCentsOk

`func (o *OrderDataAttributes) GetFeesAmountCentsOk() (*int32, bool)`

GetFeesAmountCentsOk returns a tuple with the FeesAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesAmountCents

`func (o *OrderDataAttributes) SetFeesAmountCents(v int32)`

SetFeesAmountCents sets FeesAmountCents field to given value.

### HasFeesAmountCents

`func (o *OrderDataAttributes) HasFeesAmountCents() bool`

HasFeesAmountCents returns a boolean if a field has been set.

### GetFeesAmountFloat

`func (o *OrderDataAttributes) GetFeesAmountFloat() float32`

GetFeesAmountFloat returns the FeesAmountFloat field if non-nil, zero value otherwise.

### GetFeesAmountFloatOk

`func (o *OrderDataAttributes) GetFeesAmountFloatOk() (*float32, bool)`

GetFeesAmountFloatOk returns a tuple with the FeesAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesAmountFloat

`func (o *OrderDataAttributes) SetFeesAmountFloat(v float32)`

SetFeesAmountFloat sets FeesAmountFloat field to given value.

### HasFeesAmountFloat

`func (o *OrderDataAttributes) HasFeesAmountFloat() bool`

HasFeesAmountFloat returns a boolean if a field has been set.

### GetFormattedFeesAmount

`func (o *OrderDataAttributes) GetFormattedFeesAmount() string`

GetFormattedFeesAmount returns the FormattedFeesAmount field if non-nil, zero value otherwise.

### GetFormattedFeesAmountOk

`func (o *OrderDataAttributes) GetFormattedFeesAmountOk() (*string, bool)`

GetFormattedFeesAmountOk returns a tuple with the FormattedFeesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedFeesAmount

`func (o *OrderDataAttributes) SetFormattedFeesAmount(v string)`

SetFormattedFeesAmount sets FormattedFeesAmount field to given value.

### HasFormattedFeesAmount

`func (o *OrderDataAttributes) HasFormattedFeesAmount() bool`

HasFormattedFeesAmount returns a boolean if a field has been set.

### GetDutyAmountCents

`func (o *OrderDataAttributes) GetDutyAmountCents() int32`

GetDutyAmountCents returns the DutyAmountCents field if non-nil, zero value otherwise.

### GetDutyAmountCentsOk

`func (o *OrderDataAttributes) GetDutyAmountCentsOk() (*int32, bool)`

GetDutyAmountCentsOk returns a tuple with the DutyAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyAmountCents

`func (o *OrderDataAttributes) SetDutyAmountCents(v int32)`

SetDutyAmountCents sets DutyAmountCents field to given value.

### HasDutyAmountCents

`func (o *OrderDataAttributes) HasDutyAmountCents() bool`

HasDutyAmountCents returns a boolean if a field has been set.

### GetDutyAmountFloat

`func (o *OrderDataAttributes) GetDutyAmountFloat() float32`

GetDutyAmountFloat returns the DutyAmountFloat field if non-nil, zero value otherwise.

### GetDutyAmountFloatOk

`func (o *OrderDataAttributes) GetDutyAmountFloatOk() (*float32, bool)`

GetDutyAmountFloatOk returns a tuple with the DutyAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyAmountFloat

`func (o *OrderDataAttributes) SetDutyAmountFloat(v float32)`

SetDutyAmountFloat sets DutyAmountFloat field to given value.

### HasDutyAmountFloat

`func (o *OrderDataAttributes) HasDutyAmountFloat() bool`

HasDutyAmountFloat returns a boolean if a field has been set.

### GetFormattedDutyAmount

`func (o *OrderDataAttributes) GetFormattedDutyAmount() string`

GetFormattedDutyAmount returns the FormattedDutyAmount field if non-nil, zero value otherwise.

### GetFormattedDutyAmountOk

`func (o *OrderDataAttributes) GetFormattedDutyAmountOk() (*string, bool)`

GetFormattedDutyAmountOk returns a tuple with the FormattedDutyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedDutyAmount

`func (o *OrderDataAttributes) SetFormattedDutyAmount(v string)`

SetFormattedDutyAmount sets FormattedDutyAmount field to given value.

### HasFormattedDutyAmount

`func (o *OrderDataAttributes) HasFormattedDutyAmount() bool`

HasFormattedDutyAmount returns a boolean if a field has been set.

### GetSkusCount

`func (o *OrderDataAttributes) GetSkusCount() int32`

GetSkusCount returns the SkusCount field if non-nil, zero value otherwise.

### GetSkusCountOk

`func (o *OrderDataAttributes) GetSkusCountOk() (*int32, bool)`

GetSkusCountOk returns a tuple with the SkusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusCount

`func (o *OrderDataAttributes) SetSkusCount(v int32)`

SetSkusCount sets SkusCount field to given value.

### HasSkusCount

`func (o *OrderDataAttributes) HasSkusCount() bool`

HasSkusCount returns a boolean if a field has been set.

### GetLineItemOptionsCount

`func (o *OrderDataAttributes) GetLineItemOptionsCount() int32`

GetLineItemOptionsCount returns the LineItemOptionsCount field if non-nil, zero value otherwise.

### GetLineItemOptionsCountOk

`func (o *OrderDataAttributes) GetLineItemOptionsCountOk() (*int32, bool)`

GetLineItemOptionsCountOk returns a tuple with the LineItemOptionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemOptionsCount

`func (o *OrderDataAttributes) SetLineItemOptionsCount(v int32)`

SetLineItemOptionsCount sets LineItemOptionsCount field to given value.

### HasLineItemOptionsCount

`func (o *OrderDataAttributes) HasLineItemOptionsCount() bool`

HasLineItemOptionsCount returns a boolean if a field has been set.

### GetShipmentsCount

`func (o *OrderDataAttributes) GetShipmentsCount() int32`

GetShipmentsCount returns the ShipmentsCount field if non-nil, zero value otherwise.

### GetShipmentsCountOk

`func (o *OrderDataAttributes) GetShipmentsCountOk() (*int32, bool)`

GetShipmentsCountOk returns a tuple with the ShipmentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentsCount

`func (o *OrderDataAttributes) SetShipmentsCount(v int32)`

SetShipmentsCount sets ShipmentsCount field to given value.

### HasShipmentsCount

`func (o *OrderDataAttributes) HasShipmentsCount() bool`

HasShipmentsCount returns a boolean if a field has been set.

### GetTaxCalculationsCount

`func (o *OrderDataAttributes) GetTaxCalculationsCount() int32`

GetTaxCalculationsCount returns the TaxCalculationsCount field if non-nil, zero value otherwise.

### GetTaxCalculationsCountOk

`func (o *OrderDataAttributes) GetTaxCalculationsCountOk() (*int32, bool)`

GetTaxCalculationsCountOk returns a tuple with the TaxCalculationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationsCount

`func (o *OrderDataAttributes) SetTaxCalculationsCount(v int32)`

SetTaxCalculationsCount sets TaxCalculationsCount field to given value.

### HasTaxCalculationsCount

`func (o *OrderDataAttributes) HasTaxCalculationsCount() bool`

HasTaxCalculationsCount returns a boolean if a field has been set.

### GetPaymentSourceDetails

`func (o *OrderDataAttributes) GetPaymentSourceDetails() map[string]interface{}`

GetPaymentSourceDetails returns the PaymentSourceDetails field if non-nil, zero value otherwise.

### GetPaymentSourceDetailsOk

`func (o *OrderDataAttributes) GetPaymentSourceDetailsOk() (*map[string]interface{}, bool)`

GetPaymentSourceDetailsOk returns a tuple with the PaymentSourceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceDetails

`func (o *OrderDataAttributes) SetPaymentSourceDetails(v map[string]interface{})`

SetPaymentSourceDetails sets PaymentSourceDetails field to given value.

### HasPaymentSourceDetails

`func (o *OrderDataAttributes) HasPaymentSourceDetails() bool`

HasPaymentSourceDetails returns a boolean if a field has been set.

### GetToken

`func (o *OrderDataAttributes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OrderDataAttributes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OrderDataAttributes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *OrderDataAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetCartUrl

`func (o *OrderDataAttributes) GetCartUrl() string`

GetCartUrl returns the CartUrl field if non-nil, zero value otherwise.

### GetCartUrlOk

`func (o *OrderDataAttributes) GetCartUrlOk() (*string, bool)`

GetCartUrlOk returns a tuple with the CartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartUrl

`func (o *OrderDataAttributes) SetCartUrl(v string)`

SetCartUrl sets CartUrl field to given value.

### HasCartUrl

`func (o *OrderDataAttributes) HasCartUrl() bool`

HasCartUrl returns a boolean if a field has been set.

### GetReturnUrl

`func (o *OrderDataAttributes) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *OrderDataAttributes) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *OrderDataAttributes) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *OrderDataAttributes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetTermsUrl

`func (o *OrderDataAttributes) GetTermsUrl() string`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *OrderDataAttributes) GetTermsUrlOk() (*string, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *OrderDataAttributes) SetTermsUrl(v string)`

SetTermsUrl sets TermsUrl field to given value.

### HasTermsUrl

`func (o *OrderDataAttributes) HasTermsUrl() bool`

HasTermsUrl returns a boolean if a field has been set.

### GetPrivacyUrl

`func (o *OrderDataAttributes) GetPrivacyUrl() string`

GetPrivacyUrl returns the PrivacyUrl field if non-nil, zero value otherwise.

### GetPrivacyUrlOk

`func (o *OrderDataAttributes) GetPrivacyUrlOk() (*string, bool)`

GetPrivacyUrlOk returns a tuple with the PrivacyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyUrl

`func (o *OrderDataAttributes) SetPrivacyUrl(v string)`

SetPrivacyUrl sets PrivacyUrl field to given value.

### HasPrivacyUrl

`func (o *OrderDataAttributes) HasPrivacyUrl() bool`

HasPrivacyUrl returns a boolean if a field has been set.

### GetCheckoutUrl

`func (o *OrderDataAttributes) GetCheckoutUrl() string`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *OrderDataAttributes) GetCheckoutUrlOk() (*string, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *OrderDataAttributes) SetCheckoutUrl(v string)`

SetCheckoutUrl sets CheckoutUrl field to given value.

### HasCheckoutUrl

`func (o *OrderDataAttributes) HasCheckoutUrl() bool`

HasCheckoutUrl returns a boolean if a field has been set.

### GetPlacedAt

`func (o *OrderDataAttributes) GetPlacedAt() string`

GetPlacedAt returns the PlacedAt field if non-nil, zero value otherwise.

### GetPlacedAtOk

`func (o *OrderDataAttributes) GetPlacedAtOk() (*string, bool)`

GetPlacedAtOk returns a tuple with the PlacedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacedAt

`func (o *OrderDataAttributes) SetPlacedAt(v string)`

SetPlacedAt sets PlacedAt field to given value.

### HasPlacedAt

`func (o *OrderDataAttributes) HasPlacedAt() bool`

HasPlacedAt returns a boolean if a field has been set.

### GetApprovedAt

`func (o *OrderDataAttributes) GetApprovedAt() string`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *OrderDataAttributes) GetApprovedAtOk() (*string, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *OrderDataAttributes) SetApprovedAt(v string)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *OrderDataAttributes) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### GetCancelledAt

`func (o *OrderDataAttributes) GetCancelledAt() string`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *OrderDataAttributes) GetCancelledAtOk() (*string, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *OrderDataAttributes) SetCancelledAt(v string)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *OrderDataAttributes) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### GetPaymentUpdatedAt

`func (o *OrderDataAttributes) GetPaymentUpdatedAt() string`

GetPaymentUpdatedAt returns the PaymentUpdatedAt field if non-nil, zero value otherwise.

### GetPaymentUpdatedAtOk

`func (o *OrderDataAttributes) GetPaymentUpdatedAtOk() (*string, bool)`

GetPaymentUpdatedAtOk returns a tuple with the PaymentUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentUpdatedAt

`func (o *OrderDataAttributes) SetPaymentUpdatedAt(v string)`

SetPaymentUpdatedAt sets PaymentUpdatedAt field to given value.

### HasPaymentUpdatedAt

`func (o *OrderDataAttributes) HasPaymentUpdatedAt() bool`

HasPaymentUpdatedAt returns a boolean if a field has been set.

### GetFulfillmentUpdatedAt

`func (o *OrderDataAttributes) GetFulfillmentUpdatedAt() string`

GetFulfillmentUpdatedAt returns the FulfillmentUpdatedAt field if non-nil, zero value otherwise.

### GetFulfillmentUpdatedAtOk

`func (o *OrderDataAttributes) GetFulfillmentUpdatedAtOk() (*string, bool)`

GetFulfillmentUpdatedAtOk returns a tuple with the FulfillmentUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentUpdatedAt

`func (o *OrderDataAttributes) SetFulfillmentUpdatedAt(v string)`

SetFulfillmentUpdatedAt sets FulfillmentUpdatedAt field to given value.

### HasFulfillmentUpdatedAt

`func (o *OrderDataAttributes) HasFulfillmentUpdatedAt() bool`

HasFulfillmentUpdatedAt returns a boolean if a field has been set.

### GetRefreshedAt

`func (o *OrderDataAttributes) GetRefreshedAt() string`

GetRefreshedAt returns the RefreshedAt field if non-nil, zero value otherwise.

### GetRefreshedAtOk

`func (o *OrderDataAttributes) GetRefreshedAtOk() (*string, bool)`

GetRefreshedAtOk returns a tuple with the RefreshedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshedAt

`func (o *OrderDataAttributes) SetRefreshedAt(v string)`

SetRefreshedAt sets RefreshedAt field to given value.

### HasRefreshedAt

`func (o *OrderDataAttributes) HasRefreshedAt() bool`

HasRefreshedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *OrderDataAttributes) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *OrderDataAttributes) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *OrderDataAttributes) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *OrderDataAttributes) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OrderDataAttributes) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrderDataAttributes) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrderDataAttributes) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrderDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *OrderDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrderDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrderDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrderDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrderDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrderDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrderDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *OrderDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *OrderDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *OrderDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *OrderDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *OrderDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *OrderDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *OrderDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *OrderDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *OrderDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


