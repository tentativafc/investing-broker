CREATE TABLE IF NOT EXISTS public.asset_quotation
(
    id               VARCHAR NOT NULL,
    date             TIMESTAMP,
    symbol           VARCHAR,
    market_type      INTEGER,
    bdi_code         INTEGER,
    days_term_market INTEGER,
    min_price        DOUBLE PRECISION,
    max_price        DOUBLE PRECISION,
    open_price       DOUBLE PRECISION,
    close_price      DOUBLE PRECISION,
    volume           DOUBLE PRECISION,
    CONSTRAINT asset_quotation_pkey PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS asset_quotation_market_type_bdi_code_idx
  ON asset_quotation (market_type, bdi_code);


CREATE INDEX IF NOT EXISTS asset_quotation_symbol_date_idx
  ON asset_quotation (symbol, date);


CREATE TABLE IF NOT EXISTS public.currency_quotation
(
    id         VARCHAR NOT NULL,
    date       TIMESTAMP,
    symbol     VARCHAR,
    buy_price  DOUBLE PRECISION,
    sell_price DOUBLE PRECISION,
    CONSTRAINT currency_quotation_pkey PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS currency_quotation_symbol_date_idx
  ON currency_quotation (symbol, date);


CREATE TABLE IF NOT EXISTS public.coin_quotation
(
    id          VARCHAR NOT NULL,
    date        TIMESTAMP,
    symbol      VARCHAR,
    min_price   DOUBLE PRECISION,
    max_price   DOUBLE PRECISION,
    open_price  DOUBLE PRECISION,
    close_price DOUBLE PRECISION,
    volume      DOUBLE PRECISION,
    quantity    DOUBLE PRECISION,
    amount      DOUBLE PRECISION,
    avg_price   DOUBLE PRECISION,
    CONSTRAINT coin_quotation_pkey PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS coin_quotation_symbol_date_idx
  ON coin_quotation (symbol, date);