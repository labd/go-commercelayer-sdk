# OrderUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autorefresh** | Pointer to **bool** | Save this attribute as &#39;false&#39; if you want prevent the order to be refreshed automatically at each change (much faster). | [optional] 
**Guest** | Pointer to **bool** | Indicates if the order has been placed as guest. | [optional] 
**CustomerEmail** | Pointer to **string** | The email address of the associated customer. When creating or updating an order, this is a shortcut to find or create the associated customer by email. | [optional] 
**CustomerPassword** | Pointer to **string** | The password of the associated customer. When creating or updating an order, this is a shortcut to sign up the associated customer. | [optional] 
**LanguageCode** | Pointer to **string** | The preferred language code (ISO 639-1) to be used when communicating with the customer. This can be useful when sending the order to 3rd party marketing tools and CRMs. If the language is supported, the hosted checkout will be localized accordingly. | [optional] 
**ShippingCountryCodeLock** | Pointer to **string** | The country code that you want the shipping address to be locked to. This can be useful to make sure the shipping address belongs to a given shipping country, e.g. the one selected in a country selector page. | [optional] 
**CouponCode** | Pointer to **string** | The coupon code to be used for the order. If valid, it triggers a promotion adding a discount line item to the order. | [optional] 
**GiftCardCode** | Pointer to **string** | The gift card code (at least the first 8 characters) to be used for the order. If valid, it uses the gift card balance to pay for the order. | [optional] 
**GiftCardOrCouponCode** | Pointer to **string** | The gift card or coupon code (at least the first 8 characters) to be used for the order. If a gift card mathes, it uses the gift card balance to pay for the order. Otherwise it tries to find a valid coupon code and applies the associated discount. | [optional] 
**CartUrl** | Pointer to **string** | The cart url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**ReturnUrl** | Pointer to **string** | The return url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**TermsUrl** | Pointer to **string** | The terms and conditions url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**PrivacyUrl** | Pointer to **string** | The privacy policy url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**Archive** | Pointer to **bool** | Send this attribute if you want to archive the order. | [optional] 
**Unarchive** | Pointer to **bool** | Send this attribute if you want to unarchive the order. | [optional] 
**Place** | Pointer to **bool** | Send this attribute if you want to place the order. | [optional] 
**Cancel** | Pointer to **bool** | Send this attribute if you want to cancel a placed order. The order&#39;s authorization will be automatically voided. | [optional] 
**Approve** | Pointer to **bool** | Send this attribute if you want to approve a placed order. | [optional] 
**ApproveAndCapture** | Pointer to **bool** | Send this attribute if you want to approve and capture a placed order. | [optional] 
**Authorize** | Pointer to **bool** | Send this attribute if you want to authorize the order&#39;s payment source. | [optional] 
**AuthorizationAmountCents** | Pointer to **int32** | The authorization amount, in cents. | [optional] 
**Capture** | Pointer to **bool** | Send this attribute if you want to capture an approved order. | [optional] 
**Refund** | Pointer to **bool** | Send this attribute if you want to refund a captured order. | [optional] 
**UpdateTaxes** | Pointer to **bool** | Send this attribute if you want to force tax calculation for this order (a tax calculator must be associated to the order&#39;s market). | [optional] 
**NullifyPaymentSource** | Pointer to **bool** | Send this attribute if you want to nullify the payment source for this order. | [optional] 
**BillingAddressCloneId** | Pointer to **string** | The id of the address that you want to clone to create the order&#39;s billing address. | [optional] 
**ShippingAddressCloneId** | Pointer to **string** | The id of the address that you want to clone to create the order&#39;s shipping address. | [optional] 
**CustomerPaymentSourceId** | Pointer to **string** | The id of the customer payment source (i.e. credit card) that you want to use as the order&#39;s payment source. | [optional] 
**ShippingAddressSameAsBilling** | Pointer to **bool** | Send this attribute if you want the shipping address to be cloned from the order&#39;s billing address. | [optional] 
**BillingAddressSameAsShipping** | Pointer to **bool** | Send this attribute if you want the billing address to be cloned from the order&#39;s shipping address. | [optional] 
**CommitInvoice** | Pointer to **bool** | Send this attribute if you want commit the sales tax invoice to the associated tax calculator (currently supported by Avalara). | [optional] 
**RefundInvoice** | Pointer to **bool** | Send this attribute if you want refund the sales tax invoice to the associated tax calculator (currently supported by Avalara). | [optional] 
**SavePaymentSourceToCustomerWallet** | Pointer to **bool** | Send this attribute if you want the order&#39;s payment source to be saved in the customer&#39;s wallet as a customer payment source. | [optional] 
**SaveShippingAddressToCustomerAddressBook** | Pointer to **bool** | Send this attribute if you want the order&#39;s shipping address to be saved in the customer&#39;s address book as a customer address. | [optional] 
**SaveBillingAddressToCustomerAddressBook** | Pointer to **bool** | Send this attribute if you want the order&#39;s billing address to be saved in the customer&#39;s address book as a customer address. | [optional] 
**Refresh** | Pointer to **bool** | Send this attribute if you want to manually refresh the order. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewOrderUpdateDataAttributes

`func NewOrderUpdateDataAttributes() *OrderUpdateDataAttributes`

NewOrderUpdateDataAttributes instantiates a new OrderUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderUpdateDataAttributesWithDefaults

`func NewOrderUpdateDataAttributesWithDefaults() *OrderUpdateDataAttributes`

NewOrderUpdateDataAttributesWithDefaults instantiates a new OrderUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutorefresh

`func (o *OrderUpdateDataAttributes) GetAutorefresh() bool`

GetAutorefresh returns the Autorefresh field if non-nil, zero value otherwise.

### GetAutorefreshOk

`func (o *OrderUpdateDataAttributes) GetAutorefreshOk() (*bool, bool)`

GetAutorefreshOk returns a tuple with the Autorefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutorefresh

`func (o *OrderUpdateDataAttributes) SetAutorefresh(v bool)`

SetAutorefresh sets Autorefresh field to given value.

### HasAutorefresh

`func (o *OrderUpdateDataAttributes) HasAutorefresh() bool`

HasAutorefresh returns a boolean if a field has been set.

### GetGuest

`func (o *OrderUpdateDataAttributes) GetGuest() bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *OrderUpdateDataAttributes) GetGuestOk() (*bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *OrderUpdateDataAttributes) SetGuest(v bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *OrderUpdateDataAttributes) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *OrderUpdateDataAttributes) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *OrderUpdateDataAttributes) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *OrderUpdateDataAttributes) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *OrderUpdateDataAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetCustomerPassword

`func (o *OrderUpdateDataAttributes) GetCustomerPassword() string`

GetCustomerPassword returns the CustomerPassword field if non-nil, zero value otherwise.

### GetCustomerPasswordOk

`func (o *OrderUpdateDataAttributes) GetCustomerPasswordOk() (*string, bool)`

GetCustomerPasswordOk returns a tuple with the CustomerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPassword

`func (o *OrderUpdateDataAttributes) SetCustomerPassword(v string)`

SetCustomerPassword sets CustomerPassword field to given value.

### HasCustomerPassword

`func (o *OrderUpdateDataAttributes) HasCustomerPassword() bool`

HasCustomerPassword returns a boolean if a field has been set.

### GetLanguageCode

`func (o *OrderUpdateDataAttributes) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *OrderUpdateDataAttributes) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *OrderUpdateDataAttributes) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *OrderUpdateDataAttributes) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetShippingCountryCodeLock

