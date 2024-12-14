-- Aggregated historical rush metrics (for analytics)
CREATE TABLE rush_metrics_aggregated (
    metric_id SERIAL PRIMARY KEY,
    restaurant_id INTEGER NOT NULL REFERENCES restaurants(restaurant_id),
    time_bucket TIMESTAMP NOT NULL,
    average_wait_time INTEGER NOT NULL,
    average_occupancy INTEGER NOT NULL,
    data_points INTEGER NOT NULL DEFAULT 1,
    vendor_updates INTEGER NOT NULL DEFAULT 0,
    user_updates INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(restaurant_id, time_bucket)
);

-- Indexes
CREATE INDEX idx_rush_metrics_agg_time ON rush_metrics_aggregated(restaurant_id, time_bucket);
