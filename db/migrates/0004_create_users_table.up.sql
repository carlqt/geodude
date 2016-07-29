CREATE TABLE users(
  id serial primary key not null,
  email varchar(255),
  username varchar(25),
  first_name varchar(255),
  last_name varchar(255),
  contact_number char(15),
  password char(60),
  role varchar(10),
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);
