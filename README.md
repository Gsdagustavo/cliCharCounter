---

# cliCharCounter

**cliCharCounter** is a simple command-line tool written in Go that counts the occurrences of a specific letter in a
text file.

---

## **Features**

* Read any text file.
* Count the occurrences of a single letter.
* Works with Unicode characters.
* Simple CLI usage with flags.

---

## **Installation**

1. Make sure you have [Go](https://golang.org/dl/) installed.
2. Clone this repository:

```bash
git clone <your-repo-url>
cd cliCharCounter
```

---

## **Usage**

Run the program with the `--file` and `--letter` flags:

```bash
go run . --file path/to/file.txt --letter a
```

**Example output:**

```
Occurrences of the letter a in the file: 12
```

**Notes:**

* The `--letter` flag must be a single character.
* The `--file` flag must point to a valid text file.

---

## **Code Overview**

* `main.go`

    * Parses CLI flags (`--file` and `--letter`).
    * Opens the file and reads its contents.
    * Counts the occurrences of the given letter.
    * Prints the result.
* `readFile` function: Reads the content of a file and returns it as a string.
* `countLetter` function: Counts occurrences of a specific rune in a string.

---

## **License**

This project is open-source and free to use.

---
