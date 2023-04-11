# GETPaymentMethods200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentSourceType** | Pointer to **interface{}** | The payment source type, can be one of: &#39;AdyenPayment&#39;, &#39;AxervePayment&#39;, &#39;BraintreePayment&#39;, &#39;CheckoutComPayment&#39;, &#39;CreditCard&#39;, &#39;ExternalPayment&#39;, &#39;KlarnaPayment&#39;, &#39;PaypalPayment&#39;, &#39;SatispayPayment&#39;, &#39;StripePayment&#39;, or &#39;WireTransfer&#39;. | [optional] 
**Name** | Pointer to **interface{}** | Payment source type, titleized | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Moto** | Pointer to **interface{}** | Send this attribute if you want to mark the payment as MOTO, must be supported by payment gateway. | [optional] 
**RequireCapture** | Pointer to **interface{}** | Send this attribute if you want to require the payment capture before fulfillment. | [optional] 
**AutoCapture** | Pointer to **interface{}** | Send this attribute if you want to automatically capture the payment upon authorization. | [optional] 
**DisabledAt** | Pointer to **interface{}** | Time at which the payment method was disabled. | [optional] 
**PriceAmountCents** | Pointer to **interface{}** | The payment method&#39;s price, in cents | [optional] 
**PriceAmountFloat** | Pointer to **interface{}** | The payment method&#39;s price, float | [optional] 
**FormattedPriceAmount** | Pointer to **interface{}** | The payment method&#39;s price, formatted | [optional] 
**AutoCaptureMaxAmountCents** | Pointer to **interface{}** | Send this attribute if you want to limit automatic capture to orders for which the total amount is equal or less than the specified value, in cents. | [optional] 
**AutoCaptureMaxAmountFloat** | Pointer to **interface{}** | The automatic capture max amount, float | [optional] 
**FormattedAutoCaptureMaxAmount** | Pointer to **interface{}** | The automatic capture max amount, formatted | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETPaymentMethods200ResponseDataInnerAttributes

`func NewGETPaymentMethods200ResponseDataInnerAttributes() *GETPaymentMethods200ResponseDataInnerAttributes`

NewGETPaymentMethods200ResponseDataInnerAttributes instantiates a new GETPaymentMethods200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETPaymentMethods200ResponseDataInnerAttributesWithDefaults

`func NewGETPaymentMethods200ResponseDataInnerAttributesWithDefaults() *GETPaymentMethods200ResponseDataInnerAttributes`

NewGETPaymentMethods200ResponseDataInnerAttributesWithDefaults instantiates a new GETPaymentMethods200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentSourceType

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetPaymentSourceType() interface{}`

GetPaymentSourceType returns the PaymentSourceType field if non-nil, zero value otherwise.

### GetPaymentSourceTypeOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetPaymentSourceTypeOk() (*interface{}, bool)`

GetPaymentSourceTypeOk returns a tuple with the PaymentSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceType

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetPaymentSourceType(v interface{})`

SetPaymentSourceType sets PaymentSourceType field to given value.

### HasPaymentSourceType

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasPaymentSourceType() bool`

HasPaymentSourceType returns a boolean if a field has been set.

### SetPaymentSourceTypeNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetPaymentSourceTypeNil(b bool)`

 SetPaymentSourceTypeNil sets the value for PaymentSourceType to be an explicit nil

### UnsetPaymentSourceType
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetPaymentSourceType()`

UnsetPaymentSourceType ensures that no value is present for PaymentSourceType, not even an explicit nil
### GetName

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrencyCode

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetMoto

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetMoto() interface{}`

GetMoto returns the Moto field if non-nil, zero value otherwise.

### GetMotoOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetMotoOk() (*interface{}, bool)`

GetMotoOk returns a tuple with the Moto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoto

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetMoto(v interface{})`

SetMoto sets Moto field to given value.

### HasMoto

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasMoto() bool`

HasMoto returns a boolean if a field has been set.

### SetMotoNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetMotoNil(b bool)`

 SetMotoNil sets the value for Moto to be an explicit nil

