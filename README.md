# Social Media Backend Example

## About
An example social media backend using AWS. This project demonstrates the use of an API using a ECS cluster and RDS
backend. The API is publicly accessible using a custom domain configuration for API gateway. The API
gateway uses a VPC link to reach a services within a private VPC across multiple subnets.
Traffic is sent to an application load balancer then to ECS. ECS uses Secrets Manager to obtain credentials to connect to RDS.
  The project is designed for use with a hybrid of local and cloud development.

![sample image](./diagram.drawio.png)


## Usage

### Requirements
- Setup credentials for AWS CLI
- Create a file `./terraform/terraform.tfvars` and populate it using variables described in `./terraform/variables.tf`

### Local Development
The ECS image can be developed and tested locally using a local PostgreSQL database using docker compose.

- Run `docker compose build` to build image for the ECS container.
- Start up docker compose with `docker compose up -d` to run local PostgreSQL and ECS container.

### Cloud Development

Use infrastructure as code to create AWS services.

Initialize terraform.

`terraform init`

Create infrastructure.

`terraform apply`

Build and tag ECS image with name of your ECR repo url that was output from terraform.

`docker build -t <ecr-repo-url>:latest .`

Authenticate docker

`aws ecr get-login-password --region region | docker login --username AWS --password-stdin aws_account_id.dkr.ecr.region.amazonaws.com`

Push to ECR

`docker push <ecr-repo-url>:latest`

## API Usage

### Users

<details>
<summary>
<code>POST</code> /users
</summary>

Create a new user.

</details>



<details>
<summary>
<code>GET</code> /users/{:id}
</summary>

Get a user

</details>

<details>
<summary>
<code>PUT</code> /users/{:id}
</summary>

Update a user

</details>

<details>
<summary>
<code>DELETE</code> /users/{:id}
</summary>

Delete a user

</details>

### Posts

<details>
<summary>
<code>GET</code> /posts
</summary>

Get posts

</details>

<details>
<summary>
<code>POST</code> /posts
</summary>

Create a new post

</details>

<details>
<summary>
<code>PUT</code> /posts/{:post-id}
</summary>

Update a post

</details>

<details>
<summary>
<code>DELETE</code> /posts/{:post-id}
</summary>

Delete a post

</details>

### Comments

<details>
<summary>
<code>GET</code> /comments
</summary>

Get comments from a post with pagination

</details>

<details>
<summary>
<code>POST</code> /comments
</summary>

Create a new comment on a post

</details>

<details>
<summary>
<code>PUT</code> /comments/{:comment-id}
</summary>

Update a comment on a post

</details>

<details>
<summary>
<code>DELETE</code> /comments/{:comment-id}
</summary>

Delete a comment on a post

</details>

### Follower

<details>
<summary>
<code>GET</code> /follower
</summary>

Get follwers with pagination

</details>

<details>
<summary>
<code>POST</code> /follower
</summary>

Create a follower

</details>

<details>
<summary>
<code>DELETE</code> /follower
</summary>

Delete a following relationship

</details>


### NewsFeed

<details>
<summary>
<code>Get</code> /feed
</summary>

Get newsfeed for user with pagination

</details>

