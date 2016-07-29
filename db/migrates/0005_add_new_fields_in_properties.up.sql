ALTER TABLE properties
ADD COLUMN name varchar(50),
ADD COLUMN description varchar(255),
ADD COLUMN price numeric(10, 2),
ADD COLUMN type varchar(15),
ADD COLUMN agent_id integer REFERENCES users;
