-- Create games table with enhanced fields
CREATE TABLE games (
                       id SERIAL PRIMARY KEY,
                       game_type VARCHAR(10) NOT NULL,
                       status VARCHAR(20) NOT NULL DEFAULT 'ongoing',
                       start_time TIMESTAMPTZ,
                       end_time TIMESTAMPTZ,

                       metadata JSONB,
                       CONSTRAINT chk_game_type CHECK (game_type IN ('chess', 'fool'))
);

-- Index on game_type for performance
CREATE INDEX idx_game_type ON games(id);
