# GETGiftCardsGiftCardId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **interface{}** | The gift card status, one of &#39;draft&#39;, &#39;inactive&#39;, &#39;active&#39;, or &#39;redeemed&#39;. | [optional] 
**Code** | Pointer to **interface{}** | The gift card code UUID. If not set, it&#39;s automatically generated. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**InitialBalanceCents** | Pointer to **interface{}** | The gift card initial balance, in cents. | [optional] 
**InitialBalanceFloat** | Pointer to **interface{}** | The gift card initial balance, float. | [optional] 
**FormattedInitialBalance** | Pointer to **interface{}** | The gift card initial balance, formatted. | [optional] 
**BalanceCents** | Pointer to **interface{}** | The gift card balance, in cents. | [optional] 
**BalanceFloat** | Pointer to **interface{}** | The gift card balance, float. | [optional] 
**FormattedBalance** | Pointer to **interface{}** | The gift card balance, formatted. | [optional] 
**BalanceMaxCents** | Pointer to **interface{}** | The gift card balance max, in cents. | [optional] 
**BalanceMaxFloat** | Pointer to **interface{}** | The gift card balance max, float. | [optional] 
**FormattedBalanceMax** | Pointer to **interface{}** | The gift card balance max, formatted. | [optional] 
**BalanceLog** | Pointer to **interface{}** | The gift card balance log. Tracks all the gift card transactions. | [optional] 
**SingleUse** | Pointer to **interface{}** | Indicates if the gift card can be used only one. | [optional] 
**Rechargeable** | Pointer to **interface{}** | Indicates if the gift card can be recharged. | [optional] 
**ImageUrl** | Pointer to **interface{}** | The URL of an image that represents the gift card. | [optional] 
**ExpiresAt** | Pointer to **interface{}** | Time at which the gift card will expire. | [optional] 
**RecipientEmail** | Pointer to **interface{}** | The email address of the associated recipient. When creating or updating a gift card, this is a shortcut to find or create the associated recipient by email. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETGiftCardsGiftCardId200ResponseDataAttributes

`func NewGETGiftCardsGiftCardId200ResponseDataAttributes() *GETGiftCardsGiftCardId200ResponseDataAttributes`

NewGETGiftCardsGiftCardId200ResponseDataAttributes instantiates a new GETGiftCardsGiftCardId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETGiftCardsGiftCardId200ResponseDataAttributesWithDefaults

`func NewGETGiftCardsGiftCardId200ResponseDataAttributesWithDefaults() *GETGiftCardsGiftCardId200ResponseDataAttributes`

NewGETGiftCardsGiftCardId200ResponseDataAttributesWithDefaults instantiates a new GETGiftCardsGiftCardId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCode

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCurrencyCode

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetInitialBalanceCents

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetInitialBalanceCents() interface{}`

GetInitialBalanceCents returns the InitialBalanceCents field if non-nil, zero value otherwise.

### GetInitialBalanceCentsOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetInitialBalanceCentsOk() (*interface{}, bool)`

GetInitialBalanceCentsOk returns a tuple with the InitialBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialBalanceCents

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetInitialBalanceCents(v interface{})`

SetInitialBalanceCents sets InitialBalanceCents field to given value.

### HasInitialBalanceCents

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasInitialBalanceCents() bool`

HasInitialBalanceCents returns a boolean if a field has been set.

### SetInitialBalanceCentsNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetInitialBalanceCentsNil(b bool)`

 SetInitialBalanceCentsNil sets the value for InitialBalanceCents to be an explicit nil

### UnsetInitialBalanceCents
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetInitialBalanceCents()`

UnsetInitialBalanceCents ensures that no value is present for InitialBalanceCents, not even an explicit nil
### GetInitialBalanceFloat

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetInitialBalanceFloat() interface{}`

GetInitialBalanceFloat returns the InitialBalanceFloat field if non-nil, zero value otherwise.

### GetInitialBalanceFloatOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetInitialBalanceFloatOk() (*interface{}, bool)`

GetInitialBalanceFloatOk returns a tuple with the InitialBalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialBalanceFloat

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetInitialBalanceFloat(v interface{})`

SetInitialBalanceFloat sets InitialBalanceFloat field to given value.

### HasInitialBalanceFloat

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasInitialBalanceFloat() bool`

HasInitialBalanceFloat returns a boolean if a field has been set.

### SetInitialBalanceFloatNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetInitialBalanceFloatNil(b bool)`

 SetInitialBalanceFloatNil sets the value for InitialBalanceFloat to be an explicit nil

