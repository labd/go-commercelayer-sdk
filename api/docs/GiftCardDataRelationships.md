# GiftCardDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**AvalaraAccountDataRelationshipsMarkets**](AvalaraAccountDataRelationshipsMarkets.md) |  | [optional] 
**GiftCardRecipient** | Pointer to [**GiftCardDataRelationshipsGiftCardRecipient**](GiftCardDataRelationshipsGiftCardRecipient.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountDataRelationshipsAttachments**](AvalaraAccountDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**CustomerAddressDataRelationshipsEvents**](CustomerAddressDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewGiftCardDataRelationships

`func NewGiftCardDataRelationships() *GiftCardDataRelationships`

NewGiftCardDataRelationships instantiates a new GiftCardDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardDataRelationshipsWithDefaults

`func NewGiftCardDataRelationshipsWithDefaults() *GiftCardDataRelationships`

NewGiftCardDataRelationshipsWithDefaults instantiates a new GiftCardDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *GiftCardDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *GiftCardDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *GiftCardDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *GiftCardDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetGiftCardRecipient

`func (o *GiftCardDataRelationships) GetGiftCardRecipient() GiftCardDataRelationshipsGiftCardRecipient`

GetGiftCardRecipient returns the GiftCardRecipient field if non-nil, zero value otherwise.

### GetGiftCardRecipientOk

`func (o *GiftCardDataRelationships) GetGiftCardRecipientOk() (*GiftCardDataRelationshipsGiftCardRecipient, bool)`

GetGiftCardRecipientOk returns a tuple with the GiftCardRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardRecipient

`func (o *GiftCardDataRelationships) SetGiftCardRecipient(v GiftCardDataRelationshipsGiftCardRecipient)`

SetGiftCardRecipient sets GiftCardRecipient field to given value.

### HasGiftCardRecipient

`func (o *GiftCardDataRelationships) HasGiftCardRecipient() bool`

HasGiftCardRecipient returns a boolean if a field has been set.

### GetAttachments

`func (o *GiftCardDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GiftCardDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GiftCardDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GiftCardDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *GiftCardDataRelationships) GetEvents() CustomerAddressDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GiftCardDataRelationships) GetEventsOk() (*CustomerAddressDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GiftCardDataRelationships) SetEvents(v CustomerAddressDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GiftCardDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


