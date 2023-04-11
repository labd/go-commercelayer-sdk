# GETStripePayments200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSecret** | Pointer to **interface{}** | The Stripe payment intent client secret. Required to create a charge through Stripe.js. | [optional] 
**PublishableKey** | Pointer to **interface{}** | The Stripe publishable API key. | [optional] 
**Options** | Pointer to **interface{}** | Stripe payment options: &#39;customer&#39;, &#39;payment_method&#39;, etc. Check Stripe payment intent API for more details. | [optional] 
**PaymentMethod** | Pointer to **interface{}** | Stripe &#39;payment_method&#39;, set by webhook. | [optional] 
**MismatchedAmounts** | Pointer to **interface{}** | Indicates if the order current amount differs form the one of the created payment intent. | [optional] 
**IntentAmountCents** | Pointer to **interface{}** | The amount of the associated payment intent, in cents. | [optional] 
**IntentAmountFloat** | Pointer to **interface{}** | The amount of the associated payment intent, float. | [optional] 
**FormattedIntentAmount** | Pointer to **interface{}** | The amount of the associated payment intent, formatted. | [optional] 
**PaymentInstrument** | Pointer to **interface{}** | Information about the payment instrument used in the transaction | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETStripePayments200ResponseDataInnerAttributes

`func NewGETStripePayments200ResponseDataInnerAttributes() *GETStripePayments200ResponseDataInnerAttributes`

NewGETStripePayments200ResponseDataInnerAttributes instantiates a new GETStripePayments200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETStripePayments200ResponseDataInnerAttributesWithDefaults

`func NewGETStripePayments200ResponseDataInnerAttributesWithDefaults() *GETStripePayments200ResponseDataInnerAttributes`

NewGETStripePayments200ResponseDataInnerAttributesWithDefaults instantiates a new GETStripePayments200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSecret

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetClientSecret() interface{}`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetClientSecretOk() (*interface{}, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetClientSecret(v interface{})`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *GETStripePayments200ResponseDataInnerAttributes) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *GETStripePayments200ResponseDataInnerAttributes) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetPublishableKey

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetPublishableKey() interface{}`

GetPublishableKey returns the PublishableKey field if non-nil, zero value otherwise.

### GetPublishableKeyOk

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetPublishableKeyOk() (*interface{}, bool)`

GetPublishableKeyOk returns a tuple with the PublishableKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishableKey

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetPublishableKey(v interface{})`

SetPublishableKey sets PublishableKey field to given value.

### HasPublishableKey

`func (o *GETStripePayments200ResponseDataInnerAttributes) HasPublishableKey() bool`

HasPublishableKey returns a boolean if a field has been set.

### SetPublishableKeyNil

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetPublishableKeyNil(b bool)`

 SetPublishableKeyNil sets the value for PublishableKey to be an explicit nil

### UnsetPublishableKey
`func (o *GETStripePayments200ResponseDataInnerAttributes) UnsetPublishableKey()`

UnsetPublishableKey ensures that no value is present for PublishableKey, not even an explicit nil
### GetOptions

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GETStripePayments200ResponseDataInnerAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *GETStripePayments200ResponseDataInnerAttributes) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetPaymentMethod

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetPaymentMethod() interface{}`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetPaymentMethodOk() (*interface{}, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetPaymentMethod(v interface{})`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *GETStripePayments200ResponseDataInnerAttributes) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### SetPaymentMethodNil

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *GETStripePayments200ResponseDataInnerAttributes) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
### GetMismatchedAmounts

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetMismatchedAmounts() interface{}`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetMismatchedAmountsOk() (*interface{}, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetMismatchedAmounts(v interface{})`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *GETStripePayments200ResponseDataInnerAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### SetMismatchedAmountsNil

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetMismatchedAmountsNil(b bool)`

 SetMismatchedAmountsNil sets the value for MismatchedAmounts to be an explicit nil

### UnsetMismatchedAmounts
`func (o *GETStripePayments200ResponseDataInnerAttributes) UnsetMismatchedAmounts()`

UnsetMismatchedAmounts ensures that no value is present for MismatchedAmounts, not even an explicit nil
### GetIntentAmountCents

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetIntentAmountCents() interface{}`

