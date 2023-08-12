# MySQL-Golang-Redis pipeline

MySQL-Golang-Redis is a professional-grade Golang project designed to efficiently retrieve and aggregate data from a MySQL database and store it in a Redis database using the hash data type.

## MySQL install and configuration

The MySQL-Golang-Redis project utilizes the Go programming language to establish a connection with a MySQL database running on Windows 10. It leverages the native MySQL driver for Go to execute optimized queries and efficiently retrieve the required data.

### mysql-server Install on Windows 10

It's important to note that there may be compatibility issues when using [MySQL](https://dev.mysql.com/downloads/file/?id=518840/) server version 5.7 on Ubuntu 22.04. Due to potential conflicts or incompatibilities between the MySQL server version and the Ubuntu operating system.

### mysql-server Configuration

To access to database from another computer, we have to set user permission

```bash
mysql -u root -h <windows_machine_ip_address> -p

```
```sql
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY '<root_password>' WITH GRANT OPTION;

FLUSH PRIVILEGES;
```
### mysql-client install
```
sudo apt update
sudo apt install mysql-client
```
## Development environment setting

### Golang install

The following installation command will be used to install Go on Ubuntu:
```bash
sudo apt update
sudo apt install golang-go
```
After installation, check golang on ubuntu.
```bash
go version
```
### Redis server install
```bash
sudo apt update
sudo add-apt-repository ppa:redislabs/redis
sudo apt-get install redis
```
After install, check server and get some configuration
```bash
redis-server -v
sudo systemctl enable --now redis-server
sudo nano /etc/redis/redis.conf
```
    Find the line stating the “bind” address as “127.0.0.1”:
    Restart server.
```bash
sudo systemctl restart redis-server
```

## Code feature

By utilizing Golang's concurrency features, the project ensures optimal performance by executing multiple database queries and Redis operations concurrently. This approach maximizes throughput and minimizes latency, making it suitable for handling large-scale datasets.
### CLI script
```bash
git clone https://github.com/cozyguy/mysql-golang-redis-pipeline.git
```
```cli
go mod init <project_name>
go mod tidy
go run main.go
```

## Conclusion

The MySQL-Golang-Redis project follows best practices for error handling, logging, and configuration management. It provides a clean and modular codebase that is easy to understand, maintain, and extend. Additionally, it includes comprehensive unit tests to ensure the reliability and correctness of the implemented functionality.
Overall, MySQL-Golang-Redis is a professional-level Golang project that combines the power of MySQL and Redis to efficiently aggregate and store data, making it an ideal choice for applications requiring real-time analytics and data processing.