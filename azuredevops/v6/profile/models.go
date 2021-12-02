﻿// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package profile

import (
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azuredevops/v6"
)

// Identifies an attribute with a name and a container.
type AttributeDescriptor struct {
    // The name of the attribute.
    AttributeName *string `json:"attributeName,omitempty"`
    // The container the attribute resides in.
    ContainerName *string `json:"containerName,omitempty"`
}

// Stores a set of named profile attributes.
type AttributesContainer struct {
    // The attributes stored by the container.
    Attributes *map[string]ProfileAttribute `json:"attributes,omitempty"`
    // The name of the container.
    ContainerName *string `json:"containerName,omitempty"`
    // The maximum revision number of any attribute within the container.
    Revision *int `json:"revision,omitempty"`
}

type Avatar struct {
    IsAutoGenerated *bool `json:"isAutoGenerated,omitempty"`
    Size *AvatarSize `json:"size,omitempty"`
    TimeStamp *azuredevops.Time `json:"timeStamp,omitempty"`
    Value *[]byte `json:"value,omitempty"`
}

// Small = 34 x 34 pixels; Medium = 44 x 44 pixels; Large = 220 x 220 pixels
type AvatarSize string

type avatarSizeValuesType struct {
    Small AvatarSize
    Medium AvatarSize
    Large AvatarSize
}

var AvatarSizeValues = avatarSizeValuesType{
    Small: "small",
    Medium: "medium",
    Large: "large",
}

// A profile attribute which always has a value for each profile.
type CoreProfileAttribute struct {
    // The descriptor of the attribute.
    Descriptor *AttributeDescriptor `json:"descriptor,omitempty"`
    // The revision number of the attribute.
    Revision *int `json:"revision,omitempty"`
    // The time the attribute was last changed.
    TimeStamp *azuredevops.Time `json:"timeStamp,omitempty"`
    // The value of the attribute.
    Value interface{} `json:"value,omitempty"`
}

type CreateProfileContext struct {
    CiData *map[string]interface{} `json:"ciData,omitempty"`
    ContactWithOffers *bool `json:"contactWithOffers,omitempty"`
    CountryName *string `json:"countryName,omitempty"`
    DisplayName *string `json:"displayName,omitempty"`
    EmailAddress *string `json:"emailAddress,omitempty"`
    HasAccount *bool `json:"hasAccount,omitempty"`
    Language *string `json:"language,omitempty"`
    PhoneNumber *string `json:"phoneNumber,omitempty"`
    // The current state of the profile.
    ProfileState *ProfileState `json:"profileState,omitempty"`
}

type GeoRegion struct {
    RegionCode *string `json:"regionCode,omitempty"`
}

// A user profile.
type Profile struct {
    // The attributes of this profile.
    ApplicationContainer *AttributesContainer `json:"applicationContainer,omitempty"`
    // The core attributes of this profile.
    CoreAttributes *map[string]CoreProfileAttribute `json:"coreAttributes,omitempty"`
    // The maximum revision number of any attribute.
    CoreRevision *int `json:"coreRevision,omitempty"`
    // The unique identifier of the profile.
    Id *uuid.UUID `json:"id,omitempty"`
    // The current state of the profile.
    ProfileState *ProfileState `json:"profileState,omitempty"`
    // The maximum revision number of any attribute.
    Revision *int `json:"revision,omitempty"`
    // The time at which this profile was last changed.
    TimeStamp *azuredevops.Time `json:"timeStamp,omitempty"`
}

// A named object associated with a profile.
type ProfileAttribute struct {
    // The descriptor of the attribute.
    Descriptor *AttributeDescriptor `json:"descriptor,omitempty"`
    // The revision number of the attribute.
    Revision *int `json:"revision,omitempty"`
    // The time the attribute was last changed.
    TimeStamp *azuredevops.Time `json:"timeStamp,omitempty"`
    // The value of the attribute.
    Value interface{} `json:"value,omitempty"`
}

type ProfileAttributeBase struct {
    // The descriptor of the attribute.
    Descriptor *AttributeDescriptor `json:"descriptor,omitempty"`
    // The revision number of the attribute.
    Revision *int `json:"revision,omitempty"`
    // The time the attribute was last changed.
    TimeStamp *azuredevops.Time `json:"timeStamp,omitempty"`
    // The value of the attribute.
    Value interface{} `json:"value,omitempty"`
}

// Country/region information
type ProfileRegion struct {
    // The two-letter code defined in ISO 3166 for the country/region.
    Code *string `json:"code,omitempty"`
    // Localized country/region name
    Name *string `json:"name,omitempty"`
}

// Container of country/region information
type ProfileRegions struct {
    // List of country/region code with contact consent requirement type of notice
    NoticeContactConsentRequirementRegions *[]string `json:"noticeContactConsentRequirementRegions,omitempty"`
    // List of country/region code with contact consent requirement type of opt-out
    OptOutContactConsentRequirementRegions *[]string `json:"optOutContactConsentRequirementRegions,omitempty"`
    // List of country/regions
    Regions *[]ProfileRegion `json:"regions,omitempty"`
}

// The state of a profile.
type ProfileState string

type profileStateValuesType struct {
    Custom ProfileState
    CustomReadOnly ProfileState
    ReadOnly ProfileState
}

var ProfileStateValues = profileStateValuesType{
    // The profile is in use.
    Custom: "custom",
    // The profile is in use, but can only be read.
    CustomReadOnly: "customReadOnly",
    // The profile may only be read.
    ReadOnly: "readOnly",
}
