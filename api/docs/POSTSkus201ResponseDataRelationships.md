# POSTSkus201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingCategory** | Pointer to [**POSTShipments201ResponseDataRelationshipsShippingCategory**](POSTShipments201ResponseDataRelationshipsShippingCategory.md) |  | [optional] 
**Prices** | Pointer to [**POSTPriceLists201ResponseDataRelationshipsPrices**](POSTPriceLists201ResponseDataRelationshipsPrices.md) |  | [optional] 
**StockItems** | Pointer to [**POSTSkus201ResponseDataRelationshipsStockItems**](POSTSkus201ResponseDataRelationshipsStockItems.md) |  | [optional] 
**StockReservations** | Pointer to [**POSTLineItems201ResponseDataRelationshipsStockReservations**](POSTLineItems201ResponseDataRelationshipsStockReservations.md) |  | [optional] 
**DeliveryLeadTimes** | Pointer to [**POSTSkus201ResponseDataRelationshipsDeliveryLeadTimes**](POSTSkus201ResponseDataRelationshipsDeliveryLeadTimes.md) |  | [optional] 
**SkuOptions** | Pointer to [**POSTSkus201ResponseDataRelationshipsSkuOptions**](POSTSkus201ResponseDataRelationshipsSkuOptions.md) |  | [optional] 
**SkuListItems** | Pointer to [**POSTSkuLists201ResponseDataRelationshipsSkuListItems**](POSTSkuLists201ResponseDataRelationshipsSkuListItems.md) |  | [optional] 
**SkuLists** | Pointer to [**POSTCustomers201ResponseDataRelationshipsSkuLists**](POSTCustomers201ResponseDataRelationshipsSkuLists.md) |  | [optional] 
**Attachments** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Links** | Pointer to [**POSTOrders201ResponseDataRelationshipsLinks**](POSTOrders201ResponseDataRelationshipsLinks.md) |  | [optional] 
**Events** | Pointer to [**POSTAddresses201ResponseDataRelationshipsEvents**](POSTAddresses201ResponseDataRelationshipsEvents.md) |  | [optional] 
**Tags** | Pointer to [**POSTAddresses201ResponseDataRelationshipsTags**](POSTAddresses201ResponseDataRelationshipsTags.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 
**JwtCustomer** | Pointer to [**POSTPrices201ResponseDataRelationshipsJwtCustomer**](POSTPrices201ResponseDataRelationshipsJwtCustomer.md) |  | [optional] 
**JwtMarkets** | Pointer to [**POSTPrices201ResponseDataRelationshipsJwtMarkets**](POSTPrices201ResponseDataRelationshipsJwtMarkets.md) |  | [optional] 
**JwtStockLocations** | Pointer to [**POSTPrices201ResponseDataRelationshipsJwtStockLocations**](POSTPrices201ResponseDataRelationshipsJwtStockLocations.md) |  | [optional] 

## Methods

### NewPOSTSkus201ResponseDataRelationships

`func NewPOSTSkus201ResponseDataRelationships() *POSTSkus201ResponseDataRelationships`

NewPOSTSkus201ResponseDataRelationships instantiates a new POSTSkus201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTSkus201ResponseDataRelationshipsWithDefaults

`func NewPOSTSkus201ResponseDataRelationshipsWithDefaults() *POSTSkus201ResponseDataRelationships`

NewPOSTSkus201ResponseDataRelationshipsWithDefaults instantiates a new POSTSkus201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingCategory

`func (o *POSTSkus201ResponseDataRelationships) GetShippingCategory() POSTShipments201ResponseDataRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *POSTSkus201ResponseDataRelationships) GetShippingCategoryOk() (*POSTShipments201ResponseDataRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *POSTSkus201ResponseDataRelationships) SetShippingCategory(v POSTShipments201ResponseDataRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *POSTSkus201ResponseDataRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetPrices

`func (o *POSTSkus201ResponseDataRelationships) GetPrices() POSTPriceLists201ResponseDataRelationshipsPrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *POSTSkus201ResponseDataRelationships) GetPricesOk() (*POSTPriceLists201ResponseDataRelationshipsPrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *POSTSkus201ResponseDataRelationships) SetPrices(v POSTPriceLists201ResponseDataRelationshipsPrices)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *POSTSkus201ResponseDataRelationships) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetStockItems

`func (o *POSTSkus201ResponseDataRelationships) GetStockItems() POSTSkus201ResponseDataRelationshipsStockItems`

GetStockItems returns the StockItems field if non-nil, zero value otherwise.

### GetStockItemsOk

