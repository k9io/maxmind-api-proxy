/*
** Copyright (C) 2023 - Champ Clark III <dabeave _AT_ gmail.com>
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
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var Cached float64
var NotCached float64
var Total float64

func Maxmind_Query_IP(c *gin.Context) {

	var body []byte
	var results string

	ip_address := c.Params.ByName("ip_address")

	body = Redis_Check_Cache(ip_address)

	Total++

	if body == nil {

		results = ip_address + " was not cached"
		NotCached++

		client := &http.Client{}

		connect_url := fmt.Sprintf("%s/%s", Config.Maxmind_Url, ip_address)

		req, _ := http.NewRequest("GET", connect_url, nil)
		req.Header.Add("Authorization", "Basic "+basicAuth(Config.Maxmind_Username, Config.Maxmind_Password))
		resp, _ := client.Do(req)

		body, _ = io.ReadAll(resp.Body) // DEBUG : err check!

		str_body := string(body)

		if strings.Contains(str_body, "\"error\":") {

			/* Got an error,  what to do with it? */

			if Config.Cache_Errors == true {

				Redis_Store_Cache(str_body, ip_address)

			} else {

				log.Printf("Returned error for %s, not caching", ip_address)

			}

		} else {

			/* Store as normal */

			Redis_Store_Cache(str_body, ip_address)

		}

	} else {

		results = ip_address + " pull from cache."
		Cached++

	}

	log.Printf("[ Total Queries: %v | Cached: %v [%v%%]| Not Cached: %v [%v%%] - %s\n", Total, Cached, (Cached/Total)*100, NotCached, (NotCached/Total)*100, results)

	c.String(http.StatusOK, string(body))

}
