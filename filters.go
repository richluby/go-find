package main

import (
	"fmt"
	"os"
)

func filter_by_name(filepath os.FileInfo, filter string) (bool, error) {
	return filepath.Name() == filter, nil
}
func filter_by_type(filepath os.FileInfo, filter string) (bool, error) {
	switch filter {
	case "d":
		return filepath.IsDir(), nil
	case "f":
		return !filepath.IsDir(), nil
	default:
		return false, fmt.Errorf("Unrecognized file type : %s", filter)
	}
	return filepath.IsDir() && filter == "d", nil
}
func filter_by_iname(filepath os.FileInfo, filter string) (bool, error) {
	return true, nil
}
func filter_by_mtime(filepath os.FileInfo, filter string) (bool, error) {
	return true, nil
}