`func (o *OrderUpdateDataAttributes) GetShippingCountryCodeLock() string`

GetShippingCountryCodeLock returns the ShippingCountryCodeLock field if non-nil, zero value otherwise.

### GetShippingCountryCodeLockOk

`func (o *OrderUpdateDataAttributes) GetShippingCountryCodeLockOk() (*string, bool)`

GetShippingCountryCodeLockOk returns a tuple with the ShippingCountryCodeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCountryCodeLock

`func (o *OrderUpdateDataAttributes) SetShippingCountryCodeLock(v string)`

SetShippingCountryCodeLock sets ShippingCountryCodeLock field to given value.

### HasShippingCountryCodeLock

`func (o *OrderUpdateDataAttributes) HasShippingCountryCodeLock() bool`

HasShippingCountryCodeLock returns a boolean if a field has been set.

### GetCouponCode

`func (o *OrderUpdateDataAttributes) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *OrderUpdateDataAttributes) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *OrderUpdateDataAttributes) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *OrderUpdateDataAttributes) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.

### GetGiftCardCode

`func (o *OrderUpdateDataAttributes) GetGiftCardCode() string`

GetGiftCardCode returns the GiftCardCode field if non-nil, zero value otherwise.

### GetGiftCardCodeOk

`func (o *OrderUpdateDataAttributes) GetGiftCardCodeOk() (*string, bool)`

GetGiftCardCodeOk returns a tuple with the GiftCardCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardCode

`func (o *OrderUpdateDataAttributes) SetGiftCardCode(v string)`

SetGiftCardCode sets GiftCardCode field to given value.

### HasGiftCardCode

`func (o *OrderUpdateDataAttributes) HasGiftCardCode() bool`

HasGiftCardCode returns a boolean if a field has been set.

### GetGiftCardOrCouponCode

`func (o *OrderUpdateDataAttributes) GetGiftCardOrCouponCode() string`

GetGiftCardOrCouponCode returns the GiftCardOrCouponCode field if non-nil, zero value otherwise.

### GetGiftCardOrCouponCodeOk

`func (o *OrderUpdateDataAttributes) GetGiftCardOrCouponCodeOk() (*string, bool)`

GetGiftCardOrCouponCodeOk returns a tuple with the GiftCardOrCouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardOrCouponCode

`func (o *OrderUpdateDataAttributes) SetGiftCardOrCouponCode(v string)`

SetGiftCardOrCouponCode sets GiftCardOrCouponCode field to given value.

### HasGiftCardOrCouponCode

`func (o *OrderUpdateDataAttributes) HasGiftCardOrCouponCode() bool`

HasGiftCardOrCouponCode returns a boolean if a field has been set.

### GetCartUrl

`func (o *OrderUpdateDataAttributes) GetCartUrl() string`

GetCartUrl returns the CartUrl field if non-nil, zero value otherwise.

### GetCartUrlOk

`func (o *OrderUpdateDataAttributes) GetCartUrlOk() (*string, bool)`

GetCartUrlOk returns a tuple with the CartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartUrl

`func (o *OrderUpdateDataAttributes) SetCartUrl(v string)`

SetCartUrl sets CartUrl field to given value.

### HasCartUrl

`func (o *OrderUpdateDataAttributes) HasCartUrl() bool`

HasCartUrl returns a boolean if a field has been set.

### GetReturnUrl

`func (o *OrderUpdateDataAttributes) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *OrderUpdateDataAttributes) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *OrderUpdateDataAttributes) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *OrderUpdateDataAttributes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetTermsUrl

