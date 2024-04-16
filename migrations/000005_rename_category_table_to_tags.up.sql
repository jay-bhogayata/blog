ALTER TABLE articles DROP CONSTRAINT articles_category_id_fkey;

ALTER TABLE categories RENAME TO old_categories;

CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

INSERT INTO tags (id, name) SELECT id, name FROM old_categories;

DROP TABLE old_categories;

ALTER TABLE articles RENAME COLUMN category_id TO tag_id;
ALTER TABLE articles ADD CONSTRAINT articles_tag_id_fkey FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE SET NULL;