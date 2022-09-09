# POSTInStockSubscriptions201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "in_stock_subscriptions"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTInStockSubscriptions201ResponseDataAttributes**](POSTInStockSubscriptions201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETInStockSubscriptions200ResponseDataInnerRelationships**](GETInStockSubscriptions200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewPOSTInStockSubscriptions201ResponseData

`func NewPOSTInStockSubscriptions201ResponseData() *POSTInStockSubscriptions201ResponseData`

NewPOSTInStockSubscriptions201ResponseData instantiates a new POSTInStockSubscriptions201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTInStockSubscriptions201ResponseDataWithDefaults

`func NewPOSTInStockSubscriptions201ResponseDataWithDefaults() *POSTInStockSubscriptions201ResponseData`

NewPOSTInStockSubscriptions201ResponseDataWithDefaults instantiates a new POSTInStockSubscriptions201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTInStockSubscriptions201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTInStockSubscriptions201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTInStockSubscriptions201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTInStockSubscriptions201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTInStockSubscriptions201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTInStockSubscriptions201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTInStockSubscriptions201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTInStockSubscriptions201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTInStockSubscriptions201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTInStockSubscriptions201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTInStockSubscriptions201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTInStockSubscriptions201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTInStockSubscriptions201ResponseData) GetAttributes() POSTInStockSubscriptions201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTInStockSubscriptions201ResponseData) GetAttributesOk() (*POSTInStockSubscriptions201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTInStockSubscriptions201ResponseData) SetAttributes(v POSTInStockSubscriptions201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTInStockSubscriptions201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTInStockSubscriptions201ResponseData) GetRelationships() GETInStockSubscriptions200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTInStockSubscriptions201ResponseData) GetRelationshipsOk() (*GETInStockSubscriptions200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTInStockSubscriptions201ResponseData) SetRelationships(v GETInStockSubscriptions200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTInStockSubscriptions201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


