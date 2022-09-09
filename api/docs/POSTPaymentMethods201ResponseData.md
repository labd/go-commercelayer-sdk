# POSTPaymentMethods201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "payment_methods"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTPaymentMethods201ResponseDataAttributes**](POSTPaymentMethods201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETPaymentMethods200ResponseDataInnerRelationships**](GETPaymentMethods200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewPOSTPaymentMethods201ResponseData

`func NewPOSTPaymentMethods201ResponseData() *POSTPaymentMethods201ResponseData`

NewPOSTPaymentMethods201ResponseData instantiates a new POSTPaymentMethods201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPaymentMethods201ResponseDataWithDefaults

`func NewPOSTPaymentMethods201ResponseDataWithDefaults() *POSTPaymentMethods201ResponseData`

NewPOSTPaymentMethods201ResponseDataWithDefaults instantiates a new POSTPaymentMethods201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTPaymentMethods201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTPaymentMethods201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTPaymentMethods201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTPaymentMethods201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTPaymentMethods201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTPaymentMethods201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTPaymentMethods201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTPaymentMethods201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTPaymentMethods201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTPaymentMethods201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTPaymentMethods201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTPaymentMethods201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTPaymentMethods201ResponseData) GetAttributes() POSTPaymentMethods201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTPaymentMethods201ResponseData) GetAttributesOk() (*POSTPaymentMethods201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTPaymentMethods201ResponseData) SetAttributes(v POSTPaymentMethods201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTPaymentMethods201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTPaymentMethods201ResponseData) GetRelationships() GETPaymentMethods200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTPaymentMethods201ResponseData) GetRelationshipsOk() (*GETPaymentMethods200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTPaymentMethods201ResponseData) SetRelationships(v GETPaymentMethods200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTPaymentMethods201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


