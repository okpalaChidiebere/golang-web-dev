Remember Go Lang encourages you to think about solving a problem programmatically instead of downloading a package and using it right away the way you do them in JavaScript or maybe other languages

# 002 to 013 - Templates
 we talked about Templates. We learned how we can make web templates the way we do them in React or other Front-end framework in Go lang this time

# 014 to 041 - TCP Servers
we talked about TCP Servers

- This concept of HTTP running on top TCP is really important.
- This means to create an HTTP server(Server which responds to HTTP requests), we have to create a TCP server that handles requests formatted in a certain way. This is because when a request comes in, the request is just Text, and the text is formatted in a certain way adhering to the Hypertext transfer Protocol then our server will be able to process that text.
- Protocol are rules or standard of communication. This Protocol are written in RFCs (7230 for this project). Remember that the RFC(Request For Comments) are documents put out by the IETF(Internet Engineering Task Force). These documents basically says here are the standards we adhere to when we do HTTP stuff!

The Net package in golang allows us to do things on the Network like create serves and do things on the TCP network

We will build on top of this, to build our own server to handle request(which is just a text). The request that comes in will be formatted to adhere to rfc7230

All the TCP code you wrote, you will never ideally have to write just to create our real server. This is for you to have a good knowledge understanding of HTTP when we start to set up our own HTTP servers except you want to start to get really custom

At this stage we fully understand how to make how HTTP works with the help of TCP. We built an entire server from tcp while understanding HTTP. This is why go lang is a good language. You code using go means you really understand what you are doing. 

