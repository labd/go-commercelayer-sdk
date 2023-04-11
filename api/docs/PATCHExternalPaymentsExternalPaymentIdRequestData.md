# PATCHExternalPaymentsExternalPaymentIdRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHExternalPaymentsExternalPaymentIdRequestDataAttributes**](PATCHExternalPaymentsExternalPaymentIdRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships**](PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPATCHExternalPaymentsExternalPaymentIdRequestData

`func NewPATCHExternalPaymentsExternalPaymentIdRequestData(type_ interface{}, id interface{}, attributes PATCHExternalPaymentsExternalPaymentIdRequestDataAttributes, ) *PATCHExternalPaymentsExternalPaymentIdRequestData`

NewPATCHExternalPaymentsExternalPaymentIdRequestData instantiates a new PATCHExternalPaymentsExternalPaymentIdRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHExternalPaymentsExternalPaymentIdRequestDataWithDefaults

`func NewPATCHExternalPaymentsExternalPaymentIdRequestDataWithDefaults() *PATCHExternalPaymentsExternalPaymentIdRequestData`

NewPATCHExternalPaymentsExternalPaymentIdRequestDataWithDefaults instantiates a new PATCHExternalPaymentsExternalPaymentIdRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) GetAttributes() PATCHExternalPaymentsExternalPaymentIdRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) GetAttributesOk() (*PATCHExternalPaymentsExternalPaymentIdRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) SetAttributes(v PATCHExternalPaymentsExternalPaymentIdRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) GetRelationships() PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) GetRelationshipsOk() (*PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) SetRelationships(v PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PATCHExternalPaymentsExternalPaymentIdRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


