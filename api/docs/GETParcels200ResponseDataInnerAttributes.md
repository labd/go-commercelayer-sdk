# GETParcels200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** | Unique identifier for the parcel | [optional] 
**Weight** | Pointer to **float32** | The parcel weight, used to automatically calculate the tax rates from the available carrier accounts. | [optional] 
**UnitOfWeight** | Pointer to **string** | Can be one of &#39;gr&#39;, &#39;lb&#39;, or &#39;oz&#39; | [optional] 
**EelPfc** | Pointer to **string** | When shipping outside the US, you need to provide either an Exemption and Exclusion Legend (EEL) code or a Proof of Filing Citation (PFC). Which you need is based on the value of the goods being shipped. Value can be one of \&quot;EEL\&quot; o \&quot;PFC\&quot;. | [optional] 
**ContentsType** | Pointer to **string** | The type of item you are sending. Can be one of &#39;merchandise&#39;, &#39;gift&#39;, &#39;documents&#39;, &#39;returned_goods&#39;, &#39;sample&#39;, or &#39;other&#39;. | [optional] 
**ContentsExplanation** | Pointer to **string** | If you specify &#39;other&#39; in the &#39;contents_type&#39; attribute, you must supply a brief description in this attribute. | [optional] 
**CustomsCertify** | Pointer to **bool** | Indicates if the provided information is accurate | [optional] 
**CustomsSigner** | Pointer to **string** | This is the name of the person who is certifying that the information provided on the customs form is accurate. Use a name of the person in your organization who is responsible for this. | [optional] 
**NonDeliveryOption** | Pointer to **string** | In case the shipment cannot be delivered, this option tells the carrier what you want to happen to the parcel. You can pass either &#39;return&#39;, or &#39;abandon&#39;. The value defaults to &#39;return&#39;. If you pass &#39;abandon&#39;, you will not receive the parcel back if it cannot be delivered. | [optional] 
**RestrictionType** | Pointer to **string** | Describes if your parcel requires any special treatment or quarantine when entering the country. Can be one of &#39;none&#39;, &#39;other&#39;, &#39;quarantine&#39;, or &#39;sanitary_phytosanitary_inspection&#39;. | [optional] 
**RestrictionComments** | Pointer to **string** | If you specify &#39;other&#39; in the restriction type, you must supply a brief description of what is required. | [optional] 
**CustomsInfoRequired** | Pointer to **bool** | Indicates if the parcel requires customs info to get the shipping rates. | [optional] 
**ShippingLabelUrl** | Pointer to **string** | The shipping label url, ready to be downloaded and printed. | [optional] 
**ShippingLabelFileType** | Pointer to **string** | The shipping label file type. One of &#39;application/pdf&#39;, &#39;application/zpl&#39;, &#39;application/epl2&#39;, or &#39;image/png&#39;. | [optional] 
**ShippingLabelSize** | Pointer to **string** | The shipping label size. | [optional] 
**ShippingLabelResolution** | Pointer to **string** | The shipping label resolution. | [optional] 
**TrackingNumber** | Pointer to **string** | The tracking number associated to this parcel. | [optional] 
**TrackingStatus** | Pointer to **string** | The tracking status for this parcel, automatically updated in real time by the shipping carrier. | [optional] 
**TrackingStatusDetail** | Pointer to **string** | Additional information about the tracking status, automatically updated in real time by the shipping carrier. | [optional] 
**TrackingStatusUpdatedAt** | Pointer to **string** | Time at which the parcel&#39;s tracking status was last updated. | [optional] 
**TrackingDetails** | Pointer to **string** | The parcel&#39;s full tracking history, automatically updated in real time by the shipping carrier. | [optional] 
**CarrierWeightOz** | Pointer to **string** | The weight of the parcel as measured by the carrier in ounces (if available) | [optional] 
**SignedBy** | Pointer to **string** | The name of the person who signed for the parcel (if available) | [optional] 
**Incoterm** | Pointer to **string** | The type of Incoterm (if available). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETParcels200ResponseDataInnerAttributes

