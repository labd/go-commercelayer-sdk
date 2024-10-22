# SkuDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingCategory** | Pointer to [**ShipmentDataRelationshipsShippingCategory**](ShipmentDataRelationshipsShippingCategory.md) |  | [optional] 
**Prices** | Pointer to [**PriceFrequencyTierDataRelationshipsPrice**](PriceFrequencyTierDataRelationshipsPrice.md) |  | [optional] 
**StockItems** | Pointer to [**ReservedStockDataRelationshipsStockItem**](ReservedStockDataRelationshipsStockItem.md) |  | [optional] 
**StockReservations** | Pointer to [**LineItemDataRelationshipsStockReservations**](LineItemDataRelationshipsStockReservations.md) |  | [optional] 
**DeliveryLeadTimes** | Pointer to [**ShipmentDataRelationshipsDeliveryLeadTime**](ShipmentDataRelationshipsDeliveryLeadTime.md) |  | [optional] 
**SkuOptions** | Pointer to [**LineItemOptionDataRelationshipsSkuOption**](LineItemOptionDataRelationshipsSkuOption.md) |  | [optional] 
**SkuListItems** | Pointer to [**SkuListDataRelationshipsSkuListItems**](SkuListDataRelationshipsSkuListItems.md) |  | [optional] 
**SkuLists** | Pointer to [**BundleDataRelationshipsSkuList**](BundleDataRelationshipsSkuList.md) |  | [optional] 
**Attachments** | Pointer to [**AuthorizationDataRelationshipsAttachments**](AuthorizationDataRelationshipsAttachments.md) |  | [optional] 
**Links** | Pointer to [**OrderDataRelationshipsLinks**](OrderDataRelationshipsLinks.md) |  | [optional] 
**Events** | Pointer to [**AddressDataRelationshipsEvents**](AddressDataRelationshipsEvents.md) |  | [optional] 
**Tags** | Pointer to [**AddressDataRelationshipsTags**](AddressDataRelationshipsTags.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 
**JwtCustomer** | Pointer to [**CouponRecipientDataRelationshipsCustomer**](CouponRecipientDataRelationshipsCustomer.md) |  | [optional] 
**JwtMarkets** | Pointer to [**AvalaraAccountDataRelationshipsMarkets**](AvalaraAccountDataRelationshipsMarkets.md) |  | [optional] 
**JwtStockLocations** | Pointer to [**DeliveryLeadTimeDataRelationshipsStockLocation**](DeliveryLeadTimeDataRelationshipsStockLocation.md) |  | [optional] 

## Methods

### NewSkuDataRelationships

`func NewSkuDataRelationships() *SkuDataRelationships`

NewSkuDataRelationships instantiates a new SkuDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuDataRelationshipsWithDefaults

`func NewSkuDataRelationshipsWithDefaults() *SkuDataRelationships`

NewSkuDataRelationshipsWithDefaults instantiates a new SkuDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingCategory

`func (o *SkuDataRelationships) GetShippingCategory() ShipmentDataRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *SkuDataRelationships) GetShippingCategoryOk() (*ShipmentDataRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *SkuDataRelationships) SetShippingCategory(v ShipmentDataRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *SkuDataRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetPrices

`func (o *SkuDataRelationships) GetPrices() PriceFrequencyTierDataRelationshipsPrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *SkuDataRelationships) GetPricesOk() (*PriceFrequencyTierDataRelationshipsPrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *SkuDataRelationships) SetPrices(v PriceFrequencyTierDataRelationshipsPrice)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *SkuDataRelationships) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetStockItems

`func (o *SkuDataRelationships) GetStockItems() ReservedStockDataRelationshipsStockItem`

GetStockItems returns the StockItems field if non-nil, zero value otherwise.

### GetStockItemsOk

`func (o *SkuDataRelationships) GetStockItemsOk() (*ReservedStockDataRelationshipsStockItem, bool)`

GetStockItemsOk returns a tuple with the StockItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockItems

`func (o *SkuDataRelationships) SetStockItems(v ReservedStockDataRelationshipsStockItem)`

SetStockItems sets StockItems field to given value.

### HasStockItems

`func (o *SkuDataRelationships) HasStockItems() bool`

HasStockItems returns a boolean if a field has been set.

### GetStockReservations

`func (o *SkuDataRelationships) GetStockReservations() LineItemDataRelationshipsStockReservations`

GetStockReservations returns the StockReservations field if non-nil, zero value otherwise.

### GetStockReservationsOk

`func (o *SkuDataRelationships) GetStockReservationsOk() (*LineItemDataRelationshipsStockReservations, bool)`

GetStockReservationsOk returns a tuple with the StockReservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockReservations

`func (o *SkuDataRelationships) SetStockReservations(v LineItemDataRelationshipsStockReservations)`

SetStockReservations sets StockReservations field to given value.

### HasStockReservations

`func (o *SkuDataRelationships) HasStockReservations() bool`

HasStockReservations returns a boolean if a field has been set.

### GetDeliveryLeadTimes

`func (o *SkuDataRelationships) GetDeliveryLeadTimes() ShipmentDataRelationshipsDeliveryLeadTime`

GetDeliveryLeadTimes returns the DeliveryLeadTimes field if non-nil, zero value otherwise.

### GetDeliveryLeadTimesOk

`func (o *SkuDataRelationships) GetDeliveryLeadTimesOk() (*ShipmentDataRelationshipsDeliveryLeadTime, bool)`

GetDeliveryLeadTimesOk returns a tuple with the DeliveryLeadTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLeadTimes

`func (o *SkuDataRelationships) SetDeliveryLeadTimes(v ShipmentDataRelationshipsDeliveryLeadTime)`

