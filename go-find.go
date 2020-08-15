package main

import (
	"flag"
	"fmt"
	"container/list"
	"os"
)

var filters_g * list.List

func init_flags() {
	for element := filters_g.Front() ; element != nil; element = element.Next() {
		var e * filter_t = element.Value.(*filter_t)
		flag.StringVar(&e.Current_filter, e.Name(), e.Default(), e.Help())
	}
	flag.Parse()
}

func init_filters() {
	filters_g = list.New()
	filters_g.PushBack(New_Filter("type", filter_by_type, "", "the file type to match [d,f]"))
	filters_g.PushBack(New_Filter("name", filter_by_name, "", "the file name to match"))
	filters_g.PushBack(New_Filter("mtime", filter_by_mtime, "", "the modified file time to match (file time must be after this value)"))
	filters_g.PushBack(New_Filter("iname", filter_by_iname, "", "the case-insensitive file name to match"))
}

func get_active_filters() []*filter_t {
	var filters []*filter_t
	for element := filters_g.Front(); element != nil; element = element.Next() {
		e := element.Value.(*filter_t)
		if "" != e.Current_filter {
			filters = append(filters, e)
		}
	}
	return filters
}

func filter_path(path os.FileInfo, filters []*filter_t) {

}

func handle_path(path string, filters []*filter_t) {
	info, err := os.Lstat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error statting [%s] : %+v", path, err)
		return
	}

	if info.IsDir() {
		open_file, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening [%s] : %+v", path, open_file)
		}
		defer open_file.Close()
		names, err := open_file.Readdirnames(-1)
		for _, name := range names {
			go handle_path(name, filters)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error occurred : %+v", err)
		}
	}

	filter_path(info, filters)
}

func perform_search(filters []*filter_t) {
	files := flag.Args() // skip program name
	for _, f := range files {
		go handle_path(f, filters)
	}
}

func main () {
	init_filters()
	init_flags()
	filters := get_active_filters()
	perform_search(filters)
	os.Exit(0)
}