# POSTMarketsRequestDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Merchant** | [**POSTMarketsRequestDataRelationshipsMerchant**](POSTMarketsRequestDataRelationshipsMerchant.md) |  | 
**PriceList** | [**POSTMarketsRequestDataRelationshipsPriceList**](POSTMarketsRequestDataRelationshipsPriceList.md) |  | 
**InventoryModel** | [**POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel**](POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel.md) |  | 
**SubscriptionModel** | Pointer to [**POSTMarketsRequestDataRelationshipsSubscriptionModel**](POSTMarketsRequestDataRelationshipsSubscriptionModel.md) |  | [optional] 
**TaxCalculator** | Pointer to [**POSTMarketsRequestDataRelationshipsTaxCalculator**](POSTMarketsRequestDataRelationshipsTaxCalculator.md) |  | [optional] 
**CustomerGroup** | Pointer to [**POSTCustomersRequestDataRelationshipsCustomerGroup**](POSTCustomersRequestDataRelationshipsCustomerGroup.md) |  | [optional] 

## Methods

### NewPOSTMarketsRequestDataRelationships

`func NewPOSTMarketsRequestDataRelationships(merchant POSTMarketsRequestDataRelationshipsMerchant, priceList POSTMarketsRequestDataRelationshipsPriceList, inventoryModel POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel, ) *POSTMarketsRequestDataRelationships`

NewPOSTMarketsRequestDataRelationships instantiates a new POSTMarketsRequestDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTMarketsRequestDataRelationshipsWithDefaults

`func NewPOSTMarketsRequestDataRelationshipsWithDefaults() *POSTMarketsRequestDataRelationships`

NewPOSTMarketsRequestDataRelationshipsWithDefaults instantiates a new POSTMarketsRequestDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchant

`func (o *POSTMarketsRequestDataRelationships) GetMerchant() POSTMarketsRequestDataRelationshipsMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *POSTMarketsRequestDataRelationships) GetMerchantOk() (*POSTMarketsRequestDataRelationshipsMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *POSTMarketsRequestDataRelationships) SetMerchant(v POSTMarketsRequestDataRelationshipsMerchant)`

SetMerchant sets Merchant field to given value.


### GetPriceList

`func (o *POSTMarketsRequestDataRelationships) GetPriceList() POSTMarketsRequestDataRelationshipsPriceList`

GetPriceList returns the PriceList field if non-nil, zero value otherwise.

### GetPriceListOk

`func (o *POSTMarketsRequestDataRelationships) GetPriceListOk() (*POSTMarketsRequestDataRelationshipsPriceList, bool)`

GetPriceListOk returns a tuple with the PriceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceList

`func (o *POSTMarketsRequestDataRelationships) SetPriceList(v POSTMarketsRequestDataRelationshipsPriceList)`

SetPriceList sets PriceList field to given value.


### GetInventoryModel

`func (o *POSTMarketsRequestDataRelationships) GetInventoryModel() POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel`

GetInventoryModel returns the InventoryModel field if non-nil, zero value otherwise.

### GetInventoryModelOk

`func (o *POSTMarketsRequestDataRelationships) GetInventoryModelOk() (*POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel, bool)`

GetInventoryModelOk returns a tuple with the InventoryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryModel

`func (o *POSTMarketsRequestDataRelationships) SetInventoryModel(v POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel)`

SetInventoryModel sets InventoryModel field to given value.


### GetSubscriptionModel

`func (o *POSTMarketsRequestDataRelationships) GetSubscriptionModel() POSTMarketsRequestDataRelationshipsSubscriptionModel`

GetSubscriptionModel returns the SubscriptionModel field if non-nil, zero value otherwise.

### GetSubscriptionModelOk

`func (o *POSTMarketsRequestDataRelationships) GetSubscriptionModelOk() (*POSTMarketsRequestDataRelationshipsSubscriptionModel, bool)`

GetSubscriptionModelOk returns a tuple with the SubscriptionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionModel

`func (o *POSTMarketsRequestDataRelationships) SetSubscriptionModel(v POSTMarketsRequestDataRelationshipsSubscriptionModel)`

SetSubscriptionModel sets SubscriptionModel field to given value.

### HasSubscriptionModel

`func (o *POSTMarketsRequestDataRelationships) HasSubscriptionModel() bool`

HasSubscriptionModel returns a boolean if a field has been set.

### GetTaxCalculator

`func (o *POSTMarketsRequestDataRelationships) GetTaxCalculator() POSTMarketsRequestDataRelationshipsTaxCalculator`

GetTaxCalculator returns the TaxCalculator field if non-nil, zero value otherwise.

### GetTaxCalculatorOk

`func (o *POSTMarketsRequestDataRelationships) GetTaxCalculatorOk() (*POSTMarketsRequestDataRelationshipsTaxCalculator, bool)`

GetTaxCalculatorOk returns a tuple with the TaxCalculator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculator

`func (o *POSTMarketsRequestDataRelationships) SetTaxCalculator(v POSTMarketsRequestDataRelationshipsTaxCalculator)`

SetTaxCalculator sets TaxCalculator field to given value.

### HasTaxCalculator

`func (o *POSTMarketsRequestDataRelationships) HasTaxCalculator() bool`

HasTaxCalculator returns a boolean if a field has been set.

### GetCustomerGroup

`func (o *POSTMarketsRequestDataRelationships) GetCustomerGroup() POSTCustomersRequestDataRelationshipsCustomerGroup`

GetCustomerGroup returns the CustomerGroup field if non-nil, zero value otherwise.

### GetCustomerGroupOk

`func (o *POSTMarketsRequestDataRelationships) GetCustomerGroupOk() (*POSTCustomersRequestDataRelationshipsCustomerGroup, bool)`

GetCustomerGroupOk returns a tuple with the CustomerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerGroup

`func (o *POSTMarketsRequestDataRelationships) SetCustomerGroup(v POSTCustomersRequestDataRelationshipsCustomerGroup)`

SetCustomerGroup sets CustomerGroup field to given value.

### HasCustomerGroup

`func (o *POSTMarketsRequestDataRelationships) HasCustomerGroup() bool`

HasCustomerGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


