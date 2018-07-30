-- DROP TABLE developer CASCADE;
-- DROP TABLE auditor CASCADE;
-- DROP TABLE ico CASCADE;
-- DROP TABLE icoAuditors CASCADE;
-- DROP TABLE auditor CASCADE;

CREATE TABLE developer (
  id      SERIAL PRIMARY KEY,
  address VARCHAR(42) UNIQUE NOT NULL
);

CREATE TABLE auditor (
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
  developerID INT REFERENCES developer(id),
  status      status
);

CREATE TABLE icoAuditors (
  icoID     INT REFERENCES ico(id),
  auditorID INT REFERENCES auditor(id),
  PRIMARY KEY (icoID, auditorID)
);