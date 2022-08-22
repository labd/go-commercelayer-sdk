# PriceListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "price_lists"]
**Attributes** | [**GETPriceLists200ResponseDataInnerAttributes**](GETPriceLists200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**GETPriceLists200ResponseDataInnerRelationships**](GETPriceLists200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewPriceListData

`func NewPriceListData(type_ string, attributes GETPriceLists200ResponseDataInnerAttributes, ) *PriceListData`

NewPriceListData instantiates a new PriceListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceListDataWithDefaults

`func NewPriceListDataWithDefaults() *PriceListData`

NewPriceListDataWithDefaults instantiates a new PriceListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceListData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceListData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceListData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PriceListData) GetAttributes() GETPriceLists200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceListData) GetAttributesOk() (*GETPriceLists200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceListData) SetAttributes(v GETPriceLists200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceListData) GetRelationships() GETPriceLists200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceListData) GetRelationshipsOk() (*GETPriceLists200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceListData) SetRelationships(v GETPriceLists200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceListData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


