# POSTOrderAmountPromotionRules201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**OrderAmountCents** | Pointer to **int32** | Apply the promotion only when order is over this amount, in cents. | [optional] 

## Methods

### NewPOSTOrderAmountPromotionRules201ResponseDataAttributes

`func NewPOSTOrderAmountPromotionRules201ResponseDataAttributes() *POSTOrderAmountPromotionRules201ResponseDataAttributes`

NewPOSTOrderAmountPromotionRules201ResponseDataAttributes instantiates a new POSTOrderAmountPromotionRules201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrderAmountPromotionRules201ResponseDataAttributesWithDefaults

`func NewPOSTOrderAmountPromotionRules201ResponseDataAttributesWithDefaults() *POSTOrderAmountPromotionRules201ResponseDataAttributes`

NewPOSTOrderAmountPromotionRules201ResponseDataAttributesWithDefaults instantiates a new POSTOrderAmountPromotionRules201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOrderAmountCents

`func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetOrderAmountCents() int32`

GetOrderAmountCents returns the OrderAmountCents field if non-nil, zero value otherwise.

### GetOrderAmountCentsOk

`func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) GetOrderAmountCentsOk() (*int32, bool)`

GetOrderAmountCentsOk returns a tuple with the OrderAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmountCents

`func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) SetOrderAmountCents(v int32)`

SetOrderAmountCents sets OrderAmountCents field to given value.

### HasOrderAmountCents

`func (o *POSTOrderAmountPromotionRules201ResponseDataAttributes) HasOrderAmountCents() bool`

HasOrderAmountCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


