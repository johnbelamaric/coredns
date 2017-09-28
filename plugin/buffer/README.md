# buffer

*buffer* sets the size of the UDP receive buffer

By default, CoreDNS will not set the UDP receive buffer size, but let the kernel choose. The
*buffer* directive allows you to manually set this for the UDP socket. If this is set in
multiple server blocks that share a port number, the largest value will be used.

## Syntax

~~~ txt
buffer BYTES
~~~

**BYTES** is the number of bytes to use for the buffer

## Examples

To make your socket one megabyte:

~~~
. {
    buffer 1048576
}
~~~
