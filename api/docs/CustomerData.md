# CustomerData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETCustomers200ResponseDataInnerAttributes**](GETCustomers200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**CustomerDataRelationships**](CustomerDataRelationships.md) |  | [optional] 

## Methods

### NewCustomerData

`func NewCustomerData(type_ interface{}, attributes GETCustomers200ResponseDataInnerAttributes, ) *CustomerData`

NewCustomerData instantiates a new CustomerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerDataWithDefaults

`func NewCustomerDataWithDefaults() *CustomerData`

NewCustomerDataWithDefaults instantiates a new CustomerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CustomerData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CustomerData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *CustomerData) GetAttributes() GETCustomers200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerData) GetAttributesOk() (*GETCustomers200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerData) SetAttributes(v GETCustomers200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerData) GetRelationships() CustomerDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerData) GetRelationshipsOk() (*CustomerDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerData) SetRelationships(v CustomerDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


