-- DESTROY TABLES FIRST

DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS companies CASCADE;
DROP TABLE IF EXISTS commodities CASCADE;
DROP TABLE IF EXISTS varieties CASCADE;
DROP TABLE IF EXISTS vessels CASCADE;
DROP TABLE IF EXISTS trades CASCADE;
DROP TABLE IF EXISTS contacts CASCADE;

-- DROP TRIGGERS

DROP TRIGGER IF EXISTS update_modified_column ON users;
DROP TRIGGER IF EXISTS update_modified_column ON companies;
DROP TRIGGER IF EXISTS update_modified_column ON commodities;
DROP TRIGGER IF EXISTS update_modified_column ON varieties;
DROP TRIGGER IF EXISTS update_modified_column ON vessels;
DROP TRIGGER IF EXISTS update_modified_column ON trades;
DROP TRIGGER IF EXISTS update_modified_column ON contacts;

-- DROP TYPES

DROP TYPE IF EXISTS company_types;
DROP TYPE IF EXISTS user_types;

-- FUNCTIONS & TYPES

CREATE TYPE user_types AS ENUM ('admin', 'employee');
CREATE TYPE company_types AS ENUM ('supplier', 'buyer');

CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.modified = now();
  RETURN NEW;
END;
$$ language 'plpgsql';


-- Tables

CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL CHECK (name <> ''),
    username varchar(30) NOT NULL,
    password varchar(255) NOT NULL,
    type user_types NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TRIGGER update_modified_column
BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_modified_column();


CREATE TABLE IF NOT EXISTS commodities(
    id UUID PRIMARY KEY NOT NULL CHECK (name <> ''),
    name varchar(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TRIGGER update_modified_column
BEFORE UPDATE ON commodities FOR EACH ROW EXECUTE PROCEDURE update_modified_column();


CREATE TABLE IF NOT EXISTS varieties(
    id UUID PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL CHECK (name <> ''),
    commodity_id UUID NOT NULL,
    origin varchar(255) NOT NULL,
    specs text,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (commodity_id) REFERENCES commodities (id)
);

CREATE TRIGGER update_modified_column
BEFORE UPDATE ON varieties FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

CREATE TABLE IF NOT EXISTS companies(
    id UUID PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL CHECK (name <> ''),
    address varchar(255) NOT NULL,
    company_type company_types NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TRIGGER update_modified_column
BEFORE UPDATE ON companies FOR EACH ROW EXECUTE PROCEDURE update_modified_column();


CREATE TABLE IF NOT EXISTS contacts(
    id UUID PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL,
    position varchar(255) NOT NULL,
    office_number varchar(255),
    cell_number varchar(255),
    notes varchar(255),
    company_id UUID,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (company_id) REFERENCES companies (id)
);

CREATE TRIGGER update_modified_column
BEFORE UPDATE ON contacts FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

CREATE TABLE IF NOT EXISTS vessels(
    id UUID PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL,
    beam varchar(255),
    LOA varchar(255),
    draft varchar(255),
    status varchar(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TRIGGER update_modified_column
BEFORE UPDATE ON vessels FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

CREATE TABLE IF NOT EXISTS trades(
    id UUID PRIMARY KEY NOT NULL,
    buyer_id UUID NOT NULL,
    supplier_id UUID NOT NULL,
    variety_id UUID NOT NULL,
    vessel_id UUID,
    quantity int NOT NULL,
    bl_quantity int,
    shipment tstzrange NOT NULL,
    price int NOT NULL,
    price_note varchar (255),
    status varchar(255),
    notes varchar(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (buyer_id) REFERENCES companies (id),
    FOREIGN KEY (supplier_id) REFERENCES companies (id),
    FOREIGN KEY (variety_id) REFERENCES varieties (id),
    FOREIGN KEY (vessel_id) REFERENCES vessels (id)
);

CREATE TRIGGER update_modified_column
BEFORE UPDATE ON trades FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

-- CREATE TABLE reservation (room int, during tsrange);
-- INSERT INTO reservation VALUES
--     (1108, '[2010-01-01 14:30, 2010-01-01 15:30)');
--
-- -- Containment
-- SELECT int4range(10, 20) @> 3;
--
-- -- Overlaps
-- SELECT numrange(11.1, 22.2) && numrange(20.0, 30.0);
--
-- -- Extract the upper bound
-- SELECT upper(int8range(15, 25));
--
-- -- Compute the intersection
-- SELECT int4range(10, 20) * int4range(15, 25);
--
-- -- Is the range empty?
-- SELECT isempty(numrange(1, 5));

-- INSERT INTO tests VALUES
-- ('fab701fa-6d9b-4ba5-a14d-20a2a726371c','[2017-07-20T16:03:29+07:00, 2017-07-20T16:03:32+07:00]');
