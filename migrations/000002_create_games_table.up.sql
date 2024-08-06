-- Create games table with enhanced fields
CREATE TABLE games (
                       id SERIAL PRIMARY KEY,
                       game_type VARCHAR(10) NOT NULL,
                       duration INTERVAL,
                       status VARCHAR(20) NOT NULL DEFAULT 'ongoing',
                       start_time TIMESTAMPTZ,
                       end_time TIMESTAMPTZ,
                       score VARCHAR(50),
                       metadata JSONB,
                       CONSTRAINT chk_game_type CHECK (game_type IN ('chess', 'fool'))
);

-- Create game_users table with unique constraint
CREATE TABLE game_users (
                            game_id INT NOT NULL,
                            user_id INT NOT NULL,
                            PRIMARY KEY (game_id, user_id),
                            FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE,
                            FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Index on game_type for performance
CREATE INDEX idx_game_type ON games(id);

ALTER TABLE game_users
    ADD CONSTRAINT unique_game_user UNIQUE (game_id, user_id);
