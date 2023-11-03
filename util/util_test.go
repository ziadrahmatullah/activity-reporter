package util_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/application"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/util"
	"github.com/stretchr/testify/assert"
)

func TestProcessSocialGraph(t *testing.T) {
	t.Run("should add user in app when input Alice follows Bob", func(t *testing.T) {
		input := "Alice follows Bob"
		socialApp := application.NewSocialApp()
		util.ProcessSocialGraph(socialApp, input)

		user1, _ := socialApp.IsUserInApp("Alice")
		user2, _ := socialApp.IsUserInApp("Bob")

		assert.True(t, user1)
		assert.True(t, user2)
	})

}
