# RecurringOrderCopyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETOrderFactories200ResponseDataInnerAttributes**](GETOrderFactories200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**OrderCopyDataRelationships**](OrderCopyDataRelationships.md) |  | [optional] 

## Methods

### NewRecurringOrderCopyData

`func NewRecurringOrderCopyData(type_ interface{}, attributes GETOrderFactories200ResponseDataInnerAttributes, ) *RecurringOrderCopyData`

NewRecurringOrderCopyData instantiates a new RecurringOrderCopyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringOrderCopyDataWithDefaults

`func NewRecurringOrderCopyDataWithDefaults() *RecurringOrderCopyData`

NewRecurringOrderCopyDataWithDefaults instantiates a new RecurringOrderCopyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RecurringOrderCopyData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecurringOrderCopyData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecurringOrderCopyData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *RecurringOrderCopyData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RecurringOrderCopyData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *RecurringOrderCopyData) GetAttributes() GETOrderFactories200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RecurringOrderCopyData) GetAttributesOk() (*GETOrderFactories200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RecurringOrderCopyData) SetAttributes(v GETOrderFactories200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *RecurringOrderCopyData) GetRelationships() OrderCopyDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *RecurringOrderCopyData) GetRelationshipsOk() (*OrderCopyDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *RecurringOrderCopyData) SetRelationships(v OrderCopyDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *RecurringOrderCopyData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


