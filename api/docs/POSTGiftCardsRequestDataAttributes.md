# POSTGiftCardsRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **interface{}** | The gift card code UUID. If not set, it&#39;s automatically generated. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**BalanceCents** | **interface{}** | The gift card balance, in cents. | 
**BalanceMaxCents** | Pointer to **interface{}** | The gift card balance max, in cents. | [optional] 
**SingleUse** | Pointer to **interface{}** | Indicates if the gift card can be used only one. | [optional] 
**Rechargeable** | Pointer to **interface{}** | Indicates if the gift card can be recharged. | [optional] 
**ImageUrl** | Pointer to **interface{}** | The URL of an image that represents the gift card. | [optional] 
**ExpiresAt** | Pointer to **interface{}** | Time at which the gift card will expire. | [optional] 
**RecipientEmail** | Pointer to **interface{}** | The email address of the associated recipient. When creating or updating a gift card, this is a shortcut to find or create the associated recipient by email. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTGiftCardsRequestDataAttributes

`func NewPOSTGiftCardsRequestDataAttributes(balanceCents interface{}, ) *POSTGiftCardsRequestDataAttributes`

NewPOSTGiftCardsRequestDataAttributes instantiates a new POSTGiftCardsRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTGiftCardsRequestDataAttributesWithDefaults

`func NewPOSTGiftCardsRequestDataAttributesWithDefaults() *POSTGiftCardsRequestDataAttributes`

NewPOSTGiftCardsRequestDataAttributesWithDefaults instantiates a new POSTGiftCardsRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *POSTGiftCardsRequestDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *POSTGiftCardsRequestDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *POSTGiftCardsRequestDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *POSTGiftCardsRequestDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *POSTGiftCardsRequestDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *POSTGiftCardsRequestDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCurrencyCode

`func (o *POSTGiftCardsRequestDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *POSTGiftCardsRequestDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *POSTGiftCardsRequestDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *POSTGiftCardsRequestDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *POSTGiftCardsRequestDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *POSTGiftCardsRequestDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetBalanceCents

`func (o *POSTGiftCardsRequestDataAttributes) GetBalanceCents() interface{}`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *POSTGiftCardsRequestDataAttributes) GetBalanceCentsOk() (*interface{}, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *POSTGiftCardsRequestDataAttributes) SetBalanceCents(v interface{})`

SetBalanceCents sets BalanceCents field to given value.


### SetBalanceCentsNil

`func (o *POSTGiftCardsRequestDataAttributes) SetBalanceCentsNil(b bool)`

 SetBalanceCentsNil sets the value for BalanceCents to be an explicit nil

### UnsetBalanceCents
`func (o *POSTGiftCardsRequestDataAttributes) UnsetBalanceCents()`

UnsetBalanceCents ensures that no value is present for BalanceCents, not even an explicit nil
### GetBalanceMaxCents

`func (o *POSTGiftCardsRequestDataAttributes) GetBalanceMaxCents() interface{}`

GetBalanceMaxCents returns the BalanceMaxCents field if non-nil, zero value otherwise.

### GetBalanceMaxCentsOk

`func (o *POSTGiftCardsRequestDataAttributes) GetBalanceMaxCentsOk() (*interface{}, bool)`

GetBalanceMaxCentsOk returns a tuple with the BalanceMaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceMaxCents

`func (o *POSTGiftCardsRequestDataAttributes) SetBalanceMaxCents(v interface{})`

SetBalanceMaxCents sets BalanceMaxCents field to given value.

### HasBalanceMaxCents

`func (o *POSTGiftCardsRequestDataAttributes) HasBalanceMaxCents() bool`

HasBalanceMaxCents returns a boolean if a field has been set.

### SetBalanceMaxCentsNil

`func (o *POSTGiftCardsRequestDataAttributes) SetBalanceMaxCentsNil(b bool)`

 SetBalanceMaxCentsNil sets the value for BalanceMaxCents to be an explicit nil

### UnsetBalanceMaxCents
`func (o *POSTGiftCardsRequestDataAttributes) UnsetBalanceMaxCents()`

