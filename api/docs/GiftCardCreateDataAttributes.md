# GiftCardCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The gift card code UUID. If not set, it&#39;s automatically generated. | [optional] 
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**BalanceCents** | **int32** | The gift card balance, in cents. | 
**BalanceMaxCents** | Pointer to **string** | The gift card balance max, in cents. | [optional] 
**SingleUse** | Pointer to **bool** | Indicates if the gift card can be used only one. | [optional] 
**Rechargeable** | Pointer to **bool** | Indicates if the gift card can be recharged. | [optional] 
**ImageUrl** | Pointer to **string** | The URL of an image that represents the gift card. | [optional] 
**ExpiresAt** | Pointer to **string** | Time at which the gift card will expire. | [optional] 
**RecipientEmail** | Pointer to **string** | The email address of the associated recipient. When creating or updating a gift card, this is a shortcut to find or create the associated recipient by email. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGiftCardCreateDataAttributes

`func NewGiftCardCreateDataAttributes(balanceCents int32, ) *GiftCardCreateDataAttributes`

NewGiftCardCreateDataAttributes instantiates a new GiftCardCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardCreateDataAttributesWithDefaults

`func NewGiftCardCreateDataAttributesWithDefaults() *GiftCardCreateDataAttributes`

NewGiftCardCreateDataAttributesWithDefaults instantiates a new GiftCardCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GiftCardCreateDataAttributes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GiftCardCreateDataAttributes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GiftCardCreateDataAttributes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GiftCardCreateDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GiftCardCreateDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GiftCardCreateDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GiftCardCreateDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GiftCardCreateDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetBalanceCents

`func (o *GiftCardCreateDataAttributes) GetBalanceCents() int32`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *GiftCardCreateDataAttributes) GetBalanceCentsOk() (*int32, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *GiftCardCreateDataAttributes) SetBalanceCents(v int32)`

SetBalanceCents sets BalanceCents field to given value.


### GetBalanceMaxCents

`func (o *GiftCardCreateDataAttributes) GetBalanceMaxCents() string`

GetBalanceMaxCents returns the BalanceMaxCents field if non-nil, zero value otherwise.

### GetBalanceMaxCentsOk

`func (o *GiftCardCreateDataAttributes) GetBalanceMaxCentsOk() (*string, bool)`

GetBalanceMaxCentsOk returns a tuple with the BalanceMaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceMaxCents

`func (o *GiftCardCreateDataAttributes) SetBalanceMaxCents(v string)`

SetBalanceMaxCents sets BalanceMaxCents field to given value.

### HasBalanceMaxCents

`func (o *GiftCardCreateDataAttributes) HasBalanceMaxCents() bool`

HasBalanceMaxCents returns a boolean if a field has been set.

### GetSingleUse

`func (o *GiftCardCreateDataAttributes) GetSingleUse() bool`

GetSingleUse returns the SingleUse field if non-nil, zero value otherwise.

### GetSingleUseOk

`func (o *GiftCardCreateDataAttributes) GetSingleUseOk() (*bool, bool)`

GetSingleUseOk returns a tuple with the SingleUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUse

`func (o *GiftCardCreateDataAttributes) SetSingleUse(v bool)`

SetSingleUse sets SingleUse field to given value.

### HasSingleUse

`func (o *GiftCardCreateDataAttributes) HasSingleUse() bool`

HasSingleUse returns a boolean if a field has been set.

### GetRechargeable

`func (o *GiftCardCreateDataAttributes) GetRechargeable() bool`

GetRechargeable returns the Rechargeable field if non-nil, zero value otherwise.

### GetRechargeableOk

`func (o *GiftCardCreateDataAttributes) GetRechargeableOk() (*bool, bool)`

GetRechargeableOk returns a tuple with the Rechargeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeable

`func (o *GiftCardCreateDataAttributes) SetRechargeable(v bool)`

SetRechargeable sets Rechargeable field to given value.

### HasRechargeable

`func (o *GiftCardCreateDataAttributes) HasRechargeable() bool`

HasRechargeable returns a boolean if a field has been set.

### GetImageUrl

`func (o *GiftCardCreateDataAttributes) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GiftCardCreateDataAttributes) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GiftCardCreateDataAttributes) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GiftCardCreateDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GiftCardCreateDataAttributes) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GiftCardCreateDataAttributes) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GiftCardCreateDataAttributes) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GiftCardCreateDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetRecipientEmail

`func (o *GiftCardCreateDataAttributes) GetRecipientEmail() string`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *GiftCardCreateDataAttributes) GetRecipientEmailOk() (*string, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *GiftCardCreateDataAttributes) SetRecipientEmail(v string)`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *GiftCardCreateDataAttributes) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### GetReference

`func (o *GiftCardCreateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GiftCardCreateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GiftCardCreateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GiftCardCreateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GiftCardCreateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GiftCardCreateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GiftCardCreateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GiftCardCreateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GiftCardCreateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GiftCardCreateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GiftCardCreateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GiftCardCreateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


