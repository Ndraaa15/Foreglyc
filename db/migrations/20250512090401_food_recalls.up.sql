CREATE TABLE IF NOT EXISTS food_recalls (
   id BIGSERIAL NOT NULL,
   user_id VARCHAR(255) NOT NULL,
   food_name VARCHAR(255) NOT NULL,
   time_type VARCHAR(255) NOT NULL,
   image_url TEXT NOT NULL,
   nutritions JSONB NOT NULL,
   total_calories INT NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE,
   updated_at TIMESTAMP WITH TIME ZONE
)