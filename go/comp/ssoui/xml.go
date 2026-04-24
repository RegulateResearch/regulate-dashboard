package ssoui

import "encoding/xml"

type successResponse struct {
	XmlName     xml.Name    `xml:"serviceResponse"`
	AuthSuccess authSuccess `xml:"authenticationSuccess" validate:"required"`
}

func newSuccessResponse() successResponse {
	return successResponse{}
}

type authSuccess struct {
	XmlName  xml.Name   `xml:"authenticationSuccess" validate:"required"`
	Username string     `xml:"user" validate:"required"`
	Attr     attributes `xml:"attributes" validate:"required"`
}

type attributes struct {
	XmlName xml.Name `xml:"attributes" validate:"required"`
	Name    string   `xml:"nama" validate:"required"`
	Npm     string   `xml:"npm" validate:"required"`
}

type failureResponse struct {
	XmlName     xml.Name    `xml:"serviceResponse" validate:"required"`
	AuthFailure authFailure `xml:"authenticationFailure" validate:"required"`
}

func newFailureResponse() failureResponse {
	return failureResponse{}
}

type authFailure struct {
	XmlName xml.Name `xml:"authenticationFailure" validate:"required"`
	Code    string   `xml:"code,attr" validate:"required"`
	Message string   `xml:",chardata" validate:"required"`
}
