# POSTWireTransfers201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTAdyenPayments201ResponseDataAttributes**](POSTAdyenPayments201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETTransactions200ResponseDataInnerRelationships**](GETTransactions200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewPOSTWireTransfers201ResponseData

`func NewPOSTWireTransfers201ResponseData() *POSTWireTransfers201ResponseData`

NewPOSTWireTransfers201ResponseData instantiates a new POSTWireTransfers201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTWireTransfers201ResponseDataWithDefaults

`func NewPOSTWireTransfers201ResponseDataWithDefaults() *POSTWireTransfers201ResponseData`

NewPOSTWireTransfers201ResponseDataWithDefaults instantiates a new POSTWireTransfers201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTWireTransfers201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTWireTransfers201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTWireTransfers201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTWireTransfers201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTWireTransfers201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTWireTransfers201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTWireTransfers201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTWireTransfers201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTWireTransfers201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTWireTransfers201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTWireTransfers201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTWireTransfers201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTWireTransfers201ResponseData) GetAttributes() POSTAdyenPayments201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTWireTransfers201ResponseData) GetAttributesOk() (*POSTAdyenPayments201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTWireTransfers201ResponseData) SetAttributes(v POSTAdyenPayments201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTWireTransfers201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTWireTransfers201ResponseData) GetRelationships() GETTransactions200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTWireTransfers201ResponseData) GetRelationshipsOk() (*GETTransactions200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTWireTransfers201ResponseData) SetRelationships(v GETTransactions200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTWireTransfers201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


