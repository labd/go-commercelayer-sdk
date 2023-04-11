# PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**AllSkus** | Pointer to **interface{}** | Indicates if the rule is activated only when all of the SKUs of the list is also part of the order. | [optional] 
**MinQuantity** | Pointer to **interface{}** | The min quantity of SKUs of the list that must be also part of the order. If positive, overwrites the &#39;all_skus&#39; option. When the SKU list is manual, its items quantities are honoured. | [optional] 

## Methods

### NewPATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes

`func NewPATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes() *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes`

NewPATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes instantiates a new PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributesWithDefaults

`func NewPATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributesWithDefaults() *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes`

NewPATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributesWithDefaults instantiates a new PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAllSkus

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) GetAllSkus() interface{}`

GetAllSkus returns the AllSkus field if non-nil, zero value otherwise.

### GetAllSkusOk

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) GetAllSkusOk() (*interface{}, bool)`

GetAllSkusOk returns a tuple with the AllSkus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSkus

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) SetAllSkus(v interface{})`

SetAllSkus sets AllSkus field to given value.

### HasAllSkus

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) HasAllSkus() bool`

HasAllSkus returns a boolean if a field has been set.

### SetAllSkusNil

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) SetAllSkusNil(b bool)`

 SetAllSkusNil sets the value for AllSkus to be an explicit nil

### UnsetAllSkus
`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) UnsetAllSkus()`

UnsetAllSkus ensures that no value is present for AllSkus, not even an explicit nil
### GetMinQuantity

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) GetMinQuantity() interface{}`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) GetMinQuantityOk() (*interface{}, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) SetMinQuantity(v interface{})`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### SetMinQuantityNil

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) SetMinQuantityNil(b bool)`

 SetMinQuantityNil sets the value for MinQuantity to be an explicit nil

### UnsetMinQuantity
`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) UnsetMinQuantity()`

UnsetMinQuantity ensures that no value is present for MinQuantity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


