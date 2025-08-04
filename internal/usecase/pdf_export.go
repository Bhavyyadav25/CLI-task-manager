package usecase

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Bhavyyadav25/CLI-task-manager/internal/repo"
	"github.com/jung-kurt/gofpdf"
)

type PDFExporter struct {
	repo repo.TaskRepository
}

func NewPDFExporter(r repo.TaskRepository) *PDFExporter {
	return &PDFExporter{repo: r}
}

func (e *PDFExporter) Export(filePath string) error {
	lower := strings.ToLower(filePath)
	if !strings.HasSuffix(lower, ".pdf") {
		return fmt.Errorf("invalid file extension: %s", lower, " expected .pdf")
	}

	tasks, err := e.repo.List()
	if err.Message != "" {
		return fmt.Errorf("list tasks: %w", err)
	}

	fontPath := filepath.Join("assets", "fonts", "DejaVuSans.ttf")
	if info, ferr := os.Stat(fontPath); ferr != nil {
		return fmt.Errorf("font not found at %s: %w", fontPath, ferr)
	} else if info.IsDir() {
		return fmt.Errorf("expected font file but found directory at %s", fontPath)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetTitle("Task Overview", false)
	pdf.SetMargins(20, 20, 20)
	pdf.AddPage()

	pdf.AddUTF8Font("DejaVu", "", fontPath)

	pdf.SetFont("DejaVu", "", 18)
	pdf.SetTextColor(30, 30, 30)
	pdf.CellFormat(0, 10, "My Task Overview", "", 1, "C", false, 0, "")
	pdf.Ln(4)

	colW := []float64{20, 150}
	pdf.SetFont("DejaVu", "", 12)
	pdf.SetFillColor(200, 200, 200)
	pdf.SetDrawColor(150, 150, 150)
	pdf.CellFormat(colW[0], 8, "Status", "1", 0, "C", true, 0, "")
	pdf.CellFormat(colW[1], 8, "Task", "1", 0, "C", true, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("DejaVu", "", 11)
	fill := false
	for _, t := range tasks {
		if fill {
			pdf.SetFillColor(245, 245, 245)
		} else {
			pdf.SetFillColor(255, 255, 255)
		}
		status := "☐"
		if t.Done {
			status = "✓"
		}
		pdf.CellFormat(colW[0], 7, status, "1", 0, "C", true, 0, "")
		pdf.CellFormat(colW[1], 7, t.Description, "1", 0, "L", true, 0, "")
		pdf.Ln(-1)
		fill = !fill
	}

	pdf.Ln(5)
	pdf.SetFont("DejaVu", "", 10)
	pdf.SetTextColor(100, 100, 100)
	pdf.CellFormat(0, 5, "Generated: "+time.Now().Format("2006-01-02 15:04:05"), "", 0, "R", false, 0, "")

	if err := pdf.OutputFileAndClose(filePath); err != nil {
		return fmt.Errorf("save PDF: %w", err)
	}
	return nil
}
