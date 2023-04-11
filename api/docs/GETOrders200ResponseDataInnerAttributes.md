# GETOrders200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | Unique identifier for the order (numeric). | [optional] 
**Autorefresh** | Pointer to **interface{}** | Save this attribute as &#39;false&#39; if you want prevent the order to be refreshed automatically at each change (much faster). | [optional] 
**Status** | Pointer to **interface{}** | The order status. One of &#39;draft&#39; (default), &#39;pending&#39;, &#39;placed&#39;, &#39;approved&#39;, or &#39;cancelled&#39;. | [optional] 
**PaymentStatus** | Pointer to **interface{}** | The order&#39;s payment status. One of &#39;unpaid&#39; (default), &#39;authorized&#39;, &#39;partially_authorized&#39;, &#39;paid&#39;, &#39;partially_paid&#39;, &#39;voided&#39;, &#39;partially_voided&#39;, &#39;refunded&#39;, &#39;partially_refunded&#39; or &#39;free&#39;. | [optional] 
**FulfillmentStatus** | Pointer to **interface{}** | The order&#39;s fulfillment status. One of &#39;unfulfilled&#39; (default), &#39;in_progress&#39;, &#39;fulfilled&#39;, or &#39;not_required&#39;. | [optional] 
**Guest** | Pointer to **interface{}** | Indicates if the order has been placed as guest. | [optional] 
**Editable** | Pointer to **interface{}** | Indicates if the order can be edited. | [optional] 
**CustomerEmail** | Pointer to **interface{}** | The email address of the associated customer. When creating or updating an order, this is a shortcut to find or create the associated customer by email. | [optional] 
**LanguageCode** | Pointer to **interface{}** | The preferred language code (ISO 639-1) to be used when communicating with the customer. This can be useful when sending the order to 3rd party marketing tools and CRMs. If the language is supported, the hosted checkout will be localized accordingly. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard, automatically inherited from the order&#39;s market. | [optional] 
**TaxIncluded** | Pointer to **interface{}** | Indicates if taxes are included in the order amounts, automatically inherited from the order&#39;s price list. | [optional] 
**TaxRate** | Pointer to **interface{}** | The tax rate for this order (if calculated). | [optional] 
**FreightTaxable** | Pointer to **interface{}** | Indicates if taxes are applied to shipping costs. | [optional] 
**RequiresBillingInfo** | Pointer to **interface{}** | Indicates if the billing address associated to this order requires billing info to be present. | [optional] 
**CountryCode** | Pointer to **interface{}** | The international 2-letter country code as defined by the ISO 3166-1 standard, automatically inherited from the order&#39;s shipping address. | [optional] 
**ShippingCountryCodeLock** | Pointer to **interface{}** | The country code that you want the shipping address to be locked to. This can be useful to make sure the shipping address belongs to a given shipping country, e.g. the one selected in a country selector page. | [optional] 
**CouponCode** | Pointer to **interface{}** | The coupon code to be used for the order. If valid, it triggers a promotion adding a discount line item to the order. | [optional] 
**GiftCardCode** | Pointer to **interface{}** | The gift card code (at least the first 8 characters) to be used for the order. If valid, it uses the gift card balance to pay for the order. | [optional] 
**GiftCardOrCouponCode** | Pointer to **interface{}** | The gift card or coupon code (at least the first 8 characters) to be used for the order. If a gift card mathes, it uses the gift card balance to pay for the order. Otherwise it tries to find a valid coupon code and applies the associated discount. | [optional] 
**SubtotalAmountCents** | Pointer to **interface{}** | The sum of all the SKU line items total amounts, in cents. | [optional] 
**SubtotalAmountFloat** | Pointer to **interface{}** | The sum of all the SKU line items total amounts, float. | [optional] 
**FormattedSubtotalAmount** | Pointer to **interface{}** | The sum of all the SKU line items total amounts, formatted. | [optional] 
**ShippingAmountCents** | Pointer to **interface{}** | The sum of all the shipping costs, in cents. | [optional] 
**ShippingAmountFloat** | Pointer to **interface{}** | The sum of all the shipping costs, float. | [optional] 
**FormattedShippingAmount** | Pointer to **interface{}** | The sum of all the shipping costs, formatted. | [optional] 
**PaymentMethodAmountCents** | Pointer to **interface{}** | The payment method costs, in cents. | [optional] 
**PaymentMethodAmountFloat** | Pointer to **interface{}** | The payment method costs, float. | [optional] 
**FormattedPaymentMethodAmount** | Pointer to **interface{}** | The payment method costs, formatted. | [optional] 
**DiscountAmountCents** | Pointer to **interface{}** | The sum of all the discounts applied to the order, in cents (negative amount). | [optional] 
**DiscountAmountFloat** | Pointer to **interface{}** | The sum of all the discounts applied to the order, float. | [optional] 
**FormattedDiscountAmount** | Pointer to **interface{}** | The sum of all the discounts applied to the order, formatted. | [optional] 
**AdjustmentAmountCents** | Pointer to **interface{}** | The sum of all the adjustments applied to the order, in cents. | [optional] 
**AdjustmentAmountFloat** | Pointer to **interface{}** | The sum of all the adjustments applied to the order, float. | [optional] 
**FormattedAdjustmentAmount** | Pointer to **interface{}** | The sum of all the adjustments applied to the order, formatted. | [optional] 
**GiftCardAmountCents** | Pointer to **interface{}** | The sum of all the gift_cards applied to the order, in cents. | [optional] 
**GiftCardAmountFloat** | Pointer to **interface{}** | The sum of all the gift_cards applied to the order, float. | [optional] 
**FormattedGiftCardAmount** | Pointer to **interface{}** | The sum of all the gift_cards applied to the order, formatted. | [optional] 
**TotalTaxAmountCents** | Pointer to **interface{}** | The sum of all the taxes applied to the order, in cents. | [optional] 
**TotalTaxAmountFloat** | Pointer to **interface{}** | The sum of all the taxes applied to the order, float. | [optional] 
**FormattedTotalTaxAmount** | Pointer to **interface{}** | The sum of all the taxes applied to the order, formatted. | [optional] 
**SubtotalTaxAmountCents** | Pointer to **interface{}** | The taxes applied to the order&#39;s subtotal, in cents. | [optional] 
**SubtotalTaxAmountFloat** | Pointer to **interface{}** | The taxes applied to the order&#39;s subtotal, float. | [optional] 
**FormattedSubtotalTaxAmount** | Pointer to **interface{}** | The taxes applied to the order&#39;s subtotal, formatted. | [optional] 
**ShippingTaxAmountCents** | Pointer to **interface{}** | The taxes applied to the order&#39;s shipping costs, in cents. | [optional] 
**ShippingTaxAmountFloat** | Pointer to **interface{}** | The taxes applied to the order&#39;s shipping costs, float. | [optional] 
**FormattedShippingTaxAmount** | Pointer to **interface{}** | The taxes applied to the order&#39;s shipping costs, formatted. | [optional] 
**PaymentMethodTaxAmountCents** | Pointer to **interface{}** | The taxes applied to the order&#39;s payment method costs, in cents. | [optional] 
**PaymentMethodTaxAmountFloat** | Pointer to **interface{}** | The taxes applied to the order&#39;s payment method costs, float. | [optional] 
**FormattedPaymentMethodTaxAmount** | Pointer to **interface{}** | The taxes applied to the order&#39;s payment method costs, formatted. | [optional] 
**AdjustmentTaxAmountCents** | Pointer to **interface{}** | The taxes applied to the order adjustments, in cents. | [optional] 
**AdjustmentTaxAmountFloat** | Pointer to **interface{}** | The taxes applied to the order adjustments, float. | [optional] 
**FormattedAdjustmentTaxAmount** | Pointer to **interface{}** | The taxes applied to the order adjustments, formatted. | [optional] 
**TotalAmountCents** | Pointer to **interface{}** | The order&#39;s total amount, in cents. | [optional] 
**TotalAmountFloat** | Pointer to **interface{}** | The order&#39;s total amount, float. | [optional] 
**FormattedTotalAmount** | Pointer to **interface{}** | The order&#39;s total amount, formatted. | [optional] 
**TotalTaxableAmountCents** | Pointer to **interface{}** | The order&#39;s total taxable amount, in cents (without discounts). | [optional] 
**TotalTaxableAmountFloat** | Pointer to **interface{}** | The order&#39;s total taxable amount, float. | [optional] 
**FormattedTotalTaxableAmount** | Pointer to **interface{}** | The order&#39;s total taxable amount, formatted. | [optional] 
**SubtotalTaxableAmountCents** | Pointer to **interface{}** | The order&#39;s subtotal taxable amount, in cents (equal to subtotal_amount_cents when prices don&#39;t include taxes). | [optional] 
**SubtotalTaxableAmountFloat** | Pointer to **interface{}** | The order&#39;s subtotal taxable amount, float. | [optional] 
**FormattedSubtotalTaxableAmount** | Pointer to **interface{}** | The order&#39;s subtotal taxable amount, formatted. | [optional] 
**ShippingTaxableAmountCents** | Pointer to **interface{}** | The order&#39;s shipping taxable amount, in cents (equal to shipping_amount_cents when prices don&#39;t include taxes). | [optional] 
**ShippingTaxableAmountFloat** | Pointer to **interface{}** | The order&#39;s shipping taxable amount, float. | [optional] 
**FormattedShippingTaxableAmount** | Pointer to **interface{}** | The order&#39;s shipping taxable amount, formatted. | [optional] 
**PaymentMethodTaxableAmountCents** | Pointer to **interface{}** | The order&#39;s payment method taxable amount, in cents (equal to payment_method_amount_cents when prices don&#39;t include taxes). | [optional] 
**PaymentMethodTaxableAmountFloat** | Pointer to **interface{}** | The order&#39;s payment method taxable amount, float. | [optional] 
**FormattedPaymentMethodTaxableAmount** | Pointer to **interface{}** | The order&#39;s payment method taxable amount, formatted. | [optional] 
**AdjustmentTaxableAmountCents** | Pointer to **interface{}** | The order&#39;s adjustment taxable amount, in cents (equal to discount_adjustment_cents when prices don&#39;t include taxes). | [optional] 
**AdjustmentTaxableAmountFloat** | Pointer to **interface{}** | The order&#39;s adjustment taxable amount, float. | [optional] 
**FormattedAdjustmentTaxableAmount** | Pointer to **interface{}** | The order&#39;s adjustment taxable amount, formatted. | [optional] 
**TotalAmountWithTaxesCents** | Pointer to **interface{}** | The order&#39;s total amount (when prices include taxes) or the order&#39;s total + taxes amount (when prices don&#39;t include taxes, e.g. US Markets or B2B). | [optional] 
**TotalAmountWithTaxesFloat** | Pointer to **interface{}** | The order&#39;s total amount with taxes, float. | [optional] 
**FormattedTotalAmountWithTaxes** | Pointer to **interface{}** | The order&#39;s total amount with taxes, formatted. | [optional] 
**FeesAmountCents** | Pointer to **interface{}** | The fees amount that is applied by Commerce Layer, in cents. | [optional] 
**FeesAmountFloat** | Pointer to **interface{}** | The fees amount that is applied by Commerce Layer, float. | [optional] 
**FormattedFeesAmount** | Pointer to **interface{}** | The fees amount that is applied by Commerce Layer, formatted. | [optional] 
**DutyAmountCents** | Pointer to **interface{}** | The duty amount that is calculated by external services, in cents. | [optional] 
**DutyAmountFloat** | Pointer to **interface{}** | The duty amount that is calculated by external services, float. | [optional] 
**FormattedDutyAmount** | Pointer to **interface{}** | The duty amount that is calculated by external services, formatted. | [optional] 
**SkusCount** | Pointer to **interface{}** | The total number of SKUs in the order&#39;s line items. This can be useful to display a preview of the customer shopping cart content. | [optional] 
**LineItemOptionsCount** | Pointer to **interface{}** | The total number of line item options. This can be useful to display a preview of the customer shopping cart content. | [optional] 
**ShipmentsCount** | Pointer to **interface{}** | The total number of shipments. This can be useful to manage the shipping method(s) selection during checkout. | [optional] 
**TaxCalculationsCount** | Pointer to **interface{}** | The total number of tax calculations. This can be useful to monitor external tax service usage. | [optional] 
**PaymentSourceDetails** | Pointer to **interface{}** | An object that contains the shareable details of the order&#39;s payment source. | [optional] 
**Token** | Pointer to **interface{}** | A unique token that can be shared more securely instead of the order&#39;s id. | [optional] 
**CartUrl** | Pointer to **interface{}** | The cart url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**ReturnUrl** | Pointer to **interface{}** | The return url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**TermsUrl** | Pointer to **interface{}** | The terms and conditions url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**PrivacyUrl** | Pointer to **interface{}** | The privacy policy url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**CheckoutUrl** | Pointer to **interface{}** | The checkout url that was automatically generated for the order. Send the customers to this url to let them checkout the order securely on our hosted checkout application. | [optional] 
**PlacedAt** | Pointer to **interface{}** | Time at which the order was placed. | [optional] 
**ApprovedAt** | Pointer to **interface{}** | Time at which the order was approved. | [optional] 
**CancelledAt** | Pointer to **interface{}** | Time at which the order was cancelled. | [optional] 
**PaymentUpdatedAt** | Pointer to **interface{}** | Time at which the order&#39;s payment status was last updated. | [optional] 
**FulfillmentUpdatedAt** | Pointer to **interface{}** | Time at which the order&#39;s fulfillment status was last updated. | [optional] 
**RefreshedAt** | Pointer to **interface{}** | Last time at which an order was manually refreshed. | [optional] 
**ArchivedAt** | Pointer to **interface{}** | Time at which the resource has been archived. | [optional] 
**ExpiresAt** | Pointer to **interface{}** | Time at which an order is marked for cleanup. Any order will start with a default expire time of 2 months. Expiration is reset once a line item is added to the order. | [optional] 
**SubscriptionCreatedAt** | Pointer to **interface{}** | Time at which the order has been marked to create a subscription from its recurring line items. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *GETOrders200ResponseDataInnerAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETOrders200ResponseDataInnerAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETOrders200ResponseDataInnerAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetAutorefresh

