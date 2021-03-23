package main

// type (
// 	// CoordPoint - coordinates of click
// 	CoordPoint struct {
// 		X float64 `json:"x"`
// 		Y float64 `json:"y"`
// 	}

// 	// CoordPicSize - size of picture
// 	CoordPicSize struct {
// 		MapWidth  float64 `json:"mapWidth"`
// 		MapHeight float64 `json:"mapHeight"`
// 	}
// )

// var (
// 	// CoordPointVar - point object
// 	CoordPointVar = CoordPoint{}

// 	// CoordPicSizeVar - picture object
// 	CoordPicSizeVar = CoordPicSize{}

// 	// SPoint - point object that have standartised values
// 	SPoint = CoordPoint{}
// )

// // MapSize - updates map size and returns
// func MapSize(c *gin.Context) {
// 	mapWidth, err := strconv.ParseFloat(c.PostForm("mapWidth"), 64)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, "Failed to check map size")
// 		return
// 	}
// 	mapHeight, err := strconv.ParseFloat(c.PostForm("mapHeight"), 64)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, "Failed to check map size")
// 		return
// 	}

// 	CoordPicSizeVar.MapWidth = mapWidth
// 	CoordPicSizeVar.MapHeight = mapHeight

// 	c.JSON(http.StatusOK, gin.H{
// 		"mapWidth":  mapWidth,
// 		"mapHeight": mapHeight,
// 	})
// }

// // Coords - fills coords objects
// func Coords(c *gin.Context) {
// 	x, err := strconv.ParseFloat(c.PostForm("x"), 64)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, "Failed to check click position")
// 		return
// 	}
// 	y, err := strconv.ParseFloat(c.PostForm("y"), 64)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, "Failed to check click position")
// 		return
// 	}

// 	CoordPointVar.X = x
// 	CoordPointVar.Y = y

// 	sX := (100 * x) / CoordPicSizeVar.MapWidth
// 	sY := (100 * y) / CoordPicSizeVar.MapHeight
// 	SPoint.X = math.Round(sX*100) / 100
// 	SPoint.Y = math.Round(sY*100) / 100

// 	//fmt.Println(CoordPointVar, CoordPicSizeVar, SPoint)

// 	c.JSON(http.StatusOK, nil)
// }
