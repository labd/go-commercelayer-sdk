# PATCHStripePaymentsStripePaymentIdRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHStripePaymentsStripePaymentIdRequestDataAttributes**](PATCHStripePaymentsStripePaymentIdRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships**](PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPATCHStripePaymentsStripePaymentIdRequestData

`func NewPATCHStripePaymentsStripePaymentIdRequestData(type_ interface{}, id interface{}, attributes PATCHStripePaymentsStripePaymentIdRequestDataAttributes, ) *PATCHStripePaymentsStripePaymentIdRequestData`

NewPATCHStripePaymentsStripePaymentIdRequestData instantiates a new PATCHStripePaymentsStripePaymentIdRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHStripePaymentsStripePaymentIdRequestDataWithDefaults

`func NewPATCHStripePaymentsStripePaymentIdRequestDataWithDefaults() *PATCHStripePaymentsStripePaymentIdRequestData`

NewPATCHStripePaymentsStripePaymentIdRequestDataWithDefaults instantiates a new PATCHStripePaymentsStripePaymentIdRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PATCHStripePaymentsStripePaymentIdRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PATCHStripePaymentsStripePaymentIdRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PATCHStripePaymentsStripePaymentIdRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PATCHStripePaymentsStripePaymentIdRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PATCHStripePaymentsStripePaymentIdRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *PATCHStripePaymentsStripePaymentIdRequestData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PATCHStripePaymentsStripePaymentIdRequestData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PATCHStripePaymentsStripePaymentIdRequestData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *PATCHStripePaymentsStripePaymentIdRequestData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PATCHStripePaymentsStripePaymentIdRequestData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *PATCHStripePaymentsStripePaymentIdRequestData) GetAttributes() PATCHStripePaymentsStripePaymentIdRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PATCHStripePaymentsStripePaymentIdRequestData) GetAttributesOk() (*PATCHStripePaymentsStripePaymentIdRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PATCHStripePaymentsStripePaymentIdRequestData) SetAttributes(v PATCHStripePaymentsStripePaymentIdRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PATCHStripePaymentsStripePaymentIdRequestData) GetRelationships() PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PATCHStripePaymentsStripePaymentIdRequestData) GetRelationshipsOk() (*PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PATCHStripePaymentsStripePaymentIdRequestData) SetRelationships(v PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PATCHStripePaymentsStripePaymentIdRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


