# Create security groups

A security group acts as a virtual *firewall* that controls the traffic for one or more instances.

When you launch an instance, you associate one or more security groups with the instance.

You add rules to each security group that allow traffic to or from its associated instances.

You can modify the rules for a security group at any time; the new rules are automatically applied to all instances that are associated with the security group. When we decide whether to allow traffic to reach an instance, we evaluate all the rules from all the security groups that are associated with the instance.

## ELB (load balancer) security group (we called it loadbalancer-sg)
1. add this rule
  - HTTP TCP 80 Anywhere
1. copy *Group ID*
  
## Web tier security group (we called it webtier-sg)
1. add these rules
  - HTTP TCP 80 Custom IP ```<load-balancer-sg Group ID>```
  - SSH TCP 22 My IP
1. copy **Group ID**
1. add this rule
  - MySql TCP 3306 Custom IP ```<web-servers-sg Group ID>```

# Load balancer
1. EC2 / Load balancers / Create load balancer
  - application or classic
  - name: web-elb
  - http & https
  - default VPC
  - add two subnets
1. configure security groups
  - choose "load-balancer" security group which we just setup
1. configure routing
  - target group: web-servers-tg1
  - ping path: /ping
  - allows us to define a "healthy" web server
  - load balancer will only forward to healthy web servers
1. register targets
1. create

We created different types of security group for the different kinds of machine that we will be running.

Firstly we created a Load Balancer Security Group. The load balancer security group(called loadbalancer-sg) will accept request(HTTP traffic) from anywhere on TCP Port 80. Eg  this Load balancer can accept SSH traffic from us, client's web browser, mobile, etc

Secondly, we created another security group(we called it webtier-sg) for database and EC2. We want this to accept traffic if it is coming from the load balancer on TCP port 80. This "web-tier" security group we will set it on the EC2 machine and the Database machine. Internally the machines that has the web-tier security group can communicate with each other (eg the EC2 and database can interact with each other because we applied the web-tier security group to them) but cannot be accessed publicly except from the load balancer. If you try to load the public Ip of the EC2 machine, you will just get the request loading forever! The DNS name of the LoadBalancer is what the client can use to send requests

At the end of this demo we had our loadbalancer up and running which is connected to our EC2 instance(webserver) and you can only get to our web server through the load balancer