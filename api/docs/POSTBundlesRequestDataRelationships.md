# POSTBundlesRequestDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**POSTBillingInfoValidationRulesRequestDataRelationshipsMarket**](POSTBillingInfoValidationRulesRequestDataRelationshipsMarket.md) |  | [optional] 
**SkuList** | [**POSTBundlesRequestDataRelationshipsSkuList**](POSTBundlesRequestDataRelationshipsSkuList.md) |  | 

## Methods

### NewPOSTBundlesRequestDataRelationships

`func NewPOSTBundlesRequestDataRelationships(skuList POSTBundlesRequestDataRelationshipsSkuList, ) *POSTBundlesRequestDataRelationships`

NewPOSTBundlesRequestDataRelationships instantiates a new POSTBundlesRequestDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTBundlesRequestDataRelationshipsWithDefaults

`func NewPOSTBundlesRequestDataRelationshipsWithDefaults() *POSTBundlesRequestDataRelationships`

NewPOSTBundlesRequestDataRelationshipsWithDefaults instantiates a new POSTBundlesRequestDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTBundlesRequestDataRelationships) GetMarket() POSTBillingInfoValidationRulesRequestDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTBundlesRequestDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRulesRequestDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTBundlesRequestDataRelationships) SetMarket(v POSTBillingInfoValidationRulesRequestDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTBundlesRequestDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetSkuList

`func (o *POSTBundlesRequestDataRelationships) GetSkuList() POSTBundlesRequestDataRelationshipsSkuList`

GetSkuList returns the SkuList field if non-nil, zero value otherwise.

### GetSkuListOk

`func (o *POSTBundlesRequestDataRelationships) GetSkuListOk() (*POSTBundlesRequestDataRelationshipsSkuList, bool)`

GetSkuListOk returns a tuple with the SkuList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuList

`func (o *POSTBundlesRequestDataRelationships) SetSkuList(v POSTBundlesRequestDataRelationshipsSkuList)`

SetSkuList sets SkuList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


