# LineItemOptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "line_item_options"]
**Attributes** | [**GETLineItemOptions200ResponseDataInnerAttributes**](GETLineItemOptions200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**GETLineItemOptions200ResponseDataInnerRelationships**](GETLineItemOptions200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewLineItemOptionData

`func NewLineItemOptionData(type_ string, attributes GETLineItemOptions200ResponseDataInnerAttributes, ) *LineItemOptionData`

NewLineItemOptionData instantiates a new LineItemOptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemOptionDataWithDefaults

`func NewLineItemOptionDataWithDefaults() *LineItemOptionData`

NewLineItemOptionDataWithDefaults instantiates a new LineItemOptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LineItemOptionData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LineItemOptionData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LineItemOptionData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *LineItemOptionData) GetAttributes() GETLineItemOptions200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LineItemOptionData) GetAttributesOk() (*GETLineItemOptions200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LineItemOptionData) SetAttributes(v GETLineItemOptions200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *LineItemOptionData) GetRelationships() GETLineItemOptions200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *LineItemOptionData) GetRelationshipsOk() (*GETLineItemOptions200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *LineItemOptionData) SetRelationships(v GETLineItemOptions200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *LineItemOptionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


