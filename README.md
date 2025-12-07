# Go Blog CMS

A lightweight and efficient Blog Content Management System built with Go, Fiber, and GORM.

## Features

- **RESTful API** - Clean and intuitive API endpoints for blog management
- **Post Management** - Create and retrieve blog posts
- **Database Integration** - Uses GORM with SQLite for data persistence
- **Fiber Framework** - Built on the high-performance Fiber web framework

## API Endpoints

### Posts

- `GET /api/posts` - Retrieve all blog posts
- `POST /api/posts` - Create a new blog post

#### Post Object Structure
```json
{
  "id": 1,
  "title": "Sample Post",
  "content": "This is a sample blog post content.",
  "author": "John Doe",
  "created_at": "2025-12-07T08:30:21Z",
  "updated_at": "2025-12-07T08:30:21Z"
}
```

## Getting Started

### Prerequisites

- Go 1.16 or higher
- SQLite3 (or any other database supported by GORM)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/Go-cms-blog.git
   cd Go-cms-blog
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Configure the application by editing `config/config.yaml`

4. Run the application:
   ```bash
   go run main.go
   ```

The server will start on the configured port (default: 3000).

## Project Structure

```
.
├── config/           # Configuration files
├── internal/
│   ├── controllers/  # Request handlers
│   ├── models/       # Database models
│   ├── repository/   # Database connection and operations
│   └── routes/       # API route definitions
├── main.go           # Application entry point
└── README.md         # This file
```

## Configuration

Edit `config/config.yaml` to configure the application:

```yaml
server:
  port: 3000

database:
  dsn: blog.db
```

## Dependencies

- [Fiber](https://gofiber.io/) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [Viper](https://github.com/spf13/viper) - Configuration management

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.