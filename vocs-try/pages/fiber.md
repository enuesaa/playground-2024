# fiber
example app of go fiber.

## Links
- https://vocs.dev/docs/markdown
- https://docs.gofiber.io

## Code
:::code-group
```go [foucs]
package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New() // [!code focus]

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":3000") // [!code focus]
}
```

```go [highlight] showLineNumbers
package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New() // [!code hl]

	app.Get("/", func(c *fiber.Ctx) error {
		// [!code word:SendString]
		return c.SendString("Hello, World!")
	})
	app.Listen(":3000") // [!code hl]
}
```
:::

:::details[See more]
hello
:::
