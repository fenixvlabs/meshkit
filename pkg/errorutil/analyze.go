package errorutil

import (
	"bytes"
	"github.com/fenixvlabs/meshkit/pkg/meshlog"
	log "github.com/sirupsen/logrus"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"golang.org/x/exp/slog"
	"os"
)

func handleFile(path string, update, updateAll bool, infoAll *InfoAll, comp *Info) error {
	logger := slog.New(&meshlog.LogrusHandler{
		Logger: log.StandardLogger(),
	})

	slog.SetDefault(logger)

	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	anyValueChanged := false
	ast.Inspect(file, func(n ast.Node) bool {
		if pgkid, ok := isNewDefaultCallExpr(n); ok {
			slog.Warn("Usage of deprecated function %s.NewDefault detected.", pgkid)
			if !contains(infoAll.DeprecatedNewDefault, path) {
				infoAll.DeprecatedNewDefault = append(infoAll.DeprecatedNewDefault, path)
			}
			// If a NewDefault call expression is detected, child-nodes are not inspected.
			// This would lead to duplicates detections in case of dot-import.
			return false
		}
		if newErr, ok := isNewCallExpr(n); ok {
			name := newErr.Name
			logger.Info("New.Error(...) call detected, error code name", name)
			_, ok := infoAll.Errors[name]
			if !ok {
				infoAll.Errors[name] = []Error{}
			}
			infoAll.Errors[name] = append(infoAll.Errors[name], *newErr)
			// If a New call expression is detected, child-nodes are not inspected:
			return false
		}
		if handleValueSpec(n, update, updateAll, comp, logger, path, infoAll) {
			anyValueChanged = true
		}
		return true
	})
	if update && anyValueChanged {
		logger.Info("writing updated file")
		buf := new(bytes.Buffer)
		err = format.Node(buf, fileSet, file)
		if err != nil {
			return err
		}
		err = os.WriteFile(path, buf.Bytes(), 0600)
		if err != nil {
			return err
		}
	}
	return nil
}
