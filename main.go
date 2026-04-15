package main

import (
	"fmt"
	"io"
	"localft/internal/db"
	"localft/internal/file"
	"localft/internal/utils"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	const uploadPath = "uploads"
	const port = 8423
	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("./client/dist"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./client/dist/index.html")
	})

	apiRouter := chi.NewRouter()

	apiRouter.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		gb := int64(1024 << 30)
		r.ParseMultipartForm(gb)

		f, metadata, err := r.FormFile("file")
		if err != nil {
			log.Println(err)
			utils.WriteJson(w, 500, utils.JsonMap{
				"error": "error finding file in form data",
			})
			return
		}
		defer f.Close()

		data, err := io.ReadAll(f)
		if err != nil {
			log.Println(err)
			utils.WriteJson(w, 500, utils.JsonMap{
				"error": "error reading file",
			})
			return
		}

		tkns := strings.Split(metadata.Filename, ".")
		extension := "." + strings.ToLower(tkns[len(tkns)-1])

		name := r.FormValue("file-name")
		if len(name) == 0 {
			name = tkns[0]
		}

		path := path.Join(uploadPath, name + extension)
		fmt.Println(path)

		upload, err := os.Create(path)
		_, err = upload.Write(data)
		if err != nil {
			log.Println(err)
			utils.WriteJson(w, 500, utils.JsonMap{
				"error": "error creating or writing to upload file",
			})
			return
		}

		fileModel := file.FileModel{
			Name: name,
			Extension: extension,
			MimeType: metadata.Header.Get("Content-Type"),
			Size: metadata.Size,
			CreatedAt: time.Now(),
		}

		err = fileModel.Save()
		if err != nil {
			log.Println(err)
			utils.WriteJson(w, 500, utils.JsonMap{
				"error": "error saving file",
			})
			return
		}

		utils.WriteJson(w, 200, utils.JsonMap{
			"message": "upload successful",
			"name": name,
		})
	})

	apiRouter.Get("/files", func(w http.ResponseWriter, r *http.Request) {
		files, err := file.GetFiles()
		if err != nil {
			log.Println(err)
			utils.WriteJson(w, 500, utils.JsonMap{
				"error": "error getting files from db",
			})
			return
		}

		utils.WriteJson(w, 200, utils.JsonMap{
			"message": "fetched files successfully",
			"files": files,
		})
	})

	apiRouter.Get("/download/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println(err)
			utils.WriteJson(w, 500, utils.JsonMap{"error": "error parsing id"})
			return
		}

		fileModel, err := file.GetFileById(id)
		if err != nil {
			log.Println(err)
			utils.WriteJson(w, 500, utils.JsonMap{"error": "error getting file from db"})
			return
		}

		http.ServeFile(w, r, path.Join(uploadPath, fileModel.Name + fileModel.Extension))
	})

	apiRouter.Get("/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println(err)
			utils.WriteJson(w, 500, utils.JsonMap{"error": "error parsing id"})
			return
		}

		fileModel, err := file.GetFileById(id)
		if err != nil {
			log.Println(err)
			utils.WriteJson(w, 500, utils.JsonMap{"error": "error deleting file db entry"})
			return
		}

		err = os.Remove(path.Join(uploadPath, fileModel.Name + fileModel.Extension))
		if err != nil {
			log.Println(err)
			utils.WriteJson(w, 500, utils.JsonMap{"error": "error removing file"})
			return
		}

		file.DeleteFileById(id)

		utils.WriteJson(w, 200, utils.JsonMap{"message": "deleted file successfully"})
	})

	r.Mount("/api", apiRouter)
	
	fmt.Println("running @ http://localhost:8423")
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