`func (o *OrderUpdateDataAttributes) GetTermsUrl() string`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *OrderUpdateDataAttributes) GetTermsUrlOk() (*string, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *OrderUpdateDataAttributes) SetTermsUrl(v string)`

SetTermsUrl sets TermsUrl field to given value.

### HasTermsUrl

`func (o *OrderUpdateDataAttributes) HasTermsUrl() bool`

HasTermsUrl returns a boolean if a field has been set.

### GetPrivacyUrl

`func (o *OrderUpdateDataAttributes) GetPrivacyUrl() string`

GetPrivacyUrl returns the PrivacyUrl field if non-nil, zero value otherwise.

### GetPrivacyUrlOk

`func (o *OrderUpdateDataAttributes) GetPrivacyUrlOk() (*string, bool)`

GetPrivacyUrlOk returns a tuple with the PrivacyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyUrl

`func (o *OrderUpdateDataAttributes) SetPrivacyUrl(v string)`

SetPrivacyUrl sets PrivacyUrl field to given value.

### HasPrivacyUrl

`func (o *OrderUpdateDataAttributes) HasPrivacyUrl() bool`

HasPrivacyUrl returns a boolean if a field has been set.

### GetArchive

`func (o *OrderUpdateDataAttributes) GetArchive() bool`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *OrderUpdateDataAttributes) GetArchiveOk() (*bool, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *OrderUpdateDataAttributes) SetArchive(v bool)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *OrderUpdateDataAttributes) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetUnarchive

`func (o *OrderUpdateDataAttributes) GetUnarchive() bool`

GetUnarchive returns the Unarchive field if non-nil, zero value otherwise.

### GetUnarchiveOk

`func (o *OrderUpdateDataAttributes) GetUnarchiveOk() (*bool, bool)`

GetUnarchiveOk returns a tuple with the Unarchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnarchive

`func (o *OrderUpdateDataAttributes) SetUnarchive(v bool)`

SetUnarchive sets Unarchive field to given value.

### HasUnarchive

`func (o *OrderUpdateDataAttributes) HasUnarchive() bool`

HasUnarchive returns a boolean if a field has been set.

### GetPlace

`func (o *OrderUpdateDataAttributes) GetPlace() bool`

GetPlace returns the Place field if non-nil, zero value otherwise.

### GetPlaceOk

`func (o *OrderUpdateDataAttributes) GetPlaceOk() (*bool, bool)`

GetPlaceOk returns a tuple with the Place field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlace

`func (o *OrderUpdateDataAttributes) SetPlace(v bool)`

SetPlace sets Place field to given value.

### HasPlace

`func (o *OrderUpdateDataAttributes) HasPlace() bool`

HasPlace returns a boolean if a field has been set.

### GetCancel

`func (o *OrderUpdateDataAttributes) GetCancel() bool`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *OrderUpdateDataAttributes) GetCancelOk() (*bool, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *OrderUpdateDataAttributes) SetCancel(v bool)`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *OrderUpdateDataAttributes) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### GetApprove

