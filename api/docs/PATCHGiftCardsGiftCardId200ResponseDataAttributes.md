# PATCHGiftCardsGiftCardId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**BalanceCents** | Pointer to **int32** | The gift card balance, in cents. | [optional] 
**BalanceMaxCents** | Pointer to **string** | The gift card balance max, in cents. | [optional] 
**SingleUse** | Pointer to **bool** | Indicates if the gift card can be used only one. | [optional] 
**Rechargeable** | Pointer to **bool** | Indicates if the gift card can be recharged. | [optional] 
**ImageUrl** | Pointer to **string** | The URL of an image that represents the gift card. | [optional] 
**ExpiresAt** | Pointer to **string** | Time at which the gift card will expire. | [optional] 
**RecipientEmail** | Pointer to **string** | The email address of the associated recipient. When creating or updating a gift card, this is a shortcut to find or create the associated recipient by email. | [optional] 
**Purchase** | Pointer to **bool** | Send this attribute if you want to confirm a draft gift card. The gift card becomes &#39;inactive&#39;, waiting to be activated. | [optional] 
**Activate** | Pointer to **bool** | Send this attribute if you want to activate a gift card. | [optional] 
**Deactivate** | Pointer to **bool** | Send this attribute if you want to deactivate a gift card. | [optional] 
**BalanceChangeCents** | Pointer to **int32** | The balance change, in cents. Send a negative value to reduces the card balance by the specified amount. Send a positive value to recharge the gift card (if rechargeable). | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHGiftCardsGiftCardId200ResponseDataAttributes

`func NewPATCHGiftCardsGiftCardId200ResponseDataAttributes() *PATCHGiftCardsGiftCardId200ResponseDataAttributes`

NewPATCHGiftCardsGiftCardId200ResponseDataAttributes instantiates a new PATCHGiftCardsGiftCardId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHGiftCardsGiftCardId200ResponseDataAttributesWithDefaults

`func NewPATCHGiftCardsGiftCardId200ResponseDataAttributesWithDefaults() *PATCHGiftCardsGiftCardId200ResponseDataAttributes`

NewPATCHGiftCardsGiftCardId200ResponseDataAttributesWithDefaults instantiates a new PATCHGiftCardsGiftCardId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetBalanceCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceCents() int32`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceCentsOk() (*int32, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceCents(v int32)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### GetBalanceMaxCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceMaxCents() string`

GetBalanceMaxCents returns the BalanceMaxCents field if non-nil, zero value otherwise.

### GetBalanceMaxCentsOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceMaxCentsOk() (*string, bool)`

GetBalanceMaxCentsOk returns a tuple with the BalanceMaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceMaxCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceMaxCents(v string)`

SetBalanceMaxCents sets BalanceMaxCents field to given value.

### HasBalanceMaxCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceMaxCents() bool`

HasBalanceMaxCents returns a boolean if a field has been set.

### GetSingleUse

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetSingleUse() bool`

GetSingleUse returns the SingleUse field if non-nil, zero value otherwise.

### GetSingleUseOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetSingleUseOk() (*bool, bool)`

GetSingleUseOk returns a tuple with the SingleUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUse

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetSingleUse(v bool)`

SetSingleUse sets SingleUse field to given value.

### HasSingleUse

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasSingleUse() bool`

HasSingleUse returns a boolean if a field has been set.

### GetRechargeable

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetRechargeable() bool`

GetRechargeable returns the Rechargeable field if non-nil, zero value otherwise.

### GetRechargeableOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetRechargeableOk() (*bool, bool)`

GetRechargeableOk returns a tuple with the Rechargeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeable

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetRechargeable(v bool)`

SetRechargeable sets Rechargeable field to given value.

### HasRechargeable

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasRechargeable() bool`

HasRechargeable returns a boolean if a field has been set.

### GetImageUrl

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetRecipientEmail

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetRecipientEmail() string`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetRecipientEmailOk() (*string, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetRecipientEmail(v string)`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### GetPurchase

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetPurchase() bool`

GetPurchase returns the Purchase field if non-nil, zero value otherwise.

### GetPurchaseOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetPurchaseOk() (*bool, bool)`

GetPurchaseOk returns a tuple with the Purchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchase

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetPurchase(v bool)`

SetPurchase sets Purchase field to given value.

### HasPurchase

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasPurchase() bool`

HasPurchase returns a boolean if a field has been set.

### GetActivate

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetActivate() bool`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetActivateOk() (*bool, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetActivate(v bool)`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### GetDeactivate

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetDeactivate() bool`

GetDeactivate returns the Deactivate field if non-nil, zero value otherwise.

### GetDeactivateOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetDeactivateOk() (*bool, bool)`

GetDeactivateOk returns a tuple with the Deactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivate

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetDeactivate(v bool)`

SetDeactivate sets Deactivate field to given value.

### HasDeactivate

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasDeactivate() bool`

HasDeactivate returns a boolean if a field has been set.

### GetBalanceChangeCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceChangeCents() int32`

GetBalanceChangeCents returns the BalanceChangeCents field if non-nil, zero value otherwise.

### GetBalanceChangeCentsOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceChangeCentsOk() (*int32, bool)`

GetBalanceChangeCentsOk returns a tuple with the BalanceChangeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceChangeCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceChangeCents(v int32)`

SetBalanceChangeCents sets BalanceChangeCents field to given value.

### HasBalanceChangeCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceChangeCents() bool`

HasBalanceChangeCents returns a boolean if a field has been set.

### GetReference

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


