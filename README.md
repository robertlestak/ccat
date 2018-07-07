# ccat

ccat (curl-cat) `cat`s the contents of the provided file(s) to `stdout`, and if the provided path is a remote URL, it HTTP `GET`s the file contents.

## Build

````
go build .
cp ccat /usr/local/bin/ccat
````

## Usage

````
ccat [options] [file/url ...]
  -H value
    	Headers to send when requesting remote resources
  -b string
    	Body to send when requesting remote resources
  -c value
    	Cookies to send when requesting remote resources
  -m string
    	Method when requesting remote resources (default "GET")
  -u string
    	Basic auth to send when requesting remote resources
````

### Example

````
ccat file.txt another_local_file.txt https://example.com/file.txt
````

## A Note on the Name

At this point, "curl-cat" is probably a bit of a stretch. This script does *not* have all the power that `curl` does - currently, it can simply `GET` a remote resource and that's it.

The name was admittedly chosen because it rolled off the toung well, gave a broad idea of the usage, and was "cute". Sorry. I aim to add more features to the script to make it more "curl-like". Pull requests are always welcome.