`func (o *OrderUpdateDataAttributes) GetApprove() bool`

GetApprove returns the Approve field if non-nil, zero value otherwise.

### GetApproveOk

`func (o *OrderUpdateDataAttributes) GetApproveOk() (*bool, bool)`

GetApproveOk returns a tuple with the Approve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprove

`func (o *OrderUpdateDataAttributes) SetApprove(v bool)`

SetApprove sets Approve field to given value.

### HasApprove

`func (o *OrderUpdateDataAttributes) HasApprove() bool`

HasApprove returns a boolean if a field has been set.

### GetApproveAndCapture

`func (o *OrderUpdateDataAttributes) GetApproveAndCapture() bool`

GetApproveAndCapture returns the ApproveAndCapture field if non-nil, zero value otherwise.

### GetApproveAndCaptureOk

`func (o *OrderUpdateDataAttributes) GetApproveAndCaptureOk() (*bool, bool)`

GetApproveAndCaptureOk returns a tuple with the ApproveAndCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveAndCapture

`func (o *OrderUpdateDataAttributes) SetApproveAndCapture(v bool)`

SetApproveAndCapture sets ApproveAndCapture field to given value.

### HasApproveAndCapture

`func (o *OrderUpdateDataAttributes) HasApproveAndCapture() bool`

HasApproveAndCapture returns a boolean if a field has been set.

### GetAuthorize

`func (o *OrderUpdateDataAttributes) GetAuthorize() bool`

GetAuthorize returns the Authorize field if non-nil, zero value otherwise.

### GetAuthorizeOk

`func (o *OrderUpdateDataAttributes) GetAuthorizeOk() (*bool, bool)`

GetAuthorizeOk returns a tuple with the Authorize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorize

`func (o *OrderUpdateDataAttributes) SetAuthorize(v bool)`

SetAuthorize sets Authorize field to given value.

### HasAuthorize

`func (o *OrderUpdateDataAttributes) HasAuthorize() bool`

HasAuthorize returns a boolean if a field has been set.

### GetAuthorizationAmountCents

`func (o *OrderUpdateDataAttributes) GetAuthorizationAmountCents() int32`

GetAuthorizationAmountCents returns the AuthorizationAmountCents field if non-nil, zero value otherwise.

### GetAuthorizationAmountCentsOk

`func (o *OrderUpdateDataAttributes) GetAuthorizationAmountCentsOk() (*int32, bool)`

GetAuthorizationAmountCentsOk returns a tuple with the AuthorizationAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationAmountCents

`func (o *OrderUpdateDataAttributes) SetAuthorizationAmountCents(v int32)`

SetAuthorizationAmountCents sets AuthorizationAmountCents field to given value.

### HasAuthorizationAmountCents

`func (o *OrderUpdateDataAttributes) HasAuthorizationAmountCents() bool`

