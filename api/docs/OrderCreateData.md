# OrderCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "orders"]
**Attributes** | [**POSTOrders201ResponseDataAttributes**](POSTOrders201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**OrderCreateDataRelationships**](OrderCreateDataRelationships.md) |  | [optional] 

## Methods

### NewOrderCreateData

`func NewOrderCreateData(type_ string, attributes POSTOrders201ResponseDataAttributes, ) *OrderCreateData`

NewOrderCreateData instantiates a new OrderCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateDataWithDefaults

`func NewOrderCreateDataWithDefaults() *OrderCreateData`

NewOrderCreateDataWithDefaults instantiates a new OrderCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *OrderCreateData) GetAttributes() POSTOrders201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderCreateData) GetAttributesOk() (*POSTOrders201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderCreateData) SetAttributes(v POSTOrders201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrderCreateData) GetRelationships() OrderCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrderCreateData) GetRelationshipsOk() (*OrderCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrderCreateData) SetRelationships(v OrderCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrderCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


