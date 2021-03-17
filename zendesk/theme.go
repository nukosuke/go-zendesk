package zendesk

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// Theme is information about Zendesk Guide theme
type Theme struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Version   string    `json:"version"`
	Author    string    `json:"author"`
	Live      bool      `json:"live"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	BrandID   string    `json:"brand_id,omitempty"`
}

// ThemeListOptions is parameters used of GetThemes
type ThemeListOptions struct {
	BrandID string `url:"brand_id,omitempty"`
}

type ThemeAPI interface {
	GetThemes(ctx context.Context, opts *ThemeListOptions) ([]Theme, error)
	GetTheme(ctx context.Context, themeID string) (Theme, error)
}

func (z *Client) GetThemes(ctx context.Context, opts *ThemeListOptions) ([]Theme, error) {
	var data struct {
		Themes []Theme `json:"themes"`
	}

	if opts == nil {
		opts = &ThemeListOptions{}
	}

	u, err := addOptions("/guide/theming/themes", opts)
	if err != nil {
		return []Theme{}, err
	}

	body, err := z.get(ctx, u)
	if err != nil {
		return []Theme{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return []Theme{}, err
	}

	return data.Themes, nil
}

func (z *Client) GetTheme(ctx context.Context, themeID string) (Theme, error) {
	var data struct {
		Theme Theme `json:"theme"`
	}

	u := fmt.Sprintf("/guide/theming/themes/%s", themeID)
	body, err := z.get(ctx, u)

	if err != nil {
		return Theme{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return Theme{}, err
	}

	return data.Theme, nil
}
