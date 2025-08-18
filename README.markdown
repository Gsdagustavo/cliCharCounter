# cliCharCounter

**cliCharCounter** is a lightweight, efficient command-line tool written in Go to count occurrences of a specific letter (or Unicode character) in a text file. It’s simple to use, fast, and perfect for quick text analysis.

---

## **Features**

- Reads any text file (UTF-8 encoded).
- Counts occurrences of a single character (supports Unicode).
- User-friendly CLI with intuitive flags.
- Lightweight with minimal dependencies.
- Cross-platform compatibility (Windows, macOS, Linux).

---

## **Installation**

1. Ensure [Go](https://go.dev/dl/) (version 1.16 or later) is installed on your system.
2. Clone the repository:

   ```bash
   git clone <your-repo-url>
   cd cliCharCounter
   ```

3. Build the executable (optional, for standalone use):

   ```bash
   go build -o cliCharCounter
   ```

---

## **Usage**

Run the program using the `--file` and `--letter` flags:

```bash
go run . --file path/to/file.txt --letter a
```

Or, if built:

```bash
./cliCharCounter --file path/to/file.txt --letter a
```

### **Example Output**

```
Occurrences of the letter 'a' in the file: 12
```

### **Flags**

- `--file`: Path to the input text file (required).
- `--letter`: Single character to count (required, case-sensitive).

### **Notes**

- The `--letter` flag accepts a single character, including Unicode characters (e.g., `é`, `π`).
- The `--file` flag must point to a valid, readable text file.
- Use `-h` or `--help` for usage information:

  ```bash
  go run . --help
  ```

---

## **Code Structure**

- **`main.go`**: Core logic for parsing CLI flags, reading the file, counting characters, and displaying results.
- **`readFile`**: Reads the file content and returns it as a string, handling errors gracefully.
- **`countLetter`**: Counts occurrences of a specified rune in the text, supporting Unicode characters.

---

## **Example**

Given a file `sample.txt` with the content:

```
banana apple orange
```

Running:

```bash
go run . --file sample.txt --letter a
```

Outputs:

```
Occurrences of the letter 'a' in the file: 5
```

---

## **Contributing**

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-name`).
3. Make your changes and commit (`git commit -m "Add feature"`).
4. Push to the branch (`git push origin feature-name`).
5. Open a pull request.

Please ensure your code follows Go best practices and includes tests where applicable.

---

## **License**

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## **Contact**

For questions, suggestions, or issues, please open an issue on the [GitHub repository](<your-repo-url>).