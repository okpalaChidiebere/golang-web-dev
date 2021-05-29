# Deploying our session example

Before i deploy, i had to make sure that my code was working just fine locally. I had to test the (code)[https://github.com/okpalaChidiebere/golang-web-dev/tree/master/030_sessions/08_expire-session] by initializing the project using `go mod init expire-session.com/m` then download the third party packages used by our app using `go mod tidy`

1. change your port number from 8080 to 80

1. create your binary
  - `GOOS=linux GOARCH=amd64 go build -o [some-name]`

1. Launch an EC2 server Instance.
  - This is an (eight-step process)[https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EC2_GetStarted.html#ec2-launch-instance] 
  - You can create a keyFile at the last step of the launching an instacne or use an existing one(If you have one already). Either way a good practise is to have the keyfile in a hidden directory like .ssh/ So we can move the key file there like `mv ~/Downloads/[your-AWS].pem ~/.ssh/` then change the file to only be executable using  `chmod 400 ~/.ssh/[your-AWS].pem`

1. SSH into your server
  - ssh -i /path/to/[your].pem ubuntu@[public-DNS]:
  - I did `ssh -i ~/.ssh/[your-AWS].pem ec2-user@[public-DNS]:`  The colon at the end means connect to th the root directory (~/). You can get your public IP address from the aws console!

1. create directories to hold your code
  - for example, "wildwest" & "wildwest/templates"
  - In my case, i created `src` & `/src/templates`
 
1. copy binary to the server
  - I used `scp -i ~/.ssh/[your-AWS].pem -r [some-name] ec2-user@[IPV4-public-DNS]:src`  NOTE: we are copying into our remote src folder. If we did not specify the src and just have the colon we will be copying just to the root folder

1. copy your "templates" to the server
  - scp -i ~/.ssh/[your].pem templates/* ubuntu@[public-DNS]:/home/ubuntu/templates
  - I used `scp -i /path/to/[your].pem -r templates/* ubuntu@[IPV4-public-DNS]:src/templates` NOTE: Now we are copying our /templates locally to /src/templates that is remote

1. chmod permissions on your binary
  - i used `sudo chmod 700 [some-name]`  NOTE: this permisson will only allow our instance to read, write or execute the code.  -rwx------

1. Run your code
  - `sudo ./[some-name]`
  - check it in a browser at [public-IPV4-address]. Give it some time; about 5 mins. NOTE: You can see the logs your code wrote as well in the shell. You can use that for debugging! Very Cool!

# Persisting your application using systemd

  To run our application after the terminal session has ended, we must do the following:

  FYI: i named my executable in the aws remote instance as expiresession. see that at /030_sessions/08_expire-session

  1. Create a configuration file
    - first `cd /etc/systemd/system/` from the root directory (~/); then
    - `sudo nano /etc/systemd/system/[filename].service` I used  then copy code below into the service file

  ```
  [Unit]
  Description=Go Server

  [Service]
  ExecStart=/home/<username>/<path-to-exe>/<exe>
  WorkingDirectory=/home/<username>/<exe-working-dir>
  User=root
  Group=root
  Restart=always

  [Install]
  WantedBy=multi-user.target
  ```

  NOTE: i used 
  ExecStart=/home/ec2-user/src/expiresession    
  WorkingDirectory=/home/ec2-user/src

  1. Add the service to systemd.
    - `sudo systemctl enable [filename].service`
  1. Activate the service.
    - `sudo systemctl start [filename].service`
  1. Check if systemd started it.
    - `sudo systemctl status [filename].service`
  1. Stop systemd if so desired.
    - `sudo systemctl stop [filename].service`








# FOR EXAMPLE
  ```
  [Unit]
  Description=Go Server

  [Service]
  ExecStart=/home/ec2-user/src
  WorkingDirectory=/home/ec2-user
  User=root
  Group=root
  Restart=always

  [Install]
  WantedBy=multi-user.target
```

Things to note:
usename for your ec2 instance is always ec2-user by default. You can check it out by clicking the "Connect" button when you click into the instance in the console. You add your won custom name though and then you can use that custom name to connect using secure shell(SSH) for mac or linux users. You also connect from the console as well. Our teacher used ubuntu because used ubuntu as platform for his machine. Why on that (here)[https://stackoverflow.com/questions/33991816/ec2-ssh-permission-denied-publickey-gssapi-keyex-gssapi-with-mic]
Few things about using Nano to edit files (here)[https://serverpilot.io/docs/how-to-use-nano-to-edit-files/#:~:text=You%20can%20save%20the%20file,different%20filename%20and%20press%20ENTER.]








