# PATCHParcelsParcelIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Weight** | Pointer to **interface{}** | The parcel weight, used to automatically calculate the tax rates from the available carrier accounts. | [optional] 
**UnitOfWeight** | Pointer to **interface{}** | Can be one of &#39;gr&#39;, &#39;lb&#39;, or &#39;oz&#39; | [optional] 
**EelPfc** | Pointer to **interface{}** | When shipping outside the US, you need to provide either an Exemption and Exclusion Legend (EEL) code or a Proof of Filing Citation (PFC). Which you need is based on the value of the goods being shipped. Value can be one of \&quot;EEL\&quot; o \&quot;PFC\&quot;. | [optional] 
**ContentsType** | Pointer to **interface{}** | The type of item you are sending. Can be one of &#39;merchandise&#39;, &#39;gift&#39;, &#39;documents&#39;, &#39;returned_goods&#39;, &#39;sample&#39;, or &#39;other&#39;. | [optional] 
**ContentsExplanation** | Pointer to **interface{}** | If you specify &#39;other&#39; in the &#39;contents_type&#39; attribute, you must supply a brief description in this attribute. | [optional] 
**CustomsCertify** | Pointer to **interface{}** | Indicates if the provided information is accurate | [optional] 
**CustomsSigner** | Pointer to **interface{}** | This is the name of the person who is certifying that the information provided on the customs form is accurate. Use a name of the person in your organization who is responsible for this. | [optional] 
**NonDeliveryOption** | Pointer to **interface{}** | In case the shipment cannot be delivered, this option tells the carrier what you want to happen to the parcel. You can pass either &#39;return&#39;, or &#39;abandon&#39;. The value defaults to &#39;return&#39;. If you pass &#39;abandon&#39;, you will not receive the parcel back if it cannot be delivered. | [optional] 
**RestrictionType** | Pointer to **interface{}** | Describes if your parcel requires any special treatment or quarantine when entering the country. Can be one of &#39;none&#39;, &#39;other&#39;, &#39;quarantine&#39;, or &#39;sanitary_phytosanitary_inspection&#39;. | [optional] 
**RestrictionComments** | Pointer to **interface{}** | If you specify &#39;other&#39; in the restriction type, you must supply a brief description of what is required. | [optional] 
**CustomsInfoRequired** | Pointer to **interface{}** | Indicates if the parcel requires customs info to get the shipping rates. | [optional] 
**ShippingLabelUrl** | Pointer to **interface{}** | The shipping label url, ready to be downloaded and printed. | [optional] 
**ShippingLabelFileType** | Pointer to **interface{}** | The shipping label file type. One of &#39;application/pdf&#39;, &#39;application/zpl&#39;, &#39;application/epl2&#39;, or &#39;image/png&#39;. | [optional] 
**ShippingLabelSize** | Pointer to **interface{}** | The shipping label size. | [optional] 
**ShippingLabelResolution** | Pointer to **interface{}** | The shipping label resolution. | [optional] 
**TrackingNumber** | Pointer to **interface{}** | The tracking number associated to this parcel. | [optional] 
**TrackingStatus** | Pointer to **interface{}** | The tracking status for this parcel, automatically updated in real time by the shipping carrier. | [optional] 
**TrackingStatusDetail** | Pointer to **interface{}** | Additional information about the tracking status, automatically updated in real time by the shipping carrier. | [optional] 
**TrackingStatusUpdatedAt** | Pointer to **interface{}** | Time at which the parcel&#39;s tracking status was last updated. | [optional] 
**TrackingDetails** | Pointer to **interface{}** | The parcel&#39;s full tracking history, automatically updated in real time by the shipping carrier. | [optional] 
**CarrierWeightOz** | Pointer to **interface{}** | The weight of the parcel as measured by the carrier in ounces (if available) | [optional] 
**SignedBy** | Pointer to **interface{}** | The name of the person who signed for the parcel (if available) | [optional] 
**Incoterm** | Pointer to **interface{}** | The type of Incoterm (if available). | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHParcelsParcelIdRequestDataAttributes

