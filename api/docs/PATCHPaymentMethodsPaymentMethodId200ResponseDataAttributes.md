# PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentSourceType** | Pointer to **interface{}** | The payment source type, can be one of: &#39;AdyenPayment&#39;, &#39;AxervePayment&#39;, &#39;BraintreePayment&#39;, &#39;CheckoutComPayment&#39;, &#39;CreditCard&#39;, &#39;ExternalPayment&#39;, &#39;KlarnaPayment&#39;, &#39;PaypalPayment&#39;, &#39;SatispayPayment&#39;, &#39;StripePayment&#39;, or &#39;WireTransfer&#39;. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Moto** | Pointer to **interface{}** | Send this attribute if you want to mark the payment as MOTO, must be supported by payment gateway. | [optional] 
**RequireCapture** | Pointer to **interface{}** | Send this attribute if you want to require the payment capture before fulfillment. | [optional] 
**AutoCapture** | Pointer to **interface{}** | Send this attribute if you want to automatically capture the payment upon authorization. | [optional] 
**Disable** | Pointer to **interface{}** | Send this attribute if you want to mark the payment method as disabled. | [optional] 
**Enable** | Pointer to **interface{}** | Send this attribute if you want to mark the payment method as enabled. | [optional] 
**PriceAmountCents** | Pointer to **interface{}** | The payment method&#39;s price, in cents | [optional] 
**AutoCaptureMaxAmountCents** | Pointer to **interface{}** | Send this attribute if you want to limit automatic capture to orders for which the total amount is equal or less than the specified value, in cents. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetPaymentSourceType() interface{}`

GetPaymentSourceType returns the PaymentSourceType field if non-nil, zero value otherwise.

### GetPaymentSourceTypeOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetPaymentSourceTypeOk() (*interface{}, bool)`

GetPaymentSourceTypeOk returns a tuple with the PaymentSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceType

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetPaymentSourceType(v interface{})`

SetPaymentSourceType sets PaymentSourceType field to given value.

### HasPaymentSourceType

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasPaymentSourceType() bool`

HasPaymentSourceType returns a boolean if a field has been set.

### SetPaymentSourceTypeNil

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetPaymentSourceTypeNil(b bool)`

 SetPaymentSourceTypeNil sets the value for PaymentSourceType to be an explicit nil

### UnsetPaymentSourceType
`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) UnsetPaymentSourceType()`

UnsetPaymentSourceType ensures that no value is present for PaymentSourceType, not even an explicit nil
### GetCurrencyCode

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetMoto

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetMoto() interface{}`

GetMoto returns the Moto field if non-nil, zero value otherwise.

### GetMotoOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetMotoOk() (*interface{}, bool)`

GetMotoOk returns a tuple with the Moto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoto

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetMoto(v interface{})`

SetMoto sets Moto field to given value.

### HasMoto

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasMoto() bool`

HasMoto returns a boolean if a field has been set.

### SetMotoNil

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetMotoNil(b bool)`

 SetMotoNil sets the value for Moto to be an explicit nil

### UnsetMoto
`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) UnsetMoto()`

UnsetMoto ensures that no value is present for Moto, not even an explicit nil
### GetRequireCapture

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetRequireCapture() interface{}`

GetRequireCapture returns the RequireCapture field if non-nil, zero value otherwise.

### GetRequireCaptureOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetRequireCaptureOk() (*interface{}, bool)`

GetRequireCaptureOk returns a tuple with the RequireCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCapture

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetRequireCapture(v interface{})`

SetRequireCapture sets RequireCapture field to given value.

### HasRequireCapture

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasRequireCapture() bool`

HasRequireCapture returns a boolean if a field has been set.

### SetRequireCaptureNil

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetRequireCaptureNil(b bool)`

 SetRequireCaptureNil sets the value for RequireCapture to be an explicit nil

### UnsetRequireCapture
`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) UnsetRequireCapture()`

UnsetRequireCapture ensures that no value is present for RequireCapture, not even an explicit nil
### GetAutoCapture

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetAutoCapture() interface{}`

GetAutoCapture returns the AutoCapture field if non-nil, zero value otherwise.

### GetAutoCaptureOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetAutoCaptureOk() (*interface{}, bool)`

GetAutoCaptureOk returns a tuple with the AutoCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCapture

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetAutoCapture(v interface{})`

SetAutoCapture sets AutoCapture field to given value.

### HasAutoCapture

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasAutoCapture() bool`

HasAutoCapture returns a boolean if a field has been set.

### SetAutoCaptureNil

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetAutoCaptureNil(b bool)`

 SetAutoCaptureNil sets the value for AutoCapture to be an explicit nil

### UnsetAutoCapture
`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) UnsetAutoCapture()`

UnsetAutoCapture ensures that no value is present for AutoCapture, not even an explicit nil
### GetDisable

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetDisable() interface{}`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetDisableOk() (*interface{}, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetDisable(v interface{})`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### SetDisableNil

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetDisableNil(b bool)`

 SetDisableNil sets the value for Disable to be an explicit nil

### UnsetDisable
`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) UnsetDisable()`

UnsetDisable ensures that no value is present for Disable, not even an explicit nil
### GetEnable

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetEnable() interface{}`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetEnableOk() (*interface{}, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetEnable(v interface{})`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetPriceAmountCents

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### SetPriceAmountCentsNil

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetAutoCaptureMaxAmountCents

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetAutoCaptureMaxAmountCents() interface{}`

GetAutoCaptureMaxAmountCents returns the AutoCaptureMaxAmountCents field if non-nil, zero value otherwise.

### GetAutoCaptureMaxAmountCentsOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetAutoCaptureMaxAmountCentsOk() (*interface{}, bool)`

GetAutoCaptureMaxAmountCentsOk returns a tuple with the AutoCaptureMaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCaptureMaxAmountCents

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetAutoCaptureMaxAmountCents(v interface{})`

SetAutoCaptureMaxAmountCents sets AutoCaptureMaxAmountCents field to given value.

### HasAutoCaptureMaxAmountCents

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasAutoCaptureMaxAmountCents() bool`

HasAutoCaptureMaxAmountCents returns a boolean if a field has been set.

### SetAutoCaptureMaxAmountCentsNil

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetAutoCaptureMaxAmountCentsNil(b bool)`

 SetAutoCaptureMaxAmountCentsNil sets the value for AutoCaptureMaxAmountCents to be an explicit nil

### UnsetAutoCaptureMaxAmountCents
`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) UnsetAutoCaptureMaxAmountCents()`

UnsetAutoCaptureMaxAmountCents ensures that no value is present for AutoCaptureMaxAmountCents, not even an explicit nil
### GetReference

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


