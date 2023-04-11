# POSTFixedAmountPromotionsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTFixedAmountPromotionsRequestDataAttributes**](POSTFixedAmountPromotionsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTExternalPromotionsRequestDataRelationships**](POSTExternalPromotionsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTFixedAmountPromotionsRequestData

`func NewPOSTFixedAmountPromotionsRequestData(type_ interface{}, attributes POSTFixedAmountPromotionsRequestDataAttributes, ) *POSTFixedAmountPromotionsRequestData`

NewPOSTFixedAmountPromotionsRequestData instantiates a new POSTFixedAmountPromotionsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTFixedAmountPromotionsRequestDataWithDefaults

`func NewPOSTFixedAmountPromotionsRequestDataWithDefaults() *POSTFixedAmountPromotionsRequestData`

NewPOSTFixedAmountPromotionsRequestDataWithDefaults instantiates a new POSTFixedAmountPromotionsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTFixedAmountPromotionsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTFixedAmountPromotionsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTFixedAmountPromotionsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTFixedAmountPromotionsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTFixedAmountPromotionsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTFixedAmountPromotionsRequestData) GetAttributes() POSTFixedAmountPromotionsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTFixedAmountPromotionsRequestData) GetAttributesOk() (*POSTFixedAmountPromotionsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTFixedAmountPromotionsRequestData) SetAttributes(v POSTFixedAmountPromotionsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTFixedAmountPromotionsRequestData) GetRelationships() POSTExternalPromotionsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTFixedAmountPromotionsRequestData) GetRelationshipsOk() (*POSTExternalPromotionsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTFixedAmountPromotionsRequestData) SetRelationships(v POSTExternalPromotionsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTFixedAmountPromotionsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


