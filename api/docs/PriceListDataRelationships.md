# PriceListDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prices** | Pointer to [**PriceFrequencyTierDataRelationshipsPrice**](PriceFrequencyTierDataRelationshipsPrice.md) |  | [optional] 
**PriceListSchedulers** | Pointer to [**MarketDataRelationshipsPriceListSchedulers**](MarketDataRelationshipsPriceListSchedulers.md) |  | [optional] 
**Attachments** | Pointer to [**AuthorizationDataRelationshipsAttachments**](AuthorizationDataRelationshipsAttachments.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewPriceListDataRelationships

`func NewPriceListDataRelationships() *PriceListDataRelationships`

NewPriceListDataRelationships instantiates a new PriceListDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceListDataRelationshipsWithDefaults

`func NewPriceListDataRelationshipsWithDefaults() *PriceListDataRelationships`

NewPriceListDataRelationshipsWithDefaults instantiates a new PriceListDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrices

`func (o *PriceListDataRelationships) GetPrices() PriceFrequencyTierDataRelationshipsPrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *PriceListDataRelationships) GetPricesOk() (*PriceFrequencyTierDataRelationshipsPrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *PriceListDataRelationships) SetPrices(v PriceFrequencyTierDataRelationshipsPrice)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *PriceListDataRelationships) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetPriceListSchedulers

`func (o *PriceListDataRelationships) GetPriceListSchedulers() MarketDataRelationshipsPriceListSchedulers`

GetPriceListSchedulers returns the PriceListSchedulers field if non-nil, zero value otherwise.

### GetPriceListSchedulersOk

`func (o *PriceListDataRelationships) GetPriceListSchedulersOk() (*MarketDataRelationshipsPriceListSchedulers, bool)`

GetPriceListSchedulersOk returns a tuple with the PriceListSchedulers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListSchedulers

`func (o *PriceListDataRelationships) SetPriceListSchedulers(v MarketDataRelationshipsPriceListSchedulers)`

SetPriceListSchedulers sets PriceListSchedulers field to given value.

### HasPriceListSchedulers

`func (o *PriceListDataRelationships) HasPriceListSchedulers() bool`

HasPriceListSchedulers returns a boolean if a field has been set.

### GetAttachments

`func (o *PriceListDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PriceListDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PriceListDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *PriceListDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetVersions

`func (o *PriceListDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *PriceListDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *PriceListDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *PriceListDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


