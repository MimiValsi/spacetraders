-- +goose Up
CREATE TABLE contracts(
  cid SERIAL NOT NULL PRIMARY KEY,
  id TEXT NOT NULL,
  faction_symbol TEXT NOT NULL,
  type TEXT NOT NULL,
  accepted BOOL NOT NULL DEFAULT false,
  fulfilled BOOL NOT NULL DEFAULT false,
  deadline_to_accept TIMESTAMPTZ NOT NULL
);

CREATE TABLE terms(
  id SERIAL NOT NULL PRIMARY KEY,
  deadline TIMESTAMPTZ NOT NULL,
  contract_id INTEGER REFERENCES contracts(cid) ON DELETE CASCADE NOT NULL
);

CREATE TABLE payments(
  id SERIAL NOT NULL PRIMARY KEY,
  on_accepted INTEGER NOT NULL,
  on_fulfilled INTEGER NOT NULL,
  term_id INTEGER REFERENCES terms(id) ON DELETE CASCADE NOT NULL,
  contract_id INTEGER REFERENCES contracts(cid) ON DELETE CASCADE NOT NULL
);

CREATE TABLE delivers(
  id SERIAL NOT NULL PRIMARY KEY,
  trade_symbol TEXT NOT NULL,
  destination_symbol TEXT NOT NULL,
  units_required INTEGER NOT NULL,
  units_fulfilled INTEGER NOT NULL,
  term_id INTEGER REFERENCES terms(id) ON DELETE CASCADE NOT NULL,
  contract_id INTEGER REFERENCES contracts(cid) ON DELETE CASCADE NOT NULL
);
--

-- +goose Down
DROP TABLE contracts CASCADE;
DROP TABLE payments;
DROP TABLE delivers;
DROP TABLE terms;
--
