ALTER TABLE articles DROP CONSTRAINT articles_tag_id_fkey;
ALTER TABLE articles RENAME COLUMN tag_id TO category_id;

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

INSERT INTO categories (id, name)
SELECT id, name FROM tags;

DROP TABLE tags;

ALTER TABLE articles ADD CONSTRAINT articles_category_id_fkey FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL;