# CaptureResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**AddressResponseDataLinks**](AddressResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**CaptureResponseDataRelationships**](CaptureResponseDataRelationships.md) |  | [optional] 

## Methods

### NewCaptureResponseData

`func NewCaptureResponseData() *CaptureResponseData`

NewCaptureResponseData instantiates a new CaptureResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureResponseDataWithDefaults

`func NewCaptureResponseDataWithDefaults() *CaptureResponseData`

NewCaptureResponseDataWithDefaults instantiates a new CaptureResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CaptureResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptureResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptureResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptureResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CaptureResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptureResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptureResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CaptureResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *CaptureResponseData) GetLinks() AddressResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CaptureResponseData) GetLinksOk() (*AddressResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CaptureResponseData) SetLinks(v AddressResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CaptureResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *CaptureResponseData) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CaptureResponseData) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CaptureResponseData) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CaptureResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *CaptureResponseData) GetRelationships() CaptureResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CaptureResponseData) GetRelationshipsOk() (*CaptureResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CaptureResponseData) SetRelationships(v CaptureResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CaptureResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


