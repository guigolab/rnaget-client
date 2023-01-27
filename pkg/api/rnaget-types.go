// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"time"
)

const (
	Rnaget_authScopes = "rnaget_auth.Scopes"
)

// Service defines model for Service.
type Service struct {
	// ContactUrl URL of the contact for the provider of this service, e.g. a link to a contact form (RFC 3986 format), or an email (RFC 2368 format).
	ContactUrl *string `json:"contactUrl,omitempty"`

	// CreatedAt Timestamp describing when the service was first deployed and available (RFC 3339 format)
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// Description Description of the service. Should be human readable and provide information about the service.
	Description *string `json:"description,omitempty"`

	// DocumentationUrl URL of the documentation of this service (RFC 3986 format). This should help someone learn how to use your service, including any specifics required to access data, e.g. authentication.
	DocumentationUrl *string `json:"documentationUrl,omitempty"`

	// Environment Environment the service is running in. Use this to distinguish between production, development and testing/staging deployments. Suggested values are prod, test, dev, staging. However this is advised and not enforced.
	Environment *string `json:"environment,omitempty"`

	// Id Unique ID of this service. Reverse domain name notation is recommended, though not required. The identifier should attempt to be globally unique so it can be used in downstream aggregator services e.g. Service Registry.
	Id string `json:"id"`

	// Name Name of this service. Should be human readable.
	Name string `json:"name"`

	// Organization Organization providing the service
	Organization struct {
		// Name Name of the organization responsible for the service
		Name string `json:"name"`

		// Url URL of the website of the organization (RFC 3986 format)
		Url string `json:"url"`
	} `json:"organization"`

	// Supported A true value indicates that the corresponding route is implemented by the service.  A non-implemented route should have a false value and return a 501 error to any requests.  If any boolean property is not provided, it is assumed to be implemented. If the entire supported object is not provided, all endpoints are assumed to be implemented.
	Supported *Supported `json:"supported,omitempty"`
	Type      struct {
		Artifact *interface{} `json:"artifact,omitempty"`
		Version  *interface{} `json:"version,omitempty"`
	} `json:"type"`

	// UpdatedAt Timestamp describing when the service was last updated (RFC 3339 format)
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Version Version of the service being described. Semantic versioning is recommended, but other identifiers, such as dates or commit hashes, are also allowed. The version should be changed whenever the service is updated.
	Version string `json:"version"`
}

// Supported A true value indicates that the corresponding route is implemented by the service.  A non-implemented route should have a false value and return a 501 error to any requests.  If any boolean property is not provided, it is assumed to be implemented. If the entire supported object is not provided, all endpoints are assumed to be implemented.
type Supported struct {
	Continuous  *bool `json:"continuous,omitempty"`
	Expressions *bool `json:"expressions,omitempty"`
	Projects    *bool `json:"projects,omitempty"`
	Studies     *bool `json:"studies,omitempty"`
}

// Error General API error model
type Error struct {
	// Message Error message details
	Message *string `json:"message,omitempty"`
}

// Filter Implementation defined parameter to use for filtering collections
type Filter struct {
	// Description Detailed description of the filter
	Description *string `json:"description,omitempty"`

	// FieldType The dataType (string, float, etc.) of the filter
	FieldType *string `json:"fieldType,omitempty"`

	// Filter A unique name for the filter for use in query URLs
	Filter string `json:"filter"`

	// Values List of supported values for the filter
	Values *[]string `json:"values,omitempty"`
}

// Project The project is the top level of the model hierarchy and contains one or more studies.
type Project struct {
	// Description Detailed description of the object
	Description *string `json:"description,omitempty"`

	// Id A unique identifier assigned to this object
	Id string `json:"id"`

	// Name Short, readable name
	Name *string `json:"name,omitempty"`

	// Version Version number of the object
	Version *string `json:"version,omitempty"`
}

// Study The study is a container for one or more related RNA expression matrices.
type Study struct {
	// Description Detailed description of the object
	Description *string `json:"description,omitempty"`

	// Genome Name of the reference genome build used for aligning samples in the study
	Genome *string `json:"genome,omitempty"`

	// Id A unique identifier assigned to this object
	Id string `json:"id"`

	// Name Short, readable name
	Name *string `json:"name,omitempty"`

	// ParentProjectID ID of the project containing the study
	ParentProjectID *string `json:"parentProjectID,omitempty"`

	// Version Version number of the object
	Version *string `json:"version,omitempty"`
}

