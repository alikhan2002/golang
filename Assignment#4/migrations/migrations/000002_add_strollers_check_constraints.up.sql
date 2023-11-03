ALTER TABLE strollers ADD CONSTRAINT strollers_price_check CHECK (price > 0);
ALTER TABLE strollers ADD CONSTRAINT strollers_title_check CHECK (title is not null);
ALTER TABLE strollers ADD CONSTRAINT ages_length_check CHECK (ages is not null);