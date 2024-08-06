package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jlaffaye/ftp"
)

func main() {
	// Replace with your FTP server details
	ftpServer := "ftp.example.com:21"
	username := "your-username"
	password := "your-password"

	// Connect to the FTP server
	conn, err := ftp.Dial(ftpServer)
	if err != nil {
		log.Fatalf("Failed to connect to FTP server: %v", err)
	}
	defer conn.Quit()

	// Login to the FTP server
	err = conn.Login(username, password)
	if err != nil {
		log.Fatalf("Failed to login to FTP server: %v", err)
	}

	// Upload a file
	err = uploadFile(conn, "local-file.txt", "remote-file.txt")
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}
	fmt.Println("File uploaded successfully")

	// Download a file
	err = downloadFile(conn, "remote-file.txt", "downloaded-file.txt")
	if err != nil {
		log.Fatalf("Failed to download file: %v", err)
	}
	fmt.Println("File downloaded successfully")
}

func uploadFile(conn *ftp.ServerConn, localFilePath, remoteFilePath string) error {
	// Open the local file
	file, err := os.Open(localFilePath)
	if err != nil {
		return fmt.Errorf("failed to open local file: %w", err)
	}
	defer file.Close()

	// Upload the file to the FTP server
	err = conn.Stor(remoteFilePath, file)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	return nil
}

func downloadFile(conn *ftp.ServerConn, remoteFilePath, localFilePath string) error {
	// Download the file from the FTP server
	response, err := conn.Retr(remoteFilePath)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}
	defer response.Close()

	// Create the local file
	file, err := os.Create(localFilePath)
	if err != nil {
		return fmt.Errorf("failed to create local file: %w", err)
	}
	defer file.Close()

	// Copy the response data to the local file
	_, err = io.Copy(file, response)
	if err != nil {
		return fmt.Errorf("failed to copy file data: %w", err)
	}

	return nil
}
