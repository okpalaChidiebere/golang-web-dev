Remember Go Lang encourages you to think about solving a problem programmatically instead of downloading a package and using it right away the way you do them in JavaScript or maybe other languages

# 002 to 013 - Templates
 we talked about Templates. We learned how we can make web templates the way we do them in React or other Front-end framework in Go lang this time

# 014 to 016 - TCP Servers
we talked about TCP Servers

- This concept of HTTP running on top TCP is really important.
- This means to create an HTTP server(Server which responds to HTTP requests), we have to create a TCP server that handles requests formatted in a certain way. This is because when a request comes in, the request is just Text, and the text is formatted in a certain way adhering to the Hypertext transfer Protocol then our server will be able to process that text.
- Protocol are rules or standard of communication. This Protocol are written in RFCs (7230 for this project). Remember that the RFC(Request For Comments) are documents put out by the IETF(Internet Engineering Task Force). These documents basically says here are the standards we adhere to when we do HTTP stuff!

The Net package in golang allows us to do things on the Network like create serves and do things on the TCP network

We will build on top of this, to build our own server to handle request(which is just a text). The request that comes in will be formatted to adhere to rfc7230

All the TCP code you wrote, you will never ideally have to write just to create our real server. This is for you to have a good knowledge understanding of HTTP when we start to set up our own HTTP servers except you want to start to get really custom

At this stage we fully understand how to make how HTTP works with the help of TCP. We built an entire server from tcp while understanding HTTP. This is why go lang is a good language. You code using go means you really understand what you are doing. 

# 017
We talked about understanding the net.http package

Main takeaways are

It is important to (Handler interface)[https://golang.org/pkg/net/http/#Handler]
Any custom type is also said to be a Handler type as well if it has a method with the signature (ServeHTTP(ResponseWriter, *Request)) attached to it. So that custom type can be said to implement the handler interface as well

ListenAndServer() method takes in a Handler. This ListenAndServer allow your server to be able to listen to HTTP request coming in from a certain port number. Port 80 for prod and 8080 for testing

- We understood that the ServerHTTP method listens for Request and writes back responses. We looked at the pointer to a http request (*http.Request) and the ResponseWriter argumements in the ServerHTTP method. 
- We understood that the pointer to a Request is a struct with a whole bunch of fields. We looked at fileds like Form and PostForm. We learned that ParseForm() methods need to be called to be able to access these two fields. We also learned abut the Method, URL, Header, Host, ContentLength fields as well
- We saw how to set different Headers for our response that we are writing back to the client as well

***********************
Next we will understand the type ServerMux. When you have a pointer to a mux, it is also of type Handler. Therefore we can pass the ServerMux to the ListenAndServer() method
ServerMux allows us to set our route and resposnses to those routes as well

# 018 - 022

Routing or Multiplexing

We will create a Multiplexer to respond to different request in different ways. 

We want to answer the question when people make request to our server at different path or url, how do we want to respond to those different urls and server different code for each of those urls?

We look at examples from the Least Elegant to the Best way to set up your Multiplexer

Three ways to make a multiplexer for your code
- by using a NewServeMux() method. This creates a mux that gives us a pointer to a ServeMux which has several methods attached to it. ServeHTTP is one of the methods. This means our mux is a Handler that we can pas to ListenAndServe(). Handle method is attached to it as well which allows us to pass in a route and Handler
- by using a DefaultServeMux where we pass in nil to the ListenAndserve. We use Handle(check out 020_HandleFunc) and HandleFunc (check out 020_HandleFunc) to attach routes to the DefaultServeMux. They are both http(net/http) package level methods
- by using third library packages. The third party library we explored was julius schmidt


https://pkg.go.dev/ is where you can search for standard packages or third party packages as well. The more popular third party packages will have more imports

FEW FYIs
- When you see http.ListenAndServe(":8080", nil) know that we are using the default multiplexer 
- When you see http.ListenAndServe(":8080", mux) know that someone created their own multiplexer. Probably using the NewServeMux() method like we have seen

****
The Entry point of understanding the net/http package is:
- type Handler which is an interface that has the ServeHTTP() method with a ResponseWriter and a Pointer to a Request (017)
- ListenAndServer takes a Handler (017)
- Next thing we want to do in our server is routing. The way we do Routing is with a Multiplexer(ServerMux). (018 - 022)

022 is more of Hands On Exercies. The Code in there dont have much comments. like the class lessons