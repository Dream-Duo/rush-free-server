-- Define ENUM types
DO $$ BEGIN
    CREATE TYPE restaurant_status_enum AS ENUM('active', 'temporarily_closed', 'permanently_closed', 'opening_soon');
EXCEPTION WHEN duplicate_object THEN
    -- Do nothing, type already exists
END $$;

DO $$ BEGIN
    CREATE TYPE image_type_enum AS ENUM('exterior', 'interior', 'food', 'menu', 'other');
EXCEPTION WHEN duplicate_object THEN
    -- Do nothing, type already exists
END $$;

-- Restaurants table
CREATE TABLE restaurants (
    restaurant_id SERIAL PRIMARY KEY,
    vendor_id INTEGER NOT NULL REFERENCES vendors(vendor_id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    cuisine_type VARCHAR(100),
    address TEXT NOT NULL,
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    phone VARCHAR(20),
    email VARCHAR(255),
    seating_capacity INTEGER,
    status restaurant_status_enum NOT NULL DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Restaurant operating hours
CREATE TABLE restaurant_hours (
    hours_id SERIAL PRIMARY KEY,
    restaurant_id INTEGER NOT NULL REFERENCES restaurants(restaurant_id),
    day_of_week INTEGER NOT NULL CHECK (day_of_week BETWEEN 0 AND 6),
    opening_time TIME NOT NULL,
    closing_time TIME NOT NULL,
    is_closed BOOLEAN DEFAULT false,
    special_hours_note TEXT,
    UNIQUE(restaurant_id, day_of_week)
);

-- Restaurant images
CREATE TABLE restaurant_images (
    image_id SERIAL PRIMARY KEY,
    restaurant_id INTEGER NOT NULL REFERENCES restaurants(restaurant_id),
    image_url VARCHAR(255) NOT NULL,
    image_type image_type_enum,
    is_primary BOOLEAN DEFAULT false,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Restaurant amenities
CREATE TABLE restaurant_amenities (
    amenity_id SERIAL PRIMARY KEY,
    restaurant_id INTEGER NOT NULL REFERENCES restaurants(restaurant_id),
    amenity_name VARCHAR(100),
    is_available BOOLEAN DEFAULT true
);

-- Restaurant ratings and reviews
CREATE TABLE restaurant_reviews (
    review_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(user_id),
    restaurant_id INTEGER NOT NULL REFERENCES restaurants(restaurant_id),
    rating INTEGER CHECK (rating BETWEEN 1 AND 5),
    review_text TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_restaurants_location ON restaurants(latitude, longitude);
CREATE INDEX idx_restaurant_hours_search ON restaurant_hours(day_of_week, opening_time, closing_time);
