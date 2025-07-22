
# 👻 DeepSearch

DeepSearch is an LLM (Large Language Model) powered search engine application that allows users to search for information from various sources. This project was created using a backend developed with the **GoFiber** framework, a **Svelte**-based frontend, and a **PostgreSQL** database. It also offers a rich user interface with modern tools like Tailwind CSS.

---

## 🚀 Features

- **LLM-Powered Summarization and Analysis**: Summarizes and analyzes search results.
- **Multiple Search Engine Support**: Fetches data from search engines like Google, Yandex, and Bing.
- **Sqlite Database**: Used to store and manage search results.
- **Fiber Framework**: A fast and scalable backend.
- **Svelte Frontend**: A user-friendly and high-performance interface.
- **Tailwind CSS**: A modern and stylish design.



## 🛠️ Installation

### Requirements

- **Go** (v1.19+)
- **Node.js** (v16+)
- **PostgreSQL** (v13+)

### Steps

1.  **Clone the Repository**:
    ```bash
    git clone https://github.com/username/deepsearch.git
    cd deepsearch
    ```

2.  **Install Backend Dependencies**:
    ```bash
    go mod tidy
    ```

3.  **Install Frontend Dependencies**:
    ```bash
    cd web
    npm install
    ```

4.  **Set Up the Database**:
    Create a database in PostgreSQL and update the `dsn` value in the `config/server.ini` file:
    ```ini
    [db]
    dsn="postgresql://user:password@localhost:5432/deepsearch"
    ```

5.  **Set Up Search Engine and LLM API Keys**:
    Fill in the `key` and `gemini` fields in the `config/search.ini` file:
    ```ini
    [serpapi]
    key = "YOUR_SERPAPI_KEY"

    [ai]
    gemini = "YOUR_LLM_API_KEY"
    ```

6.  **Run the Backend**:
    ```bash
    go run main.go
    ```

7.  **Run the Frontend**:
    ```bash
    cd web
    npm run dev
    ```

8.  **Open the Application**:
    Go to `http://localhost:3000` in your browser.


## 🔧 Configuration

### `server.ini`
Contains API and database settings:
```ini
[api]
port = ":3000"

[db]
dsn = "postgresql://user:password@localhost:5432/deepsearch"
```

### `search.ini`
Contains search engine and LLM settings:
```ini
[serpapi]
key = "YOUR_SERPAPI_KEY"
google = true
yandex = true
bing = true

[ai]
gemini = "YOUR_GEMINI_API_KEY"
prompt = "Based on the data here, I want you to extract a summary and analyze it..."
```


## 📜 Usage

1.  **Perform a Search**: Enter a query on the main page and click the "Search" button.
2.  **View the Results**: The search results are displayed in a summarized and analyzed form.
3.  **Error Messages**: If an error occurs, an appropriate message is shown to the user.


## 🤝 Contributing

If you want to contribute, please submit a **pull request** or open an **issue**. All contributions are welcome!


## 🌟 Support

If you like this project, please support it by giving it a ⭐! 😊
