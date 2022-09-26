# ReturnLineItemResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**AddressResponseDataLinks**](AddressResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**ReturnLineItemResponseDataRelationships**](ReturnLineItemResponseDataRelationships.md) |  | [optional] 

## Methods

### NewReturnLineItemResponseData

`func NewReturnLineItemResponseData() *ReturnLineItemResponseData`

NewReturnLineItemResponseData instantiates a new ReturnLineItemResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnLineItemResponseDataWithDefaults

`func NewReturnLineItemResponseDataWithDefaults() *ReturnLineItemResponseData`

NewReturnLineItemResponseDataWithDefaults instantiates a new ReturnLineItemResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReturnLineItemResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReturnLineItemResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReturnLineItemResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReturnLineItemResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ReturnLineItemResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReturnLineItemResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReturnLineItemResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReturnLineItemResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *ReturnLineItemResponseData) GetLinks() AddressResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReturnLineItemResponseData) GetLinksOk() (*AddressResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReturnLineItemResponseData) SetLinks(v AddressResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ReturnLineItemResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *ReturnLineItemResponseData) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ReturnLineItemResponseData) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ReturnLineItemResponseData) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ReturnLineItemResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *ReturnLineItemResponseData) GetRelationships() ReturnLineItemResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ReturnLineItemResponseData) GetRelationshipsOk() (*ReturnLineItemResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ReturnLineItemResponseData) SetRelationships(v ReturnLineItemResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ReturnLineItemResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


