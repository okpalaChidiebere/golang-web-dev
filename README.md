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

# 023 - 025

We saw some other functions that we can use to server files like http.ServeContent, http.ServeFile, io.Copy. These are just for so that i am aware of them and what they do. But most of the time for real projects you may not need them.

In practise, for most of the time, we use http.FileServer as handler for serving files(espacially if its a static website) or http.StripPrefix as handler for serving files stored in a particular directory espacially for exposing those files ou are serving from a single route in your server

We learned about http.FileServer which allows us to serve a whole directory and also stripping the prefix to take care of routing stuff with the help of http.StripPrefix. So we can specify a certain directory and have certain control over what is asked for and where we serve it from. This is important becasie we want to be able to just serve the assets. We don't want our go code to be served up like the FileServer does. 
Inside the assets we will have css, javascript, images. In addition you may have templates which will be in a folder called templates

We learned how to serve a static website. The idea here is we can take a static website make with react, vue, etc and put it up somewhere and it is running with go. 

We saw how log.fatal & http.Error works

We learned the http.notFoundHandler for handling favicon.ico requests

# 027 - 030

We saw two ways of managing state for your application
- One to maintain state is to pass a unique id value each time to the url of a page so that we can know who the user is each time. You probably will go with this option, if the user dont allow you to write cookies to their browser
- Another way is to use cookies and sessions

State is a persistent awareness of who is communicating with the server. For example, when someone logs in, the server knows who is logged and know if they have access to certain information or not.

Sessions are basically where someone can log in, do stuff with authorized access credentials to certain areas, and then they can finally logout.  

We learned AGAIN how to pass form data values through the url or request body

We will learn how to write cookie(s) to the client's browser and also how to read cookies that are sent from the client to the server.

While learning cookies, we understood how redirects works as well
We the client sends a request to the server to get a resource at a location though a URL, the server can choose to redirect that request to another location. Maybe because, that resource is moved to another location or maybe the server does some processing at that location and then redirects the request to another location (like you send a POST request on form submit, it gets processed in a location and then the user will not be send to another location to show the success page)

** How session works **
Clients makes request to the server, and we want the client to send a unique identifier (sessionID) to that server. In our server, we could associated that SessionID to a UserID (we have a table that stores a sessionID for every userID). Then we can now take the userID and get anything we want about the user. We could store that Unique ID(sessionID) in a cookie and knowing that cookie are domain specific, whenever the client send a request to our server, if there is a our cookie on their browser, they will send that cookie as well. Then our server can grab that value and we can uniquely identify them. 

The SessionID are stored and deleted as the user logs in and out from our server. But the UserID and other information is always there in our DB table

To Logout the user we delete/expire their cookie and then delete their session entry from the session table!

To add permissions in our server, we just added an extra Role attribute to a user struct. Then in any of the functions that need to run only based on role, we check if the user role is equal to some certain value

Expiring a session
This means after some amount certain time of inactivity, you have not clicked on any links for a certain amount of time within the website, then it will log you out automatically
- One way to expire a session is to set the maxAge on the cookie to a certain value in seconds. Eg We basically set 10 minutes on their cookie, this mean they have just 10 minutes to click on a link, otherwise we log them out
- Another way is to keep track of each session on the server(db) and when was the last time of each activity. So occasionally we will go through the table and clean up the sessions that has not been active in the past 10 mins

IN ADDITION ( 030_session/09_middleware)
You can check out how Middleware works on our server. We can use middleware to check for authorization on an enpoint before it gets executed