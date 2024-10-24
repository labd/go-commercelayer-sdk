# PATCHMarketsMarketId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The market&#39;s internal name. | [optional] 
**Code** | Pointer to **interface{}** | A string that you can use to identify the market (must be unique within the environment). | [optional] 
**FacebookPixelId** | Pointer to **interface{}** | The Facebook Pixed ID. | [optional] 
**CheckoutUrl** | Pointer to **interface{}** | The checkout URL for this market. | [optional] 
**ExternalPricesUrl** | Pointer to **interface{}** | The URL used to overwrite prices by an external source. | [optional] 
**ExternalOrderValidationUrl** | Pointer to **interface{}** | The URL used to validate orders by an external source. | [optional] 
**ShippingCostCutoff** | Pointer to **interface{}** | When specified indicates the maximum number of shipping line items with cost that will be added to an order. | [optional] 
**Disable** | Pointer to **interface{}** | Send this attribute if you want to mark this resource as disabled. | [optional] 
**Enable** | Pointer to **interface{}** | Send this attribute if you want to mark this resource as enabled. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHMarketsMarketId200ResponseDataAttributes

`func NewPATCHMarketsMarketId200ResponseDataAttributes() *PATCHMarketsMarketId200ResponseDataAttributes`

NewPATCHMarketsMarketId200ResponseDataAttributes instantiates a new PATCHMarketsMarketId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHMarketsMarketId200ResponseDataAttributesWithDefaults

`func NewPATCHMarketsMarketId200ResponseDataAttributesWithDefaults() *PATCHMarketsMarketId200ResponseDataAttributes`

NewPATCHMarketsMarketId200ResponseDataAttributesWithDefaults instantiates a new PATCHMarketsMarketId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHMarketsMarketId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCode

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *PATCHMarketsMarketId200ResponseDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetFacebookPixelId

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetFacebookPixelId() interface{}`

GetFacebookPixelId returns the FacebookPixelId field if non-nil, zero value otherwise.

### GetFacebookPixelIdOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetFacebookPixelIdOk() (*interface{}, bool)`

GetFacebookPixelIdOk returns a tuple with the FacebookPixelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookPixelId

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetFacebookPixelId(v interface{})`

SetFacebookPixelId sets FacebookPixelId field to given value.

### HasFacebookPixelId

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasFacebookPixelId() bool`

HasFacebookPixelId returns a boolean if a field has been set.

### SetFacebookPixelIdNil

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetFacebookPixelIdNil(b bool)`

 SetFacebookPixelIdNil sets the value for FacebookPixelId to be an explicit nil

### UnsetFacebookPixelId
`func (o *PATCHMarketsMarketId200ResponseDataAttributes) UnsetFacebookPixelId()`

UnsetFacebookPixelId ensures that no value is present for FacebookPixelId, not even an explicit nil
### GetCheckoutUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetCheckoutUrl() interface{}`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetCheckoutUrlOk() (*interface{}, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetCheckoutUrl(v interface{})`

SetCheckoutUrl sets CheckoutUrl field to given value.

### HasCheckoutUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasCheckoutUrl() bool`

HasCheckoutUrl returns a boolean if a field has been set.

### SetCheckoutUrlNil

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetCheckoutUrlNil(b bool)`

 SetCheckoutUrlNil sets the value for CheckoutUrl to be an explicit nil

### UnsetCheckoutUrl
`func (o *PATCHMarketsMarketId200ResponseDataAttributes) UnsetCheckoutUrl()`

UnsetCheckoutUrl ensures that no value is present for CheckoutUrl, not even an explicit nil
### GetExternalPricesUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetExternalPricesUrl() interface{}`

GetExternalPricesUrl returns the ExternalPricesUrl field if non-nil, zero value otherwise.

### GetExternalPricesUrlOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetExternalPricesUrlOk() (*interface{}, bool)`

GetExternalPricesUrlOk returns a tuple with the ExternalPricesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPricesUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetExternalPricesUrl(v interface{})`

SetExternalPricesUrl sets ExternalPricesUrl field to given value.

### HasExternalPricesUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasExternalPricesUrl() bool`

HasExternalPricesUrl returns a boolean if a field has been set.

### SetExternalPricesUrlNil

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetExternalPricesUrlNil(b bool)`

 SetExternalPricesUrlNil sets the value for ExternalPricesUrl to be an explicit nil

### UnsetExternalPricesUrl
`func (o *PATCHMarketsMarketId200ResponseDataAttributes) UnsetExternalPricesUrl()`

UnsetExternalPricesUrl ensures that no value is present for ExternalPricesUrl, not even an explicit nil
### GetExternalOrderValidationUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetExternalOrderValidationUrl() interface{}`

GetExternalOrderValidationUrl returns the ExternalOrderValidationUrl field if non-nil, zero value otherwise.

### GetExternalOrderValidationUrlOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetExternalOrderValidationUrlOk() (*interface{}, bool)`

GetExternalOrderValidationUrlOk returns a tuple with the ExternalOrderValidationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOrderValidationUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetExternalOrderValidationUrl(v interface{})`

SetExternalOrderValidationUrl sets ExternalOrderValidationUrl field to given value.

### HasExternalOrderValidationUrl

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasExternalOrderValidationUrl() bool`

HasExternalOrderValidationUrl returns a boolean if a field has been set.

### SetExternalOrderValidationUrlNil

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetExternalOrderValidationUrlNil(b bool)`

 SetExternalOrderValidationUrlNil sets the value for ExternalOrderValidationUrl to be an explicit nil

### UnsetExternalOrderValidationUrl
`func (o *PATCHMarketsMarketId200ResponseDataAttributes) UnsetExternalOrderValidationUrl()`

UnsetExternalOrderValidationUrl ensures that no value is present for ExternalOrderValidationUrl, not even an explicit nil
### GetShippingCostCutoff

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetShippingCostCutoff() interface{}`

GetShippingCostCutoff returns the ShippingCostCutoff field if non-nil, zero value otherwise.

### GetShippingCostCutoffOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetShippingCostCutoffOk() (*interface{}, bool)`

GetShippingCostCutoffOk returns a tuple with the ShippingCostCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCostCutoff

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetShippingCostCutoff(v interface{})`

SetShippingCostCutoff sets ShippingCostCutoff field to given value.

### HasShippingCostCutoff

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasShippingCostCutoff() bool`

HasShippingCostCutoff returns a boolean if a field has been set.

### SetShippingCostCutoffNil

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetShippingCostCutoffNil(b bool)`

 SetShippingCostCutoffNil sets the value for ShippingCostCutoff to be an explicit nil

### UnsetShippingCostCutoff
`func (o *PATCHMarketsMarketId200ResponseDataAttributes) UnsetShippingCostCutoff()`

UnsetShippingCostCutoff ensures that no value is present for ShippingCostCutoff, not even an explicit nil
### GetDisable

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetDisable() interface{}`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetDisableOk() (*interface{}, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetDisable(v interface{})`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### SetDisableNil

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetDisableNil(b bool)`

 SetDisableNil sets the value for Disable to be an explicit nil

### UnsetDisable
`func (o *PATCHMarketsMarketId200ResponseDataAttributes) UnsetDisable()`

UnsetDisable ensures that no value is present for Disable, not even an explicit nil
### GetEnable

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetEnable() interface{}`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetEnableOk() (*interface{}, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetEnable(v interface{})`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *PATCHMarketsMarketId200ResponseDataAttributes) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetReference

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHMarketsMarketId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHMarketsMarketId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHMarketsMarketId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHMarketsMarketId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


