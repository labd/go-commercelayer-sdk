# POSTFreeShippingPromotions201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The promotion&#39;s internal name. | 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Exclusive** | Pointer to **interface{}** | Indicates if the promotion will be applied exclusively, based on its priority score. | [optional] 
**Priority** | Pointer to **interface{}** | The priority assigned to the promotion (lower means higher priority). | [optional] 
**StartsAt** | **interface{}** | The activation date/time of this promotion. | 
**ExpiresAt** | **interface{}** | The expiration date/time of this promotion (must be after starts_at). | 
**TotalUsageLimit** | Pointer to **interface{}** | The total number of times this promotion can be applied. When &#39;null&#39; it means promotion can be applied infinite times. | [optional] 
**Disable** | Pointer to **interface{}** | Send this attribute if you want to mark this resource as disabled. | [optional] 
**Enable** | Pointer to **interface{}** | Send this attribute if you want to mark this resource as enabled. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTFreeShippingPromotions201ResponseDataAttributes

`func NewPOSTFreeShippingPromotions201ResponseDataAttributes(name interface{}, startsAt interface{}, expiresAt interface{}, ) *POSTFreeShippingPromotions201ResponseDataAttributes`

NewPOSTFreeShippingPromotions201ResponseDataAttributes instantiates a new POSTFreeShippingPromotions201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTFreeShippingPromotions201ResponseDataAttributesWithDefaults

`func NewPOSTFreeShippingPromotions201ResponseDataAttributesWithDefaults() *POSTFreeShippingPromotions201ResponseDataAttributes`

NewPOSTFreeShippingPromotions201ResponseDataAttributesWithDefaults instantiates a new POSTFreeShippingPromotions201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrencyCode

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetExclusive

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetExclusive() interface{}`

GetExclusive returns the Exclusive field if non-nil, zero value otherwise.

### GetExclusiveOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetExclusiveOk() (*interface{}, bool)`

GetExclusiveOk returns a tuple with the Exclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusive

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetExclusive(v interface{})`

SetExclusive sets Exclusive field to given value.

### HasExclusive

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) HasExclusive() bool`

HasExclusive returns a boolean if a field has been set.

### SetExclusiveNil

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetExclusiveNil(b bool)`

 SetExclusiveNil sets the value for Exclusive to be an explicit nil

### UnsetExclusive
`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) UnsetExclusive()`

UnsetExclusive ensures that no value is present for Exclusive, not even an explicit nil
### GetPriority

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetPriority() interface{}`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetPriorityOk() (*interface{}, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetPriority(v interface{})`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetStartsAt

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetStartsAt() interface{}`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetStartsAt(v interface{})`

SetStartsAt sets StartsAt field to given value.


### SetStartsAtNil

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetExpiresAt

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetTotalUsageLimit

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetTotalUsageLimit() interface{}`

GetTotalUsageLimit returns the TotalUsageLimit field if non-nil, zero value otherwise.

### GetTotalUsageLimitOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetTotalUsageLimitOk() (*interface{}, bool)`

GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageLimit

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetTotalUsageLimit(v interface{})`

SetTotalUsageLimit sets TotalUsageLimit field to given value.

### HasTotalUsageLimit

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) HasTotalUsageLimit() bool`

HasTotalUsageLimit returns a boolean if a field has been set.

### SetTotalUsageLimitNil

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetTotalUsageLimitNil(b bool)`

 SetTotalUsageLimitNil sets the value for TotalUsageLimit to be an explicit nil

### UnsetTotalUsageLimit
`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) UnsetTotalUsageLimit()`

UnsetTotalUsageLimit ensures that no value is present for TotalUsageLimit, not even an explicit nil
### GetDisable

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetDisable() interface{}`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetDisableOk() (*interface{}, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetDisable(v interface{})`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### SetDisableNil

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetDisableNil(b bool)`

 SetDisableNil sets the value for Disable to be an explicit nil

### UnsetDisable
`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) UnsetDisable()`

UnsetDisable ensures that no value is present for Disable, not even an explicit nil
### GetEnable

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetEnable() interface{}`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetEnableOk() (*interface{}, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetEnable(v interface{})`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetReference

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


