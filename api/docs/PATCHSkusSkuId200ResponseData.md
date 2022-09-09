# PATCHSkusSkuId200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "skus"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**PATCHSkusSkuId200ResponseDataAttributes**](PATCHSkusSkuId200ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETSkus200ResponseDataInnerRelationships**](GETSkus200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewPATCHSkusSkuId200ResponseData

`func NewPATCHSkusSkuId200ResponseData() *PATCHSkusSkuId200ResponseData`

NewPATCHSkusSkuId200ResponseData instantiates a new PATCHSkusSkuId200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHSkusSkuId200ResponseDataWithDefaults

`func NewPATCHSkusSkuId200ResponseDataWithDefaults() *PATCHSkusSkuId200ResponseData`

NewPATCHSkusSkuId200ResponseDataWithDefaults instantiates a new PATCHSkusSkuId200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PATCHSkusSkuId200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PATCHSkusSkuId200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PATCHSkusSkuId200ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PATCHSkusSkuId200ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PATCHSkusSkuId200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PATCHSkusSkuId200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PATCHSkusSkuId200ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PATCHSkusSkuId200ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *PATCHSkusSkuId200ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PATCHSkusSkuId200ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PATCHSkusSkuId200ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PATCHSkusSkuId200ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *PATCHSkusSkuId200ResponseData) GetAttributes() PATCHSkusSkuId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PATCHSkusSkuId200ResponseData) GetAttributesOk() (*PATCHSkusSkuId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PATCHSkusSkuId200ResponseData) SetAttributes(v PATCHSkusSkuId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PATCHSkusSkuId200ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *PATCHSkusSkuId200ResponseData) GetRelationships() GETSkus200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PATCHSkusSkuId200ResponseData) GetRelationshipsOk() (*GETSkus200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PATCHSkusSkuId200ResponseData) SetRelationships(v GETSkus200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PATCHSkusSkuId200ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


