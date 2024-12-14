-- Users table
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    phone_number VARCHAR(20) UNIQUE,
    email VARCHAR(255) UNIQUE,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    phone_verified BOOLEAN DEFAULT false,
    phone_verification_code VARCHAR(6),
    phone_verification_expires_at TIMESTAMP,
    preferences JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP
);

-- OAuth connections for users
CREATE TABLE user_oauth_connections (
    connection_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(user_id),
    provider VARCHAR(50) NOT NULL,
    provider_user_id VARCHAR(255) NOT NULL,
    access_token TEXT,
    refresh_token TEXT,
    token_expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(provider, provider_user_id)
);

-- Indexes for authentication
CREATE INDEX idx_user_phone ON users(phone_number) WHERE phone_verified = true;
CREATE INDEX idx_user_oauth ON user_oauth_connections(provider, provider_user_id);
