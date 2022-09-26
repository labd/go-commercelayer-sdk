# CustomerCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**POSTCustomers201ResponseDataAttributes**](POSTCustomers201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CustomerCreateDataRelationships**](CustomerCreateDataRelationships.md) |  | [optional] 

## Methods

### NewCustomerCreateData

`func NewCustomerCreateData(type_ string, attributes POSTCustomers201ResponseDataAttributes, ) *CustomerCreateData`

NewCustomerCreateData instantiates a new CustomerCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerCreateDataWithDefaults

`func NewCustomerCreateDataWithDefaults() *CustomerCreateData`

NewCustomerCreateDataWithDefaults instantiates a new CustomerCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *CustomerCreateData) GetAttributes() POSTCustomers201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerCreateData) GetAttributesOk() (*POSTCustomers201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerCreateData) SetAttributes(v POSTCustomers201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerCreateData) GetRelationships() CustomerCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerCreateData) GetRelationshipsOk() (*CustomerCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerCreateData) SetRelationships(v CustomerCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


