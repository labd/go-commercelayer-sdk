# PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The promotion&#39;s internal name. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Exclusive** | Pointer to **interface{}** | Indicates if the promotion will be applied exclusively, based on its priority score. | [optional] 
**Priority** | Pointer to **interface{}** | The priority assigned to the promotion (lower means higher priority). | [optional] 
**StartsAt** | Pointer to **interface{}** | The activation date/time of this promotion. | [optional] 
**ExpiresAt** | Pointer to **interface{}** | The expiration date/time of this promotion (must be after starts_at). | [optional] 
**TotalUsageLimit** | Pointer to **interface{}** | The total number of times this promotion can be applied. When &#39;null&#39; it means promotion can be applied infinite times. | [optional] 
**Disable** | Pointer to **interface{}** | Send this attribute if you want to mark this resource as disabled. | [optional] 
**Enable** | Pointer to **interface{}** | Send this attribute if you want to mark this resource as enabled. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**FixedAmountCents** | Pointer to **interface{}** | The discount fixed amount to be applied, in cents. | [optional] 

## Methods

### NewPATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes

`func NewPATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes() *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes`

NewPATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes instantiates a new PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributesWithDefaults

`func NewPATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributesWithDefaults() *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes`

NewPATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributesWithDefaults instantiates a new PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrencyCode

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetExclusive

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetExclusive() interface{}`

GetExclusive returns the Exclusive field if non-nil, zero value otherwise.

### GetExclusiveOk

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetExclusiveOk() (*interface{}, bool)`

GetExclusiveOk returns a tuple with the Exclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusive

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetExclusive(v interface{})`

SetExclusive sets Exclusive field to given value.

### HasExclusive

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasExclusive() bool`

HasExclusive returns a boolean if a field has been set.

### SetExclusiveNil

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetExclusiveNil(b bool)`

 SetExclusiveNil sets the value for Exclusive to be an explicit nil

### UnsetExclusive
`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) UnsetExclusive()`

UnsetExclusive ensures that no value is present for Exclusive, not even an explicit nil
### GetPriority

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetPriority() interface{}`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetPriorityOk() (*interface{}, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetPriority(v interface{})`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetStartsAt

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetStartsAt() interface{}`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetStartsAt(v interface{})`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetExpiresAt

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetTotalUsageLimit

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetTotalUsageLimit() interface{}`

GetTotalUsageLimit returns the TotalUsageLimit field if non-nil, zero value otherwise.

### GetTotalUsageLimitOk

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetTotalUsageLimitOk() (*interface{}, bool)`

GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageLimit

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetTotalUsageLimit(v interface{})`

SetTotalUsageLimit sets TotalUsageLimit field to given value.

### HasTotalUsageLimit

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasTotalUsageLimit() bool`

HasTotalUsageLimit returns a boolean if a field has been set.

### SetTotalUsageLimitNil

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetTotalUsageLimitNil(b bool)`

 SetTotalUsageLimitNil sets the value for TotalUsageLimit to be an explicit nil

### UnsetTotalUsageLimit
`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) UnsetTotalUsageLimit()`

UnsetTotalUsageLimit ensures that no value is present for TotalUsageLimit, not even an explicit nil
### GetDisable

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetDisable() interface{}`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetDisableOk() (*interface{}, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetDisable(v interface{})`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### SetDisableNil

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetDisableNil(b bool)`

 SetDisableNil sets the value for Disable to be an explicit nil

### UnsetDisable
`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) UnsetDisable()`

UnsetDisable ensures that no value is present for Disable, not even an explicit nil
### GetEnable

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetEnable() interface{}`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetEnableOk() (*interface{}, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetEnable(v interface{})`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetReference

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFixedAmountCents

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetFixedAmountCents() interface{}`

GetFixedAmountCents returns the FixedAmountCents field if non-nil, zero value otherwise.

### GetFixedAmountCentsOk

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) GetFixedAmountCentsOk() (*interface{}, bool)`

GetFixedAmountCentsOk returns a tuple with the FixedAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmountCents

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetFixedAmountCents(v interface{})`

SetFixedAmountCents sets FixedAmountCents field to given value.

### HasFixedAmountCents

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) HasFixedAmountCents() bool`

HasFixedAmountCents returns a boolean if a field has been set.

### SetFixedAmountCentsNil

`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) SetFixedAmountCentsNil(b bool)`

 SetFixedAmountCentsNil sets the value for FixedAmountCents to be an explicit nil

### UnsetFixedAmountCents
`func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes) UnsetFixedAmountCents()`

UnsetFixedAmountCents ensures that no value is present for FixedAmountCents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


