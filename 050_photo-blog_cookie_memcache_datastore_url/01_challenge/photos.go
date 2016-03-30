package mem

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

func uploadPhoto(src multipart.File, id string, req *http.Request) error {
	defer src.Close()
	fName := getSha(src) + ".jpg"
	return addPhoto(fName, id, req)
}

func addPhoto(fName string, id string, req *http.Request) error {

	// DATASTORE
	md, err := retrieveDstore(id, req)
	if err != nil {
		return err
	}
	md.Pictures = append(md.Pictures, fName)
	err = storeDstore(md, id, req)
	if err != nil {
		log.Println("ERROR addPhoto storeDstore", err)
		return err
	}

	// MEMCACHE
	var mc model
	mc, err = retrieveMemc(id, req)
	if err != nil {
		log.Println("ERROR addPhoto retrieveMemc", err)
		return err
	}
	mc.Pictures = append(mc.Pictures, fName)
	err = storeMemc(mc, id, req)
	if err != nil {
		log.Println("ERROR addPhoto storeMemc", err)
		return err
	}

	return nil
}

func getSha(src multipart.File) string {
	h := sha1.New()
	io.Copy(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}
