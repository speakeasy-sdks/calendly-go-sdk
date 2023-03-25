// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
)

type LocationType string

const (
	LocationTypeInPersonMeeting          LocationType = "InPersonMeeting"
	LocationTypeOutboundCall             LocationType = "OutboundCall"
	LocationTypeInboundCall              LocationType = "InboundCall"
	LocationTypeGoogleConference         LocationType = "GoogleConference"
	LocationTypeZoomConference           LocationType = "ZoomConference"
	LocationTypeGotoMeetingConference    LocationType = "GotoMeetingConference"
	LocationTypeMicrosoftTeamsConference LocationType = "MicrosoftTeamsConference"
	LocationTypeCustomLocation           LocationType = "CustomLocation"
	LocationTypeInviteeSpecifiedLocation LocationType = "InviteeSpecifiedLocation"
	LocationTypeWebExConference          LocationType = "WebExConference"
)

type Location struct {
	InPersonMeeting          *InPersonMeeting
	OutboundCall             *OutboundCall
	InboundCall              *InboundCall
	GoogleConference         *GoogleConference
	ZoomConference           *ZoomConference
	GotoMeetingConference    *GotoMeetingConference
	MicrosoftTeamsConference *MicrosoftTeamsConference
	CustomLocation           *CustomLocation
	InviteeSpecifiedLocation *InviteeSpecifiedLocation
	WebExConference          *WebExConference

	Type LocationType
}

func CreateLocationInPersonMeeting(inPersonMeeting InPersonMeeting) Location {
	typ := LocationTypeInPersonMeeting

	return Location{
		InPersonMeeting: &inPersonMeeting,
		Type:            typ,
	}
}

func CreateLocationOutboundCall(outboundCall OutboundCall) Location {
	typ := LocationTypeOutboundCall

	return Location{
		OutboundCall: &outboundCall,
		Type:         typ,
	}
}

func CreateLocationInboundCall(inboundCall InboundCall) Location {
	typ := LocationTypeInboundCall

	return Location{
		InboundCall: &inboundCall,
		Type:        typ,
	}
}

func CreateLocationGoogleConference(googleConference GoogleConference) Location {
	typ := LocationTypeGoogleConference

	return Location{
		GoogleConference: &googleConference,
		Type:             typ,
	}
}

func CreateLocationZoomConference(zoomConference ZoomConference) Location {
	typ := LocationTypeZoomConference

	return Location{
		ZoomConference: &zoomConference,
		Type:           typ,
	}
}

func CreateLocationGotoMeetingConference(gotoMeetingConference GotoMeetingConference) Location {
	typ := LocationTypeGotoMeetingConference

	return Location{
		GotoMeetingConference: &gotoMeetingConference,
		Type:                  typ,
	}
}

func CreateLocationMicrosoftTeamsConference(microsoftTeamsConference MicrosoftTeamsConference) Location {
	typ := LocationTypeMicrosoftTeamsConference

	return Location{
		MicrosoftTeamsConference: &microsoftTeamsConference,
		Type:                     typ,
	}
}

func CreateLocationCustomLocation(customLocation CustomLocation) Location {
	typ := LocationTypeCustomLocation

	return Location{
		CustomLocation: &customLocation,
		Type:           typ,
	}
}

func CreateLocationInviteeSpecifiedLocation(inviteeSpecifiedLocation InviteeSpecifiedLocation) Location {
	typ := LocationTypeInviteeSpecifiedLocation

	return Location{
		InviteeSpecifiedLocation: &inviteeSpecifiedLocation,
		Type:                     typ,
	}
}

func CreateLocationWebExConference(webExConference WebExConference) Location {
	typ := LocationTypeWebExConference

	return Location{
		WebExConference: &webExConference,
		Type:            typ,
	}
}

func (u *Location) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	inPersonMeeting := new(InPersonMeeting)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inPersonMeeting); err == nil {
		u.InPersonMeeting = inPersonMeeting
		u.Type = LocationTypeInPersonMeeting
		return nil
	}

	outboundCall := new(OutboundCall)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&outboundCall); err == nil {
		u.OutboundCall = outboundCall
		u.Type = LocationTypeOutboundCall
		return nil
	}

	inboundCall := new(InboundCall)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inboundCall); err == nil {
		u.InboundCall = inboundCall
		u.Type = LocationTypeInboundCall
		return nil
	}

	googleConference := new(GoogleConference)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&googleConference); err == nil {
		u.GoogleConference = googleConference
		u.Type = LocationTypeGoogleConference
		return nil
	}

	zoomConference := new(ZoomConference)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&zoomConference); err == nil {
		u.ZoomConference = zoomConference
		u.Type = LocationTypeZoomConference
		return nil
	}

	gotoMeetingConference := new(GotoMeetingConference)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&gotoMeetingConference); err == nil {
		u.GotoMeetingConference = gotoMeetingConference
		u.Type = LocationTypeGotoMeetingConference
		return nil
	}

	microsoftTeamsConference := new(MicrosoftTeamsConference)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&microsoftTeamsConference); err == nil {
		u.MicrosoftTeamsConference = microsoftTeamsConference
		u.Type = LocationTypeMicrosoftTeamsConference
		return nil
	}

	customLocation := new(CustomLocation)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&customLocation); err == nil {
		u.CustomLocation = customLocation
		u.Type = LocationTypeCustomLocation
		return nil
	}

	inviteeSpecifiedLocation := new(InviteeSpecifiedLocation)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inviteeSpecifiedLocation); err == nil {
		u.InviteeSpecifiedLocation = inviteeSpecifiedLocation
		u.Type = LocationTypeInviteeSpecifiedLocation
		return nil
	}

	webExConference := new(WebExConference)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&webExConference); err == nil {
		u.WebExConference = webExConference
		u.Type = LocationTypeWebExConference
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Location) MarshalJSON() ([]byte, error) {
	if u.InPersonMeeting != nil {
		return json.Marshal(u.InPersonMeeting)
	}

	if u.OutboundCall != nil {
		return json.Marshal(u.OutboundCall)
	}

	if u.InboundCall != nil {
		return json.Marshal(u.InboundCall)
	}

	if u.GoogleConference != nil {
		return json.Marshal(u.GoogleConference)
	}

	if u.ZoomConference != nil {
		return json.Marshal(u.ZoomConference)
	}

	if u.GotoMeetingConference != nil {
		return json.Marshal(u.GotoMeetingConference)
	}

	if u.MicrosoftTeamsConference != nil {
		return json.Marshal(u.MicrosoftTeamsConference)
	}

	if u.CustomLocation != nil {
		return json.Marshal(u.CustomLocation)
	}

	if u.InviteeSpecifiedLocation != nil {
		return json.Marshal(u.InviteeSpecifiedLocation)
	}

	if u.WebExConference != nil {
		return json.Marshal(u.WebExConference)
	}

	return nil, nil
}