package exception

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/utils/validation"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

type ResponseError struct {
	Errors *any `json:"errors,omitempty"`
}

type ResponseSuccess struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func Err(c *fiber.Ctx, err error) error {
	switch err.(type) {
	case *RespError:
	case *json.UnmarshalTypeError:
		message := fmt.Sprintf("type field must be %s", err.(*json.UnmarshalTypeError).Type.String())
		err = UnprocessableEntity(map[string]map[string]string{
			err.(*json.UnmarshalTypeError).Field: {
				"unprocess_entity": message,
			},
		})
	case *json.SyntaxError:
		err = BadRequest(map[string][]string{
			"unexpected": {
				"unexpected end of JSON input",
			},
		})
	case *pq.Error:
		err = InternalServerError(err.Error())
	case validator.ValidationErrors:
		mapBad := make(map[string][]string)
		for _, e := range err.(validator.ValidationErrors) {
			mapBad[e.Field()] = append(mapBad[e.Field()], validation.MsgForTag(e.Tag(), e.Param()))
		}
		err = BadRequest(mapBad)
	default:
		if errors.Is(err, context.DeadlineExceeded) {
			err = RequestTimeOut("request time out")
		} else {
			err = InternalServerError(err.Error())
		}
	}
	return c.Status(err.(*RespError).Code).JSON(ResponseError{Errors: &err.(*RespError).Message})
}
