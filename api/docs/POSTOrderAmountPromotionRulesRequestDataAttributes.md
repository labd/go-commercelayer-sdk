# POSTOrderAmountPromotionRulesRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**OrderAmountCents** | Pointer to **interface{}** | Apply the promotion only when order is over this amount, in cents. | [optional] 

## Methods

### NewPOSTOrderAmountPromotionRulesRequestDataAttributes

`func NewPOSTOrderAmountPromotionRulesRequestDataAttributes() *POSTOrderAmountPromotionRulesRequestDataAttributes`

NewPOSTOrderAmountPromotionRulesRequestDataAttributes instantiates a new POSTOrderAmountPromotionRulesRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrderAmountPromotionRulesRequestDataAttributesWithDefaults

`func NewPOSTOrderAmountPromotionRulesRequestDataAttributesWithDefaults() *POSTOrderAmountPromotionRulesRequestDataAttributes`

NewPOSTOrderAmountPromotionRulesRequestDataAttributesWithDefaults instantiates a new POSTOrderAmountPromotionRulesRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOrderAmountCents

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) GetOrderAmountCents() interface{}`

GetOrderAmountCents returns the OrderAmountCents field if non-nil, zero value otherwise.

### GetOrderAmountCentsOk

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) GetOrderAmountCentsOk() (*interface{}, bool)`

GetOrderAmountCentsOk returns a tuple with the OrderAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmountCents

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) SetOrderAmountCents(v interface{})`

SetOrderAmountCents sets OrderAmountCents field to given value.

### HasOrderAmountCents

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) HasOrderAmountCents() bool`

HasOrderAmountCents returns a boolean if a field has been set.

### SetOrderAmountCentsNil

`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) SetOrderAmountCentsNil(b bool)`

 SetOrderAmountCentsNil sets the value for OrderAmountCents to be an explicit nil

### UnsetOrderAmountCents
`func (o *POSTOrderAmountPromotionRulesRequestDataAttributes) UnsetOrderAmountCents()`

UnsetOrderAmountCents ensures that no value is present for OrderAmountCents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