### UnsetMoto
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetMoto()`

UnsetMoto ensures that no value is present for Moto, not even an explicit nil
### GetRequireCapture

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetRequireCapture() interface{}`

GetRequireCapture returns the RequireCapture field if non-nil, zero value otherwise.

### GetRequireCaptureOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetRequireCaptureOk() (*interface{}, bool)`

GetRequireCaptureOk returns a tuple with the RequireCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCapture

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetRequireCapture(v interface{})`

SetRequireCapture sets RequireCapture field to given value.

### HasRequireCapture

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasRequireCapture() bool`

HasRequireCapture returns a boolean if a field has been set.

### SetRequireCaptureNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetRequireCaptureNil(b bool)`

 SetRequireCaptureNil sets the value for RequireCapture to be an explicit nil

### UnsetRequireCapture
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetRequireCapture()`

UnsetRequireCapture ensures that no value is present for RequireCapture, not even an explicit nil
### GetAutoCapture

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetAutoCapture() interface{}`

GetAutoCapture returns the AutoCapture field if non-nil, zero value otherwise.

### GetAutoCaptureOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetAutoCaptureOk() (*interface{}, bool)`

GetAutoCaptureOk returns a tuple with the AutoCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCapture

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetAutoCapture(v interface{})`

SetAutoCapture sets AutoCapture field to given value.

### HasAutoCapture

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasAutoCapture() bool`

HasAutoCapture returns a boolean if a field has been set.

### SetAutoCaptureNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetAutoCaptureNil(b bool)`

 SetAutoCaptureNil sets the value for AutoCapture to be an explicit nil

