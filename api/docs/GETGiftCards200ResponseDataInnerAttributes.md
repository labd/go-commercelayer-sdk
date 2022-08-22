# GETGiftCards200ResponseDataInnerAttributes

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

### NewGETGiftCards200ResponseDataInnerAttributes

`func NewGETGiftCards200ResponseDataInnerAttributes() *GETGiftCards200ResponseDataInnerAttributes`

NewGETGiftCards200ResponseDataInnerAttributes instantiates a new GETGiftCards200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETGiftCards200ResponseDataInnerAttributesWithDefaults

`func NewGETGiftCards200ResponseDataInnerAttributesWithDefaults() *GETGiftCards200ResponseDataInnerAttributes`

NewGETGiftCards200ResponseDataInnerAttributesWithDefaults instantiates a new GETGiftCards200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetInitialBalanceCents

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetInitialBalanceCents() int32`

GetInitialBalanceCents returns the InitialBalanceCents field if non-nil, zero value otherwise.

### GetInitialBalanceCentsOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetInitialBalanceCentsOk() (*int32, bool)`

GetInitialBalanceCentsOk returns a tuple with the InitialBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialBalanceCents

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetInitialBalanceCents(v int32)`

SetInitialBalanceCents sets InitialBalanceCents field to given value.

### HasInitialBalanceCents

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasInitialBalanceCents() bool`

HasInitialBalanceCents returns a boolean if a field has been set.

### GetInitialBalanceFloat

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetInitialBalanceFloat() float32`

GetInitialBalanceFloat returns the InitialBalanceFloat field if non-nil, zero value otherwise.

### GetInitialBalanceFloatOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetInitialBalanceFloatOk() (*float32, bool)`

GetInitialBalanceFloatOk returns a tuple with the InitialBalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialBalanceFloat

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetInitialBalanceFloat(v float32)`

SetInitialBalanceFloat sets InitialBalanceFloat field to given value.

### HasInitialBalanceFloat

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasInitialBalanceFloat() bool`

HasInitialBalanceFloat returns a boolean if a field has been set.

### GetFormattedInitialBalance

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetFormattedInitialBalance() string`

GetFormattedInitialBalance returns the FormattedInitialBalance field if non-nil, zero value otherwise.

### GetFormattedInitialBalanceOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetFormattedInitialBalanceOk() (*string, bool)`

GetFormattedInitialBalanceOk returns a tuple with the FormattedInitialBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedInitialBalance

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetFormattedInitialBalance(v string)`

SetFormattedInitialBalance sets FormattedInitialBalance field to given value.

### HasFormattedInitialBalance

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasFormattedInitialBalance() bool`

HasFormattedInitialBalance returns a boolean if a field has been set.

### GetBalanceCents

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetBalanceCents() int32`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetBalanceCentsOk() (*int32, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetBalanceCents(v int32)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### GetBalanceFloat

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetBalanceFloat() float32`

GetBalanceFloat returns the BalanceFloat field if non-nil, zero value otherwise.

### GetBalanceFloatOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetBalanceFloatOk() (*float32, bool)`

GetBalanceFloatOk returns a tuple with the BalanceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceFloat

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetBalanceFloat(v float32)`

SetBalanceFloat sets BalanceFloat field to given value.

### HasBalanceFloat

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasBalanceFloat() bool`

HasBalanceFloat returns a boolean if a field has been set.

### GetFormattedBalance

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetFormattedBalance() string`

GetFormattedBalance returns the FormattedBalance field if non-nil, zero value otherwise.

### GetFormattedBalanceOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetFormattedBalanceOk() (*string, bool)`

GetFormattedBalanceOk returns a tuple with the FormattedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedBalance

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetFormattedBalance(v string)`

SetFormattedBalance sets FormattedBalance field to given value.

### HasFormattedBalance

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasFormattedBalance() bool`

HasFormattedBalance returns a boolean if a field has been set.

### GetBalanceMaxCents

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetBalanceMaxCents() string`

GetBalanceMaxCents returns the BalanceMaxCents field if non-nil, zero value otherwise.

### GetBalanceMaxCentsOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetBalanceMaxCentsOk() (*string, bool)`

GetBalanceMaxCentsOk returns a tuple with the BalanceMaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceMaxCents

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetBalanceMaxCents(v string)`

SetBalanceMaxCents sets BalanceMaxCents field to given value.

### HasBalanceMaxCents

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasBalanceMaxCents() bool`

HasBalanceMaxCents returns a boolean if a field has been set.

### GetBalanceMaxFloat

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetBalanceMaxFloat() float32`

GetBalanceMaxFloat returns the BalanceMaxFloat field if non-nil, zero value otherwise.

### GetBalanceMaxFloatOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetBalanceMaxFloatOk() (*float32, bool)`

GetBalanceMaxFloatOk returns a tuple with the BalanceMaxFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceMaxFloat

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetBalanceMaxFloat(v float32)`

SetBalanceMaxFloat sets BalanceMaxFloat field to given value.

### HasBalanceMaxFloat

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasBalanceMaxFloat() bool`

HasBalanceMaxFloat returns a boolean if a field has been set.

### GetFormattedBalanceMax

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetFormattedBalanceMax() string`

GetFormattedBalanceMax returns the FormattedBalanceMax field if non-nil, zero value otherwise.

### GetFormattedBalanceMaxOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetFormattedBalanceMaxOk() (*string, bool)`

GetFormattedBalanceMaxOk returns a tuple with the FormattedBalanceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedBalanceMax

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetFormattedBalanceMax(v string)`

SetFormattedBalanceMax sets FormattedBalanceMax field to given value.

### HasFormattedBalanceMax

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasFormattedBalanceMax() bool`

HasFormattedBalanceMax returns a boolean if a field has been set.

### GetBalanceLog

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetBalanceLog() []map[string]interface{}`

GetBalanceLog returns the BalanceLog field if non-nil, zero value otherwise.

### GetBalanceLogOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetBalanceLogOk() (*[]map[string]interface{}, bool)`

GetBalanceLogOk returns a tuple with the BalanceLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceLog

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetBalanceLog(v []map[string]interface{})`

SetBalanceLog sets BalanceLog field to given value.

### HasBalanceLog

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasBalanceLog() bool`

HasBalanceLog returns a boolean if a field has been set.

### GetSingleUse

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetSingleUse() bool`

GetSingleUse returns the SingleUse field if non-nil, zero value otherwise.

### GetSingleUseOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetSingleUseOk() (*bool, bool)`

GetSingleUseOk returns a tuple with the SingleUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUse

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetSingleUse(v bool)`

SetSingleUse sets SingleUse field to given value.

### HasSingleUse

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasSingleUse() bool`

HasSingleUse returns a boolean if a field has been set.

### GetRechargeable

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetRechargeable() bool`

GetRechargeable returns the Rechargeable field if non-nil, zero value otherwise.

### GetRechargeableOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetRechargeableOk() (*bool, bool)`

GetRechargeableOk returns a tuple with the Rechargeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeable

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetRechargeable(v bool)`

SetRechargeable sets Rechargeable field to given value.

### HasRechargeable

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasRechargeable() bool`

HasRechargeable returns a boolean if a field has been set.

### GetImageUrl

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetRecipientEmail

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetRecipientEmail() string`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetRecipientEmailOk() (*string, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetRecipientEmail(v string)`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### GetId

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETGiftCards200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETGiftCards200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETGiftCards200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


