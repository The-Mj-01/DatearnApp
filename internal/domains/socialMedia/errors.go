package socialMedia

import "errors"

var (
	SocialMediaNotFound error = errors.New("request social media does not found")
	NameNotFound        error = errors.New("request social media name does not found")
)
