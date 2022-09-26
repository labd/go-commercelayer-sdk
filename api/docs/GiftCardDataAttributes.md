# GiftCardDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The gift card status, one of &#39;draft&#39;, &#39;inactive&#39;, &#39;active&#39;, or &#39;redeemed&#39;. | [optional] 
**Code** | Pointer to **string** | The gift card code UUID. If not set, it&#39;s automatically generated. | [optional] 
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**InitialBalanceCents** | Pointer to **int32** | The gift card initial balance, in cents. | [optional] 
**InitialBalanceFloat** | Pointer to **float32** | The gift card initial balance, float. | [optional] 
**FormattedInitialBalance** | Pointer to **string** | The gift card initial balance, formatted. | [optional] 
**BalanceCents** | Pointer to **int32** | The gift card balance, in cents. | [optional] 
**BalanceFloat** | Pointer to **float32** | The gift card balance, float. | [optional] 
**FormattedBalance** | Pointer to **string** | The gift card balance, formatted. | [optional] 
**BalanceMaxCents** | Pointer to **string** | The gift card balance max, in cents. | [optional] 
**BalanceMaxFloat** | Pointer to **float32** | The gift card balance max, float. | [optional] 
**FormattedBalanceMax** | Pointer to **string** | The gift card balance max, formatted. | [optional] 
**BalanceLog** | Pointer to **[]map[string]interface{}** | The gift card balance log. Tracks all the gift card transactions. | [optional] 
**SingleUse** | Pointer to **bool** | Indicates if the gift card can be used only one. | [optional] 
**Rechargeable** | Pointer to **bool** | Indicates if the gift card can be recharged. | [optional] 
**ImageUrl** | Pointer to **string** | The URL of an image that represents the gift card. | [optional] 
**ExpiresAt** | Pointer to **string** | Time at which the gift card will expire. | [optional] 
**RecipientEmail** | Pointer to **string** | The email address of the associated recipient. When creating or updating a gift card, this is a shortcut to find or create the associated recipient by email. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGiftCardDataAttributes

`func NewGiftCardDataAttributes() *GiftCardDataAttributes`

NewGiftCardDataAttributes instantiates a new GiftCardDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardDataAttributesWithDefaults

`func NewGiftCardDataAttributesWithDefaults() *GiftCardDataAttributes`

NewGiftCardDataAttributesWithDefaults instantiates a new GiftCardDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GiftCardDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GiftCardDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GiftCardDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GiftCardDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *GiftCardDataAttributes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GiftCardDataAttributes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GiftCardDataAttributes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GiftCardDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GiftCardDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GiftCardDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GiftCardDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GiftCardDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetInitialBalanceCents

`func (o *GiftCardDataAttributes) GetInitialBalanceCents() int32`

GetInitialBalanceCents returns the InitialBalanceCents field if non-nil, zero value otherwise.

### GetInitialBalanceCentsOk

`func (o *GiftCardDataAttributes) GetInitialBalanceCentsOk() (*int32, bool)`

GetInitialBalanceCentsOk returns a tuple with the InitialBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialBalanceCents

`func (o *GiftCardDataAttributes) SetInitialBalanceCents(v int32)`

SetInitialBalanceCents sets InitialBalanceCents field to given value.

### HasInitialBalanceCents

`func (o *GiftCardDataAttributes) HasInitialBalanceCents() bool`

HasInitialBalanceCents returns a boolean if a field has been set.

### GetInitialBalanceFloat

`func (o *GiftCardDataAttributes) GetInitialBalanceFloat() float32`

GetInitialBalanceFloat returns the InitialBalanceFloat field if non-nil, zero value otherwise.

### GetInitialBalanceFloatOk

`func (o *GiftCardDataAttributes) GetInitialBalanceFloatOk() (*float32, bool)`

GetInitialBalanceFloatOk returns a tuple with the InitialBalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialBalanceFloat

`func (o *GiftCardDataAttributes) SetInitialBalanceFloat(v float32)`

SetInitialBalanceFloat sets InitialBalanceFloat field to given value.