// Ticket URL and type for data files
type Ticket struct {
	// FileType Type of file. Examples include: loom, tsv
	FileType *string `json:"fileType,omitempty"`

	// Headers For HTTPS URLs, the server may supply a JSON object containing one or more string key-value pairs which the client MUST supply verbatim as headers with any request to the URL. For example, if headers is `{"Authorization": "Bearer xxxx"}`, then the client must supply the header `Authorization: Bearer xxxx` with the HTTPS request to the URL.
	Headers *map[string]interface{} `json:"headers,omitempty"`

	// Md5 MD5 digest of the file
	Md5 *string `json:"md5,omitempty"`

	// StudyID ID of containing study
	StudyID *string `json:"studyID,omitempty"`

	// Units Units for the values. Examples include: TPM, FPKM, counts
	Units string `json:"units"`

	// Url An `https:` URL to download file
	Url string `json:"url"`

	// Version Version number of the object
	Version *string `json:"version,omitempty"`
}

// ChrParam defines model for chrParam.
type ChrParam = string

// EndParam defines model for endParam.
type EndParam = int32

// FeatureIDListParam defines model for featureIDListParam.
type FeatureIDListParam = []string

// FeatureMaxParam defines model for featureMaxParam.
type FeatureMaxParam = float32

// FeatureMinParam defines model for featureMinParam.
type FeatureMinParam = float32

// FeatureNameListParam defines model for featureNameListParam.
type FeatureNameListParam = []string

// ProjectIDParam defines model for projectIDParam.
type ProjectIDParam = string

// SampleIDListParam defines model for sampleIDListParam.
type SampleIDListParam = []string

// StartParam defines model for startParam.
type StartParam = int32

// StudyIDParam defines model for studyIDParam.
type StudyIDParam = string

// UnitsParam defines model for unitsParam.
type UnitsParam = string

// VersionParam defines model for versionParam.
type VersionParam = string

// GetContinuousFileParams defines parameters for GetContinuousFile.
type GetContinuousFileParams struct {
	// Format Data format to return
	Format string `form:"format" json:"format"`

	// ProjectID project to filter by
	ProjectID *ProjectIDParam `form:"projectID,omitempty" json:"projectID,omitempty"`

	// StudyID study to filter by
	StudyID *StudyIDParam `form:"studyID,omitempty" json:"studyID,omitempty"`

	// Version version to filter by
	Version *VersionParam `form:"version,omitempty" json:"version,omitempty"`

	// SampleIDList return only values for listed sampleIDs
	SampleIDList *SampleIDListParam `form:"sampleIDList,omitempty" json:"sampleIDList,omitempty"`

	// Chr The refererence to which start and end apply in the form chr? where ? is the specific ID of the chromosome (ex. chr1, chrX).
	Chr *ChrParam `form:"chr,omitempty" json:"chr,omitempty"`

	// Start The start position of the range on the sequence, 0-based, inclusive.
	Start *StartParam `form:"start,omitempty" json:"start,omitempty"`

	// End The end position of the range on the sequence, 0-based, exclusive.
	End *EndParam `form:"end,omitempty" json:"end,omitempty"`
}

// GetContinuousTicketParams defines parameters for GetContinuousTicket.
type GetContinuousTicketParams struct {
	// Format Data format to return
	Format string `form:"format" json:"format"`

	// ProjectID project to filter by
	ProjectID *ProjectIDParam `form:"projectID,omitempty" json:"projectID,omitempty"`

	// StudyID study to filter by
	StudyID *StudyIDParam `form:"studyID,omitempty" json:"studyID,omitempty"`

	// Version version to filter by
	Version *VersionParam `form:"version,omitempty" json:"version,omitempty"`

	// SampleIDList return only values for listed sampleIDs
	SampleIDList *SampleIDListParam `form:"sampleIDList,omitempty" json:"sampleIDList,omitempty"`

	// Chr The refererence to which start and end apply in the form chr? where ? is the specific ID of the chromosome (ex. chr1, chrX).
	Chr *ChrParam `form:"chr,omitempty" json:"chr,omitempty"`

	// Start The start position of the range on the sequence, 0-based, inclusive.
	Start *StartParam `form:"start,omitempty" json:"start,omitempty"`

	// End The end position of the range on the sequence, 0-based, exclusive.
	End *EndParam `form:"end,omitempty" json:"end,omitempty"`
}

// GetContinuousFileByIdParams defines parameters for GetContinuousFileById.
type GetContinuousFileByIdParams struct {
	// Chr The refererence to which start and end apply in the form chr? where ? is the specific ID of the chromosome (ex. chr1, chrX).
	Chr *ChrParam `form:"chr,omitempty" json:"chr,omitempty"`

	// Start The start position of the range on the sequence, 0-based, inclusive.
	Start *StartParam `form:"start,omitempty" json:"start,omitempty"`

	// End The end position of the range on the sequence, 0-based, exclusive.
	End *EndParam `form:"end,omitempty" json:"end,omitempty"`
}

