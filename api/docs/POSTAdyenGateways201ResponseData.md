# POSTAdyenGateways201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "adyen_gateways"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTAdyenGateways201ResponseDataAttributes**](POSTAdyenGateways201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**POSTAdyenGateways201ResponseDataRelationships**](POSTAdyenGateways201ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTAdyenGateways201ResponseData

`func NewPOSTAdyenGateways201ResponseData() *POSTAdyenGateways201ResponseData`

NewPOSTAdyenGateways201ResponseData instantiates a new POSTAdyenGateways201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTAdyenGateways201ResponseDataWithDefaults

`func NewPOSTAdyenGateways201ResponseDataWithDefaults() *POSTAdyenGateways201ResponseData`

NewPOSTAdyenGateways201ResponseDataWithDefaults instantiates a new POSTAdyenGateways201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTAdyenGateways201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTAdyenGateways201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTAdyenGateways201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTAdyenGateways201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTAdyenGateways201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTAdyenGateways201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTAdyenGateways201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTAdyenGateways201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTAdyenGateways201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTAdyenGateways201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTAdyenGateways201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTAdyenGateways201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTAdyenGateways201ResponseData) GetAttributes() POSTAdyenGateways201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTAdyenGateways201ResponseData) GetAttributesOk() (*POSTAdyenGateways201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTAdyenGateways201ResponseData) SetAttributes(v POSTAdyenGateways201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTAdyenGateways201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTAdyenGateways201ResponseData) GetRelationships() POSTAdyenGateways201ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTAdyenGateways201ResponseData) GetRelationshipsOk() (*POSTAdyenGateways201ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTAdyenGateways201ResponseData) SetRelationships(v POSTAdyenGateways201ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTAdyenGateways201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


