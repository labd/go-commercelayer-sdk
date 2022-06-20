# BundleCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**AvalaraAccountDataRelationshipsMarkets**](AvalaraAccountDataRelationshipsMarkets.md) |  | [optional] 
**SkuList** | [**BundleDataRelationshipsSkuList**](BundleDataRelationshipsSkuList.md) |  | 

## Methods

### NewBundleCreateDataRelationships

`func NewBundleCreateDataRelationships(skuList BundleDataRelationshipsSkuList, ) *BundleCreateDataRelationships`

NewBundleCreateDataRelationships instantiates a new BundleCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleCreateDataRelationshipsWithDefaults

`func NewBundleCreateDataRelationshipsWithDefaults() *BundleCreateDataRelationships`

NewBundleCreateDataRelationshipsWithDefaults instantiates a new BundleCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *BundleCreateDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *BundleCreateDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *BundleCreateDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *BundleCreateDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetSkuList

`func (o *BundleCreateDataRelationships) GetSkuList() BundleDataRelationshipsSkuList`

GetSkuList returns the SkuList field if non-nil, zero value otherwise.

### GetSkuListOk

`func (o *BundleCreateDataRelationships) GetSkuListOk() (*BundleDataRelationshipsSkuList, bool)`

GetSkuListOk returns a tuple with the SkuList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuList

`func (o *BundleCreateDataRelationships) SetSkuList(v BundleDataRelationshipsSkuList)`

SetSkuList sets SkuList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


