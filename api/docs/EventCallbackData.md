# EventCallbackData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "event_callbacks"]
**Attributes** | [**GETEventCallbacks200ResponseDataInnerAttributes**](GETEventCallbacks200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**GETEventCallbacks200ResponseDataInnerRelationships**](GETEventCallbacks200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewEventCallbackData

`func NewEventCallbackData(type_ string, attributes GETEventCallbacks200ResponseDataInnerAttributes, ) *EventCallbackData`

NewEventCallbackData instantiates a new EventCallbackData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventCallbackDataWithDefaults

`func NewEventCallbackDataWithDefaults() *EventCallbackData`

NewEventCallbackDataWithDefaults instantiates a new EventCallbackData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventCallbackData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventCallbackData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventCallbackData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EventCallbackData) GetAttributes() GETEventCallbacks200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EventCallbackData) GetAttributesOk() (*GETEventCallbacks200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EventCallbackData) SetAttributes(v GETEventCallbacks200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *EventCallbackData) GetRelationships() GETEventCallbacks200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *EventCallbackData) GetRelationshipsOk() (*GETEventCallbacks200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *EventCallbackData) SetRelationships(v GETEventCallbacks200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *EventCallbackData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


