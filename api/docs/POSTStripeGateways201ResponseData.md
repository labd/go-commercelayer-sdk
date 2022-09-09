# POSTStripeGateways201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "stripe_gateways"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTStripeGateways201ResponseDataAttributes**](POSTStripeGateways201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETStripeGateways200ResponseDataInnerRelationships**](GETStripeGateways200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewPOSTStripeGateways201ResponseData

`func NewPOSTStripeGateways201ResponseData() *POSTStripeGateways201ResponseData`

NewPOSTStripeGateways201ResponseData instantiates a new POSTStripeGateways201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStripeGateways201ResponseDataWithDefaults

`func NewPOSTStripeGateways201ResponseDataWithDefaults() *POSTStripeGateways201ResponseData`

NewPOSTStripeGateways201ResponseDataWithDefaults instantiates a new POSTStripeGateways201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTStripeGateways201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTStripeGateways201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTStripeGateways201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTStripeGateways201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTStripeGateways201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTStripeGateways201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTStripeGateways201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTStripeGateways201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTStripeGateways201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTStripeGateways201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTStripeGateways201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTStripeGateways201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTStripeGateways201ResponseData) GetAttributes() POSTStripeGateways201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTStripeGateways201ResponseData) GetAttributesOk() (*POSTStripeGateways201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTStripeGateways201ResponseData) SetAttributes(v POSTStripeGateways201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTStripeGateways201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTStripeGateways201ResponseData) GetRelationships() GETStripeGateways200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTStripeGateways201ResponseData) GetRelationshipsOk() (*GETStripeGateways200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTStripeGateways201ResponseData) SetRelationships(v GETStripeGateways200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTStripeGateways201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


