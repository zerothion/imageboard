CREATE TABLE entities(
    entity_id  UUID  UNIQUE  NOT NULL  PRIMARY KEY  DEFAULT gen_random_uuid(),
    created_at TIMESTAMP  NOT NULL  DEFAULT now(),
    deleted_at TIMESTAMP  DEFAULT NULL
);
CREATE TABLE users(
    -- User --
    user_id UUID  UNIQUE  NOT NULL  PRIMARY KEY,
    name    TEXT,

    handle   TEXT UNIQUE,
    email    TEXT,
    password TEXT,

    -- Constraints --
    FOREIGN KEY (user_id)
        REFERENCES entities(entity_id)
        ON DELETE CASCADE
);
