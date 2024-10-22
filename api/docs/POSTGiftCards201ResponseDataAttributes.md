# POSTGiftCards201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **interface{}** | The gift card code UUID. If not set, it&#39;s automatically generated. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**BalanceCents** | **interface{}** | The gift card balance, in cents. | 
**BalanceMaxCents** | Pointer to **interface{}** | The gift card balance max, in cents. | [optional] 
**SingleUse** | Pointer to **interface{}** | Indicates if the gift card can be used only one. | [optional] 
**Rechargeable** | Pointer to **interface{}** | Indicates if the gift card can be recharged. | [optional] 
**DistributeDiscount** | Pointer to **interface{}** | Indicates if redeemed gift card amount is distributed for tax calculation. | [optional] 
**ImageUrl** | Pointer to **interface{}** | The URL of an image that represents the gift card. | [optional] 
**ExpiresAt** | Pointer to **interface{}** | Time at which the gift card will expire. | [optional] 
**RecipientEmail** | Pointer to **interface{}** | The email address of the associated recipient. When creating or updating a gift card, this is a shortcut to find or create the associated recipient by email. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTGiftCards201ResponseDataAttributes

`func NewPOSTGiftCards201ResponseDataAttributes(balanceCents interface{}, ) *POSTGiftCards201ResponseDataAttributes`

NewPOSTGiftCards201ResponseDataAttributes instantiates a new POSTGiftCards201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTGiftCards201ResponseDataAttributesWithDefaults

`func NewPOSTGiftCards201ResponseDataAttributesWithDefaults() *POSTGiftCards201ResponseDataAttributes`

NewPOSTGiftCards201ResponseDataAttributesWithDefaults instantiates a new POSTGiftCards201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *POSTGiftCards201ResponseDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *POSTGiftCards201ResponseDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *POSTGiftCards201ResponseDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *POSTGiftCards201ResponseDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *POSTGiftCards201ResponseDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *POSTGiftCards201ResponseDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCurrencyCode

`func (o *POSTGiftCards201ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *POSTGiftCards201ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *POSTGiftCards201ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *POSTGiftCards201ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *POSTGiftCards201ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *POSTGiftCards201ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetBalanceCents

`func (o *POSTGiftCards201ResponseDataAttributes) GetBalanceCents() interface{}`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *POSTGiftCards201ResponseDataAttributes) GetBalanceCentsOk() (*interface{}, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *POSTGiftCards201ResponseDataAttributes) SetBalanceCents(v interface{})`

SetBalanceCents sets BalanceCents field to given value.


### SetBalanceCentsNil

`func (o *POSTGiftCards201ResponseDataAttributes) SetBalanceCentsNil(b bool)`

 SetBalanceCentsNil sets the value for BalanceCents to be an explicit nil

### UnsetBalanceCents
`func (o *POSTGiftCards201ResponseDataAttributes) UnsetBalanceCents()`

UnsetBalanceCents ensures that no value is present for BalanceCents, not even an explicit nil
### GetBalanceMaxCents

`func (o *POSTGiftCards201ResponseDataAttributes) GetBalanceMaxCents() interface{}`

GetBalanceMaxCents returns the BalanceMaxCents field if non-nil, zero value otherwise.

### GetBalanceMaxCentsOk

`func (o *POSTGiftCards201ResponseDataAttributes) GetBalanceMaxCentsOk() (*interface{}, bool)`

GetBalanceMaxCentsOk returns a tuple with the BalanceMaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceMaxCents

`func (o *POSTGiftCards201ResponseDataAttributes) SetBalanceMaxCents(v interface{})`

SetBalanceMaxCents sets BalanceMaxCents field to given value.

### HasBalanceMaxCents

`func (o *POSTGiftCards201ResponseDataAttributes) HasBalanceMaxCents() bool`

HasBalanceMaxCents returns a boolean if a field has been set.

### SetBalanceMaxCentsNil

`func (o *POSTGiftCards201ResponseDataAttributes) SetBalanceMaxCentsNil(b bool)`

 SetBalanceMaxCentsNil sets the value for BalanceMaxCents to be an explicit nil

### UnsetBalanceMaxCents
`func (o *POSTGiftCards201ResponseDataAttributes) UnsetBalanceMaxCents()`

UnsetBalanceMaxCents ensures that no value is present for BalanceMaxCents, not even an explicit nil
### GetSingleUse

`func (o *POSTGiftCards201ResponseDataAttributes) GetSingleUse() interface{}`

GetSingleUse returns the SingleUse field if non-nil, zero value otherwise.

### GetSingleUseOk

`func (o *POSTGiftCards201ResponseDataAttributes) GetSingleUseOk() (*interface{}, bool)`

GetSingleUseOk returns a tuple with the SingleUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUse

`func (o *POSTGiftCards201ResponseDataAttributes) SetSingleUse(v interface{})`