### UnsetAutoCapture
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetAutoCapture()`

UnsetAutoCapture ensures that no value is present for AutoCapture, not even an explicit nil
### GetDisabledAt

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetDisabledAt() interface{}`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetDisabledAtOk() (*interface{}, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetDisabledAt(v interface{})`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### SetDisabledAtNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetDisabledAtNil(b bool)`

 SetDisabledAtNil sets the value for DisabledAt to be an explicit nil

### UnsetDisabledAt
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetDisabledAt()`

UnsetDisabledAt ensures that no value is present for DisabledAt, not even an explicit nil
### GetPriceAmountCents

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### SetPriceAmountCentsNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetPriceAmountFloat

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetPriceAmountFloat() interface{}`

GetPriceAmountFloat returns the PriceAmountFloat field if non-nil, zero value otherwise.

### GetPriceAmountFloatOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetPriceAmountFloatOk() (*interface{}, bool)`

GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountFloat

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetPriceAmountFloat(v interface{})`

SetPriceAmountFloat sets PriceAmountFloat field to given value.

### HasPriceAmountFloat

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasPriceAmountFloat() bool`

HasPriceAmountFloat returns a boolean if a field has been set.

### SetPriceAmountFloatNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetPriceAmountFloatNil(b bool)`

 SetPriceAmountFloatNil sets the value for PriceAmountFloat to be an explicit nil

### UnsetPriceAmountFloat
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetPriceAmountFloat()`

UnsetPriceAmountFloat ensures that no value is present for PriceAmountFloat, not even an explicit nil
### GetFormattedPriceAmount

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetFormattedPriceAmount() interface{}`

GetFormattedPriceAmount returns the FormattedPriceAmount field if non-nil, zero value otherwise.

### GetFormattedPriceAmountOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetFormattedPriceAmountOk() (*interface{}, bool)`

GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmount

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetFormattedPriceAmount(v interface{})`

SetFormattedPriceAmount sets FormattedPriceAmount field to given value.

### HasFormattedPriceAmount

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasFormattedPriceAmount() bool`

HasFormattedPriceAmount returns a boolean if a field has been set.

### SetFormattedPriceAmountNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetFormattedPriceAmountNil(b bool)`

 SetFormattedPriceAmountNil sets the value for FormattedPriceAmount to be an explicit nil

### UnsetFormattedPriceAmount
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetFormattedPriceAmount()`

UnsetFormattedPriceAmount ensures that no value is present for FormattedPriceAmount, not even an explicit nil
### GetAutoCaptureMaxAmountCents

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetAutoCaptureMaxAmountCents() interface{}`

GetAutoCaptureMaxAmountCents returns the AutoCaptureMaxAmountCents field if non-nil, zero value otherwise.

### GetAutoCaptureMaxAmountCentsOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetAutoCaptureMaxAmountCentsOk() (*interface{}, bool)`

GetAutoCaptureMaxAmountCentsOk returns a tuple with the AutoCaptureMaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCaptureMaxAmountCents

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetAutoCaptureMaxAmountCents(v interface{})`

SetAutoCaptureMaxAmountCents sets AutoCaptureMaxAmountCents field to given value.

### HasAutoCaptureMaxAmountCents

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasAutoCaptureMaxAmountCents() bool`

HasAutoCaptureMaxAmountCents returns a boolean if a field has been set.

### SetAutoCaptureMaxAmountCentsNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetAutoCaptureMaxAmountCentsNil(b bool)`

 SetAutoCaptureMaxAmountCentsNil sets the value for AutoCaptureMaxAmountCents to be an explicit nil

### UnsetAutoCaptureMaxAmountCents
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetAutoCaptureMaxAmountCents()`

UnsetAutoCaptureMaxAmountCents ensures that no value is present for AutoCaptureMaxAmountCents, not even an explicit nil
### GetAutoCaptureMaxAmountFloat

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetAutoCaptureMaxAmountFloat() interface{}`

GetAutoCaptureMaxAmountFloat returns the AutoCaptureMaxAmountFloat field if non-nil, zero value otherwise.

### GetAutoCaptureMaxAmountFloatOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetAutoCaptureMaxAmountFloatOk() (*interface{}, bool)`

GetAutoCaptureMaxAmountFloatOk returns a tuple with the AutoCaptureMaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCaptureMaxAmountFloat

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetAutoCaptureMaxAmountFloat(v interface{})`

SetAutoCaptureMaxAmountFloat sets AutoCaptureMaxAmountFloat field to given value.

### HasAutoCaptureMaxAmountFloat

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasAutoCaptureMaxAmountFloat() bool`

HasAutoCaptureMaxAmountFloat returns a boolean if a field has been set.

### SetAutoCaptureMaxAmountFloatNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetAutoCaptureMaxAmountFloatNil(b bool)`

 SetAutoCaptureMaxAmountFloatNil sets the value for AutoCaptureMaxAmountFloat to be an explicit nil

### UnsetAutoCaptureMaxAmountFloat
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetAutoCaptureMaxAmountFloat()`

UnsetAutoCaptureMaxAmountFloat ensures that no value is present for AutoCaptureMaxAmountFloat, not even an explicit nil
### GetFormattedAutoCaptureMaxAmount

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetFormattedAutoCaptureMaxAmount() interface{}`

GetFormattedAutoCaptureMaxAmount returns the FormattedAutoCaptureMaxAmount field if non-nil, zero value otherwise.

### GetFormattedAutoCaptureMaxAmountOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetFormattedAutoCaptureMaxAmountOk() (*interface{}, bool)`

GetFormattedAutoCaptureMaxAmountOk returns a tuple with the FormattedAutoCaptureMaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAutoCaptureMaxAmount

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetFormattedAutoCaptureMaxAmount(v interface{})`

SetFormattedAutoCaptureMaxAmount sets FormattedAutoCaptureMaxAmount field to given value.

### HasFormattedAutoCaptureMaxAmount

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasFormattedAutoCaptureMaxAmount() bool`

HasFormattedAutoCaptureMaxAmount returns a boolean if a field has been set.

### SetFormattedAutoCaptureMaxAmountNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetFormattedAutoCaptureMaxAmountNil(b bool)`

 SetFormattedAutoCaptureMaxAmountNil sets the value for FormattedAutoCaptureMaxAmount to be an explicit nil

### UnsetFormattedAutoCaptureMaxAmount
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetFormattedAutoCaptureMaxAmount()`

UnsetFormattedAutoCaptureMaxAmount ensures that no value is present for FormattedAutoCaptureMaxAmount, not even an explicit nil
### GetCreatedAt

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETPaymentMethods200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETPaymentMethods200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


