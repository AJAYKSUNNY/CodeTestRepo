package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const maxScore = 100
const maxTweetLength = 140
const movieTitleLength = 25
const addedCharLength = 3

// Review represents a review object
type Review struct {
	MovieID int     `json:"movie_id"`
	Title   string  `json:"title"`
	Review  string  `json:"review"`
	Score   float64 `json:"score"`
}

// Movie represents a movie object
type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year,omitempty"`
}

// This function reads the values from the file and unmarshalls it into its corresponding struct
func loadJSON(filePath string, v interface{}) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, v)
	if err != nil {
		return err
	}
	return nil
}

func getStarRatingFromScore(score, max float64) string {

	stars := (score) / (max) * 5
	fullStars := int(stars)
	remainder := stars - float64(fullStars)

	var result string
	for i := 0; i < fullStars; i++ {
		result += "*"
	}
	if remainder > 0.5 {
		result += "*"
	}
	if remainder > 0.0 && remainder <= 0.5 {
		result += "½"
	}
	return result
}

func composeTweet(movieTitle, year, review string, score float64) string {

	stars := getStarRatingFromScore(score, maxScore) // assuming 100 is the maximum score
	tweet := fmt.Sprintf("%s%s: %s %s", movieTitle, year, review, stars)

	if len(tweet) > maxTweetLength {
		for len(tweet) > maxTweetLength {
			if len(movieTitle) > movieTitleLength {
				movieTitle = movieTitle[:movieTitleLength]
			} else {
				review = review[:maxTweetLength-len(movieTitle)-len(year)-len(stars)-addedCharLength] // addedCharLength is the extra char count used to format the output string (2 - space ,1 - :)
			}
			tweet = fmt.Sprintf("%s%s: %s %s", movieTitle, year, review, stars)
		}
	}
	return tweet
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go reviews.json movies.json")
		return
	}

	reviewsFile := os.Args[1]
	moviesFile := os.Args[2]

	// Load review and movie data
	var reviews []Review
	var movies []Movie
	var response []string
	if err := loadJSON(reviewsFile, &reviews); err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err := loadJSON(moviesFile, &movies); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a map of movie titles and years
	movieTitles := make(map[string]Movie)
	for _, movie := range movies {
		movieTitles[movie.Title] = movie
	}

	// Process reviews and output tweets
	for i, review := range reviews {
		movieYear := ""
		movie := movieTitles[review.Title]
		if movie.Year != 0 {
			movieYear = " (" + strconv.Itoa(movie.Year) + ")"
		}
		tweet := composeTweet(review.Title, movieYear, review.Review, review.Score)
		if i < len(reviews) {
			tweet = `"` + tweet + `"`
		}
		if i < len(reviews)-1 {
			tweet += ","
		}
		response = append(response, tweet)
	}
	fmt.Println(response)
}
