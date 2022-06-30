# EventDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastEventCallbacks** | Pointer to [**EventDataRelationshipsLastEventCallbacks**](EventDataRelationshipsLastEventCallbacks.md) |  | [optional] 
**Webhooks** | Pointer to [**EventCallbackDataRelationshipsWebhook**](EventCallbackDataRelationshipsWebhook.md) |  | [optional] 

## Methods

### NewEventDataRelationships

`func NewEventDataRelationships() *EventDataRelationships`

NewEventDataRelationships instantiates a new EventDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDataRelationshipsWithDefaults

`func NewEventDataRelationshipsWithDefaults() *EventDataRelationships`

NewEventDataRelationshipsWithDefaults instantiates a new EventDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastEventCallbacks

`func (o *EventDataRelationships) GetLastEventCallbacks() EventDataRelationshipsLastEventCallbacks`

GetLastEventCallbacks returns the LastEventCallbacks field if non-nil, zero value otherwise.

### GetLastEventCallbacksOk

`func (o *EventDataRelationships) GetLastEventCallbacksOk() (*EventDataRelationshipsLastEventCallbacks, bool)`

GetLastEventCallbacksOk returns a tuple with the LastEventCallbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventCallbacks

`func (o *EventDataRelationships) SetLastEventCallbacks(v EventDataRelationshipsLastEventCallbacks)`

SetLastEventCallbacks sets LastEventCallbacks field to given value.

### HasLastEventCallbacks

`func (o *EventDataRelationships) HasLastEventCallbacks() bool`

HasLastEventCallbacks returns a boolean if a field has been set.

### GetWebhooks

`func (o *EventDataRelationships) GetWebhooks() EventCallbackDataRelationshipsWebhook`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *EventDataRelationships) GetWebhooksOk() (*EventCallbackDataRelationshipsWebhook, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *EventDataRelationships) SetWebhooks(v EventCallbackDataRelationshipsWebhook)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *EventDataRelationships) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


