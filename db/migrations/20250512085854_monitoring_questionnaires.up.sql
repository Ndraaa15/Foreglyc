CREATE TABLE IF NOT EXISTS monitoring_questionnaires (
   id BIGSERIAL PRIMARY KEY,
   glucometer_monitoring_id BIGINT REFERENCES glucometer_monitorings(id) ON DELETE CASCADE,
   questionnaires JSONB NOT NULL,
   management_type VARCHAR(255) NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE,
   updated_at TIMESTAMP WITH TIME ZONE
);
