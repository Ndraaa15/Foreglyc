CREATE TABLE IF NOT EXISTS glucometer_monitoring_preferences (
   user_id VARCHAR(255) NOT NULL PRIMARY KEY,
   start_wake_up_time TIME NOT NULL,
   end_wake_up_time TIME NOT NULL,
   physical_activity_days TEXT[],
   start_sleep_time TIME NOT NULL,
   end_sleep_time TIME NOT NULL,
   hypoglycemia_accute_threshold DECIMAL(5, 2) NOT NULL,
   hypoglycemia_chronic_threshold DECIMAL(5, 2) NOT NULL,
   hyperglycemia_accute_threshold DECIMAL(5, 2) NOT NULL,
   hyperglycemia_chronic_threshold DECIMAL(5, 2) NOT NULL,
   send_notification BOOLEAN DEFAULT FALSE,
   created_at TIMESTAMP WITH TIME ZONE,
   updated_at TIMESTAMP WITH TIME ZONE
);