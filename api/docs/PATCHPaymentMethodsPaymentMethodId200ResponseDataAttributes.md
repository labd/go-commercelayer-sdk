# PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentSourceType** | Pointer to **string** | The payment source type, can be one of: &#39;AdyenPayment&#39;, &#39;BraintreePayment&#39;, &#39;CheckoutComPayment&#39;, &#39;CreditCard&#39;, &#39;ExternalPayment&#39;, &#39;KlarnaPayment&#39;, &#39;PaypalPayment&#39;, &#39;StripePayment&#39;, or &#39;WireTransfer&#39;. | [optional] 
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Moto** | Pointer to **bool** | Send this attribute if you want to mark the payment as MOTO, must be supported by payment gateway. | [optional] 
**Disable** | Pointer to **bool** | Send this attribute if you want to mark the payment method as disabled. | [optional] 
**Enable** | Pointer to **bool** | Send this attribute if you want to mark the payment method as enabled. | [optional] 
**PriceAmountCents** | Pointer to **int32** | The payment method&#39;s price, in cents | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes

`func NewPATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes() *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes`

NewPATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes instantiates a new PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHPaymentMethodsPaymentMethodId200ResponseDataAttributesWithDefaults

`func NewPATCHPaymentMethodsPaymentMethodId200ResponseDataAttributesWithDefaults() *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes`

NewPATCHPaymentMethodsPaymentMethodId200ResponseDataAttributesWithDefaults instantiates a new PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentSourceType

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetPaymentSourceType() string`

GetPaymentSourceType returns the PaymentSourceType field if non-nil, zero value otherwise.

### GetPaymentSourceTypeOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetPaymentSourceTypeOk() (*string, bool)`

GetPaymentSourceTypeOk returns a tuple with the PaymentSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceType

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetPaymentSourceType(v string)`

SetPaymentSourceType sets PaymentSourceType field to given value.

### HasPaymentSourceType

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasPaymentSourceType() bool`

HasPaymentSourceType returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetMoto

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetMoto() bool`

GetMoto returns the Moto field if non-nil, zero value otherwise.

### GetMotoOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetMotoOk() (*bool, bool)`

GetMotoOk returns a tuple with the Moto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoto

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetMoto(v bool)`

SetMoto sets Moto field to given value.

### HasMoto

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasMoto() bool`

HasMoto returns a boolean if a field has been set.

### GetDisable

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEnable

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetPriceAmountCents

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetPriceAmountCents() int32`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetPriceAmountCentsOk() (*int32, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetPriceAmountCents(v int32)`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### GetReference

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


