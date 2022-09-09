# LineItemOptionCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "line_item_options"]
**Attributes** | [**POSTLineItemOptions201ResponseDataAttributes**](POSTLineItemOptions201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**LineItemOptionCreateDataRelationships**](LineItemOptionCreateDataRelationships.md) |  | [optional] 

## Methods

### NewLineItemOptionCreateData

`func NewLineItemOptionCreateData(type_ string, attributes POSTLineItemOptions201ResponseDataAttributes, ) *LineItemOptionCreateData`

NewLineItemOptionCreateData instantiates a new LineItemOptionCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemOptionCreateDataWithDefaults

`func NewLineItemOptionCreateDataWithDefaults() *LineItemOptionCreateData`

NewLineItemOptionCreateDataWithDefaults instantiates a new LineItemOptionCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LineItemOptionCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LineItemOptionCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LineItemOptionCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *LineItemOptionCreateData) GetAttributes() POSTLineItemOptions201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LineItemOptionCreateData) GetAttributesOk() (*POSTLineItemOptions201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LineItemOptionCreateData) SetAttributes(v POSTLineItemOptions201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *LineItemOptionCreateData) GetRelationships() LineItemOptionCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *LineItemOptionCreateData) GetRelationshipsOk() (*LineItemOptionCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *LineItemOptionCreateData) SetRelationships(v LineItemOptionCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *LineItemOptionCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


