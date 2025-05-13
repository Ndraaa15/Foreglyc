CREATE TABLE IF NOT EXISTS dietary_plans (
   user_id VARCHAR(255) NOT NULL PRIMARY KEY,
   live_with VARCHAR(255) NOT NULL,
   breakfast_time TIME NOT NULL,
   lunch_time TIME NOT NULL,
   dinner_time TIME NOT NULL,
   morning_snack_time TIME NOT NULL,
   afternoon_snack_time TIME NOT NULL,
   is_use_insuline BOOLEAN DEFAULT FALSE,
   insulise_questionnaires JSONB,
   total_daily_insuline DECIMAL,
   meal_plan_type VARCHAR(255) NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE,
   updated_at TIMESTAMP WITH TIME ZONE
);