# PATCHMarketsMarketIdRequestDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Merchant** | Pointer to [**POSTMarketsRequestDataRelationshipsMerchant**](POSTMarketsRequestDataRelationshipsMerchant.md) |  | [optional] 
**PriceList** | Pointer to [**POSTMarketsRequestDataRelationshipsPriceList**](POSTMarketsRequestDataRelationshipsPriceList.md) |  | [optional] 
**InventoryModel** | Pointer to [**POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel**](POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel.md) |  | [optional] 
**SubscriptionModel** | Pointer to [**POSTMarketsRequestDataRelationshipsSubscriptionModel**](POSTMarketsRequestDataRelationshipsSubscriptionModel.md) |  | [optional] 
**TaxCalculator** | Pointer to [**POSTMarketsRequestDataRelationshipsTaxCalculator**](POSTMarketsRequestDataRelationshipsTaxCalculator.md) |  | [optional] 
**CustomerGroup** | Pointer to [**POSTCustomersRequestDataRelationshipsCustomerGroup**](POSTCustomersRequestDataRelationshipsCustomerGroup.md) |  | [optional] 

## Methods

### NewPATCHMarketsMarketIdRequestDataRelationships

`func NewPATCHMarketsMarketIdRequestDataRelationships() *PATCHMarketsMarketIdRequestDataRelationships`

NewPATCHMarketsMarketIdRequestDataRelationships instantiates a new PATCHMarketsMarketIdRequestDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHMarketsMarketIdRequestDataRelationshipsWithDefaults

`func NewPATCHMarketsMarketIdRequestDataRelationshipsWithDefaults() *PATCHMarketsMarketIdRequestDataRelationships`

NewPATCHMarketsMarketIdRequestDataRelationshipsWithDefaults instantiates a new PATCHMarketsMarketIdRequestDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchant

`func (o *PATCHMarketsMarketIdRequestDataRelationships) GetMerchant() POSTMarketsRequestDataRelationshipsMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *PATCHMarketsMarketIdRequestDataRelationships) GetMerchantOk() (*POSTMarketsRequestDataRelationshipsMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *PATCHMarketsMarketIdRequestDataRelationships) SetMerchant(v POSTMarketsRequestDataRelationshipsMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *PATCHMarketsMarketIdRequestDataRelationships) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetPriceList

`func (o *PATCHMarketsMarketIdRequestDataRelationships) GetPriceList() POSTMarketsRequestDataRelationshipsPriceList`

GetPriceList returns the PriceList field if non-nil, zero value otherwise.

### GetPriceListOk

`func (o *PATCHMarketsMarketIdRequestDataRelationships) GetPriceListOk() (*POSTMarketsRequestDataRelationshipsPriceList, bool)`

GetPriceListOk returns a tuple with the PriceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceList

`func (o *PATCHMarketsMarketIdRequestDataRelationships) SetPriceList(v POSTMarketsRequestDataRelationshipsPriceList)`

SetPriceList sets PriceList field to given value.

### HasPriceList

`func (o *PATCHMarketsMarketIdRequestDataRelationships) HasPriceList() bool`

HasPriceList returns a boolean if a field has been set.

### GetInventoryModel

`func (o *PATCHMarketsMarketIdRequestDataRelationships) GetInventoryModel() POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel`

GetInventoryModel returns the InventoryModel field if non-nil, zero value otherwise.

### GetInventoryModelOk

`func (o *PATCHMarketsMarketIdRequestDataRelationships) GetInventoryModelOk() (*POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel, bool)`

GetInventoryModelOk returns a tuple with the InventoryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryModel

`func (o *PATCHMarketsMarketIdRequestDataRelationships) SetInventoryModel(v POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel)`

SetInventoryModel sets InventoryModel field to given value.

### HasInventoryModel

`func (o *PATCHMarketsMarketIdRequestDataRelationships) HasInventoryModel() bool`

HasInventoryModel returns a boolean if a field has been set.

### GetSubscriptionModel

`func (o *PATCHMarketsMarketIdRequestDataRelationships) GetSubscriptionModel() POSTMarketsRequestDataRelationshipsSubscriptionModel`

GetSubscriptionModel returns the SubscriptionModel field if non-nil, zero value otherwise.

### GetSubscriptionModelOk

`func (o *PATCHMarketsMarketIdRequestDataRelationships) GetSubscriptionModelOk() (*POSTMarketsRequestDataRelationshipsSubscriptionModel, bool)`

GetSubscriptionModelOk returns a tuple with the SubscriptionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionModel

`func (o *PATCHMarketsMarketIdRequestDataRelationships) SetSubscriptionModel(v POSTMarketsRequestDataRelationshipsSubscriptionModel)`

SetSubscriptionModel sets SubscriptionModel field to given value.

### HasSubscriptionModel

`func (o *PATCHMarketsMarketIdRequestDataRelationships) HasSubscriptionModel() bool`

HasSubscriptionModel returns a boolean if a field has been set.

### GetTaxCalculator

`func (o *PATCHMarketsMarketIdRequestDataRelationships) GetTaxCalculator() POSTMarketsRequestDataRelationshipsTaxCalculator`

GetTaxCalculator returns the TaxCalculator field if non-nil, zero value otherwise.

### GetTaxCalculatorOk

`func (o *PATCHMarketsMarketIdRequestDataRelationships) GetTaxCalculatorOk() (*POSTMarketsRequestDataRelationshipsTaxCalculator, bool)`

GetTaxCalculatorOk returns a tuple with the TaxCalculator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculator

`func (o *PATCHMarketsMarketIdRequestDataRelationships) SetTaxCalculator(v POSTMarketsRequestDataRelationshipsTaxCalculator)`

SetTaxCalculator sets TaxCalculator field to given value.

### HasTaxCalculator

`func (o *PATCHMarketsMarketIdRequestDataRelationships) HasTaxCalculator() bool`

HasTaxCalculator returns a boolean if a field has been set.

### GetCustomerGroup

`func (o *PATCHMarketsMarketIdRequestDataRelationships) GetCustomerGroup() POSTCustomersRequestDataRelationshipsCustomerGroup`

GetCustomerGroup returns the CustomerGroup field if non-nil, zero value otherwise.

### GetCustomerGroupOk

`func (o *PATCHMarketsMarketIdRequestDataRelationships) GetCustomerGroupOk() (*POSTCustomersRequestDataRelationshipsCustomerGroup, bool)`

GetCustomerGroupOk returns a tuple with the CustomerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerGroup

`func (o *PATCHMarketsMarketIdRequestDataRelationships) SetCustomerGroup(v POSTCustomersRequestDataRelationshipsCustomerGroup)`

SetCustomerGroup sets CustomerGroup field to given value.

### HasCustomerGroup

`func (o *PATCHMarketsMarketIdRequestDataRelationships) HasCustomerGroup() bool`

HasCustomerGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


