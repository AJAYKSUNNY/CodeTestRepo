package main

import (
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	// Save original os.Args
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Mock os.Args with custom arguments
	os.Args = []string{"test", "reviews_test.json", "movies_test.json"}

	main()
}

func Test_getStarRatingFromScore(t *testing.T) {
	type args struct {
		score float64
		max   float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Get Star Rating From Score Test Case 1",
			args: args{
				score: 0,
				max:   100,
			},
			want: "",
		},
		{
			name: "Get Star Rating From Score Test Case 2",
			args: args{
				score: 30,
				max:   100,
			},
			want: "*½",
		},
		{
			name: "Get Star Rating From Score Test Case 3",
			args: args{
				score: 57,
				max:   100,
			},
			want: "***",
		},
		{
			name: "Get Star Rating From Score Test Case 4",
			args: args{
				score: 93,
				max:   100,
			},
			want: "*****",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStarRatingFromScore(tt.args.score, tt.args.max); got != tt.want {
				t.Errorf("getStarRatingFromScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_composeTweet(t *testing.T) {
	type args struct {
		movieTitle string
		year       string
		review     string
		score      float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Compose Tweet Test case 1",
			args: args{
				movieTitle: "The Matrix",
				year:       "",
				review:     `With mind-bending visual effects and iconic action sequences, "The Matrix" delivers a thought-provoking exploration of reality, identity, and freedom.`,
				score:      100,
			},
			want: `The Matrix: With mind-bending visual effects and iconic action sequences, "The Matrix" delivers a thought-provoking exploration of rea *****`,
		},
		{
			name: "Compose Tweet Test case 2",
			args: args{
				movieTitle: "Avatar",
				year:       "(2019)",
				review:     `With groundbreaking visual effects and a captivating narrative, "Avatar" explores themes of environmentalism, colonization, and the clash of civilizations.`,
				score:      84,
			},
			want: `Avatar(2019): With groundbreaking visual effects and a captivating narrative, "Avatar" explores themes of environmentalism, colonizat ****½`,
		},
		{
			name: "Compose Tweet Test case 3",
			args: args{
				movieTitle: "Dr. Strangelove or How I Learned to Stop Worrying and Love the Bomb",
				year:       "(2019)",
				review:     `Dr. Strangelove brilliantly satirizes the absurdity of nuclear warfare. Kubrick's dark humor and sharp wit make this film a timeless classic. A must-watch for its insightful commentary on the Cold War era`,
				score:      78,
			},
			want: `Dr. Strangelove or How I (2019): Dr. Strangelove brilliantly satirizes the absurdity of nuclear warfare. Kubrick's dark humor and sharp ****`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := composeTweet(tt.args.movieTitle, tt.args.year, tt.args.review, tt.args.score); got != tt.want {
				t.Errorf("composeTweet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadJSON(t *testing.T) {
	type args struct {
		filePath string
		v        interface{}
	}
	var movie []Movie
	var review []Review
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Load Json Test Case Movies File",
			args: args{
				filePath: "movies_test.json",
				v:        &movie,
			},
			wantErr: false,
		},
		{
			name: "Load Json Test Case Reviews File",
			args: args{
				filePath: "reviews_test.json",
				v:        &review,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := loadJSON(tt.args.filePath, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("loadJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
