package gin_handlers

import (
	"github.com/gin-gonic/gin"
)

// GetNikita req.
func (gh GinHandlers) GetNikita(c *gin.Context) {
	c.HTML(200, "hNikita.html", gin.H{
		"Banner": "Mantra",
	})
}

// GetElena req.
func (gh GinHandlers) GetElena(c *gin.Context) {
	c.HTML(200, "hElena.html", gin.H{
		"Banner": "Mantra",
	})
}

// GetMikhail req.
func (gh GinHandlers) GetMikhail(c *gin.Context) {
	c.HTML(200, "hMikhail.html", gin.H{
		"Banner": "Mantra",
	})
}