`func NewGETParcels200ResponseDataInnerAttributes() *GETParcels200ResponseDataInnerAttributes`

NewGETParcels200ResponseDataInnerAttributes instantiates a new GETParcels200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETParcels200ResponseDataInnerAttributesWithDefaults

`func NewGETParcels200ResponseDataInnerAttributesWithDefaults() *GETParcels200ResponseDataInnerAttributes`

NewGETParcels200ResponseDataInnerAttributesWithDefaults instantiates a new GETParcels200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETParcels200ResponseDataInnerAttributes) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETParcels200ResponseDataInnerAttributes) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETParcels200ResponseDataInnerAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetWeight

`func (o *GETParcels200ResponseDataInnerAttributes) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GETParcels200ResponseDataInnerAttributes) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *GETParcels200ResponseDataInnerAttributes) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetUnitOfWeight

`func (o *GETParcels200ResponseDataInnerAttributes) GetUnitOfWeight() string`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetUnitOfWeightOk() (*string, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *GETParcels200ResponseDataInnerAttributes) SetUnitOfWeight(v string)`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *GETParcels200ResponseDataInnerAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### GetEelPfc

`func (o *GETParcels200ResponseDataInnerAttributes) GetEelPfc() string`

GetEelPfc returns the EelPfc field if non-nil, zero value otherwise.

### GetEelPfcOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetEelPfcOk() (*string, bool)`

GetEelPfcOk returns a tuple with the EelPfc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEelPfc

`func (o *GETParcels200ResponseDataInnerAttributes) SetEelPfc(v string)`

SetEelPfc sets EelPfc field to given value.

### HasEelPfc

`func (o *GETParcels200ResponseDataInnerAttributes) HasEelPfc() bool`

HasEelPfc returns a boolean if a field has been set.

### GetContentsType

`func (o *GETParcels200ResponseDataInnerAttributes) GetContentsType() string`

GetContentsType returns the ContentsType field if non-nil, zero value otherwise.

### GetContentsTypeOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetContentsTypeOk() (*string, bool)`

GetContentsTypeOk returns a tuple with the ContentsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsType

`func (o *GETParcels200ResponseDataInnerAttributes) SetContentsType(v string)`

SetContentsType sets ContentsType field to given value.

### HasContentsType

`func (o *GETParcels200ResponseDataInnerAttributes) HasContentsType() bool`

HasContentsType returns a boolean if a field has been set.

### GetContentsExplanation

`func (o *GETParcels200ResponseDataInnerAttributes) GetContentsExplanation() string`

GetContentsExplanation returns the ContentsExplanation field if non-nil, zero value otherwise.

### GetContentsExplanationOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetContentsExplanationOk() (*string, bool)`

GetContentsExplanationOk returns a tuple with the ContentsExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsExplanation

`func (o *GETParcels200ResponseDataInnerAttributes) SetContentsExplanation(v string)`

SetContentsExplanation sets ContentsExplanation field to given value.

### HasContentsExplanation

`func (o *GETParcels200ResponseDataInnerAttributes) HasContentsExplanation() bool`

HasContentsExplanation returns a boolean if a field has been set.

### GetCustomsCertify

`func (o *GETParcels200ResponseDataInnerAttributes) GetCustomsCertify() bool`

GetCustomsCertify returns the CustomsCertify field if non-nil, zero value otherwise.

### GetCustomsCertifyOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetCustomsCertifyOk() (*bool, bool)`

GetCustomsCertifyOk returns a tuple with the CustomsCertify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsCertify

`func (o *GETParcels200ResponseDataInnerAttributes) SetCustomsCertify(v bool)`

SetCustomsCertify sets CustomsCertify field to given value.

### HasCustomsCertify

`func (o *GETParcels200ResponseDataInnerAttributes) HasCustomsCertify() bool`

HasCustomsCertify returns a boolean if a field has been set.

### GetCustomsSigner

