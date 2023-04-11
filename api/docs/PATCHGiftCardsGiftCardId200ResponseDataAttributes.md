# PATCHGiftCardsGiftCardId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**BalanceCents** | Pointer to **interface{}** | The gift card balance, in cents. | [optional] 
**BalanceMaxCents** | Pointer to **interface{}** | The gift card balance max, in cents. | [optional] 
**SingleUse** | Pointer to **interface{}** | Indicates if the gift card can be used only one. | [optional] 
**Rechargeable** | Pointer to **interface{}** | Indicates if the gift card can be recharged. | [optional] 
**ImageUrl** | Pointer to **interface{}** | The URL of an image that represents the gift card. | [optional] 
**ExpiresAt** | Pointer to **interface{}** | Time at which the gift card will expire. | [optional] 
**RecipientEmail** | Pointer to **interface{}** | The email address of the associated recipient. When creating or updating a gift card, this is a shortcut to find or create the associated recipient by email. | [optional] 
**Purchase** | Pointer to **interface{}** | Send this attribute if you want to confirm a draft gift card. The gift card becomes &#39;inactive&#39;, waiting to be activated. | [optional] 
**Activate** | Pointer to **interface{}** | Send this attribute if you want to activate a gift card. | [optional] 
**Deactivate** | Pointer to **interface{}** | Send this attribute if you want to deactivate a gift card. | [optional] 
**BalanceChangeCents** | Pointer to **interface{}** | The balance change, in cents. Send a negative value to reduces the card balance by the specified amount. Send a positive value to recharge the gift card (if rechargeable). | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetBalanceCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceCents() interface{}`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceCentsOk() (*interface{}, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceCents(v interface{})`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### SetBalanceCentsNil

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceCentsNil(b bool)`

 SetBalanceCentsNil sets the value for BalanceCents to be an explicit nil

### UnsetBalanceCents
`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) UnsetBalanceCents()`

UnsetBalanceCents ensures that no value is present for BalanceCents, not even an explicit nil
### GetBalanceMaxCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceMaxCents() interface{}`

GetBalanceMaxCents returns the BalanceMaxCents field if non-nil, zero value otherwise.

### GetBalanceMaxCentsOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceMaxCentsOk() (*interface{}, bool)`

GetBalanceMaxCentsOk returns a tuple with the BalanceMaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceMaxCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceMaxCents(v interface{})`

SetBalanceMaxCents sets BalanceMaxCents field to given value.

### HasBalanceMaxCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceMaxCents() bool`

HasBalanceMaxCents returns a boolean if a field has been set.

### SetBalanceMaxCentsNil

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceMaxCentsNil(b bool)`

 SetBalanceMaxCentsNil sets the value for BalanceMaxCents to be an explicit nil

### UnsetBalanceMaxCents
`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) UnsetBalanceMaxCents()`

UnsetBalanceMaxCents ensures that no value is present for BalanceMaxCents, not even an explicit nil
### GetSingleUse

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetSingleUse() interface{}`

GetSingleUse returns the SingleUse field if non-nil, zero value otherwise.

### GetSingleUseOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetSingleUseOk() (*interface{}, bool)`

GetSingleUseOk returns a tuple with the SingleUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUse

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetSingleUse(v interface{})`

SetSingleUse sets SingleUse field to given value.

### HasSingleUse

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasSingleUse() bool`

HasSingleUse returns a boolean if a field has been set.

### SetSingleUseNil

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetSingleUseNil(b bool)`

 SetSingleUseNil sets the value for SingleUse to be an explicit nil

### UnsetSingleUse
`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) UnsetSingleUse()`

UnsetSingleUse ensures that no value is present for SingleUse, not even an explicit nil
### GetRechargeable

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetRechargeable() interface{}`

GetRechargeable returns the Rechargeable field if non-nil, zero value otherwise.

### GetRechargeableOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetRechargeableOk() (*interface{}, bool)`

