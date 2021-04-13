run "02_read-scanner"

run "06_dial-write"

/* Remember this sever "02_read-scanner" when we had a browser client send connect to it, it never close. Now it does close because client "06_dial-write" defers the conncetion close and writes to the server "02_read-scanner" and exits program so the connection gets closed at "06_dial-write". 

Now the "02_read-scanner" receives what came in on the connection, prints out the text sent, and the connection get close for just that stream (that conn is closed and breaks out of that connection scanning) so we print "Code got here" But hoever, the for loop is still running and listening for new connection so that it can launch a new go routine to handle that connection */