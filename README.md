# A simple reverse proxy example with a golang server

### Gotchas:
1. Although Docker-compose allows your to orchestrate multiple containers, communication between the containers must be set. To set this we create a "network" in the Docker-compose file. Two containers in the same network can communicate with each other using the Docker-compose name.
2. Once linked, the network if the client will refer to the server not by 127.0.0.1/localhost, but by the Docker-compose container name, in my example it is "goserver".
3. Linked containers do not access each other through the host post, but by the internal container port. Hence the envoy proxy in this example forwards requests to goserver port 7000, instead of 8080.

### Other Useful tips
1. Container orchestration is done using Docker-compose, other similar ochestration tools are the Amazon ECS Task and Kubernetes Pod. They all have the same idea in that they for networks to connect containers together on the same host, they open and close ports on the host, they hey route host ports to container ports. 
   * Check out this [tool](https://github.com/micahhausler/container-transform) that converts your orchestration files to other forms.
