package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func downloadVideo(videoURL string) error {
    // Constrói o comando para baixar o vídeo em formato MP4
    cmd := exec.Command("yt-dlp", "-f", "bestvideo[ext=mp4]+bestaudio[ext=m4a]/mp4", videoURL, "-o", "%(title)s.%(ext)s")
    
    // Define a saída padrão e a saída de erros para os.Stdout para imprimir no console em tempo real
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stdout

    // Executa o comando
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("yt-dlp failed: %s", err)
    }

    return nil
}

func main() {
    // Solicita ao usuário o URL do vídeo
    fmt.Println("Enter the YouTube video URL:")
    var videoURL string
    if _, err := fmt.Scanln(&videoURL); err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    // Verifica se a URL é válida (neste exemplo, verifica se não está vazia)
    if strings.TrimSpace(videoURL) == "" {
        fmt.Println("Invalid URL")
        return
    }

    // Chama a função de download
    if err := downloadVideo(videoURL); err != nil {
        fmt.Println("Error:", err)
    }
}
