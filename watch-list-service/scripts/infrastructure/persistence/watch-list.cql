CREATE TABLE IF NOT EXISTS category (
    id text,
    user_id text,
    display_name text,
    create_time timestamp,
    update_time timestamp,
    active boolean,
    PRIMARY KEY (id, update_time)
) WITH CLUSTERING ORDER BY (update_time ASC);

CREATE TABLE IF NOT EXISTS category_user (
    id text,
    user_id text,
    display_name text,
    create_time timestamp,
    update_time timestamp,
    active boolean,
    PRIMARY KEY ((id, user_id), update_time)
) WITH CLUSTERING ORDER BY (update_time DESC);

CREATE TABLE IF NOT EXISTS director (
    id text,
    user_id text,
    display_name text,
    picture text,
    create_time timestamp,
    update_time timestamp,
    active boolean,
    PRIMARY KEY (id, update_time)
) WITH CLUSTERING ORDER BY (update_time DESC);

CREATE TABLE IF NOT EXISTS director_user (
    id text,
    user_id text,
    display_name text,
    picture text,
    create_time timestamp,
    update_time timestamp,
    active boolean,
    PRIMARY KEY ((id, user_id), update_time)
) WITH CLUSTERING ORDER BY (update_time DESC);

CREATE TABLE IF NOT EXISTS movie (
    id text,
    user_id text,
    category_id text,
    director_id text,
    display_name text,
    description text,
    year int,
    picture text,
    watch_url text,
    create_time timestamp,
    update_time timestamp,
    active boolean,
    crawl_url text,
    PRIMARY KEY (id, update_time)
) WITH CLUSTERING ORDER BY (update_time DESC);

CREATE TABLE IF NOT EXISTS movie_user (
     id text,
     user_id text,
     category_id text,
     director_id text,
     display_name text,
     description text,
     year int,
     picture text,
     watch_url text,
     create_time timestamp,
     update_time timestamp,
     active boolean,
     crawl_url text,
     PRIMARY KEY ((id, user_id), update_time)
) WITH CLUSTERING ORDER BY (update_time DESC);

CREATE TABLE IF NOT EXISTS movie_category (
      id text,
      user_id text,
      category_id text,
      director_id text,
      display_name text,
      description text,
      year int,
      picture text,
      watch_url text,
      create_time timestamp,
      update_time timestamp,
      active boolean,
      crawl_url text,
      PRIMARY KEY ((id, category_id), update_time)
) WITH CLUSTERING ORDER BY (update_time DESC);

CREATE TABLE IF NOT EXISTS movie_director (
      id text,
      user_id text,
      category_id text,
      director_id text,
      display_name text,
      description text,
      year int,
      picture text,
      watch_url text,
      create_time timestamp,
      update_time timestamp,
      active boolean,
      crawl_url text,
      PRIMARY KEY ((id, director_id), update_time)
) WITH CLUSTERING ORDER BY (update_time DESC);