`func (o *POSTSkus201ResponseDataRelationships) GetStockItemsOk() (*POSTSkus201ResponseDataRelationshipsStockItems, bool)`

GetStockItemsOk returns a tuple with the StockItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockItems

`func (o *POSTSkus201ResponseDataRelationships) SetStockItems(v POSTSkus201ResponseDataRelationshipsStockItems)`

SetStockItems sets StockItems field to given value.

### HasStockItems

`func (o *POSTSkus201ResponseDataRelationships) HasStockItems() bool`

HasStockItems returns a boolean if a field has been set.

### GetStockReservations

`func (o *POSTSkus201ResponseDataRelationships) GetStockReservations() POSTLineItems201ResponseDataRelationshipsStockReservations`

GetStockReservations returns the StockReservations field if non-nil, zero value otherwise.

### GetStockReservationsOk

`func (o *POSTSkus201ResponseDataRelationships) GetStockReservationsOk() (*POSTLineItems201ResponseDataRelationshipsStockReservations, bool)`

GetStockReservationsOk returns a tuple with the StockReservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockReservations

`func (o *POSTSkus201ResponseDataRelationships) SetStockReservations(v POSTLineItems201ResponseDataRelationshipsStockReservations)`

SetStockReservations sets StockReservations field to given value.

### HasStockReservations

`func (o *POSTSkus201ResponseDataRelationships) HasStockReservations() bool`

HasStockReservations returns a boolean if a field has been set.

### GetDeliveryLeadTimes

`func (o *POSTSkus201ResponseDataRelationships) GetDeliveryLeadTimes() POSTSkus201ResponseDataRelationshipsDeliveryLeadTimes`

GetDeliveryLeadTimes returns the DeliveryLeadTimes field if non-nil, zero value otherwise.

### GetDeliveryLeadTimesOk

`func (o *POSTSkus201ResponseDataRelationships) GetDeliveryLeadTimesOk() (*POSTSkus201ResponseDataRelationshipsDeliveryLeadTimes, bool)`

GetDeliveryLeadTimesOk returns a tuple with the DeliveryLeadTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLeadTimes

`func (o *POSTSkus201ResponseDataRelationships) SetDeliveryLeadTimes(v POSTSkus201ResponseDataRelationshipsDeliveryLeadTimes)`

SetDeliveryLeadTimes sets DeliveryLeadTimes field to given value.

### HasDeliveryLeadTimes

`func (o *POSTSkus201ResponseDataRelationships) HasDeliveryLeadTimes() bool`

HasDeliveryLeadTimes returns a boolean if a field has been set.

### GetSkuOptions

`func (o *POSTSkus201ResponseDataRelationships) GetSkuOptions() POSTSkus201ResponseDataRelationshipsSkuOptions`

GetSkuOptions returns the SkuOptions field if non-nil, zero value otherwise.

### GetSkuOptionsOk

`func (o *POSTSkus201ResponseDataRelationships) GetSkuOptionsOk() (*POSTSkus201ResponseDataRelationshipsSkuOptions, bool)`

GetSkuOptionsOk returns a tuple with the SkuOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuOptions

`func (o *POSTSkus201ResponseDataRelationships) SetSkuOptions(v POSTSkus201ResponseDataRelationshipsSkuOptions)`

SetSkuOptions sets SkuOptions field to given value.

### HasSkuOptions

`func (o *POSTSkus201ResponseDataRelationships) HasSkuOptions() bool`

HasSkuOptions returns a boolean if a field has been set.

### GetSkuListItems

`func (o *POSTSkus201ResponseDataRelationships) GetSkuListItems() POSTSkuLists201ResponseDataRelationshipsSkuListItems`

GetSkuListItems returns the SkuListItems field if non-nil, zero value otherwise.

### GetSkuListItemsOk

`func (o *POSTSkus201ResponseDataRelationships) GetSkuListItemsOk() (*POSTSkuLists201ResponseDataRelationshipsSkuListItems, bool)`

GetSkuListItemsOk returns a tuple with the SkuListItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuListItems

`func (o *POSTSkus201ResponseDataRelationships) SetSkuListItems(v POSTSkuLists201ResponseDataRelationshipsSkuListItems)`

SetSkuListItems sets SkuListItems field to given value.

### HasSkuListItems

`func (o *POSTSkus201ResponseDataRelationships) HasSkuListItems() bool`

HasSkuListItems returns a boolean if a field has been set.

### GetSkuLists

`func (o *POSTSkus201ResponseDataRelationships) GetSkuLists() POSTCustomers201ResponseDataRelationshipsSkuLists`