### UnsetInitialBalanceFloat
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetInitialBalanceFloat()`

UnsetInitialBalanceFloat ensures that no value is present for InitialBalanceFloat, not even an explicit nil
### GetFormattedInitialBalance

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetFormattedInitialBalance() interface{}`

GetFormattedInitialBalance returns the FormattedInitialBalance field if non-nil, zero value otherwise.

### GetFormattedInitialBalanceOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetFormattedInitialBalanceOk() (*interface{}, bool)`

GetFormattedInitialBalanceOk returns a tuple with the FormattedInitialBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedInitialBalance

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetFormattedInitialBalance(v interface{})`

SetFormattedInitialBalance sets FormattedInitialBalance field to given value.

### HasFormattedInitialBalance

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasFormattedInitialBalance() bool`

HasFormattedInitialBalance returns a boolean if a field has been set.

### SetFormattedInitialBalanceNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetFormattedInitialBalanceNil(b bool)`

 SetFormattedInitialBalanceNil sets the value for FormattedInitialBalance to be an explicit nil

### UnsetFormattedInitialBalance
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetFormattedInitialBalance()`

UnsetFormattedInitialBalance ensures that no value is present for FormattedInitialBalance, not even an explicit nil
### GetBalanceCents

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceCents() interface{}`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceCentsOk() (*interface{}, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceCents(v interface{})`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### SetBalanceCentsNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceCentsNil(b bool)`

 SetBalanceCentsNil sets the value for BalanceCents to be an explicit nil

### UnsetBalanceCents
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetBalanceCents()`

UnsetBalanceCents ensures that no value is present for BalanceCents, not even an explicit nil
### GetBalanceFloat

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceFloat() interface{}`

GetBalanceFloat returns the BalanceFloat field if non-nil, zero value otherwise.

### GetBalanceFloatOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceFloatOk() (*interface{}, bool)`

GetBalanceFloatOk returns a tuple with the BalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceFloat

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceFloat(v interface{})`

SetBalanceFloat sets BalanceFloat field to given value.

### HasBalanceFloat

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceFloat() bool`

HasBalanceFloat returns a boolean if a field has been set.

### SetBalanceFloatNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceFloatNil(b bool)`

 SetBalanceFloatNil sets the value for BalanceFloat to be an explicit nil

### UnsetBalanceFloat
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetBalanceFloat()`

UnsetBalanceFloat ensures that no value is present for BalanceFloat, not even an explicit nil
### GetFormattedBalance

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetFormattedBalance() interface{}`

GetFormattedBalance returns the FormattedBalance field if non-nil, zero value otherwise.

### GetFormattedBalanceOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetFormattedBalanceOk() (*interface{}, bool)`

GetFormattedBalanceOk returns a tuple with the FormattedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedBalance

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetFormattedBalance(v interface{})`

SetFormattedBalance sets FormattedBalance field to given value.

### HasFormattedBalance

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasFormattedBalance() bool`

HasFormattedBalance returns a boolean if a field has been set.

### SetFormattedBalanceNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetFormattedBalanceNil(b bool)`

 SetFormattedBalanceNil sets the value for FormattedBalance to be an explicit nil

### UnsetFormattedBalance
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetFormattedBalance()`

UnsetFormattedBalance ensures that no value is present for FormattedBalance, not even an explicit nil
### GetBalanceMaxCents

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceMaxCents() interface{}`

GetBalanceMaxCents returns the BalanceMaxCents field if non-nil, zero value otherwise.

### GetBalanceMaxCentsOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceMaxCentsOk() (*interface{}, bool)`

GetBalanceMaxCentsOk returns a tuple with the BalanceMaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceMaxCents

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceMaxCents(v interface{})`

SetBalanceMaxCents sets BalanceMaxCents field to given value.

### HasBalanceMaxCents

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceMaxCents() bool`

HasBalanceMaxCents returns a boolean if a field has been set.

### SetBalanceMaxCentsNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceMaxCentsNil(b bool)`

 SetBalanceMaxCentsNil sets the value for BalanceMaxCents to be an explicit nil

### UnsetBalanceMaxCents
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetBalanceMaxCents()`

UnsetBalanceMaxCents ensures that no value is present for BalanceMaxCents, not even an explicit nil
### GetBalanceMaxFloat

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceMaxFloat() interface{}`

GetBalanceMaxFloat returns the BalanceMaxFloat field if non-nil, zero value otherwise.

### GetBalanceMaxFloatOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceMaxFloatOk() (*interface{}, bool)`

GetBalanceMaxFloatOk returns a tuple with the BalanceMaxFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceMaxFloat

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceMaxFloat(v interface{})`

SetBalanceMaxFloat sets BalanceMaxFloat field to given value.

### HasBalanceMaxFloat

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceMaxFloat() bool`

HasBalanceMaxFloat returns a boolean if a field has been set.

### SetBalanceMaxFloatNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceMaxFloatNil(b bool)`

 SetBalanceMaxFloatNil sets the value for BalanceMaxFloat to be an explicit nil

### UnsetBalanceMaxFloat
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetBalanceMaxFloat()`

UnsetBalanceMaxFloat ensures that no value is present for BalanceMaxFloat, not even an explicit nil
### GetFormattedBalanceMax

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetFormattedBalanceMax() interface{}`

GetFormattedBalanceMax returns the FormattedBalanceMax field if non-nil, zero value otherwise.

### GetFormattedBalanceMaxOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetFormattedBalanceMaxOk() (*interface{}, bool)`

GetFormattedBalanceMaxOk returns a tuple with the FormattedBalanceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedBalanceMax

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetFormattedBalanceMax(v interface{})`

SetFormattedBalanceMax sets FormattedBalanceMax field to given value.

### HasFormattedBalanceMax

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasFormattedBalanceMax() bool`

HasFormattedBalanceMax returns a boolean if a field has been set.

### SetFormattedBalanceMaxNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetFormattedBalanceMaxNil(b bool)`

 SetFormattedBalanceMaxNil sets the value for FormattedBalanceMax to be an explicit nil

### UnsetFormattedBalanceMax
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetFormattedBalanceMax()`

UnsetFormattedBalanceMax ensures that no value is present for FormattedBalanceMax, not even an explicit nil
### GetBalanceLog

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceLog() interface{}`

GetBalanceLog returns the BalanceLog field if non-nil, zero value otherwise.

### GetBalanceLogOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceLogOk() (*interface{}, bool)`

GetBalanceLogOk returns a tuple with the BalanceLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceLog

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceLog(v interface{})`

SetBalanceLog sets BalanceLog field to given value.

### HasBalanceLog

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceLog() bool`

HasBalanceLog returns a boolean if a field has been set.

### SetBalanceLogNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceLogNil(b bool)`

 SetBalanceLogNil sets the value for BalanceLog to be an explicit nil

### UnsetBalanceLog
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetBalanceLog()`

UnsetBalanceLog ensures that no value is present for BalanceLog, not even an explicit nil
### GetSingleUse

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetSingleUse() interface{}`

GetSingleUse returns the SingleUse field if non-nil, zero value otherwise.

### GetSingleUseOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetSingleUseOk() (*interface{}, bool)`

GetSingleUseOk returns a tuple with the SingleUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUse

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetSingleUse(v interface{})`

SetSingleUse sets SingleUse field to given value.

### HasSingleUse

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasSingleUse() bool`

HasSingleUse returns a boolean if a field has been set.

### SetSingleUseNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetSingleUseNil(b bool)`

 SetSingleUseNil sets the value for SingleUse to be an explicit nil

### UnsetSingleUse
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetSingleUse()`

UnsetSingleUse ensures that no value is present for SingleUse, not even an explicit nil
### GetRechargeable

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetRechargeable() interface{}`

GetRechargeable returns the Rechargeable field if non-nil, zero value otherwise.

### GetRechargeableOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetRechargeableOk() (*interface{}, bool)`

GetRechargeableOk returns a tuple with the Rechargeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeable

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetRechargeable(v interface{})`

SetRechargeable sets Rechargeable field to given value.

### HasRechargeable

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasRechargeable() bool`

HasRechargeable returns a boolean if a field has been set.

### SetRechargeableNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetRechargeableNil(b bool)`

 SetRechargeableNil sets the value for Rechargeable to be an explicit nil

### UnsetRechargeable
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetRechargeable()`

UnsetRechargeable ensures that no value is present for Rechargeable, not even an explicit nil
### GetImageUrl

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetExpiresAt

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetRecipientEmail

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetRecipientEmail() interface{}`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetRecipientEmailOk() (*interface{}, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetRecipientEmail(v interface{})`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### SetRecipientEmailNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetRecipientEmailNil(b bool)`

 SetRecipientEmailNil sets the value for RecipientEmail to be an explicit nil

### UnsetRecipientEmail
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetRecipientEmail()`

UnsetRecipientEmail ensures that no value is present for RecipientEmail, not even an explicit nil
### GetCreatedAt

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


