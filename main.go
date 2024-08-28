package main

import (
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Cmd struct {
	Ccmd string `json:"cmd"`
}

var data = make(map[string]string)

func main() {
	// rander := bufio.NewReader(os.Stdin)
	app := fiber.New()

	app.Post("/", db)

	app.Listen(":8080")

}

func db(c *fiber.Ctx) error {
	body := c.BodyRaw()
	cmd := Cmd{}
	json.Unmarshal(body, &cmd)
	trim := strings.TrimSpace(cmd.Ccmd)
	split := strings.Split(trim, " ")
	switch split[0] {
	case "set":
		{
			if len(split) < 3 {
				return c.SendString("enter <key> & <value>")
			}
			join := strings.Join(split[2:], " ")
			data[split[1]] = join
			return c.SendString("successfully set <key> & <value>")
		}
	case "get":
		{
			if len(split) < 2 {
				return c.SendString("enter <key>")
			}
			if val, ok := data[split[1]]; ok {
				return c.SendString(val)
			} else {
				return c.SendString("key not found")
			}
		}
	case "delete":
		{
			if len(split) < 2 {
				return c.SendString("enter <key>")
			}
			if _, ok := data[split[1]]; ok {
				delete(data, split[1])
				return c.SendString("key deleted successfully")
			} else {
				return c.SendString("key not found")
			}
		}
	default:
		{
			return c.SendString("unknown command")
		}
	}
}