HasAuthorizationAmountCents returns a boolean if a field has been set.

### GetCapture

`func (o *OrderUpdateDataAttributes) GetCapture() bool`

GetCapture returns the Capture field if non-nil, zero value otherwise.

### GetCaptureOk

`func (o *OrderUpdateDataAttributes) GetCaptureOk() (*bool, bool)`

GetCaptureOk returns a tuple with the Capture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapture

`func (o *OrderUpdateDataAttributes) SetCapture(v bool)`

SetCapture sets Capture field to given value.

### HasCapture

`func (o *OrderUpdateDataAttributes) HasCapture() bool`

HasCapture returns a boolean if a field has been set.

### GetRefund

`func (o *OrderUpdateDataAttributes) GetRefund() bool`

GetRefund returns the Refund field if non-nil, zero value otherwise.

### GetRefundOk

`func (o *OrderUpdateDataAttributes) GetRefundOk() (*bool, bool)`

GetRefundOk returns a tuple with the Refund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefund

`func (o *OrderUpdateDataAttributes) SetRefund(v bool)`

SetRefund sets Refund field to given value.

### HasRefund

`func (o *OrderUpdateDataAttributes) HasRefund() bool`

HasRefund returns a boolean if a field has been set.

### GetUpdateTaxes

`func (o *OrderUpdateDataAttributes) GetUpdateTaxes() bool`

GetUpdateTaxes returns the UpdateTaxes field if non-nil, zero value otherwise.

### GetUpdateTaxesOk

`func (o *OrderUpdateDataAttributes) GetUpdateTaxesOk() (*bool, bool)`

GetUpdateTaxesOk returns a tuple with the UpdateTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTaxes

`func (o *OrderUpdateDataAttributes) SetUpdateTaxes(v bool)`

SetUpdateTaxes sets UpdateTaxes field to given value.

### HasUpdateTaxes

`func (o *OrderUpdateDataAttributes) HasUpdateTaxes() bool`

HasUpdateTaxes returns a boolean if a field has been set.

### GetNullifyPaymentSource

`func (o *OrderUpdateDataAttributes) GetNullifyPaymentSource() bool`

GetNullifyPaymentSource returns the NullifyPaymentSource field if non-nil, zero value otherwise.

### GetNullifyPaymentSourceOk

`func (o *OrderUpdateDataAttributes) GetNullifyPaymentSourceOk() (*bool, bool)`

GetNullifyPaymentSourceOk returns a tuple with the NullifyPaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullifyPaymentSource

`func (o *OrderUpdateDataAttributes) SetNullifyPaymentSource(v bool)`

SetNullifyPaymentSource sets NullifyPaymentSource field to given value.

### HasNullifyPaymentSource

`func (o *OrderUpdateDataAttributes) HasNullifyPaymentSource() bool`

HasNullifyPaymentSource returns a boolean if a field has been set.

### GetBillingAddressCloneId

`func (o *OrderUpdateDataAttributes) GetBillingAddressCloneId() string`

GetBillingAddressCloneId returns the BillingAddressCloneId field if non-nil, zero value otherwise.

### GetBillingAddressCloneIdOk

`func (o *OrderUpdateDataAttributes) GetBillingAddressCloneIdOk() (*string, bool)`

GetBillingAddressCloneIdOk returns a tuple with the BillingAddressCloneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddressCloneId

`func (o *OrderUpdateDataAttributes) SetBillingAddressCloneId(v string)`

SetBillingAddressCloneId sets BillingAddressCloneId field to given value.

### HasBillingAddressCloneId

`func (o *OrderUpdateDataAttributes) HasBillingAddressCloneId() bool`

HasBillingAddressCloneId returns a boolean if a field has been set.

### GetShippingAddressCloneId

`func (o *OrderUpdateDataAttributes) GetShippingAddressCloneId() string`

GetShippingAddressCloneId returns the ShippingAddressCloneId field if non-nil, zero value otherwise.

### GetShippingAddressCloneIdOk

`func (o *OrderUpdateDataAttributes) GetShippingAddressCloneIdOk() (*string, bool)`

GetShippingAddressCloneIdOk returns a tuple with the ShippingAddressCloneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddressCloneId

