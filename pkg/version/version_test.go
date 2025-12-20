package version_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/sangrita-tech/golang-pkg/pkg/version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func withSavedGlobals(t *testing.T, fn func()) {
	t.Helper()

	oldV, oldC, oldD := version.Version, version.Commit, version.Date
	t.Cleanup(func() {
		version.Version, version.Commit, version.Date = oldV, oldC, oldD
	})

	fn()
}

func Test_GetInfo_WhenDateUnknown_ReturnsNilDate(t *testing.T) {
	withSavedGlobals(t, func() {
		version.Version = "v1.2.3"
		version.Commit = "abc123"
		version.Date = "unknown"

		info := version.GetInfo()

		assert.Equal(t, "v1.2.3", info.Version)
		require.NotNil(t, info.Commit)
		assert.Equal(t, "abc123", *info.Commit)
		assert.Nil(t, info.Date)
	})
}

func Test_GetInfo_WhenDateEmpty_ReturnsNilDate(t *testing.T) {
	withSavedGlobals(t, func() {
		version.Date = ""

		info := version.GetInfo()

		assert.Nil(t, info.Date)
	})
}

func Test_GetInfo_WhenDateInvalid_ReturnsNilDate(t *testing.T) {
	withSavedGlobals(t, func() {
		version.Date = "not-a-date"

		info := version.GetInfo()

		assert.Nil(t, info.Date)
	})
}

func Test_GetInfo_WithValidRFC3339Date_ParsesDate(t *testing.T) {
	withSavedGlobals(t, func() {
		version.Date = "2025-12-20T12:34:56Z"
		want, err := time.Parse(time.RFC3339, version.Date)
		require.NoError(t, err)

		info := version.GetInfo()

		require.NotNil(t, info.Date)
		assert.True(t, info.Date.Equal(want), "expected %v, got %v", want, *info.Date)
	})
}

func Test_GetInfo_WhenDateNotParsed_JSONContainsNull(t *testing.T) {
	withSavedGlobals(t, func() {
		version.Version = "v0.1.0"
		version.Commit = "deadbeef"
		version.Date = "unknown"

		info := version.GetInfo()
		b, err := json.Marshal(info)
		require.NoError(t, err)

		var m map[string]any
		require.NoError(t, json.Unmarshal(b, &m))

		_, ok := m["date"]
		require.True(t, ok, "expected 'date' key to exist in JSON")
		assert.Nil(t, m["date"])

		_, ok = m["commit"]
		require.True(t, ok, "expected 'commit' key to exist in JSON")
		assert.NotNil(t, m["commit"])
		assert.Equal(t, "deadbeef", m["commit"])
	})
}

func Test_GetInfo_WhenCommitUnknown_JSONContainsNullCommit(t *testing.T) {
	withSavedGlobals(t, func() {
		version.Version = "v0.1.0"
		version.Commit = "unknown"
		version.Date = "unknown"

		info := version.GetInfo()
		b, err := json.Marshal(info)
		require.NoError(t, err)

		var m map[string]any
		require.NoError(t, json.Unmarshal(b, &m))

		_, ok := m["commit"]
		require.True(t, ok, "expected 'commit' key to exist in JSON")
		assert.Nil(t, m["commit"])
	})
}
