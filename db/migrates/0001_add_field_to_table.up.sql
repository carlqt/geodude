CREATE TABLE properties(
  id serial primary key not null,
  longitude float,
  latitude float,
  address varchar(255),
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);
