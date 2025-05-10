INSERT INTO film (
    title, description, release_year, language_id,
    rental_duration, rental_rate, length, replacement_cost,
    rating, special_features, fulltext
) VALUES
      ('The Matrix', 'A computer hacker learns about the true nature of reality', 1999, 1,
       7, 4.99, 136, 19.99,
       'R', '{"Trailers","Commentaries"}', to_tsvector('The Matrix')),

      ('Inception', 'A thief who steals corporate secrets through dream-sharing technology', 2010, 1,
       5, 4.99, 148, 19.99,
       'PG-13', '{"Trailers","Commentaries"}', to_tsvector('Inception')),

      ('Amélie', 'A whimsical story about a young woman who decides to change the lives of those around her', 2001, 3,
       6, 3.99, 122, 15.99,
       'R', '{"Trailers","Commentaries"}', to_tsvector('Amélie'));