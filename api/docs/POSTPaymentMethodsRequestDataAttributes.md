# POSTPaymentMethodsRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentSourceType** | **interface{}** | The payment source type, can be one of: &#39;AdyenPayment&#39;, &#39;AxervePayment&#39;, &#39;BraintreePayment&#39;, &#39;CheckoutComPayment&#39;, &#39;CreditCard&#39;, &#39;ExternalPayment&#39;, &#39;KlarnaPayment&#39;, &#39;PaypalPayment&#39;, &#39;SatispayPayment&#39;, &#39;StripePayment&#39;, or &#39;WireTransfer&#39;. | 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Moto** | Pointer to **interface{}** | Send this attribute if you want to mark the payment as MOTO, must be supported by payment gateway. | [optional] 
**RequireCapture** | Pointer to **interface{}** | Send this attribute if you want to require the payment capture before fulfillment. | [optional] 
**AutoCapture** | Pointer to **interface{}** | Send this attribute if you want to automatically capture the payment upon authorization. | [optional] 
**PriceAmountCents** | **interface{}** | The payment method&#39;s price, in cents | 
**AutoCaptureMaxAmountCents** | Pointer to **interface{}** | Send this attribute if you want to limit automatic capture to orders for which the total amount is equal or less than the specified value, in cents. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTPaymentMethodsRequestDataAttributes

`func NewPOSTPaymentMethodsRequestDataAttributes(paymentSourceType interface{}, priceAmountCents interface{}, ) *POSTPaymentMethodsRequestDataAttributes`

NewPOSTPaymentMethodsRequestDataAttributes instantiates a new POSTPaymentMethodsRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPaymentMethodsRequestDataAttributesWithDefaults

`func NewPOSTPaymentMethodsRequestDataAttributesWithDefaults() *POSTPaymentMethodsRequestDataAttributes`

NewPOSTPaymentMethodsRequestDataAttributesWithDefaults instantiates a new POSTPaymentMethodsRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentSourceType

`func (o *POSTPaymentMethodsRequestDataAttributes) GetPaymentSourceType() interface{}`

GetPaymentSourceType returns the PaymentSourceType field if non-nil, zero value otherwise.

### GetPaymentSourceTypeOk

`func (o *POSTPaymentMethodsRequestDataAttributes) GetPaymentSourceTypeOk() (*interface{}, bool)`

GetPaymentSourceTypeOk returns a tuple with the PaymentSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceType

`func (o *POSTPaymentMethodsRequestDataAttributes) SetPaymentSourceType(v interface{})`

SetPaymentSourceType sets PaymentSourceType field to given value.


### SetPaymentSourceTypeNil

`func (o *POSTPaymentMethodsRequestDataAttributes) SetPaymentSourceTypeNil(b bool)`

 SetPaymentSourceTypeNil sets the value for PaymentSourceType to be an explicit nil

### UnsetPaymentSourceType
`func (o *POSTPaymentMethodsRequestDataAttributes) UnsetPaymentSourceType()`

UnsetPaymentSourceType ensures that no value is present for PaymentSourceType, not even an explicit nil
### GetCurrencyCode

`func (o *POSTPaymentMethodsRequestDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *POSTPaymentMethodsRequestDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *POSTPaymentMethodsRequestDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *POSTPaymentMethodsRequestDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *POSTPaymentMethodsRequestDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *POSTPaymentMethodsRequestDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetMoto

`func (o *POSTPaymentMethodsRequestDataAttributes) GetMoto() interface{}`

GetMoto returns the Moto field if non-nil, zero value otherwise.

### GetMotoOk

`func (o *POSTPaymentMethodsRequestDataAttributes) GetMotoOk() (*interface{}, bool)`

GetMotoOk returns a tuple with the Moto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoto

`func (o *POSTPaymentMethodsRequestDataAttributes) SetMoto(v interface{})`

SetMoto sets Moto field to given value.

### HasMoto

`func (o *POSTPaymentMethodsRequestDataAttributes) HasMoto() bool`

HasMoto returns a boolean if a field has been set.

### SetMotoNil

`func (o *POSTPaymentMethodsRequestDataAttributes) SetMotoNil(b bool)`

 SetMotoNil sets the value for Moto to be an explicit nil

### UnsetMoto
`func (o *POSTPaymentMethodsRequestDataAttributes) UnsetMoto()`

UnsetMoto ensures that no value is present for Moto, not even an explicit nil
### GetRequireCapture

`func (o *POSTPaymentMethodsRequestDataAttributes) GetRequireCapture() interface{}`

GetRequireCapture returns the RequireCapture field if non-nil, zero value otherwise.

### GetRequireCaptureOk

`func (o *POSTPaymentMethodsRequestDataAttributes) GetRequireCaptureOk() (*interface{}, bool)`

GetRequireCaptureOk returns a tuple with the RequireCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCapture

`func (o *POSTPaymentMethodsRequestDataAttributes) SetRequireCapture(v interface{})`

SetRequireCapture sets RequireCapture field to given value.

### HasRequireCapture

`func (o *POSTPaymentMethodsRequestDataAttributes) HasRequireCapture() bool`

HasRequireCapture returns a boolean if a field has been set.

### SetRequireCaptureNil

`func (o *POSTPaymentMethodsRequestDataAttributes) SetRequireCaptureNil(b bool)`

 SetRequireCaptureNil sets the value for RequireCapture to be an explicit nil

### UnsetRequireCapture
`func (o *POSTPaymentMethodsRequestDataAttributes) UnsetRequireCapture()`

UnsetRequireCapture ensures that no value is present for RequireCapture, not even an explicit nil
### GetAutoCapture

`func (o *POSTPaymentMethodsRequestDataAttributes) GetAutoCapture() interface{}`

GetAutoCapture returns the AutoCapture field if non-nil, zero value otherwise.

### GetAutoCaptureOk

`func (o *POSTPaymentMethodsRequestDataAttributes) GetAutoCaptureOk() (*interface{}, bool)`

GetAutoCaptureOk returns a tuple with the AutoCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCapture

`func (o *POSTPaymentMethodsRequestDataAttributes) SetAutoCapture(v interface{})`

SetAutoCapture sets AutoCapture field to given value.

### HasAutoCapture

`func (o *POSTPaymentMethodsRequestDataAttributes) HasAutoCapture() bool`

HasAutoCapture returns a boolean if a field has been set.

### SetAutoCaptureNil

`func (o *POSTPaymentMethodsRequestDataAttributes) SetAutoCaptureNil(b bool)`

 SetAutoCaptureNil sets the value for AutoCapture to be an explicit nil

### UnsetAutoCapture
`func (o *POSTPaymentMethodsRequestDataAttributes) UnsetAutoCapture()`

UnsetAutoCapture ensures that no value is present for AutoCapture, not even an explicit nil
### GetPriceAmountCents

`func (o *POSTPaymentMethodsRequestDataAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *POSTPaymentMethodsRequestDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *POSTPaymentMethodsRequestDataAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.


### SetPriceAmountCentsNil

`func (o *POSTPaymentMethodsRequestDataAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *POSTPaymentMethodsRequestDataAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetAutoCaptureMaxAmountCents

`func (o *POSTPaymentMethodsRequestDataAttributes) GetAutoCaptureMaxAmountCents() interface{}`

GetAutoCaptureMaxAmountCents returns the AutoCaptureMaxAmountCents field if non-nil, zero value otherwise.

### GetAutoCaptureMaxAmountCentsOk

`func (o *POSTPaymentMethodsRequestDataAttributes) GetAutoCaptureMaxAmountCentsOk() (*interface{}, bool)`

GetAutoCaptureMaxAmountCentsOk returns a tuple with the AutoCaptureMaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCaptureMaxAmountCents

`func (o *POSTPaymentMethodsRequestDataAttributes) SetAutoCaptureMaxAmountCents(v interface{})`

SetAutoCaptureMaxAmountCents sets AutoCaptureMaxAmountCents field to given value.

### HasAutoCaptureMaxAmountCents

`func (o *POSTPaymentMethodsRequestDataAttributes) HasAutoCaptureMaxAmountCents() bool`

HasAutoCaptureMaxAmountCents returns a boolean if a field has been set.

### SetAutoCaptureMaxAmountCentsNil

`func (o *POSTPaymentMethodsRequestDataAttributes) SetAutoCaptureMaxAmountCentsNil(b bool)`

 SetAutoCaptureMaxAmountCentsNil sets the value for AutoCaptureMaxAmountCents to be an explicit nil

### UnsetAutoCaptureMaxAmountCents
`func (o *POSTPaymentMethodsRequestDataAttributes) UnsetAutoCaptureMaxAmountCents()`

UnsetAutoCaptureMaxAmountCents ensures that no value is present for AutoCaptureMaxAmountCents, not even an explicit nil
### GetReference

`func (o *POSTPaymentMethodsRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTPaymentMethodsRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTPaymentMethodsRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTPaymentMethodsRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTPaymentMethodsRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTPaymentMethodsRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTPaymentMethodsRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTPaymentMethodsRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTPaymentMethodsRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTPaymentMethodsRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTPaymentMethodsRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTPaymentMethodsRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTPaymentMethodsRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTPaymentMethodsRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTPaymentMethodsRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTPaymentMethodsRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTPaymentMethodsRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTPaymentMethodsRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


