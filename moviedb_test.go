package moviedb

import (
	"bytes"
	"reflect"
	"testing"
)

const RAW_HTML = `
<img class="poster fade lazyautosizes lazyloaded" data-sizes="auto"
    data-src="https://image.tmdb.org/t/p/w185_and_h278_bestv2/9E2y5Q7WlCVNEhP5GiVTjhEhx1o.jpg"
    data-srcset="https://image.tmdb.org/t/p/w185_and_h278_bestv2/9E2y5Q7WlCVNEhP5GiVTjhEhx1o.jpg 1x, https://image.tmdb.org/t/p/w370_and_h556_bestv2/9E2y5Q7WlCVNEhP5GiVTjhEhx1o.jpg 2x"
    alt="It" sizes="185px"
    srcset="https://image.tmdb.org/t/p/w185_and_h278_bestv2/9E2y5Q7WlCVNEhP5GiVTjhEhx1o.jpg 1x, https://image.tmdb.org/t/p/w370_and_h556_bestv2/9E2y5Q7WlCVNEhP5GiVTjhEhx1o.jpg 2x"
    src="https://image.tmdb.org/t/p/w185_and_h278_bestv2/9E2y5Q7WlCVNEhP5GiVTjhEhx1o.jpg">
<img class="poster fade lazyautosizes lazyloaded" data-sizes="auto"
    data-src="https://image.tmdb.org/t/p/w185_and_h278_bestv2/aMpyrCizvSdc0UIMblJ1srVgAEF.jpg"
    data-srcset="https://image.tmdb.org/t/p/w185_and_h278_bestv2/aMpyrCizvSdc0UIMblJ1srVgAEF.jpg 1x, https://image.tmdb.org/t/p/w370_and_h556_bestv2/aMpyrCizvSdc0UIMblJ1srVgAEF.jpg 2x"
    alt="Blade Runner 2049" sizes="185px"
    srcset="https://image.tmdb.org/t/p/w185_and_h278_bestv2/aMpyrCizvSdc0UIMblJ1srVgAEF.jpg 1x, https://image.tmdb.org/t/p/w370_and_h556_bestv2/aMpyrCizvSdc0UIMblJ1srVgAEF.jpg 2x"
    src="https://image.tmdb.org/t/p/w185_and_h278_bestv2/aMpyrCizvSdc0UIMblJ1srVgAEF.jpg">
<img class="poster fade lazyautosizes lazyloaded" data-sizes="auto"
     data-src="https://image.tmdb.org/t/p/w185_and_h278_bestv2/8dTWj3c74RDhXfSUZpuyVNJrgS.jpg"
     data-srcset="https://image.tmdb.org/t/p/w185_and_h278_bestv2/8dTWj3c74RDhXfSUZpuyVNJrgS.jpg 1x, https://image.tmdb.org/t/p/w370_and_h556_bestv2/8dTWj3c74RDhXfSUZpuyVNJrgS.jpg 2x"
     alt="American Made" sizes="103px"
     srcset="https://image.tmdb.org/t/p/w185_and_h278_bestv2/8dTWj3c74RDhXfSUZpuyVNJrgS.jpg 1x, https://image.tmdb.org/t/p/w370_and_h556_bestv2/8dTWj3c74RDhXfSUZpuyVNJrgS.jpg 2x"
     src="https://image.tmdb.org/t/p/w185_and_h278_bestv2/8dTWj3c74RDhXfSUZpuyVNJrgS.jpg">
`

func TestCrawler(t *testing.T) {
	expectedShows := []Show{
		Show{
			Title: "It",
			Cover: "https://image.tmdb.org/t/p/w185_and_h278_bestv2/9E2y5Q7WlCVNEhP5GiVTjhEhx1o.jpg",
		},
		Show{
			Title: "Blade Runner 2049",
			Cover: "https://image.tmdb.org/t/p/w185_and_h278_bestv2/aMpyrCizvSdc0UIMblJ1srVgAEF.jpg",
		},
		Show{
			Title: "American Made",
			Cover: "https://image.tmdb.org/t/p/w185_and_h278_bestv2/8dTWj3c74RDhXfSUZpuyVNJrgS.jpg",
		},
	}

	shows := crawl(bytes.NewBufferString(RAW_HTML))

	for i, show := range shows {
		if !reflect.DeepEqual(show, expectedShows[i]) {
			t.Error(
				"expected", expectedShows[i],
				"got", show,
			)
		}
	}
}
