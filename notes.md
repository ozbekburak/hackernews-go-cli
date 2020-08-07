
// FormattedStory formats our response to show meaningful output to user
func (i Item) FormattedStory() string {
	story := fmt.Sprintf(
		"\nTime: %v\nBy: %s\nID: %d\nComment Count: %d\nScore: %d\nTitle: %s\nURL: %s\n",
		i.FormattedTime(i.Time), i.By, i.ID, i.Descendants, i.Score, i.Title, i.URL)
	return story
}

// FormattedPoll formats our response to show meaningful output to user
func (i Item) FormattedPoll() string {
	poll := fmt.Sprintf(
		"\nTime: %v\nBy: %s\nComment Count: %d\nID: %d\nScore: %d\nText: %s\nTitle: %s\n",
		i.FormattedTime(i.Time), i.By, i.Descendants, i.ID, i.Score, html.UnescapeString(i.Text), i.Title)
	return poll
}

// FormattedJob formats our response to show meaningful output to user
func (i Item) FormattedJob() string {
	job := fmt.Sprintf(
		"\nTime: %v\nBy: %s\nID: %d\nScore: %d\nText: %s\nTitle: %s\nURL: %s\n",
		i.FormattedTime(i.Time), i.By, i.ID, i.Score, html.UnescapeString(i.Text), i.Title, i.URL)
	return job
}

// FormattedComment formats our response to show meaningful output to user
func (i Item) FormattedComment() string {
	comment := fmt.Sprintf(
		"\nTime: %v\nBy: %s\nID: %d\nParentID: %d\nText: %s\n",
		i.FormattedTime(i.Time), i.By, i.ID, i.Parent, html.UnescapeString(i.Text))
	return comment
}

// FormattedAsk formats our response to show meaningful output to user
func (i Item) FormattedAsk() string {
	ask := fmt.Sprintf(
		"\nTime: %v\nBy: %s\nComment Count: %d\nID: %d\nScore: %d\nText: %s\nTitle: %s\n",
		i.FormattedTime(i.Time), i.By, i.Descendants, i.ID, i.Score, html.UnescapeString(i.Text), i.Title)
	return ask