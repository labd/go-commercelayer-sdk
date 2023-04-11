# PATCHCustomersCustomerIdRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHCustomersCustomerIdRequestDataAttributes**](PATCHCustomersCustomerIdRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTCustomersRequestDataRelationships**](POSTCustomersRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPATCHCustomersCustomerIdRequestData

`func NewPATCHCustomersCustomerIdRequestData(type_ interface{}, id interface{}, attributes PATCHCustomersCustomerIdRequestDataAttributes, ) *PATCHCustomersCustomerIdRequestData`

NewPATCHCustomersCustomerIdRequestData instantiates a new PATCHCustomersCustomerIdRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHCustomersCustomerIdRequestDataWithDefaults

`func NewPATCHCustomersCustomerIdRequestDataWithDefaults() *PATCHCustomersCustomerIdRequestData`

NewPATCHCustomersCustomerIdRequestDataWithDefaults instantiates a new PATCHCustomersCustomerIdRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PATCHCustomersCustomerIdRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PATCHCustomersCustomerIdRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PATCHCustomersCustomerIdRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PATCHCustomersCustomerIdRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PATCHCustomersCustomerIdRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *PATCHCustomersCustomerIdRequestData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PATCHCustomersCustomerIdRequestData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PATCHCustomersCustomerIdRequestData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *PATCHCustomersCustomerIdRequestData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PATCHCustomersCustomerIdRequestData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *PATCHCustomersCustomerIdRequestData) GetAttributes() PATCHCustomersCustomerIdRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PATCHCustomersCustomerIdRequestData) GetAttributesOk() (*PATCHCustomersCustomerIdRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PATCHCustomersCustomerIdRequestData) SetAttributes(v PATCHCustomersCustomerIdRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PATCHCustomersCustomerIdRequestData) GetRelationships() POSTCustomersRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PATCHCustomersCustomerIdRequestData) GetRelationshipsOk() (*POSTCustomersRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PATCHCustomersCustomerIdRequestData) SetRelationships(v POSTCustomersRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PATCHCustomersCustomerIdRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


