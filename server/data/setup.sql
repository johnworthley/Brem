DROP TABLE developers CASCADE;
DROP TABLE auditors CASCADE;
DROP TABLE ico CASCADE;
DROP TABLE icoAuditors CASCADE;
DROP TYPE status;

CREATE TABLE developers (
  id        SERIAL PRIMARY KEY,
  address   VARCHAR(42) UNIQUE NOT NULL,
  username  VARCHAR(64) UNIQUE NOT NULL
);

CREATE TABLE auditors (
  id        SERIAL PRIMARY KEY,
  address   VARCHAR(42) UNIQUE NOT NULL,
  username  VARCHAR(64) UNIQUE NOT NULL
);

CREATE TYPE status AS ENUM (
  'created',
  'opened',
  'success',
  'failed',
  'requested',
  'withdrawn',
  'overdue'
);

CREATE TABLE  ico (
  id            SERIAL PRIMARY KEY,
  address       VARCHAR(42) UNIQUE NOT NULL,
  developerID   INT REFERENCES developers(id),
  cap           INT NOT NULL,
  description   TEXT,
  closingTime   TIMESTAMP NOT NULL,
  feePercent    INT NOT NULL,
  tokenAddress  VARCHAR(42) UNIQUE NOT NULL,
  name          VARCHAR(64) UNIQUE NOT NULL,
  symbol        VARCHAR(7) NOT NULL,
  status        status
);

CREATE TABLE icoAuditors (
  icoID     INT REFERENCES ico(id),
  auditorID INT REFERENCES auditors(id),
  PRIMARY KEY (icoID, auditorID)
);