# POSTOrders201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **interface{}** | The resource&#39;s id | [optional] 
**Type** | Pointer to **interface{}** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**POSTAddresses201ResponseDataLinks**](POSTAddresses201ResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTOrdersRequestDataAttributes**](POSTOrdersRequestDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**POSTOrders201ResponseDataRelationships**](POSTOrders201ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTOrders201ResponseData

`func NewPOSTOrders201ResponseData() *POSTOrders201ResponseData`

NewPOSTOrders201ResponseData instantiates a new POSTOrders201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrders201ResponseDataWithDefaults

`func NewPOSTOrders201ResponseDataWithDefaults() *POSTOrders201ResponseData`

NewPOSTOrders201ResponseDataWithDefaults instantiates a new POSTOrders201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTOrders201ResponseData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTOrders201ResponseData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTOrders201ResponseData) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *POSTOrders201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *POSTOrders201ResponseData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *POSTOrders201ResponseData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *POSTOrders201ResponseData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTOrders201ResponseData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTOrders201ResponseData) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *POSTOrders201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *POSTOrders201ResponseData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTOrders201ResponseData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetLinks

`func (o *POSTOrders201ResponseData) GetLinks() POSTAddresses201ResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTOrders201ResponseData) GetLinksOk() (*POSTAddresses201ResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTOrders201ResponseData) SetLinks(v POSTAddresses201ResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTOrders201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTOrders201ResponseData) GetAttributes() POSTOrdersRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTOrders201ResponseData) GetAttributesOk() (*POSTOrdersRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTOrders201ResponseData) SetAttributes(v POSTOrdersRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTOrders201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTOrders201ResponseData) GetRelationships() POSTOrders201ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTOrders201ResponseData) GetRelationshipsOk() (*POSTOrders201ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTOrders201ResponseData) SetRelationships(v POSTOrders201ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTOrders201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


