package routes
import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func MainRoutes(r *gin.Engine) *gin.Engine{
	r.GET("/", func(c *gin.Context){
		// c.String(200, "!Hello World!")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/saludo/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		// c.String(200, "Hola %s", nombre)
		
		c.JSON(http.StatusOK,gin.H{
			"message" : "Hola " + nombre,
		})
	})

	return r
}