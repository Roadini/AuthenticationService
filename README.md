# Authentication Service and Social Services
This project can have two sub-classifications:
### Authentication Service
The Authentication Service is responsible for providing the user with an authentication token. This authentication token is, later on, used assure that only an authenticated user can access its contents. 
The authentication token is a JWT signed with a private key that only this server as knowledge of. Through the use of this policy, the server can guarantee that it was responsible for that signature and, therefore, guarantee the authenticity of the user. The endpoints provided by this service allow for:  
* User Creation;
* User Search;
* User Removal;
* User Login;
* User Logout;
* User Information Update.

#### Facebook
This service integrates Facebook Authentication System. Facebook retrieves a token to our app that, along with an app private token, allows information retrieval from the Facebook system.
 
### Social Services
The classic Social networks (Facebook, Instagram ...) have key features that allow users to interact with each other and provide a more sociable environment.
Likewise, the social component of this service provides the following features:
* Follow Users:
    * User Follow/Unfollow;
    * Following Allowence;
    * Followers/Following retreival;
* Publications:
    * Create/Remove Posts;
    * User Posts Retrieval;
* Likes:
    * Like/Unlike Posts;
    * Post Like Retrieval;

### Requirements
* Docker
* Docker-compose

### Deployment
#### Key Creation 
In order to deploy the project, we must first generate a key pair.
``` bash
# On the project base dir
mkdir -p src/keys
cd src/keys
ssh-keygen -t rsa -b 4096 -f sample_key.priv
# Don't add passphrase
openssl rsa -in sample_key.priv -pubout -outform PEM -out sample_key.pub
```
#### Service Deploy
Server deployment
```bash
# On the project base dir
cd src
docker build . --tag auth_api
cd ../docker
docker-compose up
```
