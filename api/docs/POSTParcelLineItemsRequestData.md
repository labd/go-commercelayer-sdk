# POSTParcelLineItemsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTParcelLineItemsRequestDataAttributes**](POSTParcelLineItemsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTParcelLineItemsRequestDataRelationships**](POSTParcelLineItemsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTParcelLineItemsRequestData

`func NewPOSTParcelLineItemsRequestData(type_ interface{}, attributes POSTParcelLineItemsRequestDataAttributes, ) *POSTParcelLineItemsRequestData`

NewPOSTParcelLineItemsRequestData instantiates a new POSTParcelLineItemsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTParcelLineItemsRequestDataWithDefaults

`func NewPOSTParcelLineItemsRequestDataWithDefaults() *POSTParcelLineItemsRequestData`

NewPOSTParcelLineItemsRequestDataWithDefaults instantiates a new POSTParcelLineItemsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTParcelLineItemsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTParcelLineItemsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTParcelLineItemsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTParcelLineItemsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTParcelLineItemsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTParcelLineItemsRequestData) GetAttributes() POSTParcelLineItemsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTParcelLineItemsRequestData) GetAttributesOk() (*POSTParcelLineItemsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTParcelLineItemsRequestData) SetAttributes(v POSTParcelLineItemsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTParcelLineItemsRequestData) GetRelationships() POSTParcelLineItemsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTParcelLineItemsRequestData) GetRelationshipsOk() (*POSTParcelLineItemsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTParcelLineItemsRequestData) SetRelationships(v POSTParcelLineItemsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTParcelLineItemsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