// GetContinuousTicketByIdParams defines parameters for GetContinuousTicketById.
type GetContinuousTicketByIdParams struct {
	// Chr The refererence to which start and end apply in the form chr? where ? is the specific ID of the chromosome (ex. chr1, chrX).
	Chr *ChrParam `form:"chr,omitempty" json:"chr,omitempty"`

	// Start The start position of the range on the sequence, 0-based, inclusive.
	Start *StartParam `form:"start,omitempty" json:"start,omitempty"`

	// End The end position of the range on the sequence, 0-based, exclusive.
	End *EndParam `form:"end,omitempty" json:"end,omitempty"`
}

// GetExpressionFileParams defines parameters for GetExpressionFile.
type GetExpressionFileParams struct {
	// Format Data format to return
	Format string `form:"format" json:"format"`

	// ProjectID project to filter by
	ProjectID *ProjectIDParam `form:"projectID,omitempty" json:"projectID,omitempty"`

	// StudyID study to filter by
	StudyID *StudyIDParam `form:"studyID,omitempty" json:"studyID,omitempty"`

	// Version version to filter by
	Version *VersionParam `form:"version,omitempty" json:"version,omitempty"`

	// SampleIDList return only values for listed sampleIDs
	SampleIDList *SampleIDListParam `form:"sampleIDList,omitempty" json:"sampleIDList,omitempty"`

	// FeatureIDList return only values for listed feature IDs
	FeatureIDList *FeatureIDListParam `form:"featureIDList,omitempty" json:"featureIDList,omitempty"`

	// FeatureNameList return only values for listed features
	FeatureNameList *FeatureNameListParam `form:"featureNameList,omitempty" json:"featureNameList,omitempty"`

	// FeatureMinValue Sets a minimum expression value for features. If set the resulting matrix will not contain any features with a value less than the provided threshold in any sample in the matrix. The resulting matrix should have values >= the supplied threshold in every cell. Using this with feature_max_value will result in matrix values in the closed interval [feature_min_value, feature_max_value].
	FeatureMinValue *FeatureMinParam `form:"feature_min_value,omitempty" json:"feature_min_value,omitempty"`

	// FeatureMaxValue Sets a maximum expression value for features. If set the resulting matrix will not contain any features with a value greater than the provided threshold in any sample in the matrix. The resulting matrix should have values <= the supplied threshold in every cell. Using this with feature_min_value will result in matrix values in the closed interval [feature_min_value, feature_max_value].
	FeatureMaxValue *FeatureMaxParam `form:"feature_max_value,omitempty" json:"feature_max_value,omitempty"`

	// Units The values in the matrix will be those corresponding to the requested units.  If present the value provided MUST match an item in the list returned by a request to /expressions/units.
	Units *UnitsParam `form:"units,omitempty" json:"units,omitempty"`
}

// GetExpressionFiltersParams defines parameters for GetExpressionFilters.
type GetExpressionFiltersParams struct {
	// Type one of `feature` or `sample` reflecting which axis to request filters for.  If blank, both will be returned
	Type *string `form:"type,omitempty" json:"type,omitempty"`
}

// GetExpressionTicketParams defines parameters for GetExpressionTicket.
type GetExpressionTicketParams struct {
	// Format Data format to return
	Format string `form:"format" json:"format"`

	// ProjectID project to filter by
	ProjectID *ProjectIDParam `form:"projectID,omitempty" json:"projectID,omitempty"`

	// StudyID study to filter by
	StudyID *StudyIDParam `form:"studyID,omitempty" json:"studyID,omitempty"`

	// Version version to filter by
	Version *VersionParam `form:"version,omitempty" json:"version,omitempty"`

	// SampleIDList return only values for listed sampleIDs
	SampleIDList *SampleIDListParam `form:"sampleIDList,omitempty" json:"sampleIDList,omitempty"`

	// FeatureIDList return only values for listed feature IDs
	FeatureIDList *FeatureIDListParam `form:"featureIDList,omitempty" json:"featureIDList,omitempty"`

	// FeatureNameList return only values for listed features
	FeatureNameList *FeatureNameListParam `form:"featureNameList,omitempty" json:"featureNameList,omitempty"`

	// FeatureMinValue Sets a minimum expression value for features. If set the resulting matrix will not contain any features with a value less than the provided threshold in any sample in the matrix. The resulting matrix should have values >= the supplied threshold in every cell. Using this with feature_max_value will result in matrix values in the closed interval [feature_min_value, feature_max_value].
	FeatureMinValue *FeatureMinParam `form:"feature_min_value,omitempty" json:"feature_min_value,omitempty"`

	// FeatureMaxValue Sets a maximum expression value for features. If set the resulting matrix will not contain any features with a value greater than the provided threshold in any sample in the matrix. The resulting matrix should have values <= the supplied threshold in every cell. Using this with feature_min_value will result in matrix values in the closed interval [feature_min_value, feature_max_value].
	FeatureMaxValue *FeatureMaxParam `form:"feature_max_value,omitempty" json:"feature_max_value,omitempty"`

	// Units The values in the matrix will be those corresponding to the requested units.  If present the value provided MUST match an item in the list returned by a request to /expressions/units.
	Units *UnitsParam `form:"units,omitempty" json:"units,omitempty"`
}

