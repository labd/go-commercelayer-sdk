# POSTBundles201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets**](GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets.md) |  | [optional] 
**SkuList** | [**GETBundles200ResponseDataInnerRelationshipsSkuList**](GETBundles200ResponseDataInnerRelationshipsSkuList.md) |  | 

## Methods

### NewPOSTBundles201ResponseDataRelationships

`func NewPOSTBundles201ResponseDataRelationships(skuList GETBundles200ResponseDataInnerRelationshipsSkuList, ) *POSTBundles201ResponseDataRelationships`

NewPOSTBundles201ResponseDataRelationships instantiates a new POSTBundles201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTBundles201ResponseDataRelationshipsWithDefaults

`func NewPOSTBundles201ResponseDataRelationshipsWithDefaults() *POSTBundles201ResponseDataRelationships`

NewPOSTBundles201ResponseDataRelationshipsWithDefaults instantiates a new POSTBundles201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTBundles201ResponseDataRelationships) GetMarket() GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTBundles201ResponseDataRelationships) GetMarketOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTBundles201ResponseDataRelationships) SetMarket(v GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTBundles201ResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetSkuList

`func (o *POSTBundles201ResponseDataRelationships) GetSkuList() GETBundles200ResponseDataInnerRelationshipsSkuList`

GetSkuList returns the SkuList field if non-nil, zero value otherwise.

### GetSkuListOk

`func (o *POSTBundles201ResponseDataRelationships) GetSkuListOk() (*GETBundles200ResponseDataInnerRelationshipsSkuList, bool)`

GetSkuListOk returns a tuple with the SkuList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuList

`func (o *POSTBundles201ResponseDataRelationships) SetSkuList(v GETBundles200ResponseDataInnerRelationshipsSkuList)`

SetSkuList sets SkuList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