GetSkuLists returns the SkuLists field if non-nil, zero value otherwise.

### GetSkuListsOk

`func (o *POSTSkus201ResponseDataRelationships) GetSkuListsOk() (*POSTCustomers201ResponseDataRelationshipsSkuLists, bool)`

GetSkuListsOk returns a tuple with the SkuLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuLists

`func (o *POSTSkus201ResponseDataRelationships) SetSkuLists(v POSTCustomers201ResponseDataRelationshipsSkuLists)`

SetSkuLists sets SkuLists field to given value.

### HasSkuLists

`func (o *POSTSkus201ResponseDataRelationships) HasSkuLists() bool`

HasSkuLists returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTSkus201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTSkus201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTSkus201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTSkus201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetLinks

`func (o *POSTSkus201ResponseDataRelationships) GetLinks() POSTOrders201ResponseDataRelationshipsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTSkus201ResponseDataRelationships) GetLinksOk() (*POSTOrders201ResponseDataRelationshipsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTSkus201ResponseDataRelationships) SetLinks(v POSTOrders201ResponseDataRelationshipsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTSkus201ResponseDataRelationships) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEvents

`func (o *POSTSkus201ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *POSTSkus201ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *POSTSkus201ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *POSTSkus201ResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTags

`func (o *POSTSkus201ResponseDataRelationships) GetTags() POSTAddresses201ResponseDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *POSTSkus201ResponseDataRelationships) GetTagsOk() (*POSTAddresses201ResponseDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *POSTSkus201ResponseDataRelationships) SetTags(v POSTAddresses201ResponseDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *POSTSkus201ResponseDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersions

`func (o *POSTSkus201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *POSTSkus201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *POSTSkus201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *POSTSkus201ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetJwtCustomer

`func (o *POSTSkus201ResponseDataRelationships) GetJwtCustomer() POSTPrices201ResponseDataRelationshipsJwtCustomer`

GetJwtCustomer returns the JwtCustomer field if non-nil, zero value otherwise.

### GetJwtCustomerOk

`func (o *POSTSkus201ResponseDataRelationships) GetJwtCustomerOk() (*POSTPrices201ResponseDataRelationshipsJwtCustomer, bool)`

GetJwtCustomerOk returns a tuple with the JwtCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtCustomer

`func (o *POSTSkus201ResponseDataRelationships) SetJwtCustomer(v POSTPrices201ResponseDataRelationshipsJwtCustomer)`

SetJwtCustomer sets JwtCustomer field to given value.

### HasJwtCustomer

`func (o *POSTSkus201ResponseDataRelationships) HasJwtCustomer() bool`

HasJwtCustomer returns a boolean if a field has been set.

### GetJwtMarkets

`func (o *POSTSkus201ResponseDataRelationships) GetJwtMarkets() POSTPrices201ResponseDataRelationshipsJwtMarkets`

GetJwtMarkets returns the JwtMarkets field if non-nil, zero value otherwise.

### GetJwtMarketsOk

`func (o *POSTSkus201ResponseDataRelationships) GetJwtMarketsOk() (*POSTPrices201ResponseDataRelationshipsJwtMarkets, bool)`

GetJwtMarketsOk returns a tuple with the JwtMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtMarkets

`func (o *POSTSkus201ResponseDataRelationships) SetJwtMarkets(v POSTPrices201ResponseDataRelationshipsJwtMarkets)`

SetJwtMarkets sets JwtMarkets field to given value.

### HasJwtMarkets

`func (o *POSTSkus201ResponseDataRelationships) HasJwtMarkets() bool`

HasJwtMarkets returns a boolean if a field has been set.

### GetJwtStockLocations

`func (o *POSTSkus201ResponseDataRelationships) GetJwtStockLocations() POSTPrices201ResponseDataRelationshipsJwtStockLocations`

GetJwtStockLocations returns the JwtStockLocations field if non-nil, zero value otherwise.

### GetJwtStockLocationsOk

`func (o *POSTSkus201ResponseDataRelationships) GetJwtStockLocationsOk() (*POSTPrices201ResponseDataRelationshipsJwtStockLocations, bool)`

GetJwtStockLocationsOk returns a tuple with the JwtStockLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtStockLocations

`func (o *POSTSkus201ResponseDataRelationships) SetJwtStockLocations(v POSTPrices201ResponseDataRelationshipsJwtStockLocations)`

SetJwtStockLocations sets JwtStockLocations field to given value.

### HasJwtStockLocations

`func (o *POSTSkus201ResponseDataRelationships) HasJwtStockLocations() bool`

HasJwtStockLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


