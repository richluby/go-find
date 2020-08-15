package main

import (
	"path/filepath"
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

func filter_path(path os.FileInfo, full_path string, filters []*filter_t) {
	print_item := true;
	for _, filter := range filters {
		do_retain, err := filter.Should_retain(path) 
		if err != nil {
			print_item = false
			fmt.Fprintf(os.Stderr, "Error occurred filtering [%s] : %+v\n", full_path, err)
			break
		}
		if !do_retain {
			print_item = false
			break
		}
	}
	if print_item {
		fmt.Println(full_path)
	}
}

func add_to_channel(channel chan int) {
	channel <- 1
}

func handle_path(full_path string, filters []*filter_t, done chan int) {
	defer add_to_channel(done)
	info, err := os.Lstat(full_path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error statting [%s] : %+v\n", full_path, err)
		return
	}

	if info.IsDir() {
		dir_done := make (chan int)
		open_file, err := os.OpenFile(full_path, os.O_RDONLY, 0700)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening [%s] : %+v\n", full_path, open_file)
		}
		defer open_file.Close()
		names, err := open_file.Readdirnames(-1)
		for _, name := range names {
			go handle_path(filepath.Join(full_path, name), filters, dir_done)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error occurred : %+v\n", err)
		}
		for range names {
			<- dir_done
		}
	}

	filter_path(info, full_path, filters)
}

func perform_search(filters []*filter_t) {
	files := flag.Args() // skip program name
	done := make (chan int)
	for _, f := range files {
		full_path, err := filepath.Abs(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed creating absolute path [%s] : %+v\n", f, err)
			done <- 1
		} else {
			go handle_path(full_path, filters, done)
		}
	}

	for range files {
		<- done
	}
}

func main () {
	init_filters()
	init_flags()
	filters := get_active_filters()
	perform_search(filters)
	os.Exit(0)
}