package smallbiz

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const BASE_URL = "http://api.sba.gov/geodata/" // countries/chn;bra/indicators/DPANUSIFS?date=2009Q1:2010Q3
type Area string
type Query string
type Response string
type Series struct {
	City,
	County,
	State Area
	Format string
}

func (s *Series) OneCity() Response {
	s = s.Querify()
	q := Query(BASE_URL + "all_data_for_city_of/" + string(s.City) + "/" + string(s.State) + "." + string(s.Format))
	return q.Request()
}

func (s *Series) AllCities() Response {
	s = s.Querify()
	q := Query(BASE_URL + "city_data_for_state_of/" + string(s.State) + "." + string(s.Format))
	return q.Request()
}

func (s *Series) OneCounty() Response {
	s = s.Querify()
	q := Query(BASE_URL + "all_data_for_county_of/" + string(s.County) + "/" + string(s.State) + "." + string(s.Format))
	return q.Request()
}

func (s *Series) AllCounties() Response {
	s = s.Querify()
	q := Query(BASE_URL + "county_data_for_state_of/" + string(s.State) + "." + string(s.Format))
	return q.Request()
}

func (s *Series) AllState() Response {
	s = s.Querify()
	q := Query(BASE_URL + "city_county_data_for_state_of/" + string(s.State) + "." + string(s.Format))
	return q.Request()
}

func (q Query) Request() Response {
	resp, err := http.Get(string(q))
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Println(string(body))
	return Response(body)
}

func (s *Series) Querify() *Series {
	s.City.Pathize()
	s.County.Pathize()
	s.State.Pathize()
	s.Format = strings.ToLower(s.Format)
	return s
}

func (a *Area) Pathize() {
	*a = Area(strings.Replace(strings.ToLower(string(*a)), " ", "%20", -1))
}

// Write Data to File (XML or JSON)
func (r Response) Write(filepath string) {
	err := ioutil.WriteFile(filepath, []byte(r), 0777)
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
