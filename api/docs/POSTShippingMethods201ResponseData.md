# POSTShippingMethods201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "shipping_methods"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTShippingMethods201ResponseDataAttributes**](POSTShippingMethods201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**POSTShippingMethods201ResponseDataRelationships**](POSTShippingMethods201ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTShippingMethods201ResponseData

`func NewPOSTShippingMethods201ResponseData() *POSTShippingMethods201ResponseData`

NewPOSTShippingMethods201ResponseData instantiates a new POSTShippingMethods201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTShippingMethods201ResponseDataWithDefaults

`func NewPOSTShippingMethods201ResponseDataWithDefaults() *POSTShippingMethods201ResponseData`

NewPOSTShippingMethods201ResponseDataWithDefaults instantiates a new POSTShippingMethods201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTShippingMethods201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTShippingMethods201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTShippingMethods201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTShippingMethods201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTShippingMethods201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTShippingMethods201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTShippingMethods201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTShippingMethods201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTShippingMethods201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTShippingMethods201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTShippingMethods201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTShippingMethods201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTShippingMethods201ResponseData) GetAttributes() POSTShippingMethods201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTShippingMethods201ResponseData) GetAttributesOk() (*POSTShippingMethods201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTShippingMethods201ResponseData) SetAttributes(v POSTShippingMethods201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTShippingMethods201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTShippingMethods201ResponseData) GetRelationships() POSTShippingMethods201ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTShippingMethods201ResponseData) GetRelationshipsOk() (*POSTShippingMethods201ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTShippingMethods201ResponseData) SetRelationships(v POSTShippingMethods201ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTShippingMethods201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


