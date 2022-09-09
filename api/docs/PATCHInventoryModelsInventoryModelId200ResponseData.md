# PATCHInventoryModelsInventoryModelId200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "inventory_models"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**PATCHInventoryModelsInventoryModelId200ResponseDataAttributes**](PATCHInventoryModelsInventoryModelId200ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETInventoryModels200ResponseDataInnerRelationships**](GETInventoryModels200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewPATCHInventoryModelsInventoryModelId200ResponseData

`func NewPATCHInventoryModelsInventoryModelId200ResponseData() *PATCHInventoryModelsInventoryModelId200ResponseData`

NewPATCHInventoryModelsInventoryModelId200ResponseData instantiates a new PATCHInventoryModelsInventoryModelId200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHInventoryModelsInventoryModelId200ResponseDataWithDefaults

`func NewPATCHInventoryModelsInventoryModelId200ResponseDataWithDefaults() *PATCHInventoryModelsInventoryModelId200ResponseData`

NewPATCHInventoryModelsInventoryModelId200ResponseDataWithDefaults instantiates a new PATCHInventoryModelsInventoryModelId200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) GetAttributes() PATCHInventoryModelsInventoryModelId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) GetAttributesOk() (*PATCHInventoryModelsInventoryModelId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) SetAttributes(v PATCHInventoryModelsInventoryModelId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) GetRelationships() GETInventoryModels200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) GetRelationshipsOk() (*GETInventoryModels200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) SetRelationships(v GETInventoryModels200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PATCHInventoryModelsInventoryModelId200ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


