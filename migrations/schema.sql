-- Create language table
CREATE TABLE IF NOT EXISTS language (
                                        language_id serial PRIMARY KEY,
                                        name varchar(20) NOT NULL,
    last_update timestamp NOT NULL DEFAULT now()
    );

-- Create film table
CREATE TABLE IF NOT EXISTS film (
                                    film_id serial PRIMARY KEY,
                                    title varchar(255) NOT NULL,
    description text,
    release_year integer,
    language_id smallint NOT NULL REFERENCES language(language_id),
    original_language_id smallint REFERENCES language(language_id),
    rental_duration smallint NOT NULL DEFAULT 3,
    rental_rate numeric(4,2) NOT NULL DEFAULT 4.99,
    length smallint,
    replacement_cost numeric(5,2) NOT NULL DEFAULT 19.99,
    rating varchar(10),
    last_update timestamp NOT NULL DEFAULT now(),
    special_features text[],
    fulltext tsvector NOT NULL
    );