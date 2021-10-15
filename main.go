package main


import (
    "net/http"
    "net/url"	
    "github.com/gin-gonic/gin"
)

type member struct{
	ID string `json:"id"`
	Name string `json:"name"`
	FavoriteTeam string `json:"team"`
}

// members slice to seed record album data.
var members = []member{
    {ID: "1", Name: "Rowland", FavoriteTeam : "Mumbai Indians"},
    {ID: "2", Name: "Neel", FavoriteTeam : "Chennai Super Kings"},
    {ID: "3", Name: "Kunal", FavoriteTeam : "Rajasthan Royals"},
    {ID: "4", Name: "Tanuj", FavoriteTeam : "Sunrisers Hyderabad"},
    {ID: "5", Name: "Daye", FavoriteTeam : "Chennai super Kings"},
}

func main() {
    router := gin.Default()
    router.GET("/members", getMembers)
    router.POST("/members", postMembers)
    router.GET("/members/:id", getMemberByID)
    router.PUT("/members/:id/:team", editMemberByID)
    router.DELETE("/members/:id", deleteMemberByID)
    router.Run("localhost:8080")
}
// getMembers responds with the list of all members as JSON.
func getMembers(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, members)
}

func postMembers(c *gin.Context) {
    var newMember member

    // Call BindJSON to bind the received JSON to newMember
    if err := c.BindJSON(&newMember); err != nil {
        return
    }

    // Add the new member to the slice.
    members = append(members, newMember)
    c.IndentedJSON(http.StatusCreated, newMember)
}

func getMemberByID(c *gin.Context) {
	id := c.Param("id")
	
	for _, a:= range members{
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return	
		}
		
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"member not found"})
}

func editMemberByID(c *gin.Context) {
	id := c.Param("id")
	team:= c.Param("team")
	var i int = 0
	for _, a:= range members{
		if a.ID == id {
			a.FavoriteTeam = team 
			c.IndentedJSON(http.StatusOK, a)
			decodedString, err := url.QueryUnescape(team)
			members[i].FavoriteTeam = decodedString
		if err != nil {
			
		}
			return	
		}
		i++
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"member not found"})
}

func deleteMemberByID(c *gin.Context) {
	id := c.Param("id")
	
	var i int = 0
	for _, a:= range members{
		if a.ID == id {			
			c.IndentedJSON(http.StatusOK, a)
			members = RemoveIndex(members, i)
		}
		i++
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"member not found"})
}

func RemoveIndex(s []member, index int) []member {
	return append(s[:index], s[index+1:]...)
}
