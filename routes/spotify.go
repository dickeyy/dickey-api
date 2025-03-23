package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
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

// HandleSpotifyWebSocket handles WebSocket connections for real-time track updates
func HandleSpotifyWebSocket(c *websocket.Conn) {
	// Get the user from query parameters
	user := c.Query("user")
	if user == "" {
		c.WriteJSON(fiber.Map{
			"error": "User is required",
		})
		c.Close()
		return
	}

	// Create a channel to track if the client has disconnected
	done := make(chan bool)
	go func() {
		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				done <- true
				return
			}
		}
	}()

	// Track the last sent track to avoid duplicate messages
	var lastTrack *NowPlayingResponse

	// Start polling for track updates
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			track, err := getCurrentTrack(user)
			if err != nil {
				c.WriteJSON(fiber.Map{
					"error": err.Error(),
				})
				continue
			}

			// Only send if the track is different from the last one
			if lastTrack == nil || !tracksEqual(lastTrack, track) {
				if err := c.WriteJSON(track); err != nil {
					return
				}
				lastTrack = track
			}
		}
	}
}

// tracksEqual compares two NowPlayingResponse structs
func tracksEqual(a, b *NowPlayingResponse) bool {
	if a == nil || b == nil {
		return false
	}
	return a.Artist == b.Artist &&
		a.Title == b.Title &&
		a.Album == b.Album &&
		a.URL == b.URL &&
		a.IsPlaying == b.IsPlaying &&
		a.ImageURL == b.ImageURL
}

// getCurrentTrack fetches the current track from Last.fm
func getCurrentTrack(user string) (*NowPlayingResponse, error) {
	response, err := http.Get(constructLastFmAPIUrl("user.getrecenttracks", user))
	if err != nil {
		return nil, fmt.Errorf("failed to get current track: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var lastFmResponse LastFmResponse
	if err := json.Unmarshal(body, &lastFmResponse); err != nil {
		return nil, fmt.Errorf("failed to parse Last.fm response: %v", err)
	}

	recentTracks := lastFmResponse.RecentTracks.Track
	if len(recentTracks) == 0 {
		return nil, fmt.Errorf("no tracks found")
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
		return nil, fmt.Errorf("nothing is playing")
	}

	return &NowPlayingResponse{
		Artist:    track.Artist.Text,
		Title:     track.Name,
		Album:     track.Album.Text,
		URL:       track.URL,
		IsPlaying: isPlaying,
		ImageURL:  imageURL,
	}, nil
}

func GetCurrentTrack(c *fiber.Ctx) error {
	// get the user from the query params
	user := c.Query("user")
	if user == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User is required",
		})
	}

	track, err := getCurrentTrack(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(track)
}

func constructLastFmAPIUrl(method string, user string) string {
	return fmt.Sprintf("%s?method=%s&user=%s&api_key=%s&format=json&limit=1", LAST_FM_API, method, user, LAST_FM_API_KEY)
}
