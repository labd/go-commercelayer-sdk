# CustomerUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "customers"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHCustomersCustomerId200ResponseDataAttributes**](PATCHCustomersCustomerId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTCustomers201ResponseDataRelationships**](POSTCustomers201ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewCustomerUpdateData

`func NewCustomerUpdateData(type_ string, id string, attributes PATCHCustomersCustomerId200ResponseDataAttributes, ) *CustomerUpdateData`

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

`func (o *CustomerUpdateData) GetAttributes() PATCHCustomersCustomerId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerUpdateData) GetAttributesOk() (*PATCHCustomersCustomerId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerUpdateData) SetAttributes(v PATCHCustomersCustomerId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerUpdateData) GetRelationships() POSTCustomers201ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerUpdateData) GetRelationshipsOk() (*POSTCustomers201ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerUpdateData) SetRelationships(v POSTCustomers201ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


