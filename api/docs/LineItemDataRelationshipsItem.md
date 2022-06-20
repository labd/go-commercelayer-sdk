# LineItemDataRelationshipsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "adjustments"]
**Id** | **string** | The resource&#39;s id | 
**Data** | [**SkuData**](SkuData.md) |  | 

## Methods

### NewLineItemDataRelationshipsItem

`func NewLineItemDataRelationshipsItem(type_ string, id string, data SkuData, ) *LineItemDataRelationshipsItem`

NewLineItemDataRelationshipsItem instantiates a new LineItemDataRelationshipsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemDataRelationshipsItemWithDefaults

`func NewLineItemDataRelationshipsItemWithDefaults() *LineItemDataRelationshipsItem`

NewLineItemDataRelationshipsItemWithDefaults instantiates a new LineItemDataRelationshipsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LineItemDataRelationshipsItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LineItemDataRelationshipsItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LineItemDataRelationshipsItem) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *LineItemDataRelationshipsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LineItemDataRelationshipsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LineItemDataRelationshipsItem) SetId(v string)`

SetId sets Id field to given value.


### GetData

`func (o *LineItemDataRelationshipsItem) GetData() SkuData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LineItemDataRelationshipsItem) GetDataOk() (*SkuData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LineItemDataRelationshipsItem) SetData(v SkuData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