`func (o *GETParcels200ResponseDataInnerAttributes) GetCustomsSigner() string`

GetCustomsSigner returns the CustomsSigner field if non-nil, zero value otherwise.

### GetCustomsSignerOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetCustomsSignerOk() (*string, bool)`

GetCustomsSignerOk returns a tuple with the CustomsSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsSigner

`func (o *GETParcels200ResponseDataInnerAttributes) SetCustomsSigner(v string)`

SetCustomsSigner sets CustomsSigner field to given value.

### HasCustomsSigner

`func (o *GETParcels200ResponseDataInnerAttributes) HasCustomsSigner() bool`

HasCustomsSigner returns a boolean if a field has been set.

### GetNonDeliveryOption

`func (o *GETParcels200ResponseDataInnerAttributes) GetNonDeliveryOption() string`

GetNonDeliveryOption returns the NonDeliveryOption field if non-nil, zero value otherwise.

### GetNonDeliveryOptionOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetNonDeliveryOptionOk() (*string, bool)`

GetNonDeliveryOptionOk returns a tuple with the NonDeliveryOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDeliveryOption

`func (o *GETParcels200ResponseDataInnerAttributes) SetNonDeliveryOption(v string)`

SetNonDeliveryOption sets NonDeliveryOption field to given value.

### HasNonDeliveryOption

`func (o *GETParcels200ResponseDataInnerAttributes) HasNonDeliveryOption() bool`

HasNonDeliveryOption returns a boolean if a field has been set.

### GetRestrictionType

`func (o *GETParcels200ResponseDataInnerAttributes) GetRestrictionType() string`

GetRestrictionType returns the RestrictionType field if non-nil, zero value otherwise.

### GetRestrictionTypeOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetRestrictionTypeOk() (*string, bool)`

GetRestrictionTypeOk returns a tuple with the RestrictionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionType

`func (o *GETParcels200ResponseDataInnerAttributes) SetRestrictionType(v string)`

SetRestrictionType sets RestrictionType field to given value.

### HasRestrictionType

`func (o *GETParcels200ResponseDataInnerAttributes) HasRestrictionType() bool`

HasRestrictionType returns a boolean if a field has been set.

### GetRestrictionComments

`func (o *GETParcels200ResponseDataInnerAttributes) GetRestrictionComments() string`

GetRestrictionComments returns the RestrictionComments field if non-nil, zero value otherwise.

### GetRestrictionCommentsOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetRestrictionCommentsOk() (*string, bool)`

GetRestrictionCommentsOk returns a tuple with the RestrictionComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionComments

`func (o *GETParcels200ResponseDataInnerAttributes) SetRestrictionComments(v string)`

SetRestrictionComments sets RestrictionComments field to given value.

### HasRestrictionComments

`func (o *GETParcels200ResponseDataInnerAttributes) HasRestrictionComments() bool`

HasRestrictionComments returns a boolean if a field has been set.

### GetCustomsInfoRequired

`func (o *GETParcels200ResponseDataInnerAttributes) GetCustomsInfoRequired() bool`

GetCustomsInfoRequired returns the CustomsInfoRequired field if non-nil, zero value otherwise.

### GetCustomsInfoRequiredOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetCustomsInfoRequiredOk() (*bool, bool)`

GetCustomsInfoRequiredOk returns a tuple with the CustomsInfoRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsInfoRequired

`func (o *GETParcels200ResponseDataInnerAttributes) SetCustomsInfoRequired(v bool)`

SetCustomsInfoRequired sets CustomsInfoRequired field to given value.

### HasCustomsInfoRequired

`func (o *GETParcels200ResponseDataInnerAttributes) HasCustomsInfoRequired() bool`

HasCustomsInfoRequired returns a boolean if a field has been set.

### GetShippingLabelUrl

`func (o *GETParcels200ResponseDataInnerAttributes) GetShippingLabelUrl() string`

GetShippingLabelUrl returns the ShippingLabelUrl field if non-nil, zero value otherwise.

### GetShippingLabelUrlOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetShippingLabelUrlOk() (*string, bool)`

GetShippingLabelUrlOk returns a tuple with the ShippingLabelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelUrl

`func (o *GETParcels200ResponseDataInnerAttributes) SetShippingLabelUrl(v string)`

SetShippingLabelUrl sets ShippingLabelUrl field to given value.

### HasShippingLabelUrl

`func (o *GETParcels200ResponseDataInnerAttributes) HasShippingLabelUrl() bool`

HasShippingLabelUrl returns a boolean if a field has been set.

### GetShippingLabelFileType

`func (o *GETParcels200ResponseDataInnerAttributes) GetShippingLabelFileType() string`

GetShippingLabelFileType returns the ShippingLabelFileType field if non-nil, zero value otherwise.

### GetShippingLabelFileTypeOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetShippingLabelFileTypeOk() (*string, bool)`

GetShippingLabelFileTypeOk returns a tuple with the ShippingLabelFileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelFileType

`func (o *GETParcels200ResponseDataInnerAttributes) SetShippingLabelFileType(v string)`

SetShippingLabelFileType sets ShippingLabelFileType field to given value.

### HasShippingLabelFileType

`func (o *GETParcels200ResponseDataInnerAttributes) HasShippingLabelFileType() bool`

HasShippingLabelFileType returns a boolean if a field has been set.

### GetShippingLabelSize

`func (o *GETParcels200ResponseDataInnerAttributes) GetShippingLabelSize() string`

GetShippingLabelSize returns the ShippingLabelSize field if non-nil, zero value otherwise.

### GetShippingLabelSizeOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetShippingLabelSizeOk() (*string, bool)`

GetShippingLabelSizeOk returns a tuple with the ShippingLabelSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelSize

`func (o *GETParcels200ResponseDataInnerAttributes) SetShippingLabelSize(v string)`

SetShippingLabelSize sets ShippingLabelSize field to given value.

### HasShippingLabelSize

`func (o *GETParcels200ResponseDataInnerAttributes) HasShippingLabelSize() bool`

HasShippingLabelSize returns a boolean if a field has been set.

### GetShippingLabelResolution

`func (o *GETParcels200ResponseDataInnerAttributes) GetShippingLabelResolution() string`

GetShippingLabelResolution returns the ShippingLabelResolution field if non-nil, zero value otherwise.

### GetShippingLabelResolutionOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetShippingLabelResolutionOk() (*string, bool)`

GetShippingLabelResolutionOk returns a tuple with the ShippingLabelResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelResolution

`func (o *GETParcels200ResponseDataInnerAttributes) SetShippingLabelResolution(v string)`

SetShippingLabelResolution sets ShippingLabelResolution field to given value.

### HasShippingLabelResolution

`func (o *GETParcels200ResponseDataInnerAttributes) HasShippingLabelResolution() bool`

HasShippingLabelResolution returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *GETParcels200ResponseDataInnerAttributes) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *GETParcels200ResponseDataInnerAttributes) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetTrackingStatus

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingStatus() string`

GetTrackingStatus returns the TrackingStatus field if non-nil, zero value otherwise.

### GetTrackingStatusOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingStatusOk() (*string, bool)`

GetTrackingStatusOk returns a tuple with the TrackingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatus

`func (o *GETParcels200ResponseDataInnerAttributes) SetTrackingStatus(v string)`

SetTrackingStatus sets TrackingStatus field to given value.

### HasTrackingStatus

`func (o *GETParcels200ResponseDataInnerAttributes) HasTrackingStatus() bool`

HasTrackingStatus returns a boolean if a field has been set.

### GetTrackingStatusDetail

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingStatusDetail() string`

GetTrackingStatusDetail returns the TrackingStatusDetail field if non-nil, zero value otherwise.

### GetTrackingStatusDetailOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingStatusDetailOk() (*string, bool)`

GetTrackingStatusDetailOk returns a tuple with the TrackingStatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatusDetail

`func (o *GETParcels200ResponseDataInnerAttributes) SetTrackingStatusDetail(v string)`

SetTrackingStatusDetail sets TrackingStatusDetail field to given value.

### HasTrackingStatusDetail

`func (o *GETParcels200ResponseDataInnerAttributes) HasTrackingStatusDetail() bool`

HasTrackingStatusDetail returns a boolean if a field has been set.

### GetTrackingStatusUpdatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingStatusUpdatedAt() string`

GetTrackingStatusUpdatedAt returns the TrackingStatusUpdatedAt field if non-nil, zero value otherwise.

### GetTrackingStatusUpdatedAtOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingStatusUpdatedAtOk() (*string, bool)`

GetTrackingStatusUpdatedAtOk returns a tuple with the TrackingStatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatusUpdatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) SetTrackingStatusUpdatedAt(v string)`

SetTrackingStatusUpdatedAt sets TrackingStatusUpdatedAt field to given value.

### HasTrackingStatusUpdatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) HasTrackingStatusUpdatedAt() bool`

HasTrackingStatusUpdatedAt returns a boolean if a field has been set.

### GetTrackingDetails

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingDetails() string`

GetTrackingDetails returns the TrackingDetails field if non-nil, zero value otherwise.

### GetTrackingDetailsOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingDetailsOk() (*string, bool)`

GetTrackingDetailsOk returns a tuple with the TrackingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingDetails

`func (o *GETParcels200ResponseDataInnerAttributes) SetTrackingDetails(v string)`

SetTrackingDetails sets TrackingDetails field to given value.

### HasTrackingDetails

`func (o *GETParcels200ResponseDataInnerAttributes) HasTrackingDetails() bool`

HasTrackingDetails returns a boolean if a field has been set.

### GetCarrierWeightOz

`func (o *GETParcels200ResponseDataInnerAttributes) GetCarrierWeightOz() string`

GetCarrierWeightOz returns the CarrierWeightOz field if non-nil, zero value otherwise.

### GetCarrierWeightOzOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetCarrierWeightOzOk() (*string, bool)`

GetCarrierWeightOzOk returns a tuple with the CarrierWeightOz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierWeightOz

`func (o *GETParcels200ResponseDataInnerAttributes) SetCarrierWeightOz(v string)`

SetCarrierWeightOz sets CarrierWeightOz field to given value.

### HasCarrierWeightOz

`func (o *GETParcels200ResponseDataInnerAttributes) HasCarrierWeightOz() bool`

HasCarrierWeightOz returns a boolean if a field has been set.

### GetSignedBy

`func (o *GETParcels200ResponseDataInnerAttributes) GetSignedBy() string`

GetSignedBy returns the SignedBy field if non-nil, zero value otherwise.

### GetSignedByOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetSignedByOk() (*string, bool)`

GetSignedByOk returns a tuple with the SignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedBy

`func (o *GETParcels200ResponseDataInnerAttributes) SetSignedBy(v string)`

SetSignedBy sets SignedBy field to given value.

### HasSignedBy

`func (o *GETParcels200ResponseDataInnerAttributes) HasSignedBy() bool`

HasSignedBy returns a boolean if a field has been set.

### GetIncoterm

`func (o *GETParcels200ResponseDataInnerAttributes) GetIncoterm() string`

GetIncoterm returns the Incoterm field if non-nil, zero value otherwise.

### GetIncotermOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetIncotermOk() (*string, bool)`

GetIncotermOk returns a tuple with the Incoterm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncoterm

`func (o *GETParcels200ResponseDataInnerAttributes) SetIncoterm(v string)`

SetIncoterm sets Incoterm field to given value.

### HasIncoterm

`func (o *GETParcels200ResponseDataInnerAttributes) HasIncoterm() bool`

HasIncoterm returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETParcels200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETParcels200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETParcels200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETParcels200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETParcels200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETParcels200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETParcels200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETParcels200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETParcels200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


