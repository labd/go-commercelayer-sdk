# PATCHPaypalPaymentsPaypalPaymentIdRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes**](PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships**](PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPATCHPaypalPaymentsPaypalPaymentIdRequestData

`func NewPATCHPaypalPaymentsPaypalPaymentIdRequestData(type_ interface{}, id interface{}, attributes PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes, ) *PATCHPaypalPaymentsPaypalPaymentIdRequestData`

NewPATCHPaypalPaymentsPaypalPaymentIdRequestData instantiates a new PATCHPaypalPaymentsPaypalPaymentIdRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHPaypalPaymentsPaypalPaymentIdRequestDataWithDefaults

`func NewPATCHPaypalPaymentsPaypalPaymentIdRequestDataWithDefaults() *PATCHPaypalPaymentsPaypalPaymentIdRequestData`

NewPATCHPaypalPaymentsPaypalPaymentIdRequestDataWithDefaults instantiates a new PATCHPaypalPaymentsPaypalPaymentIdRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) GetAttributes() PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) GetAttributesOk() (*PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) SetAttributes(v PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) GetRelationships() PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) GetRelationshipsOk() (*PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) SetRelationships(v PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


