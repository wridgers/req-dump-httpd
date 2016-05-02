# req-dump-httpd

A simple httpd for debug/testing purposes. Responds to every request by writing its HTTP/1.x wire representation (including the body). The port to listen on and an optional message can be set.

## Install

    $ go install github.com/wridgers/req-dump-httpd

## Usage

    $ req-dump-httpd -h
    Usage of req-dump-httpd.exe:
      -msg string
            the message to respond with
      -port int
            port on which to bind (default 8080)

## Example

    $ req-dump-httpd -msg="Hello world."

Visit [http://localhost:8080](http://localhost:8080), get something like this:

    Message: Hello world.

    GET / HTTP/1.1
    Host: localhost:8080
    Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
    Accept-Encoding: gzip, deflate
    Accept-Language: en-US,en;q=0.5
    Cache-Control: max-age=0
    Connection: keep-alive
    User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64; rv:45.0) Gecko/20100101 Firefox/45.0
