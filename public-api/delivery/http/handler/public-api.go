package handler

import (
	"be-service-public-api/domain"
	"be-service-public-api/helper"
	"strconv"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

type PublicHandler struct {
	PublicAPIUseCase   domain.PublicAPIUseCase
	PublicAPIMySQLRepo domain.PublicAPIMySQLRepo
}

func (ph *PublicHandler) GetAllProduct(c *fiber.Ctx) error {
	var input domain.RequestAdditionalData
	search := c.Query("search")
	if search != "" {
		input.NameSearch = &search
	}

	limit := c.Query("limit")
	if limit == "" {
		limitInt := viper.GetInt("database.default_limit_query")
		log.Info("Parameter limit not set, running with default limit query from config", limitInt)

		input.Limit = &limitInt
	} else {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
		}
		input.Limit = &limitInt
	}

	page := c.Query("page")
	if page == "" {
		pageInt := viper.GetInt("database.default_page")
		log.Info("Parameter limit not set, running with default limit query from config", pageInt)

		input.Page = &pageInt
	} else {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
		}
		input.Page = &pageInt
	}

	order := c.Query("order")
	if order == "" {
		input.Order = nil
	} else {
		input.Order = &order
	}
	res, err := ph.PublicAPIUseCase.GetAllProduct(c.Context(), input)
	if err != nil {
		if err.Error() == "Not found" {
			return helper.HttpSimpleResponse(c, fasthttp.StatusNotFound)
		}
		return err
	}

	return c.Status(fasthttp.StatusOK).JSON(res)
}

func (ph *PublicHandler) PostCheckout(c *fiber.Ctx) (err error) {
	var input domain.RequestDataCheckout
	err = c.BodyParser(&input)
	if err != nil {
		log.Errorf(err.Error())
		return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	err = ph.PublicAPIUseCase.PostCheckout(c.Context(), input)
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = Stok not found" {
			return c.Status(fasthttp.StatusBadRequest).SendString("Out of stok")
		}

		if err.Error() == "rpc error: code = Unknown desc = Your transaction not found" {
			return c.Status(fasthttp.StatusBadRequest).SendString("Your transaction not found")
		}
		return err
	}

	// log.Info(input)
	return c.SendStatus(200)
}

func (ph *PublicHandler) GetProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		log.Error(err)
		return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	res, err := ph.PublicAPIUseCase.GetProduct(c.Context(), int(id))

	if err != nil {
		if err.Error() == "Data not found" {
			return c.Status(fasthttp.StatusOK).JSON(err.Error())
		}
		return helper.HttpSimpleResponse(c, fasthttp.StatusInternalServerError)

	}

	return c.Status(fasthttp.StatusOK).JSON(res)
}

func (ph *PublicHandler) CheckStok(c *fiber.Ctx) (err error) {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		log.Error(err)
		return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	err = ph.PublicAPIUseCase.CheckStok(c.Context(), int32(int(id)))
	if err != nil {
		return c.Status(fasthttp.StatusNotFound).SendString("Out of stok")
	}
	return c.SendStatus(fasthttp.StatusOK)
}
