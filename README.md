mergepath
=========

A small utility to merge PATH lists across different platforms (e.g. Linux, OS X).

     USAGE mergepath [OPTIONS] PATH_TO_ADD PATH_TO_MODIFY

## EXAMPLES

    append work directory to existing path: mergepath . $PATH
    prepend working directory to existing path: mergepath -P . $PATH

## OPTIONS

+ *-a*, *--append* (defaults to true) If directory is not in path, append the directory to the path
+ *-d*, *--directory* The directory you want to add to the path.
+ *-h*, *-help* (defaults to false) This help document.
+ *-p*, *--path* (defaults to $PATH) The path you want to merge with.
+ *-P*, *--prepend* (defaults to false) If directory is not in path, prepend the directory to the path

## Installation

_mergepath_ can be installed with the *go get* command.

```
    go get github.com/rsdoiel/mergepath/...
```

Copyright (c) 2014 All rights reserved.
Released under the [Simplified BSD License](http://opensource.org/licenses/bsd-license.php)