`func (o *OrderUpdateDataAttributes) SetShippingAddressCloneId(v string)`

SetShippingAddressCloneId sets ShippingAddressCloneId field to given value.

### HasShippingAddressCloneId

`func (o *OrderUpdateDataAttributes) HasShippingAddressCloneId() bool`

HasShippingAddressCloneId returns a boolean if a field has been set.

### GetCustomerPaymentSourceId

`func (o *OrderUpdateDataAttributes) GetCustomerPaymentSourceId() string`

GetCustomerPaymentSourceId returns the CustomerPaymentSourceId field if non-nil, zero value otherwise.

### GetCustomerPaymentSourceIdOk

`func (o *OrderUpdateDataAttributes) GetCustomerPaymentSourceIdOk() (*string, bool)`

GetCustomerPaymentSourceIdOk returns a tuple with the CustomerPaymentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPaymentSourceId

`func (o *OrderUpdateDataAttributes) SetCustomerPaymentSourceId(v string)`

SetCustomerPaymentSourceId sets CustomerPaymentSourceId field to given value.

### HasCustomerPaymentSourceId

`func (o *OrderUpdateDataAttributes) HasCustomerPaymentSourceId() bool`

HasCustomerPaymentSourceId returns a boolean if a field has been set.

### GetShippingAddressSameAsBilling

`func (o *OrderUpdateDataAttributes) GetShippingAddressSameAsBilling() bool`

GetShippingAddressSameAsBilling returns the ShippingAddressSameAsBilling field if non-nil, zero value otherwise.

### GetShippingAddressSameAsBillingOk

`func (o *OrderUpdateDataAttributes) GetShippingAddressSameAsBillingOk() (*bool, bool)`

GetShippingAddressSameAsBillingOk returns a tuple with the ShippingAddressSameAsBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddressSameAsBilling

`func (o *OrderUpdateDataAttributes) SetShippingAddressSameAsBilling(v bool)`

SetShippingAddressSameAsBilling sets ShippingAddressSameAsBilling field to given value.

### HasShippingAddressSameAsBilling

`func (o *OrderUpdateDataAttributes) HasShippingAddressSameAsBilling() bool`

HasShippingAddressSameAsBilling returns a boolean if a field has been set.

### GetBillingAddressSameAsShipping

`func (o *OrderUpdateDataAttributes) GetBillingAddressSameAsShipping() bool`

GetBillingAddressSameAsShipping returns the BillingAddressSameAsShipping field if non-nil, zero value otherwise.

### GetBillingAddressSameAsShippingOk

`func (o *OrderUpdateDataAttributes) GetBillingAddressSameAsShippingOk() (*bool, bool)`

GetBillingAddressSameAsShippingOk returns a tuple with the BillingAddressSameAsShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddressSameAsShipping

`func (o *OrderUpdateDataAttributes) SetBillingAddressSameAsShipping(v bool)`

SetBillingAddressSameAsShipping sets BillingAddressSameAsShipping field to given value.

### HasBillingAddressSameAsShipping

`func (o *OrderUpdateDataAttributes) HasBillingAddressSameAsShipping() bool`

HasBillingAddressSameAsShipping returns a boolean if a field has been set.

### GetCommitInvoice

`func (o *OrderUpdateDataAttributes) GetCommitInvoice() bool`

GetCommitInvoice returns the CommitInvoice field if non-nil, zero value otherwise.

### GetCommitInvoiceOk

`func (o *OrderUpdateDataAttributes) GetCommitInvoiceOk() (*bool, bool)`

GetCommitInvoiceOk returns a tuple with the CommitInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitInvoice

`func (o *OrderUpdateDataAttributes) SetCommitInvoice(v bool)`

SetCommitInvoice sets CommitInvoice field to given value.

### HasCommitInvoice

`func (o *OrderUpdateDataAttributes) HasCommitInvoice() bool`

HasCommitInvoice returns a boolean if a field has been set.

### GetRefundInvoice

`func (o *OrderUpdateDataAttributes) GetRefundInvoice() bool`

GetRefundInvoice returns the RefundInvoice field if non-nil, zero value otherwise.

### GetRefundInvoiceOk

`func (o *OrderUpdateDataAttributes) GetRefundInvoiceOk() (*bool, bool)`