GetRechargeableOk returns a tuple with the Rechargeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeable

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetRechargeable(v interface{})`

SetRechargeable sets Rechargeable field to given value.

### HasRechargeable

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasRechargeable() bool`

HasRechargeable returns a boolean if a field has been set.

### SetRechargeableNil

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetRechargeableNil(b bool)`

 SetRechargeableNil sets the value for Rechargeable to be an explicit nil

### UnsetRechargeable
`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) UnsetRechargeable()`

UnsetRechargeable ensures that no value is present for Rechargeable, not even an explicit nil
### GetImageUrl

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetExpiresAt

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetRecipientEmail

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetRecipientEmail() interface{}`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetRecipientEmailOk() (*interface{}, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetRecipientEmail(v interface{})`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### SetRecipientEmailNil

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetRecipientEmailNil(b bool)`

 SetRecipientEmailNil sets the value for RecipientEmail to be an explicit nil

### UnsetRecipientEmail
`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) UnsetRecipientEmail()`

UnsetRecipientEmail ensures that no value is present for RecipientEmail, not even an explicit nil
### GetPurchase

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetPurchase() interface{}`

GetPurchase returns the Purchase field if non-nil, zero value otherwise.

### GetPurchaseOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetPurchaseOk() (*interface{}, bool)`

GetPurchaseOk returns a tuple with the Purchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchase

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetPurchase(v interface{})`

SetPurchase sets Purchase field to given value.

### HasPurchase

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasPurchase() bool`

HasPurchase returns a boolean if a field has been set.

### SetPurchaseNil

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetPurchaseNil(b bool)`

 SetPurchaseNil sets the value for Purchase to be an explicit nil

### UnsetPurchase
`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) UnsetPurchase()`

UnsetPurchase ensures that no value is present for Purchase, not even an explicit nil
### GetActivate

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetActivate() interface{}`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetActivateOk() (*interface{}, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetActivate(v interface{})`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### SetActivateNil

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetActivateNil(b bool)`

 SetActivateNil sets the value for Activate to be an explicit nil

### UnsetActivate
`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) UnsetActivate()`

UnsetActivate ensures that no value is present for Activate, not even an explicit nil
### GetDeactivate

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetDeactivate() interface{}`

GetDeactivate returns the Deactivate field if non-nil, zero value otherwise.

### GetDeactivateOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetDeactivateOk() (*interface{}, bool)`

GetDeactivateOk returns a tuple with the Deactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivate

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetDeactivate(v interface{})`

SetDeactivate sets Deactivate field to given value.

### HasDeactivate

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasDeactivate() bool`

HasDeactivate returns a boolean if a field has been set.

### SetDeactivateNil

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetDeactivateNil(b bool)`

 SetDeactivateNil sets the value for Deactivate to be an explicit nil

### UnsetDeactivate
`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) UnsetDeactivate()`

UnsetDeactivate ensures that no value is present for Deactivate, not even an explicit nil
### GetBalanceChangeCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceChangeCents() interface{}`

GetBalanceChangeCents returns the BalanceChangeCents field if non-nil, zero value otherwise.

### GetBalanceChangeCentsOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceChangeCentsOk() (*interface{}, bool)`

GetBalanceChangeCentsOk returns a tuple with the BalanceChangeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceChangeCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceChangeCents(v interface{})`

SetBalanceChangeCents sets BalanceChangeCents field to given value.

### HasBalanceChangeCents

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceChangeCents() bool`

HasBalanceChangeCents returns a boolean if a field has been set.

### SetBalanceChangeCentsNil

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceChangeCentsNil(b bool)`

 SetBalanceChangeCentsNil sets the value for BalanceChangeCents to be an explicit nil

### UnsetBalanceChangeCents
`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) UnsetBalanceChangeCents()`

UnsetBalanceChangeCents ensures that no value is present for BalanceChangeCents, not even an explicit nil
### GetReference

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


