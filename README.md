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


# 031 - 032

At 031 we learned how to connect our Go works with SQL. We provisioned a RDS(MYSQL server) at AWS, made it open to public. We had to open the the DB to public so that we can connect to the DB in our code locally. This also made it easy for us to connect to the DB through the MYSQL work bench from just standard TCP/IP connection which was ok just for testing. For your real company Prod or Dev environment, you dont your database that public



032 we learned more about AWS. 
- We learned how to create instances and how to assign security group configuration to it to make them not accessible to the public as well as and how to use key pairs
- We learned how to create a machine image from an existing instance. The existing machine is usually the good instance we created ourself that we are sure is working fine.
- We saw how to use security groups for EC2s and RDS machines
- We learned how to use create Load balancer with target groups for that load balancer
- We learned Auto scaling. Auto scaling basically helps us scale horizontally. We set rules like if our current server is using 75% percent of its capacity spin up more servers to help with load or if the server is using 50% less of its capacity, then start to drop servers because there is no work!
- We learned how to use CloudFront on Ec2s by selecting the loadbalancer "origin domain name" if you have your Ec2 to be only accessed through the loadbalancer in the CF console. Basically with CloudFront, you server will be available in every region! You can confirm this works because when you go to the CloudFront console, the button where you select your region will be disabled because cloudFront does not require region selection. So when someone makes a request, it answers that request with the server closest to them. 
- Lastly know about AWS Rout53 for domain names!

** 033_aws-scaling/02_load-balacer **
At the end of this demo we had our loadbalancer up and running which is connected to our EC2 instance(webserver) and you can only get to our web server through the load balancer. More detals on the readme

** 033_aws-scaling/03_ami **
- We created a machine image of our server and uses the image to create another instance. we now got both the EC2 server and the EC2 server that was from the image running as well. Then we observed the Load balancer switching between the two servers.
- We learned how to connect to our RDS differently from our MySQL workbench now that we have some restrictions configured to our RDS machine. The overall idea here is knowing that any machine that has the "web-tier" security group applied to it can communicate with other machines with thesame security group as well usingTCP 3306 port; we will connect through SSH to our EC2 instance, then our instance we are able to connect to the MySQL database

** 033_aws-scaling/04_hands-on **
From the hands-on exercise(033_aws-scaling/04_hands-on), we reployed a new code to our instance have some database operations in it. Few things i learned from redeploying a new code version to your instance that has the old code.
- First build the new binary for the new code version locally by running `GOOS=linux GOARCH=amd64 go build -o [some-name]`
- Connect into the remove server using ssh
- Go the the systemd service using cd /etc/systemd/system/
- Stop the current running service for your code using `sudo systemd stop [filename].service`
- Go back to the root directory of your old server by running `cd`
- Then delete the old binary using the `sudo rm [some-name]`. NOTE if the binary is not at the root folder be sure to add the path to the folder in the remove command. - Exit our of the remote server
- Copy the new binary from local to the remove server using the `scp` command
- The connect back to the remove server to start back up the service we stopped. But first will have to modify the service file IF only the new binary file has another name. Modify the "ExecStart" value in the service file
- Then reload the systemctl by running `sudo systemctl daemon-reload`
- Run `sudo systemctl start [some-name].service` to start back up the file name
- You can check the status to confirm it is running if you want

Another way to have your new code run is you can remove the systemd service for the old code and create a new one but i did not have luck with that. I had to modify the existing service.
It is important anytime you want to make a change in the systemd to stop the service, make the change, reload the systemd and then start the service back up again. DO NOT make the change when your service is running!!.

There can be a few reasons why an endpoint from your remote server you invoke will cause an error like 502 BadGateway.
- Maybe the environmental variables are not set yet and so your code will have a runtime error when the value it needs for a configuration from its environment is not there. See more on how to set environment variables on systemd service in the link below. You can even go as far as specifying a script that will run your executable. (https://askubuntu.com/questions/1063153/systemd-service-working-directory-not-change-the-directory)[https://askubuntu.com/questions/1063153/systemd-service-working-directory-not-change-the-directory]
- Or maybe there is a bug in your code overall. So make sure the code is working locally without errors before you deploy. NoTE: after you build your executable locally you don't have to move the go.mod file to the remote server. Just be sure to move the folder paths of template files if needed by your code. I will look into on how to deploy code with hexagonal architecture to remove server

# 034

We made an app that will store the names of images that users have uploaded in the cookie. So when we are getting a cookie for the user, we pull out the images names and serve those images to the user

- at /02_cookie, We first got a cookie up and running for our site. If the user has a cookie we get the cookie and display it. If the client don't have a cookie, we create a new one for the client
- at /03/store-values, We added three images names to the cookie separated by pipes. We only add the images names if they are not already in there. Right now the image names are just mock names but down the line the will be real images names that the user uploaded from the browser.
- at /04_upload-pictures, We allowed a user to upload an image, take the image and store it on the server and puts its name of the stored image in the cookie. So if a user have image names in their cookie value, we display the image on our server to the client page. NOTE: We stored the SHA1 of the image as image name. SHA is an algorithm where you take something of a large size, run it through some algorithm and get some unique output which will always be thesame if the input(can be a file) is thesame. This is good because if multiple users upload thesame image, you will store the image just once and then delete it once nobody need the image anymore.
- at /05_display-pictures, We saw how to display the images or serve files. The images will not display anymore if the client deletes their cookie
