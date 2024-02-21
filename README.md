# README

E-commerce template for an alien widget factory called `e-commerce`.

## Pages

- Home
- About
- Products
- Product Detail
- Cart
- Checkout
- Order Confirmation
- Account
- Login
- Register
- Forgot Password
- Reset Password
- Support
- 404

## Technologies

|  | Technology |
| --- | --- |
| Frontend | HTMX |
| Backend | Go |
| Database | SQLite |
| Authentication | JWT |
| Container | Docker |
| Hosting | DigitalOcean Droplet |
| Reverse Proxy | Nginx |
| SSL | Let's Encrypt |

## User Flows

### Guest

- View products
- View product details
- Add product to cart
- View cart
- Checkout
- View order confirmation
- Register
- Login
- Forgot password
- Reset password
- View support

### Authenticated

- View products
- View product details
- Add product to cart
- View cart
- Checkout
- View order confirmation
- View account
- Edit account
- Logout
- View support

## Development

### Prerequisites

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/products/docker-desktop)
- [DigitalOcean](https://www.digitalocean.com/)

### Setup

```sh
# Clone the repository
git clone

# Change into the directory
cd

# Build the container
docker build -t e-commerce .

# Run the container
docker run -p 8080:8080 e-commerce
```

### Development

```sh
# Change into the directory
cd

# Run the server
go run main.go
```

## Deployment

```sh
# Build the container for deployment to DigitalOcean droplet
docker build -t e-commerce .

# SCP the container to the DigitalOcean droplet
scp e-commerce root@<ip>:/root

# SSH into the DigitalOcean droplet
ssh root@<ip>

# Run the container
docker run -p 8080:8080 e-commerce
```

