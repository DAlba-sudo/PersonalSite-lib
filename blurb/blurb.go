package blurb

import "database/sql"

type ContentBlurb struct {
    Id int 
    Title string
    Slug string
    Hash string
}

func (c *ContentBlurb) FromRow(row *sql.Rows) error {
    err := row.Scan(
        &c.Id,
        &c.Title,
        &c.Slug,
        &c.Hash,
    )

    return err
}

func (c *ContentBlurb) Sync(db *sql.DB) error {
    if err := db.Ping(); err != nil {
        return err
    }

    exec := "INSERT INTO ContentBlurb (Title, Slug, Hash) VALUES (?, ?, ?)"
    _, err := db.Exec(exec, c.Title, c.Slug, c.Hash)
    if err != nil {
        return err
    }

    return nil
}

func GetFromSlug(db *sql.DB, slug string) (*ContentBlurb, error) {
    blurb := new(ContentBlurb)

    query := "SELECT * FROM ContentBlurb WHERE Slug == ? LIMIT 1"
    rows, err := db.Query(query, slug)
    if err != nil {
        return nil, err
    }

    if rows.Next() {
        blurb.FromRow(rows)
    }

    return blurb, nil
}

func GetAll(db *sql.DB, limit int) ([]*ContentBlurb, error) {
    blurbs := make([]*ContentBlurb, 0)
    if err := db.Ping(); err != nil {
        return nil, err
    }

    query := "SELECT * FROM ContentBlurb LIMIT ?;"
    rows, err := db.Query(query, limit)
    if err != nil {
        return nil, err
    }

    for rows.Next() {
        blurb := &ContentBlurb{}
        if err := blurb.FromRow(rows); err != nil {
            return nil, err
        }

        blurbs = append(blurbs, blurb)
    }

    return blurbs, nil
}
