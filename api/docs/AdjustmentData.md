# AdjustmentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETAdjustmentsAdjustmentId200ResponseDataAttributes**](GETAdjustmentsAdjustmentId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AdjustmentDataRelationships**](AdjustmentDataRelationships.md) |  | [optional] 

## Methods

### NewAdjustmentData

`func NewAdjustmentData(type_ interface{}, attributes GETAdjustmentsAdjustmentId200ResponseDataAttributes, ) *AdjustmentData`

NewAdjustmentData instantiates a new AdjustmentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustmentDataWithDefaults

`func NewAdjustmentDataWithDefaults() *AdjustmentData`

NewAdjustmentDataWithDefaults instantiates a new AdjustmentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AdjustmentData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdjustmentData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdjustmentData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *AdjustmentData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AdjustmentData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *AdjustmentData) GetAttributes() GETAdjustmentsAdjustmentId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AdjustmentData) GetAttributesOk() (*GETAdjustmentsAdjustmentId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AdjustmentData) SetAttributes(v GETAdjustmentsAdjustmentId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AdjustmentData) GetRelationships() AdjustmentDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AdjustmentData) GetRelationshipsOk() (*AdjustmentDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AdjustmentData) SetRelationships(v AdjustmentDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AdjustmentData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


