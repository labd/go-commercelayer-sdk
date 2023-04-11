# PATCHGiftCardsGiftCardIdRequestDataAttributes

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

### NewPATCHGiftCardsGiftCardIdRequestDataAttributes

`func NewPATCHGiftCardsGiftCardIdRequestDataAttributes() *PATCHGiftCardsGiftCardIdRequestDataAttributes`

NewPATCHGiftCardsGiftCardIdRequestDataAttributes instantiates a new PATCHGiftCardsGiftCardIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHGiftCardsGiftCardIdRequestDataAttributesWithDefaults

`func NewPATCHGiftCardsGiftCardIdRequestDataAttributesWithDefaults() *PATCHGiftCardsGiftCardIdRequestDataAttributes`

NewPATCHGiftCardsGiftCardIdRequestDataAttributesWithDefaults instantiates a new PATCHGiftCardsGiftCardIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetBalanceCents

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetBalanceCents() interface{}`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetBalanceCentsOk() (*interface{}, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetBalanceCents(v interface{})`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### SetBalanceCentsNil

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetBalanceCentsNil(b bool)`

 SetBalanceCentsNil sets the value for BalanceCents to be an explicit nil

### UnsetBalanceCents
`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) UnsetBalanceCents()`

UnsetBalanceCents ensures that no value is present for BalanceCents, not even an explicit nil
### GetBalanceMaxCents

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetBalanceMaxCents() interface{}`

GetBalanceMaxCents returns the BalanceMaxCents field if non-nil, zero value otherwise.

### GetBalanceMaxCentsOk

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetBalanceMaxCentsOk() (*interface{}, bool)`

GetBalanceMaxCentsOk returns a tuple with the BalanceMaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceMaxCents

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetBalanceMaxCents(v interface{})`

SetBalanceMaxCents sets BalanceMaxCents field to given value.

### HasBalanceMaxCents

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) HasBalanceMaxCents() bool`

HasBalanceMaxCents returns a boolean if a field has been set.

### SetBalanceMaxCentsNil

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetBalanceMaxCentsNil(b bool)`

 SetBalanceMaxCentsNil sets the value for BalanceMaxCents to be an explicit nil

### UnsetBalanceMaxCents
`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) UnsetBalanceMaxCents()`

UnsetBalanceMaxCents ensures that no value is present for BalanceMaxCents, not even an explicit nil
### GetSingleUse

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetSingleUse() interface{}`

GetSingleUse returns the SingleUse field if non-nil, zero value otherwise.

### GetSingleUseOk

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetSingleUseOk() (*interface{}, bool)`

GetSingleUseOk returns a tuple with the SingleUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUse

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetSingleUse(v interface{})`

SetSingleUse sets SingleUse field to given value.

### HasSingleUse

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) HasSingleUse() bool`

HasSingleUse returns a boolean if a field has been set.

### SetSingleUseNil

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetSingleUseNil(b bool)`

 SetSingleUseNil sets the value for SingleUse to be an explicit nil

### UnsetSingleUse
`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) UnsetSingleUse()`

UnsetSingleUse ensures that no value is present for SingleUse, not even an explicit nil
### GetRechargeable

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetRechargeable() interface{}`

GetRechargeable returns the Rechargeable field if non-nil, zero value otherwise.

### GetRechargeableOk

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetRechargeableOk() (*interface{}, bool)`

GetRechargeableOk returns a tuple with the Rechargeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeable

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetRechargeable(v interface{})`

SetRechargeable sets Rechargeable field to given value.

### HasRechargeable

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) HasRechargeable() bool`

HasRechargeable returns a boolean if a field has been set.

### SetRechargeableNil

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetRechargeableNil(b bool)`

 SetRechargeableNil sets the value for Rechargeable to be an explicit nil

### UnsetRechargeable
`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) UnsetRechargeable()`

UnsetRechargeable ensures that no value is present for Rechargeable, not even an explicit nil
### GetImageUrl

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetExpiresAt

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetRecipientEmail

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetRecipientEmail() interface{}`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetRecipientEmailOk() (*interface{}, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetRecipientEmail(v interface{})`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### SetRecipientEmailNil

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetRecipientEmailNil(b bool)`

 SetRecipientEmailNil sets the value for RecipientEmail to be an explicit nil

### UnsetRecipientEmail
`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) UnsetRecipientEmail()`

UnsetRecipientEmail ensures that no value is present for RecipientEmail, not even an explicit nil
### GetPurchase

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetPurchase() interface{}`

GetPurchase returns the Purchase field if non-nil, zero value otherwise.

### GetPurchaseOk

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetPurchaseOk() (*interface{}, bool)`

GetPurchaseOk returns a tuple with the Purchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchase

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetPurchase(v interface{})`

SetPurchase sets Purchase field to given value.

### HasPurchase

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) HasPurchase() bool`

HasPurchase returns a boolean if a field has been set.

### SetPurchaseNil

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetPurchaseNil(b bool)`

 SetPurchaseNil sets the value for Purchase to be an explicit nil

### UnsetPurchase
`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) UnsetPurchase()`

UnsetPurchase ensures that no value is present for Purchase, not even an explicit nil
### GetActivate

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetActivate() interface{}`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetActivateOk() (*interface{}, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetActivate(v interface{})`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### SetActivateNil

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetActivateNil(b bool)`

 SetActivateNil sets the value for Activate to be an explicit nil

### UnsetActivate
`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) UnsetActivate()`

UnsetActivate ensures that no value is present for Activate, not even an explicit nil
### GetDeactivate

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetDeactivate() interface{}`

GetDeactivate returns the Deactivate field if non-nil, zero value otherwise.

### GetDeactivateOk

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetDeactivateOk() (*interface{}, bool)`

GetDeactivateOk returns a tuple with the Deactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivate

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetDeactivate(v interface{})`

SetDeactivate sets Deactivate field to given value.

### HasDeactivate

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) HasDeactivate() bool`

HasDeactivate returns a boolean if a field has been set.

### SetDeactivateNil

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetDeactivateNil(b bool)`

 SetDeactivateNil sets the value for Deactivate to be an explicit nil

### UnsetDeactivate
`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) UnsetDeactivate()`

UnsetDeactivate ensures that no value is present for Deactivate, not even an explicit nil
### GetBalanceChangeCents

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetBalanceChangeCents() interface{}`

GetBalanceChangeCents returns the BalanceChangeCents field if non-nil, zero value otherwise.

### GetBalanceChangeCentsOk

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetBalanceChangeCentsOk() (*interface{}, bool)`

GetBalanceChangeCentsOk returns a tuple with the BalanceChangeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceChangeCents

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetBalanceChangeCents(v interface{})`

SetBalanceChangeCents sets BalanceChangeCents field to given value.

### HasBalanceChangeCents

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) HasBalanceChangeCents() bool`

HasBalanceChangeCents returns a boolean if a field has been set.

### SetBalanceChangeCentsNil

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetBalanceChangeCentsNil(b bool)`

 SetBalanceChangeCentsNil sets the value for BalanceChangeCents to be an explicit nil

### UnsetBalanceChangeCents
`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) UnsetBalanceChangeCents()`

UnsetBalanceChangeCents ensures that no value is present for BalanceChangeCents, not even an explicit nil
### GetReference

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHGiftCardsGiftCardIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


