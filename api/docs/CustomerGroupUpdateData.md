# CustomerGroupUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "customer_groups"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes**](PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCustomerGroupUpdateData

`func NewCustomerGroupUpdateData(type_ string, id string, attributes PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes, ) *CustomerGroupUpdateData`

NewCustomerGroupUpdateData instantiates a new CustomerGroupUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerGroupUpdateDataWithDefaults

`func NewCustomerGroupUpdateDataWithDefaults() *CustomerGroupUpdateData`

NewCustomerGroupUpdateDataWithDefaults instantiates a new CustomerGroupUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerGroupUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerGroupUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerGroupUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CustomerGroupUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerGroupUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerGroupUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CustomerGroupUpdateData) GetAttributes() PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerGroupUpdateData) GetAttributesOk() (*PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerGroupUpdateData) SetAttributes(v PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerGroupUpdateData) GetRelationships() map[string]interface{}`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerGroupUpdateData) GetRelationshipsOk() (*map[string]interface{}, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerGroupUpdateData) SetRelationships(v map[string]interface{})`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerGroupUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


