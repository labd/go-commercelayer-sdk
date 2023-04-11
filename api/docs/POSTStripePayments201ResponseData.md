# POSTStripePayments201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **interface{}** | The resource&#39;s id | [optional] 
**Type** | Pointer to **interface{}** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**POSTAddresses201ResponseDataLinks**](POSTAddresses201ResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTStripePaymentsRequestDataAttributes**](POSTStripePaymentsRequestDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**POSTAdyenPayments201ResponseDataRelationships**](POSTAdyenPayments201ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTStripePayments201ResponseData

`func NewPOSTStripePayments201ResponseData() *POSTStripePayments201ResponseData`

NewPOSTStripePayments201ResponseData instantiates a new POSTStripePayments201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStripePayments201ResponseDataWithDefaults

`func NewPOSTStripePayments201ResponseDataWithDefaults() *POSTStripePayments201ResponseData`

NewPOSTStripePayments201ResponseDataWithDefaults instantiates a new POSTStripePayments201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTStripePayments201ResponseData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTStripePayments201ResponseData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTStripePayments201ResponseData) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *POSTStripePayments201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *POSTStripePayments201ResponseData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *POSTStripePayments201ResponseData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *POSTStripePayments201ResponseData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTStripePayments201ResponseData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTStripePayments201ResponseData) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *POSTStripePayments201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *POSTStripePayments201ResponseData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTStripePayments201ResponseData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetLinks

`func (o *POSTStripePayments201ResponseData) GetLinks() POSTAddresses201ResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTStripePayments201ResponseData) GetLinksOk() (*POSTAddresses201ResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTStripePayments201ResponseData) SetLinks(v POSTAddresses201ResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTStripePayments201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTStripePayments201ResponseData) GetAttributes() POSTStripePaymentsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTStripePayments201ResponseData) GetAttributesOk() (*POSTStripePaymentsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTStripePayments201ResponseData) SetAttributes(v POSTStripePaymentsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTStripePayments201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTStripePayments201ResponseData) GetRelationships() POSTAdyenPayments201ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTStripePayments201ResponseData) GetRelationshipsOk() (*POSTAdyenPayments201ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTStripePayments201ResponseData) SetRelationships(v POSTAdyenPayments201ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTStripePayments201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


