# Glans

A simple proof of concept (POC) for a MySQL proxy, with use cases involving PHP applications (CI3, WordPress, etc.) that functions as:

- **Connection Pool**: Given that PHP is stateless (open and close connections), it is not ideal for handling traffic spikes. This proxy provides a connection pool to manage such scenarios.
- **Smart Cache (Coming Soon)**: Designed for read-intensive tables, the proxy-level cache ensures performance while maintaining validation capabilities.
- **Load Balancer (Coming Soon)**: If you have more than two instances, this proxy can handle load balancing, suitable for master-slave setups.

## Why Not Just Use Vitess?

This project is part of my internship report, completed in just four days to showcase a cool implementation. The primary goal was to implement the MySQL protocol using a library.

## How to Use

Just like any ordinary MySQL connection, simply connect to the server, which is already compatible. It's similar to `pgpool`:

```sh
mysql -h 127.0.0.1 -P 4000 -u root
```