# POSTSubscriptionModelsRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The subscription model&#39;s internal name. | 
**Strategy** | Pointer to **interface{}** | The subscription model&#39;s strategy used to generate order subscriptions: one between &#39;by_frequency&#39; (default) and &#39;by_line_items&#39;. | [optional] 
**Frequencies** | **interface{}** | An object that contains the frequencies available for this subscription model. Supported ones are &#39;hourly&#39;, &#39;daily&#39;, &#39;weekly&#39;, &#39;monthly&#39;, &#39;two-month&#39;, &#39;three-month&#39;, &#39;four-month&#39;, &#39;six-month&#39;, &#39;yearly&#39;, or your custom crontab expression (min unit is hour). | 
**AutoActivate** | Pointer to **interface{}** | Indicates if the created subscriptions will be activated considering the placed source order as its first run, default to true. | [optional] 
**AutoCancel** | Pointer to **interface{}** | Indicates if the created subscriptions will be cancelled in case the source order is cancelled, default to false. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTSubscriptionModelsRequestDataAttributes

`func NewPOSTSubscriptionModelsRequestDataAttributes(name interface{}, frequencies interface{}, ) *POSTSubscriptionModelsRequestDataAttributes`

NewPOSTSubscriptionModelsRequestDataAttributes instantiates a new POSTSubscriptionModelsRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTSubscriptionModelsRequestDataAttributesWithDefaults

`func NewPOSTSubscriptionModelsRequestDataAttributesWithDefaults() *POSTSubscriptionModelsRequestDataAttributes`

NewPOSTSubscriptionModelsRequestDataAttributesWithDefaults instantiates a new POSTSubscriptionModelsRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTSubscriptionModelsRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTSubscriptionModelsRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTSubscriptionModelsRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTSubscriptionModelsRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTSubscriptionModelsRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStrategy

`func (o *POSTSubscriptionModelsRequestDataAttributes) GetStrategy() interface{}`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *POSTSubscriptionModelsRequestDataAttributes) GetStrategyOk() (*interface{}, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *POSTSubscriptionModelsRequestDataAttributes) SetStrategy(v interface{})`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *POSTSubscriptionModelsRequestDataAttributes) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### SetStrategyNil

`func (o *POSTSubscriptionModelsRequestDataAttributes) SetStrategyNil(b bool)`

 SetStrategyNil sets the value for Strategy to be an explicit nil

### UnsetStrategy
`func (o *POSTSubscriptionModelsRequestDataAttributes) UnsetStrategy()`

UnsetStrategy ensures that no value is present for Strategy, not even an explicit nil
### GetFrequencies

`func (o *POSTSubscriptionModelsRequestDataAttributes) GetFrequencies() interface{}`

GetFrequencies returns the Frequencies field if non-nil, zero value otherwise.

### GetFrequenciesOk

`func (o *POSTSubscriptionModelsRequestDataAttributes) GetFrequenciesOk() (*interface{}, bool)`

GetFrequenciesOk returns a tuple with the Frequencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencies

`func (o *POSTSubscriptionModelsRequestDataAttributes) SetFrequencies(v interface{})`

SetFrequencies sets Frequencies field to given value.


### SetFrequenciesNil

`func (o *POSTSubscriptionModelsRequestDataAttributes) SetFrequenciesNil(b bool)`

 SetFrequenciesNil sets the value for Frequencies to be an explicit nil

### UnsetFrequencies
`func (o *POSTSubscriptionModelsRequestDataAttributes) UnsetFrequencies()`

UnsetFrequencies ensures that no value is present for Frequencies, not even an explicit nil
### GetAutoActivate

`func (o *POSTSubscriptionModelsRequestDataAttributes) GetAutoActivate() interface{}`

GetAutoActivate returns the AutoActivate field if non-nil, zero value otherwise.

### GetAutoActivateOk

`func (o *POSTSubscriptionModelsRequestDataAttributes) GetAutoActivateOk() (*interface{}, bool)`

GetAutoActivateOk returns a tuple with the AutoActivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoActivate

`func (o *POSTSubscriptionModelsRequestDataAttributes) SetAutoActivate(v interface{})`

SetAutoActivate sets AutoActivate field to given value.

### HasAutoActivate

`func (o *POSTSubscriptionModelsRequestDataAttributes) HasAutoActivate() bool`

HasAutoActivate returns a boolean if a field has been set.

### SetAutoActivateNil

`func (o *POSTSubscriptionModelsRequestDataAttributes) SetAutoActivateNil(b bool)`

 SetAutoActivateNil sets the value for AutoActivate to be an explicit nil

### UnsetAutoActivate
`func (o *POSTSubscriptionModelsRequestDataAttributes) UnsetAutoActivate()`

UnsetAutoActivate ensures that no value is present for AutoActivate, not even an explicit nil
### GetAutoCancel

`func (o *POSTSubscriptionModelsRequestDataAttributes) GetAutoCancel() interface{}`

GetAutoCancel returns the AutoCancel field if non-nil, zero value otherwise.

### GetAutoCancelOk

`func (o *POSTSubscriptionModelsRequestDataAttributes) GetAutoCancelOk() (*interface{}, bool)`

GetAutoCancelOk returns a tuple with the AutoCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCancel

`func (o *POSTSubscriptionModelsRequestDataAttributes) SetAutoCancel(v interface{})`

SetAutoCancel sets AutoCancel field to given value.

### HasAutoCancel

`func (o *POSTSubscriptionModelsRequestDataAttributes) HasAutoCancel() bool`

HasAutoCancel returns a boolean if a field has been set.

### SetAutoCancelNil

`func (o *POSTSubscriptionModelsRequestDataAttributes) SetAutoCancelNil(b bool)`

 SetAutoCancelNil sets the value for AutoCancel to be an explicit nil

### UnsetAutoCancel
`func (o *POSTSubscriptionModelsRequestDataAttributes) UnsetAutoCancel()`

UnsetAutoCancel ensures that no value is present for AutoCancel, not even an explicit nil
### GetReference

`func (o *POSTSubscriptionModelsRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTSubscriptionModelsRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTSubscriptionModelsRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTSubscriptionModelsRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTSubscriptionModelsRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTSubscriptionModelsRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTSubscriptionModelsRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTSubscriptionModelsRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTSubscriptionModelsRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTSubscriptionModelsRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTSubscriptionModelsRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTSubscriptionModelsRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTSubscriptionModelsRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTSubscriptionModelsRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTSubscriptionModelsRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTSubscriptionModelsRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTSubscriptionModelsRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTSubscriptionModelsRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


