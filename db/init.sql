
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    content text
);

CREATE TABLE embeddings (
    postId int,
    word_offset int,
    word int
);