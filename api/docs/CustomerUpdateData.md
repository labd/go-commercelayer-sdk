# CustomerUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**CustomerUpdateDataAttributes**](CustomerUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**CustomerCreateDataRelationships**](CustomerCreateDataRelationships.md) |  | [optional] 

## Methods

### NewCustomerUpdateData

`func NewCustomerUpdateData(type_ string, id string, attributes CustomerUpdateDataAttributes, ) *CustomerUpdateData`

NewCustomerUpdateData instantiates a new CustomerUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerUpdateDataWithDefaults

`func NewCustomerUpdateDataWithDefaults() *CustomerUpdateData`

NewCustomerUpdateDataWithDefaults instantiates a new CustomerUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CustomerUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CustomerUpdateData) GetAttributes() CustomerUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerUpdateData) GetAttributesOk() (*CustomerUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerUpdateData) SetAttributes(v CustomerUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerUpdateData) GetRelationships() CustomerCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerUpdateData) GetRelationshipsOk() (*CustomerCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerUpdateData) SetRelationships(v CustomerCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


