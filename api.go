/*
** Copyright (C) 2023-2024 - Champ Clark III <cclark _AT_ k9.io>
**
** This program is free software; you can redistribute it and/or modify
** it under the terms of the GNU General Public License Version 2 as
** published by the Free Software Foundation.  You may not use, modify or
** distribute this program under any other version of the GNU General
** Public License.
**
** This program is distributed in the hope that it will be useful,
** but WITHOUT ANY WARRANTY; without even the implied warranty of
** MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
** GNU General Public License for more details.
**
** You should have received a copy of the GNU General Public License
** along with this program; if not, write to the Free Software
** Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA 02111-1307, USA.
 */

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate_API() gin.HandlerFunc {
	return func(c *gin.Context) {

		api_key := c.GetHeader("API_KEY")

		if api_key != Config.API_Key {
			c.JSON(http.StatusOK, gin.H{"error": "api authentication failed"})

			if Production == false { 

			log.Printf("[ Total Queries: %v | Cached: %v [%v%%]| Not Cached: %v [%v%%] - Authentication failure.\n", Total, Cached, (Cached/Total)*100, NotCached, (NotCached/Total)*100)
			
			}

			c.Abort()
			return
		}
	}
}
