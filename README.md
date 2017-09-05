# Witch

[![GoDoc](https://godoc.org/github.com/kusti8/witch?status.svg)](https://godoc.org/github.com/kusti8/witch)

A super small wrapper around JSON to store settings in the correct location on all platforms

## Install

`go get github.com/kusti8/witch`

## Usage

First have a Config struct with all the expected values:

```
type Config struct {
    Test int
    Ing string
}
```

And then, create a struct with all the default values in it:

```
defaults := Config{1, "hello"}
```

And then either read or write, **using a pointer**:

```
witch.Write(&defaults, "kusti8", "test") // write the configuration to a file

witch.Read(&defaults, "kusti8", "test") // read the configuration, but any missing values are replaced by the defaults
```
The second and third arguments are the vendor and the name