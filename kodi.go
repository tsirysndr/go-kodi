package kodi

import "github.com/dghubble/sling"

type Client struct {
	base         *sling.Sling
	common       service
	Addons       *AddonsService
	Application  *ApplicationService
	AudioLibrary *AudioLibraryService
	Favourites   *FavouritesService
	Files        *FilesService
	Gui          *GuiService
	Input        *InputService
	Player       *PlayerService
	Playlist     *PlaylistService
	Profiles     *ProfilesService
	Pvr          *PvrService
	Settings     *SettingsService
	System       *SystemService
	Textures     *TexturesService
	VideoLibrary *VideoLibraryService
	Xbmc         *XbmcService
}

type service struct {
	client *Client
}

func NewClient(baseURL string) *Client {
	base := sling.New().Base(baseURL).Set("Accept", "application/json")
	c := &Client{}
	c.base = base
	c.common.client = c
	c.Addons = (*AddonsService)(&c.common)
	c.Application = (*ApplicationService)(&c.common)
	c.AudioLibrary = (*AudioLibraryService)(&c.common)
	c.Favourites = (*FavouritesService)(&c.common)
	c.Files = (*FilesService)(&c.common)
	c.Gui = (*GuiService)(&c.common)
	c.Input = (*InputService)(&c.common)
	c.Player = (*PlayerService)(&c.common)
	c.Playlist = (*PlaylistService)(&c.common)
	c.Profiles = (*ProfilesService)(&c.common)
	c.Pvr = (*PvrService)(&c.common)
	c.Settings = (*SettingsService)(&c.common)
	c.System = (*SystemService)(&c.common)
	c.Textures = (*TexturesService)(&c.common)
	c.VideoLibrary = (*VideoLibraryService)(&c.common)
	c.Xbmc = (*XbmcService)(&c.common)
	return c
}
