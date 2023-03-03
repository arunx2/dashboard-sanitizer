package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

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

func (do *DashboardObject) MakeCompatibleToOS() (err error) {
	switch do.Type {
	case "dashboard":
		//TODO: check if the version is greater than 7.9.3, leave the value as is if it is less than this.
		do.MigrationVersion.Dashboard = "7.9.3"
		do.SanitizePanelJSON()
		//fix some visualization references name
		var temp []References
		for i := range do.References {
			if isCompatibleObjectType(do.References[i].Type) {
				if do.References[i].Type == "visualization" {
					do.References[i].Name = getNormalizedVizName(do.References[i].Name)
				}
				temp = append(temp, do.References[i])
			}
		}
		do.References = temp
	case "visualization":
		do.MigrationVersion.Visualization = "7.9.3"
		break
	case "index-pattern":
		do.MigrationVersion.IndexPattern = "7.6.0"
		break
	}
	return
}

func getNormalizedVizName(s string) string {
	if idx := strings.Index(s, ":"); idx != -1 {
		return s[idx+1:]
	}
	return s
}

func (do *DashboardObject) IsCompatibleType() bool {
	return isCompatibleObjectType(do.Type)
}

func isCompatibleObjectType(objectType string) bool {
	switch objectType {
	case "", "lens", "map", "canvas-workpad", "canvas-element", "graph-workspace", "connector", "rule":
		return false
	}
	return true
}

// Removes all non-compatible object types from the panel json object
func (do *DashboardObject) SanitizePanelJSON() (err error) {
	var panels []map[string]interface{}

	err = json.Unmarshal([]byte(do.Attributes.PanelsJSON), &panels)
	if err != nil {
		return
	}
	var results []map[string]interface{}
	for _, panel := range panels {
		if !isCompatibleObjectType(fmt.Sprintf("%v", panel["type"])) {
			continue
		}
		results = append(results, panel)
	}
	resultBytes, er := json.Marshal(results)
	if er != nil {
		return er
	}
	do.Attributes.PanelsJSON = string(resultBytes)
	return
}
