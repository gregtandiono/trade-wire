-- DESTROY TABLES FIRST

DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS buyers CASCADE;
DROP TABLE IF EXISTS commodities CASCADE;
DROP TABLE IF EXISTS varieties CASCADE;
DROP TABLE IF EXISTS suppliers CASCADE;
DROP TABLE IF EXISTS trades CASCADE;
DROP TABLE IF EXISTS contacts CASCADE;
DROP TABLE IF EXISTS tracking CASCADE;

-- DROP TRIGGERS

DROP TRIGGER IF EXISTS update_modified_column ON users;
DROP TRIGGER IF EXISTS update_modified_column ON buyers;
DROP TRIGGER IF EXISTS update_modified_column ON commodities;
DROP TRIGGER IF EXISTS update_modified_column ON varieties;
DROP TRIGGER IF EXISTS update_modified_column ON suppliers;
DROP TRIGGER IF EXISTS update_modified_column ON buyers;
DROP TRIGGER IF EXISTS update_modified_column ON trades;
DROP TRIGGER IF EXISTS update_modified_column ON contacts;
DROP TRIGGER IF EXISTS update_modified_column ON tracking;

-- DROP TYPES

DROP TYPE IF EXISTS contact_types;
DROP TYPE IF EXISTS user_types;

-- FUNCTIONS & TYPES

CREATE TYPE contact_types AS ENUM ('supplier', 'buyer');
CREATE TYPE user_types AS ENUM ('admin', 'employee');

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


CREATE TABLE IF NOT EXISTS buyers(
    id UUID PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL CHECK (name <> ''),
    address varchar(255),
    pic jsonb,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TRIGGER update_modified_column
BEFORE UPDATE ON buyers FOR EACH ROW EXECUTE PROCEDURE update_modified_column();


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


CREATE TABLE IF NOT EXISTS suppliers(
    id UUID PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL CHECK (name <> ''),
    address varchar(255) NOT NULL,
    pic jsonb,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TRIGGER update_modified_column
BEFORE UPDATE ON suppliers FOR EACH ROW EXECUTE PROCEDURE update_modified_column();


CREATE TABLE IF NOT EXISTS trades(
    id UUID PRIMARY KEY NOT NULL,
    buyer_id UUID NOT NULL,
    supplier_id UUID NOT NULL,
    commodity_id UUID NOT NULL,
    variety_id UUID NOT NULL,
    quantity int NOT NULL,
    shipment varchar(255) NOT NULL,
    price int NOT NULL,
    price_note varchar (255),
    status varchar(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (supplier_id) REFERENCES suppliers (id),
    FOREIGN KEY (buyer_id) REFERENCES buyers (id),
    FOREIGN KEY (commodity_id) REFERENCES commodities (id),
    FOREIGN KEY (variety_id) REFERENCES varieties (id)
);

CREATE TRIGGER update_modified_column
BEFORE UPDATE ON trades FOR EACH ROW EXECUTE PROCEDURE update_modified_column();


CREATE TABLE IF NOT EXISTS contacts(
    id UUID PRIMARY KEY NOT NULL,
    contact_type contact_types NOT NULL,
    name varchar(255) NOT NULL,
    position varchar(255) NOT NULL,
    office_number varchar(255),
    cell_number varchar(255),
    notes varchar(255),
    supplier_id UUID,
    buyer_id UUID,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (supplier_id) REFERENCES suppliers (id),
    FOREIGN KEY (buyer_id) REFERENCES buyers (id)
);

CREATE TRIGGER update_modified_column
BEFORE UPDATE ON contacts FOR EACH ROW EXECUTE PROCEDURE update_modified_column();


CREATE TABLE IF NOT EXISTS tracking(
    id UUID PRIMARY KEY NOT NULL,
    trade_id UUID NOT NULL,
    vessel jsonb,
    documents jsonb,
    notes varchar(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (trade_id) REFERENCES trades (id)
);

CREATE TRIGGER update_modified_column
BEFORE UPDATE ON tracking FOR EACH ROW EXECUTE PROCEDURE update_modified_column();