`func NewPATCHParcelsParcelIdRequestDataAttributes() *PATCHParcelsParcelIdRequestDataAttributes`

NewPATCHParcelsParcelIdRequestDataAttributes instantiates a new PATCHParcelsParcelIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHParcelsParcelIdRequestDataAttributesWithDefaults

`func NewPATCHParcelsParcelIdRequestDataAttributesWithDefaults() *PATCHParcelsParcelIdRequestDataAttributes`

NewPATCHParcelsParcelIdRequestDataAttributesWithDefaults instantiates a new PATCHParcelsParcelIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeight

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetWeight() interface{}`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetWeightOk() (*interface{}, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetWeight(v interface{})`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetUnitOfWeight

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetUnitOfWeight() interface{}`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetUnitOfWeightOk() (*interface{}, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetUnitOfWeight(v interface{})`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### SetUnitOfWeightNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetUnitOfWeightNil(b bool)`

 SetUnitOfWeightNil sets the value for UnitOfWeight to be an explicit nil

### UnsetUnitOfWeight
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetUnitOfWeight()`

UnsetUnitOfWeight ensures that no value is present for UnitOfWeight, not even an explicit nil
### GetEelPfc

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetEelPfc() interface{}`

GetEelPfc returns the EelPfc field if non-nil, zero value otherwise.

### GetEelPfcOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetEelPfcOk() (*interface{}, bool)`

GetEelPfcOk returns a tuple with the EelPfc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEelPfc

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetEelPfc(v interface{})`

SetEelPfc sets EelPfc field to given value.

### HasEelPfc

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasEelPfc() bool`

HasEelPfc returns a boolean if a field has been set.

### SetEelPfcNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetEelPfcNil(b bool)`

 SetEelPfcNil sets the value for EelPfc to be an explicit nil

### UnsetEelPfc
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetEelPfc()`

UnsetEelPfc ensures that no value is present for EelPfc, not even an explicit nil
### GetContentsType

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetContentsType() interface{}`

GetContentsType returns the ContentsType field if non-nil, zero value otherwise.

### GetContentsTypeOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetContentsTypeOk() (*interface{}, bool)`

GetContentsTypeOk returns a tuple with the ContentsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsType

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetContentsType(v interface{})`

SetContentsType sets ContentsType field to given value.

### HasContentsType

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasContentsType() bool`

HasContentsType returns a boolean if a field has been set.

### SetContentsTypeNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetContentsTypeNil(b bool)`

 SetContentsTypeNil sets the value for ContentsType to be an explicit nil

### UnsetContentsType
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetContentsType()`

UnsetContentsType ensures that no value is present for ContentsType, not even an explicit nil
### GetContentsExplanation

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetContentsExplanation() interface{}`

GetContentsExplanation returns the ContentsExplanation field if non-nil, zero value otherwise.

### GetContentsExplanationOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetContentsExplanationOk() (*interface{}, bool)`

GetContentsExplanationOk returns a tuple with the ContentsExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsExplanation

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetContentsExplanation(v interface{})`

SetContentsExplanation sets ContentsExplanation field to given value.

### HasContentsExplanation

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasContentsExplanation() bool`

HasContentsExplanation returns a boolean if a field has been set.

### SetContentsExplanationNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetContentsExplanationNil(b bool)`

 SetContentsExplanationNil sets the value for ContentsExplanation to be an explicit nil

### UnsetContentsExplanation
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetContentsExplanation()`

UnsetContentsExplanation ensures that no value is present for ContentsExplanation, not even an explicit nil
### GetCustomsCertify

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetCustomsCertify() interface{}`

GetCustomsCertify returns the CustomsCertify field if non-nil, zero value otherwise.

### GetCustomsCertifyOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetCustomsCertifyOk() (*interface{}, bool)`

