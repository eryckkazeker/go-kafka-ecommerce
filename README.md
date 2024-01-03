# Go - Kafka - eCommerce
The goal of this project is to demonstrate the integration of multiple golang written services through a kafka server.

The repo contains a docker-compose file that can run a simple kafka server that will be used by the services.

# Services
## Store Service
The store service implements a simple in-memory database that holds example products. These products can be listed and an order can be placed through a REST API.

When a new order is placed, a message is dispatched through a Kafka Producer.

## Mail Service
When a new order is placed, this service sends a mock e-mail with the order id

## Warehouse Service
When a new order is placed, this service alerts the warehouse to separate and ship the ordered products.