### HasInitialBalanceFloat

`func (o *GiftCardDataAttributes) HasInitialBalanceFloat() bool`

HasInitialBalanceFloat returns a boolean if a field has been set.

### GetFormattedInitialBalance

`func (o *GiftCardDataAttributes) GetFormattedInitialBalance() string`

GetFormattedInitialBalance returns the FormattedInitialBalance field if non-nil, zero value otherwise.

### GetFormattedInitialBalanceOk

`func (o *GiftCardDataAttributes) GetFormattedInitialBalanceOk() (*string, bool)`

GetFormattedInitialBalanceOk returns a tuple with the FormattedInitialBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedInitialBalance

`func (o *GiftCardDataAttributes) SetFormattedInitialBalance(v string)`

SetFormattedInitialBalance sets FormattedInitialBalance field to given value.

### HasFormattedInitialBalance

`func (o *GiftCardDataAttributes) HasFormattedInitialBalance() bool`

HasFormattedInitialBalance returns a boolean if a field has been set.

### GetBalanceCents

`func (o *GiftCardDataAttributes) GetBalanceCents() int32`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *GiftCardDataAttributes) GetBalanceCentsOk() (*int32, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *GiftCardDataAttributes) SetBalanceCents(v int32)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *GiftCardDataAttributes) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### GetBalanceFloat

`func (o *GiftCardDataAttributes) GetBalanceFloat() float32`

GetBalanceFloat returns the BalanceFloat field if non-nil, zero value otherwise.

### GetBalanceFloatOk

`func (o *GiftCardDataAttributes) GetBalanceFloatOk() (*float32, bool)`

GetBalanceFloatOk returns a tuple with the BalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceFloat

`func (o *GiftCardDataAttributes) SetBalanceFloat(v float32)`

SetBalanceFloat sets BalanceFloat field to given value.

### HasBalanceFloat

`func (o *GiftCardDataAttributes) HasBalanceFloat() bool`

HasBalanceFloat returns a boolean if a field has been set.

### GetFormattedBalance

`func (o *GiftCardDataAttributes) GetFormattedBalance() string`

GetFormattedBalance returns the FormattedBalance field if non-nil, zero value otherwise.

### GetFormattedBalanceOk

`func (o *GiftCardDataAttributes) GetFormattedBalanceOk() (*string, bool)`

GetFormattedBalanceOk returns a tuple with the FormattedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedBalance

`func (o *GiftCardDataAttributes) SetFormattedBalance(v string)`

SetFormattedBalance sets FormattedBalance field to given value.

### HasFormattedBalance

`func (o *GiftCardDataAttributes) HasFormattedBalance() bool`

HasFormattedBalance returns a boolean if a field has been set.

### GetBalanceMaxCents

`func (o *GiftCardDataAttributes) GetBalanceMaxCents() string`

GetBalanceMaxCents returns the BalanceMaxCents field if non-nil, zero value otherwise.

### GetBalanceMaxCentsOk

`func (o *GiftCardDataAttributes) GetBalanceMaxCentsOk() (*string, bool)`

GetBalanceMaxCentsOk returns a tuple with the BalanceMaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceMaxCents

`func (o *GiftCardDataAttributes) SetBalanceMaxCents(v string)`

SetBalanceMaxCents sets BalanceMaxCents field to given value.

### HasBalanceMaxCents

`func (o *GiftCardDataAttributes) HasBalanceMaxCents() bool`

HasBalanceMaxCents returns a boolean if a field has been set.

### GetBalanceMaxFloat

`func (o *GiftCardDataAttributes) GetBalanceMaxFloat() float32`

GetBalanceMaxFloat returns the BalanceMaxFloat field if non-nil, zero value otherwise.

### GetBalanceMaxFloatOk

`func (o *GiftCardDataAttributes) GetBalanceMaxFloatOk() (*float32, bool)`

GetBalanceMaxFloatOk returns a tuple with the BalanceMaxFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceMaxFloat

`func (o *GiftCardDataAttributes) SetBalanceMaxFloat(v float32)`

