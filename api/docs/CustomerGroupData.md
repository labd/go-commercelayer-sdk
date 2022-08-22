# CustomerGroupData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "customer_groups"]
**Attributes** | [**GETCustomerGroups200ResponseDataInnerAttributes**](GETCustomerGroups200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**GETCustomerGroups200ResponseDataInnerRelationships**](GETCustomerGroups200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewCustomerGroupData

`func NewCustomerGroupData(type_ string, attributes GETCustomerGroups200ResponseDataInnerAttributes, ) *CustomerGroupData`

NewCustomerGroupData instantiates a new CustomerGroupData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerGroupDataWithDefaults

`func NewCustomerGroupDataWithDefaults() *CustomerGroupData`

NewCustomerGroupDataWithDefaults instantiates a new CustomerGroupData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerGroupData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerGroupData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerGroupData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *CustomerGroupData) GetAttributes() GETCustomerGroups200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerGroupData) GetAttributesOk() (*GETCustomerGroups200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerGroupData) SetAttributes(v GETCustomerGroups200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerGroupData) GetRelationships() GETCustomerGroups200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerGroupData) GetRelationshipsOk() (*GETCustomerGroups200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerGroupData) SetRelationships(v GETCustomerGroups200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerGroupData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


