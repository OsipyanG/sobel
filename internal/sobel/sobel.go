package sobel

import (
	"image/color"
	"math"
	"sobel/internal/img"
)

// ApplySobel применяет фильтр Собела к указанным строкам изображения
func ApplySobel(imgData *img.ImageData, startRow, endRow int) {
	// Операторы Собела для горизонтальных и вертикальных градиентов
	sobelX := [3][3]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
	sobelY := [3][3]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	// Проход по строкам изображения в указанном диапазоне
	for y := startRow; y < endRow; y++ {
		// Проход по столбцам изображения, избегая границ
		for x := 1; x < imgData.Width-1; x++ {
			var gx, gy int
			// Применение операторов Собела для текущего пикселя
			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					// Получение цветовых компонентов пикселя
					r, g, b, _ := imgData.SrcImage.At(x+kx, y+ky).RGBA()
					// Преобразование в оттенок серого
					gray := (r*299 + g*587 + b*114) / 1000
					// Вычисление горизонтального и вертикального градиентов
					gx += int(gray>>8) * sobelX[ky+1][kx+1]
					gy += int(gray>>8) * sobelY[ky+1][kx+1]
				}
			}
			// Вычисление величины градиента
			magnitude := math.Sqrt(float64(gx*gx + gy*gy))
			normalized := uint8(math.Min(255, magnitude))
			// Установка значения пикселя в результирующем изображении
			imgData.DstImage.Set(x, y, color.RGBA{normalized, normalized, normalized, 255})
		}
	}
}
