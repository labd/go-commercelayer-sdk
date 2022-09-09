# MarketDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Merchant** | Pointer to [**MarketDataRelationshipsMerchant**](MarketDataRelationshipsMerchant.md) |  | [optional] 
**PriceList** | Pointer to [**MarketDataRelationshipsPriceList**](MarketDataRelationshipsPriceList.md) |  | [optional] 
**InventoryModel** | Pointer to [**InventoryReturnLocationDataRelationshipsInventoryModel**](InventoryReturnLocationDataRelationshipsInventoryModel.md) |  | [optional] 
**TaxCalculator** | Pointer to [**MarketDataRelationshipsTaxCalculator**](MarketDataRelationshipsTaxCalculator.md) |  | [optional] 
**CustomerGroup** | Pointer to [**CustomerDataRelationshipsCustomerGroup**](CustomerDataRelationshipsCustomerGroup.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountDataRelationshipsAttachments**](AvalaraAccountDataRelationshipsAttachments.md) |  | [optional] 

## Methods

### NewMarketDataRelationships

`func NewMarketDataRelationships() *MarketDataRelationships`

NewMarketDataRelationships instantiates a new MarketDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketDataRelationshipsWithDefaults

`func NewMarketDataRelationshipsWithDefaults() *MarketDataRelationships`

NewMarketDataRelationshipsWithDefaults instantiates a new MarketDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchant

`func (o *MarketDataRelationships) GetMerchant() MarketDataRelationshipsMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *MarketDataRelationships) GetMerchantOk() (*MarketDataRelationshipsMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *MarketDataRelationships) SetMerchant(v MarketDataRelationshipsMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *MarketDataRelationships) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetPriceList

`func (o *MarketDataRelationships) GetPriceList() MarketDataRelationshipsPriceList`

GetPriceList returns the PriceList field if non-nil, zero value otherwise.

### GetPriceListOk

`func (o *MarketDataRelationships) GetPriceListOk() (*MarketDataRelationshipsPriceList, bool)`

GetPriceListOk returns a tuple with the PriceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceList

`func (o *MarketDataRelationships) SetPriceList(v MarketDataRelationshipsPriceList)`

SetPriceList sets PriceList field to given value.

### HasPriceList

`func (o *MarketDataRelationships) HasPriceList() bool`

HasPriceList returns a boolean if a field has been set.

### GetInventoryModel

`func (o *MarketDataRelationships) GetInventoryModel() InventoryReturnLocationDataRelationshipsInventoryModel`

GetInventoryModel returns the InventoryModel field if non-nil, zero value otherwise.

### GetInventoryModelOk

`func (o *MarketDataRelationships) GetInventoryModelOk() (*InventoryReturnLocationDataRelationshipsInventoryModel, bool)`

GetInventoryModelOk returns a tuple with the InventoryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryModel

`func (o *MarketDataRelationships) SetInventoryModel(v InventoryReturnLocationDataRelationshipsInventoryModel)`

SetInventoryModel sets InventoryModel field to given value.

### HasInventoryModel

`func (o *MarketDataRelationships) HasInventoryModel() bool`

HasInventoryModel returns a boolean if a field has been set.

### GetTaxCalculator

`func (o *MarketDataRelationships) GetTaxCalculator() MarketDataRelationshipsTaxCalculator`

GetTaxCalculator returns the TaxCalculator field if non-nil, zero value otherwise.

### GetTaxCalculatorOk

`func (o *MarketDataRelationships) GetTaxCalculatorOk() (*MarketDataRelationshipsTaxCalculator, bool)`

GetTaxCalculatorOk returns a tuple with the TaxCalculator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculator

`func (o *MarketDataRelationships) SetTaxCalculator(v MarketDataRelationshipsTaxCalculator)`

SetTaxCalculator sets TaxCalculator field to given value.

### HasTaxCalculator

`func (o *MarketDataRelationships) HasTaxCalculator() bool`

HasTaxCalculator returns a boolean if a field has been set.

### GetCustomerGroup

`func (o *MarketDataRelationships) GetCustomerGroup() CustomerDataRelationshipsCustomerGroup`

GetCustomerGroup returns the CustomerGroup field if non-nil, zero value otherwise.

### GetCustomerGroupOk

`func (o *MarketDataRelationships) GetCustomerGroupOk() (*CustomerDataRelationshipsCustomerGroup, bool)`

GetCustomerGroupOk returns a tuple with the CustomerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerGroup

`func (o *MarketDataRelationships) SetCustomerGroup(v CustomerDataRelationshipsCustomerGroup)`

SetCustomerGroup sets CustomerGroup field to given value.

### HasCustomerGroup

`func (o *MarketDataRelationships) HasCustomerGroup() bool`

HasCustomerGroup returns a boolean if a field has been set.

### GetAttachments

`func (o *MarketDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MarketDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MarketDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *MarketDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


