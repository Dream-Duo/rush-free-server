-- Vendors table
CREATE TABLE vendors (
    vendor_id SERIAL PRIMARY KEY,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE,
    business_name VARCHAR(255) NOT NULL,
    contact_person_name VARCHAR(200),
    phone_verified BOOLEAN DEFAULT false,
    phone_verification_code VARCHAR(6),
    phone_verification_expires_at TIMESTAMP,
    tax_id VARCHAR(50),
    business_license VARCHAR(50),
    is_verified BOOLEAN DEFAULT false,
    verification_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP
);

-- OAuth connections for vendors
CREATE TABLE vendor_oauth_connections (
    connection_id SERIAL PRIMARY KEY,
    vendor_id INTEGER NOT NULL REFERENCES vendors(vendor_id),
    provider VARCHAR(50) NOT NULL,
    provider_user_id VARCHAR(255) NOT NULL,
    access_token TEXT,
    refresh_token TEXT,
    token_expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(provider, provider_user_id)
);

-- Indexes for authentication
CREATE INDEX idx_vendor_phone ON vendors(phone_number) WHERE phone_verified = true;
CREATE INDEX idx_vendor_oauth ON vendor_oauth_connections(provider, provider_user_id);