// GetExpressionFileByIdParams defines parameters for GetExpressionFileById.
type GetExpressionFileByIdParams struct {
	// SampleIDList return only values for listed sampleIDs
	SampleIDList *SampleIDListParam `form:"sampleIDList,omitempty" json:"sampleIDList,omitempty"`

	// FeatureIDList return only values for listed feature IDs
	FeatureIDList *FeatureIDListParam `form:"featureIDList,omitempty" json:"featureIDList,omitempty"`

	// FeatureNameList return only values for listed features
	FeatureNameList *FeatureNameListParam `form:"featureNameList,omitempty" json:"featureNameList,omitempty"`

	// FeatureMinValue Sets a minimum expression value for features. If set the resulting matrix will not contain any features with a value less than the provided threshold in any sample in the matrix. The resulting matrix should have values >= the supplied threshold in every cell. Using this with feature_max_value will result in matrix values in the closed interval [feature_min_value, feature_max_value].
	FeatureMinValue *FeatureMinParam `form:"feature_min_value,omitempty" json:"feature_min_value,omitempty"`

	// FeatureMaxValue Sets a maximum expression value for features. If set the resulting matrix will not contain any features with a value greater than the provided threshold in any sample in the matrix. The resulting matrix should have values <= the supplied threshold in every cell. Using this with feature_min_value will result in matrix values in the closed interval [feature_min_value, feature_max_value].
	FeatureMaxValue *FeatureMaxParam `form:"feature_max_value,omitempty" json:"feature_max_value,omitempty"`

	// Units The values in the matrix will be those corresponding to the requested units.  If present the value provided MUST match an item in the list returned by a request to /expressions/units.
	Units *UnitsParam `form:"units,omitempty" json:"units,omitempty"`
}

// GetExpressionTicketByIdParams defines parameters for GetExpressionTicketById.
type GetExpressionTicketByIdParams struct {
	// SampleIDList return only values for listed sampleIDs
	SampleIDList *SampleIDListParam `form:"sampleIDList,omitempty" json:"sampleIDList,omitempty"`

	// FeatureIDList return only values for listed feature IDs
	FeatureIDList *FeatureIDListParam `form:"featureIDList,omitempty" json:"featureIDList,omitempty"`

	// FeatureNameList return only values for listed features
	FeatureNameList *FeatureNameListParam `form:"featureNameList,omitempty" json:"featureNameList,omitempty"`

	// FeatureMinValue Sets a minimum expression value for features. If set the resulting matrix will not contain any features with a value less than the provided threshold in any sample in the matrix. The resulting matrix should have values >= the supplied threshold in every cell. Using this with feature_max_value will result in matrix values in the closed interval [feature_min_value, feature_max_value].
	FeatureMinValue *FeatureMinParam `form:"feature_min_value,omitempty" json:"feature_min_value,omitempty"`

	// FeatureMaxValue Sets a maximum expression value for features. If set the resulting matrix will not contain any features with a value greater than the provided threshold in any sample in the matrix. The resulting matrix should have values <= the supplied threshold in every cell. Using this with feature_min_value will result in matrix values in the closed interval [feature_min_value, feature_max_value].
	FeatureMaxValue *FeatureMaxParam `form:"feature_max_value,omitempty" json:"feature_max_value,omitempty"`

	// Units The values in the matrix will be those corresponding to the requested units.  If present the value provided MUST match an item in the list returned by a request to /expressions/units.
	Units *UnitsParam `form:"units,omitempty" json:"units,omitempty"`
}

// GetProjectsParams defines parameters for GetProjects.
type GetProjectsParams struct {
	// Version version to filter by
	Version *VersionParam `form:"version,omitempty" json:"version,omitempty"`
}

// GetStudiesParams defines parameters for GetStudies.
type GetStudiesParams struct {
	// Version version to filter by
	Version *VersionParam `form:"version,omitempty" json:"version,omitempty"`
}
