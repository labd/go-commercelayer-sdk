# PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes

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
**X** | Pointer to **interface{}** | The quantity which defines the threshold for free items (works by multiple of x). | [optional] 
**Y** | Pointer to **interface{}** | The quantity which defines how many items you get for free, with the formula x-y. | [optional] 
**CheapestFree** | Pointer to **interface{}** | Indicates if the cheapest items are discounted, allowing all of the SKUs in the associated list to be eligible for counting. | [optional] 

## Methods

### NewPATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes

`func NewPATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes() *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes`

NewPATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes instantiates a new PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributesWithDefaults

`func NewPATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributesWithDefaults() *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes`

NewPATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributesWithDefaults instantiates a new PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrencyCode

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetExclusive

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetExclusive() interface{}`

GetExclusive returns the Exclusive field if non-nil, zero value otherwise.

### GetExclusiveOk

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetExclusiveOk() (*interface{}, bool)`

GetExclusiveOk returns a tuple with the Exclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusive

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetExclusive(v interface{})`

SetExclusive sets Exclusive field to given value.

### HasExclusive

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) HasExclusive() bool`

HasExclusive returns a boolean if a field has been set.

### SetExclusiveNil

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetExclusiveNil(b bool)`

 SetExclusiveNil sets the value for Exclusive to be an explicit nil

### UnsetExclusive
`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) UnsetExclusive()`

UnsetExclusive ensures that no value is present for Exclusive, not even an explicit nil
### GetPriority

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetPriority() interface{}`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetPriorityOk() (*interface{}, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetPriority(v interface{})`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetStartsAt

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetStartsAt() interface{}`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetStartsAt(v interface{})`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetExpiresAt

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetTotalUsageLimit

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetTotalUsageLimit() interface{}`

GetTotalUsageLimit returns the TotalUsageLimit field if non-nil, zero value otherwise.

### GetTotalUsageLimitOk

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetTotalUsageLimitOk() (*interface{}, bool)`

GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageLimit

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetTotalUsageLimit(v interface{})`

SetTotalUsageLimit sets TotalUsageLimit field to given value.

### HasTotalUsageLimit

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) HasTotalUsageLimit() bool`

HasTotalUsageLimit returns a boolean if a field has been set.

### SetTotalUsageLimitNil

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetTotalUsageLimitNil(b bool)`

 SetTotalUsageLimitNil sets the value for TotalUsageLimit to be an explicit nil

### UnsetTotalUsageLimit
`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) UnsetTotalUsageLimit()`

UnsetTotalUsageLimit ensures that no value is present for TotalUsageLimit, not even an explicit nil
### GetDisable

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetDisable() interface{}`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetDisableOk() (*interface{}, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetDisable(v interface{})`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### SetDisableNil

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetDisableNil(b bool)`

 SetDisableNil sets the value for Disable to be an explicit nil

### UnsetDisable
`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) UnsetDisable()`

UnsetDisable ensures that no value is present for Disable, not even an explicit nil
### GetEnable

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetEnable() interface{}`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetEnableOk() (*interface{}, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetEnable(v interface{})`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetReference

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetX

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetX() interface{}`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetXOk() (*interface{}, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetX(v interface{})`

SetX sets X field to given value.

### HasX

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) HasX() bool`

HasX returns a boolean if a field has been set.

### SetXNil

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetXNil(b bool)`

 SetXNil sets the value for X to be an explicit nil

### UnsetX
`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) UnsetX()`

UnsetX ensures that no value is present for X, not even an explicit nil
### GetY

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetY() interface{}`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetYOk() (*interface{}, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetY(v interface{})`

SetY sets Y field to given value.

### HasY

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) HasY() bool`

HasY returns a boolean if a field has been set.

### SetYNil

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetYNil(b bool)`

 SetYNil sets the value for Y to be an explicit nil

### UnsetY
`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) UnsetY()`

UnsetY ensures that no value is present for Y, not even an explicit nil
### GetCheapestFree

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetCheapestFree() interface{}`

GetCheapestFree returns the CheapestFree field if non-nil, zero value otherwise.

### GetCheapestFreeOk

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) GetCheapestFreeOk() (*interface{}, bool)`

GetCheapestFreeOk returns a tuple with the CheapestFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheapestFree

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetCheapestFree(v interface{})`

SetCheapestFree sets CheapestFree field to given value.

### HasCheapestFree

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) HasCheapestFree() bool`

HasCheapestFree returns a boolean if a field has been set.

### SetCheapestFreeNil

`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) SetCheapestFreeNil(b bool)`

 SetCheapestFreeNil sets the value for CheapestFree to be an explicit nil

### UnsetCheapestFree
`func (o *PATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes) UnsetCheapestFree()`

UnsetCheapestFree ensures that no value is present for CheapestFree, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


