-- DROP TABLE developers CASCADE;
-- DROP TABLE auditors CASCADE;
-- DROP TABLE ico CASCADE;
-- DROP TABLE icoAuditors CASCADE;
-- DROP TABLE auditors CASCADE;

CREATE TABLE developers (
  id      SERIAL PRIMARY KEY,
  address VARCHAR(42) UNIQUE NOT NULL
);

CREATE TABLE auditors (
  id      SERIAL PRIMARY KEY,
  address VARCHAR(42) UNIQUE NOT NULL
);

CREATE TYPE status AS ENUM (
  'created',
  'opened',
  'success',
  'failed',
  'requested',
  'withdrawn'
);

CREATE TABLE  ico (
  id          SERIAL PRIMARY KEY,
  address     VARCHAR(42) UNIQUE NOT NULL,
  developerID INT REFERENCES developers(id),
  status      status
);

CREATE TABLE icoAuditors (
  icoID     INT REFERENCES ico(id),
  auditorID INT REFERENCES auditors(id),
  PRIMARY KEY (icoID, auditorID)
);