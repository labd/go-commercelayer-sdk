# POSTInventoryReturnLocationsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTInventoryReturnLocationsRequestDataAttributes**](POSTInventoryReturnLocationsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTInventoryReturnLocationsRequestDataRelationships**](POSTInventoryReturnLocationsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTInventoryReturnLocationsRequestData

`func NewPOSTInventoryReturnLocationsRequestData(type_ interface{}, attributes POSTInventoryReturnLocationsRequestDataAttributes, ) *POSTInventoryReturnLocationsRequestData`

NewPOSTInventoryReturnLocationsRequestData instantiates a new POSTInventoryReturnLocationsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTInventoryReturnLocationsRequestDataWithDefaults

`func NewPOSTInventoryReturnLocationsRequestDataWithDefaults() *POSTInventoryReturnLocationsRequestData`

NewPOSTInventoryReturnLocationsRequestDataWithDefaults instantiates a new POSTInventoryReturnLocationsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTInventoryReturnLocationsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTInventoryReturnLocationsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTInventoryReturnLocationsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTInventoryReturnLocationsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTInventoryReturnLocationsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTInventoryReturnLocationsRequestData) GetAttributes() POSTInventoryReturnLocationsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTInventoryReturnLocationsRequestData) GetAttributesOk() (*POSTInventoryReturnLocationsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTInventoryReturnLocationsRequestData) SetAttributes(v POSTInventoryReturnLocationsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTInventoryReturnLocationsRequestData) GetRelationships() POSTInventoryReturnLocationsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTInventoryReturnLocationsRequestData) GetRelationshipsOk() (*POSTInventoryReturnLocationsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTInventoryReturnLocationsRequestData) SetRelationships(v POSTInventoryReturnLocationsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTInventoryReturnLocationsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


