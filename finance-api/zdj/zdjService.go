package zdj

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/FelixAnna/web-service-dlw/common/filesystem"
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj/repository"
	"github.com/gin-gonic/gin"
)

var repo repository.ZdjRepo

func init() {
	repo = &repository.ZdjSqlServerRepo{}
}

func GetAll(c *gin.Context) {
	results, _ := repo.Search(&entity.Criteria{Page: 1, Size: 20})
	c.JSON(http.StatusOK, results)
}

func Search(c *gin.Context) {
	//get result from somewhere
	var criteria entity.Criteria
	if err := c.BindJSON(&criteria); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	results, _ := repo.Search(&criteria)

	c.JSON(http.StatusOK, results)
}

func MemoryCosty(c *gin.Context) {
	//get result from somewhere
	times := c.DefaultQuery("times", "1000000")
	itimes, err := strconv.ParseInt(times, 10, 32)
	if err != nil {
		itimes = 100000
	}

	results := make(map[int]int, itimes)

	for i := 1; i <= int(itimes); i++ {
		if i <= 2 {
			results[i] = i
			continue
		}

		results[i] = results[i-1] + results[i-2]
	}

	c.JSON(http.StatusOK, results)
}

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	version := c.DefaultQuery("version", "2021")
	iversion, err := strconv.ParseInt(version, 10, 32)
	if err != nil {
		iversion = 2021
	}

	log.Println(file.Filename)

	tempPath := getTempPath()
	c.SaveUploadedFile(file, tempPath)
	defer os.Remove(tempPath)

	lines := filesystem.ReadLines(tempPath)

	//convert to model list
	models := parseModel(lines, int(iversion))

	//save to somewhere
	err = repo.Append(&models)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "Data uploaded.")
}

func Delete(c *gin.Context) {
	sid := c.Param("id")
	sversion := c.DefaultQuery("version", "2021")
	id, err := strconv.ParseInt(sid, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	version, err := strconv.ParseInt(sversion, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = repo.Delete(int(id), int(version))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, "Data deleted.")
}

func getTempPath() string {
	path := os.TempDir() + "\\" + time.Now().Format("20000102132435") + ".txt"

	return path
}

func parseModel(lines []string, version int) []entity.Zhidaojia {
	start := 0
	for lines[start] != "1" {
		start++
	}

	results := make([]entity.Zhidaojia, 0)
	for start < len(lines) {
		content := lines[start]
		if content[0] == '-' && content[len(content)-1] == '-' {
			start++
			continue
		}

		model := buildModel(lines[start : start+5]...)
		model.Version = version
		results = append(results, model)
		start = start + 5
	}

	return results
}

func buildModel(values ...string) entity.Zhidaojia {
	model := entity.Zhidaojia{}

	id, _ := strconv.ParseInt(values[0], 10, 32)
	price, _ := strconv.ParseInt(values[4], 10, 32)

	model.Id = int(id)
	model.Distrct = values[1]
	model.Street = values[2]
	model.Community = values[3]
	model.Price = int(price)

	return model
}
