CREATE TABLE videos (
    id UUID PRIMARY KEY,
    original_filename TEXT NOT NULL,
    content_type TEXT NOT NULL,
    size BIGINT NOT NULL,
    status TEXT NOT NULL,
    s3_raw_key TEXT NOT NULL,
    s3_hls_key TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);