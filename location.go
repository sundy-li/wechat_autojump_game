package wechat_autojump_game

import (
	"image"
	"math"
)

type location struct {
	img image.Image
}

func NewLocation(img image.Image) *location {
	return &location{img: img}
}

var (
	grade float64 = 255.0 / 452.0
)

//根据阈值找到中心点
func (l *location) judgeByRGBRange() (currentX, currentY, nextX, nextY int) {
	bounds := l.img.Bounds()
	maxY := 0
	minX := 100000
	// r: 38-51 , G: 38-48 ,  B:64-73
	for y := bounds.Max.Y - 1; y >= 320; y-- {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := l.img.At(x, y).RGBA()
			if rgbInterval(r, g, b, 38, 65, 38, 53, 70, 91) {
				maxY = max(maxY, y)
				minX = min(minX, x)
			}
		}
	}
	currentY = maxY - 16
	if currentY == 0 {
		return
	}
	currentX = minX + 40
	for y := 320; y < currentY; y++ {

		r, g, b, _ := l.img.At(0, y).RGBA()
		r >>= 8
		g >>= 8
		b >>= 8
		if nextX != 0 {
			break
		}
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			//需要排除在棋子X范围内的像素点
			if x >= currentX-40 && x <= currentX+40 {
				continue
			}

			r1, g1, b1, _ := l.img.At(x, y).RGBA()
			r1 >>= 8
			g1 >>= 8
			b1 >>= 8
			if colorDiff(int(r), int(g), int(b), int(r1), int(g1), int(b1)) >= 30 {
				nextX = x
				break
			}
		}
	}
	nextY = currentY - int(math.Abs(float64(nextX-currentX)*grade))
	return
}

func getDistance(currentX, currentY, nextX, nextY int) float64 {
	return math.Sqrt(math.Pow(float64(currentX-nextX), 2) + math.Pow(float64(currentY-nextY), 2))
}

func colorDiff(r, g, b, r1, g1, b1 int) float64 {
	return math.Sqrt(math.Pow(float64(r1-r), 2) + math.Pow(float64(g1-g), 2) + math.Pow(float64(b1-b), 2))
}

func rgbInterval(r, g, b uint32, rmin, rmax, gmin, gmax, bmin, bmax int) bool {
	return rmin <= int(r>>8) && int(r>>8) <= rmax && gmin <= int(g>>8) && int(g>>8) <= gmax && bmin <= int(b>>8) && int(b>>8) <= bmax
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
