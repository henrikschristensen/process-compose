package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rs/zerolog/log"
)

func (p *PcClient) startNamespace(name string) error {
	url := fmt.Sprintf("http://%s/namespace/start/%s", p.address, url.PathEscape(name))
	resp, err := p.client.Post(url, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	var respErr pcError
	if err = json.NewDecoder(resp.Body).Decode(&respErr); err != nil {
		log.Error().Msgf("failed to decode start namespace %s response: %v", name, err)
		return err
	}
	return errors.New(respErr.Error)
}

func (p *PcClient) stopNamespace(name string) error {
	url := fmt.Sprintf("http://%s/namespace/stop/%s", p.address, url.PathEscape(name))
	resp, err := p.client.Post(url, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	var respErr pcError
	if err = json.NewDecoder(resp.Body).Decode(&respErr); err != nil {
		log.Error().Msgf("failed to decode stop namespace %s response: %v", name, err)
		return err
	}
	return errors.New(respErr.Error)
}

func (p *PcClient) restartNamespace(name string) error {
	url := fmt.Sprintf("http://%s/namespace/restart/%s", p.address, url.PathEscape(name))
	resp, err := p.client.Post(url, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	var respErr pcError
	if err = json.NewDecoder(resp.Body).Decode(&respErr); err != nil {
		log.Error().Msgf("failed to decode restart namespace %s response: %v", name, err)
		return err
	}
	return errors.New(respErr.Error)
}

func (p *PcClient) getNamespaces() ([]string, error) {
	url := fmt.Sprintf("http://%s/namespaces", p.address)
	resp, err := p.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respErr pcError
		if err = json.NewDecoder(resp.Body).Decode(&respErr); err != nil {
			log.Error().Msgf("failed to decode get namespaces response: %v", err)
			return nil, err
		}
		return nil, errors.New(respErr.Error)
	}

	var namespaces []string
	if err = json.NewDecoder(resp.Body).Decode(&namespaces); err != nil {
		return nil, err
	}
	return namespaces, nil
}
