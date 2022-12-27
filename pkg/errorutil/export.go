package errorutil

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/slog"
	"os"
	"path/filepath"
	"strconv"
)

func Export(componentInfo *Info, infoAll *InfoAll, outputDir string) error {
	fileName := filepath.Join(outputDir, App+"_errors_export.json")
	export := ExternalAll{
		ComponentName: componentInfo.Type,
		ComponentType: componentInfo.Name,
		Errors:        make(map[string]Error),
	}
	for k, v := range infoAll.LiteralCodes {
		if len(v) > 1 {
			continue
		}
		e := v[0]
		if _, err := strconv.Atoi(e.Code); err != nil {
			continue
		}

		export.Errors[k] = Error{
			Name:                 e.Name,
			Code:                 e.Code,
			Severity:             "",
			ShortDescription:     "",
			LongDescription:      "",
			ProbableCause:        "",
			SuggestedRemediation: "",
		}

		if _, ok := infoAll.Errors[e.Name]; ok {
			if len(infoAll.Errors[e.Name]) == 1 {
				details := infoAll.Errors[e.Name][0]
				export.Errors[k] = Error{
					Name:                 details.Name,
					Code:                 e.Code,
					Severity:             details.Severity,
					ShortDescription:     details.ShortDescription,
					LongDescription:      details.LongDescription,
					ProbableCause:        details.ProbableCause,
					SuggestedRemediation: details.SuggestedRemediation,
				}
			} else {
				slog.LogAttrs(slog.LevelError, fmt.Sprintf("duplicate error details for error name '%s' and code '%s'", e.Name, e.Code), slog.String("error", e.Code), slog.String("code", e.Code))
			}
		} else {
			slog.LogAttrs(slog.LevelWarn, fmt.Sprintf("no error details found for error name '%s' and code '%s'", e.Name, e.Code), slog.String("error", e.Code), slog.String("code", e.Code))
		}
	}
	jsn, err := json.MarshalIndent(export, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsn, 0600)
}
