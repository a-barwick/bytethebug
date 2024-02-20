# README

E-commerce template for an alien widget factory.

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
| Hosting | Google Cloud Run |

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
- [Google Cloud SDK](https://cloud.google.com/sdk/docs/install)

### Setup

```sh
# Clone the repository
git clone

# Change into the directory
cd

# Build the container
docker build -t alien-widget-factory .

# Run the container
docker run -p 8080:8080 alien-widget-factory
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
# Change into the directory
cd

# Build the container
docker build -t gcr.io/alien-widget-factory/alien-widget-factory .

# Push the container to Google Cloud Registry
docker push gcr.io/alien-widget-factory/alien-widget-factory

# Deploy the container to Google Cloud Run
gcloud run deploy alien-widget-factory --image gcr.io/alien-widget-factory/alien-widget-factory --platform managed
```

