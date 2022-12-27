package errorutil

import (
	"encoding/json"
	"fmt"
	"github.com/fenixvlabs/meshkit/pkg/meshlog"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slog"
	"os"
	"path/filepath"
	"strings"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func includeFile(path string) bool {
	if strings.HasSuffix(path, "_test.go") {
		return false
	}
	if filepath.Ext(path) == ".go" {
		return true
	}
	return false
}

func isErrorGoFile(path string) bool {
	_, file := filepath.Split(path)
	return file == "error.go"
}

func Walk(globalFlags GlobalFlags, update bool, updateAll bool, errorsInfo *InfoAll) error {
	logger := slog.New(&meshlog.LogrusHandler{
		Logger: log.StandardLogger(),
	})
	slog.SetDefault(logger)

	slog.Info("Analyzing tree directory...")

	subDirsToSkip := append([]string{".git", ".github"}, globalFlags.SkipDirs...)
	slog.Info(fmt.Sprintf("root directory: %s", globalFlags.RootDir))
	slog.Info(fmt.Sprintf("output directory: %s", globalFlags.OutDir))
	slog.Info(fmt.Sprintf("info directory: %s", globalFlags.InfoDir))
	slog.Info(fmt.Sprintf("subdirs to skip: %v", subDirsToSkip))
	comp, err := New(globalFlags.InfoDir)
	if err != nil {
		return err
	}

	err = filepath.Walk(globalFlags.RootDir, func(path string, info os.FileInfo, err error) error {
		slog.LogAttrs(slog.LevelInfo, "traversing directory", slog.String("path", path))
		if err != nil {
			slog.LogAttrs(slog.LevelWarn, "failure accessing path", slog.String("error", fmt.Sprintf("%v", err)))
			return err
		}
		if info.IsDir() && contains(subDirsToSkip, info.Name()) {
			slog.LogAttrs(slog.LevelInfo, "skipping directory", slog.String("path", info.Name()))
			return filepath.SkipDir
		}
		if info.IsDir() {
			slog.Debug("handling dir")
		} else {
			if includeFile(path) {
				isErrorsGoFile := isErrorGoFile(path)
				slog.LogAttrs(slog.LevelDebug, "failure accessing file", slog.String("iserrorsfile", fmt.Sprintf("%v", isErrorsGoFile)))
				err := handleFile(path, update && isErrorsGoFile, updateAll, errorsInfo, comp)
				if err != nil {
					return err
				}
			} else {
				slog.Debug("skipping file")
			}
		}
		return nil
	})
	if update {
		err = comp.Write()
	}
	return err
}

func ExportWalk(globalFlags GlobalFlags, update bool, updateAll bool) error {
	errorsInfo := NewInfoAll()
	err := Walk(globalFlags, update, updateAll, errorsInfo)
	if err != nil {
		return nil
	}
	if update {
		errorsInfo = NewInfoAll()
		err = Walk(globalFlags, false, false, errorsInfo)
		if err != nil {
			return err
		}
	}
	jsn, err := json.MarshalIndent(errorsInfo, "", " ")
	if err != nil {
		return err
	}
	fileName := filepath.Join(globalFlags.OutDir, App+"_analyze_errors.json")
	err = os.WriteFile(fileName, jsn, 0600)
	if err != nil {
		return err
	}
	componentInfo, err := New(globalFlags.InfoDir)
	if err != nil {
		return err
	}
	return Export(componentInfo, errorsInfo, globalFlags.OutDir)
}