`func (o *GETOrders200ResponseDataInnerAttributes) GetAutorefresh() interface{}`

GetAutorefresh returns the Autorefresh field if non-nil, zero value otherwise.

### GetAutorefreshOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetAutorefreshOk() (*interface{}, bool)`

GetAutorefreshOk returns a tuple with the Autorefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutorefresh

`func (o *GETOrders200ResponseDataInnerAttributes) SetAutorefresh(v interface{})`

SetAutorefresh sets Autorefresh field to given value.

### HasAutorefresh

`func (o *GETOrders200ResponseDataInnerAttributes) HasAutorefresh() bool`

HasAutorefresh returns a boolean if a field has been set.

### SetAutorefreshNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetAutorefreshNil(b bool)`

 SetAutorefreshNil sets the value for Autorefresh to be an explicit nil

### UnsetAutorefresh
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetAutorefresh()`

UnsetAutorefresh ensures that no value is present for Autorefresh, not even an explicit nil
### GetStatus

`func (o *GETOrders200ResponseDataInnerAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETOrders200ResponseDataInnerAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETOrders200ResponseDataInnerAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetPaymentStatus

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentStatus() interface{}`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentStatusOk() (*interface{}, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentStatus(v interface{})`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### SetPaymentStatusNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentStatusNil(b bool)`

 SetPaymentStatusNil sets the value for PaymentStatus to be an explicit nil

### UnsetPaymentStatus
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetPaymentStatus()`

UnsetPaymentStatus ensures that no value is present for PaymentStatus, not even an explicit nil
### GetFulfillmentStatus

`func (o *GETOrders200ResponseDataInnerAttributes) GetFulfillmentStatus() interface{}`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFulfillmentStatusOk() (*interface{}, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *GETOrders200ResponseDataInnerAttributes) SetFulfillmentStatus(v interface{})`

SetFulfillmentStatus sets FulfillmentStatus field to given value.

### HasFulfillmentStatus

`func (o *GETOrders200ResponseDataInnerAttributes) HasFulfillmentStatus() bool`

HasFulfillmentStatus returns a boolean if a field has been set.

### SetFulfillmentStatusNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFulfillmentStatusNil(b bool)`

 SetFulfillmentStatusNil sets the value for FulfillmentStatus to be an explicit nil

### UnsetFulfillmentStatus
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFulfillmentStatus()`

UnsetFulfillmentStatus ensures that no value is present for FulfillmentStatus, not even an explicit nil
### GetGuest

`func (o *GETOrders200ResponseDataInnerAttributes) GetGuest() interface{}`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetGuestOk() (*interface{}, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *GETOrders200ResponseDataInnerAttributes) SetGuest(v interface{})`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *GETOrders200ResponseDataInnerAttributes) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### SetGuestNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetGuestNil(b bool)`

 SetGuestNil sets the value for Guest to be an explicit nil

### UnsetGuest
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetGuest()`

UnsetGuest ensures that no value is present for Guest, not even an explicit nil
### GetEditable

`func (o *GETOrders200ResponseDataInnerAttributes) GetEditable() interface{}`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetEditableOk() (*interface{}, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *GETOrders200ResponseDataInnerAttributes) SetEditable(v interface{})`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *GETOrders200ResponseDataInnerAttributes) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### SetEditableNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetEditableNil(b bool)`

 SetEditableNil sets the value for Editable to be an explicit nil

### UnsetEditable
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetEditable()`

UnsetEditable ensures that no value is present for Editable, not even an explicit nil
### GetCustomerEmail

`func (o *GETOrders200ResponseDataInnerAttributes) GetCustomerEmail() interface{}`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetCustomerEmailOk() (*interface{}, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *GETOrders200ResponseDataInnerAttributes) SetCustomerEmail(v interface{})`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *GETOrders200ResponseDataInnerAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### SetCustomerEmailNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetCustomerEmailNil(b bool)`

 SetCustomerEmailNil sets the value for CustomerEmail to be an explicit nil

### UnsetCustomerEmail
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetCustomerEmail()`

UnsetCustomerEmail ensures that no value is present for CustomerEmail, not even an explicit nil
### GetLanguageCode

`func (o *GETOrders200ResponseDataInnerAttributes) GetLanguageCode() interface{}`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetLanguageCodeOk() (*interface{}, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *GETOrders200ResponseDataInnerAttributes) SetLanguageCode(v interface{})`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *GETOrders200ResponseDataInnerAttributes) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### SetLanguageCodeNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetLanguageCodeNil(b bool)`

 SetLanguageCodeNil sets the value for LanguageCode to be an explicit nil

### UnsetLanguageCode
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetLanguageCode()`

UnsetLanguageCode ensures that no value is present for LanguageCode, not even an explicit nil
### GetCurrencyCode

`func (o *GETOrders200ResponseDataInnerAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETOrders200ResponseDataInnerAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETOrders200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetTaxIncluded

`func (o *GETOrders200ResponseDataInnerAttributes) GetTaxIncluded() interface{}`

GetTaxIncluded returns the TaxIncluded field if non-nil, zero value otherwise.

### GetTaxIncludedOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTaxIncludedOk() (*interface{}, bool)`

GetTaxIncludedOk returns a tuple with the TaxIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncluded

`func (o *GETOrders200ResponseDataInnerAttributes) SetTaxIncluded(v interface{})`

SetTaxIncluded sets TaxIncluded field to given value.

### HasTaxIncluded

`func (o *GETOrders200ResponseDataInnerAttributes) HasTaxIncluded() bool`

HasTaxIncluded returns a boolean if a field has been set.

### SetTaxIncludedNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetTaxIncludedNil(b bool)`

 SetTaxIncludedNil sets the value for TaxIncluded to be an explicit nil

### UnsetTaxIncluded
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetTaxIncluded()`

UnsetTaxIncluded ensures that no value is present for TaxIncluded, not even an explicit nil
### GetTaxRate

`func (o *GETOrders200ResponseDataInnerAttributes) GetTaxRate() interface{}`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTaxRateOk() (*interface{}, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *GETOrders200ResponseDataInnerAttributes) SetTaxRate(v interface{})`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *GETOrders200ResponseDataInnerAttributes) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetFreightTaxable

`func (o *GETOrders200ResponseDataInnerAttributes) GetFreightTaxable() interface{}`

GetFreightTaxable returns the FreightTaxable field if non-nil, zero value otherwise.

### GetFreightTaxableOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFreightTaxableOk() (*interface{}, bool)`

GetFreightTaxableOk returns a tuple with the FreightTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTaxable

`func (o *GETOrders200ResponseDataInnerAttributes) SetFreightTaxable(v interface{})`

SetFreightTaxable sets FreightTaxable field to given value.

### HasFreightTaxable

`func (o *GETOrders200ResponseDataInnerAttributes) HasFreightTaxable() bool`

HasFreightTaxable returns a boolean if a field has been set.

### SetFreightTaxableNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFreightTaxableNil(b bool)`

 SetFreightTaxableNil sets the value for FreightTaxable to be an explicit nil

### UnsetFreightTaxable
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFreightTaxable()`

UnsetFreightTaxable ensures that no value is present for FreightTaxable, not even an explicit nil
### GetRequiresBillingInfo

`func (o *GETOrders200ResponseDataInnerAttributes) GetRequiresBillingInfo() interface{}`

GetRequiresBillingInfo returns the RequiresBillingInfo field if non-nil, zero value otherwise.

### GetRequiresBillingInfoOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetRequiresBillingInfoOk() (*interface{}, bool)`

GetRequiresBillingInfoOk returns a tuple with the RequiresBillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresBillingInfo

`func (o *GETOrders200ResponseDataInnerAttributes) SetRequiresBillingInfo(v interface{})`

SetRequiresBillingInfo sets RequiresBillingInfo field to given value.

### HasRequiresBillingInfo

`func (o *GETOrders200ResponseDataInnerAttributes) HasRequiresBillingInfo() bool`

HasRequiresBillingInfo returns a boolean if a field has been set.

### SetRequiresBillingInfoNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetRequiresBillingInfoNil(b bool)`

 SetRequiresBillingInfoNil sets the value for RequiresBillingInfo to be an explicit nil

### UnsetRequiresBillingInfo
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetRequiresBillingInfo()`

UnsetRequiresBillingInfo ensures that no value is present for RequiresBillingInfo, not even an explicit nil
### GetCountryCode

`func (o *GETOrders200ResponseDataInnerAttributes) GetCountryCode() interface{}`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetCountryCodeOk() (*interface{}, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *GETOrders200ResponseDataInnerAttributes) SetCountryCode(v interface{})`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *GETOrders200ResponseDataInnerAttributes) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetShippingCountryCodeLock

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingCountryCodeLock() interface{}`

GetShippingCountryCodeLock returns the ShippingCountryCodeLock field if non-nil, zero value otherwise.

### GetShippingCountryCodeLockOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingCountryCodeLockOk() (*interface{}, bool)`

GetShippingCountryCodeLockOk returns a tuple with the ShippingCountryCodeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCountryCodeLock

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingCountryCodeLock(v interface{})`

SetShippingCountryCodeLock sets ShippingCountryCodeLock field to given value.

### HasShippingCountryCodeLock

`func (o *GETOrders200ResponseDataInnerAttributes) HasShippingCountryCodeLock() bool`

HasShippingCountryCodeLock returns a boolean if a field has been set.

### SetShippingCountryCodeLockNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingCountryCodeLockNil(b bool)`

 SetShippingCountryCodeLockNil sets the value for ShippingCountryCodeLock to be an explicit nil

### UnsetShippingCountryCodeLock
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetShippingCountryCodeLock()`

UnsetShippingCountryCodeLock ensures that no value is present for ShippingCountryCodeLock, not even an explicit nil
### GetCouponCode

`func (o *GETOrders200ResponseDataInnerAttributes) GetCouponCode() interface{}`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetCouponCodeOk() (*interface{}, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *GETOrders200ResponseDataInnerAttributes) SetCouponCode(v interface{})`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *GETOrders200ResponseDataInnerAttributes) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.

### SetCouponCodeNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetCouponCodeNil(b bool)`

 SetCouponCodeNil sets the value for CouponCode to be an explicit nil

### UnsetCouponCode
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetCouponCode()`

UnsetCouponCode ensures that no value is present for CouponCode, not even an explicit nil
### GetGiftCardCode

`func (o *GETOrders200ResponseDataInnerAttributes) GetGiftCardCode() interface{}`

GetGiftCardCode returns the GiftCardCode field if non-nil, zero value otherwise.

### GetGiftCardCodeOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetGiftCardCodeOk() (*interface{}, bool)`

GetGiftCardCodeOk returns a tuple with the GiftCardCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardCode

`func (o *GETOrders200ResponseDataInnerAttributes) SetGiftCardCode(v interface{})`

SetGiftCardCode sets GiftCardCode field to given value.

### HasGiftCardCode

`func (o *GETOrders200ResponseDataInnerAttributes) HasGiftCardCode() bool`

HasGiftCardCode returns a boolean if a field has been set.

### SetGiftCardCodeNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetGiftCardCodeNil(b bool)`

 SetGiftCardCodeNil sets the value for GiftCardCode to be an explicit nil

### UnsetGiftCardCode
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetGiftCardCode()`

UnsetGiftCardCode ensures that no value is present for GiftCardCode, not even an explicit nil
### GetGiftCardOrCouponCode

`func (o *GETOrders200ResponseDataInnerAttributes) GetGiftCardOrCouponCode() interface{}`

GetGiftCardOrCouponCode returns the GiftCardOrCouponCode field if non-nil, zero value otherwise.

### GetGiftCardOrCouponCodeOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetGiftCardOrCouponCodeOk() (*interface{}, bool)`

