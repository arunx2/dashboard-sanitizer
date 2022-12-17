package model

import "time"

type References struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type DashboardObject struct {
	Attributes struct {
		Description           string `json:"description,omitempty"`
		Hits                  int    `json:"hits,omitempty"`
		KibanaSavedObjectMeta struct {
			SearchSourceJSON string `json:"searchSourceJSON,omitempty"`
		} `json:"kibanaSavedObjectMeta,omitempty"`
		OptionsJSON string `json:"optionsJSON,omitempty"`
		PanelsJSON  string `json:"panelsJSON,omitempty"`
		TimeRestore bool   `json:"timeRestore,omitempty"`
		Title       string `json:"title,omitempty"`
		//Index-patter related START
		RuntimeFieldMap string `json:"runtimeFieldMap,omitempty"`
		FieldAttrs      string `json:"fieldAttrs,omitempty"`
		FieldFormatMap  string `json:"fieldFormatMap,omitempty"`
		Fields          string `json:"fields,omitempty"`
		SourceFilters   string `json:"sourceFilters,omitempty"`
		TimeFieldName   string `json:"timeFieldName,omitempty"`
		TypeMeta        string `json:"typeMeta,omitempty"`
		//Index-patter related END
		VisualizationType string `json:"visualizationType,omitempty"`
		UIStateJSON       string `json:"uiStateJSON,omitempty"`
		VisState          string `json:"visState,omitempty"`
		Version           int    `json:"version"`
	} `json:"attributes,omitempty"`
	CoreMigrationVersion string `json:"coreMigrationVersion,omitempty"`
	ID                   string `json:"id,omitempty"`
	MigrationVersion     struct {
		Lens          string `json:"lens,omitempty"`          //should be removed for OpenSearch
		Dashboard     string `json:"dashboard,omitempty"`     // the version should be 7.9.3 or less
		IndexPattern  string `json:"index-pattern,omitempty"` // the version should be 7.6.0 or less
		Visualization string `json:"visualization,omitempty"` // the version should be 7.9.3 or less
	} `json:"migrationVersion,omitempty"`
	References []References `json:"references,omitempty"`
	Type       string       `json:"type,omitempty"`
	UpdatedAt  time.Time    `json:"updated_at,omitempty"`
	Version    string       `json:"version,omitempty"`
}
