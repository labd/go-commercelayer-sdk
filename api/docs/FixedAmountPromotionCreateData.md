# FixedAmountPromotionCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTFixedAmountPromotions201ResponseDataAttributes**](POSTFixedAmountPromotions201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ExternalPromotionCreateDataRelationships**](ExternalPromotionCreateDataRelationships.md) |  | [optional] 

## Methods

### NewFixedAmountPromotionCreateData

`func NewFixedAmountPromotionCreateData(type_ interface{}, attributes POSTFixedAmountPromotions201ResponseDataAttributes, ) *FixedAmountPromotionCreateData`

NewFixedAmountPromotionCreateData instantiates a new FixedAmountPromotionCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedAmountPromotionCreateDataWithDefaults

`func NewFixedAmountPromotionCreateDataWithDefaults() *FixedAmountPromotionCreateData`

NewFixedAmountPromotionCreateDataWithDefaults instantiates a new FixedAmountPromotionCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FixedAmountPromotionCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FixedAmountPromotionCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FixedAmountPromotionCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *FixedAmountPromotionCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FixedAmountPromotionCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *FixedAmountPromotionCreateData) GetAttributes() POSTFixedAmountPromotions201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FixedAmountPromotionCreateData) GetAttributesOk() (*POSTFixedAmountPromotions201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FixedAmountPromotionCreateData) SetAttributes(v POSTFixedAmountPromotions201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *FixedAmountPromotionCreateData) GetRelationships() ExternalPromotionCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FixedAmountPromotionCreateData) GetRelationshipsOk() (*ExternalPromotionCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FixedAmountPromotionCreateData) SetRelationships(v ExternalPromotionCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FixedAmountPromotionCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


