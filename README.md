# Delete Product Microservice

This microservice handles the deletion of products from the ToyShop platform. It is part of the product management domain and uses MongoDB as the data store.

## Technologies Used

- Go (Golang)
- Gin (web framework)
- MongoDB
- Docker
- GitHub Actions

## Getting Started

### Prerequisites

- Go >= 1.18
- MongoDB
- Git

### Installation

```bash
git clone https://github.com/andrespaida/delete_product.git
cd delete_product
go mod tidy
```

### Environment Variables

Create a `.env` file in the root directory with the following content:

```env
PORT=4004
MONGO_URI=mongodb://your_mongo_host:27017
DB_NAME=toyshop_db
COLLECTION_NAME=products
```

### Running the Service

```bash
go run main.go
```

The service will be running at `http://localhost:4004`.

## Available Endpoint

### DELETE `/products/:id`

Deletes a product with the specified ID.

#### Example Response:

```json
{
  "message": "Product deleted successfully"
}
```

## Docker

To build and run the service using Docker:

```bash
docker build -t delete-product .
docker run -p 4004:4004 --env-file .env delete-product
```

## GitHub Actions Deployment

This project includes a GitHub Actions workflow for automatic deployment to an EC2 instance. Configure the following secrets in your GitHub repository:

- `EC2_HOST`
- `EC2_USERNAME`
- `EC2_KEY`
- `EC2_PORT` (optional)

## License

This project is licensed under the MIT License.