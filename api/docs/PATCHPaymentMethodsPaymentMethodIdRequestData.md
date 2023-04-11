# PATCHPaymentMethodsPaymentMethodIdRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHPaymentMethodsPaymentMethodIdRequestDataAttributes**](PATCHPaymentMethodsPaymentMethodIdRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships**](PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPATCHPaymentMethodsPaymentMethodIdRequestData

`func NewPATCHPaymentMethodsPaymentMethodIdRequestData(type_ interface{}, id interface{}, attributes PATCHPaymentMethodsPaymentMethodIdRequestDataAttributes, ) *PATCHPaymentMethodsPaymentMethodIdRequestData`

NewPATCHPaymentMethodsPaymentMethodIdRequestData instantiates a new PATCHPaymentMethodsPaymentMethodIdRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHPaymentMethodsPaymentMethodIdRequestDataWithDefaults

`func NewPATCHPaymentMethodsPaymentMethodIdRequestDataWithDefaults() *PATCHPaymentMethodsPaymentMethodIdRequestData`

NewPATCHPaymentMethodsPaymentMethodIdRequestDataWithDefaults instantiates a new PATCHPaymentMethodsPaymentMethodIdRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) GetAttributes() PATCHPaymentMethodsPaymentMethodIdRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) GetAttributesOk() (*PATCHPaymentMethodsPaymentMethodIdRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) SetAttributes(v PATCHPaymentMethodsPaymentMethodIdRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) GetRelationships() PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) GetRelationshipsOk() (*PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) SetRelationships(v PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PATCHPaymentMethodsPaymentMethodIdRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