GetCustomsCertifyOk returns a tuple with the CustomsCertify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsCertify

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetCustomsCertify(v interface{})`

SetCustomsCertify sets CustomsCertify field to given value.

### HasCustomsCertify

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasCustomsCertify() bool`

HasCustomsCertify returns a boolean if a field has been set.

### SetCustomsCertifyNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetCustomsCertifyNil(b bool)`

 SetCustomsCertifyNil sets the value for CustomsCertify to be an explicit nil

### UnsetCustomsCertify
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetCustomsCertify()`

UnsetCustomsCertify ensures that no value is present for CustomsCertify, not even an explicit nil
### GetCustomsSigner

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetCustomsSigner() interface{}`

GetCustomsSigner returns the CustomsSigner field if non-nil, zero value otherwise.

### GetCustomsSignerOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetCustomsSignerOk() (*interface{}, bool)`

GetCustomsSignerOk returns a tuple with the CustomsSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsSigner

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetCustomsSigner(v interface{})`

SetCustomsSigner sets CustomsSigner field to given value.

### HasCustomsSigner

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasCustomsSigner() bool`

HasCustomsSigner returns a boolean if a field has been set.

### SetCustomsSignerNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetCustomsSignerNil(b bool)`

 SetCustomsSignerNil sets the value for CustomsSigner to be an explicit nil

### UnsetCustomsSigner
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetCustomsSigner()`

UnsetCustomsSigner ensures that no value is present for CustomsSigner, not even an explicit nil
### GetNonDeliveryOption

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetNonDeliveryOption() interface{}`

GetNonDeliveryOption returns the NonDeliveryOption field if non-nil, zero value otherwise.

### GetNonDeliveryOptionOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetNonDeliveryOptionOk() (*interface{}, bool)`

GetNonDeliveryOptionOk returns a tuple with the NonDeliveryOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDeliveryOption

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetNonDeliveryOption(v interface{})`

SetNonDeliveryOption sets NonDeliveryOption field to given value.

### HasNonDeliveryOption

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasNonDeliveryOption() bool`

HasNonDeliveryOption returns a boolean if a field has been set.

### SetNonDeliveryOptionNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetNonDeliveryOptionNil(b bool)`

 SetNonDeliveryOptionNil sets the value for NonDeliveryOption to be an explicit nil

### UnsetNonDeliveryOption
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetNonDeliveryOption()`

UnsetNonDeliveryOption ensures that no value is present for NonDeliveryOption, not even an explicit nil
### GetRestrictionType

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetRestrictionType() interface{}`

GetRestrictionType returns the RestrictionType field if non-nil, zero value otherwise.

### GetRestrictionTypeOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetRestrictionTypeOk() (*interface{}, bool)`

GetRestrictionTypeOk returns a tuple with the RestrictionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionType

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetRestrictionType(v interface{})`

SetRestrictionType sets RestrictionType field to given value.

### HasRestrictionType

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasRestrictionType() bool`

HasRestrictionType returns a boolean if a field has been set.

### SetRestrictionTypeNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetRestrictionTypeNil(b bool)`

 SetRestrictionTypeNil sets the value for RestrictionType to be an explicit nil

### UnsetRestrictionType
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetRestrictionType()`

UnsetRestrictionType ensures that no value is present for RestrictionType, not even an explicit nil
### GetRestrictionComments

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetRestrictionComments() interface{}`

GetRestrictionComments returns the RestrictionComments field if non-nil, zero value otherwise.

### GetRestrictionCommentsOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetRestrictionCommentsOk() (*interface{}, bool)`

GetRestrictionCommentsOk returns a tuple with the RestrictionComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionComments

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetRestrictionComments(v interface{})`

SetRestrictionComments sets RestrictionComments field to given value.

### HasRestrictionComments

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasRestrictionComments() bool`

HasRestrictionComments returns a boolean if a field has been set.

### SetRestrictionCommentsNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetRestrictionCommentsNil(b bool)`

 SetRestrictionCommentsNil sets the value for RestrictionComments to be an explicit nil

