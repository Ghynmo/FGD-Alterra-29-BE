package newsapi

import (
	"context"
	"encoding/json"
	apinews "fgd-alterra-29/business/api_news"
	"io/ioutil"
	"net/http"
)

type NewsAPIRepository struct {
	httpClient http.Client
}

func NewAPINewsRepository(httpClient http.Client) apinews.Repository {
	return &NewsAPIRepository{
		httpClient: http.Client{},
	}
}

func (newsAR *NewsAPIRepository) GetAPINews(ctx context.Context, apikey string) (apinews.Domain, error) {

	if apikey == "" {
		apikey = "8be9f10a83634066b8fcaa84908a8802"
	}

	link := "https://newsapi.org/v2/top-headlines?country=id&apiKey=" + apikey
	response, err_api := http.Get(link)

	if err_api != nil {
		return apinews.Domain{}, err_api
	}

	responseData, err_read := ioutil.ReadAll(response.Body)

	if err_read != nil {
		return apinews.Domain{}, err_read
	}

	defer response.Body.Close()
	var data ArticlesParent
	json.Unmarshal(responseData, &data)

	return data.ToDomain(), nil
}
