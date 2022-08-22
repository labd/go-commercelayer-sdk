# POSTExternalPayments201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "external_payments"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTExternalPayments201ResponseDataAttributes**](POSTExternalPayments201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**POSTAdyenPayments201ResponseDataRelationships**](POSTAdyenPayments201ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTExternalPayments201ResponseData

`func NewPOSTExternalPayments201ResponseData() *POSTExternalPayments201ResponseData`

NewPOSTExternalPayments201ResponseData instantiates a new POSTExternalPayments201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTExternalPayments201ResponseDataWithDefaults

`func NewPOSTExternalPayments201ResponseDataWithDefaults() *POSTExternalPayments201ResponseData`

NewPOSTExternalPayments201ResponseDataWithDefaults instantiates a new POSTExternalPayments201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTExternalPayments201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTExternalPayments201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTExternalPayments201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTExternalPayments201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTExternalPayments201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTExternalPayments201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTExternalPayments201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTExternalPayments201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTExternalPayments201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTExternalPayments201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTExternalPayments201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTExternalPayments201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTExternalPayments201ResponseData) GetAttributes() POSTExternalPayments201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTExternalPayments201ResponseData) GetAttributesOk() (*POSTExternalPayments201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTExternalPayments201ResponseData) SetAttributes(v POSTExternalPayments201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTExternalPayments201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTExternalPayments201ResponseData) GetRelationships() POSTAdyenPayments201ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTExternalPayments201ResponseData) GetRelationshipsOk() (*POSTAdyenPayments201ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTExternalPayments201ResponseData) SetRelationships(v POSTAdyenPayments201ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTExternalPayments201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


