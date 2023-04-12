# RecurringOrderCopyCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTRecurringOrderCopies201ResponseDataAttributes**](POSTRecurringOrderCopies201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**RecurringOrderCopyCreateDataRelationships**](RecurringOrderCopyCreateDataRelationships.md) |  | [optional] 

## Methods

### NewRecurringOrderCopyCreateData

`func NewRecurringOrderCopyCreateData(type_ interface{}, attributes POSTRecurringOrderCopies201ResponseDataAttributes, ) *RecurringOrderCopyCreateData`

NewRecurringOrderCopyCreateData instantiates a new RecurringOrderCopyCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringOrderCopyCreateDataWithDefaults

`func NewRecurringOrderCopyCreateDataWithDefaults() *RecurringOrderCopyCreateData`

NewRecurringOrderCopyCreateDataWithDefaults instantiates a new RecurringOrderCopyCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RecurringOrderCopyCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecurringOrderCopyCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecurringOrderCopyCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *RecurringOrderCopyCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RecurringOrderCopyCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *RecurringOrderCopyCreateData) GetAttributes() POSTRecurringOrderCopies201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RecurringOrderCopyCreateData) GetAttributesOk() (*POSTRecurringOrderCopies201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RecurringOrderCopyCreateData) SetAttributes(v POSTRecurringOrderCopies201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *RecurringOrderCopyCreateData) GetRelationships() RecurringOrderCopyCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *RecurringOrderCopyCreateData) GetRelationshipsOk() (*RecurringOrderCopyCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *RecurringOrderCopyCreateData) SetRelationships(v RecurringOrderCopyCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *RecurringOrderCopyCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


