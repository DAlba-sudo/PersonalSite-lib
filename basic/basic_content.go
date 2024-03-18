package basic 

import (
	"database/sql"
)

type BasicContent struct {
    Id int
    Title string
    Desc string
    Body string
    Type int
    Hash string
}

func GetFromHash(db *sql.DB, hash string) (*BasicContent, error) {
    content := new(BasicContent) 
    query := "SELECT * FROM BasicContent WHERE Hash == ? LIMIT 1"
    rows, err := db.Query(query, hash)
    if err != nil {
        return nil, err
    }

    if rows.Next() {
        content.FromRow(rows)
    }

    return content, nil
}

func (content *BasicContent) Sync(db *sql.DB) error {
    if err := db.Ping(); err != nil {
        return err
    }

    exec := "INSERT INTO BasicContent (Id, Title, Desc, Body, Type, Hash) VALUES (?, ?, ?, ?, ?, ?)"
    _, err := db.Exec(exec, content.Id, content.Title, content.Desc, content.Body, content.Type, content.Hash)
    if err != nil {
        return err
    }
    return nil
}

func (content *BasicContent) FromRow(row *sql.Rows) error {
    err := row.Scan(
        &content.Id,
        &content.Title,
        &content.Desc,
        &content.Body,
        &content.Type,
        &content.Hash,
    )
    if err != nil {
        return err
    }

    return nil
}

func GetId(db *sql.DB, id int) *BasicContent {
    content := &BasicContent{}
    query := "SELECT * FROM BasicContent WHERE ID = ?"

    rows, err := db.Query(query, id)
    if err != nil {
        return nil
    }

    if rows.Next() {
        err := content.FromRow(rows)
        if err != nil {
            return nil
        }
    }

    return content
}

func GetAll(db *sql.DB) ([]*BasicContent, error) {
    contents := make([]*BasicContent, 0)

    query := "SELECT * FROM BasicContent"
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }

    for rows.Next() {
        content := &BasicContent{}
        err = content.FromRow(rows)
        if err != nil {
            return contents, err
        }

        contents = append(contents, content)
    }

    
    return contents, nil
}
