CREATE TABLE m_payment_methods (
   id VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
   name VARCHAR(128),
   description LONGTEXT,
   image VARCHAR(255) default 'default-male.png',,
   created_at DECIMAL NOT NULL,
   created_by VARCHAR(64),
   updated_at DECIMAL NOT NULL,
   updated_by VARCHAR(64),
   deleted_at DECIMAL,
   deleted_by VARCHAR(64)
);