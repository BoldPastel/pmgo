package pmgo

import (
	"encoding/xml"
)

//ref: https://developer.salesforce.com/docs/atlas.en-us.196.0.api_meta.meta/api_meta/meta_permissionset.htm#PermissionSetExternalDataSourceAccess_title
type Permissions struct {
	XMLName                    xml.Name                     `PermissionSet`
	xmlns                      string                       "http://soap.sforce.com/2006/04/metadata"
	ApplicationVisibilities    []ApplicationVisibilities    `xml:"applicationVisibilities"`
	ClassAccesses              []ClassAccesses              `xml:"classAccesses"`
	CustomPermissions          []CustomPermissions          `xml:"customPermissions"`
	Description                string                       `xml:"description"`
	ExternalDataSourceAccesses []ExternalDataSourceAccesses `xml:"externalDataSourceAccesses"`
	FieldPermissions           []FieldPermissions           `xml:"fieldPermissions"`
	HasActivationRequired      string                       `xml:"hasActivationRequired"`
	Label                      string                       `xml:"label"`
	ObjectPermissions          []ObjectPermissions          `xml:"objectPermissions"`
	PageAccesses               []PageAccesses               `xml:"pageAccesses"`
	RecordTypeVisibilities     []RecordTypeVisibilities     `xml:"recordTypeVisibilities"`
	TabSettings                []TabSettings                `xml:"tabSettings"`
	//UserLicense                string                       `xml:"userLicense"`
	UserPermissions []UserPermissions `xml:"userPermissions"`
}

type ExternalDataSourceAccesses struct {
	Enabled            string `xml:"enabled"`
	ExternalDataSource string `xml:"externalDataSource"`
}

type PageAccesses struct {
	ApexPage string `xml:"apexPage"`
	Enabled  string `xml:"enabled"`
}

type UserPermissions struct {
	Enabled string `xml:"enabled"`
	Name    string `xml:"name"`
}
type CustomPermissions struct {
	Enabled string `xml:"enabled"`
	Name    string `xml:"name"`
}

type RecordTypeVisibilities struct {
	RecordType string `xml:"recordType"`
	Visible    string `xml:"visible"`
}
type ApplicationVisibilities struct {
	Application string `xml:"application"`
	Visible     string `xml:"visible"`
}
type ClassAccesses struct {
	ApexClass string `xml:"apexClass"`
	Visible   string `xml:"enabled"`
}

type FieldPermissions struct {
	Editable string `xml:"editable"`
	Field    string `xml:"field"`
	Readable string `xml:"readable"`
}
type ObjectPermissions struct {
	AllowCreate      string `xml:"allowCreate"`
	AllowDelete      string `xml:"allowDelete"`
	AllowEdit        string `xml:"allowEdit"`
	AllowRead        string `xml:"allowRead"`
	ModifyAllRecords string `xml:"modifyAllRecords"`
	Object           string `xml:"object"`
	ViewAllRecords   string `xml:"viewAllRecords"`
}

type TabSettings struct {
	Tab        string `xml:"tab"`
	Visibility string `xml:"visibility"`
}
