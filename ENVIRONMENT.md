# Environment Variables

Currently, the application uses **default configuration** suitable for development. 

However, for production or custom setups, we recommend implementing environment variables. The system is designed to be easily extensible to support `.env` configuration.

## Recommended Configuration (Future Implementation)

Create a `.env` file in the root directory:

```ini
# Server Configuration
PORT=3000
APP_ENV=development # development | production

# Database Configuration
# Currently hardcoded to SQLite ("votepeace.db")
# Future Use:
# DB_HOST=localhost
# DB_USER=root
# DB_PASS=password
# DB_NAME=votepeace

# Security
# Secret key for JWT signing / Session encryption
JWT_SECRET=your-super-secret-key-change-me

# Client URL (CORS)
CLIENT_URL=http://localhost:5173
```

## Current Defaults

- **Port**: `:3000`
- **Database**: SQLite (`votepeace.db` in root)
- **CORS**: Allows `http://localhost:5173`
- **Admin Seed**:
    - NIK: `0000000000000000`
    - Password: `admin123`
