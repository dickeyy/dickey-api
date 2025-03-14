package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Last.fm api keys are designed to be public so we can hard code the key here
var LAST_FM_API = "https://ws.audioscrobbler.com/2.0/"
var LAST_FM_API_KEY = "8f3b65a08dc19fd2fd52daee66e27a1f"

// LastFmResponse represents the response from Last.fm API
type LastFmResponse struct {
	RecentTracks struct {
		Track []struct {
			Artist struct {
				Text string `json:"#text"`
			} `json:"artist"`
			Name  string `json:"name"`
			URL   string `json:"url"`
			Album struct {
				Text string `json:"#text"`
			} `json:"album"`
			Image []struct {
				Text string `json:"#text"`
				Size string `json:"size"`
			} `json:"image"`
			Attr struct {
				NowPlaying string `json:"nowplaying"`
			} `json:"@attr,omitempty"`
		} `json:"track"`
	} `json:"recenttracks"`
}

// NowPlayingResponse represents the response we'll send to the client
type NowPlayingResponse struct {
	Artist    string `json:"artist"`
	Title     string `json:"title"`
	Album     string `json:"album"`
	URL       string `json:"url"`
	IsPlaying bool   `json:"isPlaying"`
	ImageURL  string `json:"imageUrl"`
}

func GetCurrentTrack(c *fiber.Ctx) error {
	// get the user from the query params
	user := c.Query("user")
	if user == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User is required",
		})
	}

	response, err := http.Get(constructLastFmAPIUrl("user.getrecenttracks", user))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get current track",
		})
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read response body",
		})
	}

	var lastFmResponse LastFmResponse
	if err := json.Unmarshal(body, &lastFmResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse Last.fm response",
		})
	}

	recentTracks := lastFmResponse.RecentTracks.Track
	if len(recentTracks) == 0 {
		return c.JSON(fiber.Map{
			"error": "No tracks found",
		})
	}

	track := recentTracks[0]

	// Find the largest image URL (extralarge if available)
	imageURL := ""
	if len(track.Image) > 0 {
		imageURL = track.Image[0].Text // Default to first image
		for _, img := range track.Image {
			if img.Size == "extralarge" {
				imageURL = img.Text
				break
			}
		}
	}

	// Check if the track is currently playing
	isPlaying := track.Attr.NowPlaying == "true"

	if !isPlaying {
		return c.JSON(fiber.Map{
			"error": "Nothing is playing",
		})
	}

	nowPlaying := NowPlayingResponse{
		Artist:    track.Artist.Text,
		Title:     track.Name,
		Album:     track.Album.Text,
		URL:       track.URL,
		IsPlaying: isPlaying,
		ImageURL:  imageURL,
	}

	// Return both the formatted response and the raw Last.fm data
	return c.JSON(nowPlaying)
}

func constructLastFmAPIUrl(method string, user string) string {
	return fmt.Sprintf("%s?method=%s&user=%s&api_key=%s&format=json&limit=1", LAST_FM_API, method, user, LAST_FM_API_KEY)
}
