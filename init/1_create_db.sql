CREATE TABLE account (
    id SERIAL,
    user_name TEXT NOT NULL,
    created_at TIME NOT NULL DEFAULT NOW()
);