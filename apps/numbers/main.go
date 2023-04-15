package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mholtzscher/fun-with-monorepos/math"
	"strconv"
)

func main() {
	r := gin.Default()
	r.GET("/factorial/:num", func(c *gin.Context) {
		num, err := strconv.Atoi(c.Param("num"))
		if err != nil {
			c.JSON(400, gin.H{
				"error": "invalid number",
			})
			return
		}
		c.JSON(200, gin.H{
			"result": math.Factorial(num),
		})
	})
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
