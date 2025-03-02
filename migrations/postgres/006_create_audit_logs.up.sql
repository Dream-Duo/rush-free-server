-- Audit logs for tracking user and vendor actions
CREATE TABLE audit_logs (
    log_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id),
    vendor_id INTEGER REFERENCES vendors(vendor_id),
    action TEXT NOT NULL,
    details JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);