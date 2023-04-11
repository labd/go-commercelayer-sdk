# POSTPricesRequestDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceList** | [**POSTMarketsRequestDataRelationshipsPriceList**](POSTMarketsRequestDataRelationshipsPriceList.md) |  | 
**Sku** | [**POSTInStockSubscriptionsRequestDataRelationshipsSku**](POSTInStockSubscriptionsRequestDataRelationshipsSku.md) |  | 
**PriceTiers** | Pointer to [**POSTPricesRequestDataRelationshipsPriceTiers**](POSTPricesRequestDataRelationshipsPriceTiers.md) |  | [optional] 

## Methods

### NewPOSTPricesRequestDataRelationships

`func NewPOSTPricesRequestDataRelationships(priceList POSTMarketsRequestDataRelationshipsPriceList, sku POSTInStockSubscriptionsRequestDataRelationshipsSku, ) *POSTPricesRequestDataRelationships`

NewPOSTPricesRequestDataRelationships instantiates a new POSTPricesRequestDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPricesRequestDataRelationshipsWithDefaults

`func NewPOSTPricesRequestDataRelationshipsWithDefaults() *POSTPricesRequestDataRelationships`

NewPOSTPricesRequestDataRelationshipsWithDefaults instantiates a new POSTPricesRequestDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceList

`func (o *POSTPricesRequestDataRelationships) GetPriceList() POSTMarketsRequestDataRelationshipsPriceList`

GetPriceList returns the PriceList field if non-nil, zero value otherwise.

### GetPriceListOk

`func (o *POSTPricesRequestDataRelationships) GetPriceListOk() (*POSTMarketsRequestDataRelationshipsPriceList, bool)`

GetPriceListOk returns a tuple with the PriceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceList

`func (o *POSTPricesRequestDataRelationships) SetPriceList(v POSTMarketsRequestDataRelationshipsPriceList)`

SetPriceList sets PriceList field to given value.


### GetSku

`func (o *POSTPricesRequestDataRelationships) GetSku() POSTInStockSubscriptionsRequestDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *POSTPricesRequestDataRelationships) GetSkuOk() (*POSTInStockSubscriptionsRequestDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *POSTPricesRequestDataRelationships) SetSku(v POSTInStockSubscriptionsRequestDataRelationshipsSku)`

SetSku sets Sku field to given value.


### GetPriceTiers

`func (o *POSTPricesRequestDataRelationships) GetPriceTiers() POSTPricesRequestDataRelationshipsPriceTiers`

GetPriceTiers returns the PriceTiers field if non-nil, zero value otherwise.

### GetPriceTiersOk

`func (o *POSTPricesRequestDataRelationships) GetPriceTiersOk() (*POSTPricesRequestDataRelationshipsPriceTiers, bool)`

GetPriceTiersOk returns a tuple with the PriceTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTiers

`func (o *POSTPricesRequestDataRelationships) SetPriceTiers(v POSTPricesRequestDataRelationshipsPriceTiers)`

SetPriceTiers sets PriceTiers field to given value.

### HasPriceTiers

`func (o *POSTPricesRequestDataRelationships) HasPriceTiers() bool`

HasPriceTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


