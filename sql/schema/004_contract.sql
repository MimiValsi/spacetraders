-- +goose Up
CREATE TABLE payments(
  id SERIAL NOT NULL PRIMARY KEY,
  on_accepted INTEGER NOT NULL,
  on_fulfilled INTEGER NOT NULL
);

CREATE TABLE delivers(
  id SERIAL NOT NULL PRIMARY KEY,
  trade_symbol TEXT NOT NULL,
  destination_symbol TEXT NOT NULL,
  units_required INTEGER NOT NULL,
  units_fulfilled INTEGER NOT NULL
);

CREATE TABLE terms(
  id SERIAL NOT NULL PRIMARY KEY,
  deadline TIMESTAMPTZ NOT NULL,
  payment_id INTEGER REFERENCES payments(id) ON DELETE CASCADE NOT NULL,
  deliver_id INTEGER REFERENCES delivers(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE contracts(
  id TEXT NOT NULL PRIMARY KEY,
  faction_symbol TEXT NOT NULL,
  type TEXT NOT NULL,
  terms_id INTEGER REFERENCES terms(id) ON DELETE CASCADE NOT NULL,
  accepted BOOL NOT NULL DEFAULT false,
  fulfilled BOOL NOT NULL DEFAULT false,
  deadline_to_accept TIMESTAMPTZ NOT NULL
);
--

-- +goose Down
DROP TABLE contracts;

DROP TABLE terms;

DROP TABLE payments;

DROP TABLE delivers;
--
