package gin_handlers

import (
	"github.com/gin-gonic/gin"
)

// Bio returns your bio Slogan.
func (gh GinHandlers) Bio(c *gin.Context) {
	c.HTML(200, "bio.html", gin.H{
		"Banner": "Mantra",
	})
}

// GetBio returns your Mantra.
func (gh GinHandlers) GetBio(c *gin.Context) {
	c.HTML(200, "bioGet.html", gin.H{
		"Banner": "Mantra 1",
	})
}
