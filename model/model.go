package model

import "time"

type DashboardObject struct {
	Attributes struct {
		Description           string `json:"description"`
		Hits                  int    `json:"hits"`
		KibanaSavedObjectMeta struct {
			SearchSourceJSON string `json:"searchSourceJSON"`
		} `json:"kibanaSavedObjectMeta"`
		OptionsJSON string `json:"optionsJSON"`
		PanelsJSON  string `json:"panelsJSON"`
		TimeRestore bool   `json:"timeRestore"`
		Title       string `json:"title"`
		//Index-patter related START
		RuntimeFieldMap string `json:"runtimeFieldMap"`
		FieldAttrs      string `json:"fieldAttrs"`
		FieldFormatMap  string `json:"fieldFormatMap"`
		Fields          string `json:"fields"`
		SourceFilters   string `json:"sourceFilters"`
		TimeFieldName   string `json:"timeFieldName"`
		TypeMeta        string `json:"typeMeta"`
		//Index-patter related END
		VisualizationType string `json:"visualizationType"`
		UIStateJSON       string `json:"uiStateJSON"`
		VisState          string `json:"visState"`
		Version           int    `json:"version"`
	} `json:"attributes"`
	CoreMigrationVersion string `json:"coreMigrationVersion"`
	ID                   string `json:"id"`
	MigrationVersion     struct {
		Lens          string `json:"lens"`          //should be removed for OpenSearch
		Dashboard     string `json:"dashboard"`     // the version should be 7.9.3 or less
		IndexPattern  string `json:"index-pattern"` // the version should be 7.6.0 or less
		Visualization string `json:"visualization"` // the version should be 7.9.3 or less
	} `json:"migrationVersion"`
	References []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"references"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
	Version   string    `json:"version"`
}