GetGiftCardOrCouponCodeOk returns a tuple with the GiftCardOrCouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardOrCouponCode

`func (o *GETOrders200ResponseDataInnerAttributes) SetGiftCardOrCouponCode(v interface{})`

SetGiftCardOrCouponCode sets GiftCardOrCouponCode field to given value.

### HasGiftCardOrCouponCode

`func (o *GETOrders200ResponseDataInnerAttributes) HasGiftCardOrCouponCode() bool`

HasGiftCardOrCouponCode returns a boolean if a field has been set.

### SetGiftCardOrCouponCodeNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetGiftCardOrCouponCodeNil(b bool)`

 SetGiftCardOrCouponCodeNil sets the value for GiftCardOrCouponCode to be an explicit nil

### UnsetGiftCardOrCouponCode
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetGiftCardOrCouponCode()`

UnsetGiftCardOrCouponCode ensures that no value is present for GiftCardOrCouponCode, not even an explicit nil
### GetSubtotalAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalAmountCents() interface{}`

GetSubtotalAmountCents returns the SubtotalAmountCents field if non-nil, zero value otherwise.

### GetSubtotalAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalAmountCentsOk() (*interface{}, bool)`

GetSubtotalAmountCentsOk returns a tuple with the SubtotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalAmountCents(v interface{})`

SetSubtotalAmountCents sets SubtotalAmountCents field to given value.

### HasSubtotalAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasSubtotalAmountCents() bool`

HasSubtotalAmountCents returns a boolean if a field has been set.

### SetSubtotalAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalAmountCentsNil(b bool)`

 SetSubtotalAmountCentsNil sets the value for SubtotalAmountCents to be an explicit nil

### UnsetSubtotalAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetSubtotalAmountCents()`

UnsetSubtotalAmountCents ensures that no value is present for SubtotalAmountCents, not even an explicit nil
### GetSubtotalAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalAmountFloat() interface{}`

GetSubtotalAmountFloat returns the SubtotalAmountFloat field if non-nil, zero value otherwise.

### GetSubtotalAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalAmountFloatOk() (*interface{}, bool)`

GetSubtotalAmountFloatOk returns a tuple with the SubtotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalAmountFloat(v interface{})`

SetSubtotalAmountFloat sets SubtotalAmountFloat field to given value.

### HasSubtotalAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasSubtotalAmountFloat() bool`

HasSubtotalAmountFloat returns a boolean if a field has been set.

### SetSubtotalAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalAmountFloatNil(b bool)`

 SetSubtotalAmountFloatNil sets the value for SubtotalAmountFloat to be an explicit nil

### UnsetSubtotalAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetSubtotalAmountFloat()`

UnsetSubtotalAmountFloat ensures that no value is present for SubtotalAmountFloat, not even an explicit nil
### GetFormattedSubtotalAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedSubtotalAmount() interface{}`

GetFormattedSubtotalAmount returns the FormattedSubtotalAmount field if non-nil, zero value otherwise.

### GetFormattedSubtotalAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedSubtotalAmountOk() (*interface{}, bool)`

GetFormattedSubtotalAmountOk returns a tuple with the FormattedSubtotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSubtotalAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedSubtotalAmount(v interface{})`

SetFormattedSubtotalAmount sets FormattedSubtotalAmount field to given value.

### HasFormattedSubtotalAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedSubtotalAmount() bool`

HasFormattedSubtotalAmount returns a boolean if a field has been set.

### SetFormattedSubtotalAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedSubtotalAmountNil(b bool)`

 SetFormattedSubtotalAmountNil sets the value for FormattedSubtotalAmount to be an explicit nil

### UnsetFormattedSubtotalAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedSubtotalAmount()`

UnsetFormattedSubtotalAmount ensures that no value is present for FormattedSubtotalAmount, not even an explicit nil
### GetShippingAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingAmountCents() interface{}`

GetShippingAmountCents returns the ShippingAmountCents field if non-nil, zero value otherwise.

### GetShippingAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingAmountCentsOk() (*interface{}, bool)`

GetShippingAmountCentsOk returns a tuple with the ShippingAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingAmountCents(v interface{})`

SetShippingAmountCents sets ShippingAmountCents field to given value.

### HasShippingAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasShippingAmountCents() bool`

HasShippingAmountCents returns a boolean if a field has been set.

### SetShippingAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingAmountCentsNil(b bool)`

 SetShippingAmountCentsNil sets the value for ShippingAmountCents to be an explicit nil

### UnsetShippingAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetShippingAmountCents()`

UnsetShippingAmountCents ensures that no value is present for ShippingAmountCents, not even an explicit nil
### GetShippingAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingAmountFloat() interface{}`

GetShippingAmountFloat returns the ShippingAmountFloat field if non-nil, zero value otherwise.

### GetShippingAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingAmountFloatOk() (*interface{}, bool)`

GetShippingAmountFloatOk returns a tuple with the ShippingAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingAmountFloat(v interface{})`

SetShippingAmountFloat sets ShippingAmountFloat field to given value.

### HasShippingAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasShippingAmountFloat() bool`

HasShippingAmountFloat returns a boolean if a field has been set.

### SetShippingAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingAmountFloatNil(b bool)`

 SetShippingAmountFloatNil sets the value for ShippingAmountFloat to be an explicit nil

### UnsetShippingAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetShippingAmountFloat()`

UnsetShippingAmountFloat ensures that no value is present for ShippingAmountFloat, not even an explicit nil
### GetFormattedShippingAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedShippingAmount() interface{}`

GetFormattedShippingAmount returns the FormattedShippingAmount field if non-nil, zero value otherwise.

### GetFormattedShippingAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedShippingAmountOk() (*interface{}, bool)`

GetFormattedShippingAmountOk returns a tuple with the FormattedShippingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedShippingAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedShippingAmount(v interface{})`

SetFormattedShippingAmount sets FormattedShippingAmount field to given value.

### HasFormattedShippingAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedShippingAmount() bool`

HasFormattedShippingAmount returns a boolean if a field has been set.

### SetFormattedShippingAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedShippingAmountNil(b bool)`

 SetFormattedShippingAmountNil sets the value for FormattedShippingAmount to be an explicit nil

### UnsetFormattedShippingAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedShippingAmount()`

UnsetFormattedShippingAmount ensures that no value is present for FormattedShippingAmount, not even an explicit nil
### GetPaymentMethodAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodAmountCents() interface{}`

GetPaymentMethodAmountCents returns the PaymentMethodAmountCents field if non-nil, zero value otherwise.

### GetPaymentMethodAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodAmountCentsOk() (*interface{}, bool)`

GetPaymentMethodAmountCentsOk returns a tuple with the PaymentMethodAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodAmountCents(v interface{})`

SetPaymentMethodAmountCents sets PaymentMethodAmountCents field to given value.

### HasPaymentMethodAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentMethodAmountCents() bool`

HasPaymentMethodAmountCents returns a boolean if a field has been set.

### SetPaymentMethodAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodAmountCentsNil(b bool)`

 SetPaymentMethodAmountCentsNil sets the value for PaymentMethodAmountCents to be an explicit nil

### UnsetPaymentMethodAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetPaymentMethodAmountCents()`

UnsetPaymentMethodAmountCents ensures that no value is present for PaymentMethodAmountCents, not even an explicit nil
### GetPaymentMethodAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodAmountFloat() interface{}`

GetPaymentMethodAmountFloat returns the PaymentMethodAmountFloat field if non-nil, zero value otherwise.

### GetPaymentMethodAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodAmountFloatOk() (*interface{}, bool)`

GetPaymentMethodAmountFloatOk returns a tuple with the PaymentMethodAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodAmountFloat(v interface{})`

SetPaymentMethodAmountFloat sets PaymentMethodAmountFloat field to given value.

### HasPaymentMethodAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentMethodAmountFloat() bool`

HasPaymentMethodAmountFloat returns a boolean if a field has been set.

### SetPaymentMethodAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodAmountFloatNil(b bool)`

 SetPaymentMethodAmountFloatNil sets the value for PaymentMethodAmountFloat to be an explicit nil

### UnsetPaymentMethodAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetPaymentMethodAmountFloat()`

UnsetPaymentMethodAmountFloat ensures that no value is present for PaymentMethodAmountFloat, not even an explicit nil
### GetFormattedPaymentMethodAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedPaymentMethodAmount() interface{}`

GetFormattedPaymentMethodAmount returns the FormattedPaymentMethodAmount field if non-nil, zero value otherwise.

### GetFormattedPaymentMethodAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedPaymentMethodAmountOk() (*interface{}, bool)`

GetFormattedPaymentMethodAmountOk returns a tuple with the FormattedPaymentMethodAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPaymentMethodAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedPaymentMethodAmount(v interface{})`

SetFormattedPaymentMethodAmount sets FormattedPaymentMethodAmount field to given value.

### HasFormattedPaymentMethodAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedPaymentMethodAmount() bool`

HasFormattedPaymentMethodAmount returns a boolean if a field has been set.

### SetFormattedPaymentMethodAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedPaymentMethodAmountNil(b bool)`

 SetFormattedPaymentMethodAmountNil sets the value for FormattedPaymentMethodAmount to be an explicit nil

### UnsetFormattedPaymentMethodAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedPaymentMethodAmount()`

UnsetFormattedPaymentMethodAmount ensures that no value is present for FormattedPaymentMethodAmount, not even an explicit nil
### GetDiscountAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetDiscountAmountCents() interface{}`

GetDiscountAmountCents returns the DiscountAmountCents field if non-nil, zero value otherwise.

### GetDiscountAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetDiscountAmountCentsOk() (*interface{}, bool)`

GetDiscountAmountCentsOk returns a tuple with the DiscountAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetDiscountAmountCents(v interface{})`

SetDiscountAmountCents sets DiscountAmountCents field to given value.

### HasDiscountAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasDiscountAmountCents() bool`

HasDiscountAmountCents returns a boolean if a field has been set.

### SetDiscountAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetDiscountAmountCentsNil(b bool)`

 SetDiscountAmountCentsNil sets the value for DiscountAmountCents to be an explicit nil

### UnsetDiscountAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetDiscountAmountCents()`

UnsetDiscountAmountCents ensures that no value is present for DiscountAmountCents, not even an explicit nil
### GetDiscountAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetDiscountAmountFloat() interface{}`

GetDiscountAmountFloat returns the DiscountAmountFloat field if non-nil, zero value otherwise.

### GetDiscountAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetDiscountAmountFloatOk() (*interface{}, bool)`

GetDiscountAmountFloatOk returns a tuple with the DiscountAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetDiscountAmountFloat(v interface{})`

SetDiscountAmountFloat sets DiscountAmountFloat field to given value.

### HasDiscountAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasDiscountAmountFloat() bool`

HasDiscountAmountFloat returns a boolean if a field has been set.

### SetDiscountAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetDiscountAmountFloatNil(b bool)`

 SetDiscountAmountFloatNil sets the value for DiscountAmountFloat to be an explicit nil

### UnsetDiscountAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetDiscountAmountFloat()`

UnsetDiscountAmountFloat ensures that no value is present for DiscountAmountFloat, not even an explicit nil
### GetFormattedDiscountAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedDiscountAmount() interface{}`

GetFormattedDiscountAmount returns the FormattedDiscountAmount field if non-nil, zero value otherwise.

### GetFormattedDiscountAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedDiscountAmountOk() (*interface{}, bool)`

GetFormattedDiscountAmountOk returns a tuple with the FormattedDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedDiscountAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedDiscountAmount(v interface{})`

SetFormattedDiscountAmount sets FormattedDiscountAmount field to given value.

### HasFormattedDiscountAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedDiscountAmount() bool`

HasFormattedDiscountAmount returns a boolean if a field has been set.

### SetFormattedDiscountAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedDiscountAmountNil(b bool)`

 SetFormattedDiscountAmountNil sets the value for FormattedDiscountAmount to be an explicit nil

### UnsetFormattedDiscountAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedDiscountAmount()`

UnsetFormattedDiscountAmount ensures that no value is present for FormattedDiscountAmount, not even an explicit nil
### GetAdjustmentAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentAmountCents() interface{}`

GetAdjustmentAmountCents returns the AdjustmentAmountCents field if non-nil, zero value otherwise.

### GetAdjustmentAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentAmountCentsOk() (*interface{}, bool)`

GetAdjustmentAmountCentsOk returns a tuple with the AdjustmentAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentAmountCents(v interface{})`

SetAdjustmentAmountCents sets AdjustmentAmountCents field to given value.

### HasAdjustmentAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasAdjustmentAmountCents() bool`

HasAdjustmentAmountCents returns a boolean if a field has been set.

### SetAdjustmentAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentAmountCentsNil(b bool)`

 SetAdjustmentAmountCentsNil sets the value for AdjustmentAmountCents to be an explicit nil

### UnsetAdjustmentAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetAdjustmentAmountCents()`

UnsetAdjustmentAmountCents ensures that no value is present for AdjustmentAmountCents, not even an explicit nil
### GetAdjustmentAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentAmountFloat() interface{}`

GetAdjustmentAmountFloat returns the AdjustmentAmountFloat field if non-nil, zero value otherwise.

### GetAdjustmentAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentAmountFloatOk() (*interface{}, bool)`

GetAdjustmentAmountFloatOk returns a tuple with the AdjustmentAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentAmountFloat(v interface{})`

SetAdjustmentAmountFloat sets AdjustmentAmountFloat field to given value.

### HasAdjustmentAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasAdjustmentAmountFloat() bool`

HasAdjustmentAmountFloat returns a boolean if a field has been set.

### SetAdjustmentAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentAmountFloatNil(b bool)`

 SetAdjustmentAmountFloatNil sets the value for AdjustmentAmountFloat to be an explicit nil

### UnsetAdjustmentAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetAdjustmentAmountFloat()`

UnsetAdjustmentAmountFloat ensures that no value is present for AdjustmentAmountFloat, not even an explicit nil
### GetFormattedAdjustmentAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedAdjustmentAmount() interface{}`

GetFormattedAdjustmentAmount returns the FormattedAdjustmentAmount field if non-nil, zero value otherwise.

### GetFormattedAdjustmentAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedAdjustmentAmountOk() (*interface{}, bool)`

GetFormattedAdjustmentAmountOk returns a tuple with the FormattedAdjustmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAdjustmentAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedAdjustmentAmount(v interface{})`

SetFormattedAdjustmentAmount sets FormattedAdjustmentAmount field to given value.

### HasFormattedAdjustmentAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedAdjustmentAmount() bool`

HasFormattedAdjustmentAmount returns a boolean if a field has been set.

### SetFormattedAdjustmentAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedAdjustmentAmountNil(b bool)`

 SetFormattedAdjustmentAmountNil sets the value for FormattedAdjustmentAmount to be an explicit nil

### UnsetFormattedAdjustmentAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedAdjustmentAmount()`

UnsetFormattedAdjustmentAmount ensures that no value is present for FormattedAdjustmentAmount, not even an explicit nil
### GetGiftCardAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetGiftCardAmountCents() interface{}`

GetGiftCardAmountCents returns the GiftCardAmountCents field if non-nil, zero value otherwise.

### GetGiftCardAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetGiftCardAmountCentsOk() (*interface{}, bool)`

GetGiftCardAmountCentsOk returns a tuple with the GiftCardAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetGiftCardAmountCents(v interface{})`

SetGiftCardAmountCents sets GiftCardAmountCents field to given value.

### HasGiftCardAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasGiftCardAmountCents() bool`

HasGiftCardAmountCents returns a boolean if a field has been set.

### SetGiftCardAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetGiftCardAmountCentsNil(b bool)`

 SetGiftCardAmountCentsNil sets the value for GiftCardAmountCents to be an explicit nil

### UnsetGiftCardAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetGiftCardAmountCents()`

UnsetGiftCardAmountCents ensures that no value is present for GiftCardAmountCents, not even an explicit nil
### GetGiftCardAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetGiftCardAmountFloat() interface{}`

GetGiftCardAmountFloat returns the GiftCardAmountFloat field if non-nil, zero value otherwise.

### GetGiftCardAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetGiftCardAmountFloatOk() (*interface{}, bool)`

GetGiftCardAmountFloatOk returns a tuple with the GiftCardAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetGiftCardAmountFloat(v interface{})`

SetGiftCardAmountFloat sets GiftCardAmountFloat field to given value.

### HasGiftCardAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasGiftCardAmountFloat() bool`

HasGiftCardAmountFloat returns a boolean if a field has been set.

### SetGiftCardAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetGiftCardAmountFloatNil(b bool)`

 SetGiftCardAmountFloatNil sets the value for GiftCardAmountFloat to be an explicit nil

### UnsetGiftCardAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetGiftCardAmountFloat()`

UnsetGiftCardAmountFloat ensures that no value is present for GiftCardAmountFloat, not even an explicit nil
### GetFormattedGiftCardAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedGiftCardAmount() interface{}`

GetFormattedGiftCardAmount returns the FormattedGiftCardAmount field if non-nil, zero value otherwise.

### GetFormattedGiftCardAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedGiftCardAmountOk() (*interface{}, bool)`

GetFormattedGiftCardAmountOk returns a tuple with the FormattedGiftCardAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedGiftCardAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedGiftCardAmount(v interface{})`

SetFormattedGiftCardAmount sets FormattedGiftCardAmount field to given value.

### HasFormattedGiftCardAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedGiftCardAmount() bool`

HasFormattedGiftCardAmount returns a boolean if a field has been set.

### SetFormattedGiftCardAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedGiftCardAmountNil(b bool)`

 SetFormattedGiftCardAmountNil sets the value for FormattedGiftCardAmount to be an explicit nil

### UnsetFormattedGiftCardAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedGiftCardAmount()`

UnsetFormattedGiftCardAmount ensures that no value is present for FormattedGiftCardAmount, not even an explicit nil
### GetTotalTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalTaxAmountCents() interface{}`

GetTotalTaxAmountCents returns the TotalTaxAmountCents field if non-nil, zero value otherwise.

### GetTotalTaxAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalTaxAmountCentsOk() (*interface{}, bool)`

GetTotalTaxAmountCentsOk returns a tuple with the TotalTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalTaxAmountCents(v interface{})`

SetTotalTaxAmountCents sets TotalTaxAmountCents field to given value.

### HasTotalTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasTotalTaxAmountCents() bool`

HasTotalTaxAmountCents returns a boolean if a field has been set.

### SetTotalTaxAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalTaxAmountCentsNil(b bool)`

 SetTotalTaxAmountCentsNil sets the value for TotalTaxAmountCents to be an explicit nil

### UnsetTotalTaxAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetTotalTaxAmountCents()`

UnsetTotalTaxAmountCents ensures that no value is present for TotalTaxAmountCents, not even an explicit nil
### GetTotalTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalTaxAmountFloat() interface{}`

GetTotalTaxAmountFloat returns the TotalTaxAmountFloat field if non-nil, zero value otherwise.

### GetTotalTaxAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalTaxAmountFloatOk() (*interface{}, bool)`

GetTotalTaxAmountFloatOk returns a tuple with the TotalTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalTaxAmountFloat(v interface{})`

SetTotalTaxAmountFloat sets TotalTaxAmountFloat field to given value.

### HasTotalTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasTotalTaxAmountFloat() bool`

HasTotalTaxAmountFloat returns a boolean if a field has been set.

### SetTotalTaxAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalTaxAmountFloatNil(b bool)`

 SetTotalTaxAmountFloatNil sets the value for TotalTaxAmountFloat to be an explicit nil

### UnsetTotalTaxAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetTotalTaxAmountFloat()`

UnsetTotalTaxAmountFloat ensures that no value is present for TotalTaxAmountFloat, not even an explicit nil
### GetFormattedTotalTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedTotalTaxAmount() interface{}`

GetFormattedTotalTaxAmount returns the FormattedTotalTaxAmount field if non-nil, zero value otherwise.

### GetFormattedTotalTaxAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedTotalTaxAmountOk() (*interface{}, bool)`

GetFormattedTotalTaxAmountOk returns a tuple with the FormattedTotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedTotalTaxAmount(v interface{})`

SetFormattedTotalTaxAmount sets FormattedTotalTaxAmount field to given value.

### HasFormattedTotalTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedTotalTaxAmount() bool`

HasFormattedTotalTaxAmount returns a boolean if a field has been set.

### SetFormattedTotalTaxAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedTotalTaxAmountNil(b bool)`

 SetFormattedTotalTaxAmountNil sets the value for FormattedTotalTaxAmount to be an explicit nil

### UnsetFormattedTotalTaxAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedTotalTaxAmount()`

UnsetFormattedTotalTaxAmount ensures that no value is present for FormattedTotalTaxAmount, not even an explicit nil
### GetSubtotalTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalTaxAmountCents() interface{}`

GetSubtotalTaxAmountCents returns the SubtotalTaxAmountCents field if non-nil, zero value otherwise.

### GetSubtotalTaxAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalTaxAmountCentsOk() (*interface{}, bool)`

GetSubtotalTaxAmountCentsOk returns a tuple with the SubtotalTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalTaxAmountCents(v interface{})`

SetSubtotalTaxAmountCents sets SubtotalTaxAmountCents field to given value.

### HasSubtotalTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasSubtotalTaxAmountCents() bool`

HasSubtotalTaxAmountCents returns a boolean if a field has been set.

### SetSubtotalTaxAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalTaxAmountCentsNil(b bool)`

 SetSubtotalTaxAmountCentsNil sets the value for SubtotalTaxAmountCents to be an explicit nil

### UnsetSubtotalTaxAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetSubtotalTaxAmountCents()`

UnsetSubtotalTaxAmountCents ensures that no value is present for SubtotalTaxAmountCents, not even an explicit nil
### GetSubtotalTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalTaxAmountFloat() interface{}`

GetSubtotalTaxAmountFloat returns the SubtotalTaxAmountFloat field if non-nil, zero value otherwise.

### GetSubtotalTaxAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalTaxAmountFloatOk() (*interface{}, bool)`

GetSubtotalTaxAmountFloatOk returns a tuple with the SubtotalTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalTaxAmountFloat(v interface{})`

SetSubtotalTaxAmountFloat sets SubtotalTaxAmountFloat field to given value.

### HasSubtotalTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasSubtotalTaxAmountFloat() bool`

HasSubtotalTaxAmountFloat returns a boolean if a field has been set.

### SetSubtotalTaxAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalTaxAmountFloatNil(b bool)`

 SetSubtotalTaxAmountFloatNil sets the value for SubtotalTaxAmountFloat to be an explicit nil

### UnsetSubtotalTaxAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetSubtotalTaxAmountFloat()`

UnsetSubtotalTaxAmountFloat ensures that no value is present for SubtotalTaxAmountFloat, not even an explicit nil
### GetFormattedSubtotalTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedSubtotalTaxAmount() interface{}`

GetFormattedSubtotalTaxAmount returns the FormattedSubtotalTaxAmount field if non-nil, zero value otherwise.

### GetFormattedSubtotalTaxAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedSubtotalTaxAmountOk() (*interface{}, bool)`

GetFormattedSubtotalTaxAmountOk returns a tuple with the FormattedSubtotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSubtotalTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedSubtotalTaxAmount(v interface{})`

SetFormattedSubtotalTaxAmount sets FormattedSubtotalTaxAmount field to given value.

### HasFormattedSubtotalTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedSubtotalTaxAmount() bool`

HasFormattedSubtotalTaxAmount returns a boolean if a field has been set.

### SetFormattedSubtotalTaxAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedSubtotalTaxAmountNil(b bool)`

 SetFormattedSubtotalTaxAmountNil sets the value for FormattedSubtotalTaxAmount to be an explicit nil

### UnsetFormattedSubtotalTaxAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedSubtotalTaxAmount()`

UnsetFormattedSubtotalTaxAmount ensures that no value is present for FormattedSubtotalTaxAmount, not even an explicit nil
### GetShippingTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingTaxAmountCents() interface{}`

GetShippingTaxAmountCents returns the ShippingTaxAmountCents field if non-nil, zero value otherwise.

### GetShippingTaxAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingTaxAmountCentsOk() (*interface{}, bool)`

GetShippingTaxAmountCentsOk returns a tuple with the ShippingTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingTaxAmountCents(v interface{})`

SetShippingTaxAmountCents sets ShippingTaxAmountCents field to given value.

### HasShippingTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasShippingTaxAmountCents() bool`

HasShippingTaxAmountCents returns a boolean if a field has been set.

### SetShippingTaxAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingTaxAmountCentsNil(b bool)`

 SetShippingTaxAmountCentsNil sets the value for ShippingTaxAmountCents to be an explicit nil

### UnsetShippingTaxAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetShippingTaxAmountCents()`

UnsetShippingTaxAmountCents ensures that no value is present for ShippingTaxAmountCents, not even an explicit nil
### GetShippingTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingTaxAmountFloat() interface{}`

GetShippingTaxAmountFloat returns the ShippingTaxAmountFloat field if non-nil, zero value otherwise.

### GetShippingTaxAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingTaxAmountFloatOk() (*interface{}, bool)`

GetShippingTaxAmountFloatOk returns a tuple with the ShippingTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingTaxAmountFloat(v interface{})`

SetShippingTaxAmountFloat sets ShippingTaxAmountFloat field to given value.

### HasShippingTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasShippingTaxAmountFloat() bool`

HasShippingTaxAmountFloat returns a boolean if a field has been set.

### SetShippingTaxAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingTaxAmountFloatNil(b bool)`

 SetShippingTaxAmountFloatNil sets the value for ShippingTaxAmountFloat to be an explicit nil

### UnsetShippingTaxAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetShippingTaxAmountFloat()`

UnsetShippingTaxAmountFloat ensures that no value is present for ShippingTaxAmountFloat, not even an explicit nil
### GetFormattedShippingTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedShippingTaxAmount() interface{}`

GetFormattedShippingTaxAmount returns the FormattedShippingTaxAmount field if non-nil, zero value otherwise.

### GetFormattedShippingTaxAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedShippingTaxAmountOk() (*interface{}, bool)`

GetFormattedShippingTaxAmountOk returns a tuple with the FormattedShippingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedShippingTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedShippingTaxAmount(v interface{})`

SetFormattedShippingTaxAmount sets FormattedShippingTaxAmount field to given value.

### HasFormattedShippingTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedShippingTaxAmount() bool`

HasFormattedShippingTaxAmount returns a boolean if a field has been set.

### SetFormattedShippingTaxAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedShippingTaxAmountNil(b bool)`

 SetFormattedShippingTaxAmountNil sets the value for FormattedShippingTaxAmount to be an explicit nil

### UnsetFormattedShippingTaxAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedShippingTaxAmount()`

UnsetFormattedShippingTaxAmount ensures that no value is present for FormattedShippingTaxAmount, not even an explicit nil
### GetPaymentMethodTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodTaxAmountCents() interface{}`

GetPaymentMethodTaxAmountCents returns the PaymentMethodTaxAmountCents field if non-nil, zero value otherwise.

### GetPaymentMethodTaxAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodTaxAmountCentsOk() (*interface{}, bool)`

GetPaymentMethodTaxAmountCentsOk returns a tuple with the PaymentMethodTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodTaxAmountCents(v interface{})`

SetPaymentMethodTaxAmountCents sets PaymentMethodTaxAmountCents field to given value.

### HasPaymentMethodTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentMethodTaxAmountCents() bool`

HasPaymentMethodTaxAmountCents returns a boolean if a field has been set.

### SetPaymentMethodTaxAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodTaxAmountCentsNil(b bool)`

 SetPaymentMethodTaxAmountCentsNil sets the value for PaymentMethodTaxAmountCents to be an explicit nil

### UnsetPaymentMethodTaxAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetPaymentMethodTaxAmountCents()`

UnsetPaymentMethodTaxAmountCents ensures that no value is present for PaymentMethodTaxAmountCents, not even an explicit nil
### GetPaymentMethodTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodTaxAmountFloat() interface{}`

GetPaymentMethodTaxAmountFloat returns the PaymentMethodTaxAmountFloat field if non-nil, zero value otherwise.

### GetPaymentMethodTaxAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodTaxAmountFloatOk() (*interface{}, bool)`

GetPaymentMethodTaxAmountFloatOk returns a tuple with the PaymentMethodTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodTaxAmountFloat(v interface{})`

SetPaymentMethodTaxAmountFloat sets PaymentMethodTaxAmountFloat field to given value.

### HasPaymentMethodTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentMethodTaxAmountFloat() bool`

HasPaymentMethodTaxAmountFloat returns a boolean if a field has been set.

### SetPaymentMethodTaxAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodTaxAmountFloatNil(b bool)`

 SetPaymentMethodTaxAmountFloatNil sets the value for PaymentMethodTaxAmountFloat to be an explicit nil

### UnsetPaymentMethodTaxAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetPaymentMethodTaxAmountFloat()`

UnsetPaymentMethodTaxAmountFloat ensures that no value is present for PaymentMethodTaxAmountFloat, not even an explicit nil
### GetFormattedPaymentMethodTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedPaymentMethodTaxAmount() interface{}`

GetFormattedPaymentMethodTaxAmount returns the FormattedPaymentMethodTaxAmount field if non-nil, zero value otherwise.

### GetFormattedPaymentMethodTaxAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedPaymentMethodTaxAmountOk() (*interface{}, bool)`

GetFormattedPaymentMethodTaxAmountOk returns a tuple with the FormattedPaymentMethodTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPaymentMethodTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedPaymentMethodTaxAmount(v interface{})`

SetFormattedPaymentMethodTaxAmount sets FormattedPaymentMethodTaxAmount field to given value.

### HasFormattedPaymentMethodTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedPaymentMethodTaxAmount() bool`

HasFormattedPaymentMethodTaxAmount returns a boolean if a field has been set.

### SetFormattedPaymentMethodTaxAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedPaymentMethodTaxAmountNil(b bool)`

 SetFormattedPaymentMethodTaxAmountNil sets the value for FormattedPaymentMethodTaxAmount to be an explicit nil

### UnsetFormattedPaymentMethodTaxAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedPaymentMethodTaxAmount()`

UnsetFormattedPaymentMethodTaxAmount ensures that no value is present for FormattedPaymentMethodTaxAmount, not even an explicit nil
### GetAdjustmentTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentTaxAmountCents() interface{}`

GetAdjustmentTaxAmountCents returns the AdjustmentTaxAmountCents field if non-nil, zero value otherwise.

### GetAdjustmentTaxAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentTaxAmountCentsOk() (*interface{}, bool)`

GetAdjustmentTaxAmountCentsOk returns a tuple with the AdjustmentTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentTaxAmountCents(v interface{})`

SetAdjustmentTaxAmountCents sets AdjustmentTaxAmountCents field to given value.

### HasAdjustmentTaxAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasAdjustmentTaxAmountCents() bool`

HasAdjustmentTaxAmountCents returns a boolean if a field has been set.

### SetAdjustmentTaxAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentTaxAmountCentsNil(b bool)`

 SetAdjustmentTaxAmountCentsNil sets the value for AdjustmentTaxAmountCents to be an explicit nil

### UnsetAdjustmentTaxAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetAdjustmentTaxAmountCents()`

UnsetAdjustmentTaxAmountCents ensures that no value is present for AdjustmentTaxAmountCents, not even an explicit nil
### GetAdjustmentTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentTaxAmountFloat() interface{}`

GetAdjustmentTaxAmountFloat returns the AdjustmentTaxAmountFloat field if non-nil, zero value otherwise.

### GetAdjustmentTaxAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentTaxAmountFloatOk() (*interface{}, bool)`

GetAdjustmentTaxAmountFloatOk returns a tuple with the AdjustmentTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentTaxAmountFloat(v interface{})`

SetAdjustmentTaxAmountFloat sets AdjustmentTaxAmountFloat field to given value.

### HasAdjustmentTaxAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasAdjustmentTaxAmountFloat() bool`

HasAdjustmentTaxAmountFloat returns a boolean if a field has been set.

### SetAdjustmentTaxAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentTaxAmountFloatNil(b bool)`

 SetAdjustmentTaxAmountFloatNil sets the value for AdjustmentTaxAmountFloat to be an explicit nil

### UnsetAdjustmentTaxAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetAdjustmentTaxAmountFloat()`

UnsetAdjustmentTaxAmountFloat ensures that no value is present for AdjustmentTaxAmountFloat, not even an explicit nil
### GetFormattedAdjustmentTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedAdjustmentTaxAmount() interface{}`

GetFormattedAdjustmentTaxAmount returns the FormattedAdjustmentTaxAmount field if non-nil, zero value otherwise.

### GetFormattedAdjustmentTaxAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedAdjustmentTaxAmountOk() (*interface{}, bool)`

GetFormattedAdjustmentTaxAmountOk returns a tuple with the FormattedAdjustmentTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAdjustmentTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedAdjustmentTaxAmount(v interface{})`

SetFormattedAdjustmentTaxAmount sets FormattedAdjustmentTaxAmount field to given value.

### HasFormattedAdjustmentTaxAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedAdjustmentTaxAmount() bool`

HasFormattedAdjustmentTaxAmount returns a boolean if a field has been set.

### SetFormattedAdjustmentTaxAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedAdjustmentTaxAmountNil(b bool)`

 SetFormattedAdjustmentTaxAmountNil sets the value for FormattedAdjustmentTaxAmount to be an explicit nil

### UnsetFormattedAdjustmentTaxAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedAdjustmentTaxAmount()`

UnsetFormattedAdjustmentTaxAmount ensures that no value is present for FormattedAdjustmentTaxAmount, not even an explicit nil
### GetTotalAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalAmountCents() interface{}`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalAmountCentsOk() (*interface{}, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalAmountCents(v interface{})`

SetTotalAmountCents sets TotalAmountCents field to given value.

### HasTotalAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasTotalAmountCents() bool`

HasTotalAmountCents returns a boolean if a field has been set.

### SetTotalAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalAmountCentsNil(b bool)`

 SetTotalAmountCentsNil sets the value for TotalAmountCents to be an explicit nil

### UnsetTotalAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetTotalAmountCents()`

UnsetTotalAmountCents ensures that no value is present for TotalAmountCents, not even an explicit nil
### GetTotalAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalAmountFloat() interface{}`

GetTotalAmountFloat returns the TotalAmountFloat field if non-nil, zero value otherwise.

### GetTotalAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalAmountFloatOk() (*interface{}, bool)`

GetTotalAmountFloatOk returns a tuple with the TotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalAmountFloat(v interface{})`

SetTotalAmountFloat sets TotalAmountFloat field to given value.

### HasTotalAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasTotalAmountFloat() bool`

HasTotalAmountFloat returns a boolean if a field has been set.

### SetTotalAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalAmountFloatNil(b bool)`

 SetTotalAmountFloatNil sets the value for TotalAmountFloat to be an explicit nil

### UnsetTotalAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetTotalAmountFloat()`

UnsetTotalAmountFloat ensures that no value is present for TotalAmountFloat, not even an explicit nil
### GetFormattedTotalAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedTotalAmount() interface{}`

GetFormattedTotalAmount returns the FormattedTotalAmount field if non-nil, zero value otherwise.

### GetFormattedTotalAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedTotalAmountOk() (*interface{}, bool)`

GetFormattedTotalAmountOk returns a tuple with the FormattedTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedTotalAmount(v interface{})`

SetFormattedTotalAmount sets FormattedTotalAmount field to given value.

### HasFormattedTotalAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedTotalAmount() bool`

HasFormattedTotalAmount returns a boolean if a field has been set.

### SetFormattedTotalAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedTotalAmountNil(b bool)`

 SetFormattedTotalAmountNil sets the value for FormattedTotalAmount to be an explicit nil

### UnsetFormattedTotalAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedTotalAmount()`

UnsetFormattedTotalAmount ensures that no value is present for FormattedTotalAmount, not even an explicit nil
### GetTotalTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalTaxableAmountCents() interface{}`

GetTotalTaxableAmountCents returns the TotalTaxableAmountCents field if non-nil, zero value otherwise.

### GetTotalTaxableAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalTaxableAmountCentsOk() (*interface{}, bool)`

GetTotalTaxableAmountCentsOk returns a tuple with the TotalTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalTaxableAmountCents(v interface{})`

SetTotalTaxableAmountCents sets TotalTaxableAmountCents field to given value.

### HasTotalTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasTotalTaxableAmountCents() bool`

HasTotalTaxableAmountCents returns a boolean if a field has been set.

### SetTotalTaxableAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalTaxableAmountCentsNil(b bool)`

 SetTotalTaxableAmountCentsNil sets the value for TotalTaxableAmountCents to be an explicit nil

### UnsetTotalTaxableAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetTotalTaxableAmountCents()`

UnsetTotalTaxableAmountCents ensures that no value is present for TotalTaxableAmountCents, not even an explicit nil
### GetTotalTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalTaxableAmountFloat() interface{}`

GetTotalTaxableAmountFloat returns the TotalTaxableAmountFloat field if non-nil, zero value otherwise.

### GetTotalTaxableAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalTaxableAmountFloatOk() (*interface{}, bool)`

GetTotalTaxableAmountFloatOk returns a tuple with the TotalTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalTaxableAmountFloat(v interface{})`

SetTotalTaxableAmountFloat sets TotalTaxableAmountFloat field to given value.

### HasTotalTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasTotalTaxableAmountFloat() bool`

HasTotalTaxableAmountFloat returns a boolean if a field has been set.

### SetTotalTaxableAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalTaxableAmountFloatNil(b bool)`

 SetTotalTaxableAmountFloatNil sets the value for TotalTaxableAmountFloat to be an explicit nil

### UnsetTotalTaxableAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetTotalTaxableAmountFloat()`

UnsetTotalTaxableAmountFloat ensures that no value is present for TotalTaxableAmountFloat, not even an explicit nil
### GetFormattedTotalTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedTotalTaxableAmount() interface{}`

GetFormattedTotalTaxableAmount returns the FormattedTotalTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedTotalTaxableAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedTotalTaxableAmountOk() (*interface{}, bool)`

GetFormattedTotalTaxableAmountOk returns a tuple with the FormattedTotalTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedTotalTaxableAmount(v interface{})`

SetFormattedTotalTaxableAmount sets FormattedTotalTaxableAmount field to given value.

### HasFormattedTotalTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedTotalTaxableAmount() bool`

HasFormattedTotalTaxableAmount returns a boolean if a field has been set.

### SetFormattedTotalTaxableAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedTotalTaxableAmountNil(b bool)`

 SetFormattedTotalTaxableAmountNil sets the value for FormattedTotalTaxableAmount to be an explicit nil

### UnsetFormattedTotalTaxableAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedTotalTaxableAmount()`

UnsetFormattedTotalTaxableAmount ensures that no value is present for FormattedTotalTaxableAmount, not even an explicit nil
### GetSubtotalTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalTaxableAmountCents() interface{}`

GetSubtotalTaxableAmountCents returns the SubtotalTaxableAmountCents field if non-nil, zero value otherwise.

### GetSubtotalTaxableAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalTaxableAmountCentsOk() (*interface{}, bool)`

GetSubtotalTaxableAmountCentsOk returns a tuple with the SubtotalTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalTaxableAmountCents(v interface{})`

SetSubtotalTaxableAmountCents sets SubtotalTaxableAmountCents field to given value.

### HasSubtotalTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasSubtotalTaxableAmountCents() bool`

HasSubtotalTaxableAmountCents returns a boolean if a field has been set.

### SetSubtotalTaxableAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalTaxableAmountCentsNil(b bool)`

 SetSubtotalTaxableAmountCentsNil sets the value for SubtotalTaxableAmountCents to be an explicit nil

### UnsetSubtotalTaxableAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetSubtotalTaxableAmountCents()`

UnsetSubtotalTaxableAmountCents ensures that no value is present for SubtotalTaxableAmountCents, not even an explicit nil
### GetSubtotalTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalTaxableAmountFloat() interface{}`

GetSubtotalTaxableAmountFloat returns the SubtotalTaxableAmountFloat field if non-nil, zero value otherwise.

### GetSubtotalTaxableAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubtotalTaxableAmountFloatOk() (*interface{}, bool)`

GetSubtotalTaxableAmountFloatOk returns a tuple with the SubtotalTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalTaxableAmountFloat(v interface{})`

SetSubtotalTaxableAmountFloat sets SubtotalTaxableAmountFloat field to given value.

### HasSubtotalTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasSubtotalTaxableAmountFloat() bool`

HasSubtotalTaxableAmountFloat returns a boolean if a field has been set.

### SetSubtotalTaxableAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubtotalTaxableAmountFloatNil(b bool)`

 SetSubtotalTaxableAmountFloatNil sets the value for SubtotalTaxableAmountFloat to be an explicit nil

### UnsetSubtotalTaxableAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetSubtotalTaxableAmountFloat()`

UnsetSubtotalTaxableAmountFloat ensures that no value is present for SubtotalTaxableAmountFloat, not even an explicit nil
### GetFormattedSubtotalTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedSubtotalTaxableAmount() interface{}`

GetFormattedSubtotalTaxableAmount returns the FormattedSubtotalTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedSubtotalTaxableAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedSubtotalTaxableAmountOk() (*interface{}, bool)`

GetFormattedSubtotalTaxableAmountOk returns a tuple with the FormattedSubtotalTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSubtotalTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedSubtotalTaxableAmount(v interface{})`

SetFormattedSubtotalTaxableAmount sets FormattedSubtotalTaxableAmount field to given value.

### HasFormattedSubtotalTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedSubtotalTaxableAmount() bool`

HasFormattedSubtotalTaxableAmount returns a boolean if a field has been set.

### SetFormattedSubtotalTaxableAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedSubtotalTaxableAmountNil(b bool)`

 SetFormattedSubtotalTaxableAmountNil sets the value for FormattedSubtotalTaxableAmount to be an explicit nil

### UnsetFormattedSubtotalTaxableAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedSubtotalTaxableAmount()`

UnsetFormattedSubtotalTaxableAmount ensures that no value is present for FormattedSubtotalTaxableAmount, not even an explicit nil
### GetShippingTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingTaxableAmountCents() interface{}`

GetShippingTaxableAmountCents returns the ShippingTaxableAmountCents field if non-nil, zero value otherwise.

### GetShippingTaxableAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingTaxableAmountCentsOk() (*interface{}, bool)`

GetShippingTaxableAmountCentsOk returns a tuple with the ShippingTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingTaxableAmountCents(v interface{})`

SetShippingTaxableAmountCents sets ShippingTaxableAmountCents field to given value.

### HasShippingTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasShippingTaxableAmountCents() bool`

HasShippingTaxableAmountCents returns a boolean if a field has been set.

### SetShippingTaxableAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingTaxableAmountCentsNil(b bool)`

 SetShippingTaxableAmountCentsNil sets the value for ShippingTaxableAmountCents to be an explicit nil

### UnsetShippingTaxableAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetShippingTaxableAmountCents()`

UnsetShippingTaxableAmountCents ensures that no value is present for ShippingTaxableAmountCents, not even an explicit nil
### GetShippingTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingTaxableAmountFloat() interface{}`

GetShippingTaxableAmountFloat returns the ShippingTaxableAmountFloat field if non-nil, zero value otherwise.

### GetShippingTaxableAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetShippingTaxableAmountFloatOk() (*interface{}, bool)`

GetShippingTaxableAmountFloatOk returns a tuple with the ShippingTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingTaxableAmountFloat(v interface{})`

SetShippingTaxableAmountFloat sets ShippingTaxableAmountFloat field to given value.

### HasShippingTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasShippingTaxableAmountFloat() bool`

HasShippingTaxableAmountFloat returns a boolean if a field has been set.

### SetShippingTaxableAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetShippingTaxableAmountFloatNil(b bool)`

 SetShippingTaxableAmountFloatNil sets the value for ShippingTaxableAmountFloat to be an explicit nil

### UnsetShippingTaxableAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetShippingTaxableAmountFloat()`

UnsetShippingTaxableAmountFloat ensures that no value is present for ShippingTaxableAmountFloat, not even an explicit nil
### GetFormattedShippingTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedShippingTaxableAmount() interface{}`

GetFormattedShippingTaxableAmount returns the FormattedShippingTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedShippingTaxableAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedShippingTaxableAmountOk() (*interface{}, bool)`

GetFormattedShippingTaxableAmountOk returns a tuple with the FormattedShippingTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedShippingTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedShippingTaxableAmount(v interface{})`

SetFormattedShippingTaxableAmount sets FormattedShippingTaxableAmount field to given value.

### HasFormattedShippingTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedShippingTaxableAmount() bool`

HasFormattedShippingTaxableAmount returns a boolean if a field has been set.

### SetFormattedShippingTaxableAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedShippingTaxableAmountNil(b bool)`

 SetFormattedShippingTaxableAmountNil sets the value for FormattedShippingTaxableAmount to be an explicit nil

### UnsetFormattedShippingTaxableAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedShippingTaxableAmount()`

UnsetFormattedShippingTaxableAmount ensures that no value is present for FormattedShippingTaxableAmount, not even an explicit nil
### GetPaymentMethodTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodTaxableAmountCents() interface{}`

GetPaymentMethodTaxableAmountCents returns the PaymentMethodTaxableAmountCents field if non-nil, zero value otherwise.

### GetPaymentMethodTaxableAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodTaxableAmountCentsOk() (*interface{}, bool)`

GetPaymentMethodTaxableAmountCentsOk returns a tuple with the PaymentMethodTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodTaxableAmountCents(v interface{})`

SetPaymentMethodTaxableAmountCents sets PaymentMethodTaxableAmountCents field to given value.

### HasPaymentMethodTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentMethodTaxableAmountCents() bool`

HasPaymentMethodTaxableAmountCents returns a boolean if a field has been set.

### SetPaymentMethodTaxableAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodTaxableAmountCentsNil(b bool)`

 SetPaymentMethodTaxableAmountCentsNil sets the value for PaymentMethodTaxableAmountCents to be an explicit nil

### UnsetPaymentMethodTaxableAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetPaymentMethodTaxableAmountCents()`

UnsetPaymentMethodTaxableAmountCents ensures that no value is present for PaymentMethodTaxableAmountCents, not even an explicit nil
### GetPaymentMethodTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodTaxableAmountFloat() interface{}`

GetPaymentMethodTaxableAmountFloat returns the PaymentMethodTaxableAmountFloat field if non-nil, zero value otherwise.

### GetPaymentMethodTaxableAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentMethodTaxableAmountFloatOk() (*interface{}, bool)`

GetPaymentMethodTaxableAmountFloatOk returns a tuple with the PaymentMethodTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodTaxableAmountFloat(v interface{})`

SetPaymentMethodTaxableAmountFloat sets PaymentMethodTaxableAmountFloat field to given value.

### HasPaymentMethodTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentMethodTaxableAmountFloat() bool`

HasPaymentMethodTaxableAmountFloat returns a boolean if a field has been set.

### SetPaymentMethodTaxableAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentMethodTaxableAmountFloatNil(b bool)`

 SetPaymentMethodTaxableAmountFloatNil sets the value for PaymentMethodTaxableAmountFloat to be an explicit nil

### UnsetPaymentMethodTaxableAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetPaymentMethodTaxableAmountFloat()`

UnsetPaymentMethodTaxableAmountFloat ensures that no value is present for PaymentMethodTaxableAmountFloat, not even an explicit nil
### GetFormattedPaymentMethodTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedPaymentMethodTaxableAmount() interface{}`

GetFormattedPaymentMethodTaxableAmount returns the FormattedPaymentMethodTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedPaymentMethodTaxableAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedPaymentMethodTaxableAmountOk() (*interface{}, bool)`

GetFormattedPaymentMethodTaxableAmountOk returns a tuple with the FormattedPaymentMethodTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPaymentMethodTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedPaymentMethodTaxableAmount(v interface{})`

SetFormattedPaymentMethodTaxableAmount sets FormattedPaymentMethodTaxableAmount field to given value.

### HasFormattedPaymentMethodTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedPaymentMethodTaxableAmount() bool`

HasFormattedPaymentMethodTaxableAmount returns a boolean if a field has been set.

### SetFormattedPaymentMethodTaxableAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedPaymentMethodTaxableAmountNil(b bool)`

 SetFormattedPaymentMethodTaxableAmountNil sets the value for FormattedPaymentMethodTaxableAmount to be an explicit nil

### UnsetFormattedPaymentMethodTaxableAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedPaymentMethodTaxableAmount()`

UnsetFormattedPaymentMethodTaxableAmount ensures that no value is present for FormattedPaymentMethodTaxableAmount, not even an explicit nil
### GetAdjustmentTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentTaxableAmountCents() interface{}`

GetAdjustmentTaxableAmountCents returns the AdjustmentTaxableAmountCents field if non-nil, zero value otherwise.

### GetAdjustmentTaxableAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentTaxableAmountCentsOk() (*interface{}, bool)`

GetAdjustmentTaxableAmountCentsOk returns a tuple with the AdjustmentTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentTaxableAmountCents(v interface{})`

SetAdjustmentTaxableAmountCents sets AdjustmentTaxableAmountCents field to given value.

### HasAdjustmentTaxableAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasAdjustmentTaxableAmountCents() bool`

HasAdjustmentTaxableAmountCents returns a boolean if a field has been set.

### SetAdjustmentTaxableAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentTaxableAmountCentsNil(b bool)`

 SetAdjustmentTaxableAmountCentsNil sets the value for AdjustmentTaxableAmountCents to be an explicit nil

### UnsetAdjustmentTaxableAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetAdjustmentTaxableAmountCents()`

UnsetAdjustmentTaxableAmountCents ensures that no value is present for AdjustmentTaxableAmountCents, not even an explicit nil
### GetAdjustmentTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentTaxableAmountFloat() interface{}`

GetAdjustmentTaxableAmountFloat returns the AdjustmentTaxableAmountFloat field if non-nil, zero value otherwise.

### GetAdjustmentTaxableAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetAdjustmentTaxableAmountFloatOk() (*interface{}, bool)`

GetAdjustmentTaxableAmountFloatOk returns a tuple with the AdjustmentTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentTaxableAmountFloat(v interface{})`

SetAdjustmentTaxableAmountFloat sets AdjustmentTaxableAmountFloat field to given value.

### HasAdjustmentTaxableAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasAdjustmentTaxableAmountFloat() bool`

HasAdjustmentTaxableAmountFloat returns a boolean if a field has been set.

### SetAdjustmentTaxableAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetAdjustmentTaxableAmountFloatNil(b bool)`

 SetAdjustmentTaxableAmountFloatNil sets the value for AdjustmentTaxableAmountFloat to be an explicit nil

### UnsetAdjustmentTaxableAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetAdjustmentTaxableAmountFloat()`

UnsetAdjustmentTaxableAmountFloat ensures that no value is present for AdjustmentTaxableAmountFloat, not even an explicit nil
### GetFormattedAdjustmentTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedAdjustmentTaxableAmount() interface{}`

GetFormattedAdjustmentTaxableAmount returns the FormattedAdjustmentTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedAdjustmentTaxableAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedAdjustmentTaxableAmountOk() (*interface{}, bool)`

GetFormattedAdjustmentTaxableAmountOk returns a tuple with the FormattedAdjustmentTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAdjustmentTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedAdjustmentTaxableAmount(v interface{})`

SetFormattedAdjustmentTaxableAmount sets FormattedAdjustmentTaxableAmount field to given value.

### HasFormattedAdjustmentTaxableAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedAdjustmentTaxableAmount() bool`

HasFormattedAdjustmentTaxableAmount returns a boolean if a field has been set.

### SetFormattedAdjustmentTaxableAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedAdjustmentTaxableAmountNil(b bool)`

 SetFormattedAdjustmentTaxableAmountNil sets the value for FormattedAdjustmentTaxableAmount to be an explicit nil

### UnsetFormattedAdjustmentTaxableAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedAdjustmentTaxableAmount()`

UnsetFormattedAdjustmentTaxableAmount ensures that no value is present for FormattedAdjustmentTaxableAmount, not even an explicit nil
### GetTotalAmountWithTaxesCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalAmountWithTaxesCents() interface{}`

GetTotalAmountWithTaxesCents returns the TotalAmountWithTaxesCents field if non-nil, zero value otherwise.

### GetTotalAmountWithTaxesCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalAmountWithTaxesCentsOk() (*interface{}, bool)`

GetTotalAmountWithTaxesCentsOk returns a tuple with the TotalAmountWithTaxesCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountWithTaxesCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalAmountWithTaxesCents(v interface{})`

SetTotalAmountWithTaxesCents sets TotalAmountWithTaxesCents field to given value.

### HasTotalAmountWithTaxesCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasTotalAmountWithTaxesCents() bool`

HasTotalAmountWithTaxesCents returns a boolean if a field has been set.

### SetTotalAmountWithTaxesCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalAmountWithTaxesCentsNil(b bool)`

 SetTotalAmountWithTaxesCentsNil sets the value for TotalAmountWithTaxesCents to be an explicit nil

### UnsetTotalAmountWithTaxesCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetTotalAmountWithTaxesCents()`

UnsetTotalAmountWithTaxesCents ensures that no value is present for TotalAmountWithTaxesCents, not even an explicit nil
### GetTotalAmountWithTaxesFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalAmountWithTaxesFloat() interface{}`

GetTotalAmountWithTaxesFloat returns the TotalAmountWithTaxesFloat field if non-nil, zero value otherwise.

### GetTotalAmountWithTaxesFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTotalAmountWithTaxesFloatOk() (*interface{}, bool)`

GetTotalAmountWithTaxesFloatOk returns a tuple with the TotalAmountWithTaxesFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountWithTaxesFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalAmountWithTaxesFloat(v interface{})`

SetTotalAmountWithTaxesFloat sets TotalAmountWithTaxesFloat field to given value.

### HasTotalAmountWithTaxesFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasTotalAmountWithTaxesFloat() bool`

HasTotalAmountWithTaxesFloat returns a boolean if a field has been set.

### SetTotalAmountWithTaxesFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetTotalAmountWithTaxesFloatNil(b bool)`

 SetTotalAmountWithTaxesFloatNil sets the value for TotalAmountWithTaxesFloat to be an explicit nil

### UnsetTotalAmountWithTaxesFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetTotalAmountWithTaxesFloat()`

UnsetTotalAmountWithTaxesFloat ensures that no value is present for TotalAmountWithTaxesFloat, not even an explicit nil
### GetFormattedTotalAmountWithTaxes

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedTotalAmountWithTaxes() interface{}`

GetFormattedTotalAmountWithTaxes returns the FormattedTotalAmountWithTaxes field if non-nil, zero value otherwise.

### GetFormattedTotalAmountWithTaxesOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedTotalAmountWithTaxesOk() (*interface{}, bool)`

GetFormattedTotalAmountWithTaxesOk returns a tuple with the FormattedTotalAmountWithTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalAmountWithTaxes

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedTotalAmountWithTaxes(v interface{})`

SetFormattedTotalAmountWithTaxes sets FormattedTotalAmountWithTaxes field to given value.

### HasFormattedTotalAmountWithTaxes

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedTotalAmountWithTaxes() bool`

HasFormattedTotalAmountWithTaxes returns a boolean if a field has been set.

### SetFormattedTotalAmountWithTaxesNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedTotalAmountWithTaxesNil(b bool)`

 SetFormattedTotalAmountWithTaxesNil sets the value for FormattedTotalAmountWithTaxes to be an explicit nil

### UnsetFormattedTotalAmountWithTaxes
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedTotalAmountWithTaxes()`

UnsetFormattedTotalAmountWithTaxes ensures that no value is present for FormattedTotalAmountWithTaxes, not even an explicit nil
### GetFeesAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetFeesAmountCents() interface{}`

GetFeesAmountCents returns the FeesAmountCents field if non-nil, zero value otherwise.

### GetFeesAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFeesAmountCentsOk() (*interface{}, bool)`

GetFeesAmountCentsOk returns a tuple with the FeesAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetFeesAmountCents(v interface{})`

SetFeesAmountCents sets FeesAmountCents field to given value.

### HasFeesAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasFeesAmountCents() bool`

HasFeesAmountCents returns a boolean if a field has been set.

### SetFeesAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFeesAmountCentsNil(b bool)`

 SetFeesAmountCentsNil sets the value for FeesAmountCents to be an explicit nil

### UnsetFeesAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFeesAmountCents()`

UnsetFeesAmountCents ensures that no value is present for FeesAmountCents, not even an explicit nil
### GetFeesAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetFeesAmountFloat() interface{}`

GetFeesAmountFloat returns the FeesAmountFloat field if non-nil, zero value otherwise.

### GetFeesAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFeesAmountFloatOk() (*interface{}, bool)`

GetFeesAmountFloatOk returns a tuple with the FeesAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetFeesAmountFloat(v interface{})`

SetFeesAmountFloat sets FeesAmountFloat field to given value.

### HasFeesAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasFeesAmountFloat() bool`

HasFeesAmountFloat returns a boolean if a field has been set.

### SetFeesAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFeesAmountFloatNil(b bool)`

 SetFeesAmountFloatNil sets the value for FeesAmountFloat to be an explicit nil

### UnsetFeesAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFeesAmountFloat()`

UnsetFeesAmountFloat ensures that no value is present for FeesAmountFloat, not even an explicit nil
### GetFormattedFeesAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedFeesAmount() interface{}`

GetFormattedFeesAmount returns the FormattedFeesAmount field if non-nil, zero value otherwise.

### GetFormattedFeesAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedFeesAmountOk() (*interface{}, bool)`

GetFormattedFeesAmountOk returns a tuple with the FormattedFeesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedFeesAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedFeesAmount(v interface{})`

SetFormattedFeesAmount sets FormattedFeesAmount field to given value.

### HasFormattedFeesAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedFeesAmount() bool`

HasFormattedFeesAmount returns a boolean if a field has been set.

### SetFormattedFeesAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedFeesAmountNil(b bool)`

 SetFormattedFeesAmountNil sets the value for FormattedFeesAmount to be an explicit nil

### UnsetFormattedFeesAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedFeesAmount()`

UnsetFormattedFeesAmount ensures that no value is present for FormattedFeesAmount, not even an explicit nil
### GetDutyAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) GetDutyAmountCents() interface{}`

GetDutyAmountCents returns the DutyAmountCents field if non-nil, zero value otherwise.

### GetDutyAmountCentsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetDutyAmountCentsOk() (*interface{}, bool)`

GetDutyAmountCentsOk returns a tuple with the DutyAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) SetDutyAmountCents(v interface{})`

SetDutyAmountCents sets DutyAmountCents field to given value.

### HasDutyAmountCents

`func (o *GETOrders200ResponseDataInnerAttributes) HasDutyAmountCents() bool`

HasDutyAmountCents returns a boolean if a field has been set.

### SetDutyAmountCentsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetDutyAmountCentsNil(b bool)`

 SetDutyAmountCentsNil sets the value for DutyAmountCents to be an explicit nil

### UnsetDutyAmountCents
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetDutyAmountCents()`

UnsetDutyAmountCents ensures that no value is present for DutyAmountCents, not even an explicit nil
### GetDutyAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) GetDutyAmountFloat() interface{}`

GetDutyAmountFloat returns the DutyAmountFloat field if non-nil, zero value otherwise.

### GetDutyAmountFloatOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetDutyAmountFloatOk() (*interface{}, bool)`

GetDutyAmountFloatOk returns a tuple with the DutyAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) SetDutyAmountFloat(v interface{})`

SetDutyAmountFloat sets DutyAmountFloat field to given value.

### HasDutyAmountFloat

`func (o *GETOrders200ResponseDataInnerAttributes) HasDutyAmountFloat() bool`

HasDutyAmountFloat returns a boolean if a field has been set.

### SetDutyAmountFloatNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetDutyAmountFloatNil(b bool)`

 SetDutyAmountFloatNil sets the value for DutyAmountFloat to be an explicit nil

### UnsetDutyAmountFloat
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetDutyAmountFloat()`

UnsetDutyAmountFloat ensures that no value is present for DutyAmountFloat, not even an explicit nil
### GetFormattedDutyAmount

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedDutyAmount() interface{}`

GetFormattedDutyAmount returns the FormattedDutyAmount field if non-nil, zero value otherwise.

### GetFormattedDutyAmountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFormattedDutyAmountOk() (*interface{}, bool)`

GetFormattedDutyAmountOk returns a tuple with the FormattedDutyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedDutyAmount

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedDutyAmount(v interface{})`

SetFormattedDutyAmount sets FormattedDutyAmount field to given value.

### HasFormattedDutyAmount

`func (o *GETOrders200ResponseDataInnerAttributes) HasFormattedDutyAmount() bool`

HasFormattedDutyAmount returns a boolean if a field has been set.

### SetFormattedDutyAmountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFormattedDutyAmountNil(b bool)`

 SetFormattedDutyAmountNil sets the value for FormattedDutyAmount to be an explicit nil

### UnsetFormattedDutyAmount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFormattedDutyAmount()`

UnsetFormattedDutyAmount ensures that no value is present for FormattedDutyAmount, not even an explicit nil
### GetSkusCount

`func (o *GETOrders200ResponseDataInnerAttributes) GetSkusCount() interface{}`

GetSkusCount returns the SkusCount field if non-nil, zero value otherwise.

### GetSkusCountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetSkusCountOk() (*interface{}, bool)`

GetSkusCountOk returns a tuple with the SkusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusCount

`func (o *GETOrders200ResponseDataInnerAttributes) SetSkusCount(v interface{})`

SetSkusCount sets SkusCount field to given value.

### HasSkusCount

`func (o *GETOrders200ResponseDataInnerAttributes) HasSkusCount() bool`

HasSkusCount returns a boolean if a field has been set.

### SetSkusCountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetSkusCountNil(b bool)`

 SetSkusCountNil sets the value for SkusCount to be an explicit nil

### UnsetSkusCount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetSkusCount()`

UnsetSkusCount ensures that no value is present for SkusCount, not even an explicit nil
### GetLineItemOptionsCount

`func (o *GETOrders200ResponseDataInnerAttributes) GetLineItemOptionsCount() interface{}`

GetLineItemOptionsCount returns the LineItemOptionsCount field if non-nil, zero value otherwise.

### GetLineItemOptionsCountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetLineItemOptionsCountOk() (*interface{}, bool)`

GetLineItemOptionsCountOk returns a tuple with the LineItemOptionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemOptionsCount

`func (o *GETOrders200ResponseDataInnerAttributes) SetLineItemOptionsCount(v interface{})`

SetLineItemOptionsCount sets LineItemOptionsCount field to given value.

### HasLineItemOptionsCount

`func (o *GETOrders200ResponseDataInnerAttributes) HasLineItemOptionsCount() bool`

HasLineItemOptionsCount returns a boolean if a field has been set.

### SetLineItemOptionsCountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetLineItemOptionsCountNil(b bool)`

 SetLineItemOptionsCountNil sets the value for LineItemOptionsCount to be an explicit nil

### UnsetLineItemOptionsCount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetLineItemOptionsCount()`

UnsetLineItemOptionsCount ensures that no value is present for LineItemOptionsCount, not even an explicit nil
### GetShipmentsCount

`func (o *GETOrders200ResponseDataInnerAttributes) GetShipmentsCount() interface{}`

GetShipmentsCount returns the ShipmentsCount field if non-nil, zero value otherwise.

### GetShipmentsCountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetShipmentsCountOk() (*interface{}, bool)`

GetShipmentsCountOk returns a tuple with the ShipmentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentsCount

`func (o *GETOrders200ResponseDataInnerAttributes) SetShipmentsCount(v interface{})`

SetShipmentsCount sets ShipmentsCount field to given value.

