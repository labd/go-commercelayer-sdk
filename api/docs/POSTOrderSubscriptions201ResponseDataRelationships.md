# POSTOrderSubscriptions201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets**](GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets.md) |  | [optional] 
**SourceOrder** | [**GETAdyenPayments200ResponseDataInnerRelationshipsOrder**](GETAdyenPayments200ResponseDataInnerRelationshipsOrder.md) |  | 

## Methods

### NewPOSTOrderSubscriptions201ResponseDataRelationships

`func NewPOSTOrderSubscriptions201ResponseDataRelationships(sourceOrder GETAdyenPayments200ResponseDataInnerRelationshipsOrder, ) *POSTOrderSubscriptions201ResponseDataRelationships`

NewPOSTOrderSubscriptions201ResponseDataRelationships instantiates a new POSTOrderSubscriptions201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrderSubscriptions201ResponseDataRelationshipsWithDefaults

`func NewPOSTOrderSubscriptions201ResponseDataRelationshipsWithDefaults() *POSTOrderSubscriptions201ResponseDataRelationships`

NewPOSTOrderSubscriptions201ResponseDataRelationshipsWithDefaults instantiates a new POSTOrderSubscriptions201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetMarket() GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetMarketOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetMarket(v GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetSourceOrder

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetSourceOrder() GETAdyenPayments200ResponseDataInnerRelationshipsOrder`

GetSourceOrder returns the SourceOrder field if non-nil, zero value otherwise.

### GetSourceOrderOk

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetSourceOrderOk() (*GETAdyenPayments200ResponseDataInnerRelationshipsOrder, bool)`

GetSourceOrderOk returns a tuple with the SourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOrder

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetSourceOrder(v GETAdyenPayments200ResponseDataInnerRelationshipsOrder)`

SetSourceOrder sets SourceOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


