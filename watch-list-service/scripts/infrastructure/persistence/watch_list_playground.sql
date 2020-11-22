INSERT INTO director(id, user_id, display_name) VALUES ('8EFckHhch1uHsZ8Q',
                                                        'ee154507-c93d-42f6-bbd9-a32e47151057',
                                                        'Josh Safdie');

INSERT INTO director(id, user_id, display_name) VALUES ('sf9inHnsluMRioeV',
                                                        'ee154507-c93d-42f6-bbd9-a32e47151057',
                                                        'Benny Safdie');

INSERT INTO director(id, user_id, display_name) VALUES ('hklWm5SFHDxe8R8m',
                                                        'ee154507-c93d-42f6-bbd9-a32e47151057',
                                                        'Orson Welles');

INSERT INTO director(id, user_id, display_name) VALUES ('L_y48xcwxoj9PPlp',
                                                        'ee154507-c93d-42f6-bbd9-a32e47151057',
                                                        'Denis Villeneuve');

SELECT * FROM director;

INSERT INTO movie(id, user_id, display_name, year) VALUES ('zwfXK5PwCY7fTz31', 'ee154507-c93d-42f6-bbd9-a32e47151057',
                                                           'Uncut Gems', 2019);

INSERT INTO movie(id, user_id, display_name, year) VALUES ('WtnFqyElemXlKmvK', 'ee154507-c93d-42f6-bbd9-a32e47151057',
                                                           'Good Time', 2017);

INSERT INTO movie(id, user_id, display_name, year) VALUES ('nMHQPO74mMmApRqM', 'ee154507-c93d-42f6-bbd9-a32e47151057',
                                                           'Blade Runner 2049', 2017);

INSERT INTO movie(id, user_id, display_name, year) VALUES ('IYhcjNB8_9zooXbJ', 'ee154507-c93d-42f6-bbd9-a32e47151057',
                                                           'Citizen Kane', 1941);

SELECT * FROM movie;

INSERT INTO movie_directors VALUES ('zwfXK5PwCY7fTz31', '8EFckHhch1uHsZ8Q');
INSERT INTO movie_directors VALUES ('zwfXK5PwCY7fTz31', 'sf9inHnsluMRioeV');
INSERT INTO movie_directors VALUES ('WtnFqyElemXlKmvK', '8EFckHhch1uHsZ8Q');
INSERT INTO movie_directors VALUES ('IYhcjNB8_9zooXbJ', 'hklWm5SFHDxe8R8m');
INSERT INTO movie_directors VALUES ('nMHQPO74mMmApRqM', 'L_y48xcwxoj9PPlp');

SELECT * FROM movie_directors;

-- Movies with directors
SELECT m.id as movie_id, m.display_name as movie_name, d.id as director_id, d.display_name as director_name
FROM director as d, movie as m, movie_directors as md WHERE d.id = md.director_id AND m.id = md.movie_id
ORDER BY movie_name;

SELECT * FROM movie_by_directors;

-- Movies by director
SELECT * FROM movie INNER JOIN movie_directors md ON movie.id = md.movie_id AND director_id = 'hklWm5SFHDxe8R8m';