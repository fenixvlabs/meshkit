package errorutil

import (
	"encoding/json"
	"golang.org/x/exp/slog"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func New(dir string) (*Info, error) {
	info := Info{file: filepath.Join(dir, filename)}
	slog.Debug("reading", slog.String("file", info.file))
	file, err := os.ReadFile(info.file)
	if err != nil {
		return &info, err
	}

	err = json.Unmarshal([]byte(file), &info)
	return &info, err
}

// Write writes the component info back to file.
func (i *Info) Write() error {
	jsn, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return err
	}
	slog.Debug("writing", slog.String("file", i.file))
	return os.WriteFile(i.file, jsn, 0600)
}

// GetNextErrorCode returns the next error code (an int) as a string, and increments to the next error code.
func (i *Info) GetNextErrorCode() string {
	s := strconv.Itoa(i.NextErrorCode)
	i.NextErrorCode = i.NextErrorCode + 1
	return s
}

func isErrorCodeVarName(name string) bool {
	matched, _ := regexp.MatchString("^Err[A-Z].+Code$", name)
	return matched
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