UnsetBalanceMaxCents ensures that no value is present for BalanceMaxCents, not even an explicit nil
### GetSingleUse

`func (o *POSTGiftCardsRequestDataAttributes) GetSingleUse() interface{}`

GetSingleUse returns the SingleUse field if non-nil, zero value otherwise.

### GetSingleUseOk

`func (o *POSTGiftCardsRequestDataAttributes) GetSingleUseOk() (*interface{}, bool)`

GetSingleUseOk returns a tuple with the SingleUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUse

`func (o *POSTGiftCardsRequestDataAttributes) SetSingleUse(v interface{})`

SetSingleUse sets SingleUse field to given value.

### HasSingleUse

`func (o *POSTGiftCardsRequestDataAttributes) HasSingleUse() bool`

HasSingleUse returns a boolean if a field has been set.

### SetSingleUseNil

`func (o *POSTGiftCardsRequestDataAttributes) SetSingleUseNil(b bool)`

 SetSingleUseNil sets the value for SingleUse to be an explicit nil

### UnsetSingleUse
`func (o *POSTGiftCardsRequestDataAttributes) UnsetSingleUse()`

UnsetSingleUse ensures that no value is present for SingleUse, not even an explicit nil
### GetRechargeable

`func (o *POSTGiftCardsRequestDataAttributes) GetRechargeable() interface{}`

GetRechargeable returns the Rechargeable field if non-nil, zero value otherwise.

### GetRechargeableOk

`func (o *POSTGiftCardsRequestDataAttributes) GetRechargeableOk() (*interface{}, bool)`

GetRechargeableOk returns a tuple with the Rechargeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeable

`func (o *POSTGiftCardsRequestDataAttributes) SetRechargeable(v interface{})`

SetRechargeable sets Rechargeable field to given value.

### HasRechargeable

`func (o *POSTGiftCardsRequestDataAttributes) HasRechargeable() bool`

HasRechargeable returns a boolean if a field has been set.

### SetRechargeableNil

`func (o *POSTGiftCardsRequestDataAttributes) SetRechargeableNil(b bool)`

 SetRechargeableNil sets the value for Rechargeable to be an explicit nil

### UnsetRechargeable
`func (o *POSTGiftCardsRequestDataAttributes) UnsetRechargeable()`

UnsetRechargeable ensures that no value is present for Rechargeable, not even an explicit nil
### GetImageUrl

`func (o *POSTGiftCardsRequestDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *POSTGiftCardsRequestDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *POSTGiftCardsRequestDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *POSTGiftCardsRequestDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *POSTGiftCardsRequestDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *POSTGiftCardsRequestDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetExpiresAt

`func (o *POSTGiftCardsRequestDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *POSTGiftCardsRequestDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *POSTGiftCardsRequestDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *POSTGiftCardsRequestDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *POSTGiftCardsRequestDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *POSTGiftCardsRequestDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetRecipientEmail

`func (o *POSTGiftCardsRequestDataAttributes) GetRecipientEmail() interface{}`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *POSTGiftCardsRequestDataAttributes) GetRecipientEmailOk() (*interface{}, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *POSTGiftCardsRequestDataAttributes) SetRecipientEmail(v interface{})`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *POSTGiftCardsRequestDataAttributes) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### SetRecipientEmailNil

`func (o *POSTGiftCardsRequestDataAttributes) SetRecipientEmailNil(b bool)`

 SetRecipientEmailNil sets the value for RecipientEmail to be an explicit nil

### UnsetRecipientEmail
`func (o *POSTGiftCardsRequestDataAttributes) UnsetRecipientEmail()`

UnsetRecipientEmail ensures that no value is present for RecipientEmail, not even an explicit nil
### GetReference

`func (o *POSTGiftCardsRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTGiftCardsRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTGiftCardsRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTGiftCardsRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTGiftCardsRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTGiftCardsRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTGiftCardsRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTGiftCardsRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTGiftCardsRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTGiftCardsRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTGiftCardsRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTGiftCardsRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTGiftCardsRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTGiftCardsRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTGiftCardsRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTGiftCardsRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTGiftCardsRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTGiftCardsRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


