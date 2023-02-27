# OpenEarth

OpenEarth is the API to discover and explore all open resources about Earth Observation. OpenEarth collects and indexes digital assets on Earth from open data sources and makes them available through a GraphQL API. 

So far the groups of resources are:

- Datasets 
  - [NASA](https://www.earthdata.nasa.gov/)
  - [ESA](https://www.esa.int/)
- Outer Space Objects
  - [United NationsOffice for Outer Space Affairs](https://www.unoosa.org)
- Nasa EarthData
- [Nasa RSS feeds](https://www.nasa.gov/content/nasa-rss-feeds/)

## Playground

**Endpoint:** [OpenEarth GraphQL Playground](https://openearth.vercel.app/api/playground)

```graphql endpoint
https://openearth.vercel.app/api/playground
```

## Query examples:

```graphql
query($category: String, $topic_name: String, $term: String) {
  outer_space_objects(term: $term){
    name
    country
    function
    status
  }
  datasets(category: $category) {
    category
    description  
  }
  nasa_earthdata(topicName: $topic_name) {
    url
    topics {
      name
      url
      rss
      subtopics {	
        name
      }
    }
  }
  nasa_rss_feeds {
    name
    feed_type
  }
}
```

## Response example

```json
{
    "data": {
        "outer_space_objects": [
            {
                "name": "(CAROLINE) - Nusat 10 ",
                "country": "Uruguay",
                "function": "Earth observation",
                "status": "in orbit"
            }
        ],
        "datasets": [
            {
                "category": "EARTH SCIENCE",
                "status": "PLANNED",
                "provider": "AMD_USAPDC",
                "distributors": [
                    "United States Antarctic Program Data Center (USAP-DC)",
                    "Biological and Chemical Oceanography Data Management Office (BCO-DMO)"
                ],
                "created_at": "2020-08-31T00:00:00.000Z",
                "start_at": "2018-07-11T00:00:00.000Z",
                "description": "\"The Omnivores Dilemma\": The Effect of Autumn Diet on Winter Physiology and Condition of Juvenile Antarctic Krill - Antarctic krill are essential in the Southern Ocean as they support vast numbers of marine mammals, seabirds and fishes, some of which feed almost exclusively on krill. Antarctic krill also constitute a target species for industrial fisheries in the Southern Ocean. The success of Antarctic krill populations is largely determined by the ability of their young to survive the long, dark winter, where food is extremely scarce. To survive the long-dark winter, young Antarctic krill must have a high-quality diet in autumn. However, warming in certain parts of Antarctica is changing the dynamics and quality of the polar food web, resulting in a shift in the type of food available to young krill in autumn. It is not yet clear how these dynamic changes are affecting the ability of krill to survive the winter. This project aims to fill an important gap in current knowledge on an understudied stage of the Antarctic krill life cycle, the 1-year old juveniles. The results derived from this work will contribute to the development of improved bioenergetic, population and ecosystem models, and will advance current scientific understanding of this critical Antarctic species. This CAREER project's core education and outreach objectives seek to enhance education and increase diversity within STEM fields. An undergraduate course will be developed that will integrate undergraduate research and writing in way that promotes authentic scientific inquiry and analysis of original research data by the students, and that enhances their communication skills. A graduate course will be developed that will promote students' skills in communicating their own research to a non-scientific audience.  Graduate students will be supported through the proposed study and will gain valuable research experience. Traditionally underserved undergraduate students will be recruited to conduct independent research under the umbrella of the larger project. Throughout each field season, the research team will maintain a weekly blog that will include short videos, photographs and text highlighting the research, as well as their experiences living and working in Antarctica. The aim of the blog will be to engage the public and increase awareness and understanding of Antarctic ecosystems and the impact of warming, and of the scientific process of research and discovery.\r\n\r\n\r\n\r\nIn this 5-year CAREER project, the investigator will use a combination of empirical and theoretical techniques to assess the effects of diet on 1-year old krill in autumn-winter. The research is centered on four hypotheses: (H1) autumn diet affects 1-year old krill physiology and condition at the onset of winter; (H2) autumn diet has an effect on winter physiology and condition of 1-year old krill under variable winter food conditions; (H3) the rate of change in physiology and condition of 1-year old krill from autumn to winter is dependent on autumn diet; and (H4) the winter energy budget of 1-year old krill will vary between years and will be dependent on autumn diet. Long-term feeding experiments and in situ sampling will be used to measure changes in the physiology and condition of krill in relation to their diet and feeding environment. Empirically-derived data will be used to develop theoretical models of growth rates and energy budgets to determine how diet will influence the overwinter survival of 1-year old krill. The research will be integrated with an education and outreach plan to (1) develop engaging undergraduate and graduate courses, (2) train and develop young scientists for careers in polar research, and (3) engage the public and increase their awareness and understanding.\r\n\r\n\r\n\r\nThis award reflects NSF's statutory mission and has been deemed worthy of support through evaluation using the Foundation's intellectual merit and broader impacts review criteria."
            }
        ],
        "nasa_earthdata": {
            "url": "https://www.earthdata.nasa.gov/topics",
            "topics": [
                {
                    "name": "Atmosphere",
                    "url": "https://www.earthdata.nasa.gov/topics/atmosphere",
                    "rss": "https://www.earthdata.nasa.gov/topics/rss/Atmosphere",
                    "subtopics": [
                        {
                            "name": "Aerosols"
                        },
                        {
                            "name": "Air Quality"
                        }
                    ]
                }
            ]
        },
        "nasa_rss_feeds": [
            {
                "name": "Space Station News",
                "url": "https://www.nasa.gov/rss/dyn/shuttle_station.rss",
                "feed_type": "Topical News Feeds"
            }
        ]
    }
}
```