SetBalanceMaxFloat sets BalanceMaxFloat field to given value.

### HasBalanceMaxFloat

`func (o *GiftCardDataAttributes) HasBalanceMaxFloat() bool`

HasBalanceMaxFloat returns a boolean if a field has been set.

### GetFormattedBalanceMax

`func (o *GiftCardDataAttributes) GetFormattedBalanceMax() string`

GetFormattedBalanceMax returns the FormattedBalanceMax field if non-nil, zero value otherwise.

### GetFormattedBalanceMaxOk

`func (o *GiftCardDataAttributes) GetFormattedBalanceMaxOk() (*string, bool)`

GetFormattedBalanceMaxOk returns a tuple with the FormattedBalanceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedBalanceMax

`func (o *GiftCardDataAttributes) SetFormattedBalanceMax(v string)`

SetFormattedBalanceMax sets FormattedBalanceMax field to given value.

### HasFormattedBalanceMax

`func (o *GiftCardDataAttributes) HasFormattedBalanceMax() bool`

HasFormattedBalanceMax returns a boolean if a field has been set.

### GetBalanceLog

`func (o *GiftCardDataAttributes) GetBalanceLog() []map[string]interface{}`

GetBalanceLog returns the BalanceLog field if non-nil, zero value otherwise.

### GetBalanceLogOk

`func (o *GiftCardDataAttributes) GetBalanceLogOk() (*[]map[string]interface{}, bool)`

GetBalanceLogOk returns a tuple with the BalanceLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceLog

`func (o *GiftCardDataAttributes) SetBalanceLog(v []map[string]interface{})`

SetBalanceLog sets BalanceLog field to given value.

### HasBalanceLog

`func (o *GiftCardDataAttributes) HasBalanceLog() bool`

HasBalanceLog returns a boolean if a field has been set.

### GetSingleUse

`func (o *GiftCardDataAttributes) GetSingleUse() bool`

GetSingleUse returns the SingleUse field if non-nil, zero value otherwise.

### GetSingleUseOk

`func (o *GiftCardDataAttributes) GetSingleUseOk() (*bool, bool)`

GetSingleUseOk returns a tuple with the SingleUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUse

`func (o *GiftCardDataAttributes) SetSingleUse(v bool)`

SetSingleUse sets SingleUse field to given value.

### HasSingleUse

`func (o *GiftCardDataAttributes) HasSingleUse() bool`

HasSingleUse returns a boolean if a field has been set.

### GetRechargeable

`func (o *GiftCardDataAttributes) GetRechargeable() bool`

GetRechargeable returns the Rechargeable field if non-nil, zero value otherwise.

### GetRechargeableOk

`func (o *GiftCardDataAttributes) GetRechargeableOk() (*bool, bool)`

GetRechargeableOk returns a tuple with the Rechargeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeable

`func (o *GiftCardDataAttributes) SetRechargeable(v bool)`

SetRechargeable sets Rechargeable field to given value.

### HasRechargeable

`func (o *GiftCardDataAttributes) HasRechargeable() bool`

HasRechargeable returns a boolean if a field has been set.

### GetImageUrl

`func (o *GiftCardDataAttributes) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GiftCardDataAttributes) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GiftCardDataAttributes) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GiftCardDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GiftCardDataAttributes) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GiftCardDataAttributes) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GiftCardDataAttributes) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GiftCardDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetRecipientEmail

`func (o *GiftCardDataAttributes) GetRecipientEmail() string`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *GiftCardDataAttributes) GetRecipientEmailOk() (*string, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *GiftCardDataAttributes) SetRecipientEmail(v string)`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *GiftCardDataAttributes) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### GetId

`func (o *GiftCardDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GiftCardDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GiftCardDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GiftCardDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GiftCardDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GiftCardDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GiftCardDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GiftCardDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GiftCardDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GiftCardDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GiftCardDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GiftCardDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GiftCardDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GiftCardDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GiftCardDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GiftCardDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GiftCardDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GiftCardDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GiftCardDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GiftCardDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GiftCardDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GiftCardDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GiftCardDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GiftCardDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


