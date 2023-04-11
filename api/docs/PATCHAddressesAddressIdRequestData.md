# PATCHAddressesAddressIdRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHAddressesAddressIdRequestDataAttributes**](PATCHAddressesAddressIdRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTAddressesRequestDataRelationships**](POSTAddressesRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPATCHAddressesAddressIdRequestData

`func NewPATCHAddressesAddressIdRequestData(type_ interface{}, id interface{}, attributes PATCHAddressesAddressIdRequestDataAttributes, ) *PATCHAddressesAddressIdRequestData`

NewPATCHAddressesAddressIdRequestData instantiates a new PATCHAddressesAddressIdRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHAddressesAddressIdRequestDataWithDefaults

`func NewPATCHAddressesAddressIdRequestDataWithDefaults() *PATCHAddressesAddressIdRequestData`

NewPATCHAddressesAddressIdRequestDataWithDefaults instantiates a new PATCHAddressesAddressIdRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PATCHAddressesAddressIdRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PATCHAddressesAddressIdRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PATCHAddressesAddressIdRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PATCHAddressesAddressIdRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PATCHAddressesAddressIdRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *PATCHAddressesAddressIdRequestData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PATCHAddressesAddressIdRequestData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PATCHAddressesAddressIdRequestData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *PATCHAddressesAddressIdRequestData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PATCHAddressesAddressIdRequestData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *PATCHAddressesAddressIdRequestData) GetAttributes() PATCHAddressesAddressIdRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PATCHAddressesAddressIdRequestData) GetAttributesOk() (*PATCHAddressesAddressIdRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PATCHAddressesAddressIdRequestData) SetAttributes(v PATCHAddressesAddressIdRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PATCHAddressesAddressIdRequestData) GetRelationships() POSTAddressesRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PATCHAddressesAddressIdRequestData) GetRelationshipsOk() (*POSTAddressesRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PATCHAddressesAddressIdRequestData) SetRelationships(v POSTAddressesRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PATCHAddressesAddressIdRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


