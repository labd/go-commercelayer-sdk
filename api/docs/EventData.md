# EventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETEvents200ResponseDataInnerAttributes**](GETEvents200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**EventDataRelationships**](EventDataRelationships.md) |  | [optional] 

## Methods

### NewEventData

`func NewEventData(type_ interface{}, attributes GETEvents200ResponseDataInnerAttributes, ) *EventData`

NewEventData instantiates a new EventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDataWithDefaults

`func NewEventDataWithDefaults() *EventData`

NewEventDataWithDefaults instantiates a new EventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *EventData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *EventData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *EventData) GetAttributes() GETEvents200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EventData) GetAttributesOk() (*GETEvents200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EventData) SetAttributes(v GETEvents200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *EventData) GetRelationships() EventDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *EventData) GetRelationshipsOk() (*EventDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *EventData) SetRelationships(v EventDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *EventData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