### HasShipmentsCount

`func (o *GETOrders200ResponseDataInnerAttributes) HasShipmentsCount() bool`

HasShipmentsCount returns a boolean if a field has been set.

### SetShipmentsCountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetShipmentsCountNil(b bool)`

 SetShipmentsCountNil sets the value for ShipmentsCount to be an explicit nil

### UnsetShipmentsCount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetShipmentsCount()`

UnsetShipmentsCount ensures that no value is present for ShipmentsCount, not even an explicit nil
### GetTaxCalculationsCount

`func (o *GETOrders200ResponseDataInnerAttributes) GetTaxCalculationsCount() interface{}`

GetTaxCalculationsCount returns the TaxCalculationsCount field if non-nil, zero value otherwise.

### GetTaxCalculationsCountOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTaxCalculationsCountOk() (*interface{}, bool)`

GetTaxCalculationsCountOk returns a tuple with the TaxCalculationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationsCount

`func (o *GETOrders200ResponseDataInnerAttributes) SetTaxCalculationsCount(v interface{})`

SetTaxCalculationsCount sets TaxCalculationsCount field to given value.

### HasTaxCalculationsCount

`func (o *GETOrders200ResponseDataInnerAttributes) HasTaxCalculationsCount() bool`

HasTaxCalculationsCount returns a boolean if a field has been set.

### SetTaxCalculationsCountNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetTaxCalculationsCountNil(b bool)`

 SetTaxCalculationsCountNil sets the value for TaxCalculationsCount to be an explicit nil

### UnsetTaxCalculationsCount
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetTaxCalculationsCount()`

UnsetTaxCalculationsCount ensures that no value is present for TaxCalculationsCount, not even an explicit nil
### GetPaymentSourceDetails

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentSourceDetails() interface{}`

GetPaymentSourceDetails returns the PaymentSourceDetails field if non-nil, zero value otherwise.

### GetPaymentSourceDetailsOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentSourceDetailsOk() (*interface{}, bool)`

GetPaymentSourceDetailsOk returns a tuple with the PaymentSourceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceDetails

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentSourceDetails(v interface{})`

SetPaymentSourceDetails sets PaymentSourceDetails field to given value.

### HasPaymentSourceDetails

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentSourceDetails() bool`

HasPaymentSourceDetails returns a boolean if a field has been set.

### SetPaymentSourceDetailsNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentSourceDetailsNil(b bool)`

 SetPaymentSourceDetailsNil sets the value for PaymentSourceDetails to be an explicit nil

### UnsetPaymentSourceDetails
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetPaymentSourceDetails()`

UnsetPaymentSourceDetails ensures that no value is present for PaymentSourceDetails, not even an explicit nil
### GetToken

`func (o *GETOrders200ResponseDataInnerAttributes) GetToken() interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTokenOk() (*interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GETOrders200ResponseDataInnerAttributes) SetToken(v interface{})`

SetToken sets Token field to given value.

### HasToken

`func (o *GETOrders200ResponseDataInnerAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetCartUrl

`func (o *GETOrders200ResponseDataInnerAttributes) GetCartUrl() interface{}`

GetCartUrl returns the CartUrl field if non-nil, zero value otherwise.

### GetCartUrlOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetCartUrlOk() (*interface{}, bool)`

GetCartUrlOk returns a tuple with the CartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartUrl

`func (o *GETOrders200ResponseDataInnerAttributes) SetCartUrl(v interface{})`

SetCartUrl sets CartUrl field to given value.

### HasCartUrl

`func (o *GETOrders200ResponseDataInnerAttributes) HasCartUrl() bool`

HasCartUrl returns a boolean if a field has been set.

### SetCartUrlNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetCartUrlNil(b bool)`

 SetCartUrlNil sets the value for CartUrl to be an explicit nil

### UnsetCartUrl
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetCartUrl()`

UnsetCartUrl ensures that no value is present for CartUrl, not even an explicit nil
### GetReturnUrl

`func (o *GETOrders200ResponseDataInnerAttributes) GetReturnUrl() interface{}`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetReturnUrlOk() (*interface{}, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *GETOrders200ResponseDataInnerAttributes) SetReturnUrl(v interface{})`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *GETOrders200ResponseDataInnerAttributes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### SetReturnUrlNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetTermsUrl

`func (o *GETOrders200ResponseDataInnerAttributes) GetTermsUrl() interface{}`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetTermsUrlOk() (*interface{}, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *GETOrders200ResponseDataInnerAttributes) SetTermsUrl(v interface{})`

SetTermsUrl sets TermsUrl field to given value.

### HasTermsUrl

`func (o *GETOrders200ResponseDataInnerAttributes) HasTermsUrl() bool`

HasTermsUrl returns a boolean if a field has been set.

### SetTermsUrlNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetTermsUrlNil(b bool)`

 SetTermsUrlNil sets the value for TermsUrl to be an explicit nil

### UnsetTermsUrl
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetTermsUrl()`

UnsetTermsUrl ensures that no value is present for TermsUrl, not even an explicit nil
### GetPrivacyUrl

`func (o *GETOrders200ResponseDataInnerAttributes) GetPrivacyUrl() interface{}`

GetPrivacyUrl returns the PrivacyUrl field if non-nil, zero value otherwise.

### GetPrivacyUrlOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPrivacyUrlOk() (*interface{}, bool)`

GetPrivacyUrlOk returns a tuple with the PrivacyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyUrl

`func (o *GETOrders200ResponseDataInnerAttributes) SetPrivacyUrl(v interface{})`

SetPrivacyUrl sets PrivacyUrl field to given value.

### HasPrivacyUrl

`func (o *GETOrders200ResponseDataInnerAttributes) HasPrivacyUrl() bool`

HasPrivacyUrl returns a boolean if a field has been set.

### SetPrivacyUrlNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetPrivacyUrlNil(b bool)`

 SetPrivacyUrlNil sets the value for PrivacyUrl to be an explicit nil

### UnsetPrivacyUrl
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetPrivacyUrl()`

UnsetPrivacyUrl ensures that no value is present for PrivacyUrl, not even an explicit nil
### GetCheckoutUrl

`func (o *GETOrders200ResponseDataInnerAttributes) GetCheckoutUrl() interface{}`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetCheckoutUrlOk() (*interface{}, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *GETOrders200ResponseDataInnerAttributes) SetCheckoutUrl(v interface{})`

SetCheckoutUrl sets CheckoutUrl field to given value.

### HasCheckoutUrl

`func (o *GETOrders200ResponseDataInnerAttributes) HasCheckoutUrl() bool`

HasCheckoutUrl returns a boolean if a field has been set.

### SetCheckoutUrlNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetCheckoutUrlNil(b bool)`

 SetCheckoutUrlNil sets the value for CheckoutUrl to be an explicit nil

### UnsetCheckoutUrl
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetCheckoutUrl()`

UnsetCheckoutUrl ensures that no value is present for CheckoutUrl, not even an explicit nil
### GetPlacedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetPlacedAt() interface{}`

GetPlacedAt returns the PlacedAt field if non-nil, zero value otherwise.

### GetPlacedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPlacedAtOk() (*interface{}, bool)`

GetPlacedAtOk returns a tuple with the PlacedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetPlacedAt(v interface{})`

SetPlacedAt sets PlacedAt field to given value.

### HasPlacedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasPlacedAt() bool`

HasPlacedAt returns a boolean if a field has been set.

### SetPlacedAtNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetPlacedAtNil(b bool)`

 SetPlacedAtNil sets the value for PlacedAt to be an explicit nil

### UnsetPlacedAt
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetPlacedAt()`

UnsetPlacedAt ensures that no value is present for PlacedAt, not even an explicit nil
### GetApprovedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetApprovedAt() interface{}`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetApprovedAtOk() (*interface{}, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetApprovedAt(v interface{})`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### SetApprovedAtNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetApprovedAtNil(b bool)`

 SetApprovedAtNil sets the value for ApprovedAt to be an explicit nil

### UnsetApprovedAt
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetApprovedAt()`

UnsetApprovedAt ensures that no value is present for ApprovedAt, not even an explicit nil
### GetCancelledAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetCancelledAt() interface{}`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetCancelledAtOk() (*interface{}, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetCancelledAt(v interface{})`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### SetCancelledAtNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetCancelledAtNil(b bool)`

 SetCancelledAtNil sets the value for CancelledAt to be an explicit nil

### UnsetCancelledAt
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetCancelledAt()`

UnsetCancelledAt ensures that no value is present for CancelledAt, not even an explicit nil
### GetPaymentUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentUpdatedAt() interface{}`

GetPaymentUpdatedAt returns the PaymentUpdatedAt field if non-nil, zero value otherwise.

### GetPaymentUpdatedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetPaymentUpdatedAtOk() (*interface{}, bool)`

GetPaymentUpdatedAtOk returns a tuple with the PaymentUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentUpdatedAt(v interface{})`

SetPaymentUpdatedAt sets PaymentUpdatedAt field to given value.

### HasPaymentUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasPaymentUpdatedAt() bool`

HasPaymentUpdatedAt returns a boolean if a field has been set.

### SetPaymentUpdatedAtNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetPaymentUpdatedAtNil(b bool)`

 SetPaymentUpdatedAtNil sets the value for PaymentUpdatedAt to be an explicit nil

### UnsetPaymentUpdatedAt
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetPaymentUpdatedAt()`

UnsetPaymentUpdatedAt ensures that no value is present for PaymentUpdatedAt, not even an explicit nil
### GetFulfillmentUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetFulfillmentUpdatedAt() interface{}`

GetFulfillmentUpdatedAt returns the FulfillmentUpdatedAt field if non-nil, zero value otherwise.

### GetFulfillmentUpdatedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetFulfillmentUpdatedAtOk() (*interface{}, bool)`

GetFulfillmentUpdatedAtOk returns a tuple with the FulfillmentUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetFulfillmentUpdatedAt(v interface{})`

SetFulfillmentUpdatedAt sets FulfillmentUpdatedAt field to given value.

### HasFulfillmentUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasFulfillmentUpdatedAt() bool`

HasFulfillmentUpdatedAt returns a boolean if a field has been set.

### SetFulfillmentUpdatedAtNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetFulfillmentUpdatedAtNil(b bool)`

 SetFulfillmentUpdatedAtNil sets the value for FulfillmentUpdatedAt to be an explicit nil

### UnsetFulfillmentUpdatedAt
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetFulfillmentUpdatedAt()`

UnsetFulfillmentUpdatedAt ensures that no value is present for FulfillmentUpdatedAt, not even an explicit nil
### GetRefreshedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetRefreshedAt() interface{}`

GetRefreshedAt returns the RefreshedAt field if non-nil, zero value otherwise.

### GetRefreshedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetRefreshedAtOk() (*interface{}, bool)`

GetRefreshedAtOk returns a tuple with the RefreshedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetRefreshedAt(v interface{})`

SetRefreshedAt sets RefreshedAt field to given value.

### HasRefreshedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasRefreshedAt() bool`

HasRefreshedAt returns a boolean if a field has been set.

### SetRefreshedAtNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetRefreshedAtNil(b bool)`

 SetRefreshedAtNil sets the value for RefreshedAt to be an explicit nil

### UnsetRefreshedAt
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetRefreshedAt()`

UnsetRefreshedAt ensures that no value is present for RefreshedAt, not even an explicit nil
### GetArchivedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetArchivedAt() interface{}`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetArchivedAtOk() (*interface{}, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetArchivedAt(v interface{})`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### SetArchivedAtNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetArchivedAtNil(b bool)`

 SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil

### UnsetArchivedAt
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetArchivedAt()`

UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil
### GetExpiresAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetSubscriptionCreatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubscriptionCreatedAt() interface{}`

GetSubscriptionCreatedAt returns the SubscriptionCreatedAt field if non-nil, zero value otherwise.

### GetSubscriptionCreatedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetSubscriptionCreatedAtOk() (*interface{}, bool)`

GetSubscriptionCreatedAtOk returns a tuple with the SubscriptionCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCreatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubscriptionCreatedAt(v interface{})`

SetSubscriptionCreatedAt sets SubscriptionCreatedAt field to given value.

### HasSubscriptionCreatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasSubscriptionCreatedAt() bool`

HasSubscriptionCreatedAt returns a boolean if a field has been set.

### SetSubscriptionCreatedAtNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetSubscriptionCreatedAtNil(b bool)`

 SetSubscriptionCreatedAtNil sets the value for SubscriptionCreatedAt to be an explicit nil

### UnsetSubscriptionCreatedAt
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetSubscriptionCreatedAt()`

UnsetSubscriptionCreatedAt ensures that no value is present for SubscriptionCreatedAt, not even an explicit nil
### GetCreatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETOrders200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETOrders200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETOrders200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETOrders200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETOrders200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETOrders200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETOrders200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETOrders200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETOrders200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETOrders200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETOrders200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETOrders200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETOrders200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


