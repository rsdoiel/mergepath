/**
 * mergepath.go - merge the path variable to avoid duplicates
 *
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2014 All rights reserved.
 * @license released under the Simplified BSD License
 * see: http://opensource.org/licenses/bsd-license.php
 */
package main

import (
    "fmt"
    "os"
    "flag"
    "strings"
)

var (
    envPath string
    dir string
    help bool
    append_path = true
    prepend_path = false
)

var Usage = func(exit_code int, msg string) {
    var fh = os.Stderr

    if exit_code == 0 {
        fh = os.Stdout
    }
    fmt.Fprintf(fh, `%s
 USAGE %s [OPTIONS] PATH_TO_ADD PATH_TO_MODIFY

 EXAMPLES

    append work directory to existing path: %s . $PATH
    prepend working directory to existing path: %s -P . $PATH

 OPTIONS

`, msg, os.Args[0], os.Args[0], os.Args[0])
    flag.VisitAll(func (f *flag.Flag) {
        fmt.Fprintf(fh, "\t-%s\t(defaults to %s) %s\n", f.Name, f.DefValue, f.Usage)
    })
    fmt.Fprintf(fh, `

 Copyright (c) 2014 All rights reserved.
 Released under the Simplified BSD License
 See: http://opensource.org/licenses/bsd-license.php
`)
    os.Exit(exit_code)
}

func init() {
    const (
        path_usage = "The path you want to merge with."
        dir_usage = "The directory you want to add to the path."
        append_usage = "If directory is not in path, append the directory to the path"
        prepend_usage = "If directory is not in path, prepend the directory to the path"
        help_usage = "This help document."
    )
    envPath = os.Getenv("PATH")
    flag.StringVar(&envPath, "path", envPath, path_usage)
    flag.StringVar(&envPath, "p", envPath, path_usage)
    flag.StringVar(&dir, "directory", dir, dir_usage)
    flag.StringVar(&dir, "d", dir, dir_usage) 

    flag.BoolVar(&append_path, "append", append_path, append_usage)
    flag.BoolVar(&append_path, "a", append_path, append_usage)
    flag.BoolVar(&prepend_path, "prepend", prepend_path, prepend_usage)
    flag.BoolVar(&prepend_path, "P", prepend_path, prepend_usage)
    flag.BoolVar(&help, "help", false, help_usage)
    flag.BoolVar(&help, "h", false, help_usage)
}

func main() {
    flag.Parse()

    if help == true {
        Usage(0, "")
    }

    if flag.NArg() > 0 {
        dir = flag.Arg(0)
        if flag.NArg() == 2 {
            envPath = flag.Arg(1)
        }
    }

    if dir == "" {
        Usage(1, "Missing directory to add to path")
    }
    if prepend_path == true {
        append_path = false
    }
    if strings.Contains(envPath, dir) {
        fmt.Println(envPath)
        os.Exit(0)
    }
    if append_path == true {
        fmt.Printf("%s:%s", envPath, dir)
    } else {
        fmt.Printf("%s:%s", dir, envPath)
    }
}
