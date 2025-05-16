CREATE TABLE IF NOT EXISTS report_informations (
   id BIGSERIAL PRIMARY KEY,
   user_id VARCHAR(255) NOT NULL,
   total_blood_sugar DECIMAL NOT NULL,
   recommendation_blood_glucose TEXT NOT NULL,
   recommendation TEXT NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE,
   updated_at TIMESTAMP WITH TIME ZONE
);