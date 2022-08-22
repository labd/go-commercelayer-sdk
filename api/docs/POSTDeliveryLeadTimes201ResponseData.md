# POSTDeliveryLeadTimes201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "delivery_lead_times"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTDeliveryLeadTimes201ResponseDataAttributes**](POSTDeliveryLeadTimes201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**POSTDeliveryLeadTimes201ResponseDataRelationships**](POSTDeliveryLeadTimes201ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTDeliveryLeadTimes201ResponseData

`func NewPOSTDeliveryLeadTimes201ResponseData() *POSTDeliveryLeadTimes201ResponseData`

NewPOSTDeliveryLeadTimes201ResponseData instantiates a new POSTDeliveryLeadTimes201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTDeliveryLeadTimes201ResponseDataWithDefaults

`func NewPOSTDeliveryLeadTimes201ResponseDataWithDefaults() *POSTDeliveryLeadTimes201ResponseData`

NewPOSTDeliveryLeadTimes201ResponseDataWithDefaults instantiates a new POSTDeliveryLeadTimes201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTDeliveryLeadTimes201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTDeliveryLeadTimes201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTDeliveryLeadTimes201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTDeliveryLeadTimes201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTDeliveryLeadTimes201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTDeliveryLeadTimes201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTDeliveryLeadTimes201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTDeliveryLeadTimes201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTDeliveryLeadTimes201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTDeliveryLeadTimes201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTDeliveryLeadTimes201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTDeliveryLeadTimes201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTDeliveryLeadTimes201ResponseData) GetAttributes() POSTDeliveryLeadTimes201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTDeliveryLeadTimes201ResponseData) GetAttributesOk() (*POSTDeliveryLeadTimes201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTDeliveryLeadTimes201ResponseData) SetAttributes(v POSTDeliveryLeadTimes201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTDeliveryLeadTimes201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTDeliveryLeadTimes201ResponseData) GetRelationships() POSTDeliveryLeadTimes201ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTDeliveryLeadTimes201ResponseData) GetRelationshipsOk() (*POSTDeliveryLeadTimes201ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTDeliveryLeadTimes201ResponseData) SetRelationships(v POSTDeliveryLeadTimes201ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTDeliveryLeadTimes201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


