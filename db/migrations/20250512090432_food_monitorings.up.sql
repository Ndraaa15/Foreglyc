CREATE TABLE IF NOT EXISTS food_monitorings (
   id BIGSERIAL NOT NULL,
   user_id VARCHAR(255) NOT NULL,
   food_name VARCHAR(255) NOT NULL,
   meal_time VARCHAR(255) NOT NULL,
   image_url TEXT NOT NULL,
   nutritions JSONB NOT NULL,
   total_calory INT NOT NULL,
   total_carbohydrate INT NOT NULL,
   total_protein INT NOT NULL,
   total_fat INT NOT NULL,
   glycemic_index INT NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE,
   updated_at TIMESTAMP WITH TIME ZONE
)