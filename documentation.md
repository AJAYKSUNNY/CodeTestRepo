Overview:

    This command-line application, written in Go, reads movie review data and composes "tweets" based on the reviews. Each tweet follows a specific format and is constructed from the review data provided.

Pre-Requisites

    * Go Programming Language
        installed version - go1.20.6 darwin/arm64

    * JSON Files
        Require two json files in source project root i.e, review data (reviews.json) and movie information (movies.json) from where the data is read and used for processing tweet

    * Text Editor or IDE
        Need a text editor or an integrated development environment (IDE) to view and edit the Go source code files.I have used VS Code 

    * Terminal or Command Prompt
        Access to a terminal or command prompt to execute commands and run the Go application.


Features:

    Load JSON Data:
        The application loads review and movie data from JSON files provided as command-line arguments.
        It uses the loadJSON function to read the JSON files and unmarshal the data into corresponding struct types.

    Compose Tweets:
        Tweets are composed based on the review data retrieved.
        The composeTweet function constructs each tweet according to the specified format: Movie Title (Year): Review of the movie plus ratings
        The movie year is included if available, and the review score is converted into a star rating out of 5, using Unicode stars and half-star symbols.

    Limit Tweet Length:
        Tweets are limited to a maximum length of 140 characters.
        If a tweet exceeds the character limit, the movie title and review text are truncated appropriately to fit within the limit while preserving the format.

    Output Tweets:
        The application processes each review and outputs the corresponding tweet.
        Tweets are stored in a slice and printed to the console in JSON array format.

Usage:

    To run the application, execute the binary with two command-line arguments specifying the paths to the review and movie JSON files. For example:

    go mod init github.com/yourusername/myproject
    go run main.go reviews.json movies.json

    To run test files
        Install testing package
            go get github.com/stretchr/testify or run go mod vendor
            RUN the below command to run the test files,it should give an ok response
                go test
            To see the code coverage run , it will show the percentage of code covered in the package
                go test -cover ./...



My Approcah Towards the Task

    * I began by parsing the JSON files and storing their contents in corresponding structs for efficient data management and retrieval.

    * To facilitate easy access to movie years, I constructed a map using the movie titles as keys and the corresponding years as values, allowing me to quickly retrieve the year based on the movie name provided in the review struct.

    * Next, I iterated through the reviews, applying the necessary conditions to ensure the desired output format.

    * Utilizing the review scores provided as input, I computed the star ratings by mapping them to a five-star scale, assuming a maximum score of 100.

    * As I processed each review, I formatted the tweet according to the specified requirements and appended it to an array for storage.

    * Finally, I displayed the array containing the formatted tweets, presenting the desired output reflecting our processed reviews.