GetRefundInvoiceOk returns a tuple with the RefundInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundInvoice

`func (o *OrderUpdateDataAttributes) SetRefundInvoice(v bool)`

SetRefundInvoice sets RefundInvoice field to given value.

### HasRefundInvoice

`func (o *OrderUpdateDataAttributes) HasRefundInvoice() bool`

HasRefundInvoice returns a boolean if a field has been set.

### GetSavePaymentSourceToCustomerWallet

`func (o *OrderUpdateDataAttributes) GetSavePaymentSourceToCustomerWallet() bool`

GetSavePaymentSourceToCustomerWallet returns the SavePaymentSourceToCustomerWallet field if non-nil, zero value otherwise.

### GetSavePaymentSourceToCustomerWalletOk

`func (o *OrderUpdateDataAttributes) GetSavePaymentSourceToCustomerWalletOk() (*bool, bool)`

GetSavePaymentSourceToCustomerWalletOk returns a tuple with the SavePaymentSourceToCustomerWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavePaymentSourceToCustomerWallet

`func (o *OrderUpdateDataAttributes) SetSavePaymentSourceToCustomerWallet(v bool)`

SetSavePaymentSourceToCustomerWallet sets SavePaymentSourceToCustomerWallet field to given value.

### HasSavePaymentSourceToCustomerWallet

`func (o *OrderUpdateDataAttributes) HasSavePaymentSourceToCustomerWallet() bool`

HasSavePaymentSourceToCustomerWallet returns a boolean if a field has been set.

### GetSaveShippingAddressToCustomerAddressBook

`func (o *OrderUpdateDataAttributes) GetSaveShippingAddressToCustomerAddressBook() bool`

GetSaveShippingAddressToCustomerAddressBook returns the SaveShippingAddressToCustomerAddressBook field if non-nil, zero value otherwise.

### GetSaveShippingAddressToCustomerAddressBookOk

`func (o *OrderUpdateDataAttributes) GetSaveShippingAddressToCustomerAddressBookOk() (*bool, bool)`

GetSaveShippingAddressToCustomerAddressBookOk returns a tuple with the SaveShippingAddressToCustomerAddressBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveShippingAddressToCustomerAddressBook

`func (o *OrderUpdateDataAttributes) SetSaveShippingAddressToCustomerAddressBook(v bool)`

SetSaveShippingAddressToCustomerAddressBook sets SaveShippingAddressToCustomerAddressBook field to given value.

### HasSaveShippingAddressToCustomerAddressBook

`func (o *OrderUpdateDataAttributes) HasSaveShippingAddressToCustomerAddressBook() bool`

HasSaveShippingAddressToCustomerAddressBook returns a boolean if a field has been set.

### GetSaveBillingAddressToCustomerAddressBook

`func (o *OrderUpdateDataAttributes) GetSaveBillingAddressToCustomerAddressBook() bool`

GetSaveBillingAddressToCustomerAddressBook returns the SaveBillingAddressToCustomerAddressBook field if non-nil, zero value otherwise.

### GetSaveBillingAddressToCustomerAddressBookOk

`func (o *OrderUpdateDataAttributes) GetSaveBillingAddressToCustomerAddressBookOk() (*bool, bool)`

GetSaveBillingAddressToCustomerAddressBookOk returns a tuple with the SaveBillingAddressToCustomerAddressBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveBillingAddressToCustomerAddressBook

`func (o *OrderUpdateDataAttributes) SetSaveBillingAddressToCustomerAddressBook(v bool)`

SetSaveBillingAddressToCustomerAddressBook sets SaveBillingAddressToCustomerAddressBook field to given value.

### HasSaveBillingAddressToCustomerAddressBook

`func (o *OrderUpdateDataAttributes) HasSaveBillingAddressToCustomerAddressBook() bool`

HasSaveBillingAddressToCustomerAddressBook returns a boolean if a field has been set.

### GetRefresh

`func (o *OrderUpdateDataAttributes) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *OrderUpdateDataAttributes) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *OrderUpdateDataAttributes) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *OrderUpdateDataAttributes) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetReference

`func (o *OrderUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *OrderUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *OrderUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *OrderUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *OrderUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *OrderUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *OrderUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *OrderUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *OrderUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


