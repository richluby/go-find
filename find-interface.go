package main

import "os"

type filter_func_t func(os.FileInfo, string) (bool, error)

type filter_t struct {
	Current_filter string
	name string
	help string
	default_value string
	filter filter_func_t
}

type filter_i interface {
	Should_retain(filepath os.FileInfo) (bool, error)
	Name() string
	Help() string
	Default() string
}

func (f *filter_t) Should_retain(filepath os.FileInfo) (bool, error) {
	return f.filter(filepath, f.Current_filter)
}

func (f *filter_t) Name() string {
	return f.name
}

func (f *filter_t) Help() string {
	return f.help
}
func (f *filter_t) Default() string {
	return f.default_value
}

func New_Filter(name string,
	filter_function filter_func_t,
	default_value string,
	help string) filter_i {
	return &filter_t{ 
		name : name,
		filter : filter_function,
		help : help,
		default_value : default_value }
}