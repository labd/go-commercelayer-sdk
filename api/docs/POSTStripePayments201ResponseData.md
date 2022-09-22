# POSTStripePayments201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTStripePayments201ResponseDataAttributes**](POSTStripePayments201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETAdyenPayments200ResponseDataInnerRelationships**](GETAdyenPayments200ResponseDataInnerRelationships.md) |  | [optional] 

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

`func (o *POSTStripePayments201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTStripePayments201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTStripePayments201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTStripePayments201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTStripePayments201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTStripePayments201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTStripePayments201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTStripePayments201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTStripePayments201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTStripePayments201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTStripePayments201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTStripePayments201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTStripePayments201ResponseData) GetAttributes() POSTStripePayments201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTStripePayments201ResponseData) GetAttributesOk() (*POSTStripePayments201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTStripePayments201ResponseData) SetAttributes(v POSTStripePayments201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTStripePayments201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTStripePayments201ResponseData) GetRelationships() GETAdyenPayments200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTStripePayments201ResponseData) GetRelationshipsOk() (*GETAdyenPayments200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTStripePayments201ResponseData) SetRelationships(v GETAdyenPayments200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTStripePayments201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


