# Arxor

XOR armor for tcp traffic.
It's a proof-of-concept program for breaking traffic analysis.

## Warning

**This is not encryption.** It's easy to break this armor.
Use SSH port forwarding for better security.

## Usage

    $ arxor <listenAddr> <dialAddr>

This program will send all traffic coming from `listenAddr`
to `dialAddr` with the bytes XOR'ed.

Typical configuration:

           clear text         XOR'ed traffic         clear text
    client<---------->arxor<-----Internet----->arxor<---------->server
    `----------v----------'                    `----------v----------'
        client machine                             server machine

## Example

Suppose you have your HTTP proxy running on port 12345 on your server.
Run

    $ arxor :12346 :12345

on this server, then on the client machine, run

    $ arxor :12345 server:12346

Then point your browser's HTTP proxy to `localhost:12345`,
so that HTTP proxy traffic will be sent over the Internet with its bytes
XOR'ed, and third-party may not recognize it without specialized analysis.

## Hack

You can alter `Transformer.Transform` method to use a more sophiscated
XOR pattern (e.g., XOR with a pseudo-random sequence).