SetSingleUse sets SingleUse field to given value.

### HasSingleUse

`func (o *POSTGiftCards201ResponseDataAttributes) HasSingleUse() bool`

HasSingleUse returns a boolean if a field has been set.

### SetSingleUseNil

`func (o *POSTGiftCards201ResponseDataAttributes) SetSingleUseNil(b bool)`

 SetSingleUseNil sets the value for SingleUse to be an explicit nil

### UnsetSingleUse
`func (o *POSTGiftCards201ResponseDataAttributes) UnsetSingleUse()`

UnsetSingleUse ensures that no value is present for SingleUse, not even an explicit nil
### GetRechargeable

`func (o *POSTGiftCards201ResponseDataAttributes) GetRechargeable() interface{}`

GetRechargeable returns the Rechargeable field if non-nil, zero value otherwise.

### GetRechargeableOk

`func (o *POSTGiftCards201ResponseDataAttributes) GetRechargeableOk() (*interface{}, bool)`

GetRechargeableOk returns a tuple with the Rechargeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeable

`func (o *POSTGiftCards201ResponseDataAttributes) SetRechargeable(v interface{})`

SetRechargeable sets Rechargeable field to given value.

### HasRechargeable

`func (o *POSTGiftCards201ResponseDataAttributes) HasRechargeable() bool`

HasRechargeable returns a boolean if a field has been set.

### SetRechargeableNil

`func (o *POSTGiftCards201ResponseDataAttributes) SetRechargeableNil(b bool)`

 SetRechargeableNil sets the value for Rechargeable to be an explicit nil

### UnsetRechargeable
`func (o *POSTGiftCards201ResponseDataAttributes) UnsetRechargeable()`

UnsetRechargeable ensures that no value is present for Rechargeable, not even an explicit nil
### GetDistributeDiscount

`func (o *POSTGiftCards201ResponseDataAttributes) GetDistributeDiscount() interface{}`

GetDistributeDiscount returns the DistributeDiscount field if non-nil, zero value otherwise.

### GetDistributeDiscountOk

`func (o *POSTGiftCards201ResponseDataAttributes) GetDistributeDiscountOk() (*interface{}, bool)`

GetDistributeDiscountOk returns a tuple with the DistributeDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributeDiscount

`func (o *POSTGiftCards201ResponseDataAttributes) SetDistributeDiscount(v interface{})`

SetDistributeDiscount sets DistributeDiscount field to given value.

### HasDistributeDiscount

`func (o *POSTGiftCards201ResponseDataAttributes) HasDistributeDiscount() bool`

HasDistributeDiscount returns a boolean if a field has been set.

### SetDistributeDiscountNil

`func (o *POSTGiftCards201ResponseDataAttributes) SetDistributeDiscountNil(b bool)`

 SetDistributeDiscountNil sets the value for DistributeDiscount to be an explicit nil

### UnsetDistributeDiscount
`func (o *POSTGiftCards201ResponseDataAttributes) UnsetDistributeDiscount()`

UnsetDistributeDiscount ensures that no value is present for DistributeDiscount, not even an explicit nil
### GetImageUrl

`func (o *POSTGiftCards201ResponseDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *POSTGiftCards201ResponseDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *POSTGiftCards201ResponseDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *POSTGiftCards201ResponseDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *POSTGiftCards201ResponseDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *POSTGiftCards201ResponseDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetExpiresAt

`func (o *POSTGiftCards201ResponseDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *POSTGiftCards201ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *POSTGiftCards201ResponseDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *POSTGiftCards201ResponseDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *POSTGiftCards201ResponseDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *POSTGiftCards201ResponseDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetRecipientEmail

`func (o *POSTGiftCards201ResponseDataAttributes) GetRecipientEmail() interface{}`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *POSTGiftCards201ResponseDataAttributes) GetRecipientEmailOk() (*interface{}, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *POSTGiftCards201ResponseDataAttributes) SetRecipientEmail(v interface{})`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *POSTGiftCards201ResponseDataAttributes) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### SetRecipientEmailNil

`func (o *POSTGiftCards201ResponseDataAttributes) SetRecipientEmailNil(b bool)`

 SetRecipientEmailNil sets the value for RecipientEmail to be an explicit nil

### UnsetRecipientEmail
`func (o *POSTGiftCards201ResponseDataAttributes) UnsetRecipientEmail()`

UnsetRecipientEmail ensures that no value is present for RecipientEmail, not even an explicit nil
### GetReference

`func (o *POSTGiftCards201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTGiftCards201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTGiftCards201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTGiftCards201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTGiftCards201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTGiftCards201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTGiftCards201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTGiftCards201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTGiftCards201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTGiftCards201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTGiftCards201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTGiftCards201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTGiftCards201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTGiftCards201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTGiftCards201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTGiftCards201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTGiftCards201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTGiftCards201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


