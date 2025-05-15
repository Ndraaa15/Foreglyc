CREATE TABLE IF NOT EXISTS dietary_informations (
   user_id VARCHAR(255) NOT NULL PRIMARY KEY,
   total_calory INT NOT NULL,
   total_breakfast_calory INT NOT NULL,
   total_snack_calory INT NOT NULL,
   total_lunch_calory INT NOT NULL,
   total_dinner_calory INT NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE,
   updated_at TIMESTAMP WITH TIME ZONE
)