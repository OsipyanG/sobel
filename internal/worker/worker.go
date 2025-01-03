package worker

import (
	"sobel/internal/img"
	"sobel/internal/sobel"
	"sync"
)

type SobelTask struct {
	StartRow int
	EndRow   int
}

func worker(tasks chan SobelTask, imgData *img.ImageData) {
	for task := range tasks {
		sobel.ApplySobel(imgData, task.StartRow, task.EndRow)
	}
}

func RunWorkerPool(numWorkers int, imgData *img.ImageData) {
	tasks := make(chan SobelTask, numWorkers)

	wg := sync.WaitGroup{}
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			worker(tasks, imgData)
			wg.Done()
		}()
	}

	rowsPerWorker := imgData.Height / numWorkers
	for i := 0; i < numWorkers; i++ {
		startRow := i * rowsPerWorker
		endRow := startRow + rowsPerWorker
		if i == numWorkers-1 {
			endRow = imgData.Height
		}
		tasks <- SobelTask{StartRow: startRow, EndRow: endRow}
	}
	close(tasks)

	wg.Wait()
}
