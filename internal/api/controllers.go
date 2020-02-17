package api

import (
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grissius/foxymoron/internal/core"
	"github.com/xanzy/go-gitlab"
)

func parseIsoDate(date string) time.Time {
	t, err := time.Parse(time.RFC3339, date)

	if err != nil {
		return time.Now()
	}
	return t
}

// List projects
// @Tags Projects
// @Summary List all available unarchived projects
// @Produce json
// @Success 200 {array} gitlab.Project
// @Router /projects [get]
// @Security ApiKey
// @Security GitLabURL
func getProjectsController(c *gin.Context) {
	projects := core.FetchProjects(c.MustGet("client").(*gitlab.Client))
	c.JSON(200, projects)
}

// List commits
// @Tags Commits
// @Summary List commit from all available projects within range
// @Produce json
// @Success 200 {array} gitlab.Commit
// @Router /commits [get]
// @Security ApiKey
// @Security GitLabURL
func getCommitsController(c *gin.Context) {
	from := parseIsoDate(c.Query("from"))
	to := parseIsoDate(c.Query("to"))
	message, _ := regexp.Compile(c.Query("message"))
	commits := core.FetchCommits(c.MustGet("client").(*gitlab.Client), &core.FetchCommitsOptions{From: &from, To: &to, WithStats: true, MessageRegex: message})
	c.JSON(200, commits)
}

// Commit statistics
// @Tags Statistics
// @Summary Get statistics for commits within range
// @Produce json
// @Success 200 {array} core.Stats
// @Router /statistics [get]
// @Security ApiKey
// @Security GitLabURL
func getStatisticsController(c *gin.Context) {
	from := parseIsoDate(c.Query("from"))
	to := parseIsoDate(c.Query("to"))
	stats := core.CommitsToStats(core.FetchCommits(c.MustGet("client").(*gitlab.Client), &core.FetchCommitsOptions{From: &from, To: &to, WithStats: true, MessageRegex: nil}))

	c.JSON(200, stats)
}