GetIntentAmountCents returns the IntentAmountCents field if non-nil, zero value otherwise.

### GetIntentAmountCentsOk

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetIntentAmountCentsOk() (*interface{}, bool)`

GetIntentAmountCentsOk returns a tuple with the IntentAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentAmountCents

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetIntentAmountCents(v interface{})`

SetIntentAmountCents sets IntentAmountCents field to given value.

### HasIntentAmountCents

`func (o *GETStripePayments200ResponseDataInnerAttributes) HasIntentAmountCents() bool`

HasIntentAmountCents returns a boolean if a field has been set.

### SetIntentAmountCentsNil

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetIntentAmountCentsNil(b bool)`

 SetIntentAmountCentsNil sets the value for IntentAmountCents to be an explicit nil

### UnsetIntentAmountCents
`func (o *GETStripePayments200ResponseDataInnerAttributes) UnsetIntentAmountCents()`

UnsetIntentAmountCents ensures that no value is present for IntentAmountCents, not even an explicit nil
### GetIntentAmountFloat

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetIntentAmountFloat() interface{}`

GetIntentAmountFloat returns the IntentAmountFloat field if non-nil, zero value otherwise.

### GetIntentAmountFloatOk

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetIntentAmountFloatOk() (*interface{}, bool)`

GetIntentAmountFloatOk returns a tuple with the IntentAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentAmountFloat

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetIntentAmountFloat(v interface{})`

SetIntentAmountFloat sets IntentAmountFloat field to given value.

### HasIntentAmountFloat

`func (o *GETStripePayments200ResponseDataInnerAttributes) HasIntentAmountFloat() bool`

HasIntentAmountFloat returns a boolean if a field has been set.

### SetIntentAmountFloatNil

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetIntentAmountFloatNil(b bool)`

 SetIntentAmountFloatNil sets the value for IntentAmountFloat to be an explicit nil

### UnsetIntentAmountFloat
`func (o *GETStripePayments200ResponseDataInnerAttributes) UnsetIntentAmountFloat()`

UnsetIntentAmountFloat ensures that no value is present for IntentAmountFloat, not even an explicit nil
### GetFormattedIntentAmount

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetFormattedIntentAmount() interface{}`

GetFormattedIntentAmount returns the FormattedIntentAmount field if non-nil, zero value otherwise.

### GetFormattedIntentAmountOk

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetFormattedIntentAmountOk() (*interface{}, bool)`

GetFormattedIntentAmountOk returns a tuple with the FormattedIntentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedIntentAmount

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetFormattedIntentAmount(v interface{})`

SetFormattedIntentAmount sets FormattedIntentAmount field to given value.

### HasFormattedIntentAmount

`func (o *GETStripePayments200ResponseDataInnerAttributes) HasFormattedIntentAmount() bool`

HasFormattedIntentAmount returns a boolean if a field has been set.

### SetFormattedIntentAmountNil

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetFormattedIntentAmountNil(b bool)`

 SetFormattedIntentAmountNil sets the value for FormattedIntentAmount to be an explicit nil

### UnsetFormattedIntentAmount
`func (o *GETStripePayments200ResponseDataInnerAttributes) UnsetFormattedIntentAmount()`

UnsetFormattedIntentAmount ensures that no value is present for FormattedIntentAmount, not even an explicit nil
### GetPaymentInstrument

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetPaymentInstrument() interface{}`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetPaymentInstrumentOk() (*interface{}, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetPaymentInstrument(v interface{})`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *GETStripePayments200ResponseDataInnerAttributes) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### SetPaymentInstrumentNil

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetPaymentInstrumentNil(b bool)`

 SetPaymentInstrumentNil sets the value for PaymentInstrument to be an explicit nil

### UnsetPaymentInstrument
`func (o *GETStripePayments200ResponseDataInnerAttributes) UnsetPaymentInstrument()`

UnsetPaymentInstrument ensures that no value is present for PaymentInstrument, not even an explicit nil
### GetCreatedAt

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETStripePayments200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETStripePayments200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETStripePayments200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETStripePayments200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETStripePayments200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETStripePayments200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETStripePayments200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETStripePayments200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETStripePayments200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETStripePayments200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETStripePayments200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETStripePayments200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