SetDeliveryLeadTimes sets DeliveryLeadTimes field to given value.

### HasDeliveryLeadTimes

`func (o *SkuDataRelationships) HasDeliveryLeadTimes() bool`

HasDeliveryLeadTimes returns a boolean if a field has been set.

### GetSkuOptions

`func (o *SkuDataRelationships) GetSkuOptions() LineItemOptionDataRelationshipsSkuOption`

GetSkuOptions returns the SkuOptions field if non-nil, zero value otherwise.

### GetSkuOptionsOk

`func (o *SkuDataRelationships) GetSkuOptionsOk() (*LineItemOptionDataRelationshipsSkuOption, bool)`

GetSkuOptionsOk returns a tuple with the SkuOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuOptions

`func (o *SkuDataRelationships) SetSkuOptions(v LineItemOptionDataRelationshipsSkuOption)`

SetSkuOptions sets SkuOptions field to given value.

### HasSkuOptions

`func (o *SkuDataRelationships) HasSkuOptions() bool`

HasSkuOptions returns a boolean if a field has been set.

### GetSkuListItems

`func (o *SkuDataRelationships) GetSkuListItems() SkuListDataRelationshipsSkuListItems`

GetSkuListItems returns the SkuListItems field if non-nil, zero value otherwise.

### GetSkuListItemsOk

`func (o *SkuDataRelationships) GetSkuListItemsOk() (*SkuListDataRelationshipsSkuListItems, bool)`

GetSkuListItemsOk returns a tuple with the SkuListItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuListItems

`func (o *SkuDataRelationships) SetSkuListItems(v SkuListDataRelationshipsSkuListItems)`

SetSkuListItems sets SkuListItems field to given value.

### HasSkuListItems

`func (o *SkuDataRelationships) HasSkuListItems() bool`

HasSkuListItems returns a boolean if a field has been set.

### GetSkuLists

`func (o *SkuDataRelationships) GetSkuLists() BundleDataRelationshipsSkuList`

GetSkuLists returns the SkuLists field if non-nil, zero value otherwise.

### GetSkuListsOk

`func (o *SkuDataRelationships) GetSkuListsOk() (*BundleDataRelationshipsSkuList, bool)`

GetSkuListsOk returns a tuple with the SkuLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuLists

`func (o *SkuDataRelationships) SetSkuLists(v BundleDataRelationshipsSkuList)`

SetSkuLists sets SkuLists field to given value.

### HasSkuLists

`func (o *SkuDataRelationships) HasSkuLists() bool`

HasSkuLists returns a boolean if a field has been set.

### GetAttachments

`func (o *SkuDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SkuDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SkuDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SkuDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetLinks

`func (o *SkuDataRelationships) GetLinks() OrderDataRelationshipsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SkuDataRelationships) GetLinksOk() (*OrderDataRelationshipsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SkuDataRelationships) SetLinks(v OrderDataRelationshipsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SkuDataRelationships) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEvents

`func (o *SkuDataRelationships) GetEvents() AddressDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SkuDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SkuDataRelationships) SetEvents(v AddressDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SkuDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTags

`func (o *SkuDataRelationships) GetTags() AddressDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SkuDataRelationships) GetTagsOk() (*AddressDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SkuDataRelationships) SetTags(v AddressDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SkuDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersions

`func (o *SkuDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *SkuDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *SkuDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *SkuDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetJwtCustomer

`func (o *SkuDataRelationships) GetJwtCustomer() CouponRecipientDataRelationshipsCustomer`

GetJwtCustomer returns the JwtCustomer field if non-nil, zero value otherwise.

### GetJwtCustomerOk

`func (o *SkuDataRelationships) GetJwtCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool)`

GetJwtCustomerOk returns a tuple with the JwtCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtCustomer

`func (o *SkuDataRelationships) SetJwtCustomer(v CouponRecipientDataRelationshipsCustomer)`

SetJwtCustomer sets JwtCustomer field to given value.

### HasJwtCustomer

`func (o *SkuDataRelationships) HasJwtCustomer() bool`

HasJwtCustomer returns a boolean if a field has been set.

### GetJwtMarkets

`func (o *SkuDataRelationships) GetJwtMarkets() AvalaraAccountDataRelationshipsMarkets`

GetJwtMarkets returns the JwtMarkets field if non-nil, zero value otherwise.

### GetJwtMarketsOk

`func (o *SkuDataRelationships) GetJwtMarketsOk() (*AvalaraAccountDataRelationshipsMarkets, bool)`

GetJwtMarketsOk returns a tuple with the JwtMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtMarkets

`func (o *SkuDataRelationships) SetJwtMarkets(v AvalaraAccountDataRelationshipsMarkets)`

SetJwtMarkets sets JwtMarkets field to given value.

### HasJwtMarkets

`func (o *SkuDataRelationships) HasJwtMarkets() bool`

HasJwtMarkets returns a boolean if a field has been set.

### GetJwtStockLocations

`func (o *SkuDataRelationships) GetJwtStockLocations() DeliveryLeadTimeDataRelationshipsStockLocation`

GetJwtStockLocations returns the JwtStockLocations field if non-nil, zero value otherwise.

### GetJwtStockLocationsOk

`func (o *SkuDataRelationships) GetJwtStockLocationsOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool)`

GetJwtStockLocationsOk returns a tuple with the JwtStockLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtStockLocations

`func (o *SkuDataRelationships) SetJwtStockLocations(v DeliveryLeadTimeDataRelationshipsStockLocation)`

SetJwtStockLocations sets JwtStockLocations field to given value.

### HasJwtStockLocations

`func (o *SkuDataRelationships) HasJwtStockLocations() bool`

HasJwtStockLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


