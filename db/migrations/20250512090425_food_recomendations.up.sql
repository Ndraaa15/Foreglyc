CREATE TABLE IF NOT EXISTS food_recomendations (
   id BIGSERIAL NOT NULL,
   user_id VARCHAR(255) NOT NULL,
   food_name VARCHAR(255) NOT NULL,
   meal_time VARCHAR(255) NOT NULL,
   ingredients TEXT NOT NULL,
   calories_per_ingredients TEXT NOT NULL,
   image_url TEXT NOT NULL,
   total_calory INT NOT NULL,
   glycemic_index INT NOT NULL,
   date DATE NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE,
   updated_at TIMESTAMP WITH TIME ZONE
)