### UnsetRestrictionComments
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetRestrictionComments()`

UnsetRestrictionComments ensures that no value is present for RestrictionComments, not even an explicit nil
### GetCustomsInfoRequired

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetCustomsInfoRequired() interface{}`

GetCustomsInfoRequired returns the CustomsInfoRequired field if non-nil, zero value otherwise.

### GetCustomsInfoRequiredOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetCustomsInfoRequiredOk() (*interface{}, bool)`

GetCustomsInfoRequiredOk returns a tuple with the CustomsInfoRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsInfoRequired

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetCustomsInfoRequired(v interface{})`

SetCustomsInfoRequired sets CustomsInfoRequired field to given value.

### HasCustomsInfoRequired

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasCustomsInfoRequired() bool`

HasCustomsInfoRequired returns a boolean if a field has been set.

### SetCustomsInfoRequiredNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetCustomsInfoRequiredNil(b bool)`

 SetCustomsInfoRequiredNil sets the value for CustomsInfoRequired to be an explicit nil

### UnsetCustomsInfoRequired
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetCustomsInfoRequired()`

UnsetCustomsInfoRequired ensures that no value is present for CustomsInfoRequired, not even an explicit nil
### GetShippingLabelUrl

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetShippingLabelUrl() interface{}`

GetShippingLabelUrl returns the ShippingLabelUrl field if non-nil, zero value otherwise.

### GetShippingLabelUrlOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetShippingLabelUrlOk() (*interface{}, bool)`

GetShippingLabelUrlOk returns a tuple with the ShippingLabelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelUrl

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetShippingLabelUrl(v interface{})`

SetShippingLabelUrl sets ShippingLabelUrl field to given value.

### HasShippingLabelUrl

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasShippingLabelUrl() bool`

HasShippingLabelUrl returns a boolean if a field has been set.

### SetShippingLabelUrlNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetShippingLabelUrlNil(b bool)`

 SetShippingLabelUrlNil sets the value for ShippingLabelUrl to be an explicit nil

### UnsetShippingLabelUrl
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetShippingLabelUrl()`

UnsetShippingLabelUrl ensures that no value is present for ShippingLabelUrl, not even an explicit nil
### GetShippingLabelFileType

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetShippingLabelFileType() interface{}`

GetShippingLabelFileType returns the ShippingLabelFileType field if non-nil, zero value otherwise.

### GetShippingLabelFileTypeOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetShippingLabelFileTypeOk() (*interface{}, bool)`

GetShippingLabelFileTypeOk returns a tuple with the ShippingLabelFileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelFileType

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetShippingLabelFileType(v interface{})`

SetShippingLabelFileType sets ShippingLabelFileType field to given value.

### HasShippingLabelFileType

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasShippingLabelFileType() bool`

HasShippingLabelFileType returns a boolean if a field has been set.

### SetShippingLabelFileTypeNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetShippingLabelFileTypeNil(b bool)`

 SetShippingLabelFileTypeNil sets the value for ShippingLabelFileType to be an explicit nil

### UnsetShippingLabelFileType
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetShippingLabelFileType()`

UnsetShippingLabelFileType ensures that no value is present for ShippingLabelFileType, not even an explicit nil
### GetShippingLabelSize

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetShippingLabelSize() interface{}`

GetShippingLabelSize returns the ShippingLabelSize field if non-nil, zero value otherwise.

### GetShippingLabelSizeOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetShippingLabelSizeOk() (*interface{}, bool)`

GetShippingLabelSizeOk returns a tuple with the ShippingLabelSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelSize

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetShippingLabelSize(v interface{})`

SetShippingLabelSize sets ShippingLabelSize field to given value.

### HasShippingLabelSize

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasShippingLabelSize() bool`

HasShippingLabelSize returns a boolean if a field has been set.

### SetShippingLabelSizeNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetShippingLabelSizeNil(b bool)`

 SetShippingLabelSizeNil sets the value for ShippingLabelSize to be an explicit nil

