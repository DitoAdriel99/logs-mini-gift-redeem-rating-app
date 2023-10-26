# Logs Gift Redemption and Rating Application
In the Logs Mini Gift App, our Logging Service plays a vital role in ensuring the reliability, security, and transparency of our operations. Logging is not just about recording events; it's a window into the inner workings of our application.

## Features
- Real-time Logging: Logs should be generated in real-time as events occur, allowing developers and administrators to monitor system activities without delay.
## Tech Specifications

### Libraries/Frameworks Used

- Go (version 1.19)
- Gorilla Mux
- PostgreSQL
- Pubsub

### Architecture/Modularity

The project follows a clean architecture pattern, separating concerns into different layers:

- **Presentation Layer**: Contains the API handlers and controllers.
- **Service Layer**: Contains the business logic.
- **Repository Layer**: Deals with data storage and retrieval.
- **Database Layer**: Connects to the database.
- **Test Layer**: Contains unit tests and integration tests.

## Quick Start
### Installation guide
#### 1. install go version 1.19
```bash
# Please Read This Link Installation Guide of GO

# Link Download -> https://go.dev/dl/
# Link Install -> https://go.dev/doc/install

```

#### 2. Run the application
```bash
# run command :
git clone https://github.com/DitoAdriel99/logs-mini-gift-redeem-rating-app.git

# install dependency
go mod tidy

# setup env
DB_DRIVER=postgres
DB_USERNAME=        #change to your db username
DB_PASSWORD=        #change to your db password
DB_HOST=            #change to your db host
DB_PORT=            #change to your db port 
DB_DATABASE=        #change to your db name 
DB_URL=             #postgres://{DB_USERNAME}:{DB_PASSWORD}@{DB_HOST}:{DB_PORT}/{DB_DATABASE}?sslmode=disable

KEY=                #change to your key
EXPIRED=            #change to your expiration time

LOGS_SERVICE=       #change to your logs service url

# Run App
make start

# Migrate db
make migrate-up //this for up migrations
make migrate-down //this for down migrations
```