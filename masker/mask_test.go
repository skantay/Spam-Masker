package mask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpamMasker(t *testing.T) {
	testTable := []struct {
		got  string
		want string
	}{
		{
			got:  "Here's my spammy page: http://hehefouls.netHAHAHA see you.",
			want: "Here's my spammy page: http://******************* see you.",
		},
		{
			got:  "Check out this amazing offer at http://example.com! Limited time only!",
			want: "Check out this amazing offer at http://************ Limited time only!",
		},
		{
			got:  "Spam link without space:http://spam.com",
			want: "Spam link without space:http://********",
		},
		{
			got:  "http://spam.com spam link first",
			want: "http://******** spam link first",
		},
		{
			got:  "No spam links in this sentence.",
			want: "No spam links in this sentence.",
		},
	}

	for _, testCase := range testTable {

		t.Run(testCase.got, func(t *testing.T) {
			result := spamMasker(testCase.got)
			assert.Equal(t, testCase.want, result)
		})
	}
}
