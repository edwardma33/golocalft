package file

import (
	"localft/internal/db"
	"time"
)

type FileModel struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Extension string    `json:"extension"`
	MimeType  string    `json:"mimeType"`
	Size      int64     `json:"size"`
	CreatedAt time.Time `json:"createdAt"`
}

func (f FileModel) Save() error {
	db, err := db.GetDb()
	if err != nil {
		return err
	}

	q := `
		INSERT into files (name, extension, mime_type, size, created_at)
		VALUES
		($1, $2, $3, $4, $5);
	`
	_, err = db.Exec(q, f.Name, f.Extension, f.MimeType, f.Size, f.CreatedAt.UnixMilli())

	return err
}

func GetFiles() ([]FileModel, error) {
	var files []FileModel
	db, err := db.GetDb()
	if err != nil {
		return files, err
	}

	rows, err := db.Query("SELECT * FROM files;")
	if err != nil {
		return files, err
	}
	for rows.Next() {
		var f FileModel
		var ms int64
		rows.Scan(&f.Id, &f.Name, &f.Extension, &f.MimeType, &f.Size, &ms)
		f.CreatedAt = time.UnixMilli(ms)
		files = append(files, f)
	}

	return files, nil
}

func GetFileById(id int) (FileModel, error) {
	var f FileModel
	db, err := db.GetDb()
	if err != nil {
		return f, err
	}

	row := db.QueryRow("SELECT * FROM files WHERE id = $1", id)
	var ms int64
	err = row.Scan(&f.Id, &f.Name, &f.Extension, &f.MimeType, &f.Size, &ms)
	f.CreatedAt = time.UnixMilli(ms)
	return f, err
}

func DeleteFileById(id int) error {
	db, err := db.GetDb()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM files WHERE id = $1;`", id)
	return err
}
