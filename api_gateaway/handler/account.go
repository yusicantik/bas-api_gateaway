package handler

import (
	"api_gateaway/model"
	"api_gateaway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountInterface interface {
	GetAccount(*gin.Context)
	CreateAccount(*gin.Context)
	UpdateAccount(*gin.Context)
	DeleteAccount(*gin.Context)
	GetBalanceAccount(*gin.Context)
}

type accountImplement struct{}

func NewAccount() AccountInterface {
	return &accountImplement{}
}

func (a *accountImplement) GetAccount(g *gin.Context) {
	queryParam := g.Request.URL.Query()
	name := queryParam.Get("name")

	accounts := []model.Account{}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	q := orm
	if name != "" {
		q = q.Where("name = ?", name)
	}
	result := q.Find(&accounts)
	// result := orm.Find(&accounts, "name = ?", name)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Get Account Successfully",
		"data":    accounts,
	})
}

// type BodyPayloadAccount struct {
// 	AccountID string
// 	Name      string
// 	Address   string
// }

func (a *accountImplement) CreateAccount(g *gin.Context) {
	bodyPayload := model.Account{}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	err := g.BindJSON(&bodyPayload)

	result := orm.Create(&bodyPayload)
	if result.Error != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Create Account Successfully",
		"data":    bodyPayload,
	})
}

func (a *accountImplement) UpdateAccount(g *gin.Context) {
	bodyPayload := model.Account{}

	err := g.BindJSON(&bodyPayload)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	id := g.Param("id")

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	user := model.Account{}
	orm.First(&user, "account_id = ?", id)

	if user.AccountID == "" {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Data Not Found",
		})
		return
	}

	user.Name = bodyPayload.Name
	user.Username = bodyPayload.Username
	orm.Save(&user)

	g.JSON(http.StatusOK, gin.H{
		"message": "Update Account Successfully",
		"data":    user,
	})
}

func (a *accountImplement) DeleteAccount(g *gin.Context) {
	id := g.Param("id")

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()
	result := orm.Where("account_id = ?", id).Delete(&model.Account{})
	if result.Error != nil {
		g.AbortWithError(http.StatusBadRequest, result.Error)
	}
	g.JSON(http.StatusOK, gin.H{
		"message": "Delete Account Successfully",
		"data":    id,
	})
}

type BodyPayloadBalance struct{}

func (a *accountImplement) GetBalanceAccount(g *gin.Context) {
	bodyPayloadBal := BodyPayloadBalance{}

	err := g.BindJSON(&bodyPayloadBal)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Hello Guys this API rest for later",
	})
}
