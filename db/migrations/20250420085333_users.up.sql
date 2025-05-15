CREATE TABLE IF NOT EXISTS users (
   id VARCHAR(255) PRIMARY KEY,
   full_name VARCHAR(255) NOT NULL,
   email VARCHAR(100) UNIQUE NOT NULL,
   password TEXT,
   photo_profile TEXT,
   is_verified BOOLEAN DEFAULT FALSE,
   body_weight DECIMAL(5, 2),
   date_of_birth DATE,
   address TEXT,
   caregiver_contact VARCHAR(255),
   auth_provider INT NOT NULL,
   level VARCHAR(255) NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE,
   updated_at TIMESTAMP WITH TIME ZONE
);