### UnsetShippingLabelSize
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetShippingLabelSize()`

UnsetShippingLabelSize ensures that no value is present for ShippingLabelSize, not even an explicit nil
### GetShippingLabelResolution

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetShippingLabelResolution() interface{}`

GetShippingLabelResolution returns the ShippingLabelResolution field if non-nil, zero value otherwise.

### GetShippingLabelResolutionOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetShippingLabelResolutionOk() (*interface{}, bool)`

GetShippingLabelResolutionOk returns a tuple with the ShippingLabelResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelResolution

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetShippingLabelResolution(v interface{})`

SetShippingLabelResolution sets ShippingLabelResolution field to given value.

### HasShippingLabelResolution

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasShippingLabelResolution() bool`

HasShippingLabelResolution returns a boolean if a field has been set.

### SetShippingLabelResolutionNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetShippingLabelResolutionNil(b bool)`

 SetShippingLabelResolutionNil sets the value for ShippingLabelResolution to be an explicit nil

### UnsetShippingLabelResolution
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetShippingLabelResolution()`

UnsetShippingLabelResolution ensures that no value is present for ShippingLabelResolution, not even an explicit nil
### GetTrackingNumber

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetTrackingNumber() interface{}`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetTrackingNumberOk() (*interface{}, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetTrackingNumber(v interface{})`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
### GetTrackingStatus

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetTrackingStatus() interface{}`

GetTrackingStatus returns the TrackingStatus field if non-nil, zero value otherwise.

### GetTrackingStatusOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetTrackingStatusOk() (*interface{}, bool)`

GetTrackingStatusOk returns a tuple with the TrackingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatus

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetTrackingStatus(v interface{})`

SetTrackingStatus sets TrackingStatus field to given value.

### HasTrackingStatus

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasTrackingStatus() bool`

HasTrackingStatus returns a boolean if a field has been set.

### SetTrackingStatusNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetTrackingStatusNil(b bool)`

 SetTrackingStatusNil sets the value for TrackingStatus to be an explicit nil

### UnsetTrackingStatus
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetTrackingStatus()`

UnsetTrackingStatus ensures that no value is present for TrackingStatus, not even an explicit nil
### GetTrackingStatusDetail

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetTrackingStatusDetail() interface{}`

GetTrackingStatusDetail returns the TrackingStatusDetail field if non-nil, zero value otherwise.

### GetTrackingStatusDetailOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetTrackingStatusDetailOk() (*interface{}, bool)`

GetTrackingStatusDetailOk returns a tuple with the TrackingStatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatusDetail

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetTrackingStatusDetail(v interface{})`

SetTrackingStatusDetail sets TrackingStatusDetail field to given value.

### HasTrackingStatusDetail

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasTrackingStatusDetail() bool`

HasTrackingStatusDetail returns a boolean if a field has been set.

### SetTrackingStatusDetailNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetTrackingStatusDetailNil(b bool)`

 SetTrackingStatusDetailNil sets the value for TrackingStatusDetail to be an explicit nil

### UnsetTrackingStatusDetail
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetTrackingStatusDetail()`

UnsetTrackingStatusDetail ensures that no value is present for TrackingStatusDetail, not even an explicit nil
### GetTrackingStatusUpdatedAt

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetTrackingStatusUpdatedAt() interface{}`

GetTrackingStatusUpdatedAt returns the TrackingStatusUpdatedAt field if non-nil, zero value otherwise.

### GetTrackingStatusUpdatedAtOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetTrackingStatusUpdatedAtOk() (*interface{}, bool)`

GetTrackingStatusUpdatedAtOk returns a tuple with the TrackingStatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatusUpdatedAt

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetTrackingStatusUpdatedAt(v interface{})`

SetTrackingStatusUpdatedAt sets TrackingStatusUpdatedAt field to given value.

### HasTrackingStatusUpdatedAt

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasTrackingStatusUpdatedAt() bool`

HasTrackingStatusUpdatedAt returns a boolean if a field has been set.

### SetTrackingStatusUpdatedAtNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetTrackingStatusUpdatedAtNil(b bool)`

 SetTrackingStatusUpdatedAtNil sets the value for TrackingStatusUpdatedAt to be an explicit nil

### UnsetTrackingStatusUpdatedAt
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetTrackingStatusUpdatedAt()`

UnsetTrackingStatusUpdatedAt ensures that no value is present for TrackingStatusUpdatedAt, not even an explicit nil
### GetTrackingDetails

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetTrackingDetails() interface{}`

GetTrackingDetails returns the TrackingDetails field if non-nil, zero value otherwise.

### GetTrackingDetailsOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetTrackingDetailsOk() (*interface{}, bool)`

GetTrackingDetailsOk returns a tuple with the TrackingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingDetails

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetTrackingDetails(v interface{})`

SetTrackingDetails sets TrackingDetails field to given value.

### HasTrackingDetails

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasTrackingDetails() bool`

HasTrackingDetails returns a boolean if a field has been set.

### SetTrackingDetailsNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetTrackingDetailsNil(b bool)`

 SetTrackingDetailsNil sets the value for TrackingDetails to be an explicit nil

### UnsetTrackingDetails
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetTrackingDetails()`

UnsetTrackingDetails ensures that no value is present for TrackingDetails, not even an explicit nil
### GetCarrierWeightOz

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetCarrierWeightOz() interface{}`

GetCarrierWeightOz returns the CarrierWeightOz field if non-nil, zero value otherwise.

### GetCarrierWeightOzOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetCarrierWeightOzOk() (*interface{}, bool)`

GetCarrierWeightOzOk returns a tuple with the CarrierWeightOz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierWeightOz

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetCarrierWeightOz(v interface{})`

SetCarrierWeightOz sets CarrierWeightOz field to given value.

### HasCarrierWeightOz

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasCarrierWeightOz() bool`

HasCarrierWeightOz returns a boolean if a field has been set.

### SetCarrierWeightOzNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetCarrierWeightOzNil(b bool)`

 SetCarrierWeightOzNil sets the value for CarrierWeightOz to be an explicit nil

### UnsetCarrierWeightOz
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetCarrierWeightOz()`

UnsetCarrierWeightOz ensures that no value is present for CarrierWeightOz, not even an explicit nil
### GetSignedBy

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetSignedBy() interface{}`

GetSignedBy returns the SignedBy field if non-nil, zero value otherwise.

### GetSignedByOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetSignedByOk() (*interface{}, bool)`

GetSignedByOk returns a tuple with the SignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedBy

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetSignedBy(v interface{})`

SetSignedBy sets SignedBy field to given value.

### HasSignedBy

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasSignedBy() bool`

HasSignedBy returns a boolean if a field has been set.

### SetSignedByNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetSignedByNil(b bool)`

 SetSignedByNil sets the value for SignedBy to be an explicit nil

### UnsetSignedBy
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetSignedBy()`

UnsetSignedBy ensures that no value is present for SignedBy, not even an explicit nil
### GetIncoterm

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetIncoterm() interface{}`

GetIncoterm returns the Incoterm field if non-nil, zero value otherwise.

### GetIncotermOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetIncotermOk() (*interface{}, bool)`

GetIncotermOk returns a tuple with the Incoterm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncoterm

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetIncoterm(v interface{})`

SetIncoterm sets Incoterm field to given value.

### HasIncoterm

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasIncoterm() bool`

HasIncoterm returns a boolean if a field has been set.

### SetIncotermNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetIncotermNil(b bool)`

 SetIncotermNil sets the value for Incoterm to be an explicit nil

### UnsetIncoterm
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetIncoterm()`

UnsetIncoterm ensures that no value is present for Incoterm, not even an explicit nil
### GetReference

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHParcelsParcelIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHParcelsParcelIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHParcelsParcelIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHParcelsParcelIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


