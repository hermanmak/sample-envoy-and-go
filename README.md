This is a sample stack using:
envoy as reverse proxy
go as service code

Gotchas:
1) Although Docker-compose allows your to orchestrate multiple containers, communication between the containers must be set. To set this we create a "network" in the Docker-compose file. Two containers in the same network can communicate with each other using the Docker-compose name.
2) Once linked, the network if the client will refer to the server not by 127.0.0.1/localhost, but by the Docker-compose container name, in my example it is "goserver".
3) Linked containers do not access each other through the host post, but by the internal container port. Hence the envoy proxy in this example forwards requests to goserver port 7000, instead of 8080.
