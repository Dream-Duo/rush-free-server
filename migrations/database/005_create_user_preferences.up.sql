-- User Saved Places
CREATE TABLE user_saved_places (
    saved_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(user_id),
    restaurant_id INTEGER NOT NULL REFERENCES restaurants(restaurant_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, restaurant_id)
);

-- User Preferences 
CREATE TABLE user_preferences (
    preference_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(user_id),
    preference_key VARCHAR(100),
    preference_value VARCHAR(255),
    UNIQUE(user_id, preference_key)
);