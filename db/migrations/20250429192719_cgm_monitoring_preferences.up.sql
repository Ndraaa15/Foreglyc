CREATE TABLE IF NOT EXISTS cgm_monitoring_preferences (
   user_id VARCHAR(255) NOT NULL PRIMARY KEY,
   hypoglycemia_accute_threshold DECIMAL(5, 2) NOT NULL,
   hypoglycemia_chronic_threshold DECIMAL(5, 2) NOT NULL,
   hyperglycemia_accute_threshold DECIMAL(5, 2) NOT NULL,
   hyperglycemia_chronic_threshold DECIMAL(5, 2) NOT NULL,
   send_notification BOOLEAN DEFAULT FALSE,
   created_at TIMESTAMP WITH TIME ZONE,
   updated_at TIMESTAMP WITH TIME ZONE
);