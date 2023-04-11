# PATCHOrdersOrderId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autorefresh** | Pointer to **interface{}** | Save this attribute as &#39;false&#39; if you want prevent the order to be refreshed automatically at each change (much faster). | [optional] 
**Guest** | Pointer to **interface{}** | Indicates if the order has been placed as guest. | [optional] 
**CustomerEmail** | Pointer to **interface{}** | The email address of the associated customer. When creating or updating an order, this is a shortcut to find or create the associated customer by email. | [optional] 
**CustomerPassword** | Pointer to **interface{}** | The password of the associated customer. When creating or updating an order, this is a shortcut to sign up the associated customer. | [optional] 
**LanguageCode** | Pointer to **interface{}** | The preferred language code (ISO 639-1) to be used when communicating with the customer. This can be useful when sending the order to 3rd party marketing tools and CRMs. If the language is supported, the hosted checkout will be localized accordingly. | [optional] 
**ShippingCountryCodeLock** | Pointer to **interface{}** | The country code that you want the shipping address to be locked to. This can be useful to make sure the shipping address belongs to a given shipping country, e.g. the one selected in a country selector page. | [optional] 
**CouponCode** | Pointer to **interface{}** | The coupon code to be used for the order. If valid, it triggers a promotion adding a discount line item to the order. | [optional] 
**GiftCardCode** | Pointer to **interface{}** | The gift card code (at least the first 8 characters) to be used for the order. If valid, it uses the gift card balance to pay for the order. | [optional] 
**GiftCardOrCouponCode** | Pointer to **interface{}** | The gift card or coupon code (at least the first 8 characters) to be used for the order. If a gift card mathes, it uses the gift card balance to pay for the order. Otherwise it tries to find a valid coupon code and applies the associated discount. | [optional] 
**CartUrl** | Pointer to **interface{}** | The cart url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**ReturnUrl** | Pointer to **interface{}** | The return url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**TermsUrl** | Pointer to **interface{}** | The terms and conditions url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**PrivacyUrl** | Pointer to **interface{}** | The privacy policy url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**Archive** | Pointer to **interface{}** | Send this attribute if you want to archive the order. | [optional] 
**Unarchive** | Pointer to **interface{}** | Send this attribute if you want to unarchive the order. | [optional] 
**Place** | Pointer to **interface{}** | Send this attribute if you want to place the order. | [optional] 
**Cancel** | Pointer to **interface{}** | Send this attribute if you want to cancel a placed order. The order&#39;s authorization will be automatically voided. | [optional] 
**Approve** | Pointer to **interface{}** | Send this attribute if you want to approve a placed order. | [optional] 
**ApproveAndCapture** | Pointer to **interface{}** | Send this attribute if you want to approve and capture a placed order. | [optional] 
**Authorize** | Pointer to **interface{}** | Send this attribute if you want to authorize the order&#39;s payment source. | [optional] 
**AuthorizationAmountCents** | Pointer to **interface{}** | The authorization amount, in cents. | [optional] 
**Capture** | Pointer to **interface{}** | Send this attribute if you want to capture an authorized order. | [optional] 
**Refund** | Pointer to **interface{}** | Send this attribute if you want to refund a captured order. | [optional] 
**UpdateTaxes** | Pointer to **interface{}** | Send this attribute if you want to force tax calculation for this order (a tax calculator must be associated to the order&#39;s market). | [optional] 
**NullifyPaymentSource** | Pointer to **interface{}** | Send this attribute if you want to nullify the payment source for this order. | [optional] 
**BillingAddressCloneId** | Pointer to **interface{}** | The id of the address that you want to clone to create the order&#39;s billing address. | [optional] 
**ShippingAddressCloneId** | Pointer to **interface{}** | The id of the address that you want to clone to create the order&#39;s shipping address. | [optional] 
**CustomerPaymentSourceId** | Pointer to **interface{}** | The id of the customer payment source (i.e. credit card) that you want to use as the order&#39;s payment source. | [optional] 
**ShippingAddressSameAsBilling** | Pointer to **interface{}** | Send this attribute if you want the shipping address to be cloned from the order&#39;s billing address. | [optional] 
**BillingAddressSameAsShipping** | Pointer to **interface{}** | Send this attribute if you want the billing address to be cloned from the order&#39;s shipping address. | [optional] 
**CommitInvoice** | Pointer to **interface{}** | Send this attribute if you want commit the sales tax invoice to the associated tax calculator (currently supported by Avalara). | [optional] 
**RefundInvoice** | Pointer to **interface{}** | Send this attribute if you want refund the sales tax invoice to the associated tax calculator (currently supported by Avalara). | [optional] 
**SavePaymentSourceToCustomerWallet** | Pointer to **interface{}** | Send this attribute if you want the order&#39;s payment source to be saved in the customer&#39;s wallet as a customer payment source. | [optional] 
**SaveShippingAddressToCustomerAddressBook** | Pointer to **interface{}** | Send this attribute if you want the order&#39;s shipping address to be saved in the customer&#39;s address book as a customer address. | [optional] 
**SaveBillingAddressToCustomerAddressBook** | Pointer to **interface{}** | Send this attribute if you want the order&#39;s billing address to be saved in the customer&#39;s address book as a customer address. | [optional] 
**Refresh** | Pointer to **interface{}** | Send this attribute if you want to manually refresh the order. | [optional] 
**Validate** | Pointer to **interface{}** | Send this attribute if you want to trigger the external validation for the order. | [optional] 
**CreateSubscriptions** | Pointer to **interface{}** | Send this attribute if you want to create order subscriptions from the recurring line items upon/after placing the order. Subscriptions are generated according to associated subscription model strategy. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHOrdersOrderId200ResponseDataAttributes

`func NewPATCHOrdersOrderId200ResponseDataAttributes() *PATCHOrdersOrderId200ResponseDataAttributes`

NewPATCHOrdersOrderId200ResponseDataAttributes instantiates a new PATCHOrdersOrderId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHOrdersOrderId200ResponseDataAttributesWithDefaults

`func NewPATCHOrdersOrderId200ResponseDataAttributesWithDefaults() *PATCHOrdersOrderId200ResponseDataAttributes`

NewPATCHOrdersOrderId200ResponseDataAttributesWithDefaults instantiates a new PATCHOrdersOrderId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutorefresh

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetAutorefresh() interface{}`

GetAutorefresh returns the Autorefresh field if non-nil, zero value otherwise.

### GetAutorefreshOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetAutorefreshOk() (*interface{}, bool)`

GetAutorefreshOk returns a tuple with the Autorefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutorefresh

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetAutorefresh(v interface{})`

SetAutorefresh sets Autorefresh field to given value.

### HasAutorefresh

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasAutorefresh() bool`

HasAutorefresh returns a boolean if a field has been set.

### SetAutorefreshNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetAutorefreshNil(b bool)`

 SetAutorefreshNil sets the value for Autorefresh to be an explicit nil

### UnsetAutorefresh
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetAutorefresh()`

UnsetAutorefresh ensures that no value is present for Autorefresh, not even an explicit nil
### GetGuest

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetGuest() interface{}`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetGuestOk() (*interface{}, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetGuest(v interface{})`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### SetGuestNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetGuestNil(b bool)`

 SetGuestNil sets the value for Guest to be an explicit nil

### UnsetGuest
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetGuest()`

UnsetGuest ensures that no value is present for Guest, not even an explicit nil
### GetCustomerEmail

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCustomerEmail() interface{}`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCustomerEmailOk() (*interface{}, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCustomerEmail(v interface{})`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### SetCustomerEmailNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCustomerEmailNil(b bool)`

 SetCustomerEmailNil sets the value for CustomerEmail to be an explicit nil

### UnsetCustomerEmail
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetCustomerEmail()`

UnsetCustomerEmail ensures that no value is present for CustomerEmail, not even an explicit nil
### GetCustomerPassword

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCustomerPassword() interface{}`

GetCustomerPassword returns the CustomerPassword field if non-nil, zero value otherwise.

### GetCustomerPasswordOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCustomerPasswordOk() (*interface{}, bool)`

GetCustomerPasswordOk returns a tuple with the CustomerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPassword

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCustomerPassword(v interface{})`

SetCustomerPassword sets CustomerPassword field to given value.

### HasCustomerPassword

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasCustomerPassword() bool`

HasCustomerPassword returns a boolean if a field has been set.

### SetCustomerPasswordNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCustomerPasswordNil(b bool)`

 SetCustomerPasswordNil sets the value for CustomerPassword to be an explicit nil

### UnsetCustomerPassword
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetCustomerPassword()`

UnsetCustomerPassword ensures that no value is present for CustomerPassword, not even an explicit nil
### GetLanguageCode

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetLanguageCode() interface{}`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetLanguageCodeOk() (*interface{}, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetLanguageCode(v interface{})`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### SetLanguageCodeNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetLanguageCodeNil(b bool)`

 SetLanguageCodeNil sets the value for LanguageCode to be an explicit nil

### UnsetLanguageCode
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetLanguageCode()`

UnsetLanguageCode ensures that no value is present for LanguageCode, not even an explicit nil
### GetShippingCountryCodeLock

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetShippingCountryCodeLock() interface{}`

GetShippingCountryCodeLock returns the ShippingCountryCodeLock field if non-nil, zero value otherwise.

### GetShippingCountryCodeLockOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetShippingCountryCodeLockOk() (*interface{}, bool)`

GetShippingCountryCodeLockOk returns a tuple with the ShippingCountryCodeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCountryCodeLock

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetShippingCountryCodeLock(v interface{})`

SetShippingCountryCodeLock sets ShippingCountryCodeLock field to given value.

### HasShippingCountryCodeLock

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasShippingCountryCodeLock() bool`

HasShippingCountryCodeLock returns a boolean if a field has been set.

### SetShippingCountryCodeLockNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetShippingCountryCodeLockNil(b bool)`

 SetShippingCountryCodeLockNil sets the value for ShippingCountryCodeLock to be an explicit nil

### UnsetShippingCountryCodeLock
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetShippingCountryCodeLock()`

UnsetShippingCountryCodeLock ensures that no value is present for ShippingCountryCodeLock, not even an explicit nil
### GetCouponCode

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCouponCode() interface{}`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCouponCodeOk() (*interface{}, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCouponCode(v interface{})`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.

### SetCouponCodeNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCouponCodeNil(b bool)`

 SetCouponCodeNil sets the value for CouponCode to be an explicit nil

### UnsetCouponCode
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetCouponCode()`

UnsetCouponCode ensures that no value is present for CouponCode, not even an explicit nil
### GetGiftCardCode

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetGiftCardCode() interface{}`

GetGiftCardCode returns the GiftCardCode field if non-nil, zero value otherwise.

### GetGiftCardCodeOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetGiftCardCodeOk() (*interface{}, bool)`

GetGiftCardCodeOk returns a tuple with the GiftCardCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardCode

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetGiftCardCode(v interface{})`

SetGiftCardCode sets GiftCardCode field to given value.

### HasGiftCardCode

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasGiftCardCode() bool`

HasGiftCardCode returns a boolean if a field has been set.

### SetGiftCardCodeNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetGiftCardCodeNil(b bool)`

 SetGiftCardCodeNil sets the value for GiftCardCode to be an explicit nil

### UnsetGiftCardCode
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetGiftCardCode()`

UnsetGiftCardCode ensures that no value is present for GiftCardCode, not even an explicit nil
### GetGiftCardOrCouponCode

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetGiftCardOrCouponCode() interface{}`

GetGiftCardOrCouponCode returns the GiftCardOrCouponCode field if non-nil, zero value otherwise.

### GetGiftCardOrCouponCodeOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetGiftCardOrCouponCodeOk() (*interface{}, bool)`

GetGiftCardOrCouponCodeOk returns a tuple with the GiftCardOrCouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardOrCouponCode

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetGiftCardOrCouponCode(v interface{})`

SetGiftCardOrCouponCode sets GiftCardOrCouponCode field to given value.

### HasGiftCardOrCouponCode

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasGiftCardOrCouponCode() bool`

HasGiftCardOrCouponCode returns a boolean if a field has been set.

### SetGiftCardOrCouponCodeNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetGiftCardOrCouponCodeNil(b bool)`

 SetGiftCardOrCouponCodeNil sets the value for GiftCardOrCouponCode to be an explicit nil

### UnsetGiftCardOrCouponCode
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetGiftCardOrCouponCode()`

UnsetGiftCardOrCouponCode ensures that no value is present for GiftCardOrCouponCode, not even an explicit nil
### GetCartUrl

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCartUrl() interface{}`

GetCartUrl returns the CartUrl field if non-nil, zero value otherwise.

### GetCartUrlOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCartUrlOk() (*interface{}, bool)`

GetCartUrlOk returns a tuple with the CartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartUrl

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCartUrl(v interface{})`

SetCartUrl sets CartUrl field to given value.

### HasCartUrl

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasCartUrl() bool`

HasCartUrl returns a boolean if a field has been set.

### SetCartUrlNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCartUrlNil(b bool)`

 SetCartUrlNil sets the value for CartUrl to be an explicit nil

### UnsetCartUrl
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetCartUrl()`

UnsetCartUrl ensures that no value is present for CartUrl, not even an explicit nil
### GetReturnUrl

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetReturnUrl() interface{}`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetReturnUrlOk() (*interface{}, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetReturnUrl(v interface{})`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### SetReturnUrlNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetTermsUrl

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetTermsUrl() interface{}`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetTermsUrlOk() (*interface{}, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetTermsUrl(v interface{})`

SetTermsUrl sets TermsUrl field to given value.

### HasTermsUrl

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasTermsUrl() bool`

HasTermsUrl returns a boolean if a field has been set.

### SetTermsUrlNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetTermsUrlNil(b bool)`

 SetTermsUrlNil sets the value for TermsUrl to be an explicit nil

### UnsetTermsUrl
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetTermsUrl()`

UnsetTermsUrl ensures that no value is present for TermsUrl, not even an explicit nil
### GetPrivacyUrl

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetPrivacyUrl() interface{}`

GetPrivacyUrl returns the PrivacyUrl field if non-nil, zero value otherwise.

### GetPrivacyUrlOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetPrivacyUrlOk() (*interface{}, bool)`

GetPrivacyUrlOk returns a tuple with the PrivacyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyUrl

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetPrivacyUrl(v interface{})`

SetPrivacyUrl sets PrivacyUrl field to given value.

### HasPrivacyUrl

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasPrivacyUrl() bool`

HasPrivacyUrl returns a boolean if a field has been set.

### SetPrivacyUrlNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetPrivacyUrlNil(b bool)`

 SetPrivacyUrlNil sets the value for PrivacyUrl to be an explicit nil

### UnsetPrivacyUrl
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetPrivacyUrl()`

UnsetPrivacyUrl ensures that no value is present for PrivacyUrl, not even an explicit nil
### GetArchive

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetArchive() interface{}`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetArchiveOk() (*interface{}, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetArchive(v interface{})`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### SetArchiveNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetArchiveNil(b bool)`

 SetArchiveNil sets the value for Archive to be an explicit nil

### UnsetArchive
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetArchive()`

UnsetArchive ensures that no value is present for Archive, not even an explicit nil
### GetUnarchive

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetUnarchive() interface{}`

GetUnarchive returns the Unarchive field if non-nil, zero value otherwise.

### GetUnarchiveOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetUnarchiveOk() (*interface{}, bool)`

GetUnarchiveOk returns a tuple with the Unarchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnarchive

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetUnarchive(v interface{})`

SetUnarchive sets Unarchive field to given value.

### HasUnarchive

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasUnarchive() bool`

HasUnarchive returns a boolean if a field has been set.

### SetUnarchiveNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetUnarchiveNil(b bool)`

 SetUnarchiveNil sets the value for Unarchive to be an explicit nil

### UnsetUnarchive
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetUnarchive()`

UnsetUnarchive ensures that no value is present for Unarchive, not even an explicit nil
### GetPlace

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetPlace() interface{}`

GetPlace returns the Place field if non-nil, zero value otherwise.

### GetPlaceOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetPlaceOk() (*interface{}, bool)`

GetPlaceOk returns a tuple with the Place field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlace

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetPlace(v interface{})`

SetPlace sets Place field to given value.

### HasPlace

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasPlace() bool`

HasPlace returns a boolean if a field has been set.

### SetPlaceNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetPlaceNil(b bool)`

 SetPlaceNil sets the value for Place to be an explicit nil

### UnsetPlace
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetPlace()`

UnsetPlace ensures that no value is present for Place, not even an explicit nil
### GetCancel

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCancel() interface{}`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCancelOk() (*interface{}, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCancel(v interface{})`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### SetCancelNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCancelNil(b bool)`

 SetCancelNil sets the value for Cancel to be an explicit nil

### UnsetCancel
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetCancel()`

UnsetCancel ensures that no value is present for Cancel, not even an explicit nil
### GetApprove

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetApprove() interface{}`

GetApprove returns the Approve field if non-nil, zero value otherwise.

### GetApproveOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetApproveOk() (*interface{}, bool)`

GetApproveOk returns a tuple with the Approve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprove

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetApprove(v interface{})`

SetApprove sets Approve field to given value.

### HasApprove

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasApprove() bool`

HasApprove returns a boolean if a field has been set.

### SetApproveNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetApproveNil(b bool)`

 SetApproveNil sets the value for Approve to be an explicit nil

### UnsetApprove
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetApprove()`

UnsetApprove ensures that no value is present for Approve, not even an explicit nil
### GetApproveAndCapture

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetApproveAndCapture() interface{}`

GetApproveAndCapture returns the ApproveAndCapture field if non-nil, zero value otherwise.

### GetApproveAndCaptureOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetApproveAndCaptureOk() (*interface{}, bool)`

GetApproveAndCaptureOk returns a tuple with the ApproveAndCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveAndCapture

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetApproveAndCapture(v interface{})`

SetApproveAndCapture sets ApproveAndCapture field to given value.

### HasApproveAndCapture

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasApproveAndCapture() bool`

HasApproveAndCapture returns a boolean if a field has been set.

### SetApproveAndCaptureNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetApproveAndCaptureNil(b bool)`

 SetApproveAndCaptureNil sets the value for ApproveAndCapture to be an explicit nil

### UnsetApproveAndCapture
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetApproveAndCapture()`

UnsetApproveAndCapture ensures that no value is present for ApproveAndCapture, not even an explicit nil
### GetAuthorize

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetAuthorize() interface{}`

GetAuthorize returns the Authorize field if non-nil, zero value otherwise.

### GetAuthorizeOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetAuthorizeOk() (*interface{}, bool)`

GetAuthorizeOk returns a tuple with the Authorize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorize

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetAuthorize(v interface{})`

SetAuthorize sets Authorize field to given value.

### HasAuthorize

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasAuthorize() bool`

HasAuthorize returns a boolean if a field has been set.

### SetAuthorizeNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetAuthorizeNil(b bool)`

 SetAuthorizeNil sets the value for Authorize to be an explicit nil

### UnsetAuthorize
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetAuthorize()`

UnsetAuthorize ensures that no value is present for Authorize, not even an explicit nil
### GetAuthorizationAmountCents

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetAuthorizationAmountCents() interface{}`

GetAuthorizationAmountCents returns the AuthorizationAmountCents field if non-nil, zero value otherwise.

### GetAuthorizationAmountCentsOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetAuthorizationAmountCentsOk() (*interface{}, bool)`

GetAuthorizationAmountCentsOk returns a tuple with the AuthorizationAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationAmountCents

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetAuthorizationAmountCents(v interface{})`

SetAuthorizationAmountCents sets AuthorizationAmountCents field to given value.

### HasAuthorizationAmountCents

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasAuthorizationAmountCents() bool`

HasAuthorizationAmountCents returns a boolean if a field has been set.

### SetAuthorizationAmountCentsNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetAuthorizationAmountCentsNil(b bool)`

 SetAuthorizationAmountCentsNil sets the value for AuthorizationAmountCents to be an explicit nil

### UnsetAuthorizationAmountCents
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetAuthorizationAmountCents()`

UnsetAuthorizationAmountCents ensures that no value is present for AuthorizationAmountCents, not even an explicit nil
### GetCapture

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCapture() interface{}`

GetCapture returns the Capture field if non-nil, zero value otherwise.

### GetCaptureOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCaptureOk() (*interface{}, bool)`

GetCaptureOk returns a tuple with the Capture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapture

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCapture(v interface{})`

SetCapture sets Capture field to given value.

### HasCapture

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasCapture() bool`

HasCapture returns a boolean if a field has been set.

### SetCaptureNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCaptureNil(b bool)`

 SetCaptureNil sets the value for Capture to be an explicit nil

### UnsetCapture
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetCapture()`

UnsetCapture ensures that no value is present for Capture, not even an explicit nil
### GetRefund

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetRefund() interface{}`

GetRefund returns the Refund field if non-nil, zero value otherwise.

### GetRefundOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetRefundOk() (*interface{}, bool)`

GetRefundOk returns a tuple with the Refund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefund

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetRefund(v interface{})`

SetRefund sets Refund field to given value.

### HasRefund

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasRefund() bool`

HasRefund returns a boolean if a field has been set.

### SetRefundNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetRefundNil(b bool)`

 SetRefundNil sets the value for Refund to be an explicit nil

### UnsetRefund
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetRefund()`

UnsetRefund ensures that no value is present for Refund, not even an explicit nil
### GetUpdateTaxes

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetUpdateTaxes() interface{}`

GetUpdateTaxes returns the UpdateTaxes field if non-nil, zero value otherwise.

### GetUpdateTaxesOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetUpdateTaxesOk() (*interface{}, bool)`

GetUpdateTaxesOk returns a tuple with the UpdateTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTaxes

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetUpdateTaxes(v interface{})`

SetUpdateTaxes sets UpdateTaxes field to given value.

### HasUpdateTaxes

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasUpdateTaxes() bool`

HasUpdateTaxes returns a boolean if a field has been set.

### SetUpdateTaxesNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetUpdateTaxesNil(b bool)`

 SetUpdateTaxesNil sets the value for UpdateTaxes to be an explicit nil

### UnsetUpdateTaxes
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetUpdateTaxes()`

UnsetUpdateTaxes ensures that no value is present for UpdateTaxes, not even an explicit nil
### GetNullifyPaymentSource

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetNullifyPaymentSource() interface{}`

GetNullifyPaymentSource returns the NullifyPaymentSource field if non-nil, zero value otherwise.

### GetNullifyPaymentSourceOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetNullifyPaymentSourceOk() (*interface{}, bool)`

GetNullifyPaymentSourceOk returns a tuple with the NullifyPaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullifyPaymentSource

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetNullifyPaymentSource(v interface{})`

SetNullifyPaymentSource sets NullifyPaymentSource field to given value.

### HasNullifyPaymentSource

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasNullifyPaymentSource() bool`

HasNullifyPaymentSource returns a boolean if a field has been set.

### SetNullifyPaymentSourceNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetNullifyPaymentSourceNil(b bool)`

 SetNullifyPaymentSourceNil sets the value for NullifyPaymentSource to be an explicit nil

### UnsetNullifyPaymentSource
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetNullifyPaymentSource()`

UnsetNullifyPaymentSource ensures that no value is present for NullifyPaymentSource, not even an explicit nil
### GetBillingAddressCloneId

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetBillingAddressCloneId() interface{}`

GetBillingAddressCloneId returns the BillingAddressCloneId field if non-nil, zero value otherwise.

### GetBillingAddressCloneIdOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetBillingAddressCloneIdOk() (*interface{}, bool)`

GetBillingAddressCloneIdOk returns a tuple with the BillingAddressCloneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddressCloneId

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetBillingAddressCloneId(v interface{})`

SetBillingAddressCloneId sets BillingAddressCloneId field to given value.

### HasBillingAddressCloneId

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasBillingAddressCloneId() bool`

HasBillingAddressCloneId returns a boolean if a field has been set.

### SetBillingAddressCloneIdNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetBillingAddressCloneIdNil(b bool)`

 SetBillingAddressCloneIdNil sets the value for BillingAddressCloneId to be an explicit nil

### UnsetBillingAddressCloneId
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetBillingAddressCloneId()`

UnsetBillingAddressCloneId ensures that no value is present for BillingAddressCloneId, not even an explicit nil
### GetShippingAddressCloneId

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetShippingAddressCloneId() interface{}`

GetShippingAddressCloneId returns the ShippingAddressCloneId field if non-nil, zero value otherwise.

### GetShippingAddressCloneIdOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetShippingAddressCloneIdOk() (*interface{}, bool)`

GetShippingAddressCloneIdOk returns a tuple with the ShippingAddressCloneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddressCloneId

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetShippingAddressCloneId(v interface{})`

SetShippingAddressCloneId sets ShippingAddressCloneId field to given value.

### HasShippingAddressCloneId

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasShippingAddressCloneId() bool`

HasShippingAddressCloneId returns a boolean if a field has been set.

### SetShippingAddressCloneIdNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetShippingAddressCloneIdNil(b bool)`

 SetShippingAddressCloneIdNil sets the value for ShippingAddressCloneId to be an explicit nil

### UnsetShippingAddressCloneId
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetShippingAddressCloneId()`

UnsetShippingAddressCloneId ensures that no value is present for ShippingAddressCloneId, not even an explicit nil
### GetCustomerPaymentSourceId

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCustomerPaymentSourceId() interface{}`

GetCustomerPaymentSourceId returns the CustomerPaymentSourceId field if non-nil, zero value otherwise.

### GetCustomerPaymentSourceIdOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCustomerPaymentSourceIdOk() (*interface{}, bool)`

GetCustomerPaymentSourceIdOk returns a tuple with the CustomerPaymentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPaymentSourceId

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCustomerPaymentSourceId(v interface{})`

SetCustomerPaymentSourceId sets CustomerPaymentSourceId field to given value.

### HasCustomerPaymentSourceId

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasCustomerPaymentSourceId() bool`

HasCustomerPaymentSourceId returns a boolean if a field has been set.

### SetCustomerPaymentSourceIdNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCustomerPaymentSourceIdNil(b bool)`

 SetCustomerPaymentSourceIdNil sets the value for CustomerPaymentSourceId to be an explicit nil

### UnsetCustomerPaymentSourceId
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetCustomerPaymentSourceId()`

UnsetCustomerPaymentSourceId ensures that no value is present for CustomerPaymentSourceId, not even an explicit nil
### GetShippingAddressSameAsBilling

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetShippingAddressSameAsBilling() interface{}`

GetShippingAddressSameAsBilling returns the ShippingAddressSameAsBilling field if non-nil, zero value otherwise.

### GetShippingAddressSameAsBillingOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetShippingAddressSameAsBillingOk() (*interface{}, bool)`

GetShippingAddressSameAsBillingOk returns a tuple with the ShippingAddressSameAsBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddressSameAsBilling

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetShippingAddressSameAsBilling(v interface{})`

SetShippingAddressSameAsBilling sets ShippingAddressSameAsBilling field to given value.

### HasShippingAddressSameAsBilling

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasShippingAddressSameAsBilling() bool`

HasShippingAddressSameAsBilling returns a boolean if a field has been set.

### SetShippingAddressSameAsBillingNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetShippingAddressSameAsBillingNil(b bool)`

 SetShippingAddressSameAsBillingNil sets the value for ShippingAddressSameAsBilling to be an explicit nil

### UnsetShippingAddressSameAsBilling
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetShippingAddressSameAsBilling()`

UnsetShippingAddressSameAsBilling ensures that no value is present for ShippingAddressSameAsBilling, not even an explicit nil
### GetBillingAddressSameAsShipping

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetBillingAddressSameAsShipping() interface{}`

GetBillingAddressSameAsShipping returns the BillingAddressSameAsShipping field if non-nil, zero value otherwise.

### GetBillingAddressSameAsShippingOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetBillingAddressSameAsShippingOk() (*interface{}, bool)`

GetBillingAddressSameAsShippingOk returns a tuple with the BillingAddressSameAsShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddressSameAsShipping

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetBillingAddressSameAsShipping(v interface{})`

SetBillingAddressSameAsShipping sets BillingAddressSameAsShipping field to given value.

### HasBillingAddressSameAsShipping

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasBillingAddressSameAsShipping() bool`

HasBillingAddressSameAsShipping returns a boolean if a field has been set.

### SetBillingAddressSameAsShippingNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetBillingAddressSameAsShippingNil(b bool)`

 SetBillingAddressSameAsShippingNil sets the value for BillingAddressSameAsShipping to be an explicit nil

### UnsetBillingAddressSameAsShipping
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetBillingAddressSameAsShipping()`

UnsetBillingAddressSameAsShipping ensures that no value is present for BillingAddressSameAsShipping, not even an explicit nil
### GetCommitInvoice

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCommitInvoice() interface{}`

GetCommitInvoice returns the CommitInvoice field if non-nil, zero value otherwise.

### GetCommitInvoiceOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCommitInvoiceOk() (*interface{}, bool)`

GetCommitInvoiceOk returns a tuple with the CommitInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitInvoice

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCommitInvoice(v interface{})`

SetCommitInvoice sets CommitInvoice field to given value.

### HasCommitInvoice

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasCommitInvoice() bool`

HasCommitInvoice returns a boolean if a field has been set.

### SetCommitInvoiceNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCommitInvoiceNil(b bool)`

 SetCommitInvoiceNil sets the value for CommitInvoice to be an explicit nil

### UnsetCommitInvoice
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetCommitInvoice()`

UnsetCommitInvoice ensures that no value is present for CommitInvoice, not even an explicit nil
### GetRefundInvoice

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetRefundInvoice() interface{}`

GetRefundInvoice returns the RefundInvoice field if non-nil, zero value otherwise.

### GetRefundInvoiceOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetRefundInvoiceOk() (*interface{}, bool)`

GetRefundInvoiceOk returns a tuple with the RefundInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundInvoice

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetRefundInvoice(v interface{})`

SetRefundInvoice sets RefundInvoice field to given value.

### HasRefundInvoice

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasRefundInvoice() bool`

HasRefundInvoice returns a boolean if a field has been set.

### SetRefundInvoiceNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetRefundInvoiceNil(b bool)`

 SetRefundInvoiceNil sets the value for RefundInvoice to be an explicit nil

### UnsetRefundInvoice
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetRefundInvoice()`

UnsetRefundInvoice ensures that no value is present for RefundInvoice, not even an explicit nil
### GetSavePaymentSourceToCustomerWallet

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetSavePaymentSourceToCustomerWallet() interface{}`

GetSavePaymentSourceToCustomerWallet returns the SavePaymentSourceToCustomerWallet field if non-nil, zero value otherwise.

### GetSavePaymentSourceToCustomerWalletOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetSavePaymentSourceToCustomerWalletOk() (*interface{}, bool)`

GetSavePaymentSourceToCustomerWalletOk returns a tuple with the SavePaymentSourceToCustomerWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavePaymentSourceToCustomerWallet

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetSavePaymentSourceToCustomerWallet(v interface{})`

SetSavePaymentSourceToCustomerWallet sets SavePaymentSourceToCustomerWallet field to given value.

### HasSavePaymentSourceToCustomerWallet

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasSavePaymentSourceToCustomerWallet() bool`

HasSavePaymentSourceToCustomerWallet returns a boolean if a field has been set.

### SetSavePaymentSourceToCustomerWalletNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetSavePaymentSourceToCustomerWalletNil(b bool)`

 SetSavePaymentSourceToCustomerWalletNil sets the value for SavePaymentSourceToCustomerWallet to be an explicit nil

### UnsetSavePaymentSourceToCustomerWallet
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetSavePaymentSourceToCustomerWallet()`

UnsetSavePaymentSourceToCustomerWallet ensures that no value is present for SavePaymentSourceToCustomerWallet, not even an explicit nil
### GetSaveShippingAddressToCustomerAddressBook

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetSaveShippingAddressToCustomerAddressBook() interface{}`

GetSaveShippingAddressToCustomerAddressBook returns the SaveShippingAddressToCustomerAddressBook field if non-nil, zero value otherwise.

### GetSaveShippingAddressToCustomerAddressBookOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetSaveShippingAddressToCustomerAddressBookOk() (*interface{}, bool)`

GetSaveShippingAddressToCustomerAddressBookOk returns a tuple with the SaveShippingAddressToCustomerAddressBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveShippingAddressToCustomerAddressBook

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetSaveShippingAddressToCustomerAddressBook(v interface{})`

SetSaveShippingAddressToCustomerAddressBook sets SaveShippingAddressToCustomerAddressBook field to given value.

### HasSaveShippingAddressToCustomerAddressBook

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasSaveShippingAddressToCustomerAddressBook() bool`

HasSaveShippingAddressToCustomerAddressBook returns a boolean if a field has been set.

### SetSaveShippingAddressToCustomerAddressBookNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetSaveShippingAddressToCustomerAddressBookNil(b bool)`

 SetSaveShippingAddressToCustomerAddressBookNil sets the value for SaveShippingAddressToCustomerAddressBook to be an explicit nil

### UnsetSaveShippingAddressToCustomerAddressBook
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetSaveShippingAddressToCustomerAddressBook()`

UnsetSaveShippingAddressToCustomerAddressBook ensures that no value is present for SaveShippingAddressToCustomerAddressBook, not even an explicit nil
### GetSaveBillingAddressToCustomerAddressBook

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetSaveBillingAddressToCustomerAddressBook() interface{}`

GetSaveBillingAddressToCustomerAddressBook returns the SaveBillingAddressToCustomerAddressBook field if non-nil, zero value otherwise.

### GetSaveBillingAddressToCustomerAddressBookOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetSaveBillingAddressToCustomerAddressBookOk() (*interface{}, bool)`

GetSaveBillingAddressToCustomerAddressBookOk returns a tuple with the SaveBillingAddressToCustomerAddressBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveBillingAddressToCustomerAddressBook

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetSaveBillingAddressToCustomerAddressBook(v interface{})`

SetSaveBillingAddressToCustomerAddressBook sets SaveBillingAddressToCustomerAddressBook field to given value.

### HasSaveBillingAddressToCustomerAddressBook

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasSaveBillingAddressToCustomerAddressBook() bool`

HasSaveBillingAddressToCustomerAddressBook returns a boolean if a field has been set.

### SetSaveBillingAddressToCustomerAddressBookNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetSaveBillingAddressToCustomerAddressBookNil(b bool)`

 SetSaveBillingAddressToCustomerAddressBookNil sets the value for SaveBillingAddressToCustomerAddressBook to be an explicit nil

### UnsetSaveBillingAddressToCustomerAddressBook
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetSaveBillingAddressToCustomerAddressBook()`

UnsetSaveBillingAddressToCustomerAddressBook ensures that no value is present for SaveBillingAddressToCustomerAddressBook, not even an explicit nil
### GetRefresh

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetRefresh() interface{}`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetRefreshOk() (*interface{}, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetRefresh(v interface{})`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### SetRefreshNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetRefreshNil(b bool)`

 SetRefreshNil sets the value for Refresh to be an explicit nil

### UnsetRefresh
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetRefresh()`

UnsetRefresh ensures that no value is present for Refresh, not even an explicit nil
### GetValidate

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetValidate() interface{}`

GetValidate returns the Validate field if non-nil, zero value otherwise.

### GetValidateOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetValidateOk() (*interface{}, bool)`

GetValidateOk returns a tuple with the Validate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidate

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetValidate(v interface{})`

SetValidate sets Validate field to given value.

### HasValidate

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasValidate() bool`

HasValidate returns a boolean if a field has been set.

### SetValidateNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetValidateNil(b bool)`

 SetValidateNil sets the value for Validate to be an explicit nil

### UnsetValidate
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetValidate()`

UnsetValidate ensures that no value is present for Validate, not even an explicit nil
### GetCreateSubscriptions

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCreateSubscriptions() interface{}`

GetCreateSubscriptions returns the CreateSubscriptions field if non-nil, zero value otherwise.

### GetCreateSubscriptionsOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetCreateSubscriptionsOk() (*interface{}, bool)`

GetCreateSubscriptionsOk returns a tuple with the CreateSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSubscriptions

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCreateSubscriptions(v interface{})`

SetCreateSubscriptions sets CreateSubscriptions field to given value.

### HasCreateSubscriptions

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasCreateSubscriptions() bool`

HasCreateSubscriptions returns a boolean if a field has been set.

### SetCreateSubscriptionsNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetCreateSubscriptionsNil(b bool)`

 SetCreateSubscriptionsNil sets the value for CreateSubscriptions to be an explicit nil

### UnsetCreateSubscriptions
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetCreateSubscriptions()`

UnsetCreateSubscriptions ensures that no value is present for CreateSubscriptions, not even an explicit nil
### GetReference

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHOrdersOrderId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHOrdersOrderId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


