CREATE TABLE IF NOT EXISTS category (
    id varchar(128) NOT NULL,
    user_id varchar(128) NOT NULL,
    display_name varchar(256) NOT NULL,
    create_time timestamptz NOT NULL DEFAULT now(),
    update_time timestamptz NOT NULL DEFAULT now(),
    active boolean DEFAULT TRUE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS movie (
    id varchar(128) NOT NULL,
    user_id varchar(128) NOT NULL,
    category_id varchar(128) DEFAULT NULL,
    display_name varchar(256) NOT NULL,
    description varchar(512) DEFAULT NULL,
    picture varchar(2048) DEFAULT NULL,
    watch_url varchar(2048) DEFAULT NULL,

    create_time timestamptz NOT NULL DEFAULT now(),
    update_time timestamptz NOT NULL DEFAULT now(),
    active boolean DEFAULT TRUE,
    crawl_url varchar(2048) DEFAULT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS watch_queue (
    id varchar(128) NOT NULL,
    user_id varchar(128) NOT NULL,
    create_time timestamptz NOT NULL DEFAULT now(),
    update_time timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS movie_watch_queue (
    movie_id varchar(128) NOT NULL,
    watch_queue_id varchar(128) NOT NULL,
    add_time timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT fk_movie FOREIGN KEY (movie_id) REFERENCES movie(id) ON DELETE CASCADE,
    CONSTRAINT fk_watch_queue FOREIGN KEY (watch_queue_id) REFERENCES watch_queue(id) ON DELETE CASCADE,
    PRIMARY KEY (movie_id, watch_queue_id)
);

CREATE TYPE watch_list_kind AS enum (
    'wish',
    'seen'
);

CREATE TABLE IF NOT EXISTS watch_list (
    id varchar(128) NOT NULL,
    user_id varchar(128) NOT NULL,
    category_id varchar(128) DEFAULT NULL,
    list_kind watch_list_kind NOT NULL DEFAULT 'wish',
    create_time timestamptz NOT NULL DEFAULT now(),
    update_time timestamptz NOT NULL DEFAULT now(),
    active boolean DEFAULT TRUE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS movie_watch_list (
    movie_id varchar(128) NOT NULL,
    watch_list_id varchar(128) NOT NULL,
    add_time timestamptz NOT NULL DEFAULT now(),
    total_times_seen int DEFAULT 0,
    CONSTRAINT fk_movie FOREIGN KEY (movie_id) REFERENCES movie(id) ON DELETE CASCADE,
    CONSTRAINT fk_watch_queue FOREIGN KEY (watch_list_id) REFERENCES watch_list(id) ON DELETE CASCADE,
    PRIMARY KEY (movie_id, watch_list_id)
);
