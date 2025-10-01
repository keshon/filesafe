package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/keshon/filesafe/translit"
)

func listFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	files := []string{}
	for _, e := range entries {
		if !e.IsDir() {
			files = append(files, e.Name())
		}
	}
	return files, nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

func main() {
	inDir := "in"
	outDir := "out"

	if err := os.MkdirAll(inDir, 0755); err != nil {
		fmt.Println("Ошибка создания папки in:", err)
		return
	}
	if err := os.MkdirAll(outDir, 0755); err != nil {
		fmt.Println("Ошибка создания папки out:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== FileSafe (Go) ===")
	fmt.Println("Файлы берутся из папки:", inDir)
	fmt.Println("Результат будет в папке:", outDir)
	fmt.Println("Доступные команды: list, convert, exit")

	for {
		fmt.Print("\nВведите команду: ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		switch cmd {
		case "list":
			files, err := listFiles(inDir)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
			if len(files) == 0 {
				fmt.Println("Нет файлов в папке", inDir)
				continue
			}
			fmt.Println("Список файлов:")
			for i, f := range files {
				fmt.Printf("[%d] %s -> %s\n", i+1, f, translit.TranslitUniversal(f))
			}

		case "convert":
			files, err := listFiles(inDir)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
			if len(files) == 0 {
				fmt.Println("Нет файлов в папке", inDir)
				continue
			}

			for i, f := range files {
				newName := translit.TranslitUniversal(f)
				src := filepath.Join(inDir, f)
				dst := filepath.Join(outDir, newName)
				fmt.Printf("[%d/%d] %s -> %s\n", i+1, len(files), f, newName)
				if err := copyFile(src, dst); err != nil {
					fmt.Println("Ошибка копирования:", err)
				}
			}
			fmt.Println("Готово! Файлы в", outDir)

		case "exit":
			fmt.Println("Выход.")
			return

		default:
			fmt.Println("Неизвестная команда. Доступные: list, convert, exit")
		}
	}
}
