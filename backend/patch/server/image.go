package server

import (
	"app/sqlmanager"
	"encoding/json"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

type ImgSQL struct {
	ImgURL  string `json:"imgURL"`
	ImgName string `json:"imgName"`
	ID      string `json:"id"`
}

func (i *ImgSQL) Value() (string, error) {
	b, e := json.Marshal(i)
	return string(b), e
}
func PatchImage(ctx *gin.Context) {
	log.Println("image")
	var store IIMG
	if err := ctx.ShouldBindWith(&store, binding.FormMultipart); err != nil {
		panic(err)
	}

	file := store.File
	cleanedName := strings.ReplaceAll(file.Filename, " ", "_")

	savePath := "./server/uploads/" + cleanedName

	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to save file"})
		return
	}

	// Respond with file view URL and save path
	viewURL := "/uploads/" + cleanedName

	// note: in-order to set the image to the client you have to save it on your disk first
	imageURL := "http://localhost:" + port + viewURL

	var save_ ImgSQL
	save_.ID = *store.ID
	save_.ImgURL = imageURL
	save_.ImgName = cleanedName

	userID := *store.ID
	_uuid, _ := uuid.Parse(userID)
	BUID, _ := _uuid.MarshalBinary()

	query := "UPDATE _nygpatch_ SET  img = ?  WHERE id = ?"

	// re := regexp.MustCompile(`\s+`)

	// _name := re.ReplaceAllString(*store.ImageName, "")
	// store.ImageName = &_name
	x, _ := save_.Value()
	//
	//	update _nygpatch_ set profile=JSON_SET(profile,"$.point","1") where id =UUID_TO_BIN('63922f07-422d-4c2a-afed-652e47328fba');
	m := sqlmanager.ConnectSQL{}

	cfg := Env()
	db, err := m.Init("nygpatch", "_nygpatch_", cfg)
	defer db.CloseDB()

	if err != nil {
		panic(err)
	}
	db.Prepare(query, x, BUID)

}
